
# CRUD Employee Go

This is a simple CRUD (Create, Read, Update, Delete) web application built with Go (Golang) that manages employee records.

## Features

- Add new employee
- View all employees
- Update employee information
- Delete employee

## Prerequisites

To run this project, you need to have the following installed:

- [Go](https://golang.org/doc/install) (version 1.16 or higher)
- [MySQL](https://dev.mysql.com/downloads/mysql/) (for database)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/tsaqiffatih/crud-employee-go.git
   ```

2. Initialize a Go module: :

   ```bash
   go mod init github.com/tsaqiffatih/crud-employee-go
   ```

3. Install the dependencies (if any):

   ```bash
   go mod tidy
   ```

4. Setup your MySQL database and run the following SQL script to create the employee table:

   ```sql
   CREATE DATABASE crud_employee_go;
   USE crud_employee_go;

   CREATE TABLE employee (
       id INT AUTO_INCREMENT PRIMARY KEY,
       name VARCHAR(100),
       npwp VARCHAR(15),
       address TEXT
   );
   ```

## Running the Application

1. Make sure your MySQL database is running and connected.

2. Run the Go application:

   ```bash
   go run main.go
   ```

3. Open your web browser and navigate to:

   ```
   http://localhost:8080
   ```

## Project Structure

- `controller/` - Contains the controllers for handling requests.
- `routes/` - Manages all the application routes.
- `database/` - Contains database connection and configurations.
- `views/` - HTML templates for the web interface.
- `main.go` - The entry point of the application.

## Routes

- `/employee` - List all employees.
- `/employee/create` - Create a new employee.
- `/employee/update` - Update employee information.
- `/employee/delete` - Delete an employee.

## Contributing

Feel free to submit pull requests or open issues if you want to contribute to the project.

## License

This project is open source and available under the [MIT License](LICENSE).
