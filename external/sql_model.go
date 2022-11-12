package external

import (
	"gorm.io/gorm"
	"time"
)

type SQLModel struct {
	ID              uint32    `json:"id"         gorm:"primarykey"`
	FormatCreatedAt *DateTime `json:"created_at" gorm:"-"`
	FormatUpdatedAt *DateTime `json:"updated_at" gorm:"-"`

	// gorm.Model
	CreatedAt time.Time      `json:"-" gorm:"autoCreateTime:true;column:created_at"`
	UpdatedAt time.Time      `json:"-" gorm:"autoUpdateTime:true;column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}

type SQLHardDeleteModel struct {
	// gorm.Model
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime:true;column:created_at"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime:true;column:updated_at"`
}

func (m *SQLModel) GenUID(dbType int) {
	//uid := NewUID(uint32(m.Id), dbType, 1)
	//m.FakeId = &uid
	m.FormatCreatedAt = &DateTime{
		DateTime: &m.CreatedAt,
	}
	m.FormatUpdatedAt = &DateTime{
		DateTime: &m.UpdatedAt,
	}
}
