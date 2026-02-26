package base

type ValidateRequest struct {
	UserID      string
	Title       string
	Description string
}

func Validate(req *ValidateRequest) []string {
	var res []string
	if req == nil {
		return []string{"request is nil"}
	}

	if req.UserID == "" {
		res = append(res, "user id is empty")
	}
	if req.Title == "" {
		res = append(res, "title is empty")
	}
	if req.Description == "" {
		res = append(res, "description is empty")
	}

	return res
}