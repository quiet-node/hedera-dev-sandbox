package main

import (
	"hedera-dev/initializers"
	"hedera-dev/models"
	"log"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

var (
	ha *models.HederaAcount
	hc *hedera.Client
)

func init() {
	// init env vars
	initializers.LoadEnvVars();

	// construct HederaAcount
	ha = initializers.InitHederaAccount()

	// init HederaClient
	hc = initializers.InitHederaClientForTestnet(ha)

}

func main() {
	log.Println(hc.GetOperatorAccountID())
	log.Println(hc.GetOperatorPublicKey())
	
}