**Project Name:** Task Management System

**Description:**

This project implements a user-friendly task management system using Go and MariaDB. Users can:

- Register, log in, and manage their profiles.
- Create, view, update, and delete tasks.
- Track task progress with statuses (e.g., pending, completed).

**Features:**

- **Basic User Management:** Registration, login, and profile updates.
- **Task Management:** CRUD (Create, Read, Update, Delete) operations for tasks.
- **Secure Authentication:** Token-based authentication for user access control.

**Project Setup:**

**Prerequisites:**

- Go (version 1.x or later) installed: [https://go.dev/doc/install](https://go.dev/doc/install)
- Git version control installed: [https://git-scm.com/downloads](https://git-scm.com/downloads)
- MariaDB database server: [https://mariadb.org/download/](https://mariadb.org/download/)

**Forking the Repository:**

1. Visit the project repository on GitHub: https://github.com/JordonAndersen/GOAPI
2. Click the "Fork" button in the top right corner. This will create a copy of the repository in your own GitHub account.

**Cloning Your Fork:**

1. Once forked, clone your copy of the repository to your local machine using Git:

   ```bash
   git clone https://github.com/<your-username>/<your-GOAPI-FORK>
   cd task-management-system
   ```

  

**Dependencies:**

The project's dependencies are managed using Go modules (go.mod). No additional steps are required to download dependencies; they will be automatically downloaded when you run the project.

**Running the Application:**

1. **Configure Database Connection:**

   - Update the `database.go` file with your MariaDB connection details (host, port, username, password, database name).

2. **Start the Server:**

   ```bash
   go run main.go
   ```

   This will start the Go API server, connecting to the MariaDB instance.

**API Endpoints:**

| Endpoint           | Method | Description                                                                          |
|-------------------|--------|---------------------------------------------------------------------------------------|
| `/register`       | POST   | Registers a new user. Requires username, email, and password in the request body.      |
| `/login`          | POST   | Logs in a user. Requires username and password in the request body. Returns an auth token. |
| `/users/{id}`      | GET    | Retrieves user profile details for the specified ID (requires authentication).          |
| `/users/{id}`      | PUT    | Updates user profile details (requires authentication).                               |
| `/tasks`           | GET    | Retrieves a list of all tasks for the authenticated user.                               |
| `/tasks`           | POST   | Creates a new task. Requires title, description, and status in the request body.         |
| `/tasks/{id}`      | GET    | Retrieves details for a specific task (requires authentication).                        |
| `/tasks/{id}`      | PUT    | Updates details for a specific task (requires authentication).                        |
| `/tasks/{id}`      | DELETE | Deletes a specific task (requires authentication).                                    |

**Frontend (Optional):**

This project is designed as a backend API. You can optionally create a frontend application (e.g., React, Angular) to interact with the API for a user interface.

**Testing:**

Unit and integration tests are encouraged to ensure code functionality and database interactions.

**Deployment:**

- Choose a hosting platform that supports Go applications and MariaDB (e.g., Heroku, AWS, GCP).
- Follow deployment instructions specific to the chosen platform.

**Contribution:**

We welcome pull requests for improvements, bug fixes, and new features. Please follow standard Git practices and conventions.

**License:**

This project is licensed under the (insert your preferred license, e.g., MIT, Apache). See the `LICENSE` file for details.

Project Helper videos: 
https://dev.to/divrhino/build-a-rest-api-from-scratch-with-go-and-docker-3o54
