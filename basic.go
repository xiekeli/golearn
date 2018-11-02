package main

import (
	"fmt"
	"math"
)

var aa = 3   //这不是全局变量，是包内部变量
var ss = "kk" //在函数外面不能用：=来定义

var(
	aaa = 33
	sss = "KKK"
)

func variableZeroValue(){
	var a int
	var s string
	fmt.Printf("%d %q\n",a,s)
}

func variableInitialValue(){
	var a, b int = 3,4
	var s string = "abc"
	fmt.Println(a,b,s)
}

func variableTypeDeduction(){
	var a,b,c,s = 3,4,true,"def"
	fmt.Println(a,b,c,s)
}

func variableShorter(){
	a,b,c,s := 3,4,true,"def"
	b =5
	fmt.Println(a,b,c,s)

}

func consts(){
	const (
		filename = "abc.txt"  //GOlang中的常量不用大写，就按普通变量的命名规则即可
	    a,b = 3,4
	)
	var c int
	c= int(math.Sqrt(a*a + b*b))
	fmt.Println(filename,c)

}

func enums(){
	const(
		cpp = 0
		java =1
		python = 2
		golang = 3
	)
	fmt.Println(cpp,java,python,golang)
}

func enums2(){
	const(
		cpp = iota  //表示下面都是递增的
		_           //跳过
		python
		golang
		javascript
	)
	const(
		b = 1<<(10*iota) // 按定义的公式进行递增
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp,javascript,python,golang)
	fmt.Println(b,kb,mb,gb,tb,pb)
}

func main() {
	fmt.Println("Hello world");
	variableInitialValue()
	variableZeroValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa,ss)

	consts ()
	enums()
	enums2()
}
