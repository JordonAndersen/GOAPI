# GO Task Manager - Get Organized, Get Things Done!

## Description

This project is an implementation of a task management system utilizing Go and MariaDB. It provides functionalities for managing user profiles and tasks, with a frontend developed in React. Users can seamlessly perform CRUD operations, including creating, reading, updating, and deleting tasks and user profiles through the React-based interface.

**Go Programming Language:**
Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.

[golang.org](https://golang.org)


## Database Schema

### User Table

| Column   | Type     | Description                            |
|----------|----------|----------------------------------------|
| ID       | uint     | Primary key                            |
| Username | string   | Unique username, not null              |
| Password | string   | Password, not null                     |
| Email    | string   | Email address, required                |
| Tasks    | []Task   | User's tasks, with foreign key to Task |

### Task Table

| Column     | Type       | Description                            |
|------------|------------|----------------------------------------|
| ID         | uint       | Primary key                            |
| UserID     | uint       | Foreign key to User                    |
| Title      | string     | Title of the task, required            |
| Description| string     | Description of the task                |
| Status     | string     | Status of the task                     |
| CreatedAt  | time.Time  | Creation timestamp                     |
| UpdatedAt  | time.Time  | Last update timestamp                  |

## API Endpoints

| Endpoint     | Method | Description                               |
|--------------|--------|-------------------------------------------|
| `/tasks`     | GET    | Retrieves a list of all tasks.            |
| `/tasks`     | POST   | Creates a new task.                       |
| `/tasks/:id` | GET    | Retrieves details for a specific task.    |
| `/tasks/:id` | PUT    | Updates details for a specific task.      |
| `/tasks/:id` | DELETE | Deletes a specific task.                  |
| `/users`     | GET    | Retrieves a list of all users.            |
| `/users`     | POST   | Creates a new user.                       |
| `/users/:id` | GET    | Retrieves details for a specific user.    |
| `/users/:id` | PUT    | Updates details for a specific user.      |
| `/users/:id` | DELETE | Deletes a specific user.                  |

## Project Setup

1. **Forking the Repository:**
   - Visit the project repository on GitHub: [GOAPI](https://github.com/JordonAndersen/GOAPI)
   - Click the "Fork" button in the top right corner to create a copy in your own GitHub account.

2. **Backend (Go API):**
   - **Navigate to the Backend Directory:**
     ```bash
     cd backend
     ```
   - **Start the Server:**
     ```bash
     docker-compose up --build
     ```

3. **Frontend (React):**
   - **Navigate to the Frontend Directory:**
     ```bash
     cd frontend
     ```
   - **Install Dependencies:**
     ```bash
     npm install
     ```
   - **Start the Development Server:**
     ```bash
     npm run dev
     ```

## Project Status

**Completed:**
- Backend development: Go REST API for Task Management using MariaDB and Docker.

**In Progress:**
- Frontend development: React application for user interface.

**Remaining:**
- Deployment: Choose a hosting platform and deploy the application.
- Testing: Implement unit and integration tests.
- Authentication: Implement token-based authentication.

## Resources

**Project Helper Videos:**
- [Build a REST API from scratch with Go, Docker & Postgres](https://dev.to/divrhino/build-a-rest-api-from-scratch-with-go-and-docker-3o54)
- [Build a Go REST API, React.js & TypeScript Todo App](https://www.youtube.com/watch?v=QevhhM_QfbM)
