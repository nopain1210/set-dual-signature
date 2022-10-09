package lib

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/pterm/pterm"
)

func Sign(piFile string, oiFile string, kpcFile string) {
	// Step 1: Calculate the SHA-1 hash of the PI file and OI file
	pterm.DefaultSection.Println("Step 1: Calculate the SHA-1 hash of the PI file and OI file")
	// Hash PI file
	pimd, err := calculateFileSha1Sum(piFile)
	if err != nil {
		pterm.Fatal.Printf("Error while hashing PI file: %s", err)
	}
	pterm.Info.Println("PIMD hex: H(PI) = ", hex.EncodeToString(pimd))

	// Hash OI file
	oimd, err := calculateFileSha1Sum(oiFile)
	if err != nil {
		pterm.Fatal.Printf("Error while hashing OI file: %s", err)
	}
	pterm.Info.Println("OIMD hex: H(OI) = ", hex.EncodeToString(oimd))

	// Step 2: Combine OIMD and PIMD to create the po hash
	pterm.DefaultSection.Println("Step 2: Combine OIMD and PIMD to create the po hash")
	po := append(pimd, oimd...)
	pterm.Info.Println("H(PI) || H(OI) = ", hex.EncodeToString(po))

	// Step 3: Calculate the SHA-1 hash of the po hash
	pterm.DefaultSection.Println("Step 3: Calculate the SHA-1 hash of the po hash")
	pomd := sha1.Sum(po)
	pterm.Info.Println("POMD hex: H(H(PI) || H(OI)) = ", hex.EncodeToString(pomd[:]))

	// Step 4: Sign the POMD with the customer's private key
	pterm.DefaultSection.Println("Step 4: Sign the POMD with the customer's private key (KPc)")

	privKey, err := readPrivateKeyFile(kpcFile)
	if err != nil {
		pterm.Fatal.Printf("Error while reading private key: %s", err)
	}
	signature, err := RSAPrivateEncrypt(privKey, pomd[:])
	if err != nil {
		pterm.Fatal.Printf("Error while encrypting: %s", err)
	}
	pterm.Info.Println("DS hex: E(KPc, H(H(PI) || H(OI))) = ", hex.EncodeToString(signature))
}
