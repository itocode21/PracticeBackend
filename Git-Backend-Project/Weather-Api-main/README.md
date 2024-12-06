# Weather_API 
This is my project following[ https://roadmap.sh/projects/weather-api-wrapper-service ](https://roadmap.sh/projects/weather-api-wrapper-service)
- A picture below will visualize  how the project was made.
  <p align="center" width="100%" >
    <img width="100%" src="https://assets.roadmap.sh/guest/weather-api-f8i1q.png" > 
</p>

# Presequites: 
- Install Go V1.23.1
- Install Redis

# How to use it?
1  Clone project on your machine 
```
  git clone https://github.com/itocode21/Weather-Api.git
```

2  Run the Docker file to pull Redis .If you already have Redis on your local machine , you don't need do this step   
- Build and compose up Docker
```
docker compose -f "docker-compose.yml" up -d --build
```
- Make sure your Radis is running; you can configure it in redis.go file in Connection folder
3 Install Go dependencies:
```
go mod download
```
4 Run the application:
```
go run main.go
```

# Usage
Send a GET request to the following endpoint:
```
http://localhost:3000/weather/your_location 
Example:
http://localhost:3000/weather/Tokyo
```
