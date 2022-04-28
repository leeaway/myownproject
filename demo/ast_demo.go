package demo

import "fmt"

type Helper struct {
	printer *Printer
}

//打印机
type Printer struct {
	Str string
}

//打印动作
func (p *Printer) Print() {
	fmt.Println(p.Str)
}
