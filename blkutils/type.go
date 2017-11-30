package blkutils

import (
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/go-errors/errors"
)

type Fetcher struct {
	rpcClient *rpcclient.Client
	networkParams chaincfg.Params
}

func NewFetcher(rpcClient *rpcclient.Client, netParams chaincfg.Params) (Fetcher,error){
	if rpcClient == nil{
		return Fetcher{}, errors.New("rpc client cannot be nil!")
	}
	return Fetcher{
		rpcClient:rpcClient,
		networkParams:netParams,
	}, nil
}