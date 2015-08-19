package epigo

import (
	"encoding/json"
)

type UserInfo struct {
	value            string `json:"value"`
	isVisible        bool   `json:"public"`
	isVisibleByAdmin bool   `json:"adm"`
}

type UserGroup struct {
	name  string  `json:"name"`
	title string  `json:"title"`
	size  float64 `json:"count"`
}

type UserGPA struct {
	value string `json:"gpa"`
	cycle string `json:"cycle"`
}

// Average Grade Point Average. This is some next level shit.
type UserAverageGPA struct {
	value string `json:"gpa_average"`
	cycle string `json:"cycle"`
}

type UserNetsoul struct {
	hoursActive        float64 `json:"active"`
	hoursIdle          float64 `json:"idle"`
	hoursActiveOutside float64 `json:"out_active"`
	hoursIdleOutside   float64 `json:"out_idle"`
	logNorm            float64 `json:"nslog_norm"`
}

type User struct {
	login              string              `json:"login"`
	fullName           string              `json:"title"`
	email              string              `json:"internal_email"`
	lastName           string              `json:"lastname"`
	firstName          string              `json:"firstname"`
	userInfo           map[string]UserInfo `json:"userinfo"`
	avatarURL          string              `json:"picture"`
	promotion          float64             `json:"promo"`
	semester           float64             `json:"semester"`
	uid                float64             `json:"uid"`
	gid                float64             `json:"gid"`
	location           string              `json:"location"`
	closed             bool                `json:"close"`
	promoID            string              `json:"id_promo"`
	historyID          string              `json:"id_history"`
	courseCode         string              `json:"course_code"`
	semesterCode       string              `json:"semester_code"`
	schoolID           string              `json:"school_id"`
	schoolCode         string              `json:"school_code"`
	schoolTitle        string              `json:"school_title"`
	previousPromoID    string              `json:"old_id_promo"`
	previousLocationId string              `json:"old_id_location"`
	studentYear        string              `json:"studentyear"`
	isAdmin            bool                `json:"admin"`
	isEditable         bool                `json:"editable"`
	groups             []UserGroup         `json:"groups"`
	credits            float64             `json:"credits"`
	gpa                []UserGPA           `json:"gpa"`
	averageGPA         []UserAverageGPA    `json:"average_gpa"`
	spices             float64             `json:"spice"`
	netsoul            UserNetsoul         `json:"nsstat"`
}

func (client *Client) GetUser(login string) (*User, error) {
	url := client.host + "/user/" + login + "/?format=json"
	data, err := client.GetData(url)
	if err != nil {
		return nil, err
	}
	user := new(User)
	err = json.Unmarshal(data, user)
	return user, err
}
