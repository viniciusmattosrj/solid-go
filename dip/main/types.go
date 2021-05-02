package main

type (
	Relationship int
	Person struct {
		name string
	}
	Info struct {
		from         *Person
		relationship Relationship
		to           *Person
	}
	Relationships struct {
		relations []Info
	}
	Research struct {
		browser RelationshipBrowser
	}
)
