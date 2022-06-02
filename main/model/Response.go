package model

const (
	CODE_USER_NIL  string = "UndesignatedUser"
	CODE_USER_ZERO string = "NotFoundUser"

	MESSAGE_USER_NIL  string = "userを指定してください。"
	MESSAGE_USER_ZERO string = "userが存在しません。"

	TIME_FORMAT string = "2006/01/02/15/04/05"

	FILE_PATH string = "/forBiz-cms-api/.devcontainer/.env"
)

type Response struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
}
