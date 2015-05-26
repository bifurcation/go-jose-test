package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/square/go-jose"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Phase 1: Parsing
	parsedJws, err := jose.ParseSigned(string(bytes))
	if err != nil {
		fmt.Println("parsing: FAIL")
		fmt.Println(err)
		return
	} else {
		fmt.Println("parsing: OK")
	}

	// Phase 2: Verification
	if len(parsedJws.Signatures) != 1 {
		fmt.Println("Must have exactly one signature, instead of ", len(parsedJws.Signatures))
		return
	}
	key := parsedJws.Signatures[0].Header.JsonWebKey
	_, err = parsedJws.Verify(key)
	if err != nil {
		fmt.Println("verification: FAIL")
		fmt.Println(err)
		return
	} else {
		fmt.Println("verification: OK")
	}
}
