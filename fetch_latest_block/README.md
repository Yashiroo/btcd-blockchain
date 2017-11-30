#Fetch Latest Block
In order to run this utility, you would need a full running node of `btcd`.

This utility does not connect to the blockchain network directly. Instead, it communicates with the rpc server that comes with btcd and asks for specific information.

**Configuration parameters**:
- Btcd running and rpc server activated (see config options https://godoc.org/github.com/btcsuite/btcd)
- Certificate to connect to the rpc server
- Username and password (as you have specified in btcd's configuration)
- Network to connect to (default is mainnet)