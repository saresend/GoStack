package GoStack

import "net/http"

//GetAccessToken provides easy access to tokens from
func GetAccessToken(clientID string, clientSecret string) (string, error) {
	req, err := http.NewRequest("GET", "https://stackexchange.com/oauth", nil)
	println(req)
	if err != nil {
		return "", err
	}
	return "asdfasdf", nil

}
