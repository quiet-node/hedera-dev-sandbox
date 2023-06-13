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

