package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var hashType = flag.String("n", "256", "256|384|512")

func main() {

	flag.Parse()
	hash, err := strconv.Atoi(*hashType)
	if err != nil || hash != 256 && hash != 384 && hash != 512 {
		hash = 256
	}

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		switch hash {
		case 256:
			fmt.Printf("%x\n", sha256.Sum256([]byte(text)))
		case 384:
			fmt.Printf("%x\n", sha512.Sum384([]byte(text)))
		case 512:
			fmt.Printf("%x\n", sha512.Sum512([]byte(text)))
		default:
			fmt.Printf("%x\n", sha256.Sum256([]byte(text)))
		}
	}
}
