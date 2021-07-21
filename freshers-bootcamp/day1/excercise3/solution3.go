
package main

import "fmt"

type Salary interface{
	fulltime() int
	contractor() int
	freelancer() int
}

type employees struct{
	fullTime int
	cont int
	freelance int
}
func(x employees) fulltime() int{
	return x.fullTime*500
}
func(y employees) contractor() int{
	return y.cont*100
}
func(z employees) freelancer() int{
	return z.freelance*10
}
func main(){
	var employee Salary
	employee = employees{fullTime: 28,cont: 28,freelance: 200}

	fmt.Println("fulltime employees  salary :",employee.fulltime())
	fmt.Println("contractors salary: ",employee.contractor())
	fmt.Println("freelancers salary: ",employee.freelancer())
}