package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/larrymacrich/gator/internal/database"
)

// handler_ListPosts takes optional posts [limit] (default: 2)
// and prints [limit] posts to the console
func handlerListPosts(s *state, cmd command, user database.User) error {
	if len(cmd.args) > 1 {
		errMsg := fmt.Errorf("usage: %s [limit]", cmd.name)
		return errMsg
	}
	limit := int32(2)
	if len(cmd.args) == 1 {
		limitStr := cmd.args[0]
		limit64, err := strconv.ParseInt(limitStr, 10, 32)
		if err != nil {
			errMsg := fmt.Errorf("usage: converting '%s' failed: %w", limitStr, err)
			return errMsg
		}
		limit = int32(limit64)
	}

	getPostsForUserByUserIDParams := database.GetPostsForUserByUserIDParams{
		UserID: user.ID,
		Limit:  limit,
	}
	posts, err := s.db.GetPostsForUserByUserID(context.Background(), getPostsForUserByUserIDParams)
	if err != nil {
		errMsg := fmt.Errorf("get posts for user '%s'  failed: %w", user.Name, err)
		return errMsg
	}

	fmt.Printf("found %d post(s) followed by '%s'\n", len(posts), user.Name)
	fmt.Println("=====================================")
	for _, post := range posts {
		printPost(&post)
	}

	return nil
}
