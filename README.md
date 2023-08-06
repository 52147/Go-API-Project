# Go API Project

## Project Description
The "Go API Project" is a simple RESTful API written in Go (Golang) that provides user management functionalities. It allows users to interact with the API to perform CRUD (Create, Read, Update, Delete) operations on user data stored in a PostgreSQL database.


## Features
- Get all users: Retrieve a list of all users from the database.
- Get a single user by ID: Fetch detailed information about a specific user based on their unique ID.
- Create a new user: Add a new user to the database with their name and email.
- Update an existing user: Modify the details of an existing user by providing their ID and updated information.
- Delete a user: Remove a user from the database by specifying their ID.

## Technologies Used
- Go (Golang): The primary programming language used for building the API.
- PostgreSQL: The database management system used to store and manage user data.
- net/http package: Used to handle HTTP requests and responses in the API.
- github.com/gorilla/mux package: Utilized for routing and handling URL parameters.

## Installation and Usage
To use this project, you need to have Go and PostgreSQL installed on your system. After cloning the project and setting up the database connection, you can run the API using go run main.go. The API will be accessible at http://localhost:8080.

## API Endpoints
The API offers various endpoints to interact with user data. Here are the available endpoints:

- GET /users: Retrieve a list of all users.
- GET /users/{id}: Fetch detailed information about a specific user by their ID.
- POST /users: Create a new user by providing their name and email.
- PUT /users/{id}: Update an existing user's details by specifying their ID and updated information.
- DELETE /users/{id}: Delete a user from the database based on their ID.

List the available API endpoints and their descriptions. Include details such as the HTTP methods, parameters, and response formats.
| Endpoint         | Method | Description                         |
| ---------------- | ------ | ----------------------------------- |
| /users           | GET    | Get all users                       |
| /users/{id}      | GET    | Get a single user by ID             |
| /users           | POST   | Create a new user                   |
| /users/{id}      | PUT    | Update an existing user by ID       |
| /users/{id}      | DELETE | Delete a user by ID                 |
## Contributing
Contributions to this project are welcome! If you find any issues or have suggestions for enhancements, feel free to open an issue or submit a pull request. Make sure to follow the contribution guidelines mentioned in the repository.

## License
This project is released under the MIT License. You can find the full license text in the LICENSE file.

Please modify the description to accurately reflect the details and purpose of your specific project. A comprehensive README will help users and contributors understand your project better and encourage collaboration.
