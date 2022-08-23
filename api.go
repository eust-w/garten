package garten

import "garten/vm_pool"

type Vm struct {
	ip       string
	password string
	user     string
}

func (v Vm) Upload() error {
	return nil
}

func (v Vm) ZguestInit() error {
	return nil
}

func (v Vm) UploadObject(path string) error {
	return nil
}

func (v Vm) RunCmd(cmd string) (string, error) {
	return "", nil
}

func NewVm() (Vm, error) {
	vmInstance := Vm{}
	ipMnIp := ""
	ipMnPort := ""
	ipMnInstance, err := vm_pool.NewIpMn(ipMnIp, ipMnPort)
	if err != nil {
		return Vm{}, err
	}
	ipInstance, err := vm_pool.NewIp(ipMnInstance)
	if err != nil {
		return Vm{}, err
	}
	vmInstance.ip = ipInstance.GetIp()
	return vmInstance, nil
}
