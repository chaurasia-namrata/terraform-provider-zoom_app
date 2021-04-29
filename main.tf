terraform {
  required_providers {
    zoom = {
      version = "0.2"
      source  = "hashicorp.com/edu/zoom"
    }
  }
}

provider "zoom" {
  address = "https://api.zoom.us/v2/users"
  token   = "[access_token]"
}


resource "zoom_user" "user1" {
   email      = "USER_MAIL@gmail.com"
   first_name = "[FIRSTNAME]"
   last_name  = "[LASTNAME]"
   status = "activate"
}

data "zoom_user" "user1" {
   id = "USER_MAIL@gmail.com"
}


output "user1" {
   value = data.zoom_user.user1
 }
