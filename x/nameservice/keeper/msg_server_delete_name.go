package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"nameservice/x/nameservice/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

)

func (k msgServer) DeleteName(goCtx context.Context, msg *types.MsgDeleteName) (*types.MsgDeleteNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	whois, isFound := k.GetWhois(ctx, msg.Name)

	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Name doesn't exist")
	}

	if !(whois.Owner == msg.Creator) {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
    }

    // Otherwise, remove the name information from the store
    k.RemoveWhois(ctx, msg.Name)
    return &types.MsgDeleteNameResponse{}, nil
}
