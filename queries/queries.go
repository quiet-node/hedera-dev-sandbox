package queries

import (
	"hedera-dev/common"

	"github.com/hashgraph/hedera-sdk-go/v2"
)

// @dev retrieve account balance
//
// @param acountId *hedera.AccountID
//
// @param client *hedera.Client
//
// @return hedera.AccountBalance
func GetBalanceQuery(accountId hedera.AccountID, hc *hedera.Client) hedera.AccountBalance {
	// prepare query
	query := hedera.NewAccountBalanceQuery().SetAccountID(accountId)

	// Sign with client operator private key and submit the query to a Hedera network
	accountBal, err := query.Execute(hc)
	common.HandleException(err)

	return accountBal
}
