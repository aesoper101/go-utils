package golangx

import (
	"fmt"
	"github.com/aesoper101/go-utils/execx"
	"os"
	"os/exec"
	"strings"
)

// GoInstall is a functionx that installs a package from a remote repository.
// returns an error if something goes wrong.
func GoInstall(repository string) error {
	if err := execx.LookPath("go"); err != nil {
		return err
	}

	env := os.Environ()

	if !IsGO111ModuleOn() {
		env = append(env, "GO111MODULE=on")
	}

	env = append(env, fmt.Sprintf("GOPROXY=%s", GoProxy()))

	if !strings.Contains(repository, "@") {
		repository += "@latest"
	}

	cmd := exec.Command("go", "install", repository)
	cmd.Env = env
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
