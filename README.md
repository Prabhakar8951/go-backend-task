\# Go Backend Development Task



\## Overview

This project is a simple Go backend application built using the Gin framework.

It provides REST APIs to create and retrieve users.



\## Tech Stack

\- Go (Golang)

\- Gin Web Framework

\- Postman (API Testing)



\## Project Structure

go-backend-task/

├── main.go

├── go.mod

├── go.sum

├── controllers/

│ └── user\_controller.go

├── models/

│ └── user.go

├── routes/

│ └── routes.go





\## API Endpoints

\- POST /users – Create a new user

\- GET /users – Get all users



\## Setup and Run Instructions



\### 1. Install Go

Download Go from:

https://go.dev/dl/



Verify installation:

go version



\### 2. Clone the Repository

go mod tidy



\### 4. Run the Application

go run main.go



The server will start at:

http://localhost:8080



\### 5. Test APIs

Use Postman:

\- POST http://localhost:8080/users

\- GET  http://localhost:8080/users



\## Sample Request Body

```json

{

&nbsp; "id": 1,

&nbsp; "name": "Prabhakar",

&nbsp; "email": "prabhakar@gmail.com"

}



Output



The API returns JSON responses confirming successful creation and retrieval of users.



➡️ \*\*Save\*\* and close Notepad.



---



\## 🔹 STEP 3: CONFIRM README EXISTS

Run:

```bat

dir

You MUST now see:

README.md



