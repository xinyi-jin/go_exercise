package balance

import "strconv"

type Instance struct {
	host string
	port int
}

func NewInstance(host string, port int) *Instance {
	return &Instance{
		host: host,
		port: port,
	}
}

func (i *Instance) GetHost() string {
	return i.host
}

func (i *Instance) GetPort() int {
	return i.port
}

func (i *Instance) String() string {
	return i.host + ":" + strconv.Itoa(i.port)
}
