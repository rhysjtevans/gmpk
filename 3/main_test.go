package main

import "testing"

func TestGetJsonObject(t *testing.T) {
    json := `{"a":{"b":{"c":"d"}}}`
    result := GetJsonObject(json,"a.b.c") 
    
    if result != "d" {
        t.Error("Expected a.b.c==d but got", result)
    }
}
