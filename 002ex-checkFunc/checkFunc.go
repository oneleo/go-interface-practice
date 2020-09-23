package main

import (
	"fmt"
)

type GameInterface interface {
	Update()
}

type MyGame struct {
}

func (mg *MyGame) Update() {
	fmt.Println("Execute MyGame.Update() Method")
}
func (mg *MyGame) Draw() {
	fmt.Println("Execute MyGame.Draw() Method")
}

type DumperGameWithDraw struct {
	gameInterface GameInterface
}

func (dgwd *DumperGameWithDraw) Update() {
	fmt.Println("Execute DumperGameWithDraw.Update() Method")
}
func (dgwd *DumperGameWithDraw) Draw() {
	fmt.Println("Execute DumperGameWithDraw.Draw() Method")
	dgwd.gameInterface.(interface{ Draw() }).Draw()
}

type DumperGame struct {
	gameInterface GameInterface
}

func (dg *DumperGame) Update() {
	fmt.Println("Execute DumperGameWithDraw.Update() Method")
}

func runGame(gi GameInterface) {
	//???
}

func chkFunc(gi GameInterface) {
	if _, ok := gi.(interface{ Draw() error }); ok {
		runGame(&DumperGameWithDraw{gameInterface: gi})
	}
	runGame(&DumperGame{gameInterface: gi})
}

func main() {
	mg := &MyGame{}
	chkFunc(mg)
}
