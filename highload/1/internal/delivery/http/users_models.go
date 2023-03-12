package http

import "errors"

type UserRegisterReq struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Age        int32  `json:"age"`
	Biography  string `json:"biography"`
	City       string `json:"city"`
	Password   string `json:"password"`
}

type UserRegisterResp struct {
	UserID int64 `json:"user_id"`
}

type GetByIDReq struct {
	ID int64 `uri:"id"`
}

func (user *GetByIDReq) Validate() error {
	if user.ID <= 0 {
		return errors.New("id is not present")
	}
	return nil
}

type Login struct {
	ID       int64  `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResp struct {
	Token string `json:"token"`
}

type Error struct {
	Error string `json:"error"`
}
