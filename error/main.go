package main

import "fmt"

func main() {
	defer func() {
		i := recover()
		if i != nil {
			fmt.Printf("发生错误了%s", "panic")
		}
	}()
	panicTest("")
	empInterface("1111")
}

/*
自定义错误
*/
func errorTest(p string) (a string, b error) {
	if len(p) == 0 {
		return "", fmt.Errorf("长度为%d", 0)
	}
	return p, nil
}

/*
断言
*/
func empInterface(a interface{}) {
	if s, ok := a.(string); ok {
		fmt.Printf("传入的值是%s", s)
	} else {

	}
}

func panicTest(s string) {
	if len(s) == 0 {
		panic("发生错误")
	}
	fmt.Printf("%s", "panic之后的代码")
}
