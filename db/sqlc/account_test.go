package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/jbiers/e-wallet/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	Username := util.RandomString(6)
	Type := util.RandomType()
	Document := util.RandomDocument(Type)
	Email := util.RandomEmail(Username)
	Password := "password"

	args := CreateAccountParams{
		Username,
		Type,
		Document,
		Email,
		Password,
	}

	account, err := testQueries.CreateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	var balanceValue int64 = 0

	require.Equal(t, args.Username, account.Username)
	require.Equal(t, args.Type, account.Type)
	require.Equal(t, args.Document, account.Document)
	require.Equal(t, args.Email, account.Email)
	require.Equal(t, args.Password, account.Password)
	require.Equal(t, balanceValue, account.Balance)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Username, account2.Username)
	require.Equal(t, account1.Email, account2.Email)
	require.Equal(t, account1.Type, account2.Type)
	require.Equal(t, account1.Document, account2.Document)
	require.Equal(t, account1.Password, account2.Password)
	require.Equal(t, account1.Balance, account2.Balance)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomInt(0, 10000),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Username, account2.Username)
	require.Equal(t, account1.Email, account2.Email)
	require.Equal(t, account1.Type, account2.Type)
	require.Equal(t, account1.Document, account2.Document)
	require.Equal(t, account1.Password, account2.Password)
	require.Equal(t, arg.Balance, account2.Balance)
}
