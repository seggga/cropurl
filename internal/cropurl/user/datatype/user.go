package datatype

type User struct {
	UserID   *string `json:"user_id"`
	UserName *string `json:"user_name"`
	UserPass *string `json:"user_pass"`
	UserData *string `json:"user_data"`
}
