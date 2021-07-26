package main

import (
	"fmt"
	"sync"
)

func calulateFrequencyInString(inputByte byte, inputString string,frequencyChannel chan int)  {
	sum:=0
	for i:=0;i<len(inputString);i++{
		if inputString[i]==inputByte {
			sum++
		}
	}
	frequencyChannel <-sum
}

func calculateFrequency(inputByte byte,setOfStrings []string ,wg *sync.WaitGroup) {
	defer wg.Done()
	lengthOfStringSet := len(setOfStrings)
	sum:=0
	for i:=0;i<lengthOfStringSet;i++{
		frequencyChannel := make(chan int,1)
		go calulateFrequencyInString(inputByte,setOfStrings[i],frequencyChannel)
		value := <-frequencyChannel
		sum += value
	}
	fmt.Printf("%c : %d\n",inputByte,sum)
}
func main() {

	fmt.Println("Input number of strings")
	var numOfStrings int
	fmt.Scanf("%d", &numOfStrings)
	setOfStrings := []string{}

	var wg sync.WaitGroup


	for i:=0; i<numOfStrings;i++{
		var s string
		fmt.Scanf("%s",&s)
		setOfStrings = append(setOfStrings,s)
	}

	for c := 'a';c<='z';c++{
		wg.Add(1)
		go calculateFrequency(byte(c),setOfStrings,&wg)
	}
	wg.Wait()

}

