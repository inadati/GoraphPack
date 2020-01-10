package Greeter

import (
	"log"
)

type Greeter struct {}

func Summon() *Greeter {
	return &Greeter{}
}

func (self *Greeter) Hello() {
	log.Println("Hello Servant!!")
}