package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("/Users/Micha/Desktop/FillWordTemplate/Template.rtf")
	check(err)
	fmt.Print(string(dat))
	f, err := os.Open("/Users/Micha/Desktop/FillWordTemplate/Template.rtf")
	check(err)

	scanner := bufio.NewScanner(f)
	letter := scanner.Text()
	// Set the Split method to ScanWords.
	//scanner.Split(bufio.ScanWords)

	// Scan all words from the file.
	//for scanner.Scan() {
	//	letter := scanner.Text()
	//	fmt.Println(letter)
	//}
	// Prepare some data to insert into the template.
	type Recipient struct {
		Name, Essen string
		Bringdienst string
	}
	var recipients = []Recipient{
		{"Louisa", "Döner", "Türke"},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
}
