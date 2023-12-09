package mapper

import (
	"github.com/onsi/gomega"
	"regexp"
	"strconv"
	"time"

	"github.com/kubeedge/kubeedge/tests/e2e/constants"
	"github.com/kubeedge/kubeedge/tests/e2e/utils"

	"github.com/onsi/ginkgo/v2"
	//"k8s.io/kubernetes/test/e2e/framework"
)

var _ = GroupDescribe("Modbus TCP Mapper test in E2E scenario", func() {

	ginkgo.Context("Testing TCP Modbus Mapper", func() {

		ginkgo.BeforeEach(func() {
			err := utils.StartDevices(constants.RunModbusTCPDevice, constants.GetModbusDeviceContainerID)
			if err != nil {
				log.Error("Fail to run tcp device")
				gomega.Expect(err).Should(BeNil())
			}
		})

		ginkgo.AfterEach(func() {
			err := utils.StopAndDeleteMapper(constants.GetModbusMapperContainerID)
			if err != nil {
				log.Error("Fail to stop and delete modbus mapper container")
				Expect(err).Should(BeNil())
			}
			err = utils.StopAndDeleteDevice(constants.GetModbusDeviceContainerID)
			if err != nil {
				log.Error("Fail to stop ans delete modbus device container")
				gomega.Expect(err).Should(BeNil())
			}
		})

		ginkgo.It("test about writing data into holding register - int16", func() {
			// run mapper container
			err := utils.RunMapper(constants.RunTCPModbusMapper, writeInt16, constants.GetModbusMapperContainerID)
			if err != nil {
				log.Error("Fail to run mapper ")
			}
			time.Sleep(5 * time.Second)

			containerLog, _ := utils.ReadDockerLog(constants.GetModbusMapperContainerID, readTwoLines)
			reg1 := regexp.MustCompile(`Get the alarming-temperature value as ([0-9.-]+)`)
			temperatureValue := reg1.FindStringSubmatch(containerLog)

			Expect(temperatureValue[1]).Should(Equal(strconv.Itoa(30)))
		})

		ginkgo.It("test about reading holding value", func() {

			err := utils.RunMapper(constants.RunTCPModbusMapper, readHolding, constants.GetModbusMapperContainerID)
			if err != nil {
				log.Error("Fail to run mapper ")
			}
			time.Sleep(5 * time.Second)

			containerLog, _ := utils.ReadDockerLog(constants.GetModbusMapperContainerID, readTwoLines)
			reg1 := regexp.MustCompile(`Get the temperature value as ([0-9.-]+)`)
			temperatureValue := reg1.FindStringSubmatch(containerLog)

			deviceLog, _ := utils.ReadDockerLog(constants.GetModbusDeviceContainerID, readEightLines)
			reg2 := regexp.MustCompile(`temperature value is ([0-9.-]+)`)
			temperatureValue2 := reg2.FindStringSubmatch(deviceLog)

			gomega.Expect(temperatureValue[1]).Should(Equal(temperatureValue2[1]))
		})

	})

	ginkgo.Context("Testing Tcp Modbus Mapper In Negative Cases", func() {
		ginkgo.BeforeEach(func() {
			err := utils.StartDevices(constants.RunModbusTCPErrorDevice, constants.GetModbusDeviceContainerID)
			if err != nil {
				log.Error("Fail to run tcp device")
				gomega.Expect(err).Should(BeNil())
			}

		})

		ginkgo.AfterEach(func() {
			err := utils.StopAndDeleteMapper(constants.GetModbusMapperContainerID)
			if err != nil {
				log.Error("Fail to stop and delete modbus mapper container")
				Expect(err).Should(BeNil())
			}
			err = utils.StopAndDeleteDevice(constants.GetModbusDeviceContainerID)
			if err != nil {
				log.Error("Fail to stop ans delete modbus device container")
				gomega.Expect(err).Should(BeNil())
			}

		})

		ginkgo.It("test negative case about modbus tcp error code 4 ", func() {
			err := utils.RunMapper(constants.RunTCPModbusMapper, errorCode4, constants.GetModbusMapperContainerID)
			if err != nil {
				log.Error("Fail to run mapper ")
			}
			time.Sleep(5 * time.Second)

			containerLog, err := utils.ReadDockerLog(constants.GetModbusMapperContainerID, readTwoLines)
			if err != nil {
				log.Error("Fail to read container log")
			}
			reg1 := regexp.MustCompile(`exception.*4`)
			result := reg1.FindStringSubmatch(containerLog)
			gomega.Expect(len(result)).Should(Equal(1))
		})

		ginkgo.It("test negative case about modbus tcp timeout error", func() {
			err := utils.RunMapper(constants.RunTCPModbusMapper, errorTimeout, constants.GetModbusMapperContainerID)
			if err != nil {
				log.Error("Fail to run mapper ")
			}
			time.Sleep(10 * time.Second)

			containerLog, err := utils.ReadDockerLog(constants.GetModbusMapperContainerID, readSixLines)
			if err != nil {
				log.Error("Fail to read container log")
			}
			reg1 := regexp.MustCompile(`timeout`)
			result := reg1.FindStringSubmatch(containerLog)
			gomega.Expect(len(result)).Should(Equal(1))
		})

	})
})
