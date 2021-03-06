package node

import (
	iface "github.com/ipfs/interface-go-ipfs-core"
	"github.com/joincloud/peers-touch-go/peer"
	"github.com/joincloud/peers-touch-go/pubsub"
	"github.com/libp2p/go-libp2p"
)

const (
	TypeEnd Type = iota + 1
	TypeRegistry
	TypeRelaying
)

const (
	TypeDefault = TypeEnd
)

type Options struct {
	Adds        []string
	Directory   string
	Host        peer.Host
	ID          string
	IPFS        iface.CoreAPI
	PeerID      peer.PeerID
	Broker      pubsub.Broker
	HostOptions []HostOption
	Type        Type
}

type Type int

type Option func(options *Options)

func Host(host peer.Host) Option {
	return func(options *Options) {
		options.Host = host
	}
}

func ID(id string) Option {
	return func(options *Options) {
		options.ID = id
	}
}

func IPFS(is iface.CoreAPI) Option {
	return func(options *Options) {
		options.IPFS = is
	}
}

func Broker(b pubsub.Broker) Option {
	return func(options *Options) {
		options.Broker = b
	}
}

func HostOptions(hps ...HostOption) Option {
	return func(options *Options) {
		options.HostOptions = hps
	}
}

func Typ(typ Type) Option {
	return func(options *Options) {
		options.Type = typ
	}
}

type HostOption = libp2p.Option
