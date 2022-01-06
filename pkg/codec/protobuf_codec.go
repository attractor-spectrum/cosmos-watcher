package watcher

import (
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/gogo/protobuf/proto"

	//injectiveapp "github.com/InjectiveLabs/injective-oracle-scaffold/injective-chain/app"
	cosmoscodectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmoscryptoed "github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cosmoscryptomultisig "github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	cosmoscryptosecp "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cosmoscryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	cosmosibcexported "github.com/cosmos/cosmos-sdk/x/ibc/core/exported"
	cosmosibcclients "github.com/cosmos/cosmos-sdk/x/ibc/light-clients/07-tendermint/types"

	auction "github.com/InjectiveLabs/sdk-go/chain/auction/types"
	keyscodec "github.com/InjectiveLabs/sdk-go/chain/crypto/codec"
	evm "github.com/InjectiveLabs/sdk-go/chain/evm/types"
	exchange "github.com/InjectiveLabs/sdk-go/chain/exchange/types"
	insurance "github.com/InjectiveLabs/sdk-go/chain/insurance/types"
	ocr "github.com/InjectiveLabs/sdk-go/chain/ocr/types"
	oracle "github.com/InjectiveLabs/sdk-go/chain/oracle/types"
	peggy "github.com/InjectiveLabs/sdk-go/chain/peggy/types"
	chaintypes "github.com/InjectiveLabs/sdk-go/chain/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	//authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramproposaltypes "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibcapplicationtypes "github.com/cosmos/ibc-go/v2/modules/apps/transfer/types"
	ibccoretypes "github.com/cosmos/ibc-go/v2/modules/core/types"
	//feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
)

func RegisterInterfacesAndImpls(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	impls := getMessageImplementations()
	interfaceRegistry.RegisterImplementations((*cosmostypes.Msg)(nil), impls...)
	injectiveRegisterInterfaces(interfaceRegistry)
	registerTypes(interfaceRegistry)
}

func injectiveRegisterInterfaces(interfaceRegistry cosmoscodectypes.InterfaceRegistry) {
	//injectiveapp.ModuleBasics.RegisterInterfaces(interfaceRegistry)

	//interfaceRegistry := types.NewInterfaceRegistry()
	keyscodec.RegisterInterfaces(interfaceRegistry)
	std.RegisterInterfaces(interfaceRegistry)
	exchange.RegisterInterfaces(interfaceRegistry)
	insurance.RegisterInterfaces(interfaceRegistry)
	auction.RegisterInterfaces(interfaceRegistry)
	oracle.RegisterInterfaces(interfaceRegistry)
	evm.RegisterInterfaces(interfaceRegistry)
	peggy.RegisterInterfaces(interfaceRegistry)
	ocr.RegisterInterfaces(interfaceRegistry)
	chaintypes.RegisterInterfaces(interfaceRegistry)

	// more cosmos types
	authtypes.RegisterInterfaces(interfaceRegistry)
	//authztypes.RegisterInterfaces(interfaceRegistry)
	vestingtypes.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	crisistypes.RegisterInterfaces(interfaceRegistry)
	distributiontypes.RegisterInterfaces(interfaceRegistry)
	evidencetypes.RegisterInterfaces(interfaceRegistry)
	govtypes.RegisterInterfaces(interfaceRegistry)
	paramproposaltypes.RegisterInterfaces(interfaceRegistry)
	ibcapplicationtypes.RegisterInterfaces(interfaceRegistry)
	ibccoretypes.RegisterInterfaces(interfaceRegistry)
	slashingtypes.RegisterInterfaces(interfaceRegistry)
	stakingtypes.RegisterInterfaces(interfaceRegistry)
	upgradetypes.RegisterInterfaces(interfaceRegistry)
	//feegranttypes.RegisterInterfaces(interfaceRegistry)
}

func registerTypes(interfaceRegistry cosmoscodectypes.InterfaceRegistry) { // todo: need to nest. Maybe we can remove it. Old code
	interfaceRegistry.RegisterInterface("cosmos.crypto.PubKey", (*cosmoscryptotypes.PubKey)(nil))
	interfaceRegistry.RegisterImplementations((*cosmoscryptotypes.PubKey)(nil), &cosmoscryptoed.PubKey{})
	interfaceRegistry.RegisterImplementations((*cosmoscryptotypes.PubKey)(nil), &cosmoscryptosecp.PubKey{})
	interfaceRegistry.RegisterImplementations((*cosmoscryptotypes.PubKey)(nil), &cosmoscryptomultisig.LegacyAminoPubKey{})

	interfaceRegistry.RegisterImplementations((*cosmosibcexported.ClientState)(nil), &cosmosibcclients.ClientState{})
	interfaceRegistry.RegisterImplementations((*cosmosibcexported.ConsensusState)(nil), &cosmosibcclients.ConsensusState{})
	interfaceRegistry.RegisterImplementations((*cosmosibcexported.Header)(nil), &cosmosibcclients.Header{})
	interfaceRegistry.RegisterImplementations((*cosmosibcexported.Misbehaviour)(nil), &cosmosibcclients.Misbehaviour{})
}

func getMessageImplementations() []proto.Message {
	var impls []proto.Message
	cosmosMessages := getCosmosMessages()
	impls = append(impls, cosmosMessages...)
	return impls
}

func getCosmosMessages() []proto.Message {
	cosmosMessages := []proto.Message{
		//&cosmostypes.ServiceMsg{}, // do i need it? cosmostypes.RegisterInterfaces don't exist ServiceMsg
	}
	return cosmosMessages
}
