package test_hasher

import (
	"fmt"
	"os"
	"github.com/a2-ai-tech-training/internal/hasher"
)

type hash_path struct {
	arg string
	fullPath string
}

func test_hasher() {
	type test struct {
		input get.args
		options []os.Args
		want string
		wantErr error
		shouldSucceed bool
		context string
	}
	possibleArgs := []values{
		{
			ArgStr: ""
			Str: "outputs/d41d8cd98f00b204e9800998ecf8427e.txt"
		},
		{
			ArgStr: "-vvv"
			Str: "outputs/0cc598961dc9ae41055f83d0950544f3.txt"
		},
		{
			ArgStr: "--json"
			Str: "outputs/31d5e31c2e1e0b3c281f5194f130287e.txt"
		}
	}
	tests := []test{
		
	}
	fmt.Println("ayo")
}
