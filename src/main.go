package main

import (
	"calendar/src/utils"
	"fmt"
	"os"
)

func main() {
	utils.LoadEnv()
	fmt.Println(os.Getenv("PORT"))
}
