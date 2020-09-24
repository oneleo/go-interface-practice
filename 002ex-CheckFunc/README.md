# Check Function
 使用 Interface 來確認 Struct 是否有實現指定的方法

## 需求 & 情境
- 確認 Struct 是否有實現指定方法（Draw() 方法）
- 若有就執行指定方法

```go
// ---------- 第一個結構 ----------
// MyGame 只實現 Update() 方法
type MyGame struct {
}
func (mg *MyGame) Update() {
	fmt.Println("Execute MyGame.Update() Method")
}

// ---------- 第二個結構 ----------
// MyGameWithDraw 有實現 Update() 及 Draw() 方法
type MyGameWithDraw struct {
}
func (mg *MyGameWithDraw) Update() {
	fmt.Println("Execute MyGameWithDraw.Update() Method")
}
func (mg *MyGameWithDraw) Draw() {
	fmt.Println("Execute MyGameWithDraw.Draw() Method")
}
```

## 解法
- 使用 interface 來接 struct 及確認是否有指定方法

```go
type IGame interface {
	Update()
}

func RunGame(gi IGame) {
	gi.Update()
	if game, ok := gi.(interface{ Draw() }); ok {
		game.Draw()
	}
}
```

## 輸出

```go
func main() {
	mg := &MyGame{}
	RunGame(mg) // Execute Update() Method

	mgwd := &MyGameWithDraw{}
	RunGame(mgwd) // Execute Update() and Draw() Method
}
```

## References
- [2020/02/09 - github.com/hajimehoshi/ebiten/run.go](https://github.com/hajimehoshi/ebiten/blob/master/run.go#L313)