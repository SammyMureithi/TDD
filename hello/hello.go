package main

import "fmt"

func main() {
	fmt.Println(Hello("world","Jop"))
}
const greeting ="Hello, "
func Hello(name string, language string) string {
	if name == ""{
		name= "World"
	}
	prefix:=greeting
	switch language{
	case "Spanish":
		prefix ="Hola, "
		
	case "French":
		prefix ="Bonjour, "

	}
	
	return prefix+name
}