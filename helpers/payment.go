package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Pay ...
type Pay struct {
	Token        string
	PayOrder     payOrder
	Email        string
	Phone        string
	FirstName    string
	LastName     string
	OrderID      int
	PaymentToken string
}

// PayResponse ...
type payResponse struct {
	Token string `json:"token"`
}

type payOrder struct {
	Token           string `json:"auth_token"`
	DeliveryNeeded  string `json:"delivery_needed"`
	Amount          string `json:"amount_cents"`
	Currency        string `json:"currency"`
	Items           []int  `json:"items"`
}

type billingData struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone_number"`
	Apartment  string `json:"apartment" default:"NA"`
	Floor      string `json:"floor" default:"NA"`
	Street     string `json:"street" default:"NA"`
	Building   string `json:"building" default:"NA"`
	Shipping   string `json:"shipping_method" default:"NA"`
	PostalCode string `json:"postal_code" default:"NA"`
	Country    string `json:"country" default:"NA"`
	City       string `json:"city" default:"NA"`
	State      string `json:"state" default:"NA"`
}

type tokenOrder struct {
	Token         string `json:"auth_token"`
	Amount        string `json:"amount_cents"`
	Currency      string `json:"currency"`
	OrderID       int    `json:"order_id"`
	Expiration    int    `json:"expiration"`
	IntegrationID int    `json:"integration_id"`
	BillingData   billingData  `json:"billing_data"`
}
type orderData struct {
	ID int `json:"id"`
}

//PayAuth ...
func (p *Pay) PayAuth() {
	url := "https://accept.paymobsolutions.com/api/auth/tokens"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{"api_key":"ZXlKaGJHY2lPaUpJVXpVeE1pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SmpiR0Z6Y3lJNklrMWxjbU5vWVc1MElpd2ljSEp2Wm1sc1pWOXdheUk2TWpVeU1UWXNJbTVoYldVaU9pSnBibWwwYVdGc0luMC5FUU45aGtNeGlkRlR3YllYcElkOGRTMm81N19rOWVJcFdGNkR1NExLU2Vlellnb2dQMUYyN1BiMXQ1eWNNc1FIdkhkU1R3dzUyTFR4QzJCVlNEY3FHUQ=="}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
	res := &payResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		fmt.Println("response ERR:", err)

	}
	fmt.Println("response Body:", res.Token)

	p.Token = res.Token
}

//PlaceOrder ...
func (p *Pay) PlaceOrder(price int, orderID int) {
	url := "https://accept.paymobsolutions.com/api/ecommerce/orders"
	fmt.Println("URL:>", url)
	emptyArr := make([]int, 0)
	request := payOrder{
		Amount:          strconv.Itoa(price),
		Currency:        "EGP",
		DeliveryNeeded:  "false",
		Token:           p.Token,
		Items:           emptyArr,
	}
	jsonStr, _ := json.Marshal(&request)
	fmt.Println("Req Body Of Order", string(jsonStr))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	p.PayOrder = request
	orderData := &orderData{}
	json.Unmarshal(body, orderData)
	p.OrderID = orderData.ID
}

//GetToken Token ...
func (p *Pay) GetToken() {
	url := "https://accept.paymobsolutions.com/api/acceptance/payment_keys"
	fmt.Println("URL:>", url)

	billing := billingData{
		p.Email,
		p.FirstName,
		p.LastName,
		p.Phone,
		"NA",
		"NA",
                "NA",
                "NA",
                "NA",
                "NA",
                "NA",
                "NA",
                "NA",

	}
	request := &tokenOrder{
		Amount:        "100",
		BillingData:   billing,
		Currency:      "EGP",
		Token:         p.Token,
		Expiration:    7200000,
		OrderID:       p.OrderID,
		IntegrationID: 54470,
	}

	jsonStr, _ := json.Marshal(request)
	fmt.Println("TTEEEEEEEEEEEEE", string(jsonStr))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	pResp := &payResponse{}
	json.Unmarshal(body, pResp)
	p.PaymentToken = pResp.Token
}

//BuildIFrame ...
func (p *Pay) BuildIFrame() string {
	return fmt.Sprintf("https://accept.paymobsolutions.com/api/acceptance/iframes/%s?payment_token=%s", "74140", p.PaymentToken)
}
