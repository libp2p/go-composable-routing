package api

import (
	"github.com/libp2p/go-routing-language/syntax"
)

type Router interface {
	Route(syntax.Node) (syntax.Node, error)
}
