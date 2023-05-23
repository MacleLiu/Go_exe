// 表单的有状态排序
// 尚未完成
package main

import (
	"fmt"
	"sort"
)

type Book struct {
	Name   string
	Artist string
	Year   int
}

type byFuncs func(i, j int) bool

type bookSlice struct {
	lists     []*Book
	lessFuncs []byFuncs
}

func (b bookSlice) Len() int           { return len(b.lists) }
func (b bookSlice) Swap(i, j int)      { b.lists[i], b.lists[j] = b.lists[j], b.lists[i] }
func (b bookSlice) Less(i, j int) bool { return false }

func (b bookSlice) byName(i, j int) bool   { return b.lists[i].Name < b.lists[j].Name }
func (b bookSlice) byArtist(i, j int) bool { return b.lists[i].Artist < b.lists[j].Artist }
func (b bookSlice) byYear(i, j int) bool   { return b.lists[i].Year < b.lists[j].Year }

func main() {
	var books = []*Book{
		{"The Go Programming Language", "Alan and Brian", 2017},
		{"Data Structures", "Mark", 2004},
		{"Python", "Tom", 2017},
		{"The C Programming Language", "Jerry", 2010},
		{"Network", "Tom", 2022},
	}
	var bs bookSlice
	bs = bookSlice{books, []byFuncs{bs.byName, bs.byArtist, bs.byYear}}
	sort.Sort(bs)
	for _, b := range bs.lists {
		fmt.Println(b.Name, b.Artist, b.Year)
	}
}
