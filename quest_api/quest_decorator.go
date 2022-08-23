package quest_api

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

const (
	GetApi = iota
	DeleteApi
	PutApi
	PostApi
)

type MyApi struct {
	sessions  string
	headers   map[string]interface{}
	urlPrefix string
	user      string
	password  string
}

func Decorator(z MyApi, funcName string, para map[string]interface{}) (any, error) {
	m, _ := reflect.TypeOf(z).MethodByName(funcName)
	v := make([]reflect.Value, 0)
	v = append(v, reflect.ValueOf(z))
	v = append(v, reflect.ValueOf(para))
	p := m.Func.Call(v)
	apiType := p[0].Int()
	url := p[1].String()
	header := p[2].Interface().(map[string]interface{})
	body := p[3].Interface().(map[string]interface{})

	switch apiType {
	case PutApi:
		return put(url, header, body)
	case DeleteApi:
		return delete(url, header, body)
	case PostApi:
		return post(url, header, body)
	case GetApi:
		return get(url, header, body)
	}
	return nil, errors.New("unknown api type")
}

func NewApi(para map[string]string) (*MyApi, error) {
	z := MyApi{}
	url := "http://" + para["ip"] + ":" + para["port"] + "/zstack/v1/"
	z.urlPrefix = url
	loginUrl := url + "accounts/login"
	passwordBytes := sha512.Sum512([]byte(para["password"]))
	password := hex.EncodeToString(passwordBytes[:])
	z.password = password
	user := para["admin"]
	z.user = user
	loginHeader := make(map[string]interface{}, 0)
	body := map[string]interface{}{"logInByAccount": map[string]string{"password": password, "accountName": user}}
	out, err := put(loginUrl, loginHeader, body)
	if err != nil {
		return nil, err
	}
	if k, OK := out["inventory"]; OK {
		z.sessions = k.(map[string]interface{})["uuid"].(string)
		z.headers = map[string]interface{}{"Authorization": "OAuth " + z.sessions}
		return &z, nil
	}
	if k, OK := out["error"]; OK {
		return nil, errors.New(k.(map[string]interface{})["details"].(string))
	}
	return nil, errors.New("unknown error")
}

func (Z MyApi) CreateZone(para map[string]interface{}) (int, string, map[string]interface{}, map[string]interface{}) {
	apiType := PostApi
	url := Z.urlPrefix + "zones"
	header := Z.headers
	body := map[string]interface{}{
		"params": map[string]interface{}{
			"name":        para["zone_name"],
			"description": para["description"],
		},
		"systemTags": nil,
		"userTags":   nil,
	}
	return apiType, url, header, body
}

func (Z MyApi) RevertVmFromSnapshotGroup(para map[string]interface{}) (int, string, map[string]interface{}, map[string]interface{}) {
	apiType := PutApi
	url := Z.urlPrefix + "volume-snapshots/group/" + para["uuid"].(string) + "/actions"
	header := Z.headers
	body := map[string]interface{}{
		"revertVmFromSnapshotGroup": map[string]interface{}{},
		"systemTags":                nil,
		"userTags":                  nil,
	}
	return apiType, url, header, body
}

func (Z MyApi) StopVmInstance(para map[string]interface{}) (int, string, map[string]interface{}, map[string]interface{}) {
	apiType := PutApi
	url := Z.urlPrefix + "vm-instances/" + para["uuid"].(string) + "/actions"
	header := Z.headers
	body := map[string]interface{}{
		"stopVmInstance": map[string]string{
			"type": "grace",
		},
		"systemTags": nil,
		"userTags":   nil,
	}
	return apiType, url, header, body
}

func (Z MyApi) QueryVmInstance(para map[string]interface{}) (int, string, map[string]interface{}, map[string]interface{}) {
	apiType := GetApi
	url := Z.urlPrefix + "vm-instances/" + para["uuid"].(string)
	header := Z.headers
	body := map[string]interface{}{}
	return apiType, url, header, body
}

func (Z MyApi) DetachTagFromResources(para map[string]interface{}) (int, string, map[string]interface{}, map[string]interface{}) {
	apiType := DeleteApi
	url := Z.urlPrefix + "tags/" + para["tagUuid"].(string) + "/resources?resourceUuids=" + para["resourceUuids"].(string)
	header := Z.headers
	body := map[string]interface{}{}
	return apiType, url, header, body
}

func (Z MyApi) AttachTagToResources(para map[string]interface{}) (int, string, map[string]interface{}, map[string]interface{}) {
	apiType := PostApi
	url := Z.urlPrefix + "tags/" + para["tagUuid"].(string) + "/resources"
	header := Z.headers
	body := map[string]interface{}{
		"params": map[string]interface{}{
			"resourceUuids": []string{
				para["uuid"].(string)},
		},
		"systemTags": nil,
		"userTags":   nil,
	}
	return apiType, url, header, body
}

func (Z MyApi) QueryTag(para map[string]interface{}) (int, string, map[string]interface{}, map[string]interface{}) {
	apiType := GetApi
	url := Z.urlPrefix + "tags/" + para["uuid"].(string)
	header := Z.headers
	body := map[string]interface{}{}
	return apiType, url, header, body
}

func (Z MyApi) QuerySystemTag(para map[string]interface{}) (int, string, map[string]interface{}, map[string]interface{}) {
	apiType := GetApi
	url := Z.urlPrefix + "system-tags/" + para["uuid"].(string)
	header := Z.headers
	body := map[string]interface{}{}
	return apiType, url, header, body
}

func (Z MyApi) ZqlVmTagList(para map[string]interface{}) (int, string, map[string]interface{}, map[string]interface{}) {
	apiType := GetApi
	zql := "query usertag where tagPatternUuid in (query accountResourceRef.resourceUuid where resourceType='TagPatternVO')"
	rawUrl := Z.urlPrefix + "zql"
	params := url.Values{}
	Url, err := url.Parse(rawUrl)
	if err != nil {
		panic(err)
	}
	params.Set("zql", zql)
	Url.RawQuery = params.Encode()
	header := Z.headers
	body := map[string]interface{}{}
	return apiType, Url.String(), header, body
}

func (Z MyApi) CreateVmInstance(para map[string]interface{}) (int, string, map[string]interface{}, map[string]interface{}) {
	apiType := PostApi
	url := Z.urlPrefix + "vm-instances"
	header := Z.headers
	body := map[string]interface{}{
		//	"params": {
		//		"name": name,
		//		"instanceOfferingUuid": instance_offering_uuid,
		//		"imageUuid": image_uuid,
		//		"l3NetworkUuids": [l3_net_uuid],
		//		"dataDiskOfferingUuids": datadisk_offering_uuid,
		//		'primaryStorageUuidForRootVolume': storage_uuid_for_root,
		//		"clusterUuid": cluster_uuid,
		//		"description": des,
		//		"strategy": strategy,
		//		"hostUuid": host_uuid
		//	},
		//	"systemTags": []string
		//	"staticIp::{l3_net_uuid}::{ip}".format(l3_net_uuid=l3_net_uuid, ip=ip),
		//],
		//	"userTags":   nil,
	}
	return apiType, url, header, body
}

func put(url string, header map[string]interface{}, body map[string]interface{}) (map[string]interface{}, error) {
	bo, _ := json.Marshal(body)
	b := strings.NewReader(string(bo))
	req, _ := http.NewRequest("PUT", url, b)
	for k, v := range header {
		req.Header.Set(k, v.(string))
	}
	v := &http.Client{}
	resp, err := v.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	p := make(map[string]interface{}, 0)
	err = json.Unmarshal(out, &p)
	return p, err
}
func get(url string, header map[string]interface{}, body map[string]interface{}) (any, error) {
	bo, _ := json.Marshal(body)
	b := strings.NewReader(string(bo))
	req, _ := http.NewRequest("GET", url, b)
	for k, v := range header {
		req.Header.Set(k, v.(string))
	}
	v := &http.Client{}
	resp, err := v.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var p any
	err = json.Unmarshal(out, &p)
	return p, err
}
func post(url string, header map[string]interface{}, body map[string]interface{}) (map[string]interface{}, error) {
	bo, _ := json.Marshal(body)
	b := strings.NewReader(string(bo))
	req, _ := http.NewRequest("POST", url, b)
	for k, v := range header {
		req.Header.Set(k, v.(string))
	}
	v := &http.Client{}
	resp, err := v.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	p := make(map[string]interface{}, 0)
	err = json.Unmarshal(out, &p)
	return p, err
}
func delete(url string, header map[string]interface{}, body map[string]interface{}) (map[string]interface{}, error) {
	bo, _ := json.Marshal(body)
	b := strings.NewReader(string(bo))
	req, _ := http.NewRequest("DELETE", url, b)
	for k, v := range header {
		req.Header.Set(k, v.(string))
	}
	v := &http.Client{}
	resp, err := v.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	p := make(map[string]interface{}, 0)
	err = json.Unmarshal(out, &p)
	return p, err
}
