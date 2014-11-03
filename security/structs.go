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

	PermSet struct {
		RolePerms []RolePerm `xml:"rolePermissions"`
	}

	RolePerm struct {
		Role Role `xml:"role"`
		Permissions []string `xml:"permissions"`

	}

	Role struct {
		Id int `xml:"id,attr"`
		Name string `xml:"name,attr"`
	}

)

