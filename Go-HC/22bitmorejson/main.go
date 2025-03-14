package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON video")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{"ReactJS ", 299, "Google.com", "abc123", []string{"web-dev", "js"}},
		{"GO", 299, "Google.com", "efg123", []string{"Go-related", "js"}},
		{"RUST", 299, "Google.com", "hij123", []string{"Google", "js"}},
	}

	//package this data as JSON data

	// finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s/n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	       {
                "coursename": "ReactJS",
                "Price": 299,
                "website": "Google.com",
                "tags": ["web-dev","js"]
        	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)

	} else {
		fmt.Println("JSON was not valid")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is  %v and value is %v and Type is: %T\n", k, v, v)

	}
}
