# Private Struct
 使用私有（private）struct 時，可使用公有（public）interface 來接收

## 需求 & 情境
- 建立自定義 `position` struct
- 強制使用者需先執行 `NewPosition()` 初始化函數
- 欲達成此限制，需設置為**私有** struct（名稱開頭為**小寫**字母）

```go
type position struct {
	x, y, z float64
}

// NewPosition returns an initial position
func NewPosition(pos ...float64) Position {
	if len(pos) == 3 {
		return position{pos[0], pos[1], pos[2]}
	}
	return position{0.0, 0.0, 0.0}
}

func (p position) GetX() float64 {
	return p.x
}
func (p position) GetY() float64 {
	return p.y
}
func (p position) GetZ() float64 {
	return p.z
}
```

## 問題
- 私有 struct 不會被 go doc 工具建立說明文件
- 使用者無法查看該 struct 的內容細節

## 解法
- 加入一個**公開的** interface

```go
// Position defines necessary functions for a position.
type Position interface {
	GetX() float64
	GetY() float64
	GetZ() float64
}
```

## 輸出
```go
func main() {
	fmt.Println(NewPosition(1.0, 2.0, 3.0)) // {1 2 3}
	fmt.Println(NewPosition())              // {0 0 0}
}
```

## 已知問題
- 若 func 使用 Pointer Receiver method（意即 `func (p *position) GetX() float64 {`） 時，須使用 Pointer Type variable（意即 `return &position{0.0, 0.0, 0.0}`）回傳給介面，否則將會出現下方錯誤訊息
- 詳情可參考[這篇](https://mileslin.github.io/2020/08/Golang/為什麼-Pointer-Receiver-不能使用-Value-Type-賦值給-Interface-Value/)文章

```sh
# command-line-arguments
.\main.go:19:18: cannot use position literal (type position) as type Position in return argument:
        position does not implement Position (GetX method has pointer receiver)
.\main.go:21:17: cannot use position literal (type position) as type Position in return argument:
        position does not implement Position (GetX method has pointer receiver)
```

## References
- [2020/09/14 - 在 Go 語言中，我為什麼使用介面](https://mp.weixin.qq.com/s/AMgCzCG_096iaCdtOJIBBA)
- [2020/08/01 - [Go] 為什麼 Pointer Receiver 不能使用 Value Type 賦值給 Interface Value](https://mileslin.github.io/2020/08/Golang/為什麼-Pointer-Receiver-不能使用-Value-Type-賦值給-Interface-Value/)