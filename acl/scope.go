package acl

// PermissionScope is an "enum" for the different filesystem permissions
// supported.
type PermissionScope string

const (
	PermissionScopeDownload  PermissionScope = "download"
	PermissionScopeUpload                    = "upload"
	PermissionScopeRename                    = "rename"
	PermissionScopeRenameOwn                 = "renameown"
	PermissionScopeDelete                    = "delete"
	PermissionScopeDeleteOwn                 = "deleteown"
	PermissionScopeResume                    = "resume"
	PermissionScopeResumeOwn                 = "resumeown"
	PermissionScopeMakeDir                   = "makedir"
	PermissionScopeHideUser                  = "hide_user"
	PermissionScopeHideGroup                 = "hide_group"
	PermissionScopePrivate                   = "private"
)

var StringToPermissionScope = map[string]PermissionScope{
	string(PermissionScopeDownload):  PermissionScopeDownload,
	string(PermissionScopeUpload):    PermissionScopeUpload,
	string(PermissionScopeRename):    PermissionScopeRename,
	string(PermissionScopeRenameOwn): PermissionScopeRenameOwn,
	string(PermissionScopeDelete):    PermissionScopeDelete,
	string(PermissionScopeDeleteOwn): PermissionScopeDeleteOwn,
	string(PermissionScopeResume):    PermissionScopeResume,
	string(PermissionScopeResumeOwn): PermissionScopeResumeOwn,
	string(PermissionScopeMakeDir):   PermissionScopeMakeDir,
	string(PermissionScopeHideUser):  PermissionScopeHideUser,
	string(PermissionScopeHideGroup): PermissionScopeHideGroup,
	string(PermissionScopePrivate):   PermissionScopePrivate,
}
