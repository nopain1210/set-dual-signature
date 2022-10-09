package lib

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"github.com/pterm/pterm"
)

func Verify(pimdHex string, oiFile string, kucFile string, signatureHex string) {
	// Step 1: Decrypt the signature with the customer's public key to get the POMD
	pterm.DefaultSection.Println("Step 1: Decrypt the signature with the customer's public key to get the POMD")
	pubKey, err := readPublicKeyFile(kucFile)
	if err != nil {
		pterm.Fatal.Printf("Error while reading public key: %s", err)
	}
	signature, err := hex.DecodeString(signatureHex)
	if err != nil {
		pterm.Fatal.Printf("Error while decoding signatureHex: %s", err)
	}
	pomd := RsaPublicDecrypt(pubKey, signature)
	pterm.Info.Println("POMD hex decrypted from signature: ", hex.EncodeToString(pomd))

	// Step 2: Calculate POMD from PIMD and OI file
	pterm.DefaultSection.Println("Step 2: Calculate POMD from PIMD and OI file")
	pimd, err := hex.DecodeString(pimdHex)
	if err != nil {
		pterm.Fatal.Printf("Error while decoding pimdHex: %s", err)
	}
	oimd, err := calculateFileSha1Sum(oiFile)
	if err != nil {
		pterm.Fatal.Printf("Error while hashing OI file: %s", err)
	}
	po := append(pimd, oimd...)
	sentPomd := sha1.Sum(po)
	pterm.Info.Println("POMD hex calculate from PIMD and OI: ", hex.EncodeToString(sentPomd[:]))

	// Step 3: Compare the POMD from the signature with the POMD calculated from the PIMD and OI file
	pterm.DefaultSection.Println("Step 3: Compare the POMD from the signature with the POMD calculated from the PIMD and OI file")
	if bytes.Compare(pomd, sentPomd[:]) == 0 {
		pterm.Success.Println("Signature is valid")
	} else {
		pterm.Error.Println("Signature is invalid")
	}
}
