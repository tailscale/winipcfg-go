/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"os/exec"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"
)

// I wish we didn't have to do this. netiohlp.dll (what's used by netsh.exe) has some nice tricks with writing directly
// to the registry and the nsi kernel object, but it's not clear copying those makes for a stable interface. WMI doesn't
// work with v6. CMI isn't in Windows 7.
func runNetsh(cmds []string) error {
	system32, err := windows.GetSystemDirectory()
	if err != nil {
		return err
	}
	cmd := exec.Command(system32 + "\\netsh.exe") // I wish we could append (, "-f", "CONIN$") but Go sets up the process context wrong.
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return errors.New(fmt.Sprintf("runNetsh stdin pipe - %v", err))
	}
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, strings.Join(append(cmds, "exit\r\n"), "\r\n"))
	}()
	output, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(fmt.Sprintf("runNetsh run - %v", err))
	}
	// Horrible kludges, sorry.
	cleaned := bytes.ReplaceAll(output, []byte("netsh>"), []byte{})
	cleaned = bytes.ReplaceAll(cleaned, []byte("There are no Domain Name Servers (DNS) configured on this computer."), []byte{})
	cleaned = bytes.TrimSpace(cleaned)
	if len(cleaned) != 0 {
		return errors.New(fmt.Sprintf("runNetsh returned error strings.\ninput:\n%s\noutput\n:%s",
			strings.Join(cmds, "\n"), bytes.ReplaceAll(output, []byte{'\r', '\n'}, []byte{'\n'})))
	}
	return nil
}

func flushDnsCmds(ifc *Interface) []string {
	return []string{
		fmt.Sprintf("interface ipv4 set dnsservers name=%d source=static address=none validate=no register=both", ifc.Index),
		fmt.Sprintf("interface ipv6 set dnsservers name=%d source=static address=none validate=no register=both", ifc.Ipv6IfIndex),
	}
}

func addDnsCmds(ifc *Interface, dnses []net.IP) []string {
	cmds := make([]string, len(dnses))
	j := 0
	for i := 0; i < len(dnses); i++ {
		if v4 := dnses[i].To4(); v4 != nil {
			cmds[j] = fmt.Sprintf("interface ipv4 add dnsservers name=%d address=%s validate=no", ifc.Index, v4.String())
		} else if v6 := dnses[i].To16(); v6 != nil {
			cmds[j] = fmt.Sprintf("interface ipv6 add dnsservers name=%d address=%s validate=no", ifc.Ipv6IfIndex, v6.String())
		} else {
			continue
		}
		j++
	}
	return cmds[:j]
}
