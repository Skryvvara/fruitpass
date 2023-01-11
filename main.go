package main

import (
	"bufio"
	"embed"
	"encoding/base64"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path"
	"time"
)

//go:embed all:data
var data embed.FS

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(filename string) ([]string, error) {
	file, err := data.Open(path.Join("data", filename))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func getRandomEntry(list []string) string {
	return list[rand.Intn(len(list)-1)]
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

	fruitList, err := readLines("fruit.txt")
	if err != nil {
		panic(err)
	}

	colorList, err := readLines("colors.txt")
	if err != nil {
		panic(err)
	}

	adjectiveList, err := readLines("adjectives.txt")
	if err != nil {
		panic(err)
	}

	password := fmt.Sprintf("%d%s-%s-%s!", rand.Intn(9), getRandomEntry(adjectiveList), getRandomEntry(colorList), getRandomEntry(fruitList))
	fmt.Println(password)

	if *encode {
		encodedPassword := base64.StdEncoding.EncodeToString([]byte(password))
		fmt.Println(encodedPassword)
	}
}
