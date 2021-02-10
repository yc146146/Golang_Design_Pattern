package main

import "fmt"


//解决软件可拓展性
//统一接口 链接不同工具

type API interface {
	Say(name string)string
}

func NewAPI(str string)API{
	if str == "en"{
		return &English{}
	}else if str == "cn"{
		return &Chinese{}
	}else{
		return nil
	}
}

type Chinese struct {}

func (*Chinese)Say(name string)string{
	return "你哈"+name
}

type English struct {}

func (*English)Say(name string)string{
	return "hello"+name
}

type Japanese struct {}

func (*Japanese)Say(name string)string{
	return "koniqiwa"+name
}

func main1(){
	api := NewAPI("cn")
	server := api.Say("ppop")
	fmt.Println(server)
}

func main(){
	api := NewAPI("en")
	server := api.Say("ppop")
	fmt.Println(server)
}