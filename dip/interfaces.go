package main

type RelationshipBrowser interface {
	FindAllChildrenOn(name string) []*Person
}
