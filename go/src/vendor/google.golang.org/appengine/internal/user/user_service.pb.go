// Code generated by protoc-gen-go.
// source: google.golang.org/appengine/internal/user/user_service.proto
// DO NOT EDIT!

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/appengine/internal/user/user_service.proto

It has these top-level messages:
	UserServiceError
	CreateLoginURLRequest
	CreateLoginURLResponse
	CreateLogoutURLRequest
	CreateLogoutURLResponse
	GetOAuthUserRequest
	GetOAuthUserResponse
	CheckOAuthSignatureRequest
	CheckOAuthSignatureResponse
*/
package user

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type UserServiceError_ErrorCode int32

const (
	UserServiceError_OK                    UserServiceError_ErrorCode = 0
	UserServiceError_REDIRECT_URL_TOO_LONG UserServiceError_ErrorCode = 1
	UserServiceError_NOT_ALLOWED           UserServiceError_ErrorCode = 2
	UserServiceError_OAUTH_INVALID_TOKEN   UserServiceError_ErrorCode = 3
	UserServiceError_OAUTH_INVALID_REQUEST UserServiceError_ErrorCode = 4
	UserServiceError_OAUTH_ERROR           UserServiceError_ErrorCode = 5
)

var UserServiceError_ErrorCode_name = map[int32]string{
	0: "OK",
	1: "REDIRECT_URL_TOO_LONG",
	2: "NOT_ALLOWED",
	3: "OAUTH_INVALID_TOKEN",
	4: "OAUTH_INVALID_REQUEST",
	5: "OAUTH_ERROR",
}
var UserServiceError_ErrorCode_value = map[string]int32{
	"OK": 0,
	"REDIRECT_URL_TOO_LONG": 1,
	"NOT_ALLOWED":           2,
	"OAUTH_INVALID_TOKEN":   3,
	"OAUTH_INVALID_REQUEST": 4,
	"OAUTH_ERROR":           5,
}

func (x UserServiceError_ErrorCode) Enum() *UserServiceError_ErrorCode {
	p := new(UserServiceError_ErrorCode)
	*p = x
	return p
}
func (x UserServiceError_ErrorCode) String() string {
	return proto.EnumName(UserServiceError_ErrorCode_name, int32(x))
}
func (x *UserServiceError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(UserServiceError_ErrorCode_value, data, "UserServiceError_ErrorCode")
	if err != nil {
		return err
	}
	*x = UserServiceError_ErrorCode(value)
	return nil
}

type UserServiceError struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *UserServiceError) Reset()         { *m = UserServiceError{} }
func (m *UserServiceError) String() string { return proto.CompactTextString(m) }
func (*UserServiceError) ProtoMessage()    {}

type CreateLoginURLRequest struct {
	DestinationUrl    *string `protobuf:"bytes,1,req,name=destination_url" json:"destination_url,omitempty"`
	AuthDomain        *string `protobuf:"bytes,2,opt,name=auth_domain" json:"auth_domain,omitempty"`
	FederatedIdentity *string `protobuf:"bytes,3,opt,name=federated_identity,def=" json:"federated_identity,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *CreateLoginURLRequest) Reset()         { *m = CreateLoginURLRequest{} }
func (m *CreateLoginURLRequest) String() string { return proto.CompactTextString(m) }
func (*CreateLoginURLRequest) ProtoMessage()    {}

func (m *CreateLoginURLRequest) GetDestinationUrl() string {
	if m != nil && m.DestinationUrl != nil {
		return *m.DestinationUrl
	}
	return ""
}

func (m *CreateLoginURLRequest) GetAuthDomain() string {
	if m != nil && m.AuthDomain != nil {
		return *m.AuthDomain
	}
	return ""
}

func (m *CreateLoginURLRequest) GetFederatedIdentity() string {
	if m != nil && m.FederatedIdentity != nil {
		return *m.FederatedIdentity
	}
	return ""
}

type CreateLoginURLResponse struct {
	LoginUrl         *string `protobuf:"bytes,1,req,name=login_url" json:"login_url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateLoginURLResponse) Reset()         { *m = CreateLoginURLResponse{} }
func (m *CreateLoginURLResponse) String() string { return proto.CompactTextString(m) }
func (*CreateLoginURLResponse) ProtoMessage()    {}

func (m *CreateLoginURLResponse) GetLoginUrl() string {
	if m != nil && m.LoginUrl != nil {
		return *m.LoginUrl
	}
	return ""
}

type CreateLogoutURLRequest struct {
	DestinationUrl   *string `protobuf:"bytes,1,req,name=destination_url" json:"destination_url,omitempty"`
	AuthDomain       *string `protobuf:"bytes,2,opt,name=auth_domain" json:"auth_domain,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateLogoutURLRequest) Reset()         { *m = CreateLogoutURLRequest{} }
func (m *CreateLogoutURLRequest) String() string { return proto.CompactTextString(m) }
func (*CreateLogoutURLRequest) ProtoMessage()    {}

func (m *CreateLogoutURLRequest) GetDestinationUrl() string {
	if m != nil && m.DestinationUrl != nil {
		return *m.DestinationUrl
	}
	return ""
}

func (m *CreateLogoutURLRequest) GetAuthDomain() string {
	if m != nil && m.AuthDomain != nil {
		return *m.AuthDomain
	}
	return ""
}

type CreateLogoutURLResponse struct {
	LogoutUrl        *string `protobuf:"bytes,1,req,name=logout_url" json:"logout_url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateLogoutURLResponse) Reset()         { *m = CreateLogoutURLResponse{} }
func (m *CreateLogoutURLResponse) String() string { return proto.CompactTextString(m) }
func (*CreateLogoutURLResponse) ProtoMessage()    {}

func (m *CreateLogoutURLResponse) GetLogoutUrl() string {
	if m != nil && m.LogoutUrl != nil {
		return *m.LogoutUrl
	}
	return ""
}

type GetOAuthUserRequest struct {
	Scope            *string  `protobuf:"bytes,1,opt,name=scope" json:"scope,omitempty"`
	Scopes           []string `protobuf:"bytes,2,rep,name=scopes" json:"scopes,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *GetOAuthUserRequest) Reset()         { *m = GetOAuthUserRequest{} }
func (m *GetOAuthUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetOAuthUserRequest) ProtoMessage()    {}

func (m *GetOAuthUserRequest) GetScope() string {
	if m != nil && m.Scope != nil {
		return *m.Scope
	}
	return ""
}

func (m *GetOAuthUserRequest) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

type GetOAuthUserResponse struct {
	Email            *string  `protobuf:"bytes,1,req,name=email" json:"email,omitempty"`
	UserId           *string  `protobuf:"bytes,2,req,name=user_id" json:"user_id,omitempty"`
	AuthDomain       *string  `protobuf:"bytes,3,req,name=auth_domain" json:"auth_domain,omitempty"`
	UserOrganization *string  `protobuf:"bytes,4,opt,name=user_organization,def=" json:"user_organization,omitempty"`
	IsAdmin          *bool    `protobuf:"varint,5,opt,name=is_admin,def=0" json:"is_admin,omitempty"`
	ClientId         *string  `protobuf:"bytes,6,opt,name=client_id,def=" json:"client_id,omitempty"`
	Scopes           []string `protobuf:"bytes,7,rep,name=scopes" json:"scopes,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *GetOAuthUserResponse) Reset()         { *m = GetOAuthUserResponse{} }
func (m *GetOAuthUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetOAuthUserResponse) ProtoMessage()    {}

const Default_GetOAuthUserResponse_IsAdmin bool = false

func (m *GetOAuthUserResponse) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

func (m *GetOAuthUserResponse) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func (m *GetOAuthUserResponse) GetAuthDomain() string {
	if m != nil && m.AuthDomain != nil {
		return *m.AuthDomain
	}
	return ""
}

func (m *GetOAuthUserResponse) GetUserOrganization() string {
	if m != nil && m.UserOrganization != nil {
		return *m.UserOrganization
	}
	return ""
}

func (m *GetOAuthUserResponse) GetIsAdmin() bool {
	if m != nil && m.IsAdmin != nil {
		return *m.IsAdmin
	}
	return Default_GetOAuthUserResponse_IsAdmin
}

func (m *GetOAuthUserResponse) GetClientId() string {
	if m != nil && m.ClientId != nil {
		return *m.ClientId
	}
	return ""
}

func (m *GetOAuthUserResponse) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

type CheckOAuthSignatureRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CheckOAuthSignatureRequest) Reset()         { *m = CheckOAuthSignatureRequest{} }
func (m *CheckOAuthSignatureRequest) String() string { return proto.CompactTextString(m) }
func (*CheckOAuthSignatureRequest) ProtoMessage()    {}

type CheckOAuthSignatureResponse struct {
	OauthConsumerKey *string `protobuf:"bytes,1,req,name=oauth_consumer_key" json:"oauth_consumer_key,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CheckOAuthSignatureResponse) Reset()         { *m = CheckOAuthSignatureResponse{} }
func (m *CheckOAuthSignatureResponse) String() string { return proto.CompactTextString(m) }
func (*CheckOAuthSignatureResponse) ProtoMessage()    {}

func (m *CheckOAuthSignatureResponse) GetOauthConsumerKey() string {
	if m != nil && m.OauthConsumerKey != nil {
		return *m.OauthConsumerKey
	}
	return ""
}

func init() {
}
