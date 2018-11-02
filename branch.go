package main

import (
	"io/ioutil"
	"fmt"

)

//switch后面可以不带表达式
func grade(score int) string{
	g:=""
	switch {
	case score < 0 || score >100:
		panic(fmt.Sprintf(
			"Wrong score :%d",score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}


func main() {
	const filename = "abc.txt"
	//contents,err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//}else{
	//	fmt.Printf("%s\n",contents)
	//}

	//IF后面可以跟多个语句的写法
	if contents,err := ioutil.ReadFile(filename);err != nil{
		fmt.Println(err)
	}else{
		fmt.Printf("%s\n",contents)  //这里的contents变量只在IF语句范围内有效
	}

	fmt.Println(
		grade(0),
	grade(80),
	grade(90))

}
