CREATE DATABASE timekeeper;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username text NOT NULL UNIQUE,
    password text NOT NULL
);

CREATE UNIQUE INDEX users_pkey ON users(id int4_ops);
CREATE UNIQUE INDEX users_username_key ON users(username text_ops);

CREATE TABLE pay_periods (
  id SERIAL PRIMARY KEY,
  started_at timestamp without time zone NOT NULL,
  ended_at timestamp without time zone NOT NULL,
  user_id integer NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT started_before_ended_check CHECK (started_at < ended_at)
);

CREATE UNIQUE INDEX pay_periods_pkey ON pay_periods(id int4_ops);

CREATE TABLE projects (
  id SERIAL PRIMARY KEY,
  name text NOT NULL UNIQUE CHECK (name <> ''::text),
  code text NOT NULL UNIQUE CHECK (code <> ''::text),
	user_id integer REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE UNIQUE INDEX projects_pkey ON projects(id int4_ops);
CREATE UNIQUE INDEX projects_name_key ON projects(name text_ops);
CREATE UNIQUE INDEX projects_code_key ON projects(code text_ops);

CREATE OR REPLACE FUNCTION ValidStart(started_at TIMESTAMP, pay_period_id INT) RETURNS BOOLEAN AS $$
	DECLARE started_at DATE;
    BEGIN
  		SELECT pay_periods.started_at INTO started_at FROM pay_periods WHERE id=$2;
        RETURN started_at <= $1;
    END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION ValidEnd(ended_at TIMESTAMP, pay_period_id INT) RETURNS BOOLEAN AS $$
	DECLARE ended_at DATE;
    BEGIN
  		SELECT pay_periods.ended_at INTO ended_at FROM pay_periods WHERE id=$2;
        RETURN $1 <= ended_at;
    END;
$$ LANGUAGE plpgsql;

-- work_blocks joins pay_periods and projects --
CREATE TABLE work_blocks (
  id SERIAL PRIMARY KEY,
	project_id integer NOT NULL REFERENCES projects(id) ON DELETE CASCADE ON UPDATE CASCADE,
	pay_period_id integer NOT NULL REFERENCES pay_periods(id) ON DELETE CASCADE ON UPDATE CASCADE,
  hours numeric NOT NULL,
  started_at timestamp without time zone NOT NULL,
  ended_at timestamp without time zone NOT NULL,
  CONSTRAINT valid_start_check CHECK (validstart(started_at, pay_period_id)),
  CONSTRAINT valid_end_check CHECK (validend(ended_at, pay_period_id)),
  CONSTRAINT valid_end_check CHECK (validend(ended_at, pay_period_id))
);

CREATE UNIQUE INDEX work_blocks_pkey ON work_blocks(id int4_ops);
