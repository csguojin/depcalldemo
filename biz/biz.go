package biz

import (
	"fmt"

	"github.com/csguojin/depcalldemo/biz/biz3"
	"github.com/spf13/viper"
)

func Hello() {
	fmt.Println("biz.Hello")
}

func BizTest2() {
	fmt.Println("biz.test.2")
	viper.GetString("test")
	biz3.Biz3Test1()
	var s S
	s.Key()
	s.Print()
}

type S struct {
	Name string
}

func (s S) Key() {
	fmt.Println("S.Key")
}

func (s *S) Print() {
	fmt.Println("S.Key")
}
