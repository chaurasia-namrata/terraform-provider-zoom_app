package server

type Item struct {
	Id        string `json:"id"`
	EmailId   string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}


type User struct {
	Id        string `json:id`  
	EmailId   string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Type      int `json:"type"`
	RoleName  string `json:"role_name"`
	Pmi       int `json:"pmi"`
	UsePmi    bool `json:"use_pmi"`
	PersonalMeetingUrl  string `json:"personal_meeting_url"`
	TimeZone  string `json:"timezone"`
	Verified  int `json:"verified"`
	Dept      string `json:"dept"`
	CreatedAt string `json:"created_at"`
	LastLoginTime  string `json:"last_login_time"`
	PicUrl  string `json:"pic_url"`
	HostKey  string `json:"host_key"`
	CmsUserId  string `json:"cms_user_id"`
	Jid  string `json:"jid"`
	AccountId  string `json:"account_id"`
	Language  string `json:"language"`
	PhoneCountry  string `json:"phone_country"`
	PhoneNumber  string `json:"phone_number"`
	Status  string `json:"status"`
	JobTitle  string `json:"job_title"`
	Location  string `json:"location"`
	RoleId  string `json:"role_id"`

}

type NewUser struct {
	Action   string   `json:"action"`
	UserInfo UserInfo `json:"user_info"`
}
type UserInfo struct {
	EmailId   string `json:"email"`
	Type      int    `json:"type"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UpdateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
