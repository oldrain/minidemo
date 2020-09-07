package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	contentTypeK = "Content-Type"
	contentTypeJsonV = "application/json; charset:UTF-8;"
	contentTypeUrlV = "application/x-www-form-urlencoded; charset:UTF-8;"
)

func HttpPostJson(url string, headers interface{}, body interface{}) (string, error) {
	reqHeaders := ObjToMap(headers)
	if headers == nil {
		reqHeaders = make(map[string]interface{})
	}
	reqHeaders[contentTypeK] = contentTypeJsonV
	res, err := HttpPost(url, reqHeaders, nil, ObjToJson(body))
	var ret string
	if err != nil {
		return ret, err
	}
	if res.StatusCode != http.StatusOK {
		return ret, fmt.Errorf("statusCode: %s", ObjToJson(res))
	}
	resBody, _ := ioutil.ReadAll(res.Body)
	return string(resBody), nil
}

func HttpPostUrlencoded(url string, headers interface{}, params interface{}) (string, error) {
	reqHeaders := ObjToMap(headers)
	if headers == nil {
		reqHeaders = make(map[string]interface{})
	}
	reqHeaders[contentTypeK] = contentTypeUrlV
	reqParams := ObjToMap(params)
	res, err := HttpPost(url, reqHeaders, reqParams, "")
	var ret string
	if err != nil {
		return ret, err
	}
	if res.StatusCode != http.StatusOK {
		return ret, fmt.Errorf("statusCode: %s", ObjToJson(res))
	}
	resBody, _ := ioutil.ReadAll(res.Body)
	return string(resBody), nil
}

func HttpPostFormData(url string, headers interface{}, params interface{}) (string, error) {
	reqHeaders := ObjToMap(headers)
	reqParams := ObjToMap(params)
	res, err := HttpPost(url, reqHeaders, reqParams, "")
	var ret string
	if err != nil {
		return ret, err
	}
	if res.StatusCode != http.StatusOK {
		return ret, fmt.Errorf("statusCode: %s", ObjToJson(res))
	}
	resBody, _ := ioutil.ReadAll(res.Body)
	return string(resBody), nil
}

func HttpGetUrlencoded(url string, headers interface{}, params interface{}) (string, error) {
	reqHeaders := ObjToMap(headers)
	if headers == nil {
		reqHeaders = make(map[string]interface{})
	}
	reqHeaders[contentTypeK] = contentTypeUrlV
	reqParams := ObjToMap(params)
	res, err := HttpGet(url, reqHeaders, reqParams, "")
	var ret string
	if err != nil {
		return ret, err
	}
	if res.StatusCode != http.StatusOK {
		return ret, fmt.Errorf("statusCode: %s", ObjToJson(res))
	}
	resBody, _ := ioutil.ReadAll(res.Body)
	return string(resBody), nil
}

func HttpGetUrlFormData(url string, headers interface{}, params interface{}) (string, error) {
	reqHeaders := ObjToMap(headers)
	reqParams := ObjToMap(params)
	res, err := HttpGet(url, reqHeaders, reqParams, "")
	var ret string
	if err != nil {
		return ret, err
	}
	if res.StatusCode != http.StatusOK {
		return ret, fmt.Errorf("statusCode: %s", ObjToJson(res))
	}
	resBody, _ := ioutil.ReadAll(res.Body)
	return string(resBody), nil
}

func HttpGetUrl(url string) (string, error) {
	res, err := HttpGet(url, nil, nil, "")
	var ret string
	if err != nil {
		return ret, err
	}
	if res.StatusCode != http.StatusOK {
		return ret, fmt.Errorf("statusCode: %s", ObjToJson(res))
	}
	resBody, _ := ioutil.ReadAll(res.Body)
	return string(resBody), nil
}

func HttpPost(url string, headers map[string]interface{}, params map[string]interface{}, body string) (*http.Response, error) {
	return HttpDo(url, http.MethodPost, headers, params, body)
}

func HttpGet(url string, headers map[string]interface{}, params map[string]interface{}, body string) (*http.Response, error) {
	return HttpDo(url, http.MethodGet, headers, params, body)
}

func HttpDo(url string, httpMethod string, headers map[string]interface{}, params map[string]interface{}, body string) (*http.Response, error) {
	req, err := http.NewRequest(httpMethod, url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, err
	}
	if len(params) > 0 {
		query := req.URL.Query()
		for k, v := range params {
			if v == nil {
				continue
			}
			query.Add(k, v.(string))
		}
		req.URL.RawQuery = query.Encode()
	}
	if len(headers) > 0 {
		for k, v := range headers {
			if v == nil {
				continue
			}
			req.Header.Add(k, v.(string))
		}
	}
	client := &http.Client{}
	return client.Do(req)
}
