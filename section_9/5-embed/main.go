package main

import (
	"embed"
	"fmt"
	"log"
)

// enterprise application in Go
// ----------------------
var name = "Russel"

//go:embed public
var public embed.FS

func main() {

	//dir := "public"
	//fileName := "data.txt"
	//path := filepath.Join(dir, fileName)

	// MkdirAll => create a dir
	//err := os.MkdirAll(dir, 0755)
	//if err != nil {
	//	fmt.Println("Papka yaratishda xato:", err)
	//	log.Fatal(err)
	//	return
	//}

	// create a file
	//file, err := os.Create(path)
	//if err != nil {
	//	fmt.Println("Fayl yaratishda xato:", err)
	//	return
	//}
	//ish tugagach fileni o'chirish
	//defer file.Close()

	// write content inside the data.txt file
	//	content := `
	//	Go,
	//	Js,
	//	Ts,
	//	Ruby,
	//	Python,
	//	Java,
	//	C++,
	//	C#,
	//	Php,
	//	Swift,
	//`
	//
	//	_, err = file.WriteString(content)
	//	if err != nil {
	//		fmt.Println("Content yozishda xatolik", err)
	//		return
	//	}

	//fmt.Printf("Fayl muvaffaqiyatli yaratildi: %s\n", path)

	// read file
	data, err := public.ReadFile("public/data.txt")
	if err != nil {
		fmt.Println("Faylni o'qishda xatolik", err)
		log.Fatal(err)
	}

	fmt.Println(string(data))

	fmt.Println(name)
}
