package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gitopia/gitopia/x/gitopia/types"
)

func (k msgServer) CreateUser(goCtx context.Context, msg *types.MsgCreateUser) (*types.MsgCreateUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendUser(
		ctx,
		msg.Creator,
		msg.Username,
		msg.UsernameGithub,
		msg.AvatarUrl,
		msg.Followers,
		msg.Following,
		msg.Repositories,
		msg.Repositories_archived,
		msg.Organizations,
		msg.Starred_repos,
		msg.Subscriptions,
		msg.Email,
		msg.Bio,
		msg.CreatedAt,
		msg.UpdatedAt,
		msg.Extensions,
	)

	return &types.MsgCreateUserResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateUser(goCtx context.Context, msg *types.MsgUpdateUser) (*types.MsgUpdateUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var user = types.User{
		Creator:               msg.Creator,
		Id:                    msg.Id,
		Username:              msg.Username,
		UsernameGithub:        msg.UsernameGithub,
		AvatarUrl:             msg.AvatarUrl,
		Followers:             msg.Followers,
		Following:             msg.Following,
		Repositories:          msg.Repositories,
		Repositories_archived: msg.Repositories_archived,
		Organizations:         msg.Organizations,
		Starred_repos:         msg.Starred_repos,
		Subscriptions:         msg.Subscriptions,
		Email:                 msg.Email,
		Bio:                   msg.Bio,
		CreatedAt:             msg.CreatedAt,
		UpdatedAt:             msg.UpdatedAt,
		Extensions:            msg.Extensions,
	}

	// Checks that the element exists
	if !k.HasUser(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetUserOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetUser(ctx, user)

	return &types.MsgUpdateUserResponse{}, nil
}

func (k msgServer) DeleteUser(goCtx context.Context, msg *types.MsgDeleteUser) (*types.MsgDeleteUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasUser(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetUserOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveUser(ctx, msg.Id)

	return &types.MsgDeleteUserResponse{}, nil
}
