package transactions

import (
	"hedera-dev/common"
	"log"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

// @dev handle creating a new account
func CreateNewAccount(hc *hedera.Client) {
	// Generate new private key
	newAccountPrivateKey, err := hedera.PrivateKeyGenerateEd25519()
	common.HandleException(err)

	// get pub key
	newAccountPubKey := newAccountPrivateKey.PublicKey()

	//Create new account and assign the pub key
	newAccount, err := hedera.NewAccountCreateTransaction().SetKey(newAccountPubKey).SetInitialBalance(hedera.HbarFrom(1000, hedera.HbarUnits.Tinybar)).Execute(hc)
	common.HandleException(err)

	//Request the receipt of the transaction
	receipt, err := newAccount.GetReceipt(hc)
	common.HandleException(err)

	//Get the new account ID from the receipt
	newAccountId := receipt.AccountID

	//Log the account ID
	log.Printf("The new account ID: %v\n", newAccountId)

	// verify new account balance
	accountBal := common.GetBalanceQuery(*newAccountId, hc);

	//Print the balance of tinybars
	log.Println("The account balance for the new account:", accountBal.Hbars.AsTinybar(), "th")
}