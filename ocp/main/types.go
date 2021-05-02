package main

type (
	Color   int
	Size    int
	Product struct {
		name  string
		color Color
		size  Size
	}
	Filter           struct{}
	AndSpecification struct {
		first, second Specification
	}
	ColorSpecification struct {
		color Color
	}
	SizeSpecification struct {
		size Size
	}
)
