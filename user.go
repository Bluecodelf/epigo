package epigo

import (
	"encoding/json"
)

type UserInfo struct {
	Value            string `json:"value"`
	IsVisible        bool   `json:"public"`
	IsVisibleByAdmin bool   `json:"adm"`
}

type UserGroup struct {
	Name  string  `json:"name"`
	Title string  `json:"title"`
	Size  float64 `json:"count"`
}

type UserGPA struct {
	Value string `json:"gpa"`
	Cycle string `json:"cycle"`
}

// Average Grade Point Average. This is some next level shit.
type UserAverageGPA struct {
	Value string `json:"gpa_average"`
	Cycle string `json:"cycle"`
}

type UserNetsoul struct {
	HoursActive        float64 `json:"active"`
	HoursIdle          float64 `json:"idle"`
	HoursActiveOutside float64 `json:"out_active"`
	HoursIdleOutside   float64 `json:"out_idle"`
	LogNorm            float64 `json:"nslog_norm"`
}

type User struct {
	Login              string              `json:"login"`
	FullName           string              `json:"title"`
	Email              string              `json:"internal_email"`
	LastName           string              `json:"lastname"`
	FirstName          string              `json:"firstname"`
	UserInfo           map[string]UserInfo `json:"userinfo"`
	AvatarURL          string              `json:"picture"`
	Promotion          float64             `json:"promo"`
	Semester           float64             `json:"semester"`
	UID                float64             `json:"uid"`
	GID                float64             `json:"gid"`
	Location           string              `json:"location"`
	Closed             bool                `json:"close"`
	PromoID            string              `json:"id_promo"`
	HistoryID          string              `json:"id_history"`
	CourseCode         string              `json:"course_code"`
	SemesterCode       string              `json:"semester_code"`
	SchoolID           string              `json:"school_id"`
	SchoolCode         string              `json:"school_code"`
	SchoolTitle        string              `json:"school_title"`
	PreviousPromoID    string              `json:"old_id_promo"`
	PreviousLocationId string              `json:"old_id_location"`
	StudentYear        float64             `json:"studentyear"`
	IsAdmin            bool                `json:"admin"`
	IsEditable         bool                `json:"editable"`
	Groups             []UserGroup         `json:"groups"`
	Credits            float64             `json:"credits"`
	GPA                []UserGPA           `json:"gpa"`
	AverageGPA         []UserAverageGPA    `json:"average_gpa"`
	Spices             float64             `json:"spice"`
	Netsoul            UserNetsoul         `json:"nsstat"`
}

func (client *Client) GetUser(login string) (*User, error) {
	url := client.Host + "/user/" + login + "/?format=json"
	data, err := client.GetData(url)
	if err != nil {
		return nil, err
	}
	user := new(User)
	err = json.Unmarshal(data, user)
	return user, err
}
