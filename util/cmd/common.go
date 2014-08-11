package cmd

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

func Stdin2Map() (result map[string]interface{}, err error) {

		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}

		result = make(map[string]interface{})

		err = json.Unmarshal(bytes, &result)
		
		return

}