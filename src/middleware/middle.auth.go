package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/AnggaArdhinata/backend-go/src/library"
	
)

func AuthMiddle(role ...string) Middle {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var header string
			var valid bool

			if header = r.Header.Get("Authorization"); header == "" {
				library.Response("header not provide, please login", 401, true).Send(w)
				return
			}
			
			if !strings.Contains(header, "Bearer") {
				library.Response("invalid header type", 401, true).Send(w)
				return
			}

			
			tokens := strings.Replace(header, "Bearer ", "", -1)

			checkToken, err := library.CheckToken(tokens)
			if err != nil {
				library.Response(err.Error(), 401, true).Send(w)
				return
			}
			for _, v := range role {
				if v == checkToken.Role {
					valid = true
				}
			}
			if !valid {
				library.Response("you have not access to this feature", 401, true).Send(w)
				return
			}

			log.Println("Auth Middleware Pass")

			//** share userid to controller
			ctx := context.WithValue(r.Context(), "user", checkToken.User_Id)

			//** serve next middleware
			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}
