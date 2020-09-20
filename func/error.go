package main

import "fmt"

/*在默认情况下，当发生错误（Panic）后，程序就会退出了
如果希望，当发生错误后，可以捕获到错误，并进行处理，保证程序继续执行

go 语言不支持传统的try catch finally

defer panic recover

go中可以抛出一个Panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理
*/
func test1() {
	defer func() {
		err := recover()//内置函数，捕获到程序错误
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	fmt.Println(num1 / num2)
}

func main() {
	test1()
	fmt.Println("hello world")
}