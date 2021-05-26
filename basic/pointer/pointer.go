package pointer

import "fmt"

/*
	指针地址
	指针类型
	指针取值

符号：
& 取地址
* 根据地址取值
如：
p := &v // 获取v的地址
*p 获取地址对应的值
*p = 新值 // 对地址p对应的位置进行重新赋值


*/
func TestPointer() {
	fmt.Println("Pointer")

	a := 10
	b := &a
	// a := 10, &a := 0xc00000a098
	fmt.Printf("a := %v, &a := %p\n", a, &a)
	// b值 := 10, b := 0xc00000a098, &b :=0xc000006030
	fmt.Printf("b值 := %v, b := %p, &b :=%p\n", *b, b, &b)

	/*
			简单总结
		  	1.对变量进行取地址（&）操作，可以获得这个变量(a)的指针变量(b)。
			2.指针变量的值是指针地址(即b，同&a)。
		    3.对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值（即*b得到10）。
	*/

}

/*
	指针传值和copy传值
*/
// 传递一个普通变量
func transCopy(x int) {
	x = 100
}

// 传递一个指针参数
func transPointer(x *int) {
	//x = 100 // 由于x是一个指针地址，不能直接这样赋值
	// 应该根据指针地址取值在修改
	*x = 100
}

func RunTrans() {
	x := 10
	transCopy(x)
	fmt.Println("copy传值x :=", x) // copy传值x := 10
	transPointer(&x)
	fmt.Println("pointer传值x :=", x) // Pointer传值x := 100
}

/*
指针类型
定义一个int类型的指针变量
*/

func PointerType() {
	// 定义后灭有内存空间，不能直接赋值
	//var a *int
	// 需要先初始化内存空间
	// new方法返回的是a的指针类型：func new(Type) *Type
	var a = new(int)
	// 然后在赋值
	*a = 10
	fmt.Println(*a)
}
