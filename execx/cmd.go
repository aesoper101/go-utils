package execx

import (
	"errors"
	"os/exec"
)

func LookPath(commandName ...string) error {
	if len(commandName) == 0 {
		return errors.New("no command specified")
	}
	for _, c := range commandName {
		if _, err := exec.LookPath(c); err != nil {
			return err
		}
	}
	return nil
}
