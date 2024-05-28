# GO Taskify - RESTful Task Management System

**Description:**

This project implements a task management system using Go and MariaDB. Users can:

- Create, view, update, and delete tasks.
- Manage user profiles.

**Features:**

- **Basic User Management:** CRUD operations for user profiles.
- **Task Management:** CRUD operations for tasks.


## Database Schema

### User Table

| Column       | Type      | Description                              |
|--------------|-----------|------------------------------------------|
| ID           | uint      | Primary key                              |
| Username     | string    | Unique username, not null                |
| Password     | string    | Password, not null                       |
| Email        | string    | Email address, required                  |
| Tasks        | []Task    | User's tasks, with foreign key to Task   |

### Task Table

| Column       | Type       | Description                              |
|--------------|------------|------------------------------------------|
| ID           | uint       | Primary key                              |
| UserID       | uint       | Foreign key to User                      |
| Title        | string     | Title of the task, required              |
| Description  | string     | Description of the task                  |
| Status       | string     | Status of the task                       |
| CreatedAt    | time.Time  | Creation timestamp                       |
| UpdatedAt    | time.Time  | Last update timestamp                    |


## API Endpoints

| Endpoint     | Method | Description                                                               |
|--------------|--------|---------------------------------------------------------------------------|
| `/tasks`     | GET    | Retrieves a list of all tasks.                                            |
| `/tasks`     | POST   | Creates a new task. Requires title, description, and status in the request body. |
| `/tasks/:id` | GET    | Retrieves details for a specific task.                                    |
| `/tasks/:id` | PUT    | Updates details for a specific task.                                      |
| `/tasks/:id` | DELETE | Deletes a specific task.                                                  |
| `/users`     | GET    | Retrieves a list of all users.                                            |
| `/users`     | POST   | Creates a new user. Requires username, email, and password in the request body. |
| `/users/:id` | GET    | Retrieves details for a specific user.                                    |
| `/users/:id` | PUT    | Updates details for a specific user.                                      |
| `/users/:id` | DELETE | Deletes a specific user.             


## Project Setup

**Backend Development:**

An initial backend development for the task management Go app has been completed.

**Forking the Repository:**

1. Visit the project repository on GitHub: [GOAPI](https://github.com/JordonAndersen/GOAPI)
2. Click the "Fork" button in the top right corner. This will create a copy of the repository in your own GitHub account.

**Cloning Your Fork:**

Once forked, clone your copy of the repository to your local machine using Git:

   ```bash
   git clone https://github.com/<your-username>/<your-GOAPI-FORK>
   cd <your-GOAPI-FORK>
   ```

**Dependencies:**

The project's dependencies are managed using Go modules (go.mod). No additional steps are required to download dependencies; they will be automatically downloaded when you run the project.

**Start the Server:**

   ```bash
   docker-compose up --build
   ```

   This will start the Go API server, connecting to the MariaDB instance.                                     |


## Improvement and Enhancements

**Tasks Remaining:**

- **Frontend Development:** Consider creating a frontend application (e.g., React, Angular) to provide a user interface for interacting with the API.
- **Deployment:** Choose a hosting platform that supports Go applications and MariaDB (e.g., Heroku, AWS, GCP) for deployment. Follow deployment instructions specific to the chosen platform.
- **Testing:** Implement unit and integration tests to ensure code functionality and database interactions are working as expected.
- **Authentication:** Enhance security by implementing token-based authentication for user access control. This can be achieved using JWT (JSON Web Tokens) or OAuth2.


## Resources

**Project Helper videos:**

- [Video 1](https://www.youtube.com/watch?v=QevhhM_QfbM)
- [Video 2](https://dev.to/divrhino/build-a-rest-api-from-scratch-with-go-and-docker-3o54)
