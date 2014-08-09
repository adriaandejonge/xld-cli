package repo

type (
	List struct {
		CIs []Ci `xml:"ci"`
	}

	Ci struct {
		CiRef string `xml:"ref,attr"`
		Type  string `xml:"type,attr"`
	}
)

type ByVersion [][]string

func (versions ByVersion) Len() int {
	return len(versions)
}
