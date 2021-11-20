package requirement

import (
	"antri-in-cli/vars"
	"net"
)

func CheckPort80() (interface{}, string, string) {
	var pesanSukses string = ""
	var pesanError string = ""
	var data string = ""

	port := "80"
	ln, errors := net.Listen("tcp", ":"+port)
	if errors != nil {
		pesanError = "Port 80 tidak tersedia"
		return nil, "", pesanError
	}
	ln.Close()
	pesanSukses = "Port 80 tersedia"
	data = vars.PORT_80
	return data, pesanSukses, ""
}
