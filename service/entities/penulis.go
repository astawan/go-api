package entities

const TableNamePenulis = "Penulis"

type Penulis struct {
	ID   *int    `gorm:"column:id" json:"id"`
	Name *string `gorm:"column:name" json:"name"`
}

func (*Penulis) TableName() string {
	return TableNamePenulis
}
