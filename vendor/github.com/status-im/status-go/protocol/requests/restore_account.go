package requests

import (
	"errors"
)

var ErrRestoreAccountInvalidMnemonic = errors.New("restore-account: invalid mnemonic")

type RestoreAccount struct {
	Mnemonic    string `json:"mnemonic"`
	FetchBackup bool   `json:"fetchBackup"`
	CreateAccount
}

func (c *RestoreAccount) Validate() error {
	if len(c.Mnemonic) == 0 {
		return ErrRestoreAccountInvalidMnemonic
	}

	return c.CreateAccount.Validate(&CreateAccountValidation{
		AllowEmptyDisplayName: true,
	})
}
