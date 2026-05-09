package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println("Path learning in go")
	// Join for add files each other
	path1 := filepath.Join("C:", "Users", "John", "Documents", "file.txt")
	fmt.Println(path1)

	path2 := filepath.Join("config", "app.yml")
	fmt.Println(path2)

	fmt.Println("Base:", filepath.Base(path1)) // Get the base name of the 2-path}
	fmt.Println("Ext:", filepath.Ext(path1))

	dirtyDir := "./users/./dir/../other_dir/./file.txt"
	fmt.Println(filepath.Clean(dirtyDir)) // Clean
}
