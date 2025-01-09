package main

import (
	"fmt"
	"os"
)

func main() {
	myInput := os.Getenv("INPUT_MYINPUT")

	output := fmt.Sprintf("Hola %s", myInput)

	fmt.Println(fmt.Sprintf(`::set-output name=myOutput::%s`, output))
}
