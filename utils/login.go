package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/Mmx233/BitSrunLoginGo/pkg/srun"
	"github.com/sirupsen/logrus"
)

func Login(username, password, ip string) error {
	callback := "callback"                                 //随机字符串就行
	challenge, err := getChallenge(username, ip, callback) //先获取challenge
	if err != nil {
		return err
	}
	//固定参数
	action := "login"
	ac_id := "12"
	n := "200"
	typ := "1"
	enc := "srun_bx1"

	//加密密码
	encrypted_pwd := EncryptPassword(challenge, password)

	//构造info并加密
	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Ip       string `json:"ip"`
		Acid     string `json:"acid"`
		EncVer   string `json:"enc_ver"`
	}
	data.Username = username
	data.Password = password
	data.Ip = ip
	data.Acid = ac_id
	data.EncVer = enc
	info := EncodeUserInfo(data, challenge)

	//构造chcksum
	chcksum := srun.Sha1(challenge + username +
		challenge + encrypted_pwd +
		challenge + ac_id +
		challenge + ip +
		challenge + n +
		challenge + typ +
		challenge + info,
	)

	//构造请求参数
	params := url.Values{}
	params.Add("action", action)
	params.Add("ac_id", ac_id)
	params.Add("n", n)
	params.Add("type", typ)
	params.Add("ip", ip)
	params.Add("username", username)
	params.Add("password", "{MD5}"+encrypted_pwd)
	params.Add("info", info)
	params.Add("chksum", chcksum)
	params.Add("callback", callback)
	base_url := "https://net.szu.edu.cn/cgi-bin/srun_portal"

	//发送请求
	resp, err := http.Get(base_url + "?" + params.Encode())
	if err != nil {
		logrus.Fatalf("发送登录请求失败: %v", err)
		return err
	}
	defer resp.Body.Close()

	//读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatalf("读取响应内容失败: %v", err)
		return err
	}
	if err := parseLoginBody(body[len(callback)+1 : len(body)-1]); err != nil {
		logrus.Fatalf("解析登录响应失败: %v", err)
	}
	return nil
}

func getChallenge(username, ip, callback string) (string, error) {

	//构造请求参数
	base_url := "https://net.szu.edu.cn/cgi-bin/get_challenge"
	params := url.Values{}
	params.Add("username", username)
	params.Add("ip", ip)
	params.Add("callback", callback)

	//发送请求
	resp, err := http.Get(base_url + "?" + params.Encode())
	if err != nil {
		logrus.Fatalf("发送获取challenge请求失败: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	//读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatalf("读取challenge请求响应内容失败: %v", err)
		return "", err
	}

	//解析JSON响应内容
	var challenge struct {
		Challenge string `json:"challenge"`
	}
	if err = json.Unmarshal(body[len(callback)+1:len(body)-1], &challenge); err != nil {
		logrus.Fatalf("解析challenge响应内容失败: %v", err)
		return "", err
	}
	return challenge.Challenge, nil
}

func parseLoginBody(body []byte) error {
	var data struct {
		Res      string `json:"res"`
		SucMsg   string `json:"suc_msg"`
		ErrorMsg string `json:"error_msg"`
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}
	if data.Res != "ok" {
		logrus.Errorf("登录失败, error_msg : %s", data.ErrorMsg)
	} else if data.SucMsg == "ip_already_online_error" {
		logrus.Infof("已登录过 suc_msg: %s", data.SucMsg)
	} else {
		logrus.Infof("登录成功 suc_msg: %s", data.SucMsg)
	}
	return nil
}
