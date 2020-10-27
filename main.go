package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	fileA, errA := os.Create("ascendente.txt")
	fileD, errD := os.Create("descendente.txt")
	if errA != nil {
		fmt.Println(errA)
		return
	}

	if errD != nil {
		fmt.Println(errD)
		return
	}

	defer fileA.Close()
	defer fileD.Close()

	var v int
	var t string
	var s []string
	fmt.Println("Cuantas strings quieres ingresar?")
	fmt.Scanln(&v)
	fmt.Println("Ingresa las ", v , " strings")
	for i := 0; i < v; i++ {
		fmt.Scanln(&t)
		s = append(s, t)
	}

	sort.Strings(s)
	for i := 0; i < v; i++ {
		fileA.WriteString(s[i] + "\n")
	}

	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	for i := 0; i < v; i++ {
		fileD.WriteString(s[i] + "\n")
	}

}
