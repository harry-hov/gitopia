package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gitopia/gitopia/x/gitopia/types"
	"github.com/gitopia/gitopia/x/gitopia/utils"
)

func (k msgServer) AddRepositoryBackupRef(goCtx context.Context, msg *types.MsgAddRepositoryBackupRef) (*types.MsgAddRepositoryBackupRefResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetUser(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("creator (%v) doesn't exist", msg.Creator))
	}

	address, err := k.ResolveAddress(ctx, msg.RepositoryId.Id)
	if err != nil {
		return nil, err
	}

	repository, found := k.GetAddressRepository(ctx, address.address, msg.RepositoryId.Name)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("repository (%v/%v) doesn't exist", msg.RepositoryId.Id, msg.RepositoryId.Name))
	}

	if !k.HavePermission(ctx, msg.Creator, repository, types.RepositoryBackupPermission) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("user (%v) doesn't have permission to perform this operation", msg.Creator))
	}

	var backup *types.RepositoryBackup
	if i, found := utils.RepositoryBackupExists(repository.Backups, msg.Store); found {
		backup = repository.Backups[i]
	} else {
		backup = new(types.RepositoryBackup)
		backup.Store = msg.Store
		repository.Backups = append(repository.Backups, backup)
	}
	backup.Refs = append(backup.Refs, msg.Ref)

	repository.UpdatedAt = ctx.BlockTime().Unix()
	k.SetRepository(ctx, repository)

	return &types.MsgAddRepositoryBackupRefResponse{}, nil
}

func (k msgServer) UpdateRepositoryBackupRef(goCtx context.Context, msg *types.MsgUpdateRepositoryBackupRef) (*types.MsgUpdateRepositoryBackupRefResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetUser(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("creator (%v) doesn't exist", msg.Creator))
	}

	address, err := k.ResolveAddress(ctx, msg.RepositoryId.Id)
	if err != nil {
		return nil, err
	}

	repository, found := k.GetAddressRepository(ctx, address.address, msg.RepositoryId.Name)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("repository (%v/%v) doesn't exist", msg.RepositoryId.Id, msg.RepositoryId.Name))
	}

	if !k.HavePermission(ctx, msg.Creator, repository, types.RepositoryBackupPermission) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("user (%v) doesn't have permission to perform this operation", msg.Creator))
	}

	var backup *types.RepositoryBackup
	if i, found := utils.RepositoryBackupExists(repository.Backups, msg.Store); found {
		backup = repository.Backups[i]
	} else {
		backup = new(types.RepositoryBackup)
		backup.Store = msg.Store
		repository.Backups = append(repository.Backups, backup)
	}

	if len(backup.Refs) == 0 {
		backup.Refs = []string{msg.Ref}
	} else {
		backup.Refs[0] = msg.Ref
	}

	repository.UpdatedAt = ctx.BlockTime().Unix()
	k.SetRepository(ctx, repository)

	return &types.MsgUpdateRepositoryBackupRefResponse{}, nil
}