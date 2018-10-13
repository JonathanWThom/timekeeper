CREATE DATABASE timekeeper;

CREATE OR REPLACE FUNCTION ValidStart(started_at TIMESTAMP, pay_period_id INT) RETURNS BOOLEAN AS $$
	DECLARE started_on DATE;
    BEGIN
  		SELECT pay_periods.started_on INTO started_on FROM pay_periods WHERE id=$2;
        RETURN $1 > started_on;
    END;
$$ LANGUAGE plpgsql;

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
  name text NOT NULL UNIQUE CHECK (name <> ''::text),
  code text NOT NULL UNIQUE CHECK (code <> ''::text)
);

CREATE UNIQUE INDEX projects_pkey ON projects(id int4_ops);

-- work_blocks joins pay_periods and projects --
CREATE TABLE work_blocks (
  id SERIAL PRIMARY KEY,
  project_id integer NOT NULL REFERENCES projects(id),
  pay_period_id integer NOT NULL REFERENCES pay_periods(id),
  hours numeric NOT NULL,
  started_at timestamp without time zone NOT NULL,
  ended_at timestamp without time zone NOT NULL,
  CONSTRAINT valid_start_check CHECK (validstart(started_at, pay_period_id))
);

CREATE UNIQUE INDEX work_blocks_pkey ON work_blocks(id int4_ops);
