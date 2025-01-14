package main

import (
	"flag"

	"github.com/nnothing1/szu-srun-login/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	var username string
	var password string
	var ip string
	flag.StringVar(&username, "u", "", "Username")
	flag.StringVar(&password, "p", "", "Password")
	// flag.StringVar(&ip, "i", "", "IP Address") 测试了下学校似乎不检测ip这个参数，所以注释掉了
	flag.Parse()
	if username == "" || password == "" {
		flag.Usage()
		return
	}
	if err := utils.Login(username, password, ip); err != nil {
		logrus.Fatalf("登录出错: %v", err)
	}
}
