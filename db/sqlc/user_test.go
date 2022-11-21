package db

import (
	"context"
	"testing"
	"time"

	"github.com/JILSE7/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) (User, error) {
	hashedPassword, err := utils.HashPassword(utils.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       utils.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       utils.RandomOwner(),
		Email:          utils.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChanged.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user, err
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)

}

func TestGetUser(t *testing.T) {
	// create account
	user, err := createRandomUser(t)

	// Get user
	searchuser, err := testQueries.GetUser(context.Background(), user.Username)

	require.NotEmpty(t, searchuser)
	require.NoError(t, err)

	require.Equal(t, user.Username, searchuser.Username)
	require.Equal(t, user.FullName, searchuser.FullName)
	require.Equal(t, user.Email, searchuser.Email)
	require.Equal(t, user.HashedPassword, searchuser.HashedPassword)
	require.Equal(t, user.CreatedAt, searchuser.CreatedAt)
	require.NotZero(t, searchuser.CreatedAt)

	require.WithinDuration(t, user.CreatedAt, searchuser.CreatedAt, time.Second)
	require.WithinDuration(t, user.PasswordChanged, searchuser.PasswordChanged, time.Second)

}
