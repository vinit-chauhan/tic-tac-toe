# Tic-Tac-Toe App

This project is a **Tic-Tac-Toe** game built with a **React Native** frontend and a **Golang** backend. The app supports both single-player and multiplayer game modes and features a clean UI with responsive gameplay.

## Table of Contents
- [Project Structure](#project-structure)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
  - [Requirements](#requirements)
  - [Installation](#installation)

## Project Structure
```
tic-tac-toe-app/ 
 │
 ├── backend/ # Golang backend code
 │ └── (all backend files)
 ├── frontend/ # React Native frontend code
 │ └── (all frontend files)
 ├── docker-compose.yml # Docker Compose setup for local dev
 ├── README.md # Main documentation
 └── .gitignore # Gitignore for both frontend and backend
```
## Technologies Used

### Backend (Golang)
- **Go**: Golang for the backend API.
- **REST API**: Communication between the backend and frontend.
- **Postgres**: If storing game history, player profiles, etc.

### Frontend (React Native)
- **React Native**: For cross-platform mobile app development.
- **React Navigation**: For navigation between app screens.
- **Axios**: For making API requests to the backend.

### Docker
- To containerize the app for consistent development and deployment.
- A multi-staged build was used for small container image sizes.

### Redis [TODO]
- To store game state for faster access

### Prometheus
- To monitor the application.

## TODOs:
- [ ] Beautify UI
- [ ] Add docker container for frontend code ( web ) 
- [ ] History of games
- [ ] Use Redis to store game state
- [ ] Use long polling instead of continuous polling
- [ ] Switch to web socket for game state events
- [ ] Create CI steps to build and release for Android and iOS

## Getting Started

### Requirements
- **Node.js**: Required to run the React Native app.
- **Go**: To run the backend server.
- **Docker**: For containerization.

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/your-username/tic-tac-toe-app.git
   cd tic-tac-toe-app
   docker compose up -d
   ```
