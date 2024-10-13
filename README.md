# Coeffee Trace Backend

Welcome to the **Coeffe Trace** project! This backend application serves as a bridge connecting drivers, farmers, and merchants. Built with modern tools and technologies, it provides a robust API for managing interactions between these roles. 

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [Docker Setup](#docker-setup)
- [Contributing](#contributing)
- [License](#license)

## Features

- **User Management**: Registration, login, and profile management for users (drivers, farmers, merchants).
- **Order Management**: Create, accept, and manage orders between farmers and drivers.
- **Chat Functionality**: Real-time chat support between users.
- **Forum & Resource Sharing**: Post and reply to blogs, share resources.
- **Transaction Tracking**: Monitor and manage transactions between users.

## Technologies Used

- **Golang**: The primary programming language used to build the application.
- **Gin**: A high-performance web framework for Go.
- **MongoDB**: NoSQL database for storing user data and transactions.
- **Docker**: Containerization platform to streamline deployment.
- **Swagger**: API documentation tool for interactive exploration of the API.

## Getting Started

To set up the project locally, follow these steps:

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/yourusername/coffeetrace.git
   cd coffeetrace
   ```

2. **Install Dependencies**:
   Ensure you have Go installed, then run:
   ```bash
   go mod download
   ```

3. **Set Up Environment Variables**:
   Create a `.env` file in the root directory and define your environment variables.

4. **Run the Application**:
   ```bash
   go run ./cmd/main.go
   ```

5. **Access API**: 
   Open your browser and navigate to `http://localhost:4000`.

## API Endpoints

The application provides the following main routes:

### User Routes
- `POST /user/register`: Register a new user.
- `POST /user/login`: User login.
- `GET /user/me`: Get current user's information.

### Order Routes
- `POST /order/create`: Create a new order.
- `POST /order/farmer/:id/accept`: Accept an order as a farmer.
- `GET /order/getmyorders`: Retrieve orders for the current user.

### Chat Routes
- `POST /chat/`: Create a new chat.
- `GET /chat/:id`: Retrieve all chats.

### Forum Routes
- `POST /forum/post`: Create a new forum post.
- `GET /forum/getAllBlog`: Get all forum posts.

### Product Routes
- `POST /product/create`: Add a new product.
- `GET /product/getall`: Retrieve all products.

## Docker Setup

To run the application using Docker, follow these steps:

1. **Build the Docker Image**:
   ```bash
   docker build -t coffeetrace .
   ```

2. **Run the Docker Container**:
   ```bash
   docker run -d -p 4000:4000 --env-file .env coffeetrace
   ```

3. **Access the Application**:
   Open your browser and navigate to `http://localhost:4000`.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you find any bugs or have suggestions for improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---

Thank you for your interest in **Coeffe Trace**! We hope you find this backend solution useful for your needs. Happy coding!