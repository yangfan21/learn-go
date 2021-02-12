package extension

import (
	"fmt"
	"testing"
)

func TestGoExtension(t *testing.T) {
	dog := new(Dog)
	//dog.Speak()
	dog.SpeakTo("ABC")
}

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

// 内嵌有种继承的错觉，但不是继承，应为并不支持父类方法的重载
type Dog struct {
	Pet
	//p *Pet
}

//func (d *Dog) Speak() {
//	//d.Speak()
//	fmt.Print("Dog")
//}
//
//func (d *Dog) SpeakTo(host string) {
//	d.Speak()
//	fmt.Println(" ", host)
//}
