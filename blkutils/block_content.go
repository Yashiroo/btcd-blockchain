package blkutils

import (
	"github.com/sirupsen/logrus"
	"github.com/btcsuite/btcd/wire"
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

