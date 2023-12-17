
SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;


CREATE TABLE public."app" (
    "ID" SERIAL PRIMARY KEY NOT NULL,
    "Name" text NOT NULL,
    "Secret" text NOT NULL
);


ALTER TABLE public."app" OWNER TO postgres;

CREATE TABLE public."user" (
    "ID" SERIAL PRIMARY KEY NOT NULL,
    "Email" text NOT NULL UNIQUE,
    "Gender" text,
    "PhoneNumber" text,
    "Password" bytea NOT NULL
);

ALTER TABLE public."user" OWNER TO postgres;