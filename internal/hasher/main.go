package main

import (
	"bufio"
	"crypto/md5"
	"embed"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

//go:embed outputs/*
var f embed.FS

func hasher(args []string) string {
	argus := args
	hasher := md5.New()

	for _, v := range argus {
		hasher.Write([]byte(v))
	}

	hash := hex.EncodeToString(hasher.Sum(nil))

	hash_path := fmt.Sprintf("outputs/%s.txt", hash)
	fmt.Println(hash_path)
	return hash_path
}

func main() {

	args := os.Args[1:]
	//fmt.Println(arguments)
	hpath := hasher(args)
	fmt.Printf("%s", hpath)
	fmt.Println()
	/*hasher := md5.New()

	for _, v := range arguments {
		hasher.Write([]byte(v))
	}

	hash := hex.EncodeToString(hasher.Sum(nil))

	hash_path := fmt.Sprintf("outputs/%s.txt", hash)
	*/
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
