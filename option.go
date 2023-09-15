package sf

type Sfacg struct {
	cookie      string
	deviceToken string
	apiKey      string
	apiBaseUrl  string
}
type SfacgOptions interface {
	apply(*Sfacg)
}
type optionFunc func(*Sfacg)

func (f optionFunc) apply(s *Sfacg) {
	f(s)
}
func Cookie(cookie string) SfacgOptions {
	return optionFunc(func(s *Sfacg) {
		s.cookie = cookie
	})
}
func DeviceToken(deviceToken string) SfacgOptions {
	return optionFunc(func(s *Sfacg) {
		s.deviceToken = deviceToken
	})
}
func ApiKey(apiKey string) SfacgOptions {
	return optionFunc(func(s *Sfacg) {
		s.apiKey = apiKey
	})
}
func ApiBaseUrl(apiBaseUrl string) SfacgOptions {
	return optionFunc(func(s *Sfacg) {
		s.apiBaseUrl = apiBaseUrl
	})
}
