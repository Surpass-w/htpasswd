### nginx 文件资源服务 helm 包
+ 生成身份验证的密码
```shell
git clone https://github.com/Surpass-w/htpasswd.git
```
+ 下载resource 资源包
```shell
git clone https://github.com/Surpass-w/resource.git
```

+ 启动服务
```shell
## 修改对应的配置
helm install -n namespace resource ./resource
```