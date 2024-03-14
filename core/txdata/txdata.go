package txdata

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"main/common/config"
	"time"
)

func Gettxbyhash(hash string) (error, uint64, string, uint64, time.Time, string, uint64, uint64) {
	chainid, _ := config.Client.ChainID(context.Background())
	latestnum, _ := config.Client.BlockNumber(context.Background())
	txrecipert, err := config.Client.TransactionReceipt(context.Background(), common.HexToHash(hash))
	if err != nil {
		return err, 0, "", 0, time.Now(), "", 0, 0
	}
	blocknum := txrecipert.BlockNumber.Uint64()
	blockhash := txrecipert.BlockHash.String()
	gas := txrecipert.GasUsed
	tx, _, _ := config.Client.TransactionByHash(context.Background(), common.HexToHash(hash))
	txtime := tx.Time().UTC()
	contract := tx.To().Hex()

	return nil, chainid.Uint64(), blockhash, blocknum, txtime, contract, gas, latestnum - blocknum
}
