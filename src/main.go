package main

import (
	"calendar/src/utils"
	"fmt"
)

func main() {
	utils.LoadEnv()
	t, _ := utils.CreateToken(4)
	r, _ := utils.ParseToken(t)
	fmt.Println(r)
}
