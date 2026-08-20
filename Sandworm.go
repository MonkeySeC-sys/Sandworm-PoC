package main

import (
	"unsafe"
)

var shellcode = []byte{
	//Defanged
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
