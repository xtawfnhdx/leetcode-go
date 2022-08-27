package main

type color byte

const (
	black color = iota
	red
	yello
)

func test(c color) {
	println(c)
}
func main() {

	switch x := 14; x {
	default:
		println("defautl")
	case 0:

	case 1:
		println("23")
	case 2:
		println("this 2")
	}

	//test(red)
	//const m = 100
	//test(m)        //只要不超出枚举类型对应类型的取值范围，就都是可以的
	//c := 11_22_234 //可以用下杠连接数字，是长数字可读性更高
	//println(c)

}
