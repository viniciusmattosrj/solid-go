package main

type Relationship int

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

type Relationships struct {
	relations []Info
}

type Research struct {
	relationships Relationships
}

