package models

type Post struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  string `json:"user_id"`
	User    User   `json:"user";"gorm:foreignKey:UserId"`
}
