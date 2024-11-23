package sqlc

import (
	"context"
	"github.com/dbacilio88/go/pkg/components/helpers"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

/**
*
* users_test
* <p>
* users_test file
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

func TestQueries_ListUsers(t *testing.T) {
	var user User
	for i := 0; i < 100; i++ {
		user = createUserRandom(t)
	}

	param := ListUsersParams{
		Role:   user.Role,
		Limit:  10,
		Offset: 0,
	}

	users, err := store.ListUsers(context.Background(), param)
	require.NoError(t, err)
	require.NotEmpty(t, users)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}

func TestQueries_DeleteUser(t *testing.T) {
	user := createUserRandom(t)
	err := store.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)
	_, err = store.GetUser(context.Background(), user.ID)
	require.Error(t, err)
}

func TestQueries_UpdateUser(t *testing.T) {
	user := createUserRandom(t)
	param := UpdateUserParams{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
	}
	db, err := store.UpdateUser(context.Background(), param)
	require.NoError(t, err)
	require.NotEmpty(t, db)
	require.NotZero(t, db.ID)
	require.NotZero(t, db.CreatedAt)
	require.Equal(t, user.ID, db.ID)
	require.Equal(t, user.Role, db.Role)
	require.Equal(t, user.Password, db.Password)
	require.Equal(t, user.Username, db.Username)
	require.Equal(t, user.CreatedAt, db.CreatedAt)
	require.WithinDuration(t, user.CreatedAt, db.CreatedAt, time.Second)
}

func TestQueries_GetUser(t *testing.T) {
	user := createUserRandom(t)
	db, err := store.GetUser(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotNil(t, db)
	require.Equal(t, user, db)
	require.NotZero(t, db.ID)
	require.NotZero(t, db.CreatedAt)
	require.Equal(t, user.ID, db.ID)
	require.Equal(t, user.Role, db.Role)
	require.Equal(t, user.Password, db.Password)
	require.Equal(t, user.Username, db.Username)
}

func TestQueries_CreateUser(t *testing.T) {
	createUserRandom(t)
}

func createUserRandom(t *testing.T) User {
	params := CreateUserParams{
		Username: helpers.RandomUser(),
		Password: helpers.RandomPassword(),
		Role:     helpers.RandomRole(),
	}
	user, err := store.CreateUser(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotZero(t, user.ID)
	require.Equal(t, params.Username, user.Username)
	require.Equal(t, params.Role, user.Role)
	require.Equal(t, params.Password, user.Password)
	return user
}
