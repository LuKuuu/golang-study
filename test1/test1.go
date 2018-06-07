package main

import (
	"fmt"
	"reflect"
)

var i = make([]int, 2, 10)



func main(){

	a := make([]byte,2,2)

	a = []byte{1,2}

	fmt.Println(a)

	a = append(a, a...)

	fmt.Println(a)

	/*	for j := 0;j < len(i); j++{
		fmt.Println(i[j])
	}
	word := "witch"
	fmt.Print(word)
	c := []byte(word)
	c[0] = 'b'
	word =string(c)
	fmt.Print(" becomes ")
	fmt.Println(word)

	MapTest()*/

	var cat1 cat
	cat1.name = "name"
	cat1.age = 9
	cat1.master = "master"

	PrintType(cat1)
	cat1.Print()


}

func Replace(word string, index int, letter string){
}

func MapTest(){
	m := make(map[int] string)
	m[124]="Chengyang Li"
	m[409]="Kun Qina"
	fmt.Println(len(m))

}

type cat struct{
	name string "The name of cat"
	age int "the age of cat"
	master string "the name of master"
}

func CreateCat(){
	var cat1 cat
	cat1.name = "name"
	cat1.age = 9
	cat1.master = "master"

	var cat2 *cat
	cat2 = new(cat)
	cat2.name = "cat2"
	cat2.age = 2
	cat2.master = "master2"
}

func ReflectType(cat0 cat, index int){
	cat0Type := reflect.TypeOf(cat0)
	index_field := cat0Type.Field(index)
	fmt.Printf("%v\n",index_field.Tag)
}

func PrintType(cat0 cat){

	for i :=0; i<3; i++{
		ReflectType(cat0, i)
	}
}

func (this cat) Print(){
	for i :=0; i<3; i++{
		ReflectType(this, i)
	}
}



