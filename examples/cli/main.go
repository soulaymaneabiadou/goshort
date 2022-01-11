package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/soulaymaneabiadou/goshort"
)

func main() {
	fmt.Print("Would you like to create or look up?(c/l) ")
	var a string
	fmt.Scanln(&a)

	switch a {
	case "l":
		fmt.Print("Enter code: ")
		var c string
		fmt.Scanln(&c)
		u, err := goshort.GetUrl(c)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(u.ShortUrl)

	case "c":
		fmt.Print("Enter long URL: ")
		var lu string
		fmt.Scanln(&lu)
		u, _ := goshort.ShortenUrl(lu)
		r, _ := json.Marshal(u)
		fmt.Println(string(r))

	default:
		fmt.Println("Unsupported choice!")
	}

}
