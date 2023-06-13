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