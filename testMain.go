package main

import initTest "initTest"
import t3Test "t3Test"
import (
	"fmt"
	"os"
	"runtime"
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

	typeofFunction()
	constFunction()
	variableFunction()
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