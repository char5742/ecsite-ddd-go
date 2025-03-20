package userhandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	sharetypes "github.com/char5742/ecsite-ddd-go/internal/share/domain/types"
	shareinfra "github.com/char5742/ecsite-ddd-go/internal/share/infra"
	userdomain "github.com/char5742/ecsite-ddd-go/internal/user/domain"
	userpersistence "github.com/char5742/ecsite-ddd-go/internal/user/persistence"
	usertypes "github.com/char5742/ecsite-ddd-go/internal/user/types"
	userworkflows "github.com/char5742/ecsite-ddd-go/internal/user/types/workflows"
	userusecases "github.com/char5742/ecsite-ddd-go/internal/user/usecases"
)

type RegisterUserHandler struct{}

func (h RegisterUserHandler) Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	command, err := h.req2command(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taken := userpersistence.NewIsEmailTaken(ctx)

	usecase := userusecases.NewRegisterUserUsecase(taken, func(vu userworkflows.ValidateUser, ru userworkflows.RegistUser) userworkflows.RegisterUserWorkflow {
		return func(user userdomain.UnvalidatedUser) ([]userworkflows.RegisterUserEvent, error) {
			validatedUser, err := vu(user)
			if err != nil {
				return nil, err
			}
			registedUser, err := ru(*validatedUser)
			if err != nil {
				return nil, err
			}
			return []userworkflows.RegisterUserEvent{
				userworkflows.UserRegistered{RegistedUser: *registedUser},
			}, nil
		}
	})

	events, err := usecase(*command)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(events) == 0 {
		http.Error(w, "No events generated", http.StatusInternalServerError)
		return
	}
	for _, event := range events {
		shareinfra.PublishEvent(event)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered successfully",
	}); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func (h RegisterUserHandler) req2command(r *http.Request) (*sharetypes.Command[usertypes.RegisterUserCommand], error) {
	var command usertypes.RegisterUserCommand
	if err := json.NewDecoder(r.Body).Decode(&command); err != nil {
		return nil, err
	}
	if err := h.validate(command); err != nil {
		return nil, err
	}
	return &sharetypes.Command[usertypes.RegisterUserCommand]{
		Context:    r.Context(),
		Timestamp:  time.Now(),
		IdentityID: nil,
		Data:       command,
	}, nil
}
func (h RegisterUserHandler) validate(command usertypes.RegisterUserCommand) error {
	if command.FirstName == "" {
		return fmt.Errorf("first name is required")
	}
	if command.LastName == "" {
		return fmt.Errorf("last name is required")
	}
	if command.Email == "" {
		return fmt.Errorf("email is required")
	}
	if command.Zipcode == "" {
		return fmt.Errorf("zipcode is required")
	}
	if command.Prefecture == "" {
		return fmt.Errorf("prefecture is required")
	}
	if command.Municipalities == "" {
		return fmt.Errorf("municipalities is required")
	}
	if command.Address == "" {
		return fmt.Errorf("address is required")
	}
	if command.Telephone == "" {
		return fmt.Errorf("telephone is required")
	}
	return nil
}
