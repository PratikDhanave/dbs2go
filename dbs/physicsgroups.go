package dbs

import (
	"fmt"
)

// physicsgroups API
func (API) PhysicsGroups(params Record) []Record {
	// variables we'll use in where clause
	var args []interface{}
	where := "WHERE "

	// parse dataset argument
	physicsgroups := getValues(params, "physics_group_name")
	if len(physicsgroups) > 1 {
		msg := "The physicsgroups API does not support list of physicsgroups"
		return errorRecord(msg)
	} else if len(physicsgroups) == 1 {
		op, val := opVal(physicsgroups[0])
		cond := fmt.Sprintf(" pg.PHYSICS_GROUP_NAME %s %s", op, placeholder("physics_group_name"))
		where += addCond(where, cond)
		args = append(args, val)
	} else {
		where = "" // no arguments
	}
	// get SQL statement from static area
	stm := getSQL("physicsgroups")
	// use generic query API to fetch the results from DB
	return executeAll(stm+where, args...)
}
