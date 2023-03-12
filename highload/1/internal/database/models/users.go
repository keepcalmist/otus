package models

type UserRegisterReq struct {
	FirstName  string
	SecondName string
	Age        int32
	Biography  string
	City       string
	Password   []byte
}

type UserRegisterResp struct {
	UserID int64
}

type GetUserFilter struct {
	UserIDs []int64
}

type User struct {
	ID         int64  `db:"id"`
	FirstName  string `db:"first_name"`
	SecondName string `db:"second_name"`
	Age        int32  `db:"age"`
	Biography  string `db:"biography"`
	City       string `db:"city"`
	Password   []byte `db:"password"`
}
