package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/empowerchain/empowerchain/x/plasticcredit/types"
)

func (k Keeper) grantAccess(ctx sdk.Context, account sdk.AccAddress, msgType string) error {
	k.setKeeperAccess(ctx, account, msgType)

	return ctx.EventManager().EmitTypedEvent(&types.EventKeeperAccessGranted{
		Account: account.String(),
		MsgType: msgType,
	})
}

func (k Keeper) revokeAccess(ctx sdk.Context, account sdk.AccAddress, msgType string) error {
	k.deleteKeeperAccess(ctx, account, msgType)

	return ctx.EventManager().EmitTypedEvent(&types.EventKeeperAccessRevoked{
		Account: account.String(),
		MsgType: msgType,
	})
}

func (k Keeper) hasAccess(ctx sdk.Context, account sdk.AccAddress, msgType string) bool {
	_, found := k.getKeeperAccess(ctx, account, msgType)
	return found
}

func (k Keeper) setKeeperAccess(ctx sdk.Context, account sdk.AccAddress, msgType string) {
	store := ctx.KVStore(k.storeKey)
	skey := types.KeeperAccessKey(account, msgType)

	b := k.cdc.MustMarshal(&types.KeeperAccess{})
	store.Set(skey, b)
}

func (k Keeper) getKeeperAccess(ctx sdk.Context, account sdk.AccAddress, msgType string) (access types.KeeperAccess, found bool) {
	store := ctx.KVStore(k.storeKey)
	skey := types.KeeperAccessKey(account, msgType)
	b := store.Get(skey)
	if b == nil {
		return access, false
	}
	k.cdc.MustUnmarshal(b, &access)
	return access, true
}

func (k Keeper) deleteKeeperAccess(ctx sdk.Context, account sdk.AccAddress, msgType string) {
	store := ctx.KVStore(k.storeKey)
	skey := types.KeeperAccessKey(account, msgType)
	store.Delete(skey)
}
