package main

import (
    "github.com/tidwall/gjson"
    "flag"
    "fmt"
    "bufio"
    "os"
)
var jsonKey string

func GetJsonObject(str string, jsonKey string) string {
    if jsonKey != "" {
        return gjson.Get(str, jsonKey).String()
    }else{
        return str
    }
}


func main() {
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    flag.StringVar(&jsonKey, "key", "","A decimal notated index key of the json object you want to retrieve")
    flag.Parse()
    fmt.Println(GetJsonObject(text, jsonKey))
}
