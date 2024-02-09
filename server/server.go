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
	users    map[string]pb.User
	sections map[string]pb.Section
	tickets  map[string]pb.Ticket
	mu       sync.RWMutex // Mutex to protect concurrent access to maps
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
func (t *trainServer) ModifyUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	if req.UserID == "" {
		return nil, errors.New("Provide user id")
	} else if req.FirstName == "" {
		return nil, errors.New("Provide first name")
	} else if req.LastName == "" {
		return nil, errors.New("Provide last name")
	} else if req.Email == "" {
		return nil, errors.New("Provide email address")
	} else if IsValidEmail(req.Email) {
		return nil, errors.New("Provide valid email address")
	}
	oldData, ok := t.users[req.UserID]
	if !ok {
		return nil, errors.New("Invalid User")
	}
	for _, user := range t.users {
		if strings.ToLower(strings.TrimSpace(req.Email)) == strings.ToLower(user.Email) && req.UserID != user.UserID {
			return nil, errors.New("Email already used.")
		}
	}
	timenow := time.Now().String()
	// Store User information
	t.mu.Lock()
	defer t.mu.Unlock()
	user := pb.User{
		UserID:     req.UserID,
		FirstName:  strings.TrimSpace(req.FirstName),
		LastName:   strings.TrimSpace(req.LastName),
		Email:      strings.TrimSpace(req.Email),
		CreatedOn:  oldData.CreatedOn,
		ModifiedOn: timenow,
	}
	t.users[user.UserID] = user

	return &user, nil
}
func (t *trainServer) RemoveUser(ctx context.Context, req *pb.UseRequest) (*pb.EmptyResponse, error) {
	_, ok := t.users[req.UserID]
	if !ok {
		return nil, errors.New("Invalid User")
	}
	delete(t.users, req.UserID)
	return nil, nil
}
func (t *trainServer) CreateSection(ctx context.Context, req *pb.CreateSectionRequest) (*pb.Section, error) {
	if req.Section == "" {
		return nil, errors.New("Provide section")
	} else if req.TotalSeats <= 0 {
		return nil, errors.New("Total seats must be greater than 0")
	}
	for _, section := range t.sections {
		if strings.ToLower(strings.TrimSpace(req.Section)) == strings.ToLower(section.Section) {
			return nil, errors.New("Section name already used.")
		}
	}
	timenow := time.Now().String()
	// Store User information
	t.mu.Lock()
	defer t.mu.Unlock()
	section := pb.Section{
		SectionID:      uuid.NewString(),
		Section:        strings.TrimSpace(req.Section),
		TotalSeats:     req.TotalSeats,
		AvailableSeats: req.TotalSeats,
		CreatedOn:      timenow,
		ModifiedOn:     timenow,
	}
	t.sections[section.SectionID] = section

	return &section, nil
}
func (t *trainServer) ViewSections(ctx context.Context, req *pb.SectionRequest) (*pb.AllSections, error) {
	if req.SectionID != "" {
		section, ok := t.sections[req.SectionID]
		if ok {
			return &pb.AllSections{Sections: []*pb.Section{&section}}, nil
		} else {
			return nil, errors.New("Invalid section")
		}
	}
	allSections := []*pb.Section{}
	for _, section := range t.sections {
		allSections = append(allSections, &section)
	}
	if len(allSections) == 0 {
		return nil, errors.New("Sections not found")
	}
	return &pb.AllSections{Sections: allSections}, nil
}
func (t *trainServer) DeleteSections(ctx context.Context, req *pb.SectionRequest) (*pb.EmptyResponse, error) {
	_, ok := t.sections[req.SectionID]
	if !ok {
		return nil, errors.New("Invalid Section")
	}
	delete(t.sections, req.SectionID)
	return nil, nil
}
func (t *trainServer) ModifySections(ctx context.Context, req *pb.Section) (*pb.Section, error) {
	if req.SectionID == "" {
		return nil, errors.New("Provide section id")
	} else if req.Section == "" {
		return nil, errors.New("Provide first name")
	} else if req.TotalSeats <= 0 {
		return nil, errors.New("Total seats must be greater than 0")
	}
	oldData, ok := t.sections[req.SectionID]
	if !ok {
		return nil, errors.New("Invalid Section")
	}
	for _, section := range t.sections {
		if strings.ToLower(strings.TrimSpace(req.Section)) == strings.ToLower(section.Section) && req.SectionID != section.SectionID {
			return nil, errors.New("Section name already used.")
		}
	}
	req.AvailableSeats = 0
	if req.TotalSeats < oldData.TotalSeats {
		seatDiff := oldData.TotalSeats - req.TotalSeats
		if seatDiff > oldData.AvailableSeats {
			return nil, errors.New("Tickets already booked for some seats. cancel those first or modify seats later.")
		}
		req.AvailableSeats = oldData.AvailableSeats - seatDiff
	} else {
		seatDiff := req.TotalSeats - oldData.TotalSeats
		req.AvailableSeats = oldData.AvailableSeats + seatDiff
	}
	timenow := time.Now().String()
	// Store User information
	t.mu.Lock()
	defer t.mu.Unlock()
	section := pb.Section{
		SectionID:      req.SectionID,
		Section:        strings.TrimSpace(req.Section),
		TotalSeats:     req.TotalSeats,
		AvailableSeats: req.AvailableSeats,
		CreatedOn:      oldData.CreatedOn,
		ModifiedOn:     timenow,
	}
	t.sections[section.SectionID] = section

	return &section, nil
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
	t.tickets[req.UserID] = *req

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
