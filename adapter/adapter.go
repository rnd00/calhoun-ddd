package adapter

import (
	"lrn/calhoun-ddd/github"
	"lrn/calhoun-ddd/gitlab"
	"lrn/calhoun-ddd/mw"
)

type GitHubUserService struct {
	Client *github.Client
}

func (us *GitHubUserService) User(token string) (mw.User, error) {
	ghUser, err := us.Client.User(token)
	if err != nil {
		return mw.User{}, err
	}
	return mw.User{
		OrgIDs: ghUser.OrgIDs,
	}, nil
}

type GitLabUserService struct {
	Client *gitlab.Client
}

func (us *GitLabUserService) User(token string) (mw.User, error) {
	glUser, err := us.Client.User(token)
	if err != nil {
		return mw.User{}, err
	}
	return mw.User{
		OrgIDs: glUser.OrgIDs,
	}, nil
}
