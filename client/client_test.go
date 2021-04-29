package client

import (
	"os"
	"terraform-provider-zoom_app/server"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetItem(t *testing.T) {
	testCases := []struct {
		testName     string
		itemName     string
		seedData     map[string]server.Item
		expectErr    bool
		expectedResp *server.Item
	}{
		{
			testName: "user exists",
			itemName: "user_mail@domain.com",
			seedData: map[string]server.Item{
				"user_mail@domain.com": {
					Id:        "[id retrieved from user account]",
					EmailId:   "user_mail@domain.com",
					FirstName: "firstname",
					LastName:  "lastname",
				},
			},
			expectErr: false,
			expectedResp: &server.Item{
				Id:        "",
				EmailId:   "",
				FirstName: "",
				LastName:  "",
			},
		},

		{
			testName:     "user does not exist",
			itemName:     "",
			seedData:     nil,
			expectErr:    true,
			expectedResp: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient("https://api.zoom.us/v2/users", os.Getenv("ZOOM_TOKEN"))

			item, err := client.GetItem(tc.itemName)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResp, item)
		})
	}
}

func TestClient_NewItem(t *testing.T) {
	testCases := []struct {
		testName  string
		newItem   *server.Item
		seedData  map[string]server.Item
		expectErr bool
	}{
		{
			testName: "success",
			newItem: &server.Item{
				Id:        "",
				EmailId:   "",
				FirstName: "",
				LastName:  "",
			},
			seedData:  nil,
			expectErr: false,
		},
		{
			testName: "item already exists",
			newItem: &server.Item{
				Id:        "",
				EmailId:   "",
				FirstName: "",
				LastName:  "",
			},
			seedData: map[string]server.Item{
				"item1": {
					Id:        "",
					EmailId:   "",
					FirstName: "",
					LastName:  "",
				},
			},
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient("https://api.zoom.us/v2/users", os.Getenv("ZOOM_TOKEN"))

			err := client.NewItem(tc.newItem)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			item, err := client.GetItem(tc.newItem.EmailId)
			assert.NoError(t, err)
			assert.Equal(t, tc.newItem, item)
		})
	}
}

//update

func TestClient_UpdateItem(t *testing.T) {
	testCases := []struct {
		testName    string
		updatedItem *server.Item
		seedData    map[string]server.Item
		expectErr   bool
	}{
		{
			testName: "item exists",
			updatedItem: &server.Item{
				Id:        "",
				EmailId:   "",
				FirstName: "",
				LastName:  "",
			},
			seedData: map[string]server.Item{
				"item1": {
					Id:        "",
					EmailId:   "",
					FirstName: "",
					LastName:  "",
				},
			},
			expectErr: false,
		},
		{
			testName: "item does not exist",
			updatedItem: &server.Item{
				Id:        "",
				EmailId:   "",
				FirstName: "",
				LastName:  "",
			},
			seedData:  nil,
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient("https://api.zoom.us/v2/users", os.Getenv("ZOOM_TOKEN"))
			err := client.UpdateItem(tc.updatedItem)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			item, err := client.GetItem(tc.updatedItem.EmailId)
			assert.NoError(t, err)
			assert.Equal(t, tc.updatedItem, item)
		})
	}
}

//Delete Testing

func TestClient_DeleteItem(t *testing.T) {
	testCases := []struct {
		testName  string
		itemName  string
		seedData  map[string]server.User
		expectErr bool
	}{
		{
			testName: "user exists",
			itemName: "[email_id]",
			seedData: map[string]server.User{
				"user1": {
					Id:        "",
					EmailId:   "",
					FirstName: "",
					LastName:  "",
				},
			},
			expectErr: false,
		},


	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			client := NewClient("https://api.zoom.us/v2/users", os.Getenv("ZOOM_TOKEN"))

			err := client.DeleteItem(tc.itemName)
			log.Println(err)
			if tc.expectErr {
				assert.Error(t, err)
				return
			}
			_, err = client.GetItem(tc.itemName)
			assert.Error(t, err)
		})
	}
}
