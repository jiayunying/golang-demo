package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var xxx = map[string]string{}

func readFile(filename string) (map[string]string, error){
	bytes,err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return nil, err
	}
	if err := json.Unmarshal(bytes, &xxx); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return nil, err
	}

	return xxx, nil
}

func main(){
	xxxMap, err := readFile("xxx.json")
	if err != nil {
		fmt.Println("readFile: ", err.Error())
		return
	}

	fmt.Println(xxxMap)
}
