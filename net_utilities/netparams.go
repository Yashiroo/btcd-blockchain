package net_utilities

import (
	"github.com/btcsuite/btcd/chaincfg"
	"strings"
	"github.com/pkg/errors"
)

func NewNetworkParams(network string) (chaincfg.Params,error){
	switch strings.ToLower(network){
	case "mainnet":
		return chaincfg.MainNetParams, nil
	case "testnet":
		return chaincfg.TestNet3Params, nil
	case "simnet":
		return chaincfg.SimNetParams, nil
	case "regtest":
		return chaincfg.RegressionNetParams, nil
	default:
		return chaincfg.Params{}, errors.New("Could not parse network parameter")
	}
}
