package sendtx

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"main/common/config"
	"main/trace"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func gentx(gaslimit uint64) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA("799bae9074c08e5445fa20f6a4b104ece28933d4dd4241f1fc89e35a335af17d")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := config.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := config.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chianid, _ := config.Client.ChainID(context.Background())
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chianid)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = gaslimit   // in units
	auth.GasPrice = gasPrice
	return auth
}
func AddOrUpdateProdData(id string, ecode string, codata string, operator string, waterdata string) string {
	address := common.HexToAddress(config.TraceSC)
	instance, err := trace.NewTrace(address, config.Client)
	if err != nil {
		log.Fatalf("error creating nftcallerinstance instance:%s", err)
	}
	auth := gentx(200000)
	tx, err := instance.AddOrupdateProdData(auth, id, ecode, codata, operator, waterdata)
	if err != nil {
		fmt.Println("error creating instance")
		log.Fatal(err)
	}
	return tx.Hash().Hex()
	//fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

func GetProdData(id string) (error, string, string, string, string) {
	address := common.HexToAddress(config.TraceSC)
	instance, err := trace.NewTrace(address, config.Client)
	if err != nil {
		//log.Fatalf("error creating nftcallerinstance instance:%s", err)
		return err, "", "", "", ""
	}
	tx, err := instance.Tracedata(nil, id)
	if err != nil {
		return err, "", "", "", ""
	}
	return nil, tx.Waterdata, tx.Codata, tx.Ecode, tx.Operator
}
