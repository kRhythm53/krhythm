package main

import (
	"fmt"
	"sync"
)

var frequency = make(map[byte]int)
func calulateFrequencyInString(c byte, s string) int{
	sum:=0
	for i:=0;i<len(s);i++{
		if s[i]==c {
			sum++
		}
	}
	return sum
}

func calculateFrequency(c byte,setOfStrings []string ,wg *sync.WaitGroup) {
	defer wg.Done()
	lengthOfStringSet := len(setOfStrings)
	sum:=0
	for i:=0;i<lengthOfStringSet;i++{
		sum += calulateFrequencyInString(c,setOfStrings[i])
	}
	fmt.Printf("%c : %d\n",c,sum)
}
func main() {

	fmt.Println("Input number of strings")
	var numOfStrings int
	fmt.Scanf("%d", &numOfStrings)
	setOfStrings := []string{}

	var wg sync.WaitGroup
	wg.Add(26)

	for i:=0; i<numOfStrings;i++{
		var s string
		fmt.Scanf("%s",&s)
		setOfStrings = append(setOfStrings,s)
	}
	var c byte
	for c = 'a';c<='z';c++{
		go calculateFrequency(c,setOfStrings,&wg)
	}
	wg.Wait()

}

