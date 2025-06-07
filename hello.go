package helloworld

import "fmt"

type HelloService struct{}

func (h HelloService) Display(){
	fmt.Println("Hello From Hello Service")

}
