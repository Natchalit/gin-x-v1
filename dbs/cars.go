package dbs

import "github.com/Natchalit/gin-x-v1/pg"

var CARS = &pg.Connect{
	DBName: `dev_liyl`,
}
