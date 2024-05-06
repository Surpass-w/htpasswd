### htpasswd

+ [htpasswd](https://github.com/Surpass-w/htpasswd/blob/main/README.md)
```go
// htpasswd is a simple utility to create and update htpasswd files.
```
+ use
```shell
install -m 755 htpasswd /usr/local/bin/htpasswd
htpasswd -u surpass -p 123456 -f ./htpasswd 
# or
# 如果 不指定用户，默认用户admin,如果不指定密码，会使用openssl生成密码，并会将密码打印到控制台
htpasswd -c ./htpasswd
```