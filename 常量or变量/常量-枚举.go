package main

import "fmt"    //iota有递增效果

/*const (
	Connected = 0  表示连接成功
	Disconnected = 1  表示连接断开
	Unknown = 2  未知）*/

const (
	a = iota                //a==0
	b                       //b==1,隐式使用iota关键字，实际等同于b=iota
	c                       //c==2,实际等同于c=iota
	d,e,f = iota,iota,iota  //d=3,e=3,f=3,同一行值相同，此处不能只写一个iota
	g = iota                //g==4
	h = "h"                 //h=="h"，单独赋值，iota依旧递增为5
	i                       //i=="h",默认使用上面的赋值，iota依旧递增为6
	j = iota                //j==7
)

const z = iota              //每个单独定义的const常量中，iota都会重置，此时z==0

func main()  {
	fmt.Println(a,b,c,d,e,f,g,h,i,j,z)

}