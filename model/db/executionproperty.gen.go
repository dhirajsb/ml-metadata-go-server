// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package db

const TableNameExecutionProperty = "ExecutionProperty"

// ExecutionProperty mapped from table <ExecutionProperty>
type ExecutionProperty struct {
	ExecutionID      int64    `gorm:"column:execution_id;primaryKey" json:"-"`
	Name             string   `gorm:"column:name;primaryKey;index:idx_execution_property_string,priority:1;index:idx_execution_property_int,priority:1;index:idx_execution_property_double,priority:1" json:"-"`
	IsCustomProperty bool     `gorm:"column:is_custom_property;primaryKey;index:idx_execution_property_string,priority:2;index:idx_execution_property_int,priority:2;index:idx_execution_property_double,priority:2" json:"-"`
	IntValue         *int64   `gorm:"column:int_value;index:idx_execution_property_int,priority:3" json:"-"`
	DoubleValue      *float64 `gorm:"column:double_value;index:idx_execution_property_double,priority:3" json:"-"`
	StringValue      *string  `gorm:"column:string_value;index:idx_execution_property_string,priority:3" json:"-"`
	ByteValue        *[]byte  `gorm:"column:byte_value" json:"-"`
	ProtoValue       *[]byte  `gorm:"column:proto_value" json:"-"`
	BoolValue        *bool    `gorm:"column:bool_value" json:"-"`
}

// TableName ExecutionProperty's table name
func (*ExecutionProperty) TableName() string {
	return TableNameExecutionProperty
}
