## Golang Auth Mongo DB

## Requirements
- docker & docker-compose

## API:

### POST /auth/sign-up

Registers new user

#### Example input:
```
{
	"email": "email@gmail.com",
	"password": "qwerty123"
} 
```

### POST /auth/sign-in

Generates JWT Token

#### Example input:
```
{
	"email": "email@gmail.com",
	"password": "qwerty123"
} 
```

#### Example Response:

```
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzEwMzgyMjQuNzQ0MzI0MiwidXNlciI6eyJJRCI6IjAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMCIsIlVzZXJuYW1lIjoiemhhc2hrZXZ5Y2giLCJQYXNzd29yZCI6IjQyODYwMTc5ZmFiMTQ2YzZiZDAyNjlkMDViZTM0ZWNmYmY5Zjk3YjUifX0.3dsyKJQ-HZJxdvBMui0Mzgw6yb6If9aB8imGhxMOjsk"
} 
```

## Start 
To run app
```
make run 
```

