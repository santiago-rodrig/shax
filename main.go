package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	buff := bufio.NewReader(os.Stdin)
	bs, err := buff.ReadBytes('\n')

	if err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
	}

	c1 := sha256.Sum256(bs)
	fmt.Printf("%x\n", c1)
}
