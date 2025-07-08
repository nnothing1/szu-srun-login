# SZU-SRUN-LOGIN
简易的2025年新版深圳大学教学区校园网命令行登录工具，主要的编码加密逻辑参考自：[BitSrunLoginGo](https://github.com/Mmx233/BitSrunLoginGo)。

## 使用方式

1. 在 [releases](https://github.com/nnothing1/szu-srun-login/releases) 页面根据你的操作系统下载并解压文件。
2. 根据你的操作系统运行相应的命令。

### linux
```bash
chmod u+x szu-srun-login
./szu-srun-login -u USER_ID -p PASSWORD
```
### windows
powershell
```powershell
.\szu-srun-login.exe -u USER_ID -p PASSWORD
```
cmd
```cmd
szu-srun-login.exe -u USER_ID -p PASSWORD
```
### 定时循环登录
虽然我更推荐使用 linux 中的 crontab 来定时登录，但本工具也提供了简易的定时登录功能，你可以使用 30s, 30m, 1h 等时间格式，并搭配 nohup 等工具将其挂起在后台。注意时间间隔不要太小，否则可能 IP 会被学校 ban 。
```
./szu-srun-login -u USER_ID -p PASSWORD -t 30s
```