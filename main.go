package main

import (
	"fmt"

	Cat "interfaces.com/m/v2/utils.com/Cat"
)

type I interface {
	Get() int
	Put(int)
}

type S struct{ i int }

func (p *S) Get() int  { return p.i }
func (p *S) Put(v int) { p.i = v }

type R struct{ i int }

func (p *R) Get() int  { return p.i }
func (p *R) Put(v int) { p.i = v }

type G struct{ i int }

func (p *G) Get() int  { return p.i }
func (p *G) Put(v int) { p.i = v }

// CHECKING THE ACTUAL TYPE OF THE PARAMETER
func f(p I) {
	switch p.(type) { // p.(type) only inside of switch statement
	case *S:
		fmt.Println("Type S")
		//pType := reflect.TypeOf(p) //get type of P
		//fmt.Println(pType)
	case *R:
		fmt.Println("Type R")
	default:
		fmt.Println("Unknown Type")
	}
}

func main() {
	// s := &S{i: 1}
	// r := &R{i: 2}
	// g := &G{i: 3}
	// f(g)
	// f(s)
	// f(r)

	// intarr := Sorter.Xi([]int{5, 2, 3, 1})
	// strarr := Sorter.Xs([]string{"f", "a", "b", "c"})

	//GENERIC SORTING FUNCTION USING INTERFACES
	// Sorter.Sort(intarr)
	// Sorter.Sort(strarr)
	// fmt.Println(intarr)
	// fmt.Println(strarr)

	// fmt.Println(Max.Max(intarr))

	// GENERIC SORTING FUNCTION USING REFLECTION
	// Bubblesort.Sort(intarr)
	// fmt.Println(intarr)
	// Bubblesort.Sort(strarr)
	// fmt.Println(strarr)

	// GENERIC LIST
	// llist := LinkedList.NewList()
	// llist.AddNode(1)
	// llist.AddNode(2)
	// llist.AddNode(3)
	// llist.AddNode(4)

	// llist.Print()

	//CAT Program
	Cat.Cat()

	fmt.Println("End of Program")
}
