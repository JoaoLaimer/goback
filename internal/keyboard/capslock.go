package keyboard

import (
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	LED_NUML  = 0x00
	LED_CAPSL = 0x01

	EVIOCGLED_INT32 = 0x80044519
)

func CapsLockStatus(fd int) bool {

	var ledStatus uint32

	_, _, errno := unix.Syscall(
		unix.SYS_IOCTL,
		uintptr(fd),
		uintptr(EVIOCGLED_INT32),
		uintptr(unsafe.Pointer(&ledStatus)),
	)
	if errno != 0 {
		return false
	}

	mask := uint32(1 << LED_CAPSL)

	return (ledStatus & mask) != 0

}
