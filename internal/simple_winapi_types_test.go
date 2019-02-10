package internal

import (
	"testing"
	"unsafe"
)

func Test_GUID_Size(t *testing.T) {

	const Actual_GUID_Size = unsafe.Sizeof(GUID{})

	if Actual_GUID_Size != GUID_Size {
		t.Errorf("Size of GUID is %d, although %d is expected.", Actual_GUID_Size, GUID_Size)
	}
}
