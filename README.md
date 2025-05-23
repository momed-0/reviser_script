## Reviser script

This Go script fetches your most recent accepted LeetCode submissions and stores them in a database.

## How it works

- Loads environment variables from `.env` (unless `ENV=PROD`).
- Creates a user session using credentials from environment variables.
- Checks that DB credentials are loaded and the database is reachable.
- Fetches recent accepted LeetCode submissions for the user.
- For each submission:
  - Fetches the problem description and submission code.
  - Upserts the question and inserts the submission into db.

## Database

- Uses [Supabase](https://supabase.com/) as the backend.
- Connects to Supabase using the REST API (not the client SDK).
- Requires the following environment variables:
  - `LEETCODE_USERNAME`
  - `LEETCODE_SESSION`
  - `SUPABASE_URL`
  - `SUPABASE_ANON_KEY`

## Running

1. Copy `.env.example` to `.env` and fill in your credentials.
2. Run the script:

   ```sh
   go run main.go

## File Overview

- `main.go`: Entry point; orchestrates fetching and storing submissions.
- `internal/db`: Manages interactions with the Supabase REST API.
- `internal/leetcode`: Handles requests to the LeetCode GraphQL API.
- `internal/models`: Defines data models for users and submissions.
- `internal/request`: Provides HTTP request utilities.
- `internal/validate`: Validates credentials and database connections.
