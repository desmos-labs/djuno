package reactions

import (
	"github.com/cosmos/cosmos-sdk/codec"
	reactionstypes "github.com/desmos-labs/desmos/v4/x/reactions/types"
	"github.com/forbole/juno/v3/modules"
	"github.com/forbole/juno/v3/node"
	"google.golang.org/grpc"

	"github.com/desmos-labs/djuno/v2/database"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
	_ modules.MessageModule = &Module{}
)

// Module represents the x/fees module handler
type Module struct {
	cdc             codec.Codec
	db              *database.Db
	node            node.Node
	reactionsClient reactionstypes.QueryClient
}

// NewModule allows to build a new Module instance
func NewModule(node node.Node, grpcConnection *grpc.ClientConn, cdc codec.Codec, db *database.Db) *Module {
	return &Module{
		cdc:             cdc,
		db:              db,
		node:            node,
		reactionsClient: reactionstypes.NewQueryClient(grpcConnection),
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "reactions"
}
