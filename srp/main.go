package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type (
	Journal struct {
		entries []string
	}
	Persistence struct {
		lineSeparator string
	}
)

var entryCount = 0

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func (p *Persistence) LoadFromFile(j *Journal, filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	j.entries = strings.Split(string(content), p.lineSeparator)
}

func main() {
	j := Journal{}
	j.AddEntry("Example journal one")
	j.AddEntry("Example journal two")
	fmt.Println(j.String())

	p := Persistence{"\r\n"}
	p.SaveToFile(&j, "journal.txt")

	f := Journal{}
	p.LoadFromFile(&f, "journal.txt")
	fmt.Println(f.String())
}
