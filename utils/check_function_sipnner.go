package utils

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
)

type functionParams func() (interface{}, string, string) // data, pesan sukses, pesan gagal

func SpinnerCheck(desc string, checkProcedure functionParams) (interface{}, string) {
	var (
		e           string
		pesanSukses string
		data        interface{}
	)
	newSetCheck := []string{"✓"}
	newSetFail := []string{"⨯"}
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Prefix = desc + " "
	s.Start()
	time.Sleep(1 * time.Second)
	data, pesanSukses, e = checkProcedure()
	if e != "" {
		s.UpdateCharSet(newSetFail)
		s.Color("red")
		s.Prefix = e + " "
		time.Sleep(1 * time.Second)
		s.Restart()
		fmt.Print("\n")
		os.Exit(0)
	} else {
		s.UpdateCharSet(newSetCheck)
		s.Color("green")
		s.Prefix = "Sukses, " + strings.ToLower(pesanSukses) + " "
		s.Restart()
		time.Sleep(1 * time.Second)
		s.Stop()
	}
	fmt.Print("\n")
	return data, e
}
