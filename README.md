# Product Manager App

## Table of Contents

- [Getting Started](#getting-started)
- [Prerequisites](#prerequisites)
- [Usage](#usage)
- [Built With](#built-with)
- [License](#license)

## Getting Started

To get started with this project, follow these steps:

1. Clone the repository to your local machine:

2. Install the required dependencies:

   ```
   go mod tidy
   ```

3. Set up the database:

   - Create a PostgreSQL database.
   - Update the database connection details in main.go.

4. Run the database migrations:

   ```
   make up
   ```

5. Run the application:

   ```
   make run
   ```

## Prerequisites

Before running this project, make sure you have the following prerequisites installed:

- Go: [Installation Guide](https://golang.org/doc/install)

- Migrate: [Installation Guide](https://github.com/golang-migrate/migrate#installation)

- PostgreSQL: [Installation Guide](https://www.postgresql.org/download/)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
