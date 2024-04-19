# Chain

亦称： 职责链模式、命令链、CoR、Chain of Command、Chain of Responsibility

责任链模式是一种行为设计模式， 允许你将请求沿着处理者链进行发送。 收到请求后， 每个处理者均可对请求进行处理， 或将其传递给链上的下个处理者。

## Demo

```go
chain := &SensitiveWordFilterChain{}
chain.AddFilter(&AdSensitiveWordFilter{})
assert.Equal(t, false, chain.Filter("test"))

chain.AddFilter(&PoliticalWordFilter{})
assert.Equal(t, true, chain.Filter("test"))
```