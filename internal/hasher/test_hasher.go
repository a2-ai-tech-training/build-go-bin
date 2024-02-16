package test_hasher

import (
	"fmt"
	"github.com/a2-ai-tech-training/internal/hasher"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type hash_path struct {
	arg      string
	fullPath string
}

func test_hasher() {
	/*
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
	*/
	test1 := []string{""}
	test2 := []string{"--json"}
	test3 := []string{"-vvv"}

	assert.Equal(hasher(test1), "outputs/d41d8cd98f00b204e9800998ecf8427e.txt")
	assert.Equal(hasher(test2), "outputs/31d5e31c2e1e0b3c281f5194f130287e.txt")
	assert.Equal(hasher(test3), "outputs/0cc598961dc9ae41055f83d0950544f3.txt")
	fmt.Println("ayo")
}
