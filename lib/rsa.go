package lib

import (
	"crypto/rsa"
	"math/big"
)

func RsaPublicDecrypt(pubKey *rsa.PublicKey, data []byte) []byte {
	c := new(big.Int)
	m := new(big.Int)
	m.SetBytes(data)
	e := big.NewInt(int64(pubKey.E))
	c.Exp(m, e, pubKey.N)
	out := c.Bytes()
	skip := 0
	for i := 2; i < len(out); i++ {
		if i+1 >= len(out) {
			break
		}
		if out[i] == 0xff && out[i+1] == 0 {
			skip = i + 2
			break
		}
	}
	return out[skip:]
}

func copyWithLeftPad(dest, src []byte) {
	numPaddingBytes := len(dest) - len(src)
	for i := 0; i < numPaddingBytes; i++ {
		dest[i] = 0
	}
	copy(dest[numPaddingBytes:], src)
}

func RSAPrivateEncrypt(privKey *rsa.PrivateKey, msg []byte) ([]byte, error) {
	k := (privKey.N.BitLen() + 7) / 8
	if len(msg) > k-11 {
		return nil, rsa.ErrMessageTooLong
	}

	em := make([]byte, k)
	em[1] = 1
	for i := 2; i < len(em)-len(msg)-1; i++ {
		em[i] = 0xff
	}
	mm := em[len(em)-len(msg):]
	em[len(em)-len(msg)-1] = 0
	copy(mm, msg)

	m := new(big.Int).SetBytes(em)
	c := new(big.Int).Exp(m, privKey.D, privKey.N)

	copyWithLeftPad(em, c.Bytes())
	return em, nil
}
