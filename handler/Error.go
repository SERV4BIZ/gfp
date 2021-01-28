package handler

import (
	"fmt"
	"log"
	"runtime"
)

// Error is check error and print it to console
func Error(err error) bool {
	if err != nil {
		pc, fn, line, _ := runtime.Caller(1)
		fmt.Println("* * * * * * * * * * * * * * * * * * * *")
		log.Println("[ ERROR ]")
		fmt.Println("* * * * * * * * * * * * * * * * * * * *")
		fmt.Println(fmt.Sprint("[ Func ] : ", runtime.FuncForPC(pc).Name()))
		fmt.Println(fmt.Sprint("[ File ] : ", fn))
		fmt.Println(fmt.Sprint("[ Line ] : ", line))
		fmt.Println(fmt.Sprint("[ Mess ] : ", err))
		fmt.Println("* * * * * * * * * * * * * * * * * * * *")
		return true
	}
	return false
}
