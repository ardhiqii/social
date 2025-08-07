package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/ardhiqii/social/internal/store"
)

var usernames = []string{
	"alex_brown842",
	"fiona_lee920",
	"charlie_miller317",
	"dana_smith476",
	"hannah_johnson102",
	"ivan_garcia783",
	"bella_davis611",
	"george_williams358",
	"julia_martin264",
	"ethan_jones730",
	"alex_smith592",
	"fiona_martin823",
	"charlie_brown137",
	"dana_lee449",
	"hannah_davis255",
	"ivan_williams911",
	"bella_jones680",
	"george_miller344",
	"julia_garcia578",
	"ethan_johnson106",
	"alex_miller903",
	"fiona_garcia771",
	"charlie_davis128",
	"dana_johnson666",
	"hannah_brown228",
	"ivan_smith315",
	"bella_martin704",
	"george_jones593",
	"julia_lee806",
	"ethan_williams459",
	"alex_davis942",
	"fiona_johnson604",
	"charlie_martin785",
	"dana_brown363",
	"hannah_smith231",
	"ivan_miller199",
	"bella_lee734",
	"george_garcia487",
	"julia_smith320",
	"ethan_martin868",
	"alex_jones500",
	"fiona_williams221",
	"charlie_garcia678",
	"dana_miller140",
	"hannah_martin964",
	"ivan_johnson281",
	"bella_brown902",
	"george_davis154",
	"julia_jones312",
	"ethan_lee447",
}

var titles = []string{
	"Why Go Is Awesome",
	"Mastering Concurrency",
	"Intro to REST APIs",
	"Understanding Pointers",
	"Error Handling in Go",
	"Goroutines Made Simple",
	"Working with JSON",
	"Using Context in Go",
	"Structs vs Interfaces",
	"Testing in Go",
	"Writing Clean Code",
	"Go Modules Explained",
	"Build a CLI in Go",
	"Database Access with SQL",
	"Memory Management Tips",
	"Go vs Python: A Comparison",
	"Logging Best Practices",
	"Creating Middleware",
	"Handling Timeouts",
	"Deploying Go Apps",
}

var contents = []string{
	"Go is a statically typed, compiled language known for its simplicity and performance. It’s great for building scalable backend services.",
	"Concurrency in Go is made easy with goroutines and channels, allowing you to handle multiple tasks efficiently.",
	"Building REST APIs in Go is straightforward with packages like `net/http` and popular frameworks like Gin or Echo.",
	"Pointers give you the power to reference and manipulate memory directly, enabling more efficient data handling in your applications.",
	"Proper error handling in Go uses simple patterns like returning errors as values, promoting clarity and control in your code.",
	"Goroutines are lightweight threads managed by the Go runtime, allowing you to perform async operations without complex thread management.",
	"JSON is widely used for data exchange, and Go provides powerful built-in support through the `encoding/json` package.",
	"Context in Go is essential for managing timeouts, cancellations, and deadlines, especially in web servers and APIs.",
	"Understanding the difference between structs and interfaces helps you write flexible, maintainable Go code.",
	"Testing in Go is built into the standard library, making it easy to write unit tests with clear output and coverage tracking.",
	"Clean code is about readability and simplicity. Go encourages this with its minimal syntax and standard formatting tools like `gofmt`.",
	"Go modules are the modern dependency management system, replacing GOPATH and making package versioning much easier.",
	"You can build powerful command-line tools in Go using the `flag` package or libraries like Cobra and urfave/cli.",
	"Accessing databases like PostgreSQL in Go is efficient using `database/sql` along with drivers and ORMs like GORM or sqlx.",
	"Understanding how Go handles memory, garbage collection, and allocation is key to writing optimized applications.",
	"Comparing Go and Python reveals strengths and trade-offs, especially in performance, syntax, and ecosystem maturity.",
	"Effective logging helps with debugging and monitoring. Libraries like `logrus` and `zap` enhance Go's basic `log` package.",
	"Middleware functions let you hook into request/response cycles, useful for logging, authentication, and more in HTTP servers.",
	"Timeouts prevent your app from hanging on long-running operations. Go's `context.WithTimeout` makes this easy to implement.",
	"Deploying Go apps is simple due to its single-binary output. You can use Docker or direct server deployment for production.",
}

var tags = []string{
	"go",
	"golang",
	"programming",
	"backend",
	"api",
	"rest",
	"web-development",
	"concurrency",
	"pointers",
	"errors",
	"testing",
	"cli",
	"json",
	"database",
	"sql",
	"context",
	"middleware",
	"performance",
	"deployment",
	"clean-code",
}

var contentComments = []string{
	"Great post! I finally understand how goroutines work.",
	"Thanks for explaining pointers — that really cleared things up!",
	"I was struggling with context cancellations, but this helped a lot.",
	"Awesome breakdown of error handling in Go. Simple and clear!",
	"This article saved me so much time. Thanks for writing it!",
	"Would love to see a follow-up on advanced concurrency patterns.",
	"Can you provide an example using GORM with PostgreSQL?",
	"Nice tutorial! Any recommendations for Go logging libraries?",
	"I didn’t know about `context.WithTimeout` — very useful tip.",
	"Clear and concise explanation. Keep up the good work!",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}

	posts := generatePosts(200, users)

	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)

	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment:", err)
			return
		}
	}

	log.Println("Seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password: "123123",
		}
	}
	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		var tempTags []string
		countTag := rand.Intn(len(tags))
		for j := 0; j < countTag; j++ {
			tempTags = append(tempTags, tags[rand.Intn(len(tags))])
		}

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags:    tempTags,
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	comments := make([]*store.Comment, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		post := posts[rand.Intn(len(posts))]

		comments[i] = &store.Comment{
			PostID:  post.ID,
			UserID:  user.ID,
			Content: contentComments[rand.Intn(len(contentComments))],
		}
	}
	return comments
}
