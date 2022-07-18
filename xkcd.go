package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type Episode struct {
	Num       int    `json:"num"`
	Trascript string `json:"transcript"`
}

var matchedList = template.Must(template.New("xkcd List").Parse(`
	<table>
	<tr style='text-align: left'>
	<th>URL</th>
	<th>Trascript</th>
	</tr>
	<tr>
	<td><a href='https://xkcd.com/{{.Num}}/info.0.json'>{{.Num}}</a></td>
	<td>{{.Trascript}}</td>
	</tr>
	</table>
`))

func main() {

	// Get search term from command line
	keyword := os.Args[1]
	files, err := ioutil.ReadDir("offlineIndex/") // relative path
	if err != nil {
		fmt.Println("1")
		log.Fatal(err)
	}
	for _, file := range files { // seach on each file
		// Open our jsonFile
		path := "offlineIndex/" + file.Name()
		jsonFile, err := os.Open(path)
		if err != nil {
			fmt.Println("2")
			fmt.Println(err)
		}
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		// do Unmarshal
		textBytes := []byte(byteValue)
		episode := Episode{}
		err = json.Unmarshal(textBytes, &episode)
		if err != nil {
			fmt.Println(err)
		}
		// search if there in a match
		if strings.Contains(episode.Trascript, keyword) {
			/* 			URL := "https://xkcd.com/" + strconv.Itoa(episode.Num) + "/info.0.json"
			   			fmt.Println(URL)
			   			fmt.Println("--------------------------------------------")
			   			fmt.Println(episode.Trascript)
			   			fmt.Println("--------------------------------------------") */
			if err := matchedList.Execute(os.Stdout, episode); err != nil {
				log.Fatal(err)
			}
		}
	}

	/* //!-template

	   func main() {
	   	result, err := github.SearchIssues(os.Args[1:])
	   	if err != nil {
	   		log.Fatal(err)
	   	}

	   } */
}
