package initializers

import (
	"hedera-dev/common"
	"hedera-dev/models"
	"os"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/joho/godotenv"
)

// @dev Loads environment variables
func LoadEnvVars() {
	err := godotenv.Load();
	common.HandleException(err)
}

// @dev handle retrieving accountId and privateKey
// 
// @return *models.HederaAccount
// 
// @resource: https://docs.hedera.com/hedera/getting-started/environment-set-up
func InitHederaAccount() (*models.HederaAcount) {
	// prepare placeholder
	ha := &models.HederaAcount{}

	var err error

	// retrieve accountId
	ha.AccountId, err = hedera.AccountIDFromString(os.Getenv("HED_ACCOUNT_ID"))
	common.HandleException(err)
	
	// retrieve privateKey
	ha.PrivateKey, err = hedera.PrivateKeyFromString(os.Getenv("HED_PRIVATE_KEY"))
	common.HandleException(err)
	
	return ha
}

// @dev handle create a client for test net
// 
// @return *hedera.Client
// 
// @resource: https://docs.hedera.com/hedera/getting-started/environment-set-up
func InitHederaClientForTestnet(ha *models.HederaAcount) *hedera.Client {
	// init client
	hc := hedera.ClientForTestnet()

	// set up operator
	hc.SetOperator(ha.AccountId, ha.PrivateKey)

	// @notice To avoid encountering the INSUFFICIENT_TX_FEE error while conducting 
	// 			transactions,  you can adjust the maximum transaction fee limit through the 
	// 			.setDefaultMaxTransactionFee() method. Similarly, the maximum query payment 
	// 			can be adjusted using the .setDefaultMaxQueryPayment() method. 
	hc.SetDefaultMaxTransactionFee(hedera.HbarFrom(100, hedera.HbarUnits.Hbar))
	hc.SetDefaultMaxQueryPayment(hedera.HbarFrom(50, hedera.HbarUnits.Hbar))

	return hc
}