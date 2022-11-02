// Package shhchecker provides realisation of Cheker interface
// for checking ssh-files: "known_hosts", "authorized_keys"

package shhchecker

import (
	"fmt"
	"os"
	"path"

	"github.com/s-vvardenfell/psyllium/internal/checker"
	"github.com/sirupsen/logrus"
)

const (
	hosts           = "known_hosts"
	keys            = "authorized_keys"
	sshDirNotExists = "ssh dir does not exists"
)

type SSHChecker struct {
	opts checker.CheckOptions
}

func New(opts checker.CheckOptions) checker.Checker {
	return &SSHChecker{opts: opts}
}

// мб должен возвр []CheckResult по кол-ву файлов
// марш CheckResult в json
// отправка CheckResult в канал
func (s *SSHChecker) Check() checker.CheckResult {
	// res := checker.CheckResult{}

	home, err := os.UserHomeDir()
	if err != nil {
		logrus.Fatal(err)
	}

	sshDir := path.Join(home, ".ssh")

	ex := exists(sshDir)
	if !ex {
		return checker.CheckResult{
			CheckNotRequired: true,
			Msg:              sshDirNotExists,
			Err:              err,
		}
	}

	if ex {
		for _, file := range s.opts.FileNames {
			switch {
			case file == hosts:
				if exists(path.Join(sshDir, hosts)) {
					s.checkKnownHosts()
				}

			case file == keys:
				if exists(path.Join(sshDir, keys)) {
					s.checkAuthKeys()
				}
			}
		}
	}

	return checker.CheckResult{}
}

func (s *SSHChecker) checkKnownHosts() {
	fmt.Println("checkKnownHosts works!")
}

func (s *SSHChecker) checkAuthKeys() {
	fmt.Println("checkAuthKeys works!")
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}
