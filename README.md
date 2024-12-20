# Todo App

### Project Description

Todo App is a backend service for managing To-Do Lists. The service enables users to create task lists, manage them, and add or edit individual tasks.  

This project is built using the Gin framework in Go, and the user authentication functionality is decoupled into a separate gRPC service, located in the repository [`dm1tl/sso`](https://github.com/dm1tl/sso).

---

### Features

#### Authentication
- Sign-up: Register a new user (`POST /auth/sign-up`)
- Sign-in: Log in to the service (`POST /auth/sign-in`)

#### Task Lists
- Create a list: Add a new task list (`POST /api/lists/`)
- Get all lists: Retrieve all task lists (`GET /api/lists/`)
- Get a list by ID: Retrieve a specific task list by its ID (`GET /api/lists/:id`)
- Update a list: Update a task list by its ID (`PUT /api/lists/:id`)
- Delete a list: Delete a task list by its ID (`DELETE /api/lists/:id`)

#### Tasks within a List
- Add a task to a list: Add a task to a specific list (`POST /api/lists/:id/items`)
- Get all tasks in a list: Retrieve all tasks in a specific list (`GET /api/lists/:id/items`)

#### Independent Tasks
- Get a task by ID: Retrieve an independent task by its ID (`GET /api/items/:id`)
- Update a task: Update an independent task by its ID (`PUT /api/items/:id`)
- Delete a task: Delete an independent task by its ID (`DELETE /api/items/:id`)

---

### Architecture
- Authentication is implemented as a separate gRPC microservice for better scalability and modularity.