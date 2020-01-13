package main

type Region struct {
	region string
	Hotels []Hotel
}

type Hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
}
