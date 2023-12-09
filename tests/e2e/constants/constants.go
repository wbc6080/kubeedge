package constants

import "time"

const (
	Interval = 5 * time.Second
	Timeout  = 10 * time.Minute

	E2ELabelKey   = "kubeedge"
	E2ELabelValue = "e2e-test"

	NodeName                   = "edge-node"
	MakeModbusDevice           = "cd $GOPATH/src/github.com/kubeedge/mappers-go;make device modbus package"
	CheckModbusDeviceImage     = "docker images | grep modbus-simulator"
	RunModbusTCPDevice         = "docker run -d --network host modbus-simulator:v1.0-linux-amd64 tcp"
	RunModbusTCPErrorDevice    = "docker run -d --network host modbus-simulator:v1.0-linux-amd64 tcperror"
	RunModbusRTUDevice         = "docker run -d -v /dev:/dev modbus-simulator:v1.0-linux-amd64 rtu"
	RunModbusRTUErrorDevice    = "docker run -d -v /dev:/dev modbus-simulator:v1.0-linux-amd64 rtuerror"
	RunSocat                   = "docker run -d -v /dev:/dev alpine/socat:1.7.3.4-r0 -d -d pty,raw,echo=0,link=/dev/ttyS001 pty,raw,echo=0,link=/dev/ttyS002"
	GetSocat                   = "docker ps | grep alpine/socat"
	GetModbusDeviceContainerID = "docker ps | grep modbus-simulator"
	DeleteModbusDevice         = "docker rmi modbus-simulator:v1.0-linux-amd64"

	MakeModbusMapperProject  = "cd /root/push/wbc/kubeedge/staging/src/github.com/kubeedge/mapper-framework;make generate modbus;"
	BuildModbusMapperProject = "cp -r /root/push/wbc/e2etest/modbus/driver/*  /root/push/wbc/kubeedge/staging/src/github.com/kubeedge/modbus/driver/ ;cp /root/push/wbc/e2etest/modbus/config.yaml  /root/push/wbc/kubeedge/staging/src/github.com/kubeedge/modbus/"
	MakeModbusMapperImage    = "cd $GOPATH/src/github.com/kubeedge/kubeedge/staging/src/github.com/kubeedge/modbus;docker build -t modbus-e2e-mapper:v1.0.0 ."
	CheckModbusMapperImage   = "docker images | grep modbus-e2e-mapper"
	DeleteModbusMapperImage  = "docker rmi modbus-e2e-mapper:v1.0.0"

	RunTCPModbusMapper = "docker run -d --network host -v $GOPATH/src/github.com/kubeedge/mappers-go/tests/e2e/modbus/devicesprofiles/%s.json:/opt/kubeedge/deviceProfile.json modbus-mapper:v1.0-linux-amd64"
	RunRTUModbusMapper = "docker run -d --network host -v /dev:/dev -v $GOPATH/src/github.com/kubeedge/mappers-go/tests/e2e/modbus/devicesprofiles/%s.json:/opt/kubeedge/deviceProfile.json modbus-mapper:v1.0-linux-amd64"
	//GetModbusMapperContainerID = "docker ps | grep modbus-mapper"

	MakeModbusMapperContainer = "docker run -d -v /etc/kubeedge:/etc/kubeedge modbus-e2e-mapper:v1.0.0 ./main --v 4 --config-file config.yaml"
	GetModbusMapperContainer  = "docker ps | grep modbus-e2e-mapper:v1.0.0"
	DeleteMapperContainer     = "docker stop `docker ps |grep mapper |awk '{print $1}'`; docker rm `docker ps -a|grep mapper |awk '{print $1}'`"
)

var (
	// KubeEdgeE2ELabel labels resources created during e2e testing
	KubeEdgeE2ELabel = map[string]string{
		E2ELabelKey: E2ELabelValue,
	}
)
