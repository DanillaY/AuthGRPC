--
-- PostgreSQL database dump
--

-- Dumped from database version 15.3
-- Dumped by pg_dump version 15.3

-- Started on 2023-12-16 10:12:01

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

--
-- TOC entry 215 (class 1259 OID 16564)
-- Name: App; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."App" (
    "ID" integer SERIAL PRIMARY KEY NOT NULL,
    "Name" text NOT NULL,
    "Secret" text NOT NULL
);


ALTER TABLE public."App" OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 16554)
-- Name: User; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."User" (
    "ID" integer SERIAL PRIMARY KEY NOT NULL,
    "Email" text NOT NULL,
    "Gender" text,
    "PhoneNumber" text,
    "Password" bytea NOT NULL
);


ALTER TABLE public."User" OWNER TO postgres;

--
-- TOC entry 3179 (class 2606 OID 16570)
-- Name: App App_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."App"
    ADD CONSTRAINT "App_pkey" PRIMARY KEY ("ID");


--
-- TOC entry 3177 (class 2606 OID 16560)
-- Name: User User_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."User"
    ADD CONSTRAINT "User_pkey" PRIMARY KEY ("ID");


-- Completed on 2023-12-16 10:12:01

--
-- PostgreSQL database dump complete
--

