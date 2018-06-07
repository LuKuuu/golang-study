package main

import (
	"fmt"
	"time"
	"strconv"
)



func main(){
	ch1 := make(chan string)

	go SendData(ch1)
	time.Sleep(2*1e9)
	go GetData(ch1)

	time.Sleep(5 * 1e9)
}

func SendData(ch chan string){
	for i := 0; i < 10; i++{
		ch <- strconv.Itoa(i)
	}

}

func GetData(ch chan string){
	time.Sleep(2e9)
	var input string
	for{
		input = <- ch
		fmt.Println(input)
	}
}