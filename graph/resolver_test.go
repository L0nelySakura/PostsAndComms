package graph

import (
	"context"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	r := &Resolver{
		storage: *NewStorage(),
	}

	post, err := r.Mutation().CreatePost(
		context.Background(),
		"Test Title",
		"Test Content",
		"Test Author",
		nil,
	)

	assert.NoError(t, err)
	assert.NotEmpty(t, post.ID)
	assert.Equal(t, "Test Title", post.Title)
	assert.True(t, post.CommentsEnabled)

	comment1, err := r.Mutation().CreateComment(
		context.Background(),
		"Test Comment1",
		"Test Author",
		post.ID,
		"",
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, comment1.ID)
	assert.Equal(t, "Test Comment1", comment1.Content)
	assert.Empty(t, comment1.Parentid)

	comment2, err := r.Mutation().CreateComment(
		context.Background(),
		"Test Comment2",
		"Test Author",
		post.ID,
		comment1.ID,
	)
	assert.NoError(t, err)
	assert.NotEmpty(t, comment2.ID)
	assert.Equal(t, comment1.ID, comment2.Parentid)
}