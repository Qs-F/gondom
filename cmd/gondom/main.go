package main

import (
	"bufio"
	"fmt"
	"github.com/Qs-F/gondom"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	t := time.Now().Unix()
	b := bufio.NewWriter(os.Stdout)
	if len(os.Args) >= 2 {
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Print("Arg is not safe")
			os.Exit(1)
		}
		if len(os.Args) == 3 {
			if os.Args[2] == "-u" {
				g := gondom.MakeURL(n, int64(t))
				// println(g)
				fmt.Fprintf(b, g)
				_ = b.Flush()
			}
		} else {
			g := gondom.Make(n, int64(t))
			// print(g)
			fmt.Fprintf(b, g)
			_ = b.Flush()
		}
	} else {
		log.Print("args is not enough")
		os.Exit(1)
	}
}
