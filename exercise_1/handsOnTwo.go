package main

import "fmt"

type person struct{
	fname string
	lname string
}

type secretAgent struct{
	person 
	licenseToKill bool
}

func (p person) pSpeak(){
	fmt.Println(p.fname+p.lname)
}
func (sa secretAgent) saSpeak(){
	fmt.Println(sa.fname+sa.lname+string(sa.licenseToKill))
}

func main (){
	p := person{
		"noboru",
		"nakahara",
	}
	sa := secretAgent{
		person{
			"nanami",
			"aikawa",
		},
		true
	}
}