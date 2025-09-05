package responses

type MemberResponse struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Email      string `json:"emil"`
	AvatarID   uint   `json:"avatar_id"`
	AvatarIcon string `json:"avatar_icon"`
}
