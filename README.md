# GuessingGame

## Tech Stack

- **Backend**: Go with Fiber framework
- **Database**: SQLite
- **Documentation**: Swagger
- **Deployment**: Docker & Docker Compose

## Project Structure

```
├── backend/                # Backend (Go + Fiber)
│   ├── database/           # SQLite database files   
│   ├── handlers/           # API handlers  
│   ├── middlewares/        # Middleware functions  
│   ├── models/             # Data models  
│   ├── utils/              # Utility functions (Token handling)
│   ├── main.go             # Entry point  
│   ├── go.mod, go.sum      # Go dependencies  
│   └── Dockerfile          # Backend container config  
├── frontend/               # Frontend (React + Vite)  
│   ├── src/                # Source code  
│   │   ├── components/     # UI components  
│   │   ├── assets/         # Static assets  
│   │   ├── lib/            # Utility functions  
│   ├── package.json        # Project metadata  
│   ├── Dockerfile          # Frontend container config  
│   └── vite.config.js      # Vite configuration  
├── docker-compose.yml      # Docker Compose setup  
└── README.md               # Project documentation  
```

## API Routes

- `GET /`: Test endpoint
- `GET /swagger/*`: Swagger documentation
- **User Management**:
  - `POST /users/login`: User login
  - `POST /users`: Create new user
  - `GET /users`: Get all users
  - `GET /users/:id`: Get user by ID
  - `PUT /users`: Update user (requires authentication)
  - `DELETE /users`: Delete user (requires authentication)
- **Game**:
  - `POST /guess/:guess`: Submit a guess (requires authentication)

## Getting Started

### Running the Application

1. Clone the repository:
   ```
   git clone https://github.com/b1ngoinho11/GuessingGame.git
   cd guessinggame
   ```

2. Start the application using Docker Compose:
   ```
   docker-compose up --build
   ```

3. Access the application:
   - Frontend: http://127.0.0.1:8080
   - Backend API: http://127.0.0.1:3000
   - Swagger Documentation: http://127.0.0.1:3000/swagger/
