package main

import (
	"fmt"
	"io"
	"os"
)

func func1(s string) string {
	s1 := s
	if s == "world" {
		return "if"
	} else {
		return "else"
	}
	return s1
}

//func CopyFile(dstName, srcName string) (written int64, err error) {
//	src, err := os.Open(srcName)
//	if err != nil {
//		fmt.Println("srcName err!")
//		return
//	}
//
//	dst, err := os.Create(dstName)
//	if err != nil {			// 如果err != nil直接return了,但是src资源没有被释放掉, 可以使用defer避免这种问题
//		fmt.Println("dstName err!")
//		return
//	}
//
//	written, err = io.Copy(dst, src)
//	dst.Close()
//	src.Close()
//	return
//}

func CopyFileWithDefer(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()	// 使用defer,

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

// 规则一 当defer被声明时，其参数就会被实时解析
func DeferTest1() {
	fmt.Println("DeferTest1")
	a := 0
	defer fmt.Println(a) 	// defer函数会在return之后被调用,但是这里打印的a为0
							// 因为这个变量(a)在defer被声明的时候，就已经确定其确定的值了。
	a++
	return
}

// 规则二 defer执行顺序为先进后出
func DeferTest2() {
	fmt.Println("DeferTest2")
	defer fmt.Println()
	for i := 0; i < 10; i++ {
		defer fmt.Print(i, " ")  // Possible resource leak, 'defer' is called in a for loop less... (Ctrl+F1)
							//Inspection info: Reports defer statements inside loops.
	}
}

// 规则三 defer可以读取有名返回值
func DeferTest3() (i int) {
	fmt.Println("DeferTest3")
	func() {
		i = i + 5
		fmt.Println(i) // 输出为5, 此处i=0
		}()
	defer func() {
			i = i + 7
			fmt.Println("defer", i)	// 输出为8, 因为defer在return后调用,所以此处i为1
		} () // 执行+7操作
							     // 在return后i + 7使返回值为8 ?
	return 1
}

func DeferTest4() {
	println("DeferTest4")
	arr := [5]string{"Go", "Java", "JavaScript", "c++", "Python"}
	for _, value := range arr {
		defer func() {
			fmt.Printf("%s -->",value)
		}()	// defer先解析,最后value值为"Python",所以循环打印"Python"
	}
	fmt.Println("从我开始执行")
}

func DeferTest5() {
	fmt.Println("\nDeferTest5")
	i:=0
	defer func(i int) {
		for ;i<5 ; i++ {
			fmt.Printf("%d-->",i)
		}
	}(i) // 在匿名函数中传入的参数i在defer函数未执行之前，就已经得到了保存。
	func(i int) {
		for; i<5; i++ {
			fmt.Printf("%d-->", i)
		}
	}(i)
	fmt.Println()
	i++
}

func main() {
	defer fmt.Println("\ndefer func1 out: ", func1("world"))	// 先进后出
	fmt.Println("hello, world!")
	//CopyFile("/home/alex/Desktop/b.txt", "/home/alex/Desktop/a.txt")
	CopyFileWithDefer("/Users/chriszhu/WorkPlace/go/issue/copy.txt",
				      "/Users/chriszhu/WorkPlace/go/issue/defer.go")
	DeferTest1()
	DeferTest2()
	fmt.Println(DeferTest3())
	DeferTest4()
	DeferTest5()
}