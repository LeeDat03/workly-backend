-- +goose Up
-- +goose StatementBegin

CREATE TYPE status_enum AS ENUM ('PENDING', 'ACCEPT', 'REJECT');
CREATE TYPE owner_type_enum AS ENUM ('USER', 'COMPANY', 'VISITOR');
CREATE TYPE degree_enum AS ENUM ('ASSOCIATE', 'BACHELOR', 'MASTER', 'DOCTOR');


CREATE TABLE IF NOT EXISTS industries (
  id BIGSERIAL PRIMARY KEY,
  industry VARCHAR(200) NOT NULL
);


CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL PRIMARY KEY,
  industry_id BIGINT NOT NULL,
  name VARCHAR(200) NOT NULL DEFAULT '',
  email VARCHAR(200) NOT NULL DEFAULT '',
  phone VARCHAR(50) NOT NULL DEFAULT '',
  avatar VARCHAR(100) DEFAULT NULL,
  cover_photo VARCHAR(100) DEFAULT NULL,
  role SMALLINT NOT NULL,
  about VARCHAR(2000) DEFAULT NULL,
  created_at TIMESTAMP(6) WITHOUT TIME ZONE NOT NULL,
  modified_at TIMESTAMP(6) WITHOUT TIME ZONE DEFAULT NULL,
  CONSTRAINT FK_users_industries FOREIGN KEY (industry_id) REFERENCES industries (id) ON UPDATE CASCADE
);


CREATE TABLE IF NOT EXISTS companies (
  id BIGSERIAL PRIMARY KEY,
  industry_id BIGINT NOT NULL,
  owner_id BIGINT NOT NULL DEFAULT 0,
  name VARCHAR(100) NOT NULL DEFAULT '',
  overview TEXT NOT NULL,
  website_url TEXT NOT NULL,
  founded_year DATE NOT NULL,
  logo_url VARCHAR(200) NOT NULL DEFAULT '',
  created_at TIMESTAMP(6) WITHOUT TIME ZONE NOT NULL,
  modified_at TIMESTAMP(6) WITHOUT TIME ZONE DEFAULT NULL,
  created_by BIGINT DEFAULT NULL,
  modified_by BIGINT DEFAULT NULL,
  CONSTRAINT FK_companies_industries FOREIGN KEY (industry_id) REFERENCES industries (id) ON UPDATE CASCADE,
  CONSTRAINT FK_companies_users FOREIGN KEY (owner_id) REFERENCES users (id) ON UPDATE CASCADE
);


CREATE TABLE IF NOT EXISTS company_admin (
  user_id BIGINT NOT NULL,
  company_id BIGINT NOT NULL,
  status status_enum NOT NULL /* 'PENDING|ACCEPT|REJECT' */,
  PRIMARY KEY (user_id, company_id)
);

CREATE TABLE IF NOT EXISTS locations (
  id BIGSERIAL PRIMARY KEY,
  owner_type owner_type_enum NOT NULL /* 'USER|COMPANY|VISITOR' */,
  owner_id BIGINT DEFAULT NULL,
  street_address VARCHAR(100) NOT NULL,
  city VARCHAR(100) NOT NULL,
  district VARCHAR(200) NOT NULL,
  UNIQUE (owner_type, owner_id)
);


CREATE TABLE IF NOT EXISTS skills (
  id BIGINT PRIMARY KEY DEFAULT 0,
  name TEXT NOT NULL
);


CREATE TABLE IF NOT EXISTS user_educations (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL DEFAULT 0,
  school BIGINT NOT NULL,
  degree degree_enum NOT NULL /* 'ASSOCIATE|BACHELOR|MASTER|DOCTOR' */,
  major VARCHAR(50) NOT NULL DEFAULT '0',
  description TEXT NOT NULL,
  start_date DATE NOT NULL,
  end_date DATE DEFAULT NULL,
  created_at TIMESTAMP(6) WITHOUT TIME ZONE NOT NULL,
  modified_at TIMESTAMP(6) WITHOUT TIME ZONE DEFAULT NULL,
  CONSTRAINT FK_user_educations_users FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE
);


CREATE TABLE IF NOT EXISTS user_skills (
  user_id BIGINT NOT NULL,
  skill_id BIGINT NOT NULL,
  PRIMARY KEY (user_id, skill_id)
);


CREATE TABLE IF NOT EXISTS user_work_experience (
  id BIGSERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  industry_id BIGINT NOT NULL,
  company_id BIGINT DEFAULT NULL,
  company_name VARCHAR(200) DEFAULT NULL,
  employee_type SMALLINT DEFAULT NULL,
  start_date DATE NOT NULL,
  end_date DATE DEFAULT NULL,
  user_id BIGINT NOT NULL,
  location_id BIGINT NOT NULL,
  description TEXT NOT NULL,
  created_at TIMESTAMP(6) WITHOUT TIME ZONE NOT NULL,
  modified_at TIMESTAMP(6) WITHOUT TIME ZONE NOT NULL,
  UNIQUE (company_id, user_id),
  CONSTRAINT FK__industries FOREIGN KEY (industry_id) REFERENCES industries (id) ON UPDATE CASCADE,
  CONSTRAINT FK__users FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE,
  CONSTRAINT FK_user_work_experience_companies FOREIGN KEY (company_id) REFERENCES companies (id) ON UPDATE CASCADE
);


CREATE TABLE IF NOT EXISTS user_work_experience_skill (
  work_exp_id BIGINT NOT NULL,
  skill_id BIGINT NOT NULL,
  PRIMARY KEY (work_exp_id, skill_id)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS user_work_experience_skill;
DROP TABLE IF EXISTS user_work_experience;
DROP TABLE IF EXISTS user_skills;
DROP TABLE IF EXISTS user_educations;
DROP TABLE IF EXISTS skills;
DROP TABLE IF EXISTS locations;
DROP TABLE IF EXISTS company_admin;
DROP TABLE IF EXISTS companies;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS industries;
DROP TYPE IF EXISTS status_enum;
DROP TYPE IF EXISTS owner_type_enum;
DROP TYPE IF EXISTS degree_enum;

-- +goose StatementEnd