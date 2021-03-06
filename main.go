package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var algorithms = [3]uint{256, 384, 512}

func algorithmIsValid(algorithm uint) bool {
	for _, a := range algorithms {
		if algorithm == a {
			return true
		}
	}

	return false
}

func main() {
	algorithm := flag.Uint("algo", 256, "-algo=256")
	flag.Parse()

	if !algorithmIsValid(*algorithm) {
		fmt.Printf("Invalid algorithm: %d\nValid values are: 256, 384, and 512\n\n", *algorithm)
		flag.PrintDefaults()
	}

	buff := bufio.NewReader(os.Stdin)
	bs, err := buff.ReadBytes('\n')

	if err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
	}

	switch *algorithm {
	case 256:
		c1 := sha256.Sum256(bs)
		fmt.Printf("%x\n", c1)
	case 384:
		c1 := sha512.Sum384(bs)
		fmt.Printf("%x\n", c1)
	case 512:
		c1 := sha512.Sum512(bs)
		fmt.Printf("%x\n", c1)
	}
}
