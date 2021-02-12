package mystruct

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestMyStruct(t *testing.T) {
	e := Employee{
		Id:   "1",
		Name: "TestA",
		Age:  12,
	}
	// 返回指针
	e1 := new(Employee)
	e1.Age = 21
	e1.Name = "Test2"
	e1.Id = "2"

	t.Logf("e1 : %T", e1)
	t.Logf("e : %T", e)
}

func (e Employee) GetName() string {
	fmt.Printf("Name Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s - Name:%s - Age :%d ", e.Id, e.Name, e.Age)
}

func (e *Employee) GetNewName() string {
	fmt.Printf("New Name Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s / Name:%s / Age :%d ", e.Id, e.Name, e.Age)
}

func TestGetName(t *testing.T) {
	e := Employee{
		Id:   "1",
		Name: "TestA",
		Age:  12,
	}
	fmt.Printf("Befor Name Address is %x\n", unsafe.Pointer(&e.Name))
	//t.Logf("get name func: %s", e.GetName())
	t.Logf("get new name func: %s", e.GetNewName())
	fmt.Printf("after Name Address is %x\n", unsafe.Pointer(&e.Name))

}
