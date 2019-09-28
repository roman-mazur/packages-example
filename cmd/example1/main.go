package main

import (
	"github.com/roman-mazur/packages-example/tools"
	"log"
)

type Dog struct {
	Logger *log.Logger
}

func (d *Dog) Bark() {
	d.Logger.Println("Bark")
}

type Duck struct {
	Logger *log.Logger
}

func (d *Duck) Fly() {
	d.Logger.Println("Fly")
}

func main() {
	dog := &Dog{Logger: tools.NewLogger("dog")}
	duck := &Duck{Logger: tools.NewLogger("duck")}

	dog.Bark()
	duck.Fly()
}
