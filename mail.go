package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/jander/golog/logger"
	"github.com/kardianos/service"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	// 代码写在这儿
	for {
		func() {
			cmd := exec.Command("C:\\bin\\frp\\frpc.exe", "-c", "C:\\bin\\frp\\frpc.ini")
			logger.Printf("启动进程: %v", cmd.Args)
			err := cmd.Run()
			if err != nil {
				logger.Printf("进程异常退出: %v", err.Error())
			}
		}()
		logger.Println("进程已退出, 10 秒后重启：")
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 1)
			fmt.Printf("...%v", i+1)
		}
		fmt.Println()

	}
}

func (p *program) Stop(s service.Service) error {
	return nil
}

/**
* MAIN函数，程序入口
 */

func main() {
	svcConfig := &service.Config{
		Name:        "Frpc",       //服务显示名称
		DisplayName: "Frpc",       //服务名称
		Description: "Frp Client", //服务描述
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		logger.Fatal(err)
	}

	if err != nil {
		logger.Fatal(err)
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			s.Install()
			logger.Println("服务安装成功")
			return
		}

		if os.Args[1] == "remove" {
			s.Uninstall()
			logger.Println("服务卸载成功")
			return
		}
	}

	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
