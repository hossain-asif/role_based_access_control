package user

import (
	"context"
	"fmt"
	utils "go_project_structure/utils"
	"net/http"
)

func RequestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received at:", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func UserRegisterRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var RequestPayload = registerUserRequest{}
		if payloadErr := utils.ReadJsonBody(r, &RequestPayload); payloadErr != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Json encoding error.", payloadErr)
			return
		}
		fmt.Println("register payload received.")

		// validation logic for the user registration payload

		// context can be used to pass the validated payload to the handler for further processing
		req_context := r.Context()                                                    // parent context -> get the context from the request
		ctx := context.WithValue(req_context, "registration_payload", RequestPayload) // create a new context with the validated payload
		r = r.WithContext(ctx)                                                        // create a new request with the new context

		next.ServeHTTP(w, r)
	})
}

func UserUpdateRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var RequestPayload = updateUserRequest{}
		if payloadErr := utils.ReadJsonBody(r, &RequestPayload); payloadErr != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Json encoding error.", payloadErr)
			return
		}
		fmt.Println("update payload received.")

		// validation logic for the user registration payload

		next.ServeHTTP(w, r)
	})
}
