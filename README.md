# go_vk

Quest is a project for managing tasks and user history.

## Installation

Clone the repository:

```bash
git clone https://github.com/geejjoo/go_vk.git
```

Navigate to the cloned directory:
```bash
cd ./app
```

Install dependencies:
```bash
go mod tidy
```

## Usage
### Running the Project
To run the project, execute the following command in your terminal:
```bash
docker compose up
```
By default, the project will be available at `http://localhost:8080`.


## Project Structure
- `cmd/`: Directory for main executable files.
  - `db/`: Directory with model for app
  - `service/`:Directory for main executable files.
- `internal/`: Directory for packages used in the application.
  - `handler/`: Package containing HTTP request handlers.
  - `repository/`: Package containing the implementation of database operations.
  - `service/`: Package containing the business logic of the application.
- `configs/`: Directory for configuration files.
- `migration/`: Directory for migration
- `schema/`: Directory for database initialization.


 # API Documentation

API is somewhat REST-like. All endpoints are at `/api/v1/`. API methods accept data and return results in JSON object.
base url http://localhost:8080/api/v1/

## Currently Implemented API Methods:

1. **Create Quest**
   - **URL:** `/api/v1/quest/create`
   - **Method:** `POST`
   - **Description:** Creates a new quest.
   - **Request Body:**
     ```json
     {
       "name": "string",
       "cost": int,
       "time": "2022-10-31T09:00:00Z"
     }
     ```
     - The `time` field should be in RFC 3339 format.
   - **Response:** Returns the ID of the created quest.

2. **Create User**
   - **URL:** `/api/v1/user/create`
   - **Method:** `POST`
   - **Description:** Creates a new user.
   - **Request Body:**
     ```json
     {
       "name": "string"
     }
     ```
   - **Response:** Returns the ID of the created user if successful.

3. **Get User Information**
   - **URL:** `/api/v1/user/info/<id>`
   - **Method:** `GET`
   - **Description:** Retrieves information about a user including their balance and completed quests.

4. **User Quest Assignment**
   - **URL:** `/api/v1/user/quest`
   - **Method:** `POST`
   - **Description:** Signals that a user has taken or completed a quest.
   - **Request Body:**
     ```json
     {
       "userId": int,
       "questId": int
     }
     ```
     - If the user has already completed the quest, returns an error.

Also have migration for data base. You can use that:
Navigate to the migration directory:
```bash
cd ./app/migration
```

Than use
```bash
psql -U postgres -d postgres -f migrate.sql
```
If you have error, you can connect to the database and insert migrations directly into the database
