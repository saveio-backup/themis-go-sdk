package utils

import (
	"errors"
	"fmt"

	"github.com/saveio/themis/account"
	"github.com/saveio/themis/crypto/keypair"
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

// Verify check the signature of data using pubKey
func Verify(pubKey keypair.PublicKey, data, signature []byte) error {
	sigObj, err := s.Deserialize(signature)
	if err != nil {
		return errors.New("invalid signature data: " + err.Error())
	}
	if !s.Verify(pubKey, data, sigObj) {
		return errors.New("signature verification failed")
	}
	return nil
}
