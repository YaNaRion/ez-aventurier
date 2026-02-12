# Architecture
## Controller
- REST API
## Router
- Handle routing for website

## Infrastructure
- Handle DB connection
- For this projet we are using a free MongoDB hosted connection

## Gateway 
- Websocket

## Views
- Website file

# How to run this projet
## Prequisite
- Required
    - Golang
    - NodeJs and npm
- Nice to have
    - Air (Golang hot reaload)

## Install dependencies (under ./server directory)
- Server
    - go mod tidy (download dependencies for the go web server)
- Website
    - npm install

## How to run the project on dev mode
- Build website: It offer hot reloading for the website, recompile the .ts file when saving
    - npm run dev
- Build and run the server
    - With hot reaload
        - air (if you want the hot reaload for the go web server)
    - Without hot reaload
        - go run main.go (build and run the server)
