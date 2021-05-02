package main

import "fmt"

const (
	Parent Relationship = iota
	Child
	Sibling
)

func (rs *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range rs.relations {
		if v.relationship == Parent &&
			v.from.name == name {
			result = append(result, rs.relations[i].to)
		}
	}
	return result
}

func (rs *Relationships) AddParentAndChild(parent, child *Person) {
	rs.relations = append(rs.relations,
		Info{parent, Parent, child})
	rs.relations = append(rs.relations,
		Info{child, Child, parent})
}

func (r *Research) Investigate() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.name)
	}
}

func main() {
	parent := &Person{"John"}
	child1 := &Person{"Chris"}
	child2 := &Person{"Matt"}
	
	relationships := &Relationships{}
	relationships.AddParentAndChild(parent, child1)
	relationships.AddParentAndChild(parent, child2)
	research := Research{relationships}
	research.Investigate()
}
