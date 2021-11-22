package main

import (
	"fmt"
	"encoding/json"
)

type Response struct{
	Page int `json:"page"`
	Words []string `json:"words,omitempty"`
}

func main(){

	res := Response{
		Page : 5,
		Words : []string{ "Pankaj", "rahul", "Chutiya"},
	}

	fmt.Printf("%#v\n", res)

	jsonBytes, _ :=json.Marshal(res)
	fmt.Println(string(jsonBytes))

	var res1 Response

	_ = json.Unmarshal(jsonBytes,&res1)

	fmt.Printf("After Marshalling %#v", res1)

}