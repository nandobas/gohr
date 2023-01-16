package quiz

import (
	"context"
	"fmt"
	"time"
)

type User struct {
	string
}

func findUserInDB(login string) *User {
	time.Sleep(1 * time.Second)
	return &User{"nono with " + login}
}

func findUser(ctx context.Context, login string) (*User, error) {
	ch := make(chan *User, 1)
	go func() {
		ch <- findUserInDB(login)
	}()

	select {
	case user := <-ch:
		return user, nil
	case <-ctx.Done():
		return nil, fmt.Errorf("timeout")
	}
}

const shortDuration = 5 * time.Second

func Search() {
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	go execSearch(ctx)
	go execSearch(ctx)
	go execSearch(ctx)
	go execSearch(ctx)
	time.Sleep(10 * time.Second)
	fmt.Printf("end")

}

func execSearch(ctx context.Context) {
	u, err := findUser(ctx, "will")
	if err != nil {
		fmt.Errorf("erro on findUser: %w\n", err)
	}
	fmt.Printf("user find: %s\n", u.string)
}
