# Decorator

亦称： 装饰者模式、装饰器模式、Wrapper、Decorator

装饰模式是一种结构型设计模式， 允许你通过将对象放入包含行为的特殊封装对象中来为原对象绑定新的行为。

## Demo

```go
sq := Square{}
csq := NewColorSquare(sq, "red")
got := csq.Draw()
```