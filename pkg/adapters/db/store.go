package sqlc

import "github.com/jackc/pgx/v5/pgxpool"

/**
*
* store
* <p>
* store file
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

type Store interface {
	Querier
}

type SqlStore struct {
	connPool *pgxpool.Pool
	*Queries
}

func NewStore(connPool *pgxpool.Pool) *SqlStore {
	return &SqlStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
