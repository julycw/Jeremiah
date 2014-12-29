package data_package_struct

//{"access_token":"ACCESS_TOKEN","expires_in":7200}
type AccessTokenPackage struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
