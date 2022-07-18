package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	var i int = 1
	for {
		var URL string = "https://xkcd.com/" + strconv.Itoa(i) + "/info.0.json"
		response, err := http.Get(URL)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		if response.StatusCode != 200 {
			fmt.Printf("response.StatusCode: %v\n", response.StatusCode)
			fmt.Println(i)
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(responseData))
		fileName := strconv.Itoa(i) + ".json"
		f, err := os.Create(fileName)

		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		_, err2 := f.WriteString(string(responseData))
		if err2 != nil {
			log.Fatal(err2)
		}
		fmt.Println("Saving", i, "...")
		i++
	}

}
