CREATE DATABASE timekeeper;

CREATE TABLE users (
  id SERIAL PRIMARY KEY
);

CREATE UNIQUE INDEX users_pkey ON users(id int4_ops);

CREATE TABLE pay_periods (
  id SERIAL PRIMARY KEY,
  started_on date NOT NULL,
  ended_on date NOT NULL,
  user_id integer NOT NULL REFERENCES users(id)
);

CREATE UNIQUE INDEX pay_periods_pkey ON pay_periods(id int4_ops);

CREATE TABLE projects (
  id SERIAL PRIMARY KEY,
  name text NOT NULL UNIQUE,
  code text NOT NULL UNIQUE
);

CREATE UNIQUE INDEX projects_pkey ON projects(id int4_ops);

-- work_blocks joins pay_periods and projects --
CREATE TABLE work_blocks (
  id SERIAL PRIMARY KEY,
  project_id integer NOT NULL REFERENCES projects(id),
  pay_period_id integer NOT NULL REFERENCES pay_periods(id),
  hours numeric NOT NULL,
  started_at timestamp without time zone NOT NULL,
  ended_at timestamp without time zone NOT NULL
);

CREATE UNIQUE INDEX work_blocks_pkey ON work_blocks(id int4_ops);
