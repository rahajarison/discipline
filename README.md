# Technological stack
## Front
Language: Javascript
Framework: Vue 3 + Vite

## Back
Language: Go  
Web Framework: Gin (performance) or Echo (fine-tuning)
Database: PostgreSQL  
ORM: GORM  
Hosting: GCP Cloud Run / Kubernetes with Docker  
Background tasks: Goroutines  
Unit Testing: Testify  
Monitoring: Prometheus + Grafana  

Conclusion: Why choose PostgreSQL?
Complex transaction management: PostgreSQL excels in managing complex transactions and relationships between multiple entities in a robust and reliable way.
Advanced data type support: If you have advanced needs like geospatial data management, JSON documents or custom types, PostgreSQL is an excellent choice.
Performance for heavy loads: It is particularly well-suited for applications with concurrent transactions and high concurrency.
Extensions and tools ecosystem: PostgreSQL has a rich ecosystem of extensions that can easily extend its capabilities.

---- 
Secrets management: For a production environment, it's recommended to use a secrets manager to avoid exposing passwords in docker-compose.yml files. This can be done using Docker secrets or tools like HashiCorp Vault.

Persisting data: The postgres_data volume allows data persistence between restarts. Make sure critical data is stored in persistent volumes to prevent data loss.

Logs and monitoring: You can also configure tools like Prometheus or Grafana to monitor your Docker services. For critical applications, you might want to consider integrating performance monitoring tools.

## Frontend
### Intro
This template should help get you started developing with Vue 3 in Vite. The template uses Vue 3 `<script setup>` SFCs, check out the [script setup docs](https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup) to learn more.

Learn more about IDE Support for Vue in the [Vue Docs Scaling up Guide](https://vuejs.org/guide/scaling-up/tooling.html#ide-support).


### Commands
`npm install`  
`npm run dev`

For production build: `npm run build`

## Backend
### Structure
- `/cmd` : Application entry points
- `/internal` : Business logic
- `/pkg` : Reusable libraries
- `/api` : API configuration and handlers
- `/config` : Configuration files
- `/migrations` : DB migration scripts
- `/test` : Unit tests

### Commands
To get an interactive Go shell
`docker-compose exec backend /bin/bash`
`docker run -it --rm -v $(pwd):/app golang:1.18-alpine /bin/sh`

Start development with Docker Compose
Build the image and start the container: Open your terminal in the directory containing the docker-compose.yml file and run the following commands:

`docker-compose up --build`
This will:

Build the image from the Dockerfile.
Start both your Go backend and PostgreSQL in separate containers.
Note: If you don't want to rebuild the image every time, you can use `docker-compose up` without --build.

Access the application: Once the container is running, you can access your Go application by visiting http://localhost:8080.

docker-compose exec backend go run main.go
Stop and restart the container:
You can stop and restart services without losing your local changes with:

`docker-compose down`
`docker-compose up`

#### Dependencies
`go get -u gorm.io/gorm`
`go get -u gorm.io/driver/postgres`

### Differences between dockerfiles

##### Go mod / Go sum
`COPY go.mod go.sum ./`: This command copies only the go.mod and go.sum files into the container. This allows Docker to cache these files during image building. In other words, only Go dependencies are downloaded into the image, and this process is cached.
`RUN go mod download`: This command downloads all Go dependencies defined in go.mod and go.sum without touching the project source code. This means dependencies are only downloaded when go.mod or go.sum files change.

Benefits:
This improves performance by caching Go dependency downloads. This means if you don't modify your dependencies, Docker won't need to download dependencies every time you change the source code. This is particularly useful during local development, where you typically only modify your code without touching dependencies.

##### Go mod tidy
`COPY . .`: This command copies the entire working directory (source code and configuration files) into the container. This means you're copying your entire application, including source code.
`RUN go mod tidy && go build -o main .`: This command does two things:
`go mod tidy`: This command cleans up the go.mod file by adding missing dependencies and removing unused ones. It's useful for keeping your go.mod file clean.
`go build -o main .`: This command compiles the Go project into an executable and places it in a file named main.

Benefits:
Less caching: Every time you modify a source file (like a Go file), Docker needs to copy all files and run go mod tidy and go build. This is a simpler but less optimized approach in terms of performance if you're only modifying code without changing dependencies.
Simple for small projects: This can be suitable for simple projects or when you're developing iteratively and want everything compiled each time.

Disadvantages of version 2:
This approach can be less efficient in terms of caching. For example, Docker will need to download dependencies each time, even if go.mod and go.sum haven't changed.
