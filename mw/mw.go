package mw

import (
	"net/http"
)

type User struct {
	OrgIDs []string
}

type UserService interface {
	User(token string) (User, error)
}

func AuthMiddleware(client UserService, reqOrgID string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		user, err := client.User(token)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		permit := false
		for _, orgID := range user.OrgIDs {
			if orgID == reqOrgID {
				permit = true
				break
			}
		}
		if !permit {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
