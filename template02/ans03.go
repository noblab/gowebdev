package main

type Item struct {
	Name    string
	Descrip string
	Price   float32
}

type Meal struct {
	Meal  string
	Items []Item
}

type Menu []Meal
