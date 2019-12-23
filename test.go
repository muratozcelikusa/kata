package main

import (
	"bufio"
	"fmt"
	"os"
	//"regexp"
	"strings"
)

func main() {
	//rx:=regexp.MustCompile(`[^a-z]+`)
	args:=os.Args[1:]
	if len(args)!=1 {
		fmt.Println("Enter search word")
		return
	}
	query:=args[0]


	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	words:=make(map[string]bool)

	for in.Scan(){
		word:=strings.ToLower(in.Text())

		if len(word)>2 {
			words[word]=true
		}
		//fmt.Println(in.Text())
	}

	//query:="sun"
	if words[query] {
		fmt.Printf("Input contain %q.\n",query)
		return
	}
	fmt.Printf("not available %q.\n",query)
	return

}