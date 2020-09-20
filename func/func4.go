package main

func main() {

	//值类型：基本数据类型int系列、float系列、bool、string、数组和结构体struct
	//	值类型默认是值传递，变量直接存储值，内存通常在栈中分配
	//引用类型：指针、slice切片、map、管道chan、interface等都是引用类型
	//	引用类型默认是引用传递，变量存储的是一个地址，这个地址对应的空间才真正存储数据，内存通常在
		//堆上分配，当没有任何变量引用这个地址时，该地址对应的数据空间就成为一个垃圾，由GC回收

	//如果需要函数内的变量能够修改函数外的变量，那么可以传入变量的地址&，函数内以指针的方式操作变量，从效果上类似引用


}
