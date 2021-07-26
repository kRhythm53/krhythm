package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func getRating(wg *sync.WaitGroup,rating chan float64) {
	defer wg.Done()
	randomRating := 10*rand.Float64()
	randomTime := rand.Intn(100)
	for i:=0;i<randomTime;i++{
		time.Sleep(time.Microsecond)
	}
	rating <- randomRating
}

func main(){
	rand.Seed(time.Now().UnixNano()) //to get different pseudo random numbers
	var averageRating float64
	var wg sync.WaitGroup
	wg.Add(200)
	for i:=1;i<=200;i++{
		rating :=  make(chan float64,1)
		getRating(&wg,rating)
		ratingValue := <- rating
		averageRating += ratingValue

	}
	wg.Wait()
	averageRating = averageRating/float64(200)
	fmt.Println(averageRating)
}