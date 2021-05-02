package isp

import "fmt"

type (
	PrinterMachine struct{}
	ScannerMachine struct{}
	MultifunctionalMachine struct{}
)

func (m PrinterMachine) Print(d Document) {
	fmt.Println("Printing...", d.Content)
}

func (m ScannerMachine) Scan(d Document) {
	fmt.Println("Scanning... ", d.Content)
}

func (m MultifunctionalMachine) Print(d Document) {
	fmt.Println("Printing...", d.Content)
}

func (m MultifunctionalMachine) Scan(d Document) {
	fmt.Println("Scanning... ", d.Content)
}

func main() {
	var p Printer
	var s Scanner
	var mfd MultifunctionalMachine

	d := Document{"My document content"}
	p = PrinterMachine{}
	p.Print(d)

	s = ScannerMachine{}
	s.Scan(d)

	mfd = MultifunctionalMachine{}
	mfd.Print(d)
	mfd.Scan(d)
}
