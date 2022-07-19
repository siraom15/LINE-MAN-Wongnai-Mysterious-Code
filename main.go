package main

import (
	base64 "encoding/base64"
	"fmt"
)

func reverse(str string) (res string) {
	for _, value := range str {
		res = string(value) + res
	}
	return
}

func main() {
	var whatIsIt string
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)
	whatIsIt = string(sd)
	fmt.Println(reverse(whatIsIt)) //Join:us:at:LINE:MAN:Wongnai
}
