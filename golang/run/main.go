package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func readCSV(filename string) ([]map[string]string, error) {
	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Ensure there is at least a header row
	if len(records) < 2 {
		return nil, fmt.Errorf("CSV file must contain at least a header and one row")
	}

	// Extract headers from the first row
	headers := records[0]

	// Convert records to a slice of maps
	var data []map[string]string
	for _, row := range records[1:] { // Skip the header row
		rowMap := make(map[string]string)
		for i, value := range row {
			rowMap[headers[i]] = value
		}
		data = append(data, rowMap)
	}

	return data, nil
}

// API Credentials
const (
	BASE_URL = "https://b2.nomadinternet.com"
)

// to be loaded at the begining of the main function
var API_TOKEN = ""

// HttpRequest makes an HTTP request and returns the response body as a string
func HttpRequestZen(method, url string, body interface{}) ([]byte, error) {
	client := &http.Client{Timeout: 90 * time.Second}

	// Convert body to JSON if not nil
	var reqBody []byte
	var err error
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %v", err)
		}
	}

	// Create request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("Authorization", "Basic auth")

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("request failed during call to zendesk url %s. Error: %v", url, err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return respBody, fmt.Errorf("error reading response: %v. body of response %s", err, string(respBody))
	}

	// Check for non-2xx status codes
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return respBody, fmt.Errorf("API error calling %s: %s - Response Body: %s", url, resp.Status, string(respBody))
	}

	return respBody, nil
}

// HttpRequest makes an HTTP request and returns the response body as a string
func HttpRequest(method, url string, body interface{}) ([]byte, error) {
	client := &http.Client{Timeout: 60 * time.Second}

	// Convert body to JSON if not nil
	var reqBody []byte
	var err error
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %v", err)
		}
	}

	// Create request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return []byte{}, fmt.Errorf("error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("Authorization", fmt.Sprintf("Splynx-EA (access_token=%s)", API_TOKEN))

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed during call to url %s. Error: %v", url, err)
	}
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return respBody, fmt.Errorf("error reading API response: %v. Response body %s", err, string(respBody))
	}

	// Check for non-2xx status codes
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return respBody, fmt.Errorf("API error calling %s: %s - Response Body: %s", url, resp.Status, string(respBody))
	}

	return respBody, nil
}

// Customer represents the JSON structure
type Customer struct {
	MRRTotal             string     `json:"mrr_total"`
	GDPRAgreed           string     `json:"gdpr_agreed"`
	ID                   int        `json:"id"`
	BillingType          string     `json:"billing_type"`
	PartnerID            int        `json:"partner_id"`
	LocationID           int        `json:"location_id"`
	AddedBy              string     `json:"added_by"`
	AddedByID            int        `json:"added_by_id"`
	Login                string     `json:"login"`
	Category             string     `json:"category"`
	Name                 string     `json:"name"`
	Email                string     `json:"email"`
	BillingEmail         string     `json:"billing_email"`
	Phone                string     `json:"phone"`
	Street1              string     `json:"street_1"`
	ZipCode              string     `json:"zip_code"`
	City                 string     `json:"city"`
	Status               string     `json:"status"`
	DateAdd              string     `json:"date_add"`
	LastOnline           string     `json:"last_online"`
	LastUpdate           string     `json:"last_update"`
	DailyPrepaidCost     string     `json:"daily_prepaid_cost"`
	GPS                  string     `json:"gps"`
	ConversionDate       string     `json:"conversion_date"`
	Street2              string     `json:"street_2"`
	AdditionalAttributes Attributes `json:"additional_attributes"`
	// CustomerLabels       []string   `json:"customer_labels"`
}

// Attributes represents the nested additional_attributes field
type Attributes struct {
	HSMAC                        string `json:"hs_mac"`
	PaymentusPin                 string `json:"paymentus_pin"`
	PortaoneCustomerIDs          string `json:"portaone_customer_ids"`
	Referrer                     string `json:"referrer"`
	ReportFirstServiceAmount     string `json:"report_first_service_amount"`
	ReportFirstServiceCancelDate string `json:"report_first_service_cancel_date"`
	SelfRegistrationComment      string `json:"self_registration_comment"`
	SocialID                     string `json:"social_id"`
	SplynxAddonAgentsAgent       string `json:"splynx_addon_agents_agent"`
	SplynxAddonResellersReseller string `json:"splynx_addon_resellers_reseller"`
}

// Administrator represents the JSON structure
type Administrator struct {
	ID        int    `json:"id"`
	PartnerID int    `json:"partner_id"`
	RoleName  string `json:"role_name"`
	Login     string `json:"login"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Timeout   int    `json:"timeout"`
	LastIP    string `json:"last_ip"`
	LastDT    string `json:"last_dt"`
	// FilterPartnerID  *int             `json:"filter_partner_id"`
	RouterAccess    string          `json:"router_access"`
	Phone           string          `json:"phone"`
	SendFromMyName  string          `json:"send_from_my_name"`
	CalendarColor   string          `json:"calendar_color"`
	UpdatedAt       string          `json:"updated_at"`
	AdditionalAttrs AdditionalAttrs `json:"additional_attributes"`
	AvatarAPIURL    *string         `json:"avatar_api_url"`
	DownloadAvatar  *string         `json:"download_avatar_link"`
}

// AdditionalAttrs represents the additional_attributes field
type AdditionalAttrs struct {
	Cashdesk string `json:"cashdesk"`
}

// Ticket represents the main JSON structure
type Ticket struct {
	CustomerID int     `json:"customer_id,omitempty"`
	AssignTo   int     `json:"assign_to,omitempty"`
	StatusID   int     `json:"status_id"`
	Subject    string  `json:"subject"`
	Priority   string  `json:"priority"`
	Closed     bool    `json:"closed"`
	CreatedAt  string  `json:"created_at"`
	Message    Message `json:"message"`
}

// Message represents the nested message object
type Message struct {
	Message    string `json:"message"`
	CustomerID int    `json:"customer_id"`
}

// Ticket represents the main JSON structure
type CustomerTicket struct {
	CustomerID *int          `json:"customer_id,omitempty"` // Pointer to omit if empty
	AssignTo   *int          `json:"assign_to,omitempty"`   // Pointer to omit if empty
	StatusID   int           `json:"status_id"`
	Subject    string        `json:"subject"`
	Priority   string        `json:"priority"`
	Closed     bool          `json:"closed"`
	CreatedAt  string        `json:"created_at"`
	Message    TicketMessage `json:"message"`
}

// Message represents the nested message object
type TicketMessage struct {
	Message string `json:"message"`
}

// InArray checks if a value exists in a slice
func InArray[T comparable](val T, arr []T) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

// SendTicket is a function that will send the customer ticket to the API
func SendTicket(ticket CustomerTicket, ticketId string, wg *sync.WaitGroup) {
	defer wg.Done()

	getResponse, err := HttpRequestZen("GET", "https://lrlabs.zendesk.com/api/v2/tickets/"+ticketId+"/comments", nil)

	var ticketComments CommentsResponse
	err = json.Unmarshal(getResponse, &ticketComments)

	if err != nil {
		fmt.Printf("failed to GET ticket comments for ticket with ID %s and info %#v. error:%s. Response %s \n", ticketId, ticket, err, string(getResponse))
	} else {

		if len(ticketComments.Comments) > 0 {
			if strings.HasPrefix(ticketComments.Comments[0].Body, "Conversation with Twilio (SMS) User") && len(ticketComments.Comments) > 1 {
				ticket.Message.Message = ticketComments.Comments[1].HTMLBody

				// the subject can be something like // Conversation with Alec, usually the first chat is emoty
			} else if strings.HasPrefix(ticketComments.Comments[0].Body, ticket.Subject) && len(ticketComments.Comments) > 1 {
				ticket.Message.Message = ticketComments.Comments[1].HTMLBody
			} else {
				ticket.Message.Message = ticketComments.Comments[0].HTMLBody
			}
		}
	}

	// Endpoint to send the ticket to
	ticketURL := BASE_URL + "/api/2.0/admin/support/tickets"
	resp, err := HttpRequest("POST", ticketURL, ticket)
	if err != nil {
		fmt.Printf("Error creating ticket for ticket ID %s: Error: %v, Response :%s  \n", ticketId, err, string(resp))
		return
	}

	// Optionally, log or handle the response
	fmt.Printf("Ticket created for ticket ID %s successfully: %s \n\n", ticketId, string(resp))
}

// TokenResponse represents the structure of the JSON response
type TokenResponse struct {
	AccessToken            string `json:"access_token"`
	RefreshToken           string `json:"refresh_token"`
	AccessTokenExpiration  int64  `json:"access_token_expiration"`
	RefreshTokenExpiration int64  `json:"refresh_token_expiration"`
}

func LoginUser() string {
	// Endpoint to send the ticket to
	logiUrl := BASE_URL + "/api/2.0/admin/auth/tokens"

	// Perform a POST request
	postData := map[string]interface{}{
		"auth_type": "admin",
		"login":     "foladoyin",
		"password":  "cLMNFhIR",
	}

	resp, err := HttpRequest("POST", logiUrl, postData)
	if err != nil {
		log.Fatalf("Error logging in USER  %#v. Error %v \n", postData, err)
		return ""
	}

	var response TokenResponse
	err = json.Unmarshal(resp, &response)

	if err != nil {
		log.Fatalf("Error reading user token during login event. Response  %s. Error %v \n", string(resp), err)
		return ""
	}

	return response.AccessToken
}

func TrimMapValues(input map[string]string) map[string]string {
	trimmedMap := make(map[string]string)
	for key, value := range input {
		bom := "\ufeff"
		trimmedKey := strings.TrimSpace(strings.Trim(key, bom))
		trimmedMap[trimmedKey] = strings.TrimSpace(strings.Trim(value, bom))
	}
	return trimmedMap
}

func main() {

	API_TOKEN = LoginUser()
	getResponse, err := HttpRequest("GET", BASE_URL+"/api/2.0/admin/customers/customer", nil)
	if err != nil {
		fmt.Println("failed to GET customers info. error:", err)
		return
	}

	var customerInfo []Customer
	err = json.Unmarshal(getResponse, &customerInfo)

	if err != nil {
		fmt.Println("failed to Unmarshal customers information. error:", err)
		return
	}

	customerMailMap := map[string]Customer{}
	customerPhoneMap := map[string]Customer{}

	for _, val := range customerInfo {

		if _, ok := customerMailMap[val.Email]; !InArray(val.Email, []string{"", " "}) && !ok {
			customerMailMap[val.Email] = val
		}

		if _, ok := customerPhoneMap[val.Phone]; !InArray(val.Phone, []string{"", " "}) && !ok {
			customerPhoneMap[val.Phone] = val
		}

	}

	getAdminResponse, err := HttpRequest("GET", BASE_URL+"/api/2.0/admin/administration/administrators", nil)
	if err != nil {
		fmt.Println("failed to GET Admin info. error:", err)
		return
	}

	var adminsInfo []Administrator
	err = json.Unmarshal(getAdminResponse, &adminsInfo)

	if err != nil {
		fmt.Println("failed to Unmarshal customers information. error:", err)
		return
	}

	adminMainMap := map[string]Administrator{}

	for _, val := range adminsInfo {
		if _, ok := adminMainMap[val.Name]; !InArray(val.Name, []string{"", " "}) && !ok {
			adminMainMap[val.Name] = val
		}
	}

	fmt.Println("About reading CSV and running uploads")
	data, err := readCSV("test-run-5.csv")

	if err != nil {
		fmt.Println("failed to read csv file. Error:", err)
		return
	}

	chunkSize := 500
	totalElements := len(data)

	for i := 0; i < totalElements; i += chunkSize {
		end := i + chunkSize
		if end > totalElements {
			end = totalElements
		}

		API_TOKEN = LoginUser()

		chunk := data[i:end]

		// Channel to capture errors from goroutines
		// errCh := make(chan error, 100)
		// Limit the number of concurrent goroutines to 100
		maxGoroutines := 100
		sem := make(chan struct{}, maxGoroutines)
		var wg sync.WaitGroup

		for _, row := range chunk {
			row = TrimMapValues(row)
			email := strings.TrimSpace(strings.Trim(row["Email Address"], "'-"))
			phone := ""
			assignedTo := strings.TrimSpace(row["Assignee"])
			// '+1 (469) 600-5862, Twilio (SMS) User +17402433616
			if strings.HasPrefix(row["Requester"], "'+") || strings.HasPrefix(row["Requester"], "Twilio (SMS)") {
				phone = strings.TrimSpace(strings.Trim(row["Email Address"], "'"))
				phone = strings.TrimSpace(strings.Trim(row["Email Address"], "Twilio (SMS) User"))
			}

			ticket := CustomerTicket{}
			if email != "" {
				if em, found := customerMailMap[email]; found {
					ticket.CustomerID = &em.ID
					goto ContinueComputaion
				}
			}

			if phone != "" {
				if pho, found := customerPhoneMap[phone]; found {
					ticket.CustomerID = &pho.ID
					goto ContinueComputaion
				}
			}

		ContinueComputaion:
			if assignedTo != "" {
				if agentDetails, found := adminMainMap[assignedTo]; found {
					ticket.AssignTo = &agentDetails.ID
				}
			}

			layout := "1/2/06 15:04"

			createdTime, err := time.Parse(layout, row["Requested"])

			if err != nil {
				fmt.Printf("failed to parse the date %s. Error: %s  \n \n", row["Requested"], err)
			}
			ticket.CreatedAt = createdTime.Format("2006-01-02 15:04:05")
			ticket.Closed = true
			ticket.Priority = "medium"
			ticket.StatusID = 3
			ticket.Subject = row["Subject"]

			wg.Add(1)
			// Acquire a slot in the semaphore to limit concurrency
			sem <- struct{}{}

			go func(ticket CustomerTicket, zendeskTicketId string, wgAdd *sync.WaitGroup) {
				defer func() { <-sem }() // Release the slot in the semaphore

				// Send the ticket
				SendTicket(ticket, zendeskTicketId, wgAdd)

			}(ticket, row["ID"], &wg)
		}

		// Wait for all goroutines to finish
		wg.Wait()

		fmt.Printf("Finished all tasks for batch %d, starts at %d and stops at %d \n", i/chunkSize+1, i, end)

		fmt.Println("waiting for 4 minutes before starting another batch")
		time.Sleep(4 * time.Minute)
	}

	fmt.Println("All tickets have been correctly uploaded. Nice one nice one")
}

// CommentsResponse represents the entire response containing multiple comments
type CommentsResponse struct {
	Comments []Comment `json:"comments"`
	Count    int       `json:"count"`
}

// Comment represents an individual comment
type Comment struct {
	ID        int64  `json:"id"`
	Type      string `json:"type"`
	AuthorID  int64  `json:"author_id"`
	Body      string `json:"body"`
	HTMLBody  string `json:"html_body"`
	PlainBody string `json:"plain_body"`
	Public    bool   `json:"public"`
}

/*
step 1 - Read the zendex export
step 2 - Have two maps
	  1 - read customers information and map it with email to the info, also phone number to the info
	  2 - read administrator information and create a map holding admin name to the info
step 3 - loop though the zendex info, if customer is found via it's email or phone, then use the ID identified
	   - else create the customer and use the ID
step 4 - create the customers using go routines to make it faster
	   - call zendesk API to get the last message and use it

// Email Address address column

Uc3ZJ2biw7huUZOHBMN3AX5fOnsTHSJg7IpJcZIh
N/A
na
*/
