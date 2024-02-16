package scanner

import (
	"bufio"
	//"embed"
	"fmt"
	"log"
	"os"

	"github.com/a2-ai-tech-training/build-go-bin/internal/hasher"
)

/*
//go:embed outputs/*
var f embed.FS
*/

func scanner() {

	args := os.Args[1:]
	hpath := hasher(args)
	fmt.Printf("%s", hpath)
	fmt.Println()
	file, err := f.Open(hpath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Skips first line that contains command
	scanner.Scan()

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
