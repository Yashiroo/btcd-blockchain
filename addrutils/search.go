package addrutils

import "strings"

func AddrIsIncluded(addy string, slice []string) bool{
	for _, addr := range slice{
		if addy == addr{
			return true
		}
	}
	return false
}

func IsCoinbase(s string,slice []string) bool{
	for _, hash := range slice{
		if strings.ContainsAny(hash,s){
			return true
		}
	}
	return false
}