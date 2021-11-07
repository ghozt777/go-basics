package main
import (
	"fmt"
	"strconv"
)

// when variable are declared at the package level the colon syntax dosen't work and hence we need to use the var syntax

var message string = "Bitch try me!"
// str := "Sike u thought"  This is not allowed

var ( // another way to declare variables without using the var keyword evertime
	name = "ghozt"
	language = "c++"
	age = 22
)

var foo string = "foo package level" // package scope

var Foo string = "foo global level" // Global Scope

// Note : variables with lowercase names(starts with a lowercase letter) are local to the package and hence are unavailable to the outside world 
// Variables with Uppercase names (starts with a uppercase letter)are available to the outside world

// There are three level of Scopes in Go : 1. Block Scope 2. Package Scope 3. Global Scope


func main() {
	var i int  // one of the ways to declare a variable in go  (var keyword -> variable name -> type)
	i = 10 // Block Scope
	fmt.Println(i)
	i = 69 
	fmt.Println(i)
	var x int = 11 
	fmt.Println(x) 
	y := 12 // another way to declare a variable where the type is determined by the compiler
	fmt.Println(y)
	fmt.Printf("%v , %T" , y , y) // allows to metion the type of the variable
	fmt.Println()
	var a float32 = 69.69 // we can explicitly make a varialbe of type float32
	b := 69.96 // but this is not the case with the automatically infered type hence using this syntax we cant declare a float32 variable
	// therefore var a int = 10 this sytax provides more control
	fmt.Printf("%v , %T" , a , a)
	fmt.Println()
	fmt.Printf("%v , %T" , b , b)
	fmt.Println()
	fmt.Println(message)
	fmt.Println()
	fmt.Println(name , ":" , language , ":" , age)
	fmt.Println()
	fmt.Println(Foo)
	fmt.Println(foo) // this prints the foo variable of thepackage level
	var foo string = "foo main function level" // this shadows the foo variable from the package level
	fmt.Println(foo)
	// if there are two variables with the same name in different scopes the one with the innermost scope takes the precedence and this is known as shadowing

	// Note : If a variable in Go is declared and never used it will throw an error 

	// orphan := "I am never used" if this is never used it will throw an error

	// explicit type conersion : using conversion functions
	var e int = 13
	var c float32
	c = float32(e) // here float32() is a conversion function
	fmt.Printf("%v ,  %T" , c , c)

	// overflow 
	var d float32 = 69.6969
	// d = f Not allowed as Go sees the risk of overflow of data
	var f int = int(d) // the only way is explicit conversion
	fmt.Println()
	fmt.Printf("loss of info during conversion from : %T to %T values : %v to %v " , d , f , d , f)
	h := 43
	var str string = string(h) // VS Code shows a warning regarding the potential error

	// int to string conversion
	fmt.Println()
	fmt.Printf("conversion of %v into string yeilds : %v" , h , str) // Prints the UNICODE equivalent of the int and not the int converted into string
	// to get the desired output of converting int to string we need to use the string conversion package 'strconv'
	var strDesired = strconv.Itoa(h)
	fmt.Println() ;
	fmt.Printf("Using String conversion package %v (%T) yeilds (%T) : %v" , h , h , strDesired , strDesired)
	
}
