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
3. Run the development server:
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
go test
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