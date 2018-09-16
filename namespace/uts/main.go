package main

/*
hostname
*/

import (
	"os/exec"
	"syscall"
	"os"
	"log"
)

func main() {
	cmd := exec.Command("sh")  // 封装了os.StartProcess函数，实际是起新进程
	cmd.SysProcAttr = &syscall.SysProcAttr {
		Cloneflags: syscall.CLONE_NEWUTS,  // 指示clone()系统调用的参数，即创建哪些namespace
	}
	cmd.Stdin = os.Stdin  // 子命令的标准输出和输入
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

/*
pstree -pl
hosname -b
*/

// go 子进程启动过程：http://blog.51cto.com/allragedbody/1747146
// readlink /proc/{pid}/ns/uts：软链到是内存数据
// readlink 读取软链接：https://blog.csdn.net/diabloneo/article/details/7173438
