package out

import (
	"fmt"
	"os"
)

func Fatalln(a ...interface{}) {
	fmt.Println(a...)
	os.Exit(0)
}

func Fatalf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
	os.Exit(0)
}
