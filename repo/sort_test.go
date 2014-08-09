package repo

/*
import (
	"testing"
	"sort"
	"fmt"
)

func main() {
	toSort := [][]string{
		[]string{"1", "2", "3"},
		[]string{"1", "1", "2"},
		[]string{"1", "1", "3"},
		[]string{"1", "1", "2"},
		[]string{"2", "1", "3"},
		[]string{"1", "1", "2"},
		[]string{"1", "5", "3"},
		[]string{"1", "1", "2"},
		[]string{"1", "1", "3"},
		[]string{"1", "5", "2"},
		[]string{"5", "1", "3"},
		[]string{"4", "1", "2"},
		[]string{"3", "1", "3"},
		[]string{"2", "1", "2"},
		[]string{"1", "1", "3"},
		[]string{"1", "1", "2"},
	}
	sort.Sort(ByVersion(toSort))
	fmt.Println("Versions: ", toSort)

}

type ByVersion [][]string

func (versions ByVersion) Len() int {
    return len(versions)
}

func less(versions ByVersion, index, i, j int) bool {
	if versions[i][index] < versions[j][index] {
		return true
	} else if versions[i][index] < versions[j][index] {
		return false
	} else if index < len(versions[i]) - 2 {
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
} */
