package utils

var (
	AppVerify              = Rules{"name_cn": {NotEmpty()}, "name_en": {NotEmpty()}, "description": {NotEmpty()}}
	LoginVerify            = Rules{"captchaId": {NotEmpty()}, "captcha": {NotEmpty()}, "username": {NotEmpty()}, "password": {NotEmpty()}}
	RegisterVerify         = Rules{"user_name": {NotEmpty()}, "real_name": {NotEmpty()}, "Password": {NotEmpty()}, "type": {NotEmpty()}}
	PageInfoVerify         = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	ModuleInfoVerify       = Rules{"app_id": {NotEmpty()}, "module_id": {NotEmpty()}}
	ModuleModifyVerify     = Rules{"definition": {NotEmpty()}, "module_id": {NotEmpty()}, "ui_schema": {NotEmpty()}}
	ModuleVerify           = Rules{"name_cn": {NotEmpty()}, "name_en": {NotEmpty()}, "app_id": {NotEmpty()}}
	ParamSelectVerify      = Rules{"appId": {Ge("0")}}
	CustomerVerify         = Rules{"CustomerName": {NotEmpty()}, "CustomerPhoneData": {NotEmpty()}}
	AutoPackageVerify      = Rules{"PackageName": {NotEmpty()}}
	AuthorityVerify        = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}}
	AuthorityIdVerify      = Rules{"AuthorityId": {NotEmpty()}}
	OldAuthorityVerify     = Rules{"OldAuthorityId": {NotEmpty()}}
	ChangePasswordVerify   = Rules{"Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	SetUserAuthorityVerify = Rules{"AuthorityId": {NotEmpty()}}
)
