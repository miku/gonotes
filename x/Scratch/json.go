package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
)

type Task struct {
	XMLName  xml.Name `xml:"MyTask"`
	TaskName string   `json:"name" xml:"T"`
	size     int      `json:"size" xml:"S"`
}

func main() {
	t := Task{
		TaskName: "initialize",
	}
	fmt.Printf("%#v\n", t)
	b, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)
	fmt.Println(string(b))
	b, _ = xml.Marshal(t)
	fmt.Println(string(b))
}
