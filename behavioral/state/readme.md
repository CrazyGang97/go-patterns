# State

状态模式是一种行为设计模式， 让你能在一个对象的内部状态变化时改变其行为， 使其看上去就像改变了自身所属的类一样。

## Demo

```go
m := &Machine{state: GetLeaderApproveState()}
assert.Equal(t, "LeaderApproveState", m.GetStateName())
m.Approval()
assert.Equal(t, "FinanceApproveState", m.GetStateName())
m.Reject()
assert.Equal(t, "LeaderApproveState", m.GetStateName())
m.Approval()
assert.Equal(t, "FinanceApproveState", m.GetStateName())
m.Approval()
```