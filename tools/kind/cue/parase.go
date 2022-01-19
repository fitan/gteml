package main

import (
	"cuelang.org/go/cue"
	"cuelang.org/go/encoding/openapi"
	"fmt"
)

func main() {
	const config = `
msg:   "Hello \(place)!"
place: string | *"world" // "world" is the default.
test: string
`

	var r cue.Runtime

	instance, _ := r.Compile("test", config)

	str, _ := instance.Lookup("test").String()

	fmt.Println(str)
	//b, err := gocode.Generate("./", instance, nil)
	//if err != nil {
	//	panic(err.Error())
	//}
	//err = ioutil.WriteFile("cue_gen.go", b, 0644)
	//if err != nil {
	//	panic(err.Error())
	//}

	gen, err := openapi.Gen(instance, nil)
	if err != nil {
		return
	}
	fmt.Println(string(gen))

}
