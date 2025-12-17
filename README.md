\# Go Backend Development Task



\## Overview

This project is a backend REST API application developed using \*\*Go (Golang)\*\* and the \*\*Gin web framework\*\*.

It implements full CRUD operations for user management, including dynamic age calculation based on date of birth.

The project follows clean backend structure and REST API standards.



---



\## Tech Stack

\- \*\*Language:\*\* Go (Golang)

\- \*\*Framework:\*\* Gin

\- \*\*API Testing Tool:\*\* Postman

\- \*\*Version Control:\*\* Git \& GitHub



---



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

└── README.md





---



\## Database Schema (Logical)

| Field | Type   | Constraints |

|------|--------|------------|

| id   | SERIAL | Primary Key |

| name | TEXT   | NOT NULL    |

| dob  | DATE   | NOT NULL    |



\*(Note: Current implementation uses in-memory storage. Schema represents intended DB design.)\*



---



\## API Endpoints



\### Create User

\*\*POST\*\* `/users`



\*\*Request:\*\*

```json

{

&nbsp; "name": "Alice",

&nbsp; "dob": "1990-05-10"

}

{

&nbsp; "id": 1,

&nbsp; "name": "Alice",

&nbsp; "dob": "1990-05-10"

}

Get User by ID



GET /users/:id



Response:

{

&nbsp; "id": 1,

&nbsp; "name": "Alice",

&nbsp; "dob": "1990-05-10",

&nbsp; "age": 35

}

Update User



PUT /users/:id



Request:

{

&nbsp; "name": "Alice Updated",

&nbsp; "dob": "1991-03-15"

}

Response:

{

&nbsp; "id": 1,

&nbsp; "name": "Alice Updated",

&nbsp; "dob": "1991-03-15"

}

List All Users



GET /users



Response:

\[

&nbsp; {

&nbsp;   "id": 1,

&nbsp;   "name": "Alice",

&nbsp;   "dob": "1990-05-10",

&nbsp;   "age": 34

&nbsp; }

]

Step 1: Install Go



Download and install Go from:

https://go.dev/dl/



Verify installation:

go version

Step 2: Clone Repository

git clone https://github.com/Prabhakar8951/go-backend-task.git

cd go-backend-task

Step 3: Install Dependencies

go mod tidy

Step 4: Run Application

go run main.go

Server will start at:

http://localhost:8080



POST http://localhost:8080/users



GET http://localhost:8080/users



PUT http://localhost:8080/users/1



