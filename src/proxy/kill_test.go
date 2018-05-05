/*
 * Radon
 *
 * Copyright 2018 The Radon Authors.
 * Code is licensed under the GPLv3.
 *
 */

package proxy

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xelabs/go-mysqlstack/driver"
	"github.com/xelabs/go-mysqlstack/sqlparser/depends/sqltypes"
	"github.com/xelabs/go-mysqlstack/xlog"
)

func TestProxyKill(t *testing.T) {
	log := xlog.NewStdLog(xlog.Level(xlog.PANIC))
	fakedbs, proxy, cleanup := MockProxy(log)
	defer cleanup()
	address := proxy.Address()
	iptable := proxy.IPTable()

	// fakedbs.
	{
		fakedbs.AddQueryPattern("use .*", &sqltypes.Result{})
		fakedbs.AddQueryPattern("create table .*", &sqltypes.Result{})
		fakedbs.AddQueryPattern("select * .*", &sqltypes.Result{})
		fakedbs.AddQueryDelay("select * from test.t1_0002", &sqltypes.Result{}, 100000000)
		fakedbs.AddQueryDelay("select * from test.t1_0004", &sqltypes.Result{}, 100000000)
	}

	// IPTables.
	{
		iptable.Add("127.0.0.1")
	}

	// create test table.
	{
		client, err := driver.NewConn("mock", "mock", address, "test", "utf8")
		assert.Nil(t, err)
		query := "create table t1(id int, b int) partition by hash(id)"
		_, err = client.FetchAll(query, -1)
		assert.Nil(t, err)
		client.Quit()
	}

	var wg sync.WaitGroup
	var clients []driver.Conn
	nums := 1
	// long query.
	{
		for i := 0; i < nums; i++ {
			client, err := driver.NewConn("mock", "mock", address, "test", "utf8")
			assert.Nil(t, err)
			wg.Add(1)
			go func(c driver.Conn) {
				defer wg.Done()
				query := "select * from t1"
				_, err = client.FetchAll(query, -1)
				log.Debug("%+v", err)
			}(client)
			clients = append(clients, client)
		}
	}

	// kill.
	{
		time.Sleep(time.Second * 1)
		for i := 0; i < nums; i++ {
			kill, err := driver.NewConn("mock", "mock", address, "test", "utf8")
			assert.Nil(t, err)
			wg.Add(1)
			go func(c driver.Conn, id uint32) {
				defer wg.Done()
				query := fmt.Sprintf("kill %d", id)
				_, err = kill.FetchAll(query, -1)
				assert.Nil(t, err)
			}(kill, clients[i].ConnectionID())
		}
	}
	wg.Wait()
}
