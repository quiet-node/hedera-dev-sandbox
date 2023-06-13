package main

import (
	"hedera-dev/initializers"
	"hedera-dev/models"
	"hedera-dev/queries"
	"log"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

var (
	oha *models.HederaAcount
	hc *hedera.Client
	recipientAccountId hedera.AccountID 
)

func init() {
	// init env vars
	initializers.LoadEnvVars();

	// construct Operator HederaAcount
	oha = initializers.InitHederaAccount()

	// init HederaClient
	hc = initializers.InitHederaClientForTestnet(oha)

	// init recipientAccountId
	recipientAccountId, _ = hedera.AccountIDFromString("0.0.14303758") // "0.0.14303758" is a random accountID created by transactions.CreateNewAccount()

}

func main() {

	// create new account transaction
	// transactions.CreateNewAccount(hc);

	// transfer hbar transaction
	// transactions.TransferHbar(oha.AccountId, recipientAccountId, hc)

	// check balance query
	opBal := queries.GetBalanceQuery(oha.AccountId, hc)
	log.Println("Operator balance:", opBal.Hbars.AsTinybar(), "th")
	reBal := queries.GetBalanceQuery(recipientAccountId, hc)
	log.Println("Recipient balance:", reBal.Hbars.AsTinybar(), "th")

}