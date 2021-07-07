package api

import (
	"context"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-routing-language/syntax"
)

// Router is the general interface of a routing system.
type Router interface {
	Route(context.Context, syntax.Node) (<-chan syntax.Node, error)
}

// CidFinder is an interface to a routing system that specializes in finding providers for a cid.
type CidFinder interface {
	FindCid(context.Context, cid.Cid) (<-chan syntax.Node, error)
}
