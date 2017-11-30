package main

import (
	"io/ioutil"
	"github.com/btcsuite/btcutil"
	"flag"
	"github.com/sirupsen/logrus"
	"runtime"
	"fmt"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/yashirooo/btcd-misc/netparamutils"
	"github.com/yashirooo/btcd-misc/blkutils"
)


func main() {
	crtPath := flag.String("cert","","Path to the .crt file")
	btcdrpc := flag.String("btcdrpc","127.0.0.1:8334","URL to connect to btcd's RPC")
	rpcUser := flag.String("rpcuser","","RPC Username")
	rpcPwd := flag.String("rpcpwd","","RPC Password")
	prefNetwork := flag.String("prefnetwork","mainnet","Network to connect to. Can be mainnet, testnet, simnet or regtest")

	flag.Parse()

	// load network params from preferred network flag
	network, err := netparams.NewNetworkParams(*prefNetwork)
	if err != nil{
		logrus.Fatalf("Could not parse pref network")
	}
	// use all cores
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu)
	// read certificate
	cert,err := ioutil.ReadFile(*crtPath)
	if err != nil{
		logrus.Fatalf("error reading certificate")
	}
	rpcClient, err := rpcclient.New(
		&rpcclient.ConnConfig{
			Host:*btcdrpc,
			DisableTLS:false,
			User:*rpcUser,
			Pass:*rpcPwd,
			DisableAutoReconnect:false,
			DisableConnectOnNew:false,
			Certificates:cert,
			// use websocket
			Endpoint:"ws",
		},
		// add handlers here that you want to use later to receive notifications (later on, you can only use handlers
		// you have activated here)
		&rpcclient.NotificationHandlers{
			OnAccountBalance: func(account string, balance btcutil.Amount, confirmed bool) {
				logrus.Printf("New balance for account %s: %v", account,
					balance)
			},
		})
	if err != nil {
		logrus.Fatalf("error creating btcrpc client: %s",err)
	}
	defer rpcClient.Shutdown()


	// prepare our fetcher
	fetcher, err := blkutils.NewFetcher(rpcClient,network)
	if err != nil{
		logrus.Fatalf("Error creating a new fetcher: %s",err)
	}
	// Retrieve the current number of blocks in the bitcoin blockchain
	blockCount,err := rpcClient.GetBlockCount()
	if err != nil{
		logrus.Fatalf("Error retrieving block count: %s",err)
	}
	fmt.Printf("Block count is currently: %d\n",blockCount)

	// fetch latest block
	block, err := fetcher.RetrieveLatestBlock()
	if err != nil{
		logrus.Fatalf("Error retrieving latest block: %s",err)
	}
	//print hash of first transaction in this block
	if len(block.Transactions) == 0{
		fmt.Println("This block has no transactions!")
		return
	}
	fmt.Printf("First tx hash: %s",block.Transactions[0].TxHash().String())
}
