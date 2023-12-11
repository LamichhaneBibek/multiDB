# Multi-Database Connection App with Golang

This is a demonstration app showcasing how to establish connections to multiple databases (PostgreSQL, MySQL, SQLite) in a Go (Golang) application. The app uses a simple user model for illustration.

## Prerequisites

- Go installed on your machine
- PostgreSQL, MySQL, and SQLite databases installed and running

## Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/multi-db-connection-app.git
    ```

2. Navigate to the project directory:

    ```bash
    cd multi-db-connection-app
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

4. Set up your database configurations:

    - Copy the `config.example.yaml` file to `config.yaml`.
    - Update the database connection details in the `config.yaml` file for PostgreSQL, MySQL, and SQLite.

## Usage

Run the application:

```bash
go run main.go
```

The application will connect to the specified databases and perform basic operations using the user model.
Project Structure

    main.go: Entry point of the application.
    db: Package containing database configuration and connection logic.
    model: Package containing the user model definition.

## Database Connection

The application establishes connections to PostgreSQL, MySQL, and SQLite databases using the provided configurations in the config.yaml file. The database connections are managed in the db package.
User Model

The model package contains a simple User struct definition with fields like Username, and Password. This model is used to interact with the connected databases.
Configuration

Database configurations are stored in the config.yaml file. Update this file with your database connection details before running the application.


## Contributing

Feel free to contribute by opening issues or submitting pull requests.

## License

This project is licensed under the MIT License.
