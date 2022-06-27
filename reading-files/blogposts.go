package blogposts


import (
	"io/fs"
	"io"
	"bufio"
	"strings"
	"fmt"
	
)


type Post struct {
	Title       string
	Description string 
	Tags        []string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)



func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")

	if err != nil {
		return nil, err
	}

	var posts []Post


	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)

	}

	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {

	postFile, err := fileSystem.Open(fileName)

	if err != nil {
		return Post{}, err
	}

	defer postFile.Close()

	
	return newPost(postFile)
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)


	readMetaLine := func(tagName string) string {
		scanner.Scan()

		fmt.Println(strings.TrimPrefix(scanner.Text(), tagName))
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	fmt.Println(readMetaLine(descriptionSeparator))

	return Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
		Tags:        strings.Split(readMetaLine(tagsSeparator), ", "),
	}, nil

}