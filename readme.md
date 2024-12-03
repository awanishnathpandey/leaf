# Leaf

A modern backend project that utilizes a range of tools and libraries for type-safe database queries, GraphQL API, migration management, and more.

## Key Technologies

- gqlgen: GraphQL-first type-safe server.
- Goose: Database migrations management.
- Go Fiber: High-performance web framework using fastHTTP.
- Go Mail: Mail service integration, supports multiple receipients with attachments
- JWT v5: Application authentication + custom OAuth token validation
- pgx v5: PostgreSQL driver with connection pooling.
- sqlc: Type-safe database query generation.
- dotenv: Manages environment variables.
- Validator v10: Input validation with custom messages.
- zerolog: Efficient logging for events and errors.

## Project Setup

Follow the steps below to set up your development environment and get started with the Leaf project.

### 1. Set up PostgreSQL Database

1. Install PostgreSQL (version 17 recommended). You can download it from the official PostgreSQL website (https://www.postgresql.org/download/).
2. Create a new database by running the following commands in your PostgreSQL CLI or a database management tool:
```sql
CREATE DATABASE db_leaf;
```
3. Configure your environment variables: Add the PostgreSQL connection string (DSN) to your `.env` file:
```env
DATABASE_URL=postgres://postgres:password@localhost:5432/db_leaf?sslmode=disable
```
4. Start PostgreSQL: You can start the PostgreSQL service manually by running:
```cmd
pg_ctl -D "/path/to/your/pg_data" start
```
5. Create a migration: To create a new migration, such as a `users` table, run:
```sql
goose -dir ./db/migrations create create_users_table sql
```
6. Apply migrations: Run the following command to apply migrations to your PostgreSQL database:
```cmd
goose -dir ./db/migrations postgres "postgres://postgres:password@localhost:5432/db_leaf?sslmode=disable" up
```
7. Rollback a migration: If you want to undo the last migration, use:
```cmd
goose -dir ./db/migrations postgres "postgres://postgres:password@localhost:5432/db_leaf?sslmode=disable" down
```
8. Reset all migrations: To reset and rollback all migrations, run:
```cmd
goose -dir ./db/migrations postgres "postgres://postgres:password@localhost:5432/db_leaf?sslmode=disable" reset
```
9. Start the application: Once migrations are applied, run the application:


---

### 2. Install SQLC

SQLC is used to generate type-safe database queries. Follow these steps to set it up:

1. Install SQLC: Download and install SQLC to generate database queries:

2. Configure SQLC: The `sqlc.yml` file contains configuration for your database schema and queries.

3. Define the database schema: Place your SQL schema files in the `db/schema` directory.

4. Define your queries: Write SQL queries in the `db/queries` directory.

5. Generate models and queries: Run the following command to generate type-safe database code:
```cmd
sqlc generate
```

The generated models, queries, and queriers will be placed in the `db/generated` directory.

---

### 3. Set up GraphQL with gqlgen

`gqlgen` is used to generate a type-safe GraphQL API. It provides tools for automatic generation of resolvers, types, and models.

1. Install gqlgen: The project already has `gqlgen` installed. It is used to generate GraphQL queries, mutations, models, resolvers, inputs, and more.

2. Configure gqlgen: The `gqlgen.yml` file contains configuration for the GraphQL models and resolvers.

3. Generate GraphQL code: To generate the types, resolvers, and models, run:
```cmd
go run github.com/99designs/gqlgen generate
```

4. GraphQL Schemas: All GraphQL schemas are located in the `graph/schema/*.graphqls` directory. Example: `user.graphqls`. The cursor pagination is added to the existing models along with sort and filter types for each models.

5. Generated Models: The models will be generated in `graph/model/models_gen.go`. You can separate them into individual files for better organization (e.g., `user.go` for user-related models).

6. Generated Resolvers: The resolvers are created in `graph/resolvers/{name}.resolvers.go`. Example: `user.resolvers.go`.

---

## Additional Notes

- Environment Variables: You can use a `.env` file to store and manage your environment variables securely.
- Custom Validation: Use `validator.v10` for custom validation of user input in GraphQL mutations and queries.
- Logging: The `zerolog` library is used for efficient logging of events and errors throughout the application.
