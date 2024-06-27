### htpasswd

+ [htpasswd](https://github.com/Surpass-w/htpasswd/blob/main/README.md)
```go
// htpasswd is a simple utility to create and update htpasswd files.
```
+ use
```shell
install -m 755 htpasswd /usr/local/bin/htpasswd
htpasswd -u surpass -p 123456 -f ./htpasswd 
```
+ 提供一个helm chart 包示例
```shell
helm install -n your_namespace resource ./resource

```