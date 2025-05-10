package main

import (
	"flag"
	"time"

	"github.com/nnothing1/szu-srun-login/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	var username string
	var password string
	var ip string
	var loop string
	flag.StringVar(&username, "u", "", "Username")
	flag.StringVar(&password, "p", "", "Password")
	// flag.StringVar(&ip, "i", "", "IP Address") 测试了下学校似乎不检测ip这个参数，所以注释掉了
	flag.StringVar(&loop, "t", "", "loop login interval time (30s, 30m, 1h etc)")

	flag.Parse()
	if username == "" || password == "" {
		flag.Usage()
		return
	}
	var ticker *time.Ticker
	if loop != "" {
		interval, err := time.ParseDuration(loop)
		if err != nil {
			logrus.Fatalf("解析 login interval error: %s", err.Error())
		}
		ticker = time.NewTicker(interval)
		defer ticker.Stop()
	}

	if err := utils.Login(username, password, ip); err != nil {
		logrus.Errorf("登录出错: %v", err)
	}

	if ticker == nil {
		return
	}

	for range ticker.C {
		if err := utils.Login(username, password, ip); err != nil {
			logrus.Errorf("登录出错: %v", err)
		}
	}
}
