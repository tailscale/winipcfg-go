// Code generated by 'go generate'; DO NOT EDIT.

package winipcfg

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modiphlpapi = windows.NewLazySystemDLL("iphlpapi.dll")

	procGetAdaptersAddresses            = modiphlpapi.NewProc("GetAdaptersAddresses")
	procInitializeIpInterfaceEntry      = modiphlpapi.NewProc("InitializeIpInterfaceEntry")
	procGetIpInterfaceEntry             = modiphlpapi.NewProc("GetIpInterfaceEntry")
	procGetIpInterfaceTable             = modiphlpapi.NewProc("GetIpInterfaceTable")
	procSetIpInterfaceEntry             = modiphlpapi.NewProc("SetIpInterfaceEntry")
	procFreeMibTable                    = modiphlpapi.NewProc("FreeMibTable")
	procGetIfEntry2Ex                   = modiphlpapi.NewProc("GetIfEntry2Ex")
	procGetIfTable2Ex                   = modiphlpapi.NewProc("GetIfTable2Ex")
	procGetUnicastIpAddressTable        = modiphlpapi.NewProc("GetUnicastIpAddressTable")
	procGetUnicastIpAddressEntry        = modiphlpapi.NewProc("GetUnicastIpAddressEntry")
	procSetUnicastIpAddressEntry        = modiphlpapi.NewProc("SetUnicastIpAddressEntry")
	procInitializeUnicastIpAddressEntry = modiphlpapi.NewProc("InitializeUnicastIpAddressEntry")
	procCreateUnicastIpAddressEntry     = modiphlpapi.NewProc("CreateUnicastIpAddressEntry")
	procDeleteUnicastIpAddressEntry     = modiphlpapi.NewProc("DeleteUnicastIpAddressEntry")
	procGetAnycastIpAddressTable        = modiphlpapi.NewProc("GetAnycastIpAddressTable")
	procGetAnycastIpAddressEntry        = modiphlpapi.NewProc("GetAnycastIpAddressEntry")
	procCreateAnycastIpAddressEntry     = modiphlpapi.NewProc("CreateAnycastIpAddressEntry")
	procDeleteAnycastIpAddressEntry     = modiphlpapi.NewProc("DeleteAnycastIpAddressEntry")
	procGetIpForwardTable2              = modiphlpapi.NewProc("GetIpForwardTable2")
	procGetIpForwardEntry2              = modiphlpapi.NewProc("GetIpForwardEntry2")
	procInitializeIpForwardEntry        = modiphlpapi.NewProc("InitializeIpForwardEntry")
	procCreateIpForwardEntry2           = modiphlpapi.NewProc("CreateIpForwardEntry2")
	procSetIpForwardEntry2              = modiphlpapi.NewProc("SetIpForwardEntry2")
	procDeleteIpForwardEntry2           = modiphlpapi.NewProc("DeleteIpForwardEntry2")
	procNotifyIpInterfaceChange         = modiphlpapi.NewProc("NotifyIpInterfaceChange")
	procNotifyUnicastIpAddressChange    = modiphlpapi.NewProc("NotifyUnicastIpAddressChange")
	procNotifyRouteChange2              = modiphlpapi.NewProc("NotifyRouteChange2")
	procCancelMibChangeNotify2          = modiphlpapi.NewProc("CancelMibChangeNotify2")
)

func getAdaptersAddresses(Family uint32, Flags uint32, Reserved uintptr, AdapterAddresses *wtIpAdapterAddresses, SizePointer *uint32) (result uint32) {
	r0, _, _ := syscall.Syscall6(procGetAdaptersAddresses.Addr(), 5, uintptr(Family), uintptr(Flags), uintptr(Reserved), uintptr(unsafe.Pointer(AdapterAddresses)), uintptr(unsafe.Pointer(SizePointer)), 0)
	result = uint32(r0)
	return
}

func initializeIpInterfaceEntry(Row *wtMibIpinterfaceRow) (result int32) {
	r0, _, _ := syscall.Syscall(procInitializeIpInterfaceEntry.Addr(), 1, uintptr(unsafe.Pointer(Row)), 0, 0)
	result = int32(r0)
	return
}

func getIpInterfaceEntry(Row *wtMibIpinterfaceRow) (result int32) {
	r0, _, _ := syscall.Syscall(procGetIpInterfaceEntry.Addr(), 1, uintptr(unsafe.Pointer(Row)), 0, 0)
	result = int32(r0)
	return
}

func getIpInterfaceTable(Family AddressFamily, Table unsafe.Pointer) (result int32) {
	r0, _, _ := syscall.Syscall(procGetIpInterfaceTable.Addr(), 2, uintptr(Family), uintptr(Table), 0)
	result = int32(r0)
	return
}

func setIpInterfaceEntry(Row *wtMibIpinterfaceRow) (result int32) {
	r0, _, _ := syscall.Syscall(procSetIpInterfaceEntry.Addr(), 1, uintptr(unsafe.Pointer(Row)), 0, 0)
	result = int32(r0)
	return
}

func freeMibTable(memory unsafe.Pointer) {
	syscall.Syscall(procFreeMibTable.Addr(), 1, uintptr(memory), 0, 0)
	return
}

func getIfEntry2Ex(Level MibIfEntryLevel, Row *wtMibIfRow2) (result int32) {
	r0, _, _ := syscall.Syscall(procGetIfEntry2Ex.Addr(), 2, uintptr(Level), uintptr(unsafe.Pointer(Row)), 0)
	result = int32(r0)
	return
}

func getIfTable2Ex(Level MibIfEntryLevel, Table unsafe.Pointer) (result int32) {
	r0, _, _ := syscall.Syscall(procGetIfTable2Ex.Addr(), 2, uintptr(Level), uintptr(Table), 0)
	result = int32(r0)
	return
}

func getUnicastIpAddressTable(Family AddressFamily, Table unsafe.Pointer) (result int32) {
	r0, _, _ := syscall.Syscall(procGetUnicastIpAddressTable.Addr(), 2, uintptr(Family), uintptr(Table), 0)
	result = int32(r0)
	return
}

func getUnicastIpAddressEntry(Row *wtMibUnicastipaddressRow) (result int32) {
	r0, _, _ := syscall.Syscall(procGetUnicastIpAddressEntry.Addr(), 1, uintptr(unsafe.Pointer(Row)), 0, 0)
	result = int32(r0)
	return
}

func setUnicastIpAddressEntry(Row *wtMibUnicastipaddressRow) (result int32) {
	r0, _, _ := syscall.Syscall(procSetUnicastIpAddressEntry.Addr(), 1, uintptr(unsafe.Pointer(Row)), 0, 0)
	result = int32(r0)
	return
}

func initializeUnicastIpAddressEntry(Row *wtMibUnicastipaddressRow) (result int32) {
	r0, _, _ := syscall.Syscall(procInitializeUnicastIpAddressEntry.Addr(), 1, uintptr(unsafe.Pointer(Row)), 0, 0)
	result = int32(r0)
	return
}

func createUnicastIpAddressEntry(Row *wtMibUnicastipaddressRow) (result int32) {
	r0, _, _ := syscall.Syscall(procCreateUnicastIpAddressEntry.Addr(), 1, uintptr(unsafe.Pointer(Row)), 0, 0)
	result = int32(r0)
	return
}

func deleteUnicastIpAddressEntry(Row *wtMibUnicastipaddressRow) (result int32) {
	r0, _, _ := syscall.Syscall(procDeleteUnicastIpAddressEntry.Addr(), 1, uintptr(unsafe.Pointer(Row)), 0, 0)
	result = int32(r0)
	return
}

func getAnycastIpAddressTable(Family AddressFamily, Table unsafe.Pointer) (result int32) {
	r0, _, _ := syscall.Syscall(procGetAnycastIpAddressTable.Addr(), 2, uintptr(Family), uintptr(Table), 0)
	result = int32(r0)
	return
}

func getAnycastIpAddressEntry(Row *wtMibAnycastipaddressRow) (result int32) {
	r0, _, _ := syscall.Syscall(procGetAnycastIpAddressEntry.Addr(), 1, uintptr(unsafe.Pointer(Row)), 0, 0)
	result = int32(r0)
	return
}

func createAnycastIpAddressEntry(Row *wtMibAnycastipaddressRow) (result int32) {
	r0, _, _ := syscall.Syscall(procCreateAnycastIpAddressEntry.Addr(), 1, uintptr(unsafe.Pointer(Row)), 0, 0)
	result = int32(r0)
	return
}

func deleteAnycastIpAddressEntry(Row *wtMibAnycastipaddressRow) (result int32) {
	r0, _, _ := syscall.Syscall(procDeleteAnycastIpAddressEntry.Addr(), 1, uintptr(unsafe.Pointer(Row)), 0, 0)
	result = int32(r0)
	return
}

func getIpForwardTable2(family AddressFamily, table unsafe.Pointer) (result int32) {
	r0, _, _ := syscall.Syscall(procGetIpForwardTable2.Addr(), 2, uintptr(family), uintptr(table), 0)
	result = int32(r0)
	return
}

func getIpForwardEntry2(route *wtMibIpforwardRow2) (result int32) {
	r0, _, _ := syscall.Syscall(procGetIpForwardEntry2.Addr(), 1, uintptr(unsafe.Pointer(route)), 0, 0)
	result = int32(r0)
	return
}

func initializeIpForwardEntry(route *wtMibIpforwardRow2) (result int32) {
	r0, _, _ := syscall.Syscall(procInitializeIpForwardEntry.Addr(), 1, uintptr(unsafe.Pointer(route)), 0, 0)
	result = int32(r0)
	return
}

func createIpForwardEntry2(route *wtMibIpforwardRow2) (result int32) {
	r0, _, _ := syscall.Syscall(procCreateIpForwardEntry2.Addr(), 1, uintptr(unsafe.Pointer(route)), 0, 0)
	result = int32(r0)
	return
}

func setIpForwardEntry2(route *wtMibIpforwardRow2) (result int32) {
	r0, _, _ := syscall.Syscall(procSetIpForwardEntry2.Addr(), 1, uintptr(unsafe.Pointer(route)), 0, 0)
	result = int32(r0)
	return
}

func deleteIpForwardEntry2(route *wtMibIpforwardRow2) (result int32) {
	r0, _, _ := syscall.Syscall(procDeleteIpForwardEntry2.Addr(), 1, uintptr(unsafe.Pointer(route)), 0, 0)
	result = int32(r0)
	return
}

func notifyIpInterfaceChange(Family AddressFamily, Callback uintptr, CallerContext uintptr, InitialNotification bool, NotificationHandle unsafe.Pointer) (result int32) {
	var _p0 uint32
	if InitialNotification {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r0, _, _ := syscall.Syscall6(procNotifyIpInterfaceChange.Addr(), 5, uintptr(Family), uintptr(Callback), uintptr(CallerContext), uintptr(_p0), uintptr(NotificationHandle), 0)
	result = int32(r0)
	return
}

func notifyUnicastIpAddressChange(Family AddressFamily, Callback uintptr, CallerContext uintptr, InitialNotification bool, NotificationHandle unsafe.Pointer) (result int32) {
	var _p0 uint32
	if InitialNotification {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r0, _, _ := syscall.Syscall6(procNotifyUnicastIpAddressChange.Addr(), 5, uintptr(Family), uintptr(Callback), uintptr(CallerContext), uintptr(_p0), uintptr(NotificationHandle), 0)
	result = int32(r0)
	return
}

func notifyRouteChange2(Family AddressFamily, Callback uintptr, CallerContext uintptr, InitialNotification bool, NotificationHandle unsafe.Pointer) (result int32) {
	var _p0 uint32
	if InitialNotification {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r0, _, _ := syscall.Syscall6(procNotifyRouteChange2.Addr(), 5, uintptr(Family), uintptr(Callback), uintptr(CallerContext), uintptr(_p0), uintptr(NotificationHandle), 0)
	result = int32(r0)
	return
}

func cancelMibChangeNotify2(NotificationHandle uintptr) (result int32) {
	r0, _, _ := syscall.Syscall(procCancelMibChangeNotify2.Addr(), 1, uintptr(NotificationHandle), 0, 0)
	result = int32(r0)
	return
}
