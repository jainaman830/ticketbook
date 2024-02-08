// server.go

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	pb "project/ticketbook/ticket/generated"
)

type trainServer struct {
	// You may want to define some storage for users and seats
	// For demonstration purposes, I'll just use maps
	users    map[string]pb.User
	sections map[string]pb.Section
	tickets  map[string]pb.Ticket // Map user email to ticket
	mu       sync.RWMutex         // Mutex to protect concurrent access to maps
	pb.UnimplementedTrainTicketingServer
}

func IsValidEmail(email string) bool {
	// Regular expression for validating email addresses
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func (t *trainServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {

	if req.FirstName == "" {
		return nil, errors.New("Provide first name")
	} else if req.LastName == "" {
		return nil, errors.New("Provide last name")
	} else if req.Email == "" {
		return nil, errors.New("Provide email address")
	} else if IsValidEmail(req.Email) {
		return nil, errors.New("Provide valid email address")
	}
	for _, user := range t.users {
		if strings.ToLower(strings.TrimSpace(req.Email)) == strings.ToLower(user.Email) {
			return nil, errors.New("Email already used.")
		}
	}
	timenow := time.Now().String()
	// Store User information
	t.mu.Lock()
	defer t.mu.Unlock()
	user := pb.User{
		UserID:     uuid.NewString(),
		FirstName:  strings.TrimSpace(req.FirstName),
		LastName:   strings.TrimSpace(req.LastName),
		Email:      strings.TrimSpace(req.Email),
		CreatedOn:  timenow,
		ModifiedOn: timenow,
	}
	t.users[user.UserID] = user

	return &user, nil
}
func (t *trainServer) GetUsers(ctx context.Context, req *pb.UseRequest) (*pb.AllUsers, error) {

	if req.UserID != "" {
		user, ok := t.users[req.UserID]
		if ok {
			return &pb.AllUsers{Users: []*pb.User{&user}}, nil
		} else {
			return nil, errors.New("Invalid User")
		}
	}
	allUsers := []*pb.User{}
	for _, user := range t.users {
		allUsers = append(allUsers, &user)
	}
	if len(allUsers) == 0 {
		return nil, errors.New("Users not found")
	}
	return &pb.AllUsers{Users: allUsers}, nil
}

func (t *trainServer) PurchaseTicket(ctx context.Context, req *pb.Ticket) (*pb.Ticket, error) {
	// Assign seat based on section
	section := "A"
	// Logic to assign section
	if len(t.tickets) > 0 {
		section = "B"
	}

	seat := fmt.Sprintf("%s%d", section, len(t.tickets)+1)
	req.SeatNumber = seat

	// Store ticket information
	t.mu.Lock()
	defer t.mu.Unlock()
	t.tickets[req.User.Email] = *req

	return req, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTrainTicketingServer(grpcServer, &trainServer{
		users:    make(map[string]pb.User),
		tickets:  make(map[string]pb.Ticket),
		sections: make(map[string]pb.Section),
	})
	grpcServer.Serve(lis)
}
