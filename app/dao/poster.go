// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"cosmonaut_api/app/dao/internal"
)

// posterDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type posterDao struct {
	*internal.PosterDao
}

var (
	// Poster is globally public accessible object for table poster operations.
	Poster posterDao
)

func init() {
	Poster = posterDao{
		internal.NewPosterDao(),
	}
}

// Fill with you ideas below.
