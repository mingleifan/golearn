package main

import (
	"golearn/com/25/4-init/lib1"
	mylib2 "golearn/com/25/4-init/lib2"
	//匿名导入包但并不使用 -> [_ "golearn/com/25/4-init/lib2"]
	//导入包别名 -> [myLib2 "golearn/com/25/4-init/lib2"]
	//导入包为当前main包的方法 -> [. "golearn/com/25/4-init/lib2"] -> 就不需要 包名.func, 直接使用func ex:尽量不使用
)

func main() {
	lib1.FuncLib1()
	mylib2.FuncLib2()
}
