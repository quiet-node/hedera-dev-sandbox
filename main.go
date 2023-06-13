package main

import (
	"hedera-dev/initializers"
	"hedera-dev/models"
	"log"
)

var (
	ha *models.HederaAcount
)

func init() {
	// init env vars
	initializers.LoadEnvVars();

	// construct HederaAcount
	ha = initializers.InitHederaAccount()

}

func main() {
	log.Println(ha.AccountId)
	log.Println(ha.PrivateKey)
}