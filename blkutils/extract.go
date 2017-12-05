package blkutils

import (
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/txscript"
	"github.com/go-errors/errors"
)

// ExtractRandomPublicKeyFromPKScript extracts addresses from the given script and returns them
// Note that this function requires the configuration for the chain being used
func (f Fetcher) ExtractRandomPublicKeyFromPKScript(script []byte) ([]btcutil.Address, error){
	//
	_,addresses,_,err := txscript.ExtractPkScriptAddrs(script, &f.networkParams)
	if err != nil{
		return btcutil.AddressPubKey{}, err
	}
	if len(addresses) == 0{
		return btcutil.Address{}, errors.New("Script contains no valid address!")
	}
	return addresses, nil
}
