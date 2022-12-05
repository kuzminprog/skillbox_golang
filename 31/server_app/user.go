package user_app

type RequestCreate struct {
	Name    string   `json:"name"`
	Age     string   `json:"age"`
	Friends []string `json:"friends"`
}

type RequestMakeFriend struct {
	SourceID string `json:"source_id"`
	TargetID string `json:"target_id"`
}

type RequestDeleteUser struct {
	TargetID string `json:"target_id"`
}

type RequestAge struct {
	NewAge string `json:"new age"`
}
