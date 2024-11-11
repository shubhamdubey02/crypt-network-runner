package api

import (
	"github.com/MetalBlockchain/coreth/plugin/evm"
	"github.com/MetalBlockchain/metalgo/api/admin"
	"github.com/MetalBlockchain/metalgo/api/health"
	"github.com/MetalBlockchain/metalgo/api/info"
	"github.com/MetalBlockchain/metalgo/api/ipcs"
	"github.com/MetalBlockchain/metalgo/api/keystore"
	"github.com/MetalBlockchain/metalgo/indexer"
	"github.com/MetalBlockchain/metalgo/vms/avm"
	"github.com/MetalBlockchain/metalgo/vms/platformvm"
)

// Issues API calls to a node
// TODO: byzantine api. check if appropriate. improve implementation.
type Client interface {
	PChainAPI() platformvm.Client
	XChainAPI() avm.Client
	XChainWalletAPI() avm.WalletClient
	CChainAPI() evm.Client
	CChainEthAPI() EthClient // ethclient websocket wrapper that adds mutexed calls, and lazy conn init (on first call)
	InfoAPI() info.Client
	HealthAPI() health.Client
	IpcsAPI() ipcs.Client
	KeystoreAPI() keystore.Client
	AdminAPI() admin.Client
	PChainIndexAPI() indexer.Client
	CChainIndexAPI() indexer.Client
	// TODO add methods
}
