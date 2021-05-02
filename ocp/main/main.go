package main

import (
	"fmt"
	"solid-go/ocp"
)

const (
	red ocp.Color = iota
	green
	blue
)

const (
	small ocp.Size = iota
	medium
	large
)

func (a ocp.AndSpecification) IsSatisfied(p *ocp.Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

func (f *ocp.Filter) Filter(products []ocp.Product, spec ocp.Specification) []*ocp.Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}

	return result
}

func (c ocp.ColorSpecification) IsSatisfied(p *ocp.Product) bool {
	return p.color == c.color
}

func (s ocp.SizeSpecification) IsSatisfied(p *ocp.Product) bool {
	return p.size == s.size
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	f := Filter{}
	greenSpecification := ColorSpecification{green}
	for _, v := range f.Filter(products, greenSpecification) {
		fmt.Printf("%s is green ", v.name)
	}

	sizeSpecification := SizeSpecification{large}
	lgSpec := AndSpecification{greenSpecification, sizeSpecification}
	for _, v := range f.Filter(products, lgSpec) {
		fmt.Printf(" %s is large and green ", v.name)
	}
}
