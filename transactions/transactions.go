package transactions

import (
	"hedera-dev/common"
	"hedera-dev/queries"
	"log"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

// @dev handle creating a new account with an initial balance of 1000 th
func CreateNewAccount(hc *hedera.Client) {
	// Generate new private key
	newAccountPrivateKey, err := hedera.PrivateKeyGenerateEd25519()
	common.HandleException(err)

	// get pub key
	newAccountPubKey := newAccountPrivateKey.PublicKey()

	// prepare transaction creating new account and assign the pub key
	createTransaction := hedera.NewAccountCreateTransaction().SetKey(newAccountPubKey).SetInitialBalance(hedera.HbarFrom(1000, hedera.HbarUnits.Tinybar))

	txRes, err := createTransaction.Execute(hc)
	common.HandleException(err)

	// Request the txReceipt of the transaction
	txReceipt, err := txRes.GetReceipt(hc)
	common.HandleException(err)

	//Get the new account ID from the txReceipt
	newAccountId := txReceipt.AccountID

	//Log the account ID
	log.Printf("The new account ID: %v\n", newAccountId)

	// verify new account balance
	accountBal := queries.GetBalanceQuery(*newAccountId, hc);

	//Print the balance of tinybars
	log.Println("The account balance for the new account:", accountBal.Hbars.AsTinybar(), "th")
}

// @dev handle transfering hbar to a recipient
// 
// @param senderAccountId hedera.AccountID
// 
// @param receiverAccountId hedera.AccountID
// 
// @param client *hedera.Client
func TransferHbar(senderAccountId, receiverAccountId hedera.AccountID, hc *hedera.Client) {
	// prepare transaction
	transferTransaction := hedera.NewTransferTransaction().
						AddHbarTransfer(senderAccountId, hedera.HbarFrom(-1000, hedera.HbarUnits.Tinybar)).
						AddHbarTransfer(receiverAccountId, hedera.HbarFrom(1000, hedera.HbarUnits.Tinybar))

	// execute the transaction
	txRes, err := transferTransaction.Execute(hc)
	common.HandleException(err)

	// Request the txReceipt of the transaction
	txReceipt, err := txRes.GetReceipt(hc)
	common.HandleException(err)

	// verify if the transaction reached consensus by the network based on the status of the receipt
	txStatus := txReceipt.Status
	log.Println("The transfer transaction's consensus status: ", txStatus.String())
	
}