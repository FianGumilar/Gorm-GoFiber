package models

type Tag struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type TagResponseWithPost struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Posts []Post `json:"posts" gorm:"many2many:post_tags";ForeignKey:"ID";joinForeignKey:"TagID";References:ID;joinReferences:"PostID"`
}

func (TagResponseWithPost) TableName() string {
	return "tags"
}
