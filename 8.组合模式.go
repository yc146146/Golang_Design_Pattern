package main


type Component interface {
	Parent() Component
	SetParent(component Component)
	Name()string
	SetName(string)
	AddChild(Component)
	Print(string)
}
