# Template

模板方法模式是一种行为设计模式， 它在超类中定义了一个算法的框架， 允许子类在不修改结构的情况下重写算法的特定步骤。

## Demo

```go
smsOTP := &Sms{}
o := Otp{
    iOtp: smsOTP,
}
o.genAndSendOTP(4)
fmt.Println("")
emailOTP := &Email{}
o = Otp{
    iOtp: emailOTP,
}
o.genAndSendOTP(4)
```