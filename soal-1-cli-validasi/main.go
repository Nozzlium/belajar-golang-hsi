package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ErrInvalidInput = errors.New("input tidak valid")
var ErrUnderage = errors.New("umur tidak valid (minimal 18 tahun)")

func main() {
	fmt.Print("Nama: ")
	reader := bufio.NewReader(os.Stdin)
	var name string
	name, err := reader.ReadString('\n')
	if err != nil || name == "\n" {
		fmt.Printf("Error: %s\n", ErrInvalidInput)
		return
	}
	name = strings.TrimSpace(name)
	fmt.Print("Umur: ")
	ageText, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error: %s\n", ErrInvalidInput)
		return
	}
	age, err := strconv.Atoi(strings.TrimSpace(ageText))
	if err != nil {
		fmt.Printf("Error: %s\n", ErrInvalidInput)
		return
	}
	if age < 18 {
		fmt.Printf("Error: %s\n", ErrUnderage)
		return
	}
	fmt.Printf("Selamat datang, %s!\n", name)
}
