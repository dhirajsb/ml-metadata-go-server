// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package db

const TableNameParentType = "ParentType"

// ParentType mapped from table <ParentType>
type ParentType struct {
	TypeID       int64 `gorm:"column:type_id;primaryKey" json:"-"`
	ParentTypeID int64 `gorm:"column:parent_type_id;primaryKey" json:"-"`
}

// TableName ParentType's table name
func (*ParentType) TableName() string {
	return TableNameParentType
}
