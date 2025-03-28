package utils

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	clientset "k8s.io/client-go/kubernetes"
)

func MakeMapperImages(makeMapperProject, getModbusCode, buildModbusMapperProject, makeMapperImage string) error {
	// build mapper project
	cmd := exec.Command("sh", "-c", makeMapperProject)
	if err := PrintCmdOutput(cmd); err != nil {
		return err
	}

	cmd = exec.Command("sh", "-c", getModbusCode)
	if err := PrintCmdOutput(cmd); err != nil {
		return err
	}

	cmd = exec.Command("sh", "-c", buildModbusMapperProject)
	if err := PrintCmdOutput(cmd); err != nil {
		return err
	}

	// check images exist
	Infof("begin build mapper images")
	cmd = exec.Command("sh", "-c", makeMapperImage)
	if err := PrintCmdOutput(cmd); err != nil {
		return err
	}
	return nil
}

func CheckMapperImage(checkMapperImage string) error {
	cmd := exec.Command("sh", "-c", checkMapperImage)
	if err := PrintCmdOutput(cmd); err != nil {
		return err
	}
	return nil
}

// run mapper
func RunMapper(runMapper, checkMapperRun string) error {
	Infof("run mapper image on docker")
	time.Sleep(1 * time.Second)
	cmd := exec.Command("sh", "-c", runMapper)
	if err := PrintCmdOutput(cmd); err != nil {
		return err
	}

	time.Sleep(5 * time.Second)

	Infof("check whether mapper container run successfully")
	cmd = exec.Command("sh", "-c", checkMapperRun)
	if err := PrintCmdOutput(cmd); err != nil {
		return err
	}
	return nil
}

// stop mapper container
func RemoveMapperContainer(deleteMapperContainer string) error {
	Infof("stop mapper container running")
	cmd := exec.Command("sh", "-c", deleteMapperContainer)
	if err := PrintCmdOutput(cmd); err != nil {
		return err
	}
	return nil
}

// delete mapper image
func RemoveMapperImage(deleteMapperImage string) error {
	cmd := exec.Command("sh", "-c", deleteMapperImage)
	if err := PrintCmdOutput(cmd); err != nil {
		return err
	}
	return nil
}

// delete mapper project
func RemoveMapperProject(deleteMapperProject, deleteModbusCode string) error {
	cmd := exec.Command("sh", "-c", deleteMapperProject)
	if err := PrintCmdOutput(cmd); err != nil {
		return err
	}
	cmd = exec.Command("sh", "-c", deleteModbusCode)
	if err := PrintCmdOutput(cmd); err != nil {
		return err
	}
	return nil
}

func CreateMapperDeployment(c clientset.Interface, replica int32, deplName string) *v1.PodList {
	ginkgo.By(fmt.Sprintf("create deployment %s", deplName))
	d := NewMapperDeployment(replica)
	_, err := CreateDeployment(c, d)
	gomega.Expect(err).To(gomega.BeNil())

	ginkgo.By(fmt.Sprintf("get deployment %s", deplName))
	_, err = GetDeployment(c, v1.NamespaceDefault, deplName)
	gomega.Expect(err).To(gomega.BeNil())

	time.Sleep(time.Second * 10)

	ginkgo.By(fmt.Sprintf("get pod for deployment %s", deplName))
	labelSelector := labels.SelectorFromSet(map[string]string{"app": deplName})
	podList, err := GetPods(c, v1.NamespaceDefault, labelSelector, nil)
	gomega.Expect(err).To(gomega.BeNil())
	gomega.Expect(podList).NotTo(gomega.BeNil())

	ginkgo.By(fmt.Sprintf("wait for pod of deployment %s running", deplName))
	WaitForPodsRunning(c, podList, 240*time.Second)

	return podList
}
