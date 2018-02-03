package models

import (
	"github.com/timkellogg/five_three_one/services/authentication"
	"github.com/timkellogg/five_three_one/services/database"
	"github.com/timkellogg/five_three_one/services/session"
)

var (
	// Database - persistant storage
	Database = database.NewDatabase().Store

	// Session - in memory storage
	Session = session.NewSession().Memcache.Client

	// Authentication - provides jwt auth service
	Authentication = &authentication.AuthService{}
)
