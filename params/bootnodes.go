// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Pirl network.
var MainnetBootnodes = []string{
	// Vantum Foundation Go Bootnodes
  "enode://4f23b87f5453ad15debb45ae0095f3dca529d5040072b39e3fc22ce9e3d8a7220d453bfd7667a8e9592f68bf28ac5378b7bbac1d580dfd941dd40db139026525@178.128.165.202:44443",
  "enode://c597f21ca80e9bea69ca19a70d45bdf04feab91ce1cc2832f60c868319b3a497ff4b6f6131a3fae39e4e299d31baaecb16d30903aa1208b32f1a20c3f1808571@41.58.175.62:44443",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:30303",    // US-Azure geth
	"enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:30303",  // US-Azure parity
	"enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", // Parity
	"enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303", // @gpip
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://4f23b87f5453ad15debb45ae0095f3dca529d5040072b39e3fc22ce9e3d8a7220d453bfd7667a8e9592f68bf28ac5378b7bbac1d580dfd941dd40db139026525@178.128.165.202:44443",
  	"enode://c597f21ca80e9bea69ca19a70d45bdf04feab91ce1cc2832f60c868319b3a497ff4b6f6131a3fae39e4e299d31baaecb16d30903aa1208b32f1a20c3f1808571@41.58.175.62:44443",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://4f23b87f5453ad15debb45ae0095f3dca529d5040072b39e3fc22ce9e3d8a7220d453bfd7667a8e9592f68bf28ac5378b7bbac1d580dfd941dd40db139026525@178.128.165.202:44443",
  	"enode://c597f21ca80e9bea69ca19a70d45bdf04feab91ce1cc2832f60c868319b3a497ff4b6f6131a3fae39e4e299d31baaecb16d30903aa1208b32f1a20c3f1808571@41.58.175.62:44443",
}
