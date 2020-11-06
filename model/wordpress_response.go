package model

import "github.com/metamemelord/portfolio-website/pkg/core"

type WordpressResponse struct {
	Posts []WordPressPost `json:"posts"`
}

type WordPressPost struct {
	ID          int    `json:"ID"`
	OriginalUrl string `json:"URL"`
	Author      struct {
		Name string `json:"name"`
	} `author:"author"`
	Content string                 `json:"content"`
	Date    string                 `json:"date"`
	Tags    map[string]interface{} `json:"tags"`
	Excerpt string                 `json:"excerpt"`
	Title   string                 `json:"title"`
}

type Post struct {
	ID            int      `json:"_id"`
	Title         string   `json:"title"`
	Subtitle      string   `json:"subtitle"`
	Content       string   `json:"content"`
	Author        string   `json:"author"`
	Date          string   `json:"date"`
	Tags          []string `json:"tags"`
	AuthorContact string   `json:"author_context"`
	Excerpt       string   `json:"excerpt"`
}

func wordPressPostToPost(wpp *WordPressPost) *Post {
	post := &Post{
		ID:            wpp.ID,
		Title:         wpp.Title,
		Excerpt:       wpp.Excerpt,
		Subtitle:      wpp.Excerpt,
		Content:       wpp.Content,
		Author:        wpp.Author.Name,
		Date:          wpp.Date,
		AuthorContact: core.EMPTY_STRING,
		Tags:          []string{},
	}

	for tag := range wpp.Tags {
		post.Tags = append(post.Tags, tag)
	}

	return post
}

func WordPressResponseToCustomResponse(wpr WordpressResponse) []*Post {
	posts := []*Post{}
	for _, post := range wpr.Posts {
		posts = append(posts, wordPressPostToPost(&post))
	}
	return posts
}
