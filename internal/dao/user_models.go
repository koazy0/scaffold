// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"scaffold/internal/dao/internal"
)

// userModelsDao is the data access object for the table user_models.
// You can define custom methods on it to extend its functionality as needed.
type userModelsDao struct {
	*internal.UserModelsDao
}

var (
	// UserModels is a globally accessible object for table user_models operations.
	UserModels = userModelsDao{internal.NewUserModelsDao()}
)

// Add your custom methods and functionality below.
