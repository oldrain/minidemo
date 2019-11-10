package output

type UserInfoOut struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	DailyView int `json:"dailyView"`
}
