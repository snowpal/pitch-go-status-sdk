package common

type Member struct {
	ID   string `json:"ID"`
	Role string `json:"Role"`
}

type MemberReqBody struct {
	ID   string `json:"id"`
	Role string `json:"role"`
}

type PairedMember struct {
	MemberId      *string  `json:"memberID"`
	NonMemberName string   `json:"nonMemberName"`
	Items         []string `json:"items"`
}

type PullRequest struct {
	Url         string `json:"url"`
	Description string `json:"description"`
}

type ProfileName struct {
	FirstName  string `json:"FirstName"`
	MiddleName string `json:"MiddleName"`
	LastName   string `json:"LastName"`
}

type PhoneNumber struct {
	Primary     bool   `json:"primary"`
	PhoneNumber string `json:"phoneNumber"`
}

type EmailAddress struct {
	Primary      bool   `json:"primary"`
	EmailAddress string `json:"emailAddress"`
}
