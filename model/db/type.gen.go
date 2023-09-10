// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package db

const TableNameType = "Type"

// Type mapped from table <Type>
type Type struct {
	ID          int64   `gorm:"column:id;primaryKey;autoIncrement:true" json:"-"`
	Name        string  `gorm:"column:name;not null;uniqueIndex:idx_type_name,priority:1" json:"-"`
	Version     *string `gorm:"column:version" json:"-"`
	TypeKind    int32   `gorm:"column:type_kind;not null" json:"-"`
	Description *string `gorm:"column:description" json:"-"`
	InputType   *string `gorm:"column:input_type" json:"-"`
	OutputType  *string `gorm:"column:output_type" json:"-"`
	ExternalID  *string `gorm:"column:external_id;uniqueIndex:idx_type_external_id,priority:1" json:"-"`

	// relationships
	Properties []TypeProperty
}

// TableName Type's table name
func (*Type) TableName() string {
	return TableNameType
}
