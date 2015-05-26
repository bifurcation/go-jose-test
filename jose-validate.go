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

	jws, err := jose.ParseSigned(string(bytes))

	if err != nil {
		fmt.Println("validation: FAIL")
		fmt.Println(err)
		return
	} else {
		fmt.Println("validation: OK")
	}

	compact, err := jws.CompactSerialize()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(compact)
}
