package main

import (
	"blog/blogmgr"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func setBlogRouter(r *gin.RouterGroup) {
	//博客的增删改查
	r.GET("blogs/:id")
	r.POST("blogs", AddBlog)
	r.DELETE("blogs/:id")
	r.PUT("blogs/:id")
}

type AddBlogPostBody struct {
	CateId   int    `json:"cate_id" sql:"cate_id"`
	Kind     int    `json:"kind" sql:"kind"`
	Status   int    `json:"status" sql:"status"`
	Title    string `json:"title" sql:"title"`
	Path     string `json:"path" sql:"path"`
	Summary  string `json:"summary" sql:"summary"`
	Markdown string `json:"markdown" sql:"markdown"`
	RichText string `json:"richtext" sql:"richtext"`
	Allow    int    `json:"allow" sql:"allow"`
}

func AddBlog(c *gin.Context) {
	var pb AddBlogPostBody
	err := c.ShouldBindJSON(&pb)
	if err != nil {
		fileLogger.ErrErr("c.ShouldBindJSON failed", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var bmk = blogmgr.BlogKey{
		CateId:   0, //目前是瞎写的id
		Kind:     pb.Kind,
		Status:   pb.Status,
		Title:    pb.Title,
		Path:     pb.Path,
		Summary:  pb.Summary,
		Markdown: pb.Markdown,
		RichText: pb.RichText,
		Allow:    pb.Allow,
		Created:  time.Now(),
		Updated:  time.Now(),
		Creator:  2,
		Updater:  2,
	}

	blogMgr, err := blogmgr.NewBlogMgr(context.Background(), &fileLogger)
	if err != nil {
		fileLogger.ErrErr("blogmgr.NewBlogMgr failed", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = blogMgr.AddBlog(&bmk)
	if err != nil {
		fileLogger.ErrErr("blogMgr.AddBlog failed", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusBadRequest, "application/json", []byte("{}"))
}
