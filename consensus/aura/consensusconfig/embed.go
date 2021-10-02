package consensusconfig

import (
	_ "embed"
	"github.com/ledgerwatch/erigon/params"
)

//go:embed poasokol.json
var Sokol []byte

//go:embed kovan.json
var Kovan []byte

//go:embed fermion.json
var Fermion []byte

//go:embed sparta.json
var Sparta []byte

//go:embed olympus.json
var Olympus []byte

func GetConfigByChain(chainName string) []byte {
	switch chainName {
	case params.SokolChainName:
		return Sokol
	case params.KovanChainName:
		return Kovan
	case params.FermionChainName:
		return Fermion
	case params.SpartaChainName:
		return Sparta
	case params.OlympusChainName:
		return Olympus
	default:
		return Sokol
	}
}
