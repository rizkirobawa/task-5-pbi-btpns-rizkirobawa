# Build an API Using GOLANG
This is the final assignment of the BTPN Syariah Fullstack Developer Project Based Internship Program held by Rakamin Academy

## Features

- User Registration
- User Login
- Added Photo Data
- Displays Photo Data
- Updating Photo Data
- Deleting Photo Data
- Updating User Data
- Deleting User Accounts

## Modules required in this project

```go
go get -u github.com/gin-gonic/gin
go get -u github.com/asaskevich/govalidator
go get -u gorm.io/gorm
go get -u github.com/golang-jwt/jwt/v5
go get -u gorm.io/driver/mysql
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/google/uuid
```

# project structure

- `app/`: Application related files.
- `controllers/`: HTTP request controller or manager files.
- `database/`: Files related to database settings and connections.
- `helpers/` : Files for build JWT Token
- `middlewares/`: Files for application middleware.
- `models/`: Definition of models or data structures.
- `main.go`: Main file for running the application.
- `go.mod` : all modules installed in this project
- `go.sum`
- `README.md`

## How to Run a Project

Make sure to follow these steps to run this project in your development environment.

1. Install all required modules using the `go get` in command.
2. Set the database connection parameters (if required) in the `database/database.go` file.
3. Run the command `go run main.go` to start the server.

## Test API

#### Register User
```
POST /users/register 
```
```json
{
  "username": "username",
  "email": "email@gmail.com",
  "password": "password"
}
```

#### Login

```
  POST /users/login
```

```json
{
  "email": "email@gmail.com",
  "password": "password"
}
```

#### Update user

```
  PUT /users/:userId
```

```json
{
  "username": "newusername",
  "email": "newemail@gmail.com",
  "password": "newpassword"
}
```

#### Delete user

```
  DELETE /users/:userId
```

#### Add photos

```
  POST /photos
```

```json
{
  "title": "example",
  "caption": "amazing",
  "photourl": "https://example.com/example.jpg"
}
```

#### Update photos

```
  PUT /photos/:photoId
```

```json
{
  "title": "newphoto",
  "caption": "#newcaption",
  "photo_url": "http://example.com/newphoto.jpg"
}
```

#### Show photos

```
  GET /photos
```

```json
{
    "data": [
        {
            "id": "2435d7ee-3412-4b5a-8249-97e1041c37d1",
            "title": "example",
            "caption": "caption",
            "photourl": "https://example.com/example.jpg",
            "userid": "5bc51574-e63e-4b04-b1c9-706633288dc8"
        }
    ],
    "message": "success"
}
```

#### Delete photos

```
  DELETE "/photos/:photoId"
```

