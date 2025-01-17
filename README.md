# Pack Calculator Project

A pack calculator application designed to calculate the optimal distribution of items into predefined pack sizes. The project includes a backend implemented in Go with a REST API and Swagger documentation, and a frontend implemented in React with TailwindCSS for styling.

## Directory Structure

### Backend
```
backend/
├── api/
│   └── handlers/
│       ├── calculate.go        # Handles the calculation API endpoint
│       └── packs.go            # Handles the pack sizes update API endpoint
├── cmd/
│   └── server/
│       └── main.go             # Application entry point
├── internal/
│   ├── calculator/
│   │   ├── calculator.go       # Core logic for pack calculations
│   │   └── calculator_test.go  # Unit tests for calculator logic
│   ├── config/
│   │   ├── config.go           # Configuration loader
│   │   └── pack_sizes.json     # Default pack sizes
│   └── models/
│       └── pack_sizes.go       # Model definitions for pack sizes
├── pkg/
│   └── utils/
│       └── utils.go            # Utility functions for validation
├── test/
│   └── integration/
│       └── api_test.go         # Integration tests for API endpoints
```

### Frontend
```
frontend/
├── public/                     # Static assets
├── src/
│   ├── components/             # React components for UI
│   ├── pages/                  # Pages for the application
│   └── App.jsx                 # Main application file
├── package.json                # Project dependencies and scripts
├── tailwind.config.js          # TailwindCSS configuration
├── Dockerfile                  # Dockerfile for containerizing the frontend
```

### Docker and Compose
```
docker-compose.yml               # Compose file for running backend and frontend together
```

---

## Running the Project

### Prerequisites
1. Ensure Docker and Docker Compose are installed on your system.
2. Install Node.js (18+) if you want to run the frontend locally.

### Running with Docker (Recommended)
To start both the backend and frontend services together:
```bash
docker-compose up --build
```

To run only one service:
#### Start Backend
```bash
docker-compose up backend
```

#### Start Frontend
```bash
docker-compose up frontend
```

### Running Locally
#### Backend
1. Navigate to the `backend` directory:
   ```bash
   cd backend
   ```
2. Install Go modules:
   ```bash
   go mod download
   ```
3. Start the backend server:
   ```bash
   go run cmd/server/main.go
   ```
4. Swagger documentation will be available at [http://localhost:8080/docs/index.html](http://localhost:8080/docs/index.html).

#### Frontend
1. Navigate to the `frontend` directory:
   ```bash
   cd frontend
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Start the development server:
   ```bash
   npm run dev
   ```
4. Access the application at [http://localhost:5173](http://localhost:5173).

---

## API Documentation
The backend includes Swagger documentation accessible at `/docs/index.html`. It provides detailed information about the available endpoints, request/response formats, and error handling.

---

## Testing
### Backend
To run unit and integration tests for the backend:
```bash
cd backend
go test ./...
