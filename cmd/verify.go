// Package cmd /*
package cmd

import (
	"dual-signature/lib"
	"github.com/spf13/cobra"
)

var PimdHex string
var SignatureHex string
var KucFile string

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify OI",
	Run: func(cmd *cobra.Command, args []string) {
		lib.Verify(PimdHex, OiFile, KucFile, SignatureHex)
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)
	verifyCmd.Flags().StringVar(&PimdHex, "pimd_hex", "", "PIMD hex")
	_ = verifyCmd.MarkFlagRequired("pimd_hex")
	verifyCmd.Flags().StringVar(&OiFile, "oi_file", "", "Path to the OI File")
	_ = verifyCmd.MarkFlagRequired("oi_file")
	verifyCmd.Flags().StringVar(&KucFile, "kuc_file", "", "Customer's public key")
	_ = verifyCmd.MarkFlagRequired("kuc_file")
	verifyCmd.Flags().StringVar(&SignatureHex, "signature_hex", "", "Signature hex")
	_ = verifyCmd.MarkFlagRequired("signature_hex")
}
