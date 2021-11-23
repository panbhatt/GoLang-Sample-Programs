package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
    "encoding/json"
)


type ToDo struct {
    UserId int `json:"userId"`
    Id int `json:"id`
    Title string `json:"title"`
    Completed bool `json:"completed"`
}

func main() {

    const url = "https://jsonplaceholder.typicode.com/todos/"
    res,err := http.Get(url)

    if err != nil {
        fmt.Println("AN Error occur, while hitting the web service")
        
        os.Exit(1)
    }

    defer res.Body.Close()

    if res.StatusCode == http.StatusOK {
        body, err:= ioutil.ReadAll(res.Body)

        if err != nil {
            fmt.Println(os.Stderr,err)
            os.Exit(-1)
        }
        //fmt.Println(string(body))

        todos :=  []ToDo{}

        er :=json.Unmarshal(body, &todos)

        if er != nil {
            fmt.Println("Parsing to JSON Failed", er)
            os.Exit(-1)
        }

        for i, item := range todos {
            fmt.Printf("%d ID = %d Title = %s  Completed = %t\n ", i+1, item.Id, item.Title, item.Completed)
        }
    }


}