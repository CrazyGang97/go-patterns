# Strategy

策略模式是一种行为设计模式， 它能让你定义一系列算法， 并将每种算法分别放入独立的类中， 以使算法的对象能够相互替换。

## Demo

```go
func Test_demo(t *testing.T) {
    // 假设这里获取数据，以及数据是否敏感
    data, sensitive := getData()
    strategyType := "file"
    if sensitive {
        strategyType = "encrypt_file"
	}

    storage, err := NewStorageStrategy(strategyType)
    assert.NoError(t, err)
    
    dir, err := os.MkdirTemp("", "strategy")
    assert.NoError(t, err)
    
    assert.NoError(t, storage.Save(dir+"/test.txt", data))
}

// getData 获取数据的方法
// 返回数据，以及数据是否敏感
func getData() ([]byte, bool) {
    return []byte("test data"), false
}
```