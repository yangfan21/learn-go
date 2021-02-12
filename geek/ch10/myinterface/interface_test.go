package myinterface

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHello("Go"))
}

type Programmer interface {
	WriteHello(str string) string
}
type GoProgrammer struct {
}

func (p GoProgrammer) WriteHello(str string) string {
	return fmt.Sprintf("Hello %s Programmer", str)
}
