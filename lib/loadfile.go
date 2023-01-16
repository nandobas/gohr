package lib

import (
	"fmt"
	"log"
	"os"
)

func LoadFromFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Panic(err)
	}

	data := make([]byte, 100)
	if _, err := file.Read(data); err != nil {
		log.Panic(err)
	}

	fmt.Println(string(data))

	names := []string{"nando", "gaby", "alice", "luisa", "cibelle"}

	for i, name := range names {
		fmt.Printf("%d: %s \n", i, name)
	}

	names = append(names, "bernardo")

	fmt.Println(names, len(names), cap(names))

	idades := make(map[string]uint)

	idades["nando"] = 42
	idades["gaby"] = 35

	fmt.Println(idades)
}
