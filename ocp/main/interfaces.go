package main

type Specification interface {
	IsSatisfied(p *Product) bool
}
