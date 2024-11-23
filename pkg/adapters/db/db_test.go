package sqlc

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"testing"
)

/**
*
* db_test
* <p>
* db_test file
*
* Copyright (c) 2024 All rights reserved.
*
* This source code is protected by copyright and may not be reproduced,
* distributed, modified, or used in any form without the express written
* permission of the copyright owner.
*
* @author christian
* @author dbacilio88@outlook.es
* @since 5/08/2024
*
 */

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"
)

var store Store

func TestMain(m *testing.M) {
	pool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		return
	}
	store = NewStore(pool)
	os.Exit(m.Run())
}
