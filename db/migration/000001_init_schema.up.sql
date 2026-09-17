CREATE TABLE "departments" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(100) NOT NULL UNIQUE,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "students" (
  "id" SERIAL PRIMARY KEY,
  "first_name" VARCHAR(50) NOT NULL,
  "last_name" VARCHAR(50) NOT NULL,
  "email" VARCHAR(100) UNIQUE NOT NULL,
  "department_id" INT REFERENCES "departments"("id") ON DELETE SET NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);

CREATE TABLE "professors" (
  "id" SERIAL PRIMARY KEY,
  "first_name" VARCHAR(50) NOT NULL,
  "last_name" VARCHAR(50) NOT NULL,
  "email" VARCHAR(100) UNIQUE NOT NULL,
  "department_id" INT REFERENCES "departments"("id") ON DELETE SET NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
);