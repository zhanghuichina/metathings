// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: add_role_to_group_on_domain.proto

/*
Package identity is a generated protocol buffer package.

It is generated from these files:
	add_role_to_group_on_domain.proto
	add_role_to_group_on_project.proto
	add_role_to_user_on_domain.proto
	add_role_to_user_on_project.proto
	add_user_to_group.proto
	application_credential.proto
	change_password.proto
	check_role_in_group_on_domain.proto
	check_role_in_group_on_project.proto
	check_role_in_user_on_domain.proto
	check_role_in_user_on_project.proto
	check_user_in_group.proto
	create_application_credential.proto
	create_domain.proto
	create_group.proto
	create_project.proto
	create_region.proto
	create_role.proto
	create_user.proto
	delete_application_credential.proto
	delete_domain.proto
	delete_group.proto
	delete_project.proto
	delete_region.proto
	delete_role.proto
	delete_user.proto
	domain.proto
	get_application_credential.proto
	get_domain.proto
	get_group.proto
	get_project.proto
	get_region.proto
	get_role.proto
	get_user.proto
	group.proto
	issue_token.proto
	list_application_credentials.proto
	list_domains.proto
	list_groups.proto
	list_groups_for_user.proto
	list_projects.proto
	list_projects_for_user.proto
	list_regions.proto
	list_role_in_group_on_domain.proto
	list_roles.proto
	list_roles_for_group_on_domain.proto
	list_roles_for_group_on_project.proto
	list_roles_for_user_on_domain.proto
	list_roles_for_user_on_project.proto
	list_users.proto
	list_users_in_group.proto
	patch_domain.proto
	patch_group.proto
	patch_project.proto
	patch_region.proto
	patch_role.proto
	patch_user.proto
	project.proto
	region.proto
	remove_role_from_group_on_domain.proto
	remove_role_from_group_on_project.proto
	remove_role_from_user_on_domain.proto
	remove_role_from_user_on_project.proto
	remove_user_from_group.proto
	role.proto
	service.proto
	token.proto
	user.proto

It has these top-level messages:
	AddRoleToGroupOnDomainRequest
	AddRoleToGroupOnProjectRequest
	AddRoleToUserOnDomainRequest
	AddRoleToUserOnProjectRequest
	AddUserToGroupRequest
	ApplicationCredential
	ChangePasswordRequest
	CheckRoleInGroupOnDomainRequest
	CheckRoleInGroupOnProjectRequest
	CheckRoleInUserOnDomainRequest
	CheckRoleInUserOnProjectRequest
	CheckUserInGroupRequest
	CreateApplicationCredentialRequest
	CreateApplicationCredentialResponse
	CreateDomainRequest
	CreateDomainResponse
	CreateGroupRequest
	CreateGroupResponse
	CreateProjectRequest
	CreateProjectResponse
	CreateRegionRequest
	CreateRegionResponse
	CreateRoleRequest
	CreateRoleResponse
	CreateUserRequest
	CreateUserResponse
	DeleteApplicationCredentialRequest
	DeleteDomainRequest
	DeleteGroupRequest
	DeleteProjectRequest
	DeleteRegionRequest
	DeleteRoleRequest
	DeleteUserRequest
	Domain
	GetApplicationCredentialRequest
	GetApplicationCredentialResponse
	GetDomainRequest
	GetDomainResponse
	GetGroupRequest
	GetGroupResponse
	GetProjectRequest
	GetProjectResponse
	GetRegionRequest
	GetRegionResponse
	GetRoleRequest
	GetRoleResponse
	GetUserRequest
	GetUserResponse
	Group
	TokenScope
	PasswordPayload
	TokenPayload
	ApplicationCredentialPayload
	IssueTokenRequest
	IssueTokenResponse
	ListApplicationCredentialsRequest
	ListApplicationCredentialsResponse
	ListDomainsRequest
	ListDomainsResponse
	ListGroupsRequest
	ListGroupsResponse
	ListGroupsForUserRequest
	ListGroupsForUserResponse
	ListProjectsRequest
	ListProjectsResponse
	ListProjectsForUserRequest
	ListProjectsForUserResponse
	ListRegionsRequest
	ListRegionsResponse
	ListRoleInGroupOnDomainRequest
	ListRoleInGroupOnDomainResponse
	ListRolesRequest
	ListRolesResponse
	ListRolesForGroupOnDomainRequest
	ListRolesForGroupOnDomainResponse
	ListRolesForGroupOnProjectRequest
	ListRolesForGroupOnProjectResponse
	ListRolesForUserOnDomainRequest
	ListRolesForUserOnDomainResponse
	ListRolesForUserOnProjectRequest
	ListRolesForUserOnProjectResponse
	ListUsersRequest
	ListUsersResponse
	ListUsersInGroupRequest
	ListUsersInGroupResponse
	PatchDomainRequest
	PatchDomainResponse
	PatchGroupRequest
	PatchGroupResponse
	PatchProjectRequest
	PatchProjectResponse
	PatchRegionRequest
	PatchRegionResponse
	PatchRoleRequest
	PatchRoleResponse
	PatchUserRequest
	PatchUserResponse
	Project
	Region
	RemoveRoleFromGroupOnDomainRequest
	RemoveRoleFromGroupOnProjectRequest
	RemoveRoleFromUserOnDomainRequest
	RemoveRoleFromUserOnProjectRequest
	RemoveUserFromGroupRequest
	Role
	Token
	User
*/
package identity

import fmt "fmt"
import go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *AddRoleToGroupOnDomainRequest) Validate() error {
	if nil == this.DomainId {
		return go_proto_validators.FieldError("DomainId", fmt.Errorf("message must exist"))
	}
	if this.DomainId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.DomainId); err != nil {
			return go_proto_validators.FieldError("DomainId", err)
		}
	}
	if nil == this.GroupId {
		return go_proto_validators.FieldError("GroupId", fmt.Errorf("message must exist"))
	}
	if this.GroupId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.GroupId); err != nil {
			return go_proto_validators.FieldError("GroupId", err)
		}
	}
	if nil == this.RoleId {
		return go_proto_validators.FieldError("RoleId", fmt.Errorf("message must exist"))
	}
	if this.RoleId != nil {
		if err := go_proto_validators.CallValidatorIfExists(this.RoleId); err != nil {
			return go_proto_validators.FieldError("RoleId", err)
		}
	}
	return nil
}
