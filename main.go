package main

import (
	"os"
	"fmt"
	"flag"
	"github.com/hellodoge/lpass/lib"
)

const (
	default_length = 15
	default_iter = 0x1000
)

var masterKey, token string
var length, iter int
var uA, ua, u0, u_ bool

func main() {
	var chars string = lib.GetChars(uA, ua, u0, u_)
	var generated []byte = lib.Generate(
		[]byte(masterKey), []byte(token), iter, length)
	var pass string = lib.Translate(generated, chars)
	fmt.Println(pass)
}

func init() {

	flag.BoolVar(&uA, "A", false, "do not use uppercase characters")
	flag.BoolVar(&ua, "a", false, "do not use lowercase characters")
	flag.BoolVar(&u0, "0", false, "do not use digit characters")
	flag.BoolVar(&u_, "_", false, "do not use special characters")
	flag.IntVar(&length, "l", default_length, "password length")
	flag.IntVar(&iter, "i", default_iter, "iteration count")

	flag.Usage = func() {
		fmt.Println("Usage: pass [master key] [token or url] [args]")
		flag.PrintDefaults()
	}

	if len(os.Args) < 3 {
		flag.Usage()
		os.Exit(2)
	}

	masterKey = os.Args[1]
	token = os.Args[2]
	os.Args = os.Args[2:]

	flag.Parse()

	if length <= 0 || iter <= 0 {
		flag.Usage()
		os.Exit(2)
	}
}
