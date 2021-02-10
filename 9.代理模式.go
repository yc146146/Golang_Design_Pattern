package main

import "fmt"

type Subject interface {
	//实际业务
	Do()string

}

type RealSubject struct {

}


func (sb RealSubject)Do()string{
	return "合约"
}

type Proxy struct {
	real RealSubject
	money int
}

func (p Proxy)Do()string{
	if p.money>0{
		return p.real.Do()
	}else{
		return "请充值"
	}
}

func main() {
	var sub Subject
	sub = &Proxy{money:1000}
	fmt.Println(sub.Do())
}
