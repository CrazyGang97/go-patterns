# Adapter

适配器模式是一种结构型设计模式， 它能使接口不兼容的对象能够相互合作。

## Demo

```go
func TestAliyunClientAdapter_CreateServer(t *testing.T) {
	// 确保 adapter 实现了目标接口 
	var a ICreateServer = &AliyunClientAdapter{
		Client: AliyunClient{},
	}
	a.CreateServer(1.0, 2.0)
}

func TestAwsClientAdapter_CreateServer(t *testing.T) {
	// 确保 adapter 实现了目标接口 
	var a ICreateServer = &AwsClientAdapter{
		Client: AWSClient{},
	}
	a.CreateServer(1.0, 2.0)
}
```