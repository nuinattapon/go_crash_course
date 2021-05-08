package db

import (
	"context"
	"database/sql"
	"math/rand"
	"testing"

	"github.com/nuinattapon/go_crash_course/44_sqlc/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAuthor(t *testing.T) {
	arg := CreateAuthorParams{
		Name: util.RandomName(),
		Bio:  util.RandomBio(),
	}
	err := testQueries.CreateAuthor(context.Background(), arg)
	require.NoError(t, err)

	id, err := testQueries.LastInsertId(context.Background())
	require.NoError(t, err)
	require.NotZero(t, id)

	author, err := testQueries.GetAuthor(context.Background(), int32(id))
	require.NoError(t, err)
	require.Equal(t, arg.Name, author.Name)
	require.Equal(t, arg.Bio, author.Bio)
}

func TestGetAuthor(t *testing.T) {
	// id := int32(util.RandomInt(1, 11))

	idList, err := testQueries.ListAuthorID(context.Background(), 20)
	require.NoError(t, err)

	id := rand.Int31n(int32(len(idList)))

	author, err := testQueries.GetAuthor(context.Background(), id)

	require.NoError(t, err)
	require.NotEmpty(t, author.ID)
	require.NotZero(t, author.ID)
	require.NotEmpty(t, author.Name)
	require.NotEmpty(t, author.Bio)
}

func TestDeleteAuthor(t *testing.T) {
	arg := CreateAuthorParams{
		Name: util.RandomName(),
		Bio:  util.RandomBio(),
	}
	err := testQueries.CreateAuthor(context.Background(), arg)
	require.NoError(t, err)

	id, err := testQueries.LastInsertId(context.Background())
	require.NoError(t, err)
	require.NotZero(t, id)

	err = testQueries.DeleteAuthor(context.Background(), int32(id))
	require.NoError(t, err)

	author, err := testQueries.GetAuthor(context.Background(), int32(id))
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, author)
}

func TestUpdateAuthor(t *testing.T) {
	id := int32(util.RandomInt(1, 10))

	arg := UpdateAuthorParams{
		Name: util.RandomName(),
		Bio:  util.RandomBio(),
		ID:   id,
	}
	err := testQueries.UpdateAuthor(context.Background(), arg)
	require.NoError(t, err)

	rowsAffected, err := testQueries.RowsAffected(context.Background())
	require.NoError(t, err)
	require.EqualValues(t, rowsAffected, 1)

	author, err := testQueries.GetAuthor(context.Background(), int32(id))
	require.NoError(t, err)

	require.Equal(t, arg.Name, author.Name)
	require.Equal(t, arg.Bio, author.Bio)
	require.Equal(t, arg.ID, author.ID)
}

func TestLastInsertId(t *testing.T) {
	id, err := testQueries.LastInsertId(context.Background())
	require.NoError(t, err)
	require.NotZero(t, id)
}

func TestListAuthor(t *testing.T) {
	arg := ListAuthorsParams{
		Limit:  5,
		Offset: 5,
	}
	authors, err := testQueries.ListAuthors(context.Background(), arg)
	require.NoError(t, err)
	for _, author := range authors {
		require.NotZero(t, author.ID)
		require.NotEmpty(t, author.Name)
		require.NotEmpty(t, author.Bio)
	}

}
