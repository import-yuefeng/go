package main

import (
	"fmt"

)


func main(){

	mult_returnval(1, 2)
	mult_returnval_2(1, 2)

}

func mult_returnval(a int, b int) (x1 int, x2 int){
	c := a + b
	// 为了测试:= 语句进行初始化的操作
	x1 = c
	d := a * b
	x2 = d
	fmt.Println(x1, x2)
	return x1, x2
	// 该为返回已命名的变量
}


func mult_returnval_2(a int, b int) (int, int){
	x1 := a + b
	x2 := a * b

	return x1, x2
	// 该为返回未命名的变量


}