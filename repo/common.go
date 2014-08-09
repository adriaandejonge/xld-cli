package repo

import (
	"encoding/xml"
	"github.com/adriaandejonge/xld/util/http"
	"regexp"
	"sort"
	"strings"
)

var shorthand = map[string]string{
	"app":  "Applications",
	"env":  "Environments",
	"inf":  "Infrastructure",
	"conf": "Configuration",
}

// TODO Make this a util?
func AntiAbbreviate(ciName string) string {
	prefix := strings.SplitN(ciName, "/", 2)
	longer := shorthand[prefix[0]]

	remainder := ""

	if len(prefix) > 1 {
		remainder = prefix[1]

	}

	if longer != "" {
		ciName = longer + "/" + remainder
	}

	if strings.HasSuffix(ciName, "latest") {
		searchFor := ciName[:len(ciName)-len("latest")]
		body, err := http.Read("/repository/query?parent=" + searchFor)
		list, err := readCiList(body)
		if err != nil {
			// TODO ERROR
			return "error"
		}

		splitter := regexp.MustCompile("[\\.?\\-?_?]+")

		versions := make([][]string, 0)
		ciMap := make(map[string]Ci)
		for _, el := range list.CIs {
			version := el.CiRef[len(searchFor):]

			split := splitter.Split(version, -1)
			ciMap[strings.Join(split, ".")] = el
			versions = append(versions, split)
		}
		sort.Sort(ByVersion(versions))
		ciName = ciMap[strings.Join(versions[len(versions)-1], ".")].CiRef

	}

	return ciName
}

// Simple heuristic to avoid converting strings and integers
// Please mind: A goes before AA and B goes before AA.
// In other words, shorter strings go before longer.

// Alternative:
//matcher := regexp.MustCompile("^[0-9]+$")
//sort
//	matcher.MatchString(el) -> true / false
// -> decide string vs int compare

// TODO: if necessary, speed this up; now inefficient but not noticable

const norm string = "0000000000"

func normalizeLess(a string, b string) bool {
	return (norm + a)[len(a):] < (norm + b)[len(b):]
}

func normalizeMore(a string, b string) bool {
	return (norm + a)[len(a):] > (norm + b)[len(b):]
}

func less(versions ByVersion, index, i, j int) bool {
	if normalizeLess(versions[i][index], versions[j][index]) {
		return true
	} else if normalizeMore(versions[i][index], versions[j][index]) {
		return false
	} else if index < len(versions[i])-1 {
		return less(versions, index+1, i, j)
	} else {
		return false
	}
}

func (versions ByVersion) Less(i, j int) bool {
	return less(versions, 0, i, j)
}
func (versions ByVersion) Swap(i, j int) {
	versions[i], versions[j] = versions[j], versions[i]
}

func readCiList(body []byte) (list List, err error) {
	list = List{}
	err = xml.Unmarshal(body, &list)
	return
}
