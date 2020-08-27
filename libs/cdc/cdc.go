package cdc

import (

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/irismod/nft"
	"github.com/irismod/record"
	"github.com/irismod/service"
	"github.com/tendermint/go-amino"
	"github.com/irismod/token"
)

var (
	Cdc          *amino.Codec
	moduleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		service.AppModuleBasic{},
		nft.AppModuleBasic{},
		record.AppModuleBasic{},
		token.AppModuleBasic{},
	)
)

func init() {
	var cdc = codec.New()
	moduleBasics.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	Cdc = cdc.Seal()
}
