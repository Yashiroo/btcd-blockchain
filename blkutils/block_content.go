package blkutils

import (
	"github.com/sirupsen/logrus"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/yashirooo/btcd-misc/addrutils"
	"github.com/go-errors/errors"
	"github.com/hyperledger/fabric/gossip/util"
)

// RetrieveLatestBlock retrieves the latest validated block on the current chain
func (f Fetcher) RetrieveLatestBlock() (*wire.MsgBlock,error){
	hash,num, err := f.rpcClient.GetBestBlock()
	if err != nil{
		return nil, err
	}
	logrus.Infof("The best block is %s at height %d",hash.String(), num)

	block, err := f.rpcClient.GetBlock(hash)
	if err != nil{
		return nil, err
	}


	return block, nil
}


//RetrieveBlockAtHeight retrieves the block at the given height
func (f Fetcher) RetrieveBlockAtHeight(height int64) (*wire.MsgBlock, error){
	hash, err := f.rpcClient.GetBlockHash(height)
	if err != nil{
		return nil, err
	}
	block, err := f.rpcClient.GetBlock(hash)
	if err != nil{
		return nil, err
	}
	return block, nil
}

//ShowPeerInfo shows basic information for connected peers
func (f Fetcher) ShowPeerInfo(){
	peerInfo, err := f.rpcClient.GetPeerInfo()
	if err != nil{
		logrus.Fatalf("error getting peers info: %s",err)
	}
	for _, peer := range peerInfo{
		logrus.Infof("Peer services: %s",peer.Services)
		if peer.BanScore != 0{
			logrus.Infof("Peer ban score: %d",peer.BanScore)
		}
	}
}

//CalculateBalanceFor calculates the balance for a given pay-to address
func (f Fetcher) CalculateConfirmedBalanceFor(address btcutil.Address) (btcutil.Amount, error){

	//filter needed to search for transactions
	var filteraddrs []string = []string{address.EncodeAddress()}
	logrus.Infof("Looking for transactions for %s ..",address.EncodeAddress())
	//search for transactions involving the given address
	searchRawResult, err := f.rpcClient.SearchRawTransactionsVerbose(address,0,1000000,true,false,filteraddrs)
	if err != nil{
		logrus.Errorf("Error retrieving raw transactions: %s",err)
	}
	if len(searchRawResult) == 0{
		return 0, errors.New("No transactions found for given address!")
	}
	logrus.Infof("Found %d transactions involving %s",len(searchRawResult),address.EncodeAddress())
	var sumin btcutil.Amount
	var sumout btcutil.Amount
	//while looping over txs, let's not forget that there may be immature coinbase txs
	// we need to also take into account regular transactions that have not had over 6 confirmations
	for _,result := range searchRawResult {
		for _, in := range result.Vin {
			if !in.IsCoinBase() {
				if addrutils.AddrIsIncluded(address.EncodeAddress(), in.PrevOut.Addresses) {
					logrus.Infof("adding input %f", in.PrevOut.Value)
					amount, err := btcutil.NewAmount(in.PrevOut.Value)
					if err != nil {
						continue
					}
					sumout += amount
				}

			}
		}
		for _, out := range result.Vout {
			if addrutils.AddrIsIncluded(address.EncodeAddress(), out.ScriptPubKey.Addresses) {
				logrus.Infof("adding output %f", out.Value)
				amount, err := btcutil.NewAmount(out.Value)
				if err != nil {
					continue
				}
				sumin += amount
			}

		}
	}
	balance := sumin - sumout

	return balance, nil
}

// RandomAddressFromBlock fetches a random address from the first transaction in this block
func (f Fetcher) RandomAddressFromBlock(block *wire.MsgBlock) (btcutil.Address, error){
	//check if there are inputs
	if len(block.Transactions) != 0{
		randomTXNum := util.RandomInt(len(block.Transactions))
		tx := block.Transactions[randomTXNum]
		if len(tx.TxOut) != 0{
			randomTxOutNum := util.RandomInt(len(tx.TxOut))
			txOut := tx.TxOut[randomTxOutNum]
			addresses, err := f.ExtractRandomPublicKeyFromPKScript(txOut.PkScript)
			if err != nil{
				return nil, err
			}
			return addresses[len(addresses)/2], nil
		}
	}

	return nil, errors.New("No transactions in block!")
}


//this function creates a new blockchain and then fetches a block from the blockchain and attempts to process it
//func (f Fetcher) ProcessBlock(height int64)  {
//
//	blk, err := f.RetrieveBlockAtHeight(height)
//	if err != nil{
//		log.Fatalf("error retrieving block: %s",err)
//	}
//
//	//convert block to btcutil.block
//	block := btcutil.NewBlock(blk)
//
//	// Process a block.
//	isMainChain, isOrphan, err := f.Chain.ProcessBlock(block,
//		blockchain.BFNone)
//	if err != nil {
//		fmt.Printf("Failed to process block: %v\n", err)
//		return
//	}
//	fmt.Printf("Block accepted. Is it on the main chain?: %v\n", isMainChain)
//	fmt.Printf("Block accepted. Is it an orphan?: %v\n", isOrphan)
//}

