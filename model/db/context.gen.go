// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package db

const TableNameContext = "Context"

// Context mapped from table <Context>
type Context struct {
	ID                       int64   `gorm:"column:id;primaryKey;autoIncrement:true" json:"-"`
	TypeID                   int64   `gorm:"column:type_id;not null;uniqueIndex:UniqueContextName,priority:1" json:"-"`
	Name                     string  `gorm:"column:name;not null;uniqueIndex:UniqueContextName,priority:2" json:"-"`
	ExternalID               *string `gorm:"column:external_id;uniqueIndex:idx_context_external_id,priority:1" json:"-"`
	CreateTimeSinceEpoch     int64   `gorm:"autoCreateTime:milli;column:create_time_since_epoch;not null;index:idx_context_create_time_since_epoch,priority:1" json:"-"`
	LastUpdateTimeSinceEpoch int64   `gorm:"autoUpdateTime:milli;column:last_update_time_since_epoch;not null;index:idx_context_last_update_time_since_epoch,priority:1" json:"-"`

	// relationships
	Properties   []ContextProperty
	ContextType  Type `gorm:"foreignKey:TypeID;references:ID"`
	Attributions []Attribution `gorm:"foreignKey:ContextID;references:ID"`
	Associations []Association `gorm:"foreignKey:ContextID;references:ID"`

	// self-reference for context graphs
	Parents      []ParentContext   `gorm:"foreignKey:ContextID;references:ID"`
	Children     []ParentContext   `gorm:"foreignKey:ParentContextID;references:ID"`
}

// TableName Context's table name
func (*Context) TableName() string {
	return TableNameContext
}
