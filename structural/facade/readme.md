# Facade

外观模式是一种结构型设计模式， 能为程序库、 框架或其他复杂类提供一个简单的接口。

## Demo

```go
func TestUserService_Login(t *testing.T) {
	service := NewUserService("lily")
	user, err := service.User.Login(13001010101, 1234)
}

func TestUserService_LoginOrRegister(t *testing.T) {
	service := NewUserService("lily")
	user, err := service.LoginOrRegister(13001010101, 1234)
}

func TestUserService_Register(t *testing.T) {
	service := NewUserService("lily")
	user, err := service.User.Register(13001010101, 1234)
}
```