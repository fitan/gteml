"my-inject-env": {
        annotations: {}
        attributes: {
                appliesToWorkloads: ["*"]
        }
        description: "resource env"
        labels: {}
        type: "trait"
}

template: {
        patch: spec: template: spec: {
        	      // +patchKey=name
                containers: [{
                        name: context.name
                        "env": [
                        {"name":"MY_POD_IP","valueFrom":{"fieldRef":{"apiVersion":"v1","fieldPath":"status.podIP"}}},
                        {"name":"MY_CPU_LIMIT","valueFrom":{"resourceFieldRef":{"containerName":context.name,"resource":"limits.cpu","divisor":"0"}}},
                        {"name":"MY_MEM_LIMIT","valueFrom":{"resourceFieldRef":{"containerName":context.name,"resource":"limits.memory","divisor":"0"}}},
                        {"name":"MY_POD_NAME","valueFrom":{"fieldRef":{"apiVersion":"v1","fieldPath":"metadata.name"}}},
                        {"name":"MY_POD_NAMESPACE","valueFrom":{"fieldRef":{"apiVersion":"v1","fieldPath":"metadata.namespace"}}},
                        {"name":"MY_INSTANCE_IP","valueFrom":{"fieldRef":{"apiVersion":"v1","fieldPath":"status.podIP"}}},
                        {"name":"MY_NODE_NAME","valueFrom":{"fieldRef":{"apiVersion":"v1","fieldPath":"spec.nodeName"}}}]
                }]
        }
        parameter: {}
}