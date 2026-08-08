package main

import "fmt"

func main() {

	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}

// Print request for mac
// Printing by a HP Printer

// Print request for mac
// Printing by a EPSON Printer

// Print request for windows
// Printing by a HP Printer

// Print request for windows
// Printing by a EPSON Printer
