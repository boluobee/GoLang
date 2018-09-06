package main

import initTest "initTest"
import t3Test "t3Test"
import (
	"math"
	"fmt"
	"os"
	"runtime"
	"regexp"
)

type (
	IZ int
	FZ float64
	STR string
)

const (
	Monday=iota+1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func init() {
	println("t2Test init")
}

func main() {
	println("t2Test main")

	//typeofFunction()
	//constFunction()
	//variableFunction()

	// var a float64 =6.7345
	// var b int=IntFromFloat64(a)
	// fmt.Printf("IntFromFloat64 return %d",b)

	RegexMatch()
}

func typeofFunction(){
	var c FZ
	var a int
	a=5
	c=FZ(a)
	println(a)
	println(c)
}

func constFunction(){
	var i int
	var j string
	var z float64
	i=1
	j="abc"
	z=2
	println(i+10)
	println(z+10)
	println(j+"fff")

	println(Monday)
	println(Tuesday)
	println(Wednesday)
	println(Thursday)
	println(Friday)
	println(Saturday)

}

func variableFunction(){
	var (
		goos=runtime.GOOS
		gopath=os.Getenv("GOPATH")
	)

	fmt.Printf("this is environment :  %s \n",goos)
	fmt.Print("this is environment : %s \n",gopath)
}

func test(){
	initTest.Test()
	t3Test.Test()
} 

func IntFromFloat64(x float64)int{
	if math.MinInt32<=x&&math.MaxInt32>=x{
		whole,fraction:=math.Modf(x)
		if fraction>=0.5{
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range",x))
}

func RegexMatch(){
	reg,err:=regexp.Compile("([a-zA-Z]{1})-*(\\d{4})$")
	phone:=reg.FindSubmatch([] byte("m1811-C-3100"))

	for i:=0;i<len(phone);i++{
		fmt.Printf("i %d phone %s err %s",i,phone[i],err)
	}
	
}