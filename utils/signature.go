package utils

import (
	"fmt"

	"github.com/saveio/themis/account"
	s "github.com/saveio/themis/crypto/signature"
)

func Sign(acc *account.Account, data []byte) ([]byte, error) {
	sig, err := s.Sign(acc.SigScheme, acc.PrivateKey, data, nil)
	if err != nil {
		return nil, err
	}
	sigData, err := s.Serialize(sig)
	if err != nil {
		return nil, fmt.Errorf("signature.Serialize error:%s", err)
	}
	return sigData, nil
}
