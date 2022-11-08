package gognupg

import "context"

type EncryptionOptions struct {
	Armor bool
}

func (gpg *GnuPG) Encrypt(ctx context.Context, data []byte, recipients []string, opts *EncryptionOptions) ([]byte, error) {
	var args []string

	args = append(args, "--encrypt")
	if opts != nil && opts.Armor {
		args = append(args, "--armor")
	}

	for _, recipient := range recipients {
		args = append(args, "-r", recipient)
	}

	output, err := gpg.runGnuPG(ctx, data, args...)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (gpg *GnuPG) Decrypt(ctx context.Context, data []byte) ([]byte, error) {
	var args []string

	args = append(args, "--decrypt")

	output, err := gpg.runGnuPG(ctx, data, args...)
	if err != nil {
		return nil, err
	}

	return output, nil
}
