# Dating App API - Backend System Design
## Project Structure
```
.
├── cmd/                # Entry point for the application
│   └── datingapp/     
│       └── main.go    
├── internal/           # Internal application logic
│   ├── configs/       
│   ├── repositories/  
│   ├── models/        
│   ├── handlers/      
│   └── services/       
│   └── routes/     
├── deployments/        # Docker deployment files
│   ├── Dockerfile     
│   └── docker-compose.yml 
├── tests/              
├── build/            
├── .env            
├── go.mod             
├── go.sum             
├── LICENSE            
├── README.md          
```

### internal/
This directory contains almost all app logic, internal/ is being used here instead of pkg/ because the code are not meant to be shared across projects.

#### configs/
This is the database (SQL) and redis connection configuration directory.

#### repositories/
The repository layer responsible for handling data from/to the data sources such as SQL and redis.

#### models/
Contains model definition of entities. The models here are defined using gorm.

#### handlers/
The functions directly tied to routes to handle each specific endpoints.

#### services/
This is an intermediary layer between handler and repository. The business logic will take place here. 

#### routes/
The routes/routes.go file defines API endpoints and handles parameter passing with middleware. The URLs also grouped and protected from unauthenticated users.

### deployments/
The directory responsible for defining docker configs.

### tests/
Includes mocks for some handlers, services, and repositories for unit testing. More on unit testing on documentation.
## Running the Project

### Prerequisites
- Go (1.23)
- Postgresql 
- Redis

### Local Development Setup
1. Clone the repository:
    ```
    git clone https://github.com/ktsabit/datingapp.git
    cd datingapp
    ```
2. Install go dependencies:
   ```
   go mod download
   ```
3. Configure the database: Add the following credentials to the DB
   ```
   db: datingapp
   user: admin
   pass: admin123
   ```
4. Run the development server:
   ```
   go run cmd/datingapp/main.go
   ```
   or build and run the executable:
   ```
   go build -o build/main ./cmd/datingapp
   ./build/main
   ```
### Running Tests
To execute unit tests run the following command:
```
go test ./tests/...  -v
```

## Deploy the App with Docker

### Prerequisites
- Docker 
- Docker compose

To deploy with docker compose go to the `deployment/` directory and run the following command:
```
docker compose up -d
```
To rebuild and apply changes run the following command:
```
docker compose up -d --build
```






## LICENSE

This project is distributed under the MIT License. Read LICENSE for details.