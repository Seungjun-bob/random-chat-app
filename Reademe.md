# Random Chat App

A simple random chat application built using Go (Golang) and WebSocket. Users are matched randomly for 1:1 chat sessions, with options to start, send messages, and end the chat.

## Features

- **1:1 Random Chat Matching**: Users are paired randomly for a chat session.
- **Real-time Messaging**: Powered by WebSocket for instant communication.
- **Mobile-Friendly UI**: Designed for both desktop and mobile devices.
- **Session Management**: Allows users to disconnect and reconnect easily.
- **UUID-based Identification**: Each user is assigned a unique ID for tracking.

## Tech Stack

- **Backend**: Go (Golang)
  - `github.com/gorilla/websocket`: WebSocket library for Go.
  - `github.com/google/uuid`: UUID generation for unique user identification.
- **Frontend**: HTML, CSS, JavaScript
  - Fully responsive layout for desktop and mobile.
- **Communication**: WebSocket for real-time communication.

## Installation

### Prerequisites

- Go (version 1.18 or later)
- A web browser (Chrome, Firefox, etc.)