# go-interface-practice
 golang 語言：一個典型使用 interface 的時機

## 使用私有（private）struct 時，可使用公有（public）interface 來接收
- [001ex-PrivateStruct](./001ex-PrivateStruct)

## 使用 Interface 來確認 Struct 是否有實現指定的方法
- [002ex-CheckFunc](./002ex-CheckFunc)

## References
- [2020/10/22 - golang 補遺：指標和介面](https://mp.weixin.qq.com/s/J_eW_O8AP-on_0DAJ6P1sw)
    - interface 有自己對應的實體數據結構
    - 盡量不要用指標去指向 interface
- [2020/09/14 - 在 Go 語言中，我為什麼使用介面](https://mp.weixin.qq.com/s/AMgCzCG_096iaCdtOJIBBA)
- [2020/08/01 - [Go] 為什麼 Pointer Receiver 不能使用 Value Type 賦值給 Interface Value](https://mileslin.github.io/2020/08/Golang/為什麼-Pointer-Receiver-不能使用-Value-Type-賦值給-Interface-Value/)
- [2020/02/09 - github.com/hajimehoshi/ebiten/run.go](https://github.com/hajimehoshi/ebiten/blob/master/run.go#L313)