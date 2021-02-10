package main

import "fmt"


//新建类	不同类执行不同操作


//实际运行类的接口
type Operator interface {
	SetLeft(int)
	SetRight(int)
	Result()int
}

//工厂接口
type OperatorFactory interface {
	Create()Operator
}

//操作数
type OperatorBase struct {
	left,right int
}

//赋值
func (op*OperatorBase)SetLeft(left int){
	op.left = left
}

func (op*OperatorBase)SetRight(right int){
	op.right=right
}

//操作的抽象
type PlusOperatorFactory struct {

}

//操作类中包括操作数
type PlusOperator struct {
	*OperatorBase
}

//实际运行
func (o PlusOperator)Result()int{
	return o.right+o.left
}
func (PlusOperatorFactory)Create()Operator{
	return &PlusOperator{&OperatorBase{}}
}

//操作的抽象
type SubOperatorFactory struct {

}

//操作类中包括操作数
type SubOperator struct {
	*OperatorBase
}

//实际运行
func (o SubOperator)Result()int{
	return o.left-o.right
}

func (SubOperatorFactory)Create()Operator{
	return &SubOperator{&OperatorBase{}}
}

func main22() {
	var fac OperatorFactory
	//fac = PlusOperatorFactory{}
	fac = SubOperatorFactory{}
	op := fac.Create()
	op.SetLeft(20)
	op.SetRight(10)
	fmt.Println(op.Result())

}