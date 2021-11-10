# RegistryByuac
云影/K0uaz

## 介绍
基于注册表的BypassUac工具，原理是利用具有AutoElevate属性的程序运行时自动提升权限这一特性  
支持6种BypassUac的方式:Eventvwr, Fodhelper, Computerdefaults, Sdclt, Slui, DiskCleanup计划任务

## 适用版本
Eventvwr : 适用Windows vista, Windows 7, Windows 8, Windows 8.1, Windows 10(win10 Update preview build 15007 修复)
Fodhelper : 适用win10
Computerdefaults : 适用win10 
Sdclt : 适用win7或更高版本
Slui : 适用win8-10(Windows 10 20H1(19041) 修复)
DiskCleanup : 适用win10和win8

## 使用
RegistryByuac.exe -h
```
Support Eventvwr/ComputerDefaults/Fodhelper/Sdclt/Slui/DiskCleanup
Quick start: ./RegistryByuac -m Ev/Co/Fo/Sd/Sl/Di -e Y2FsYy5leGU=
  -e string
        File Path Or Command
  -h    Help
  -m string
        Select Ev/Co/Fo/Sd/Sl
  -v    Search For Winver
```
Tips:  
1. 使用DiskCleanup执行命令时需要添加`&& REM`来注释后面的`\system32\cleanmgr.exe`

## Bypass问题
在装有360的Win7(虚拟机)中测试了Eventvwr，无法bypass，注册表项会被提示检测(火绒和QQ管家默认设置，当时测试一点相应没有，不拦)
在默认只装有Windows Defender中测试了适用Win10(虚拟机)的几个注册表项添加，无法bypass
之前看到一个项目，用简单的白名单Computerdefaults的方式即可绕过安装了360的win10 
→因为懒，不想在win10(虚拟机)中安装360，因此没有测试适用Win10的几个方式是否能够绕过360
如果有师傅知道如何绕过360的注册表监测的话，希望能带带弟弟

## 可能去实现的
1. 可以参照Sharpbypassuac把path修改为base64的方式(已改成base64)
2. COM接口bypassuac