# Builder

生成器模式是一种创建型设计模式， 使你能够分步骤创建复杂对象。 该模式允许你使用相同的创建代码生成不同类型和形式的对象。

## Demo

```go
assembly := NewBuilder().Paint(RedColor)

familyCar := assembly.Wheels(SportsWheels).TopSpeed(50 * MPH).Build()
familyCar.Drive()

sportsCar := assembly.Wheels(SteelWheels).TopSpeed(150 * MPH).Build()
sportsCar.Drive()
```