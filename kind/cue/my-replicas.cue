"my-replicas": {
        annotations: {}
        attributes: {
                appliesToWorkloads: []
                conflictsWith: []
                podDisruptive:   false
        }
        description: "Manually scale K8s pod for your workload which follows the pod spec in path 'spec.template'."
        labels: {}
        type: "trait"
}

template: {
        // +patchStrategy=retainkeys
        patch: spec: replicas: parameter.replicas
        parameter: {
                replicas: *1 | int
        }
}