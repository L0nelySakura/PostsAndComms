package graph


import (
	"github.com/LonelySakura/surely/graph/model"
)



func (r *queryResolver) commentTree(commentid string) []*model.Comment {
	var comments []*model.Comment
	if comment, exists := r.storage.comments[commentid]; exists {
		comments = append(comments, comment)
		for _, childID := range r.storage.childrenMap[commentid] {
			comments = append(comments, r.commentTree(childID)...)
		}
	}
	return comments
}
