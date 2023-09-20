# Discord OAuth2 API

<a href='https://ko-fi.com/hackirby'><img src='https://storage.ko-fi.com/cdn/kofi3.png' width=150></a>

This repository contains a simple Discord OAuth2 API built using the Fiber web framework in the Go programming language. It also integrates with a PostgreSQL database to store user information after successful authentication through Discord. This README will guide you through setting up and using the API.

## Table of Contents
- [Discord OAuth2 API](#discord-oauth2-api)
  - [Table of Contents](#table-of-contents)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
  - [Running the Application](#running-the-application)
  - [Configuration](#configuration)
  - [Endpoints](#endpoints)
      - [Get discord OAuth2 page](#get-discord-oauth2-page)
      - [Get user access token](#get-user-access-token)
      - [Get user record in database](#get-user-record-in-database)
  - [Authentication](#authentication)
  - [Database](#database)
  - [Contributing](#contributing)
  - [License](#license)

## Getting Started
### Prerequisites
Before you begin, make sure you have the following tools and dependencies installed:

- Go (Golang): Install from [Go Downloads](https://go.dev/dl/)
- PostgreSQL: Install from [PostgreSQL Downloads](https://www.postgresql.org/download/)
- Discord Developer Application: Create one at [Discord Developer Portal](https://discord.com/developers/applications)

### Installation
1. Clone the repository to your local machine:

```
git clone https://github.com/hackirby/discord-oauth2-api-go.git
```

2. Navigate to the project directory:
```
cd discord-oauth2-api-go
```

3. Install Go dependencies:
```
go mod tidy
```

## Running the Application
```
go run ./cmd/app
```

## Configuration
Before launching the application, you must configure it. Create a .env file in the project's root directory and populate it with the following variables:

```dotenv
PORT=3333
CLIENT_ID=YOUR_DISCORD_CLIENT_ID
CLIENT_SECRET=YOUR_DISCORD_CLIENT_SECRET
REDIRECT_URI=YOUR_DISCORD_REDIRECT_URI
DATABASE_DSN="user=<USER> password=<PASSWORD> dbname=<DATABASE> sslmode=disable"
```
Ensure that you replace `YOUR_DISCORD_CLIENT_ID`, `YOUR_DISCORD_CLIENT_SECRET`, `YOUR_DISCORD_REDIRECT_URI` and `DATABASE_DSN` with your actual Discord OAuth2 credentials and PostgreSQL database details.


## Endpoints
#### Get discord OAuth2 page

```http
  GET /api/oauth2/url
```

#### Get user access token

```http
  GET /api/oauth2/token
```

| Parameter | Description                       |
| :-------- | :-------------------------------- |
| `code`      | **Required**. code from callback |

#### Get user record in database

```http
  GET /api/user/
```

| Header | Description                       |
| :-------- | :-------------------------------- |
| `Authorization`      | **Required**. user access token |


## Authentication
To authenticate users through Discord OAuth2, adhere to these steps:

- Get OAuth2 Discord page by calling /api/oauth2/url.
- Redirect the user to the page
- Upon successful authentication, the user will be redirected back to your frontend redirect URL.
- Use /api/oauth2/token to sign-in/sign-up the user and get his access token.

As an example, you can access the /api/user endpoint to retrieve their record in the Database with the access token as `Authorization` header.

## Database
User data is stored securely in a PostgreSQL database. The database schema is automatically generated upon application startup. After a successful Discord OAuth2 authentication, user data is stored in this database.

## Contributing
If you wish to contribute to this project, kindly open an issue or submit a pull request with your proposed changes. We warmly welcome contributions from the community.

## License
This project is licensed under the MIT License. For more information, please consult the LICENSE file.
