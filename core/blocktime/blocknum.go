package blocktime

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"main/common/config"
)

func GetBlockNumber(txhash string) uint64 {
	tx, err := config.Client.TransactionReceipt(context.Background(), common.HexToHash(txhash))
	if err != nil {
		log.Fatalf("err in GetBlockTime: %v", err)
	}

	return tx.BlockNumber.Uint64()
}
