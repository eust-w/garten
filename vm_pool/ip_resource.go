package vm_pool

// 1. 先申请一个IpMn，一般都是确定的
// 2. 再申请一个可用的ip地址，使用完成后释放，释放前会删除相对应的资源

type IpMn struct {
	ip   string
	port string
}

func NewIpMn(ip, port string) (IpMn, error) {
	return IpMn{ip: ip, port: port}, nil
}

func (I IpMn) GetAvailableIpList() []string {
	return nil
}

func (I IpMn) FreedIp(ip string) error {
	return nil
}

func (I IpMn) MallocIp() (string, error) {
	return "", nil
}

type Ip struct {
	ip   string
	ipMn IpMn
}

func NewIp(mn IpMn) (Ip, error) {
	ipAddress, err := mn.MallocIp()
	if err != nil {
		return Ip{}, err
	}
	ip := Ip{ip: ipAddress, ipMn: mn}
	return ip, nil
}

func (I Ip) Free() error {
	return I.ipMn.FreedIp(I.ip)
}
func (I Ip) GetIp() string {
	return I.ip
}
