package structures

type SignUpLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	ImageRef string `json:"image_ref"`
}

type UpdateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Image    string `json:"image"`
	ImageRef string `json:"image_ref"`
}
