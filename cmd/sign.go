// Package cmd /*
package cmd

import (
	"dual-signature/lib"
	"github.com/spf13/cobra"
)

var PiFile string
var OiFile string
var KpcFile string

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign files",
	Run: func(cmd *cobra.Command, args []string) {
		lib.Sign(PiFile, OiFile, KpcFile)
	},
}

func init() {
	rootCmd.AddCommand(signCmd)

	signCmd.Flags().StringVar(&PiFile, "pi_file", "", "Path to the PI File")
	_ = signCmd.MarkFlagRequired("pi_file")
	signCmd.Flags().StringVar(&OiFile, "oi_file", "", "Path to the OI File")
	_ = signCmd.MarkFlagRequired("oi_file")
	signCmd.Flags().StringVar(&KpcFile, "kpc_file", "", "Customer's private key")
	_ = signCmd.MarkFlagRequired("kpc_file")
}
