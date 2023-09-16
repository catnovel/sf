package sf

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/catnovelapi/BuilderHttpClient"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
	"time"
)

func NewSfClient(options ...SfacgOptions) *Sfacg {
	sfacgClient := &Sfacg{
		deviceToken: uuid.New().String(),
		apiKey:      "FMLxgOdsfxmN!Dt4",
		apiBaseUrl:  "https://api.sfacg.com",
	}
	for _, option := range options {
		option.apply(sfacgClient)
	}
	return sfacgClient
}

func (s *Sfacg) NewCookie(cookie string) {
	Cookie(cookie).apply(s)
}

func (s *Sfacg) sfSecurity() string {
	var (
		uuid_, _   = uuid.NewUUID()
		StrUUID    = strings.ToUpper(uuid_.String())
		sTimeStamp = strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
		newMd5     = md5.New()
	)
	newMd5.Write([]byte(StrUUID + sTimeStamp + strings.ToUpper(s.deviceToken) + s.apiKey))
	sign := strings.ToUpper(hex.EncodeToString(newMd5.Sum(nil)))
	return fmt.Sprintf("nonce=%v&timestamp=%v&devicetoken=%v&sign=%v", StrUUID, sTimeStamp, strings.ToUpper(s.deviceToken), sign)
}

func (s *Sfacg) sfacgHeader() map[string]any {
	h := make(map[string]any)
	h["SFSecurity"] = s.sfSecurity()
	h["User-Agent"] = "boluobao/4.9.62(android;16.5)/OPPO/" + s.deviceToken
	h["authorization"] = "Basic YW5kcm9pZHVzZXI6MWEjJDUxLXl0Njk7KkFjdkBxeHE="
	if s.cookie != "" {
		h["Cookie"] = s.cookie
	}
	return h
}

func (s *Sfacg) get(endURL string, params any) gjson.Result {
	return BuilderHttpClient.Get(s.apiBaseUrl+endURL, BuilderHttpClient.Body(params), BuilderHttpClient.Header(s.sfacgHeader())).Gjson()
}

func (s *Sfacg) post(endURL string, params any) BuilderHttpClient.ResponseInterfaceBuilder {
	postHeader := s.sfacgHeader()
	postHeader["Content-Type"] = "application/json"
	return BuilderHttpClient.Post(s.apiBaseUrl+endURL, BuilderHttpClient.Body(params), BuilderHttpClient.Header(postHeader))
}

func (s *Sfacg) postWeb(endURL string, params any, notCookie bool) BuilderHttpClient.ResponseInterfaceBuilder {
	headers := s.sfacgHeader()
	headers["Content-Type"] = "application/json"
	delete(headers, "SFSecurity")
	if notCookie {
		delete(headers, "Cookie")
	}
	return BuilderHttpClient.Post(s.apiBaseUrl+endURL, BuilderHttpClient.Body(params), BuilderHttpClient.Header(headers))
}

func (s *Sfacg) getWeb(webURL string, params any, notCookie bool) gjson.Result {
	headers := s.sfacgHeader()
	delete(headers, "SFSecurity")
	if notCookie {
		delete(headers, "Cookie")
	}
	return BuilderHttpClient.Get(webURL, BuilderHttpClient.Body(params), BuilderHttpClient.Header(headers)).Gjson()
}
