package main

import "fmt"

// defer — funksiya tugashidan OLDIN
// OXIRGI chaqiriladigan amal
//defer — funksiya tugashi bilan (return, panic, yoki oxirgi qator) avtomatik chaqiriladi.
//Uni yozgan joyda emas, funksiya tugaganida bajarilishini eslab qoling.

// LIFO stack - last in first out
func simpleDefer() {
	fmt.Println("Simple defer function started")
	defer fmt.Println("This will be printed at the end of the function")
	fmt.Println("Simple defer function ended")
}
func main() {
	defer fmt.Println("Main func with defer")
	fmt.Println("Defer")
	simpleDefer()

}
