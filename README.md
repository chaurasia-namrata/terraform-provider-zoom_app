# Terraform Zoom Provider

This terraform provider allows to perform Create ,Read ,Update, Delete and Deactivate Zoom User(s). 


## Requirements

* [Go](https://golang.org/doc/install) 1.16 <br>
* [Terraform](https://www.terraform.io/downloads.html) 0.13.x <br/>
* [Zoom](https://zoom.us/) Pro/Premium account (Token)


### Setup Zoom Account
 :heavy_exclamation_mark:  [IMPORTANT] : This provider can only be successfully tested on a premium paid zoom account. <br><br>

1. Create a zoom account with paid subscription (PRO Plan/Business Account). (https://zoom.us/)<br>
2. Sign in to the zoom account.<br>
3. Go to [Zoom Marketplace](https://marketplace.zoom.us/)<br>
4. Click on `Build App`. For our purpose we need to make a JWT App. <br>
5. Follow this [Create JWT Zoom App](https://marketplace.zoom.us/docs/guides/build/jwt-app) website to make an app. <br>
This app will provide us with the token which will be needed to configure our provider and make request. <br>


### Initialise Zoom Provider in local machine 
1. Clone the repository 
```
git clone [RESPOSITORY_GITHUB_LINK] && cd terraform-provider-zoom
``` 
<br>

2. Add the token generated in the JWT Zoom App to the token field in `main.tf` <br>
3. Run the following command :
 ```
go mod init terraform-provider-zoom
go mod tidy
```
4. Run `go mod vendor` to create a vendor directory that contains all the provider's dependencies. <br>
5. Run the following command to create a vendor subdirectory which will comprise of  all provider dependencies. <br>
```
~/.terraform.d/plugins/${host_name}/${namespace}/${type}/${version}/${target}
``` 
Command: 
```
mkdir -p ~/.terraform.d/plugins/hashicorp.com/zoom/0.2.0/[OS_ARCH]
```
For eg. `mkdir -p ~/.terraform.d/plugins/hashicorp.com/zoom/0.2.0/windows_amd64`<br>

6. Run `go build -o terraform-provider-zoom.exe`. This will save the binary (`.exe`) file in the main/root directory. <br>
7. Run this command to move this binary file to appropriate location.
 ```
 move terraform-provider-zoom.exe %APPDATA%\terraform.d\plugins\hashicorp.com\zoom\0.2.0\[OS_ARCH]
 ``` 
Otherwise you can manually move the file from current directory to destination directory.<br>


## Installation

1. Download required binaries <br>
2. move binary `~/.terraform.d/plugins/[architecture name]/`


## Run the Terraform provider

#### Create User
1. Add the user email, first name, last name in the respective field in `main.tf`
2. Initialize the terraform provider by the command `terraform init`
3. Check the changes applicable using `terraform plan`
4. Apply the changes using command `terraform apply`
5. Write yes to the prompt
6. You will see that a user has been successfully created and an account activation mail has been sent to the user.
7. Activate the account using the link provided in the mail.

#### Update the user
1. Update the data of the user in the `main.tf` file
2. Run the command `terraform init`
3. Check the changes applicable using `terraform plan`
4. Apply the changes using command `terraform apply`
5. Write yes to the prompt
6. You will see that a user data has been successfully updated.

#### Read the User Data
1. Add the data and resource blocks of the `main.tf` file
2. Run the command `terraform plan`
3. This will fetch you all the necessary details of the user.

#### Deactivate the user
1. Change the status of User from activate to `deactivate`.
2. Run the command `terraform apply` and write `yes` for the prompt.

#### Delete the user
1. Delete the resource block of the particular user from main.tf file 
2. Run the command `terraform apply` and write `yes` for the prompt.

#### Import a User Data
1. Write manually a resource configuration block for the User in main.tf, to which the imported object will be mapped.
2. Run the command `terraform import zoom_user.sample [user_id]`
3. Check for the attributes in the `.tfstate` file


### Testing the Provider
1. Navigate to the test file directory.
2. Run command `go test` . This command will give combined test result for the execution or errors if any failure occur.
3. If you want to see test result of each test function individually while running test in a single go, run command `go test -v`
4. To check test cover run `go test -cover`


## Example Usage
```terraform
terraform {
  required_providers {
    zoom = {
      version = "0.2"
      source  = "hashicorp.com/zoom"
    }
  }
}

provider "zoom" {
  address = "https://api.zoom.us/v2/users"
  token   = "access_token"
}

resource "zoom_user" "user1" {
   email      = "[EMAIL_ID]"
   first_name = "[USER_FIRST_NAME]"
   last_name  = "[USER_LAST_NAME]"
   status = "activate"
}

data "zoom_user" "user1" {
  id = "[EMAIL_ID]"
}

output "user1" {
  value = data.zoom_user.user1
}
```

