package main

import "fmt"

//拷贝副本  拷贝原有的对象

//原型对象需要实现的接口
type Cloneable interface {
	Clone() Cloneable
}

//原型对象的类
type PrototypeManger struct {
	prototypes map[string]Cloneable
}

//构造初始化
func NewPrototypeManger()*PrototypeManger{
	return &PrototypeManger{make(map[string]Cloneable)}
}

//抓取
func (p*PrototypeManger)Get(name string)Cloneable{
	return p.prototypes[name]
}

//设置
func (p*PrototypeManger)Set(name string, prototype Cloneable){
	p.prototypes[name]=prototype
}

type Type1 struct{
	name string
}

func (t*Type1)Clone()Cloneable{
	//开辟内存新建变量 存储指针指向内容
	//tc:=*t
	//return &tc
	return t
}

type Type2 struct{
	name string
}

func (t*Type2)Clone()Cloneable{
	tc:=*t
	return &tc
}


func main() {
	mgr := NewPrototypeManger()
	t1 := &Type1{name:"type1"}
	mgr.Set("t1", t1)
	t11 := mgr.Get("t1")
	t22 := t11.Clone()
	if t11 == t22{
		fmt.Println("浅复制")
	}else{
		fmt.Println("深复制")
	}



}