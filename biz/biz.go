package biz

import (
	"fmt"

	"github.com/spf13/viper"
)

func Hello() {
	fmt.Println("biz.Hello")
}

func BizTest2() {
	fmt.Println("biz.test.2")
	viper.GetString("test")
}
