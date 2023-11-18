package main

import (
	"encoding/json"
	"fmt"
)


type course struct{
	Name string `json:"coursename"`
	Price int
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags []string `json:"tags,omitempty"`

}
func main(){
//	EncodeJson()
	DecodeJson()


}

func EncodeJson(){
	lcoCourses := []course{
		{"ReactJS", 255, "Youtube", "adda", []string{"react", "full"}},
		{"Go", 255, "Youtube", "adda", []string{"google", "full"}},
		{"ReactJS", 255, "Youtube", "adda", nil},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson(){
	jsonDataFromWeb := []byte(`{
		"coursename": "ReactJs",
		"Price": 233,
		"website": "youtube",
		"tags": ["web-dev", "js"]
	}`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}


	// some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is : %T\n", k, v, v)
	}
}