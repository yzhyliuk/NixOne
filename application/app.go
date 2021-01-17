package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nixLevelFour/data/comment"
	"nixLevelFour/data/post"
)

type App interface {
	Start() error
}
type nixApp struct {
}

func InitApp() App {
	return nixApp{}
}
func (n nixApp) Start() error {
	n.getPosts(7)
	return nil
}

func (n nixApp) getPosts(userID int) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts?userId=%d", userID)
	requestPost, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: can't get post of user #%d from url: %s \n", userID, url)
		return
	}
	var postArray []post.Post
	err = json.NewDecoder(requestPost.Body).Decode(&postArray)
	if err != nil {
		fmt.Println("Unable to unmarshal response body")
		return
	}
	for _, pst := range postArray {
		go n.getComments(pst.ID)
		go pst.Save()
	}
}
func (n nixApp) getComments(postID int) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?postId=%d", postID)
	requestComment, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: can't get post #%d from url: %s \n", postID, url)
		return
	}
	var commentsArray []comment.Comment
	err = json.NewDecoder(requestComment.Body).Decode(&commentsArray)

	if err != nil {
		fmt.Println("Unable to unmarshal response body")
		return
	}
	for _, comm := range commentsArray {
		go comm.Save()
	}
}
