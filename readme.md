# Leaf

## Set up

- gqlgen : Graphql first type safe
- goose : db migrations
- go fiber : fastHTTP web requests
- pgx v5 : postgresql driver pgx connection pool
- sqlc : database queries type safe
- dotenv : defining environment variables
- validator v10 : input validation and custom messages
- zerolog : logger for events


## Steps to setup

### PostgreSQL Database

- Install PostgreSQL 17 and create a database.
- Set up the dsn in the enviroment variables in .env file.
- Start the PostgreSQL service or you can you below command to start from command line
  
  pg_ctl -D "Dir to pgqlc data" start

- Use below goose command and run the database migrations
  
  goose -dir ./db/migrations postgres "postgres://postgres:password@localhost:5432/db_leaf?sslmode=disable" up

- To undo the migration one step use below command
  
  goose -dir ./db/migrations postgres "postgres://postgres:password@localhost:5432/db_leaf?sslmode=disable" down

- once migration is completed. You can run
  
  go run server.go

### SQLC Install

- Download the sqlc binary to create type safe database queries
- The sqlc.yml file contains all configuration on path for schema and queries for generation.
- Define the schema in db/schema
- Define the queries in db/queries
- For this project the models, queries and querier are generated under path db/generated
- Use below command if the generated files does not exist.
  
  sqlc generate


### gqlgen command

- Project already as gqlgen installed to generate the type safe GraphQL first queries, mutations, model, resolvers, inputs, etc.
- The gqlgen.yml file contains all the congfiguration for models, resolvers.
- Use the below command you are setting adding new graphql and model.
  
  go run github.com/99designs/gqlgen generate

- The generator reads all the graphl schemas under directory graph/schema/*.graphqls
  
  example user.graphqls

- the models are generated under graph/model/models_gen.go and once its generated you can separate common models into separate model files 
  
  example user.go for all related to users

- the resolvers are generated under graph/resolvers/{name}.resolvers.go
  
  example user.resolvers.go



