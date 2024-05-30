package entities

const TableNameBuku = "Buku"

type Buku struct {
	ID          *int    `gorm:"column:id" json:"id"`
	Name        *string `gorm:"column:name" json:"name"`
	PenulisID   *int    `gorm:"column:penulisId" json:"penulisId"`
	PenulisName *string `gorm:"column:penulisName" json:"penulisName"`
}

type BukuInsert struct {
	ID        *int    `gorm:"column:id" json:"id"`
	Name      *string `gorm:"column:name" json:"name"`
	PenulisID *int    `gorm:"column:penulisId" json:"penulisId"`
}

type BukuInsertMany struct {
	Data *[]BukuInsert
}

func (*Buku) TableName() string {
	return TableNameBuku
}

func (*BukuInsert) TableName() string {
	return TableNameBuku
}
