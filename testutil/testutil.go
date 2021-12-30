package testutil

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
)

func GenerateKeyring(t *testing.T, accts ...string) keyring.Keyring {
	t.Helper()
	kr := keyring.NewInMemory()

	for _, acc := range accts {
		_, mnm, err := kr.NewMnemonic(acc, keyring.English, "", "", hd.Secp256k1)
		if err != nil {
			panic(err)
		}

		_, err = kr.NewAccount(acc, mnm, "1234", "", hd.Secp256k1)
		if err != nil {
			panic(err)
		}
	}

	return kr
}
