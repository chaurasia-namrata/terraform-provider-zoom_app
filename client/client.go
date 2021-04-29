package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"terraform-provider-zoom_app/server"
	"os"
	"log"
)


var (
    Errors = make(map[int]string)
)

func init() {
	Errors[400] = "Bad Request, StatusCode = 400"
	Errors[404] = "User Does Not Exist , StatusCode = 404"
	Errors[409] = "User Already Exist, StatusCode = 409"
	Errors[401] = "Unautharized Access, StatusCode = 401"
	Errors[429] = "User Has Sent Too Many Request, StatusCode = 429"
}



// Client -
type Client struct {
	hostname   string
	authToken  string
	httpClient *http.Client
}



// NewClient -
func NewClient(hostname string, token string) *Client {
	return &Client{
		hostname:   hostname,
		authToken:  token,
		httpClient: &http.Client{},
	}
}

func (c *Client) NewItem(item *server.Item) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		log.Println("[CREATE ERROR]: ", err)
		return err
	}
	_, err = c.httpRequest("POST", buf, item)
	if err != nil {
		log.Println("[CREATE ERROR]: ", err)
		return err
	}
	return nil
}

//Get User
func (c *Client) GetItem(name string) (*server.Item, error) {
	body, err := c.gethttpRequest(fmt.Sprintf("%v", name), "GET", bytes.Buffer{})
	if err != nil {
		log.Println("[READ ERROR]: ", err)
		return nil, err
	}
	item := &server.Item{}
	err = json.NewDecoder(body).Decode(item)
	if err != nil {
		log.Println("[READ ERROR]: ", err)
		return nil, err
	}
	return item, nil
}

func (c *Client) gethttpRequest(emailid, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, "https://api.zoom.us/v2/users/"+emailid, &body)
	if err != nil {
		log.Println("[ERROR]: ", err)
		return nil, err
	}
	var bearer = "Bearer " + os.Getenv("ZOOM_TOKEN")

	
	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("[ERROR]: ",err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("Error : %v",Errors[resp.StatusCode] )
		}
		return nil, fmt.Errorf("Error : %v ", Errors[resp.StatusCode])
	}

	return resp.Body, nil
}


//update
func (c *Client) UpdateItem(item *server.Item) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		log.Println("[UPDATE ERROR]: ", err)
		return err
	}
	_, err = c.updatehttpRequest(fmt.Sprintf("%s", item.EmailId), "PATCH", buf,item)
	if err != nil {
		log.Println("[UPDATE ERROR]: ", err)
		return err
	}
	return nil
}


func (c *Client) updatehttpRequest(path,method string, body bytes.Buffer, item *server.Item) (closer io.ReadCloser, err error) {

	updateuserjson := server.UpdateUser{
		FirstName: item.FirstName,
		LastName:  item.LastName,
	}
	updatejson, _ := json.Marshal(updateuserjson)
	payload := strings.NewReader(string(updatejson))

	req, err := http.NewRequest(method, c.updaterequestPath(path), payload)
	var str1 string
	str1 = "Bearer "
	var str2 string
	str2 = os.Getenv("ZOOM_TOKEN")
	req.Header.Add("Authorization", str1+str2)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Println("[ERROR]: ", err)
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Println("[ERROR]: ", err)
		return nil, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode <= 400 {
		return resp.Body, nil
    } else {
		return nil, fmt.Errorf("Error : %v",Errors[resp.StatusCode] )
    }
	return resp.Body, nil
}

func (c *Client) updaterequestPath(path string) string {
	return fmt.Sprintf("%s/%s",os.Getenv("ZOOM_ADDRESS"), path)

}

//Delete

func (c *Client) DeleteItem(userId string) error {
	_, err := c.deletehttpRequest(fmt.Sprintf("%s", userId), "DELETE", bytes.Buffer{})
	if err != nil {
		log.Println("[DELETE ERROR]: ", err)
		return err
	}
	return nil
}


func (c *Client) deletehttpRequest(path, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, c.deleterequestPath(path), &body)
	if err != nil {
		log.Println("[DELETE ERROR]: ", err)
		return nil, err
	}


	var str1 string
	str1 = "Bearer "
	var str2 string
	str2 = os.Getenv("ZOOM_TOKEN")
	req.Header.Add("Authorization", str1+str2)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Println("[DELETE ERROR]: ", err)
		return nil, err
	}


	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return resp.Body, nil
    } else {
		log.Println("Broken Request")
		return nil, fmt.Errorf("Error : %v",Errors[resp.StatusCode] )
    }

}

func (c *Client) deleterequestPath(path string) string {
	return fmt.Sprintf("%s/%s", os.Getenv("ZOOM_ADDRESS"), path)
}


func (c *Client) httpRequest(method string, body bytes.Buffer, item *server.Item) (closer io.ReadCloser, err error) {
	
	userjson := server.NewUser{
		Action: "create",
		UserInfo: server.UserInfo{
			EmailId:   item.EmailId,
			Type:      1,
			FirstName: item.FirstName,
			LastName:  item.LastName,
		},
	}
	reqjson, _ := json.Marshal(userjson)
	payload := strings.NewReader(string(reqjson))

	req, err := http.NewRequest(method, c.requestPath(), payload)
	var str1 string
	str1 = "Bearer "
	var str2 string
	str2 = os.Getenv("ZOOM_TOKEN")
	req.Header.Add("Authorization", str1+str2)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Println("[ERROR]: ",err)
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Println("[ERROR]: ",err)
		return nil, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return resp.Body, nil
    } else {
		return nil, fmt.Errorf("Error : %v",Errors[resp.StatusCode] )
    }
}

func (c *Client) requestPath() string {
	return fmt.Sprintf("%s?access_token=%s", os.Getenv("ZOOM_ADDRESS"), os.Getenv("ZOOM_TOKEN"))

}

// Deactivate user
func (c *Client) DeactivateUser(userId string, status string) error {
	log.Println("Changing Status of User : ", userId)
	url := fmt.Sprintf("https://api.zoom.us/v2/users/%s/status", userId)
	data := fmt.Sprintf("{\"action\":\"%s\"}", status)
	payload := strings.NewReader(data)

	req, err := http.NewRequest("PUT", url, payload)
	if err != nil {
		log.Println("[DEACTIVATE/ACTIVATE ERROR]: ",err)
		return nil
	}
	var str1 string
	str1 = "Bearer "
	var str2 string
	str2 = os.Getenv("ZOOM_TOKEN")

	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", str1+str2)
	_, err = c.httpClient.Do(req)
	if err != nil {
		log.Println("[DEACTIVATE/ACTIVATE ERROR]: ",err)
		return nil
	}
	return nil
}
