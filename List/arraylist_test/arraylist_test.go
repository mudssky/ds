package main

import (
	"fmt"
	"testing"

	"../arraylist"
)

func TestMain(t *testing.T) {

	arrlist := arraylist.New()
	if !arrlist.Empty() {
		t.Errorf("arrlist.Empty(),experted true actual len =%v", arrlist.Empty())
	}
	if arrlist.GetLen() != 0 {
		t.Errorf("arraylist.New(),experted len=0,actual len =%d", arrlist.GetLen())
	}
	arrlist.Show()
	arrlist.Append(1)
	if arrlist.GetLen() != 1 {
		t.Errorf("arrlist.Append(1),experted len=1,actual len =%d", arrlist.GetLen())
	}
	arrlist.Show()
	arrlist.Append(2)
	arrlist.Show()
	arrlist.Append(3)
	arrlist.Show()
	arrlist.Append(4)
	arrlist.Show()
	arrlist.Insert(0, 2)
	arrlist.Insert(5, 5)
	arrlist.Show()
	if val := arrlist.Pop(); val != 5 {
		t.Errorf("arrlist.Pop(),experted val=5,actual val=%d", val)
	}
	arrlist.Erase(3)
	if arrlist.GetLen() != 4 {
		t.Errorf("arrlist.Append(1),experted len=0 actual len =%d", arrlist.GetLen())
	}
	arrlist.Erase(0)
	arrlist.Erase(0)
	arrlist.Erase(0)
	arrlist.Erase(0)
	if arrlist.GetLen() != 0 {
		t.Errorf("arrlist.Append(1),experted len=0 actual len =%d", arrlist.GetLen())
	}
	if !arrlist.Empty() {
		t.Errorf("arrlist.Empty(),experted true actual len =%v", arrlist.Empty())
	}

	for i := 0; i < 100; i++ {
		arrlist.Append(i)
	}
	arrlist.Show()
	fmt.Println(arrlist.GetLen())
}
