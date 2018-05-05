/*
 * Radon
 *
 * Copyright 2018 The Radon Authors.
 * Code is licensed under the GPLv3.
 *
 */

package planner

import (
	"router"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xelabs/go-mysqlstack/sqlparser"
	"github.com/xelabs/go-mysqlstack/xlog"
)

func TestPlanner(t *testing.T) {
	log := xlog.NewStdLog(xlog.Level(xlog.PANIC))

	database := "xx"
	query := "create table A(a int)"

	route, cleanup := router.MockNewRouter(log)
	defer cleanup()

	err := route.AddForTest(database, router.MockTableAConfig())
	assert.Nil(t, err)

	node, err := sqlparser.Parse(query)
	assert.Nil(t, err)
	DDL := NewDDLPlan(log, database, query, node.(*sqlparser.DDL), route)

	{
		planTree := NewPlanTree()
		for i := 0; i < 64; i++ {
			err := planTree.Add(DDL)
			assert.Nil(t, err)
		}
		err := planTree.Build()
		assert.Nil(t, err)
	}
}
