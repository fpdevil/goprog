package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("usage: %v <filename>\n", filepath.Base(args[0]))
		return
	}

	filename := args[1]

	var myRecord Record
	err := loadFronJSON(filename, &myRecord)
	if err == nil {
		fmt.Println("JSON:", myRecord)
	} else {
		fmt.Println(err)
	}

	myRecord.Name = "Donald"
	xmlData, _ := xml.MarshalIndent(myRecord, "", "\t")
	xmlData = []byte(xml.Header + string(xmlData))
	fmt.Printf("\nxmlData:\n%s", string(xmlData))

	data := &Record{}
	err = xml.Unmarshal(xmlData, data)
	if err != nil {
		fmt.Println("Unmarshalling from XML", err)
		return
	}

	result, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error marshalling to json", err)
		return
	}

	_ = json.Unmarshal([]byte(result), &myRecord)
	fmt.Println("\njson:", myRecord)
}

func loadFronJSON(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(in)
	err = decoder.Decode(key)
	if err != nil {
		return err
	}
	in.Close()
	return nil
}
