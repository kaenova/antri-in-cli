package requirement

import (
	"antri-in-cli/vars"
	"net"
)

func CheckPort443() (interface{}, string, string) {
	var pesanSukses string = ""
	var pesanError string = ""
	var data string = ""

	port := "443"
	ln, errors := net.Listen("tcp", ":"+port)
	if errors != nil {
		pesanError = "Port 443 tidak tersedia"
		return nil, "", pesanError
	}
	ln.Close()
	pesanSukses = "Port 443 tersedia"
	data = vars.PORT_443
	return data, pesanSukses, ""
}
