package login

import (
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"strconv"
	"github.com/adriaandejonge/xld/util/intf"
)

func Do(args intf.Command) (result string, err error) {

	subs := args.Subs()
	if len(subs) != 3 {
		err = errors.New("xld login expects 3 arguments")
	} else {
		cred := subs[1] + ":" + subs[2]
		sEnc := b64.StdEncoding.EncodeToString([]byte(cred))
		result, err = Login(&LoginObject{subs[0], sEnc})
	}
	return
}

type LoginObject struct {
	Url  string
	Auth string
}

func Login(data *LoginObject) (result string, err error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://"+data.Url+"/deployit/server/info", nil)
	if err != nil {
		return
	}

	req.Header.Add("Authorization", "Basic "+data.Auth)

	resp, err := client.Do(req)
	if err != nil {
		return "error", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		err = createCache(data)
		return "Logged in Successfully", err
	} else {
		return "error", removeData(resp.StatusCode)
	}
}

func Check() (loginData *LoginObject, err error) {
	fileData, err := ioutil.ReadFile(cacheFile())
	if err != nil {
		return nil, errors.New("User is not logged in")
	}
	loginObject := LoginObject{}
	err = json.Unmarshal(fileData, &loginObject)
	return &loginObject, err
}

func cacheFile() string {
	currentUser, err := user.Current()
	if err != nil {
		// Cannot get current user's homedir, take /tmp instead
		return "/tmp"
	}
	return currentUser.HomeDir + "/.xld"
}

func createCache(data *LoginObject) (err error) {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ioutil.WriteFile(cacheFile(), []byte(js), 0777)
	return
}

func removeData(statusCode int) (err error) {
	os.Remove(cacheFile())
	err = errors.New("Login FAILED - Server returned status " + strconv.FormatInt(int64(statusCode), 10))
	return
}
