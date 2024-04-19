# Mediator

中介者模式是一种行为设计模式， 能让你减少对象之间混乱无序的依赖关系。 该模式会限制对象之间的直接交互， 迫使它们通过一个中介者对象进行合作。

## Demo

```go
usernameInput := Input("username input")
passwordInput := Input("password input")
repeatPwdInput := Input("repeat password input")

selection := Selection("登录")
d := &Dialog{
    Selection:           &selection,
    UsernameInput:       &usernameInput,
    PasswordInput:       &passwordInput,
    RepeatPasswordInput: &repeatPwdInput,
}
d.HandleEvent(&selection)

regSelection := Selection("注册")
d.Selection = &regSelection
d.HandleEvent(&regSelection)
```