package main

import "fmt"

//go语言的字符串的字节使用utf-8编码标识unicode文本，go的字符串是由单个字节连接起来的
func main() {
	var address string = "北京长城 110 hello world"
	fmt.Println(address)

	//字符串一旦赋值了，就不能再修改了，在go中，字符串是不可变的
	//address[1] = '哈'

	//字符串的两种表示形式
	//双引号，会是识别转义字符
	//反引号，以字符串的原生字符输出，包含行和特殊字符，可以实现防止攻击、输出源代码等效果
	var str string = `你好哈 
	嘿嘿 ""   \n what！`
	fmt.Println(str)

	//字符串的拼接
	var helloworld = "hello" + "world"
	helloworld += "welcome"
	fmt.Println(helloworld)

	//当一个拼接的操作很长时，分行写，，+号要留到上一行的末尾，不能放在下一行的开头
	var str2 = "hello" +
		"world" +
		"welcome"
	fmt.Println(str2)
}