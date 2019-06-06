# GoVendor

https://github.com/kardianos/govendor

```
go get -u github.com/kardianos/govendor
```

注意:

- The project must be within a `$GOPATH/src`.
- If using go1.5, ensure you set `GO15VENDOREXPERIMENT=1`.

项目导出依赖
```
# 初始化govendor
govendor init
# 从GOPATH拷贝当前代码所需依赖到当前vendor目录下
govendor add +e
```

安装项目依赖
```
govendor sync
```

添加项目依赖
```
govendor add +e
```

更新项目依赖（新增依赖 + 更新已有依赖）
```
govendor add +e
govendor update +v
```

中国特色依赖

```
Error: Remotes failed for:
	Failed for "cloud.google.com/go/civil" (failed to ping remote repo): unrecognized import path "cloud.google.com/go/civil"
	Failed for "golang.org/x/crypto/md4" (failed to ping remote repo): unrecognized import path "golang.org/x/crypto/md4"
	Failed for "golang.org/x/crypto/ssh/terminal" (failed to ping remote repo): unrecognized import path "golang.org/x/crypto/ssh/terminal"
	Failed for "golang.org/x/sys/unix" (failed to ping remote repo): unrecognized import path "golang.org/x/sys/unix"
	Failed for "golang.org/x/sys/windows" (failed to ping remote repo): unrecognized import path "golang.org/x/sys/windows"
	Failed for "golang.org/x/text/transform" (failed to ping remote repo): unrecognized import path "golang.org/x/text/transform"
	Failed for "golang.org/x/text/unicode/norm" (failed to ping remote repo): unrecognized import path "golang.org/x/text/unicode/norm"
```

```bash
mkdir -p $GOPATH/src/cloud.google.com
git clone git@github.com:GoogleCloudPlatform/google-cloud-go.git $GOPATH/src/cloud.google.com/go

mkdir -p $GOPATH/src/golang.org/x
git clone git@github.com:golang/crypto.git $GOPATH/src/golang.org/x/crypto
git clone git@github.com:golang/sys.git $GOPATH/src/golang.org/x/sys
git clone git@github.com:golang/text.git $GOPATH/src/golang.org/x/text
```

或者 科学上网方式（推荐），shadowsocks 开启http代理服务器

`go get`之前导入代理环境变量

临时设置
```
export http_proxy=http://127.0.0.1:1087
export https_proxy=http://127.0.0.1:1087
```

临时取消
```
unset http_proxy
unset https_proxy
```
