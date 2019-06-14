package commands

import (
	"github.com/ipfs/go-ipfs-cmds"

	"github.com/filecoin-project/go-filecoin/address"
)

func fromAddr(req *cmds.Request, env cmds.Environment) (address.Address, error) {
	addr, err := optionalAddr(req.Options["from"])
	if err != nil {
		return address.Undef, err
	}
	if addr.Empty() {
		return GetPorcelainAPI(env).WalletDefaultAddress()
	}
	return addr, nil
}
