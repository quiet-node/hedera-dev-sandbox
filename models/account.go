package models

import "github.com/hashgraph/hedera-sdk-go/v2"

// @notice information of the struct HederaAcount
//
// @param AccountId hedera.AccountID
//
// @param PrivateKey hedera.PrivateKey
type HederaAcount struct {
	AccountId hedera.AccountID
	PrivateKey hedera.PrivateKey
}