// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"tree-table/internal/dao/internal"
)

// internalAreaDao is internal type for wrapping internal DAO implements.
type internalAreaDao = *internal.AreaDao

// areaDao is the data access object for table area.
// You can define custom methods on it to extend its functionality as you wish.
type areaDao struct {
	internalAreaDao
}

var (
	// Area is globally public accessible object for table area operations.
	Area = areaDao{
		internal.NewAreaDao(),
	}
)

// Fill with you ideas below.
