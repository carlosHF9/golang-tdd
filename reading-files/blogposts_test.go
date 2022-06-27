package blogposts_test

import (
	blogposts "github.com/carlosHF9/blogposts"
	"testing"
	"testing/fstest"
	"errors"
	"io/fs"
	"reflect"
	
)


type StubFailingFS struct {

}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func TestNewBlogPosts(t * testing.T) {

	const (
		firstBody = `Title: Post 1 
		Description: Description 1 
		Tags: tdd, go`


		secondBody = `Title: Post 2 
		Description: Description 2 
		Tags: rust, borrow-checker`
	)

	fs := fstest.MapFS{
		"hello world.md": {Data: []byte(firstBody)},
		"hello world2.md": {Data: []byte(secondBody)},
	}

	posts, _ := blogposts.NewPostsFromFS(fs)

	

	assertPost(t, posts[0], blogposts.Post{
		Title:       " Post 1",
		Description: " Description 1",
		Tags:        []string{"tdd", "go"},
	})

	got := posts[0]
	want := blogposts.Post{Title: "Post 1"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

	// if err != nil {
	// 	t.Fatal(err)
	// }

	// if len(posts) != len(fs) {
	// 	t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	// }
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {

	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}