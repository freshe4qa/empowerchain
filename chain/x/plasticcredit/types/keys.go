package types

// TODO probably better idea to have it in keeper package so key fetching can be private

import (
	"github.com/cosmos/cosmos-sdk/internal/conv"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	// ModuleName defines the module name
	ModuleName = "plasticcredit"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_plasticcredit"
)

var (
	KeeperAccessKeyPrefix = []byte{0x01}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	IdCountersKey = "IdCounters/value/"
)

func KeeperAccessKey(account sdk.AccAddress, msgType string) []byte {
	m := conv.UnsafeStrToBytes(msgType)
	account = address.MustLengthPrefix(account)
	key := sdk.AppendLengthPrefixedBytes(KeeperAccessKeyPrefix, account, m)

	return key
}
