package winipcfg

import (
	"bytes"
	"fmt"
	"net"
	"time"
)

// https://docs.microsoft.com/en-us/windows/desktop/CIMWin32Prov/win32-networkadapterconfiguration
// Based on WMI Win32_NetworkAdapterConfiguration class.
type win32NetworkAdapterConfiguration struct {
	Caption                      string
	Description                  string
	SettingID                    string
	ArpAlwaysSourceRoute         bool
	ArpUseEtherSNAP              bool
	DatabasePath                 string
	DeadGWDetectEnabled          bool
	DefaultIPGateway             []string
	DefaultTOS                   uint8
	DefaultTTL                   uint8
	DHCPEnabled                  bool
	DHCPLeaseExpires             time.Time
	DHCPLeaseObtained            time.Time
	DHCPServer                   string
	DNSDomain                    string
	DNSDomainSuffixSearchOrder   []string
	DNSEnabledForWINSResolution  bool
	DNSHostName                  string
	DNSServerSearchOrder         []string
	DomainDNSRegistrationEnabled bool
	ForwardBufferMemory          uint32
	FullDNSRegistrationEnabled   bool
	GatewayCostMetric            []uint16
	IGMPLevel                    uint8
	Index                        uint32
	InterfaceIndex               uint32
	IPAddress                    []net.IP
	IPConnectionMetric           uint32
	IPEnabled                    bool
	IPFilterSecurityEnabled      bool
	IPPortSecurityEnabled        bool
	IPSecPermitIPProtocols       []string
	IPSecPermitTCPPorts          []string
	IPSecPermitUDPPorts          []string
	IPSubnet                     []string
	IPUseZeroBroadcast           bool
	IPXAddress                   string
	IPXEnabled                   bool
	IPXFrameType                 []uint32
	IPXMediaType                 uint32
	IPXNetworkNumber             []string
	IPXVirtualNetNumber          string
	KeepAliveInterval            uint32
	KeepAliveTime                uint32
	MACAddress                   string
	MTU                          uint32
	NumForwardPackets            uint32
	PMTUBHDetectEnabled          bool
	PMTUDiscoveryEnabled         bool
	ServiceName                  string
	TcpipNetbiosOptions          uint32
	TcpMaxConnectRetransmissions uint32
	TcpMaxDataRetransmissions    uint32
	TcpNumConnections            uint32
	TcpUseRFC1122UrgentPointer   bool
	TcpWindowSize                uint16
	WINSEnableLMHostsLookup      bool
	WINSHostLookupFile           string
	WINSPrimaryServer            string
	WINSScopeID                  string
	WINSSecondaryServer          string
}

func (nac *win32NetworkAdapterConfiguration) String() string {

	if nac == nil {
		return "<nil>"
	}

	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf(`Caption: %s
Description: %s
SettingID: %s
ArpAlwaysSourceRoute: %v
ArpUseEtherSNAP: %v
DatabasePath: %s
DeadGWDetectEnabled: %v
DefaultIPGateway:
`, nac.Caption, nac.Description, nac.SettingID, nac.ArpAlwaysSourceRoute, nac.ArpUseEtherSNAP, nac.DatabasePath,
		nac.DeadGWDetectEnabled))

	for _, item := range nac.DefaultIPGateway {
		buffer.WriteString(fmt.Sprintf("    %s\n", item))
	}

	buffer.WriteString(fmt.Sprintf(`DefaultTOS: %d
DefaultTTL: %d
DHCPEnabled: %v
DHCPLeaseExpires: %v
DHCPLeaseObtained: %v
DHCPServer: %s
DNSDomain: %s
DNSDomainSuffixSearchOrder:
`, nac.DefaultTOS, nac.DefaultTTL, nac.DHCPEnabled, nac.DHCPLeaseExpires, nac.DHCPLeaseObtained, nac.DHCPServer,
		nac.DNSDomain))

	for _, item := range nac.DNSDomainSuffixSearchOrder {
		buffer.WriteString(fmt.Sprintf("    %s\n", item))
	}

	buffer.WriteString(fmt.Sprintf(`DNSEnabledForWINSResolution: %v
DNSHostName: %s
DNSServerSearchOrder:
`, nac.DNSEnabledForWINSResolution, nac.DNSHostName))

	for _, item := range nac.DNSServerSearchOrder {
		buffer.WriteString(fmt.Sprintf("    %s\n", item))
	}

	buffer.WriteString(fmt.Sprintf(`DomainDNSRegistrationEnabled: %v
ForwardBufferMemory: %d
FullDNSRegistrationEnabled: %v
GatewayCostMetric:
`, nac.DomainDNSRegistrationEnabled, nac.ForwardBufferMemory, nac.FullDNSRegistrationEnabled))

	for _, item := range nac.GatewayCostMetric {
		buffer.WriteString(fmt.Sprintf("    %d\n", item))
	}

	buffer.WriteString(fmt.Sprintf(`IGMPLevel: %d
Index: %d
InterfaceIndex: %d
IPAddress:
`, nac.IGMPLevel, nac.Index, nac.InterfaceIndex))

	for _, item := range nac.IPAddress {
		buffer.WriteString(fmt.Sprintf("    %s\n", item.String()))
	}

	buffer.WriteString(fmt.Sprintf(`IPConnectionMetric: %d
IPEnabled: %v
IPFilterSecurityEnabled: %v
IPPortSecurityEnabled: %v
IPSecPermitIPProtocols:
`, nac.IPConnectionMetric, nac.IPEnabled, nac.IPFilterSecurityEnabled, nac.IPPortSecurityEnabled))

	for _, item := range nac.IPSecPermitIPProtocols {
		buffer.WriteString(fmt.Sprintf("    %s\n", item))
	}

	buffer.WriteString("IPSecPermitTCPPorts:\n")

	for _, item := range nac.IPSecPermitTCPPorts {
		buffer.WriteString(fmt.Sprintf("    %s\n", item))
	}

	buffer.WriteString("IPSecPermitUDPPorts:\n")

	for _, item := range nac.IPSecPermitUDPPorts {
		buffer.WriteString(fmt.Sprintf("    %s\n", item))
	}

	buffer.WriteString("IPSubnet:\n")

	for _, item := range nac.IPSubnet {
		buffer.WriteString(fmt.Sprintf("    %s\n", item))
	}

	buffer.WriteString(fmt.Sprintf(`IPUseZeroBroadcast: %v
IPXAddress: %s
IPXEnabled: %v
IPXFrameType:
`, nac.IPUseZeroBroadcast, nac.IPXAddress, nac.IPXEnabled))

	for _, item := range nac.IPXFrameType {
		buffer.WriteString(fmt.Sprintf("    %d\n", item))
	}

	buffer.WriteString(fmt.Sprintf(`IPXMediaType: %d
IPXNetworkNumber:`, nac.IPXMediaType))

	for _, item := range nac.IPXNetworkNumber {
		buffer.WriteString(fmt.Sprintf("    %s\n", item))
	}

	buffer.WriteString(fmt.Sprintf(`IPXVirtualNetNumber: %s
KeepAliveInterval: %d
KeepAliveTime: %d
MACAddress: %s
MTU: %d
NumForwardPackets: %d
PMTUBHDetectEnabled: %v
PMTUDiscoveryEnabled: %v
ServiceName: %s
TcpipNetbiosOptions: %d
TcpMaxConnectRetransmissions: %d
TcpMaxDataRetransmissions: %d
TcpNumConnections: %d
TcpUseRFC1122UrgentPointer: %v
TcpWindowSize: %d
WINSEnableLMHostsLookup: %v
WINSHostLookupFile: %s
WINSPrimaryServer: %s
WINSScopeID: %s
WINSSecondaryServer: %s
`, nac.IPXVirtualNetNumber, nac.KeepAliveInterval, nac.KeepAliveTime, nac.MACAddress, nac.MTU, nac.NumForwardPackets,
		nac.PMTUBHDetectEnabled, nac.PMTUDiscoveryEnabled, nac.ServiceName, nac.TcpipNetbiosOptions,
		nac.TcpMaxConnectRetransmissions, nac.TcpMaxDataRetransmissions, nac.TcpNumConnections, nac.TcpUseRFC1122UrgentPointer,
		nac.TcpWindowSize, nac.WINSEnableLMHostsLookup, nac.WINSHostLookupFile, nac.WINSPrimaryServer, nac.WINSScopeID,
		nac.WINSSecondaryServer))

	return buffer.String()
}
