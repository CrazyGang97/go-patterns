# Proxy

代理模式是一种结构型设计模式， 让你能够提供对象的替代品或其占位符。 代理控制着对于原对象的访问， 并允许在将请求提交给对象前后进行一些处理。

## Demo

```go
proxy := NewUserProxy(&User{})

err := proxy.Login("test", "password")
```