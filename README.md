# Tic-Tac-Toe App

This project is a **Tic-Tac-Toe** game built with a **React Native** frontend and a **Golang** backend. The app supports both single-player and multiplayer game modes and features a clean UI with responsive gameplay.

## Table of Contents
- [Project Structure](#project-structure)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
  - [Requirements](#requirements)
  - [Installation](#installation)
  - [Running the App](#running-the-app)
  - [Backend API](#backend-api)
- [Docker Setup](#docker-setup)

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
- **REST API** or **gRPC**: Communication between the backend and frontend.
- **(Optional) Database**: If storing game history, player profiles, etc.

### Frontend (React Native)
- **React Native**: For cross-platform mobile app development.
- **React Navigation**: For navigation between app screens.
- **Axios** or **Fetch**: For making API requests to the backend.

### Docker
- **Docker**: For containerizing the app for consistent development and deployment.

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
