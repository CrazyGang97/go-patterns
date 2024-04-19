# Command

亦称： 动作、事务、Action、Transaction、Command

命令模式是一种行为设计模式， 它可将请求转换为一个包含与请求相关的所有信息的独立对象。 该转换让你能根据不同的请求将方法参数化、 延迟请求执行或将其放入队列中， 且能实现可撤销操作。

## Demo

```go
// 用于测试，模拟来自客户端的事件 
eventChan := make(chan string)
go func() {
	events := []string{"start", "archive", "start", "archive", "start", "start"}
	for _, e := range events {
		eventChan <- e
	}
}()
defer close(eventChan) 
// 使用命令队列缓存命令 
commands := make(chan ICommand, 1000)
defer close(commands)

go func() {
	for {
		// 从请求或者其他地方获取相关事件参数 
		event, ok := <-eventChan
		if !ok {
			return
		}
		
		var command ICommand
		switch event {
		case "start":
			command = NewStartCommand()
			case "archive":
				command = NewArchiveCommand()
		}
		// 将命令入队 
		commands <- command
	}
}()
for {
	select {
	case c := <-commands:
		c.Execute()
		case <-time.After(1 * time.Second):
			fmt.Println("timeout 1s")
			return
	}
}
```

```go
// 用于测试，模拟来自客户端的事件
eventChan := make(chan string)
go func() {
	events := []string{"start", "archive", "start", "archive", "start", "start"}
	for _, e := range events {
		eventChan <- e
	}

}()
defer close(eventChan)

// 使用命令队列缓存命令
commands := make(chan Command, 1000)
defer close(commands)

go func() {
	for {
		// 从请求或者其他地方获取相关事件参数
		event, ok := <-eventChan
		if !ok {
			return
		}

		var command Command
		switch event {
		case "start":
			command = StartCommandFunc()
		case "archive":
			command = ArchiveCommandFunc()
		}

		// 将命令入队
		commands <- command
	}
}()

for {
	select {
	case c := <-commands:
		c()
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1s")
		return
	}
}
```