package types

type CreateAppRequest struct {
	Metadata struct {
		Name      string `json:"name"`
		NameSpace string `json:"nameSpace"`
	} `json:"metadata"`
	Components []struct {
		Name       string `json:"name"`
		Type       string `json:"type"`
		Properties struct {
			Image           string   `json:"image"`
			Cmd             []string `json:"cmd"`
			Port            uint     `json:"port"`
			Cpu             string   `json:"cpu"`
			Mem             string   `json:"mem"`
			ImagePullPolicy string   `json:"imagePullPolicy"`
			Volumes         []struct {
				Name      string `json:"name"`
				MountPath string `json:"mountPath"`
				Type      string `json:"type"`
			} `json:"volumes"`
			Env []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"env"`
			LivenessProbe  struct{} `json:"livenessProbe"`
			ReadinessProbe struct{} `json:"readinessProbe"`
		}
		Traits struct {
			Ingress []struct {
				Properties struct {
					Domain string          `json:"domain"`
					Http   map[string]uint `json:"http"`
				} `json:"properties"`
			} `json:"ingress"`

			Rollout struct {
				Properties struct {
					TargetSize     uint   `json:"targetSize"`
					RolloutBatches []uint `json:"rolloutBatches"`
				} `json:"properties"`
			} `json:"rollout"`

			CpuScaler struct {
				Properties struct {
					Min        uint `json:"min"`
					Max        uint `json:"max"`
					CpuPercent uint `json:"cpuPercent"`
				} `json:"properties"`
			} `json:"cpuScaler"`

			Labels struct {
				Properties map[string]string `json:"properties"`
			} `json:"labels"`

			Annotations struct {
				Properties map[string]string `json:"properties"`
			} `json:"annotations"`

			Sidecar struct {
				Properties struct {
					Name    string            `json:"name"`
					Image   string            `json:"image"`
					Cmd     []string          `json:"cmd"`
					Volumes map[string]string `json:"volumes"`
				} `json:"properties"`
			}
		}
	}
}
