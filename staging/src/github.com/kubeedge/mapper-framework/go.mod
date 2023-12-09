module github.com/kubeedge/Template

go 1.20

require (
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/mux v1.8.0
	github.com/kubeedge/kubeedge v1.15.0
	github.com/spf13/pflag v1.0.6-0.20210604193023-d5e0c0615ace
	golang.org/x/net v0.8.0 // indirect
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.28.1
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/klog/v2 v2.80.1
)

require (
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20220502173005-c8bf987b8c21 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)

replace (
	github.com/kubeedge/beehive v0.0.0 => github.com/kubeedge/beehive v1.13.0
	github.com/kubeedge/viaduct v0.0.0 => github.com/kubeedge/viaduct v1.13.0
	k8s.io/component-helpers => github.com/kubeedge/kubernetes/staging/src/k8s.io/component-helpers v1.26.10-kubeedge1
	k8s.io/dynamic-resource-allocation => github.com/kubeedge/kubernetes/staging/src/k8s.io/dynamic-resource-allocation v1.26.10-kubeedge1
	k8s.io/legacy-cloud-providers => github.com/kubeedge/kubernetes/staging/src/k8s.io/legacy-cloud-providers v1.26.10-kubeedge1
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.26.10
)
