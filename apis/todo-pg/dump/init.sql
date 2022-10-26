CREATE DATABASE api_todo;

-- Connect to DB and Create a Table
\c api_todo;
CREATE TABLE IF NOT EXISTS todos (
  id serial primary key,
  title varchar,
  description text,
  done bool default FALSE
);
