# Golang MySQL CRUD API

A simple CRUD (Create, Read, Update, Delete) API built using Golang and MySQL.

## Features

- Create, Read, Update, and Delete operations.
- Golang-based RESTful API.
- MySQL database for data storage.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Golang installed on your machine.
- MySQL server installed and running.
- Set up the database and tables.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/samridhi-sahu/CURD-API.git
   
2. Navigate to the directory, install dependencies i.e. gorilla/mux and mysql and build project.
3. Start the API server : The API will be accessible at http://localhost:4000. You can test the API using tools like Postman or curl.
   
## API Endpoints :

- POST /student: Create a new student.
- GET /student/{roll}: Retrieve a student by roll no.
- GET /students: Retrieve all students.
- PUT /student/{roll}: Update a student by roll no.
- DELETE /student/{roll}: Delete a student by roll no.
