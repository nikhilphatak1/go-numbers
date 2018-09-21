package main

import (
	"fmt"
	"bufio"
	"os"
	"net/http"
	"log"
	"io/ioutil"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Give me a number: ")
	num, _ := reader.ReadString('\n')
	num = strings.TrimSuffix(num, "\n")
	url := fmt.Sprintf("http://numbersapi.com/%s/math",num)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(contents))
	}

}
