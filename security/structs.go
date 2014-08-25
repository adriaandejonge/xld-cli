package security

type (
	PermissionsList struct {
		Permissions []Permission `xml:"permission"`
	}

	Permission struct {
		Level         string        `xml:"level,attr"`
		Name      		string          `xml:"permissionName,attr"`
		Root         string        `xml:"root,attr"`
	}
)

