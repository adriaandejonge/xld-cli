package metadata

type (
	List struct {
		CITypes []CIType `xml:"descriptor"`
	}

	CIType struct {
		Type         string        `xml:"type,attr"`
		Virtual      bool          `xml:"virtual,attr"`
		Versioned    bool          `xml:"versioned,attr"`
		Root         string        `xml:"root,attr"`
		Properties   []Property    `xml:"property-descriptors>property-descriptor"`
		ControlTasks []ControlTask `xml:"control-tasks>control-task"`
		Interfaces   []string      `xml:"interfaces>interface"`
		SuperTypes   []string      `xml:"superTypes>superType"`
		props        map[string]Property
	}

	Property struct {
		Name               string `xml:"name,attr"`
		Fqn                string `xml:"fqn,attr"`
		Label              string `xml:"label,attr"`
		Kind               string `xml:"kind,attr"`
		Description        string `xml:"description,attr"`
		Category           string `xml:"category,attr"`
		AsContainment      bool   `xml:"asContainment,attr"`
		Inspection         bool   `xml:"inspection,attr"`
		Required           bool   `xml:"required,attr"`
		RequiredInspection bool   `xml:"requiredInspection,attr"`
		Password           bool   `xml:"password,attr"`
		Transient          bool   `xml:"transient,attr"`
		Size               string `xml:"size,attr"`
		Default            string `xml:"default,attr"`
		Hidden             bool   `xml:"hidden,attr"`
		ReferencedType     string `xml:"referencedType"`
	}

	ControlTask struct {
		Name        string `xml:"name,attr"`
		Fqn         string `xml:"fqn,attr"`
		Label       string `xml:"label,attr"`
		Description string `xml:"description,attr"`
	}
)

func (ciType *CIType) Prop(name string) Property {
	if ciType.props == nil {
		ciType.indexProps()
	}
	return ciType.props[name]
}

func (ciType *CIType) indexProps() {
	ciType.props = make(map[string]Property)
	for _, prop := range ciType.Properties {
		ciType.props[prop.Name] = prop
	}
}
