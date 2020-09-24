package main

import (
	"fmt"
)

// ---------- 介面 ----------

// IGame 需實現 Update() 方法
type IGame interface {
	Update()
}

// ---------- 第一個結構 ----------

// MyGame 只實現 Update() 方法
type MyGame struct {
}

// Update method of MyGame
func (mg *MyGame) Update() {
	fmt.Println("Execute MyGame.Update() Method")
}

// ---------- 第二個結構 ----------

// MyGameWithDraw 有實現 Update() 及 Draw() 方法
type MyGameWithDraw struct {
}

// Update method of MyGameWithDraw
func (mg *MyGameWithDraw) Update() {
	fmt.Println("Execute MyGameWithDraw.Update() Method")
}

// Draw method of MyGameWithDraw
func (mg *MyGameWithDraw) Draw() {
	fmt.Println("Execute MyGameWithDraw.Draw() Method")
}

// RunGame 會判斷要執行的 Game 是否有實現 Draw() 方法，有就會執行
func RunGame(gi IGame) {
	gi.Update()
	if game, ok := gi.(interface{ Draw() }); ok {
		game.Draw()
	}
}

// ---------- main() ----------

func main() {
	fmt.Println("第一款 Game：")
	mg := &MyGame{}
	RunGame(mg) // Execute Update() Method

	fmt.Println("\n第二款 Game with Draw：")
	mgwd := &MyGameWithDraw{}
	RunGame(mgwd) // Execute Update() and Draw() Method
}
