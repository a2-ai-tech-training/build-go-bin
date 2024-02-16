package get

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	//"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func get(args []string) {

	args := os.Args[1:]
	folder := os.Args[1]
	arguments := os.Args[2:]
	argString := strings.Join(args, " ")

	hasher := md5.New()

	for _, v := range arguments {
		hasher.Write([]byte(v))
	}

	hash := hex.EncodeToString(hasher.Sum(nil))

	name := fmt.Sprintf("%s/outputs/%s.txt", folder, hash)

	file, err := os.Create(name)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	cmd := exec.Command(args[0], args[1:]...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	fmt.Printf("%s", b)

	wut2write := fmt.Sprintf("%s\n%s", argString, b)

	length, err := file.WriteString(wut2write)
	if err != nil {
		panic(err)
	}

	if length == 0 {
		fmt.Printf("No bytes written")
	}

	defer file.Close()

}
