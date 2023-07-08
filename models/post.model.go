package models

type Post struct {
	ID     int          `json:"id" gorm:"primaryKey"`
	Title  string       `json:"title" form:"title" gorm:"not null"`
	Body   string       `json:"body" form:"body" gorm:"not null"`
	UserId int          `json:"user_id" form:"user_id"`
	User   UserResponse `json:"user"`
	Tags   []Tag        `json:"tags" gorm:"many2many:post_tags"`
	TagID  []int        `json:"tag_id" form:"tags_id" gorm:"-"`
}

type PostResponse struct {
	Title  string `json:"title" form:"title"`
	Body   string `json::"body" form:"body"`
	UserId int    `json:"user_id" form:"user_id"`
}

type PostResponseWithTag struct {
	ID     int          `json:"id"`
	Title  string       `json:"title" form:"title"`
	Body   string       `json:"body" form:"body"`
	UserId int          `json:"user_id" form:"user_id"`
	User   UserResponse `json:"user"`
	Tags   []Tag        `json:"tags" gorm:"many2many:post_tags";ForeignKey:"ID";joinForeignKey:"PostID";References:ID;joinReferences:"TagID"`
}

func (PostResponse) TableName() string {
	return "posts"
}

func (PostResponseWithTag) TableName() string {
	return "posts"
}
