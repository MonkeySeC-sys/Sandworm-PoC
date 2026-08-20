# Sandworm-PoC
RISC-V 32-bit Linux Reboot Payload.
# Legality's
The author is not responsible for any damage caused by Sandworm, this is only for research and learning.
# Remendation 
1. Security & Prevention
Do Not Run Untrusted Binaries: Never execute Go binaries (or any code) from untrusted sources. Even if the shellcode is for a different architecture (like RISC-V), the Go code itself could contain other malicious routines (like file deletion or credential stealing) that do work on your system.
Restrict Permissions: The reboot system call typically requires root or CAP_SYS_BOOT privileges. Avoid running applications as the root user. If this code is run as a standard user, the kernel will likely block the syscall, and the program will fail.
Use Endpoint Protection: Modern EDR (Endpoint Detection and Response) and Antivirus tools flag the specific pattern of casting a byte slice to a function pointer via unsafe.Pointer, as this is a classic "loader" technique.
