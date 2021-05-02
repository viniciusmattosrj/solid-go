package main

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

/* Interfaces aggregation if it makes senses */
type MultifunctionDevice interface {
	Printer
	Scanner
}
