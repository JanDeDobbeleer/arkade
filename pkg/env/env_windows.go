//go:build windows

package env

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	dllKernel               = windows.NewLazyDLL("kernel32.dll")
	procGetNativeSystemInfo = dllKernel.NewProc("GetNativeSystemInfo")
)

type systemInfo struct {
	wProcessorArchitecture      uint16
	wReserved                   uint16
	dwPageSize                  uint32
	lpMinimumApplicationAddress uintptr
	lpMaximumApplicationAddress uintptr
	dwActiveProcessorMask       uintptr
	dwNumberOfProcessors        uint32
	dwProcessorType             uint32
	dwAllocationGranularity     uint32
	wProcessorLevel             uint16
	wProcessorRevision          uint16
}

func GetClientArch() (arch string, os string) {
	var systemInfo systemInfo
	_, _, _ = procGetNativeSystemInfo.Call(uintptr(unsafe.Pointer(&systemInfo)))

	const (
		PROCESSOR_ARCHITECTURE_INTEL = 0
		PROCESSOR_ARCHITECTURE_ARM   = 5
		PROCESSOR_ARCHITECTURE_ARM64 = 12
		PROCESSOR_ARCHITECTURE_IA64  = 6
		PROCESSOR_ARCHITECTURE_AMD64 = 9
		WINDOWS                      = "ming"
	)

	switch systemInfo.wProcessorArchitecture {
	case PROCESSOR_ARCHITECTURE_INTEL:
		if systemInfo.wProcessorLevel < 3 {
			return "i386", WINDOWS
		}
		if systemInfo.wProcessorLevel > 6 {
			return "i686", WINDOWS
		}
		return fmt.Sprintf("i%d86", systemInfo.wProcessorLevel), WINDOWS
	case PROCESSOR_ARCHITECTURE_ARM:
		return "arm", WINDOWS
	case PROCESSOR_ARCHITECTURE_ARM64:
		return "aarch64", WINDOWS
	case PROCESSOR_ARCHITECTURE_IA64:
		return "ia64", WINDOWS
	case PROCESSOR_ARCHITECTURE_AMD64:
		return "x86_64", WINDOWS
	}
	return "", WINDOWS
}
