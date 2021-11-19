package main

import (
	"lrn/calhoun-ddd/adapter"
	"lrn/calhoun-ddd/github"
	"lrn/calhoun-ddd/gitlab"
	"lrn/calhoun-ddd/mw"
	"net/http"
)

func main() {
	var myHandler http.Handler
	var us mw.UserService
	us = &adapter.GitLabUserService{
		Client: &gitlab.Client{
			Key: "1234-mnop",
		},
	}
	myHandler = mw.AuthMiddleware(us, "my-org-id", myHandler)

	var yeHandler http.Handler
	var th mw.UserService
	th = &adapter.GitHubUserService{
		Client: &github.Client{
			Key: "0987-qrst",
		},
	}
	yeHandler = mw.AuthMiddleware(th, "my-org-id", yeHandler)

}
