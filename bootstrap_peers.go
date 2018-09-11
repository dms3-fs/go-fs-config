package config

import (
	"errors"
	"fmt"

	iaddr "github.com/dms3-fs/go-fs-addr"
	// Needs to be imported so that users can import this package directly
	// and still parse the bootstrap addresses.
	_ "github.com/dms3-mft/go-multiaddr-dns"
)

// DefaultBootstrapAddresses are the hardcoded bootstrap addresses
// for DMS3FS. they are nodes run by the DMS3FS team. docs on these later.
// As with all p2p networks, bootstrap is an important security concern.
//
// NOTE: This is here -- and not inside cmd/dms3fs/init.go -- because of an
// import dependency issue. TODO: move this into a config/default/ package.
var DefaultBootstrapAddresses = []string{
	// TODO: change these addresses for dms3:
	//      use dms3 domain DNS entries
	//      change mount mount point to dms3
	//      change node address hash
	"/ip4/50.241.107.73/tcp/4101/dms3fs/QmXg5LqE4UMcx5212oD8UjY8DYGxkGg2Z8PzWHwR5F3XEg",            // www.vectavi.com
	"/ip4/50.241.107.74/tcp/4101/dms3fs/Qma628Kv2DpZhPhupP8C14FJLRYadPWEhWB9PaBfyVFPSS",            // maps.vectavi.com
	"/ip4/50.241.107.75/tcp/4101/dms3fs/QmRbhHwLo6uoHcEMnnnf8u99RB9ubetUUsuhdQXpnD8f8H",            // geos.vectavi.com
	"/ip4/50.241.107.76/tcp/4101/dms3fs/QmYxbm54mnRMnMwuJqBEEDjcq1VDL3JLffSNajWWaKw56X",            // metircs.vectavi.com
																								  // moe.i.dms3.io
																								  // larry.i.dms3.io
																								  // curly.i.dms3.io
																								  // howard.i.dms3.io
/*
	"/dnsaddr/bootstrap.libp2p.io/ipfs/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
	"/dnsaddr/bootstrap.libp2p.io/ipfs/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
	"/dnsaddr/bootstrap.libp2p.io/ipfs/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
	"/dnsaddr/bootstrap.libp2p.io/ipfs/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",
	"/ip4/104.131.131.82/tcp/4001/ipfs/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",            // mars.i.ipfs.io
	"/ip4/104.236.179.241/tcp/4001/ipfs/QmSoLPppuBtQSGwKDZT2M73ULpjvfd3aZ6ha4oFGL1KrGM",           // pluto.i.ipfs.io
	"/ip4/128.199.219.111/tcp/4001/ipfs/QmSoLSafTMBsPKadTEgaXctDQVcqN88CNLHXMkTNwMKPnu",           // saturn.i.ipfs.io
	"/ip4/104.236.76.40/tcp/4001/ipfs/QmSoLV4Bbm51jM9C4gDYZQ9Cy3U6aXMJDAbzgu2fzaDs64",             // venus.i.ipfs.io
	"/ip4/178.62.158.247/tcp/4001/ipfs/QmSoLer265NRgSp2LA3dPaeykiS1J6DifTC88f5uVQKNAd",            // earth.i.ipfs.io
	"/ip6/2604:a880:1:20::203:d001/tcp/4001/ipfs/QmSoLPppuBtQSGwKDZT2M73ULpjvfd3aZ6ha4oFGL1KrGM",  // pluto.i.ipfs.io
	"/ip6/2400:6180:0:d0::151:6001/tcp/4001/ipfs/QmSoLSafTMBsPKadTEgaXctDQVcqN88CNLHXMkTNwMKPnu",  // saturn.i.ipfs.io
	"/ip6/2604:a880:800:10::4a:5001/tcp/4001/ipfs/QmSoLV4Bbm51jM9C4gDYZQ9Cy3U6aXMJDAbzgu2fzaDs64", // venus.i.ipfs.io
	"/ip6/2a03:b0c0:0:1010::23:1001/tcp/4001/ipfs/QmSoLer265NRgSp2LA3dPaeykiS1J6DifTC88f5uVQKNAd", // earth.i.ipfs.io
*/
}

// BootstrapPeer is a peer used to bootstrap the network.
type BootstrapPeer iaddr.DMS3FSAddr

// ErrInvalidPeerAddr signals an address is not a valid peer address.
var ErrInvalidPeerAddr = errors.New("invalid peer address")

func (c *Config) BootstrapPeers() ([]BootstrapPeer, error) {
	return ParseBootstrapPeers(c.Bootstrap)
}

// DefaultBootstrapPeers returns the (parsed) set of default bootstrap peers.
// if it fails, it returns a meaningful error for the user.
// This is here (and not inside cmd/dms3fs/init) because of module dependency problems.
func DefaultBootstrapPeers() ([]BootstrapPeer, error) {
	ps, err := ParseBootstrapPeers(DefaultBootstrapAddresses)
	if err != nil {
		return nil, fmt.Errorf(`failed to parse hardcoded bootstrap peers: %s
This is a problem with the dms3fs codebase. Please report it to the dev team.`, err)
	}
	return ps, nil
}

func (c *Config) SetBootstrapPeers(bps []BootstrapPeer) {
	c.Bootstrap = BootstrapPeerStrings(bps)
}

func ParseBootstrapPeer(addr string) (BootstrapPeer, error) {
	ia, err := iaddr.ParseString(addr)
	if err != nil {
		return nil, err
	}
	return BootstrapPeer(ia), err
}

func ParseBootstrapPeers(addrs []string) ([]BootstrapPeer, error) {
	peers := make([]BootstrapPeer, len(addrs))
	var err error
	for i, addr := range addrs {
		peers[i], err = ParseBootstrapPeer(addr)
		if err != nil {
			return nil, err
		}
	}
	return peers, nil
}

func BootstrapPeerStrings(bps []BootstrapPeer) []string {
	bpss := make([]string, len(bps))
	for i, p := range bps {
		bpss[i] = p.String()
	}
	return bpss
}
