package utils

import (
	"crypto/elliptic"
	"errors"
	"fmt"
	"github.com/saveio/themis/common/log"
	"github.com/saveio/themis/crypto/ec"

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

func SignByETH(acc *account.Account, data []byte) ([]byte, error) {
	keyBytes := acc.GetEthPrivateKey()
	ecdsaKey := ec.ConstructPrivateKey(keyBytes, elliptic.P256())
	priKey := &ec.PrivateKey{
		Algorithm:  ec.ECDSA,
		Raw:        keyBytes,
		PrivateKey: ecdsaKey,
	}
	sig, err := s.Sign(acc.SigScheme, priKey, data, nil)
	if err != nil {
		log.Debugf("SignByETH error:%s", err)
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
