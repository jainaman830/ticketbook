// server.go

package main

import (
	"context"
	"errors"
	"log"
	"net"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	pb "project/ticketbook/ticket/generated"
)

type trainServer struct {
	users          map[string]*pb.User
	sections       map[string]*pb.Section
	tickets        map[string]*pb.Ticket
	seats          map[string][]int32
	allocatedSeats map[string]string
	mu             sync.RWMutex // Mutex to protect concurrent access to maps
	pb.UnimplementedTrainTicketingServer
}

func IsValidEmail(email string) bool {
	// Regular expression for validating email addresses
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
func (t *trainServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {

	if strings.TrimSpace(req.FirstName) == "" {
		return nil, errors.New("Provide first name")
	} else if strings.TrimSpace(req.LastName) == "" {
		return nil, errors.New("Provide last name")
	} else if strings.TrimSpace(req.Email) == "" {
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
	t.users[user.UserID] = &user

	return &user, nil
}
func (t *trainServer) GetUsers(ctx context.Context, req *pb.UseRequest) (*pb.AllUsers, error) {

	if strings.TrimSpace(req.UserID) != "" {
		user, ok := t.users[strings.TrimSpace(req.UserID)]
		if ok {
			return &pb.AllUsers{Users: []*pb.User{user}}, nil
		} else {
			return nil, errors.New("Invalid User")
		}
	}
	allUsers := []*pb.User{}
	for _, user := range t.users {
		allUsers = append(allUsers, user)
	}
	if len(allUsers) == 0 {
		return nil, errors.New("Users not found")
	}
	return &pb.AllUsers{Users: allUsers}, nil
}
func (t *trainServer) ModifyUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	if strings.TrimSpace(req.UserID) == "" {
		return nil, errors.New("Provide user id")
	} else if strings.TrimSpace(req.FirstName) == "" {
		return nil, errors.New("Provide first name")
	} else if strings.TrimSpace(req.LastName) == "" {
		return nil, errors.New("Provide last name")
	} else if strings.TrimSpace(req.Email) == "" {
		return nil, errors.New("Provide email address")
	} else if IsValidEmail(req.Email) {
		return nil, errors.New("Provide valid email address")
	}
	oldData, ok := t.users[strings.TrimSpace(req.UserID)]
	if !ok {
		return nil, errors.New("Invalid User")
	}
	for _, user := range t.users {
		if strings.ToLower(strings.TrimSpace(req.Email)) == strings.ToLower(user.Email) && strings.TrimSpace(req.UserID) != user.UserID {
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
	t.users[user.UserID] = &user

	return &user, nil
}
func (t *trainServer) RemoveUser(ctx context.Context, req *pb.UseRequest) (*pb.EmptyResponse, error) {
	userid := strings.TrimSpace(req.UserID)
	_, ok := t.users[userid]
	if !ok {
		return nil, errors.New("Invalid User")
	}
	_, tOk := t.tickets[userid]
	if tOk {
		return nil, errors.New("Cancel current tickets for this user then try again")
	}
	delete(t.users, userid)
	return nil, nil
}
func (t *trainServer) CreateSection(ctx context.Context, req *pb.CreateSectionRequest) (*pb.Section, error) {
	if strings.TrimSpace(req.Section) == "" {
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
	t.sections[section.SectionID] = &section
	seats := []int32{}
	for i := int32(1); i <= req.TotalSeats; i++ {
		seats = append(seats, i)
	}
	t.seats[section.SectionID] = seats

	return &section, nil
}
func (t *trainServer) ViewSections(ctx context.Context, req *pb.SectionRequest) (*pb.AllSections, error) {
	if strings.TrimSpace(req.SectionID) != "" {
		section, ok := t.sections[strings.TrimSpace(req.SectionID)]
		if ok {
			return &pb.AllSections{Sections: []*pb.Section{section}}, nil
		} else {
			return nil, errors.New("Invalid section")
		}
	}
	allSections := []*pb.Section{}
	for _, section := range t.sections {
		allSections = append(allSections, section)
	}
	if len(allSections) == 0 {
		return nil, errors.New("Sections not found")
	}
	return &pb.AllSections{Sections: allSections}, nil
}
func (t *trainServer) ModifySections(ctx context.Context, req *pb.ModifySectionRequest) (*pb.Section, error) {
	if strings.TrimSpace(req.SectionID) == "" {
		return nil, errors.New("Provide section id")
	} else if strings.TrimSpace(req.Section) == "" {
		return nil, errors.New("Provide section name")
	}
	oldData, ok := t.sections[strings.TrimSpace(req.SectionID)]
	if !ok {
		return nil, errors.New("Invalid Section")
	}
	for _, section := range t.sections {
		if strings.ToLower(strings.TrimSpace(req.Section)) == strings.ToLower(section.Section) && strings.TrimSpace(req.SectionID) != section.SectionID {
			return nil, errors.New("Section name already used.")
		}
	}
	timenow := time.Now().String()
	// Store User information
	t.mu.Lock()
	defer t.mu.Unlock()
	section := pb.Section{
		SectionID:      oldData.SectionID,
		Section:        strings.TrimSpace(req.Section),
		TotalSeats:     oldData.TotalSeats,
		AvailableSeats: oldData.AvailableSeats,
		CreatedOn:      oldData.CreatedOn,
		ModifiedOn:     timenow,
	}
	t.sections[section.SectionID] = &section

	return &section, nil
}
func (t *trainServer) PurchaseTicket(ctx context.Context, req *pb.TicketRequest) (*pb.Ticket, error) {
	if strings.TrimSpace(req.From) == "" {
		return nil, errors.New("From can not be blank")
	} else if strings.TrimSpace(req.To) == "" {
		return nil, errors.New("To can not be blank")
	} else if strings.TrimSpace(req.UserID) == "" {
		return nil, errors.New("UserID can not be blank")
	} else if req.PricePaid < 0 {
		return nil, errors.New("Price paid can not be less than 0")
	} else if strings.TrimSpace(req.From) == strings.TrimSpace(req.To) {
		return nil, errors.New("From and to can not be same")
	}
	if _, uok := t.users[strings.TrimSpace(req.UserID)]; !uok {
		return nil, errors.New("Invalid user")
	}
	if _, tok := t.tickets[strings.TrimSpace(req.UserID)]; tok {
		return nil, errors.New("Ticket already book for this user")
	}
	availableSection := ""
	for _, section := range t.sections {
		if section.AvailableSeats > 0 {
			availableSection = strings.TrimSpace(section.SectionID)
			break
		}
	}
	if availableSection == "" {
		return nil, errors.New("All seats are booked!")
	}
	seats := t.seats[availableSection]
	var seatnumber int32
	for _, seat := range seats {
		_, seatOK := t.allocatedSeats[availableSection+"_"+strconv.Itoa(int(seat))]
		if !seatOK {
			seatnumber = seat
			break
		}
	}
	timenow := time.Now().String()
	ticket := &pb.Ticket{
		TicketId:   uuid.NewString(),
		From:       strings.TrimSpace(req.From),
		To:         strings.TrimSpace(req.To),
		UserID:     strings.TrimSpace(req.UserID),
		PricePaid:  req.PricePaid,
		Section:    availableSection,
		SeatNumber: seatnumber,
		CreatedOn:  timenow,
		ModifiedOn: timenow,
	}
	// Store ticket information
	t.mu.Lock()
	defer t.mu.Unlock()
	t.tickets[req.UserID] = ticket
	section := t.sections[availableSection]
	section.AvailableSeats = section.AvailableSeats - 1
	t.sections[availableSection] = section
	t.allocatedSeats[availableSection+"_"+strconv.Itoa(int(seatnumber))] = ticket.UserID
	return ticket, nil
}
func (t *trainServer) ViewReceipt(ctx context.Context, req *pb.UseRequest) (*pb.Receipt, error) {
	userid := strings.TrimSpace(req.UserID)
	if userid == "" {
		return nil, errors.New("User id can not be blank")
	}
	user, uOk := t.users[userid]
	if !uOk {
		return nil, errors.New("Invalid user")
	}
	ticket, tOk := t.tickets[userid]
	if !tOk {
		return nil, errors.New("Ticket not found")
	}
	receipt := &pb.Receipt{
		From:       ticket.From,
		To:         ticket.To,
		User:       user,
		PricePaid:  ticket.PricePaid,
		Section:    ticket.Section,
		SeatNumber: ticket.SeatNumber,
		CreatedOn:  ticket.CreatedOn,
		ModifiedOn: ticket.ModifiedOn,
	}
	return receipt, nil
}
func (t *trainServer) ViewSeatsBySection(ctx context.Context, req *pb.SectionRequest) (*pb.SeatAllocation, error) {
	sectionId := strings.TrimSpace(req.SectionID)
	_, sOk := t.sections[sectionId]
	if !sOk {
		return nil, errors.New("Invalid section")
	}
	seats := []*pb.SeatDetails{}
	for key, userid := range t.allocatedSeats {
		keyArr := strings.Split(key, "_")
		if sectionId == keyArr[0] {
			user := t.users[userid]
			seatNumber, _ := strconv.Atoi(keyArr[1])
			seatdetail := pb.SeatDetails{
				UserName:   user.FirstName + " " + user.LastName,
				Email:      user.Email,
				SeatNumber: int32(seatNumber),
			}
			seats = append(seats, &seatdetail)
		}
	}
	if len(seats) == 0 {
		return nil, errors.New("No seats allocated")
	}
	return &pb.SeatAllocation{Tickets: seats}, nil
}
func (t *trainServer) CancelReceipt(ctx context.Context, req *pb.UseRequest) (*pb.EmptyResponse, error) {
	userid := strings.TrimSpace(req.UserID)
	ticket, ok := t.tickets[userid]
	if !ok {
		return nil, errors.New("No ticket found for user")
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	section := t.sections[ticket.Section]
	section.AvailableSeats += 1
	t.sections[ticket.Section] = section
	delete(t.allocatedSeats, ticket.Section+"_"+strconv.Itoa(int(ticket.SeatNumber)))
	delete(t.tickets, userid)
	return nil, nil
}
func (t *trainServer) ModifySeat(ctx context.Context, req *pb.ModifySeatRequest) (*pb.Ticket, error) {
	userid := strings.TrimSpace(req.UserID)
	reqSection := strings.TrimSpace(req.Section)
	if userid == "" {
		return nil, errors.New("User can not be blank")
	}
	if req.SeatNumber < 1 {
		return nil, errors.New("Invalid seat number")
	}

	ticket, ok := t.tickets[userid]
	if !ok {
		return nil, errors.New("No ticket found for user")
	}
	section, sOk := t.sections[reqSection]
	if !sOk {
		return nil, errors.New("Invalid Section")
	}
	if req.SeatNumber > section.TotalSeats {
		return nil, errors.New("Seat number can not be more than total seats")
	}
	allocateduserid, sOk := t.allocatedSeats[reqSection+"_"+strconv.Itoa(int(req.SeatNumber))]
	if sOk && allocateduserid != userid {
		return nil, errors.New("Requested seat already allocated to other user")
	}
	delete(t.allocatedSeats, ticket.Section+"_"+strconv.Itoa(int(ticket.SeatNumber)))
	if ticket.Section != reqSection {
		prevSection := t.sections[ticket.Section]
		prevSection.AvailableSeats += 1
		t.sections[ticket.Section] = prevSection

		currentSection := t.sections[reqSection]
		currentSection.AvailableSeats -= 1
		t.sections[reqSection] = currentSection
	}
	t.allocatedSeats[reqSection+"_"+strconv.Itoa(int(req.SeatNumber))] = userid
	ticket.Section = reqSection
	ticket.SeatNumber = req.SeatNumber
	ticket.ModifiedOn = time.Now().String()
	t.tickets[userid] = ticket
	return ticket, nil
}
func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTrainTicketingServer(grpcServer, &trainServer{
		users:          make(map[string]*pb.User),
		tickets:        make(map[string]*pb.Ticket),
		sections:       make(map[string]*pb.Section),
		seats:          make(map[string][]int32),
		allocatedSeats: make(map[string]string),
	})
	grpcServer.Serve(lis)
}
