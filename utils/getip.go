package utils

import (
	"errors"
	"io"
	"net/http"
	"regexp"

	"github.com/sirupsen/logrus"
)

func GetIP() (string, error) {
	url := "https://net.szu.edu.cn/srun_portal_success?ac_id=12&theme=proyx"
	resp, err := http.Get(url)
	if err != nil {
		logrus.Errorf("发送获取IP请求失败: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("读取获取IP请求响应内容失败: %v", err)
		return "", err
	}
	ipPattern := `ip\s*:\s*"([^"]+)"`
	re := regexp.MustCompile(ipPattern)
	matches := re.FindStringSubmatch(string(body))
	if len(matches) > 1 {
		return matches[1], nil
	} else {
		return "", errors.New("正则匹配并获取响应内容中的IP地址失败")
	}
}
