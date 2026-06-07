package db

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"
	"strconv"

	"github.com/Leonfarhan/simple-social/internal/store"
)

var userList = []string{
	"olivia", "liam", "emma", "noah", "ava", "elijah", "sophia", "lucas", "isabella", "mason",
	"mia", "logan", "charlotte", "ethan", "amelia", "james", "harper", "aiden", "evelyn", "jackson",
	"abigail", "sebastian", "ella", "mateo", "scarlett", "jack", "grace", "owen", "chloe", "theodore",
	"victoria", "levi", "riley", "henry", "aria", "wyatt", "lily", "julian", "aurora", "leo",
	"zoey", "hudson", "nora", "ezra", "hannah", "lincoln", "ellie", "grayson", "layla", "isaac",
}

var titleList = []string{
	"Morning coffee thoughts", "Weekend hiking plan", "Learning Go today", "Favorite street food", "Late night coding",
	"Simple productivity tips", "New book recommendation", "Travel bucket list", "Home workout routine", "Music for focus",
	"Quick dinner recipe", "Small wins today", "Movie night picks", "Garden update", "Remote work setup",
	"Daily photo dump", "Fresh playlist drop", "Study session notes", "Local cafe review", "Sunset walk story",
	"Weekend project log", "Healthy habit check", "Tech news recap", "Pet moments", "Random life update",
}

var contentList = []string{
	"Started the day with a small win.", "Trying a new routine this week.", "Found a quiet spot to recharge.",
	"Sharing notes from a recent lesson.", "Planning something fun for the weekend.", "Taking a short break before the next task.",
	"Found a simple way to stay focused.", "Cooked something quick and tasty.", "Listening to a playlist that keeps me moving.",
	"Reading a chapter that feels useful.", "Taking a walk to clear my head.", "Writing down ideas before they disappear.",
	"Enjoying slow progress and consistency.", "Testing a new setup at home.", "Saving this moment for later.",
	"Looking back at what worked today.", "Trying to keep things simple.", "Learning from a small mistake.",
	"Spending time on a tiny project.", "Checking in after a busy day.", "Finding inspiration in ordinary things.",
	"Making room for a better habit.", "Sharing a quick update with everyone.", "Ending the day with gratitude.",
	"Keeping this one short and honest.",
}

var tagsList = []string{
	"go", "backend", "api", "database", "postgres", "web", "coding", "learning",
	"productivity", "travel", "food", "music", "books", "fitness", "movies", "nature",
	"photography", "work", "study", "lifestyle", "tech", "pets", "weekend", "coffee",
	"daily",
}

var commentList = []string{
	"Nice update!", "Thanks for sharing.", "This is really helpful.", "I can relate to this.", "Great point.",
	"Love this idea.", "Keep it up!", "This made my day.", "Interesting perspective.", "Saving this one.",
	"Well said.", "That sounds fun.", "I should try this too.", "Good reminder.", "Looks awesome.",
	"Totally agree.", "This is inspiring.", "Simple but useful.", "Glad you posted this.", "Very cool.",
	"That makes sense.", "What a nice moment.", "Appreciate the update.", "This feels familiar.", "Great share.",
}

func Seed(storage store.Storage) error {
	ctx := context.Background()

	users := generateUsers(50)
	for _, user := range users {
		if err := storage.Users.Create(ctx, user); err != nil {
			return fmt.Errorf("seed users: %w", err)
		}
	}

	posts := generatePosts(25, users)
	for _, post := range posts {
		if err := storage.Posts.Create(ctx, post); err != nil {
			return fmt.Errorf("seed posts: %w", err)
		}
	}

	comments := generateComments(200, users, posts)
	for _, comment := range comments {
		if err := storage.Comments.Create(ctx, comment); err != nil {
			return fmt.Errorf("seed comments: %w", err)
		}
	}

	log.Println("Seeding complete!")
	return nil
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := range num {
		name := userList[i%len(userList)] + strconv.Itoa(i+1)

		users[i] = &store.User{
			Username: name,
			Email:    name + "@mail.com",
			Password: "123456",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := range num {
		user := users[rand.IntN(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titleList[rand.IntN(len(titleList))],
			Content: contentList[rand.IntN(len(contentList))],
			Tags: []string{
				tagsList[rand.IntN(len(tagsList))],
				tagsList[rand.IntN(len(tagsList))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	comment := make([]*store.Comment, num)

	for i := range num {
		comment[i] = &store.Comment{
			PostID:  posts[rand.IntN(len(posts))].ID,
			UserID:  users[rand.IntN(len(users))].ID,
			Content: commentList[rand.IntN(len(commentList))],
		}
	}

	return comment
}
