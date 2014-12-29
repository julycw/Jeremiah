package data_package_struct

type ErrPackage struct {
	ErrCode    int    `json:"errcode"`
	ErrMessage string `json:"errmsg"`
}
