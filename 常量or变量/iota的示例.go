package main

import "fmt"

func main()  {
//使用”_“跳过某些值
const (
	a1 = iota  //0
	a2         //1
	_
	a4         //3
)
 fmt.Println(a1,a2,a4)

//iota声明中间插队
const (
	n1 = iota     //0
	n2 = 100     //100
	n3 = iota   //2
	n4         //3
)
const n5 = iota //0
fmt.Println(n1,n2,n3,n4,n5)

/*定义数量级
（这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，
也就是由1变成了10000000000，也就是十进制的1024。
同理2<<2表示将2的二进制表示向左移2位，也就是由10变成了1000，也就是十进制的8。）*/

const (
	_ = iota
	KB = 1 << (10*iota)
	MB = 1 << (10*iota)
	GB = 1 << (10*iota)
	TB = 1 << (10*iota)
	PB = 1 << (10*iota)
)
    fmt.Println(KB,MB,GB,TB,PB)

//多个iota定义在一行

const (
	a,b = iota+1,iota+2  //1,2
	c,d                  //2,3
	e,f                  //3,4
)
    fmt.Println(a,b,c,d,e,f)
}

