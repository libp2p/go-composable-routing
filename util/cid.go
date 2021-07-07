package util

import (
	"context"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-composable-routing/api"
	"github.com/libp2p/go-routing-language/patterns"
	"github.com/libp2p/go-routing-language/syntax"
)

// CidFinderFromRouter converts a general router interface to a specific interface for finding CIDs.
func CidFinderFromRouter(router api.Router) api.CidFinder {
	return &cidFinderFromRouter{router}
}

type cidFinderFromRouter struct {
	router api.Router
}

func (x *cidFinderFromRouter) FindCid(ctx context.Context, k cid.Cid) (<-chan syntax.Node, error) {
	p := &patterns.FindCid{Cid: k}
	return x.router.Route(ctx, p.Express())
}
