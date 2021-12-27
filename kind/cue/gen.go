package main

import (
	"cuelang.org/go/cue"
	"fmt"
)

func main() {
	const config = `
"my-ports": {
        annotations: {}
        attributes: {
                appliesToWorkloads: []
                conflictsWith: []
                podDisruptive:   false
        }
        description: "add ports"
        labels: {}
        type: "trait"
}

template: {
        patch: spec: template: spec: containers: [{"ports": [for v in parameter.ports {
                containerPort: v.containerPort
                protocol: v.protocol
        }]},...]
        parameter: {
                ports: [...{containerPort: int, protocol: *"TCP" | "UDP"}]
        }
}
`

	var r cue.Runtime

	instance, _ := r.Compile("test", config)

	s, err := instance.Value().Struct()
	if err != nil {
		panic(err)
	}

	var paraDef cue.FieldInfo
	for i := 0; i < s.Len(); i++ {
		paraDef = s.Field(i)
		if paraDef.Name == "parameter" {
			break
		}
	}

	fmt.Println(paraDef.Value)

	//str, _ := instance.Lookup("msg").String()
	//
	//fmt.Println(str)
	//extract, err := jsonschema.Extract(instance, &jsonschema.Config{})
	//if err != nil {
	//	return
	//}
	//ast.Print(nil, extract)

	//b, err := gocode.Generate("./", instance, &gocode.Config{
	//	Prefix:       "",
	//	ValidateName: "",
	//	CompleteName: "",
	//	RuntimeVar:   "",
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = ioutil.WriteFile("cue_gen.go", b, 0064)

}
