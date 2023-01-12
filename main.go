package main

import (
	_ "embed"
	"encoding/base64"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

//go:embed data/fruit.txt
var fruits string

//go:embed data/colors.txt
var colors string

//go:embed data/adjectives.txt
var adjectives string

var adjectiveList = strings.Split(adjectives, "\n")
var colorList = strings.Split(colors, "\n")
var fruitList = strings.Split(fruits, "\n")

func getRandomEntry(list []string) string {
	return list[rand.Intn(len(list)-1)]
}

func GeneratePassword() string {
	return fmt.Sprintf("%d%s-%s-%s!", rand.Intn(9), getRandomEntry(adjectiveList), getRandomEntry(colorList), getRandomEntry(fruitList))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var encode *bool
	encode = flag.Bool("base64", false, "Print the the encoded version of the password alongside the generated password")
	encode = flag.Bool("b", false, "[shorthand] Print the the encoded version of the password alongside the generated password")

	var help *bool
	help = flag.Bool("help", false, "Show command usage")
	help = flag.Bool("h", false, "[shorthand] Show command usage")

	flag.Parse()

	if *help {
		fmt.Fprintf(os.Stdout, "usage %s \n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}

	password := GeneratePassword()
	fmt.Println(password)

	if *encode {
		encodedPassword := base64.StdEncoding.EncodeToString([]byte(password))
		fmt.Println(encodedPassword)
	}
}
