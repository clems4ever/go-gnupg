package gognupg

import (
	"bytes"
	"context"
	"log"
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

var GnuPGHomeEnvVarName = "GNUPGHOME"

type GnuPG struct {
	homedir string

	// pipes stderr of gnupg to the stderr of your application.
	pipeStdErr bool
}

type GnuPGOptions func(*GnuPG)

func WithHomeDir(homeDir string) GnuPGOptions {
	return func(gp *GnuPG) {
		gp.homedir = homeDir
	}
}

func WithPipeStdErr() GnuPGOptions {
	return func(gp *GnuPG) {
		gp.pipeStdErr = true
	}
}

func NewGnuPG(options ...GnuPGOptions) *GnuPG {
	gp := &GnuPG{}

	for _, opt := range options {
		opt(gp)
	}

	if gp.homedir == "" {
		gp.homedir = os.Getenv(GnuPGHomeEnvVarName)
	}
	return gp
}

func (gpg *GnuPG) runGnuPG(ctx context.Context, input []byte, args ...string) ([]byte, error) {
	var inputBuf = bytes.NewBuffer(input)
	var outputBuf bytes.Buffer

	cmd := exec.CommandContext(ctx, "gpg", args...)
	cmd.Stdin = inputBuf
	cmd.Stdout = &outputBuf
	if gpg.pipeStdErr {
		cmd.Stderr = os.Stderr
	}

	if gpg.homedir != "" {
		cmd.Env = append(cmd.Env, "GNUPGHOME="+gpg.homedir)
	}

	err := cmd.Run()
	if err != nil {
		log.Println(outputBuf.String())
		return nil, errors.Wrap(err, "failed to run command")
	}
	return outputBuf.Bytes(), nil
}
