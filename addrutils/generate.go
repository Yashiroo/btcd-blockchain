package addrutils

import (
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
)

func GenerateRandomKeyPair() (*btcutil.AddressPubKey, *btcutil.WIF, error){

	net := &chaincfg.MainNetParams
	s256 := btcec.S256()
	priv, err := btcec.NewPrivateKey(s256)
	if err != nil{
		return nil, nil, err
	}
	wif, err := btcutil.NewWIF(priv, net, false)
	if err != nil {
		return nil, nil, err
	}
	pubaddr, err := btcutil.NewAddressPubKey(priv.PubKey().SerializeUncompressed(), net)
	if err != nil {
		return nil, nil, err
	}
	return pubaddr, wif, nil
}
