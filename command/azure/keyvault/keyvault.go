package keyvault

import (
	"github.com/godevopsdev/dvps/command/crypto"
	"github.com/spf13/cobra"
)

// AzureKeyCmd command: Generate a key for Azure Key Vault
var AzureKeyCmd = &cobra.Command{
	Use:   "azureKey [keyname file]",
	Short: "Generate Key Pair that can be load in Azure Key Vault ",
	Args:  cobra.ExactArgs(1),
	Run:   azureKey,
}

func azureKey(cmd *cobra.Command, args []string) {
	keyname := args[0]
	crypto.GenerateRSA2048(keyname)
}
