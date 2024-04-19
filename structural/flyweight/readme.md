# Flyweight

享元模式是一种结构型设计模式， 它摒弃了在每个对象中保存所有数据的方式， 通过共享多个对象所共有的相同状态， 让你能在有限的内存容量中载入更多对象。

## Demo

```go
board1 := NewChessBoard()
board1.Move(1, 1, 2)
board2 := NewChessBoard()
board2.Move(2, 2, 3)

assert.Equal(t, board1.chessPieces[1].Unit, board2.chessPieces[1].Unit)
assert.Equal(t, board1.chessPieces[2].Unit, board2.chessPieces[2].Unit)
```