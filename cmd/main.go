package main

import (
	"fmt"
	. "xorm_demo/model"
)

func main() {

	user := GetUserByName("uuu")
	fmt.Println(user)
}
