package main

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}