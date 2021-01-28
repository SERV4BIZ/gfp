package handler

import (
	"fmt"
	"log"
	"runtime"
)

// Panic is check error and print it to console and stop program
func Panic(err error) {
	if err != nil {
		pc, fn, line, _ := runtime.Caller(1)
		fmt.Println("* * * * * * * * * * * * * * * * * * * *")
		log.Println("[ PANIC ERROR ]")
		fmt.Println("* * * * * * * * * * * * * * * * * * * *")
		fmt.Println(fmt.Sprint("[ Func ] : ", runtime.FuncForPC(pc).Name()))
		fmt.Println(fmt.Sprint("[ File ] : ", fn))
		fmt.Println(fmt.Sprint("[ Line ] : ", line))
		fmt.Println(fmt.Sprint("[ Mess ] : ", err))
		fmt.Println("* * * * * * * * * * * * * * * * * * * *")
		panic(err)
	}
}
