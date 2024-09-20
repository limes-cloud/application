// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: api/application/errors/application_error_reason.proto

package errors

import (
	_ "github.com/go-kratos/kratos/v2/errors"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorReason int32

const (
	ErrorReason_ParamsError               ErrorReason = 0
	ErrorReason_DatabaseError             ErrorReason = 1
	ErrorReason_TransformError            ErrorReason = 2
	ErrorReason_GetError                  ErrorReason = 3
	ErrorReason_ListError                 ErrorReason = 4
	ErrorReason_CreateError               ErrorReason = 5
	ErrorReason_ImportError               ErrorReason = 6
	ErrorReason_ExportError               ErrorReason = 7
	ErrorReason_UpdateError               ErrorReason = 8
	ErrorReason_DeleteError               ErrorReason = 9
	ErrorReason_GetTrashError             ErrorReason = 10
	ErrorReason_ListTrashError            ErrorReason = 11
	ErrorReason_DeleteTrashError          ErrorReason = 12
	ErrorReason_RevertTrashError          ErrorReason = 13
	ErrorReason_ResourceServerError       ErrorReason = 14
	ErrorReason_ForbiddenError            ErrorReason = 15
	ErrorReason_SystemError               ErrorReason = 16
	ErrorReason_GenCaptchaError           ErrorReason = 17
	ErrorReason_NotExistEmailError        ErrorReason = 18
	ErrorReason_GenCaptchaTypeError       ErrorReason = 19
	ErrorReason_VerifyCaptchaError        ErrorReason = 20
	ErrorReason_OAuthLoginError           ErrorReason = 21
	ErrorReason_NotUserError              ErrorReason = 22
	ErrorReason_NotAppScopeError          ErrorReason = 23
	ErrorReason_RsaDecodeError            ErrorReason = 24
	ErrorReason_PasswordFormatError       ErrorReason = 25
	ErrorReason_PasswordExpireError       ErrorReason = 26
	ErrorReason_PasswordError             ErrorReason = 27
	ErrorReason_UserDisableError          ErrorReason = 28
	ErrorReason_GenTokenError             ErrorReason = 29
	ErrorReason_ParseTokenError           ErrorReason = 30
	ErrorReason_RefreshTokenError         ErrorReason = 31
	ErrorReason_DisableRegisterError      ErrorReason = 32
	ErrorReason_AlreadyExistEmailError    ErrorReason = 33
	ErrorReason_AlreadyExistUsernameError ErrorReason = 34
	ErrorReason_RegisterError             ErrorReason = 35
	ErrorReason_BindError                 ErrorReason = 36
	ErrorReason_LoginError                ErrorReason = 37
	ErrorReason_ExistFeedbackError        ErrorReason = 38
	ErrorReason_ManagerServerError        ErrorReason = 39
	ErrorReason_NotPermissionError        ErrorReason = 40
	ErrorReason_AlreadyBindError          ErrorReason = 41
	ErrorReason_AppMaintenanceError       ErrorReason = 42
	ErrorReason_ChannelCloseError         ErrorReason = 43
	ErrorReason_AppNotBindChannelError    ErrorReason = 44
	ErrorReason_ChannelNotBindUserError   ErrorReason = 45
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0:  "ParamsError",
		1:  "DatabaseError",
		2:  "TransformError",
		3:  "GetError",
		4:  "ListError",
		5:  "CreateError",
		6:  "ImportError",
		7:  "ExportError",
		8:  "UpdateError",
		9:  "DeleteError",
		10: "GetTrashError",
		11: "ListTrashError",
		12: "DeleteTrashError",
		13: "RevertTrashError",
		14: "ResourceServerError",
		15: "ForbiddenError",
		16: "SystemError",
		17: "GenCaptchaError",
		18: "NotExistEmailError",
		19: "GenCaptchaTypeError",
		20: "VerifyCaptchaError",
		21: "OAuthLoginError",
		22: "NotUserError",
		23: "NotAppScopeError",
		24: "RsaDecodeError",
		25: "PasswordFormatError",
		26: "PasswordExpireError",
		27: "PasswordError",
		28: "UserDisableError",
		29: "GenTokenError",
		30: "ParseTokenError",
		31: "RefreshTokenError",
		32: "DisableRegisterError",
		33: "AlreadyExistEmailError",
		34: "AlreadyExistUsernameError",
		35: "RegisterError",
		36: "BindError",
		37: "LoginError",
		38: "ExistFeedbackError",
		39: "ManagerServerError",
		40: "NotPermissionError",
		41: "AlreadyBindError",
		42: "AppMaintenanceError",
		43: "ChannelCloseError",
		44: "AppNotBindChannelError",
		45: "ChannelNotBindUserError",
	}
	ErrorReason_value = map[string]int32{
		"ParamsError":               0,
		"DatabaseError":             1,
		"TransformError":            2,
		"GetError":                  3,
		"ListError":                 4,
		"CreateError":               5,
		"ImportError":               6,
		"ExportError":               7,
		"UpdateError":               8,
		"DeleteError":               9,
		"GetTrashError":             10,
		"ListTrashError":            11,
		"DeleteTrashError":          12,
		"RevertTrashError":          13,
		"ResourceServerError":       14,
		"ForbiddenError":            15,
		"SystemError":               16,
		"GenCaptchaError":           17,
		"NotExistEmailError":        18,
		"GenCaptchaTypeError":       19,
		"VerifyCaptchaError":        20,
		"OAuthLoginError":           21,
		"NotUserError":              22,
		"NotAppScopeError":          23,
		"RsaDecodeError":            24,
		"PasswordFormatError":       25,
		"PasswordExpireError":       26,
		"PasswordError":             27,
		"UserDisableError":          28,
		"GenTokenError":             29,
		"ParseTokenError":           30,
		"RefreshTokenError":         31,
		"DisableRegisterError":      32,
		"AlreadyExistEmailError":    33,
		"AlreadyExistUsernameError": 34,
		"RegisterError":             35,
		"BindError":                 36,
		"LoginError":                37,
		"ExistFeedbackError":        38,
		"ManagerServerError":        39,
		"NotPermissionError":        40,
		"AlreadyBindError":          41,
		"AppMaintenanceError":       42,
		"ChannelCloseError":         43,
		"AppNotBindChannelError":    44,
		"ChannelNotBindUserError":   45,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_api_application_errors_application_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_api_application_errors_application_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_api_application_errors_application_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_api_application_errors_application_error_reason_proto protoreflect.FileDescriptor

var file_api_application_errors_application_error_reason_proto_rawDesc = []byte{
	0x0a, 0x35, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a,
	0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x8e, 0x11, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x10, 0x00, 0x1a, 0x0f, 0xb2, 0x45, 0x0c, 0xe5, 0x8f, 0x82, 0xe6, 0x95, 0xb0,
	0xe9, 0x94, 0x99, 0xe8, 0xaf, 0xaf, 0x12, 0x25, 0x0a, 0x0d, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x01, 0x1a, 0x12, 0xb2, 0x45, 0x0f, 0xe6, 0x95,
	0xb0, 0xe6, 0x8d, 0xae, 0xe5, 0xba, 0x93, 0xe9, 0x94, 0x99, 0xe8, 0xaf, 0xaf, 0x12, 0x29, 0x0a,
	0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10,
	0x02, 0x1a, 0x15, 0xb2, 0x45, 0x12, 0xe6, 0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xe8, 0xbd, 0xac, 0xe6,
	0x8d, 0xa2, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x23, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x10, 0x03, 0x1a, 0x15, 0xb2, 0x45, 0x12, 0xe8, 0x8e, 0xb7, 0xe5, 0x8f,
	0x96, 0xe6, 0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x2a, 0x0a,
	0x09, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x04, 0x1a, 0x1b, 0xb2, 0x45,
	0x18, 0xe8, 0x8e, 0xb7, 0xe5, 0x8f, 0x96, 0xe5, 0x88, 0x97, 0xe8, 0xa1, 0xa8, 0xe6, 0x95, 0xb0,
	0xe6, 0x8d, 0xae, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x26, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x05, 0x1a, 0x15, 0xb2, 0x45, 0x12, 0xe5,
	0x88, 0x9b, 0xe5, 0xbb, 0xba, 0xe6, 0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4,
	0xa5, 0x12, 0x26, 0x0a, 0x0b, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x10, 0x06, 0x1a, 0x15, 0xb2, 0x45, 0x12, 0xe5, 0xaf, 0xbc, 0xe5, 0x85, 0xa5, 0xe6, 0x95, 0xb0,
	0xe6, 0x8d, 0xae, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x26, 0x0a, 0x0b, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x07, 0x1a, 0x15, 0xb2, 0x45, 0x12, 0xe5,
	0xaf, 0xbc, 0xe5, 0x87, 0xba, 0xe6, 0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4,
	0xa5, 0x12, 0x26, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x10, 0x08, 0x1a, 0x15, 0xb2, 0x45, 0x12, 0xe6, 0x9b, 0xb4, 0xe6, 0x96, 0xb0, 0xe6, 0x95, 0xb0,
	0xe6, 0x8d, 0xae, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x26, 0x0a, 0x0b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x09, 0x1a, 0x15, 0xb2, 0x45, 0x12, 0xe5,
	0x88, 0xa0, 0xe9, 0x99, 0xa4, 0xe6, 0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4,
	0xa5, 0x12, 0x31, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x73, 0x68, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x0a, 0x1a, 0x1e, 0xb2, 0x45, 0x1b, 0xe8, 0x8e, 0xb7, 0xe5, 0x8f, 0x96, 0xe5,
	0x9b, 0x9e, 0xe6, 0x94, 0xb6, 0xe7, 0xab, 0x99, 0xe6, 0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xe5, 0xa4,
	0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x38, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x73,
	0x68, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x0b, 0x1a, 0x24, 0xb2, 0x45, 0x21, 0xe8, 0x8e, 0xb7,
	0xe5, 0x8f, 0x96, 0xe5, 0x9b, 0x9e, 0xe6, 0x94, 0xb6, 0xe7, 0xab, 0x99, 0xe5, 0x88, 0x97, 0xe8,
	0xa1, 0xa8, 0xe6, 0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x34,
	0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x73, 0x68, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x0c, 0x1a, 0x1e, 0xb2, 0x45, 0x1b, 0xe5, 0x88, 0xa0, 0xe9, 0x99, 0xa4, 0xe5,
	0x9b, 0x9e, 0xe6, 0x94, 0xb6, 0xe7, 0xab, 0x99, 0xe6, 0x95, 0xb0, 0xe6, 0x8d, 0xae, 0xe5, 0xa4,
	0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x34, 0x0a, 0x10, 0x52, 0x65, 0x76, 0x65, 0x72, 0x74, 0x54, 0x72,
	0x61, 0x73, 0x68, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x0d, 0x1a, 0x1e, 0xb2, 0x45, 0x1b, 0xe8,
	0xbf, 0x98, 0xe5, 0x8e, 0x9f, 0xe5, 0x9b, 0x9e, 0xe6, 0x94, 0xb6, 0xe7, 0xab, 0x99, 0xe6, 0x95,
	0xb0, 0xe6, 0x8d, 0xae, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x2e, 0x0a, 0x13, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x10, 0x0e, 0x1a, 0x15, 0xb2, 0x45, 0x12, 0xe8, 0xb5, 0x84, 0xe6, 0xba, 0x90, 0xe6, 0x9c,
	0x8d, 0xe5, 0x8a, 0xa1, 0xe5, 0xbc, 0x82, 0xe5, 0xb8, 0xb8, 0x12, 0x2a, 0x0a, 0x0e, 0x46, 0x6f,
	0x72, 0x62, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x0f, 0x1a, 0x16,
	0xa8, 0x45, 0x93, 0x03, 0xb2, 0x45, 0x0f, 0xe6, 0x97, 0xa0, 0xe5, 0xba, 0x94, 0xe7, 0x94, 0xa8,
	0xe6, 0x9d, 0x83, 0xe9, 0x99, 0x90, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x10, 0x1a, 0x0f, 0xb2, 0x45, 0x0c, 0xe7, 0xb3, 0xbb, 0xe7,
	0xbb, 0x9f, 0xe9, 0x94, 0x99, 0xe8, 0xaf, 0xaf, 0x12, 0x2d, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x43,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x11, 0x1a, 0x18, 0xb2,
	0x45, 0x15, 0xe9, 0xaa, 0x8c, 0xe8, 0xaf, 0x81, 0xe7, 0xa0, 0x81, 0xe7, 0x94, 0x9f, 0xe6, 0x88,
	0x90, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x2d, 0x0a, 0x12, 0x4e, 0x6f, 0x74, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x12, 0x1a,
	0x15, 0xb2, 0x45, 0x12, 0xe4, 0xb8, 0x8d, 0xe5, 0xad, 0x98, 0xe5, 0x9c, 0xa8, 0xe6, 0xad, 0xa4,
	0xe9, 0x82, 0xae, 0xe7, 0xae, 0xb1, 0x12, 0x34, 0x0a, 0x13, 0x47, 0x65, 0x6e, 0x43, 0x61, 0x70,
	0x74, 0x63, 0x68, 0x61, 0x54, 0x79, 0x70, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x13, 0x1a,
	0x1b, 0xb2, 0x45, 0x18, 0xe9, 0x94, 0x99, 0xe8, 0xaf, 0xaf, 0xe7, 0x9a, 0x84, 0xe9, 0xaa, 0x8c,
	0xe8, 0xaf, 0x81, 0xe7, 0xa0, 0x81, 0xe7, 0xb1, 0xbb, 0xe5, 0x9e, 0x8b, 0x12, 0x30, 0x0a, 0x12,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x14, 0x1a, 0x18, 0xb2, 0x45, 0x15, 0xe9, 0xaa, 0x8c, 0xe8, 0xaf, 0x81, 0xe7,
	0xa0, 0x81, 0xe9, 0xaa, 0x8c, 0xe8, 0xaf, 0x81, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x30,
	0x0a, 0x0f, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x10, 0x15, 0x1a, 0x1b, 0xb2, 0x45, 0x18, 0xe4, 0xb8, 0x89, 0xe6, 0x96, 0xb9, 0xe6, 0x8e,
	0x88, 0xe6, 0x9d, 0x83, 0xe7, 0x99, 0xbb, 0xe9, 0x99, 0x86, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5,
	0x12, 0x24, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x55, 0x73, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x10, 0x16, 0x1a, 0x12, 0xb2, 0x45, 0x0f, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe4, 0xb8, 0x8d,
	0xe5, 0xad, 0x98, 0xe5, 0x9c, 0xa8, 0x12, 0x2e, 0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x41, 0x70, 0x70,
	0x53, 0x63, 0x6f, 0x70, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x17, 0x1a, 0x18, 0xb2, 0x45,
	0x15, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe6, 0x97, 0xa0, 0xe5, 0xba, 0x94, 0xe7, 0x94, 0xa8,
	0xe6, 0x9d, 0x83, 0xe9, 0x99, 0x90, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x73, 0x61, 0x44, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x18, 0x1a, 0x12, 0xb2, 0x45, 0x0f, 0x72,
	0x73, 0x61, 0xe8, 0xa7, 0xa3, 0xe5, 0xaf, 0x86, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x2e,
	0x0a, 0x13, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x19, 0x1a, 0x15, 0xb2, 0x45, 0x12, 0xe5, 0xaf, 0x86, 0xe7,
	0xa0, 0x81, 0xe6, 0xa0, 0xbc, 0xe5, 0xbc, 0x8f, 0xe9, 0x94, 0x99, 0xe8, 0xaf, 0xaf, 0x12, 0x2b,
	0x0a, 0x13, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x1a, 0x1a, 0x12, 0xb2, 0x45, 0x0f, 0xe5, 0xaf, 0x86, 0xe7,
	0xa0, 0x81, 0xe5, 0xb7, 0xb2, 0xe8, 0xbf, 0x87, 0xe6, 0x9c, 0x9f, 0x12, 0x2b, 0x0a, 0x0d, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x1b, 0x1a, 0x18,
	0xb2, 0x45, 0x15, 0xe8, 0xb4, 0xa6, 0xe6, 0x88, 0xb7, 0xe6, 0x88, 0x96, 0xe5, 0xaf, 0x86, 0xe7,
	0xa0, 0x81, 0xe9, 0x94, 0x99, 0xe8, 0xaf, 0xaf, 0x12, 0x2b, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x1c, 0x1a, 0x15,
	0xb2, 0x45, 0x12, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe5, 0xb7, 0xb2, 0xe8, 0xa2, 0xab, 0xe7,
	0xa6, 0x81, 0xe7, 0x94, 0xa8, 0x12, 0x27, 0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x1d, 0x1a, 0x14, 0xb2, 0x45, 0x11, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0xe7, 0x94, 0x9f, 0xe6, 0x88, 0x90, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x29,
	0x0a, 0x0f, 0x50, 0x61, 0x72, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x10, 0x1e, 0x1a, 0x14, 0xb2, 0x45, 0x11, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0xe8, 0xa7, 0xa3,
	0xe6, 0x9e, 0x90, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x2f, 0x0a, 0x11, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x1f,
	0x1a, 0x18, 0xa8, 0x45, 0x91, 0x03, 0xb2, 0x45, 0x11, 0xe5, 0x88, 0xb7, 0xe6, 0x96, 0xb0, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x3c, 0x0a, 0x14, 0x44, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x20, 0x1a, 0x22, 0xa8, 0x45, 0x91, 0x03, 0xb2, 0x45, 0x1b, 0xe5, 0xba, 0x94,
	0xe7, 0x94, 0xa8, 0xe5, 0xb7, 0xb2, 0xe5, 0x85, 0xb3, 0xe9, 0x97, 0xad, 0xe6, 0xb3, 0xa8, 0xe5,
	0x86, 0x8c, 0xe6, 0x9d, 0x83, 0xe9, 0x99, 0x90, 0x12, 0x2e, 0x0a, 0x16, 0x41, 0x6c, 0x72, 0x65,
	0x61, 0x64, 0x79, 0x45, 0x78, 0x69, 0x73, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x21, 0x1a, 0x12, 0xb2, 0x45, 0x0f, 0xe9, 0x82, 0xae, 0xe7, 0xae, 0xb1, 0xe5,
	0xb7, 0xb2, 0xe5, 0xad, 0x98, 0xe5, 0x9c, 0xa8, 0x12, 0x31, 0x0a, 0x19, 0x41, 0x6c, 0x72, 0x65,
	0x61, 0x64, 0x79, 0x45, 0x78, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x22, 0x1a, 0x12, 0xb2, 0x45, 0x0f, 0xe8, 0xb4, 0xa6, 0xe5,
	0x8f, 0xb7, 0xe5, 0xb7, 0xb2, 0xe5, 0xad, 0x98, 0xe5, 0x9c, 0xa8, 0x12, 0x28, 0x0a, 0x0d, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x23, 0x1a, 0x15,
	0xb2, 0x45, 0x12, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe6, 0xb3, 0xa8, 0xe5, 0x86, 0x8c, 0xe5,
	0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x24, 0x0a, 0x09, 0x42, 0x69, 0x6e, 0x64, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x24, 0x1a, 0x15, 0xb2, 0x45, 0x12, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe6,
	0xb3, 0xa8, 0xe5, 0x86, 0x8c, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x1f, 0x0a, 0x0a, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x25, 0x1a, 0x0f, 0xb2, 0x45, 0x0c,
	0xe7, 0x99, 0xbb, 0xe9, 0x99, 0x86, 0xe5, 0xa4, 0xb1, 0xe8, 0xb4, 0xa5, 0x12, 0x39, 0x0a, 0x12,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x26, 0x1a, 0x21, 0xb2, 0x45, 0x1e, 0xe5, 0xb7, 0xb2, 0xe5, 0xad, 0x98, 0xe5,
	0x9c, 0xa8, 0xe9, 0x87, 0x8d, 0xe5, 0xa4, 0x8d, 0xe7, 0x9a, 0x84, 0xe5, 0x8f, 0x8d, 0xe9, 0xa6,
	0x88, 0xe5, 0x86, 0x85, 0xe5, 0xae, 0xb9, 0x12, 0x33, 0x0a, 0x12, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x27, 0x1a,
	0x1b, 0xb2, 0x45, 0x18, 0xe7, 0xae, 0xa1, 0xe7, 0x90, 0x86, 0xe4, 0xb8, 0xad, 0xe5, 0xbf, 0x83,
	0xe6, 0x9c, 0x8d, 0xe5, 0x8a, 0xa1, 0xe5, 0xbc, 0x82, 0xe5, 0xb8, 0xb8, 0x12, 0x2a, 0x0a, 0x12,
	0x4e, 0x6f, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x10, 0x28, 0x1a, 0x12, 0xb2, 0x45, 0x0f, 0xe6, 0x97, 0xa0, 0xe8, 0xb5, 0x84, 0xe6,
	0xba, 0x90, 0xe6, 0x9d, 0x83, 0xe9, 0x99, 0x90, 0x12, 0x4f, 0x0a, 0x10, 0x41, 0x6c, 0x72, 0x65,
	0x61, 0x64, 0x79, 0x42, 0x69, 0x6e, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x29, 0x1a, 0x39,
	0xb2, 0x45, 0x36, 0xe8, 0xaf, 0xa5, 0xe8, 0xb4, 0xa6, 0xe5, 0x8f, 0xb7, 0xe5, 0xb7, 0xb2, 0xe7,
	0xbb, 0x91, 0xe5, 0xae, 0x9a, 0xe8, 0xbf, 0x87, 0xe5, 0x85, 0xb6, 0xe4, 0xbb, 0x96, 0xe5, 0xb9,
	0xb3, 0xe5, 0x8f, 0xb0, 0xef, 0xbc, 0x8c, 0xe4, 0xb8, 0x8d, 0xe8, 0x83, 0xbd, 0xe9, 0x87, 0x8d,
	0xe5, 0xa4, 0x8d, 0xe7, 0xbb, 0x91, 0xe5, 0xae, 0x9a, 0x12, 0x37, 0x0a, 0x13, 0x41, 0x70, 0x70,
	0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x10, 0x2a, 0x1a, 0x1e, 0xb2, 0x45, 0x1b, 0xe5, 0xbd, 0x93, 0xe5, 0x89, 0x8d, 0xe5, 0xba, 0x94,
	0xe7, 0x94, 0xa8, 0xe6, 0xad, 0xa3, 0xe5, 0x9c, 0xa8, 0xe7, 0xbb, 0xb4, 0xe6, 0x8a, 0xa4, 0xe4,
	0xb8, 0xad, 0x12, 0x2f, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x2b, 0x1a, 0x18, 0xb2, 0x45, 0x15, 0xe5, 0xbd,
	0x93, 0xe5, 0x89, 0x8d, 0xe6, 0xb8, 0xa0, 0xe9, 0x81, 0x93, 0xe5, 0xb7, 0xb2, 0xe5, 0x85, 0xb3,
	0xe9, 0x97, 0xad, 0x12, 0x37, 0x0a, 0x16, 0x41, 0x70, 0x70, 0x4e, 0x6f, 0x74, 0x42, 0x69, 0x6e,
	0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x2c, 0x1a,
	0x1b, 0xb2, 0x45, 0x18, 0xe5, 0xba, 0x94, 0xe7, 0x94, 0xa8, 0xe6, 0x9c, 0xaa, 0xe5, 0xbc, 0x80,
	0xe9, 0x80, 0x9a, 0xe6, 0xad, 0xa4, 0xe6, 0xb8, 0xa0, 0xe9, 0x81, 0x93, 0x12, 0x41, 0x0a, 0x17,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x6f, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x2d, 0x1a, 0x24, 0xb2, 0x45, 0x21, 0xe5, 0xbd,
	0x93, 0xe5, 0x89, 0x8d, 0xe6, 0x8e, 0x88, 0xe6, 0x9d, 0x83, 0xe6, 0xb8, 0xa0, 0xe9, 0x81, 0x93,
	0xe6, 0x9c, 0xaa, 0xe7, 0xbb, 0x91, 0xe5, 0xae, 0x9a, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0x1a,
	0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_application_errors_application_error_reason_proto_rawDescOnce sync.Once
	file_api_application_errors_application_error_reason_proto_rawDescData = file_api_application_errors_application_error_reason_proto_rawDesc
)

func file_api_application_errors_application_error_reason_proto_rawDescGZIP() []byte {
	file_api_application_errors_application_error_reason_proto_rawDescOnce.Do(func() {
		file_api_application_errors_application_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_application_errors_application_error_reason_proto_rawDescData)
	})
	return file_api_application_errors_application_error_reason_proto_rawDescData
}

var file_api_application_errors_application_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_application_errors_application_error_reason_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: errors.ErrorReason
}
var file_api_application_errors_application_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_application_errors_application_error_reason_proto_init() }
func file_api_application_errors_application_error_reason_proto_init() {
	if File_api_application_errors_application_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_application_errors_application_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_application_errors_application_error_reason_proto_goTypes,
		DependencyIndexes: file_api_application_errors_application_error_reason_proto_depIdxs,
		EnumInfos:         file_api_application_errors_application_error_reason_proto_enumTypes,
	}.Build()
	File_api_application_errors_application_error_reason_proto = out.File
	file_api_application_errors_application_error_reason_proto_rawDesc = nil
	file_api_application_errors_application_error_reason_proto_goTypes = nil
	file_api_application_errors_application_error_reason_proto_depIdxs = nil
}
