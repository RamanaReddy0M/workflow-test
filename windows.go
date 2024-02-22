// +build windows

package main

import (
    "fmt"
    "golang.org/x/sys/windows"
    "syscall"
    "unsafe"
)

func readPassword(prompt string) ([]byte, error) {
    fmt.Print(prompt)
    consoleHandle, err := windows.GetStdHandle(windows.STD_INPUT_HANDLE)
    if err != nil {
        return nil, err
    }

    var bytesRead uint32
    maxChar := 256
    buff := make([]uint16, maxChar)
    ret, _, err := windows.ProcReadConsoleW.Call(
        uintptr(consoleHandle),
        uintptr(unsafe.Pointer(&buff[0])),
        uintptr(maxChar),
        uintptr(unsafe.Pointer(&bytesRead)),
        0,
    )
    if ret == 0 {
        return nil, err
    }

    // Convert from UTF-16 to UTF-8 excluding the newline character
    utf8Bytes, err := windows.UTF16ToUTF8(buff[:bytesRead-2])
    if err != nil {
        return nil, err
    }

    return utf8Bytes, nil
}

func main() {
    password, err := readPassword("Enter password: ")
    if err != nil {
        fmt.Println("Error reading password:", err)
        return
    }

    fmt.Println("\nPassword entered:", string(password))
}
