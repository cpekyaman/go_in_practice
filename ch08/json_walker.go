package main

import (
	"encoding/json"
	"fmt"
)

var ks = []byte(`{
	"firstName" : "Alex", 
	"lastName": "Murphy",
	"age": 40,
	"education": [
		{
			"institution":"some college",
			"degree":"Bachelor Of Science"
		}
	],
	"children": [
		"jack",
		"william"
	]
}
`)

func main() {
	var f interface{}
	err := json.Unmarshal(ks, &f)
	if err != nil {
		fmt.Print(err)
		return
	}
	printJson(0, f)
}

func printJson(level int, v interface{}) {
	switch vv := v.(type) {
	case string:
		fmt.Println("is string", vv)
	case float64:
		fmt.Println("is float", vv)
	case int32:
	case int16:
		fmt.Println("is integer", vv)
	case int64:
		fmt.Println("is long", vv)
	case []interface{}:
		fmt.Println("is array")
		for i, e := range vv {
			printIndent(level)
			fmt.Print(i, " ")
			printJson(level+1, e)
		}
	case map[string]interface{}:
		fmt.Println("is an object")
		for i, e := range vv {
			printIndent(level)
			fmt.Print(i, " ")
			printJson(level+1, e)
		}
	default:
		fmt.Println("Unknown type", vv)
	}
}

func printIndent(level int) {
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}
}
