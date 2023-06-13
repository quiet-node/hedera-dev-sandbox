package common

import (
	"log"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

// @dev handle error exeptions
func HandleException(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// @dev retrieve account balance
// 
// @param acountId *hedera.AccountID
// 
// @param client *hedera.Client
// 
// @return hedera.
func GetBalanceQuery(accountId hedera.AccountID, hc *hedera.Client) hedera.AccountBalance {
	// prepare query
	query := hedera.NewAccountBalanceQuery().SetAccountID(accountId)

	// Sign with client operator private key and submit the query to a Hedera network
	accountBal, err := query.Execute(hc)
	HandleException(err)

	return accountBal
}