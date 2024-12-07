package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Dhar01/Gator/commands"
	"github.com/Dhar01/Gator/internal/database"
)

func HandlerBrowse(s *commands.State, cmd commands.Command, user database.User) error {
	limit := 2

	if len(cmd.Args) == 1 {
		if specifiedLimit, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = specifiedLimit
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}

	posts, err := s.DB.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})

	if err != nil {
		return fmt.Errorf("couldn't get posts for user: %w", err)
	}

	fmt.Printf("Found %d posts for user %s:\n\n", len(posts), user.Name)

	for _, post := range posts {
		fmt.Printf("TITLE: %s\n", post.Title)
		fmt.Printf("PubDate: %s\n", post.PublishedAt.Time)
		fmt.Printf("URL: %s\n", post.Url)
		fmt.Printf("DESCRIPTION: %s\n", post.Description.String)
		fmt.Println()
	}

	return nil
}
