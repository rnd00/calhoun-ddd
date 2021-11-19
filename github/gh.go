package github

type User struct {
	ID     string
	Email  string
	OrgIDs []string
}

type Client struct {
	Key string
}

func (c *Client) User(token string) (User, error) {

	return User{}, nil
}
