package requirement

import (
	"antri-in-cli/vars"
	"runtime"
)

func CheckOS() (interface{}, string, string) {
	var pesan string = ""
	var err string = ""
	var data string = ""
	if runtime.GOOS == "windows" {
		data = vars.OS_NAME_WINDOWS
		pesan = "Terdeteksi menggunakan Windows"
	} else if runtime.GOOS == "linux" {
		data = vars.OS_NAME_LINUX
		pesan = "Terdeteksi menggunakan Linux"
	} else {
		err = "Tidak dapat mendeteksi OS"
	}
	return data, pesan, err
}
