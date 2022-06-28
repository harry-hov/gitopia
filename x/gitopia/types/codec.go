package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgAuthorizeGitServer{}, "gitopia/AuthorizeGitServer", nil)
	// cdc.RegisterConcrete(&MsgCreateTask{}, "gitopia/CreateTask", nil)
	cdc.RegisterConcrete(&MsgUpdateTask{}, "gitopia/UpdateTask", nil)
	// cdc.RegisterConcrete(&MsgDeleteTask{}, "gitopia/DeleteTask", nil)
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateRelease{}, "gitopia/CreateRelease", nil)
	cdc.RegisterConcrete(&MsgUpdateRelease{}, "gitopia/UpdateRelease", nil)
	cdc.RegisterConcrete(&MsgDeleteRelease{}, "gitopia/DeleteRelease", nil)

	cdc.RegisterConcrete(&MsgCreatePullRequest{}, "gitopia/CreatePullRequest", nil)
	cdc.RegisterConcrete(&MsgUpdatePullRequest{}, "gitopia/UpdatePullRequest", nil)
	cdc.RegisterConcrete(&MsgUpdatePullRequestTitle{}, "gitopia/UpdatePullRequestTitle", nil)
	cdc.RegisterConcrete(&MsgUpdatePullRequestDescription{}, "gitopia/UpdatePullRequestDescription", nil)
	cdc.RegisterConcrete(&MsgSetPullRequestState{}, "gitopia/SetPullRequestState", nil)
	cdc.RegisterConcrete(&MsgAddPullRequestReviewers{}, "gitopia/AddPullRequestReviewers", nil)
	cdc.RegisterConcrete(&MsgRemovePullRequestReviewers{}, "gitopia/RemovePullRequestReviewers", nil)
	cdc.RegisterConcrete(&MsgAddPullRequestAssignees{}, "gitopia/AddPullRequestAssignees", nil)
	cdc.RegisterConcrete(&MsgRemovePullRequestAssignees{}, "gitopia/RemovePullRequestAssignees", nil)
	cdc.RegisterConcrete(&MsgAddPullRequestLabels{}, "gitopia/AddPullRequestLabels", nil)
	cdc.RegisterConcrete(&MsgRemovePullRequestLabels{}, "gitopia/RemovePullRequestLabels", nil)
	cdc.RegisterConcrete(&MsgDeletePullRequest{}, "gitopia/DeletePullRequest", nil)

	cdc.RegisterConcrete(&MsgCreateOrganization{}, "gitopia/CreateOrganization", nil)
	cdc.RegisterConcrete(&MsgRenameOrganization{}, "gitopia/RenameOrganization", nil)
	cdc.RegisterConcrete(&MsgUpdateOrganizationMember{}, "gitopia/UpdateOrganizationMember", nil)
	cdc.RegisterConcrete(&MsgRemoveOrganizationMember{}, "gitopia/RemoveOrganizationMember", nil)
	cdc.RegisterConcrete(&MsgUpdateOrganization{}, "gitopia/UpdateOrganization", nil)
	cdc.RegisterConcrete(&MsgUpdateOrganizationDescription{}, "gitopia/UpdateOrganizationDescription", nil)
	cdc.RegisterConcrete(&MsgUpdateOrganizationAvatar{}, "gitopia/UpdateOrganizationAvatar", nil)
	cdc.RegisterConcrete(&MsgDeleteOrganization{}, "gitopia/DeleteOrganization", nil)

	cdc.RegisterConcrete(&MsgCreateComment{}, "gitopia/CreateComment", nil)
	cdc.RegisterConcrete(&MsgUpdateComment{}, "gitopia/UpdateComment", nil)
	cdc.RegisterConcrete(&MsgDeleteComment{}, "gitopia/DeleteComment", nil)

	cdc.RegisterConcrete(&MsgCreateIssue{}, "gitopia/CreateIssue", nil)
	cdc.RegisterConcrete(&MsgUpdateIssue{}, "gitopia/UpdateIssue", nil)
	cdc.RegisterConcrete(&MsgUpdateIssueTitle{}, "gitopia/UpdateIssueTitle", nil)
	cdc.RegisterConcrete(&MsgUpdateIssueDescription{}, "gitopia/UpdateIssueDescription", nil)
	cdc.RegisterConcrete(&MsgToggleIssueState{}, "gitopia/ToggleIssueState", nil)
	cdc.RegisterConcrete(&MsgAddIssueAssignees{}, "gitopia/AddIssueAssignees", nil)
	cdc.RegisterConcrete(&MsgRemoveIssueAssignees{}, "gitopia/RemoveIssueAssignees", nil)
	cdc.RegisterConcrete(&MsgAddIssueLabels{}, "gitopia/AddIssueLabels", nil)
	cdc.RegisterConcrete(&MsgRemoveIssueLabels{}, "gitopia/RemoveIssueLabels", nil)
	cdc.RegisterConcrete(&MsgDeleteIssue{}, "gitopia/DeleteIssue", nil)

	cdc.RegisterConcrete(&MsgCreateRepository{}, "gitopia/CreateRepository", nil)
	cdc.RegisterConcrete(&MsgForkRepository{}, "gitopia/ForkRepository", nil)
	cdc.RegisterConcrete(&MsgForkRepositorySuccess{}, "gitopia/ForkRepositorySuccess", nil)
	cdc.RegisterConcrete(&MsgRenameRepository{}, "gitopia/RenameRepository", nil)
	cdc.RegisterConcrete(&MsgUpdateRepositoryDescription{}, "gitopia/UpdateRepositoryDescription", nil)
	cdc.RegisterConcrete(&MsgChangeOwner{}, "gitopia/ChangeOwner", nil)
	cdc.RegisterConcrete(&MsgUpdateRepositoryCollaborator{}, "gitopia/UpdateRepositoryCollaborator", nil)
	cdc.RegisterConcrete(&MsgRemoveRepositoryCollaborator{}, "gitopia/RemoveRepositoryCollaborator", nil)
	cdc.RegisterConcrete(&MsgCreateRepositoryLabel{}, "gitopia/CreateRepositoryLabel", nil)
	cdc.RegisterConcrete(&MsgUpdateRepositoryLabel{}, "gitopia/UpdateRepositoryLabel", nil)
	cdc.RegisterConcrete(&MsgDeleteRepositoryLabel{}, "gitopia/DeleteRepositoryLabel", nil)
	cdc.RegisterConcrete(&MsgSetRepositoryBranch{}, "gitopia/SetRepositoryBranch", nil)
	cdc.RegisterConcrete(&MsgMultiSetRepositoryBranch{}, "gitopia/MultiSetRepositoryBranch", nil)
	cdc.RegisterConcrete(&MsgSetDefaultBranch{}, "gitopia/SetDefaultBranch", nil)
	cdc.RegisterConcrete(&MsgDeleteBranch{}, "gitopia/DeleteBranch", nil)
	cdc.RegisterConcrete(&MsgMultiDeleteBranch{}, "gitopia/MultiDeleteBranch", nil)
	cdc.RegisterConcrete(&MsgSetRepositoryTag{}, "gitopia/SetRepositoryTag", nil)
	cdc.RegisterConcrete(&MsgMultiSetRepositoryTag{}, "gitopia/MultiSetRepositoryTag", nil)
	cdc.RegisterConcrete(&MsgDeleteTag{}, "gitopia/DeleteTag", nil)
	cdc.RegisterConcrete(&MsgMultiDeleteTag{}, "gitopia/MultiDeleteTag", nil)
	cdc.RegisterConcrete(&MsgToggleRepositoryForking{}, "gitopia/ToggleRepositoryForking", nil)
	cdc.RegisterConcrete(&MsgDeleteRepository{}, "gitopia/DeleteRepository", nil)

	cdc.RegisterConcrete(&MsgCreateUser{}, "gitopia/CreateUser", nil)
	cdc.RegisterConcrete(&MsgUpdateUser{}, "gitopia/UpdateUser", nil)
	cdc.RegisterConcrete(&MsgUpdateUserBio{}, "gitopia/UpdateUserBio", nil)
	cdc.RegisterConcrete(&MsgUpdateUserAvatar{}, "gitopia/UpdateUserAvatar", nil)
	cdc.RegisterConcrete(&MsgDeleteUser{}, "gitopia/DeleteUser", nil)
	cdc.RegisterConcrete(&MsgTransferUser{}, "gitopia/TransferUser", nil)

	cdc.RegisterConcrete(&MsgSetWhois{}, "gitopia/SetWhois", nil)
	cdc.RegisterConcrete(&MsgUpdateWhois{}, "gitopia/UpdateWhois", nil)
	cdc.RegisterConcrete(&MsgDeleteWhois{}, "gitopia/DeleteWhois", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAuthorizeGitServer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		// &MsgCreateTask{},
		&MsgUpdateTask{},
		// &MsgDeleteTask{},
	)
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateRelease{},
		&MsgUpdateRelease{},
		&MsgDeleteRelease{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePullRequest{},
		&MsgUpdatePullRequest{},
		&MsgUpdatePullRequestTitle{},
		&MsgUpdatePullRequestDescription{},
		&MsgSetPullRequestState{},
		&MsgAddPullRequestReviewers{},
		&MsgRemovePullRequestReviewers{},
		&MsgAddPullRequestAssignees{},
		&MsgRemovePullRequestAssignees{},
		&MsgAddPullRequestLabels{},
		&MsgRemovePullRequestLabels{},
		&MsgDeletePullRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateOrganization{},
		&MsgRenameOrganization{},
		&MsgUpdateOrganizationMember{},
		&MsgRemoveOrganizationMember{},
		&MsgUpdateOrganization{},
		&MsgUpdateOrganizationDescription{},
		&MsgUpdateOrganizationAvatar{},
		&MsgDeleteOrganization{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateComment{},
		&MsgUpdateComment{},
		&MsgDeleteComment{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateIssue{},
		&MsgUpdateIssue{},
		&MsgUpdateIssueTitle{},
		&MsgUpdateIssueDescription{},
		&MsgToggleIssueState{},
		&MsgAddIssueAssignees{},
		&MsgRemoveIssueAssignees{},
		&MsgAddIssueLabels{},
		&MsgRemoveIssueLabels{},
		&MsgDeleteIssue{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateRepository{},
		&MsgForkRepository{},
		&MsgForkRepositorySuccess{},
		&MsgRenameRepository{},
		&MsgUpdateRepositoryDescription{},
		&MsgChangeOwner{},
		&MsgUpdateRepositoryCollaborator{},
		&MsgRemoveRepositoryCollaborator{},
		&MsgCreateRepositoryLabel{},
		&MsgUpdateRepositoryLabel{},
		&MsgDeleteRepositoryLabel{},
		&MsgSetRepositoryBranch{},
		&MsgMultiSetRepositoryBranch{},
		&MsgSetDefaultBranch{},
		&MsgDeleteBranch{},
		&MsgMultiDeleteBranch{},
		&MsgSetRepositoryTag{},
		&MsgMultiSetRepositoryTag{},
		&MsgDeleteTag{},
		&MsgMultiDeleteTag{},
		&MsgToggleRepositoryForking{},
		&MsgDeleteRepository{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateUser{},
		&MsgUpdateUser{},
		&MsgUpdateUserBio{},
		&MsgUpdateUserAvatar{},
		&MsgDeleteUser{},
		&MsgTransferUser{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetWhois{},
		&MsgUpdateWhois{},
		&MsgDeleteWhois{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
