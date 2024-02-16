package hasher

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

/*
//go:embed outputs/*
var f embed.FS
*/

func Hasher(args []string) string {
	argus := args
	hasher := md5.New()

	for _, v := range argus {
		hasher.Write([]byte(v))
	}

	hash := hex.EncodeToString(hasher.Sum(nil))

	//hash_path := fmt.Sprintf("outputs/%s.txt", hash)
	fmt.Printf("Hash for %s:\n%s\n", args, hash)
	return hash
}
