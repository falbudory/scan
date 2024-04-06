package structs

type User struct {
	FirstName              string `json:"first_name"`
	LastName               string `json:"last_name"`
	Email                  string `json:"email"`
	PhoneNumber            string `json:"phone_number"`
	Address                string `json:"address"`
	Username               string `json:"username"`
	Password               string `json:"password"`
	ReferralCode           string `json:"referral_code"`
	NameBusiness           string `json:"name_business"`
	FullNameRepresentative string `json:"full_name_representative"`
}

type AccUser struct {
	FirstName              string `json:"first_name"`
	LastName               string `json:"last_name"`
	Email                  string `json:"email"`
	PhoneNumber            string `json:"phone_number"`
	Address                string `json:"address"`
	Username               string `json:"username"`
	Password               string `json:"password"`
	ReferralCode           string `json:"referral_code"`
	NameBusiness           string `json:"name_business"`
	FullNameRepresentative string `json:"full_name_representative"`
	TypeUserID             int    `json:"type_user_id"`
	RoleID                 int    `json:"role_id"`
}

type ReqBody struct {
	Draw   int `json:"draw"`
	Start  int `json:"start"`
	Length int `json:"length"`
	Order  []struct {
		Column int    `json:"column"`
		Dir    string `json:"dir"`
	} `json:"order"`
	Search struct {
		Value string `json:"value"`
	} `json:"search"`
}

type RoleForm struct {
	Name         string `form:"name" validate:"required,min=3,max=30"`
	PermissionID []int  `form:"permission_id" validate:"required"`
}

type FormStateUser struct {
	ID    int  `json:"id"`
	State bool `json:"state"`
}
