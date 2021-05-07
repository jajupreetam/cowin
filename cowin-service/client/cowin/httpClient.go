package cowin

import (
	"context"
	"cowin-service/utils"
	"crypto/tls"
	"errors"
	"fmt"
	slog "github.com/go-eden/slf4go"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

var httpClient *http.Client
var xSunriseChannel = "EP"
var XSunriseAppId = "DPS"

func init() {
	proxyUrl := viper.GetString("sunrise.client.proxy.url")
	readTimeout := viper.GetDuration("sunrise.client.http.readtimeout")
	if proxyUrl != "" {
		u := url.URL{}
		proxyUrl, _ := u.Parse(proxyUrl)
		transport := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
		httpClient = &http.Client{Timeout: readTimeout * time.Second, Transport: transport}
	} else {
		transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
		httpClient = &http.Client{Timeout: readTimeout * time.Second, Transport: transport}
	}
}

type header struct {
	key, value string
}

type HttpRequest struct {
	httpMethod  string
	baseUrl     string
	url         string
	body        io.Reader
	headers     []header
	queryParams url.Values
	basicAuth   basicAuth
}

type basicAuth struct {
	username string
	password string
}

func (httpRequest *HttpRequest) Get() *HttpRequest {
	httpRequest.httpMethod = "GET"
	return httpRequest
}

func (httpRequest *HttpRequest) Post(body io.Reader) *HttpRequest {
	httpRequest.httpMethod = "POST"
	return httpRequest.Body(body)
}

func (httpRequest *HttpRequest) Put(body io.Reader) *HttpRequest {
	httpRequest.httpMethod = "PUT"
	return httpRequest.Body(body)
}

func (httpRequest *HttpRequest) Delete() *HttpRequest {
	httpRequest.httpMethod = "DELETE"
	return httpRequest
}

func (httpRequest *HttpRequest) Cowin(ctx context.Context) *HttpRequest {
	var cowinUrl = viper.GetString("sunrise.client.cowin.url")
	initClient()
	addCowinRequestHeaders(ctx, httpRequest)
	httpRequest.baseUrl = cowinUrl
	return httpRequest
}

func (httpRequest *HttpRequest) Url(url string, pathParams ...interface{}) *HttpRequest {
	httpRequest.url = fmt.Sprintf(url, pathParams...)
	return httpRequest
}

func (httpRequest *HttpRequest) Body(body io.Reader) *HttpRequest {
	httpRequest.body = body
	return httpRequest
}

func (httpRequest *HttpRequest) Header(key string, value string) *HttpRequest {
	httpRequest.headers = append(httpRequest.headers, header{
		key:   key,
		value: value,
	})
	return httpRequest
}

func (httpRequest *HttpRequest) HeaderSalesId(value string) *HttpRequest {
	httpRequest.Header(utils.XSunriseSalesId, value)
	return httpRequest
}

func (httpRequest *HttpRequest) HeaderAuthentication(value string) *HttpRequest {
	httpRequest.Header(utils.XSunriseAuthentication, value)
	return httpRequest
}

func (httpRequest *HttpRequest) ContentType(value string) *HttpRequest {
	httpRequest.Header(utils.HeaderContentType, value)
	return httpRequest
}

func (httpRequest *HttpRequest) ContentTypeJson() *HttpRequest {
	httpRequest.Header(utils.HeaderContentType, utils.MIMEApplicationJSONCharsetUTF8)
	return httpRequest
}

func (httpRequest *HttpRequest) ContentTypeApplicationForm() *HttpRequest {
	httpRequest.Header(utils.HeaderContentType, utils.MIMEApplicationForm)
	return httpRequest
}

func (httpRequest *HttpRequest) AcceptJson() *HttpRequest {
	httpRequest.Header(utils.HeaderAccept, utils.MIMEApplicationJSON)
	return httpRequest
}

func (httpRequest *HttpRequest) QueryParams(queryParams url.Values) *HttpRequest {
	httpRequest.queryParams = queryParams
	return httpRequest
}

func (httpRequest *HttpRequest) BasicAuth(username string, password string) *HttpRequest {
	httpRequest.basicAuth.username = username
	httpRequest.basicAuth.password = password
	return httpRequest
}

func DoRequest(httpRequest *HttpRequest) (*http.Response, error) {
	var uri, _ = url.Parse(httpRequest.baseUrl + httpRequest.url)
	if httpRequest.queryParams != nil {
		uri.RawQuery = httpRequest.queryParams.Encode()
	}
	req, _ := http.NewRequest(httpRequest.httpMethod, uri.String(), httpRequest.body)
	for _, header := range httpRequest.headers {
		req.Header.Add(header.key, header.value)
	}

	if httpRequest.basicAuth.username != "" && httpRequest.basicAuth.password != "" {
		req.SetBasicAuth(httpRequest.basicAuth.username, httpRequest.basicAuth.password)
	}

	requestDump, _ := httputil.DumpRequest(req, true)
	slog.Debug("Request Out: ", string(requestDump))

	response, err := httpClient.Do(req)

	responseDump, _ := httputil.DumpResponse(response, true)
	slog.Debug("Response is: ", string(responseDump))

	if is4xxOr5xxError(response.Status) {
		fmt.Println("Error while doing External Server call")
		return nil, errors.New("http: 4xxOr5xxError")
	}

	return response, err
}

func DoRequestAndUnmarshal(httpRequest *HttpRequest, obj interface{}) error {
	httpRequest.AcceptJson()
	response, err := DoRequest(httpRequest)
	if err != nil {
		panic("Error while doing External Server call")
	}

	responseBodyBytes, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	err = ffjson.Unmarshal(responseBodyBytes, obj)
	if err != nil {
		slog.Error("Error while Unmarshalling:", err)
	}

	return err
}

func is4xxOr5xxError(responseStatus string) bool {
	return strings.HasPrefix(responseStatus, "4") || strings.HasPrefix(responseStatus, "5")
}

func addCowinRequestHeaders(ctx context.Context, httpRequest *HttpRequest) {
	httpRequest.
		Header("Authorization", "Bearer "+Token)
}

func initClient() {
	var readTimeout time.Duration = 30

	transport := &http.Transport{Proxy: nil, TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	httpClient = &http.Client{Timeout: readTimeout * time.Second, Transport: transport}
}
