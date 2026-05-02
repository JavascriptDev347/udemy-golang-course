package main

func main() {

	a := 10
	for i := 1; i < 10; i++ {
		a *= i
		println(a)
	}

	// while style loop
	k := 3
	for k > 0 {
		println(k)
		k--
	}
}
