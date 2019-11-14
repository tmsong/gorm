package gorm

// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in your models
//    type User struct {
//      gorm.Model
//    }
type Model struct {
	ID         int64 `json:"id"          orm:"column(id)"`
	CreateTime int64 `json:"createTime"  orm:"column(create_time)"`
	UpdateTime int64 `json:"updateTime"  orm:"column(update_time)"`
}
