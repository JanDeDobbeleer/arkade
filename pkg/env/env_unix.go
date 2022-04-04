//go:build !windows

package env

import (
	"log"
	"strings"

	execute "github.com/alexellis/go-execute/pkg/v1"
)

// GetClientArch returns a pair of arch and os
func GetClientArch() (arch string, os string) {
	task := execute.ExecTask{Command: "uname", Args: []string{"-m"}, StreamStdio: false}
	res, err := task.Execute()
	if err != nil {
		log.Println(err)
	}

	archResult := strings.TrimSpace(res.Stdout)

	taskOS := execute.ExecTask{Command: "uname", Args: []string{"-s"}, StreamStdio: false}
	resOS, errOS := taskOS.Execute()
	if errOS != nil {
		log.Println(errOS)
	}

	osResult := strings.TrimSpace(resOS.Stdout)

	return archResult, osResult
}
