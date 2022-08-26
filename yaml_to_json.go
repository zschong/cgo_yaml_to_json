package main

import "C"
import (
	"encoding/json"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

//export yaml_to_json
func yaml_to_json(path string) *C.char {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var x interface{}
	err = yaml.Unmarshal(data, &x)
	if err != nil {
		panic(err)
	}
	json_data, err := json.MarshalIndent(x, "", "    ")
	if err != nil {
		panic(err)
	}
	return C.CString(string(json_data))
}

func main() {}
