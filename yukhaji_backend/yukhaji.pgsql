--
-- PostgreSQL database dump
--

-- Dumped from database version 10.10
-- Dumped by pg_dump version 10.10

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

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: savings; Type: TABLE; Schema: public; Owner: nakama
--

CREATE TABLE public.savings (
    id integer NOT NULL,
    user_id integer,
    balance integer,
    target integer,
    start_date timestamp without time zone,
    end_date timestamp without time zone
);


ALTER TABLE public.savings OWNER TO nakama;

--
-- Name: savings_id_seq; Type: SEQUENCE; Schema: public; Owner: nakama
--

CREATE SEQUENCE public.savings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.savings_id_seq OWNER TO nakama;

--
-- Name: savings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: nakama
--

ALTER SEQUENCE public.savings_id_seq OWNED BY public.savings.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: nakama
--

CREATE TABLE public.users (
    id integer NOT NULL,
    email text,
    name text
);


ALTER TABLE public.users OWNER TO nakama;

--
-- Name: savings id; Type: DEFAULT; Schema: public; Owner: nakama
--

ALTER TABLE ONLY public.savings ALTER COLUMN id SET DEFAULT nextval('public.savings_id_seq'::regclass);


--
-- Data for Name: savings; Type: TABLE DATA; Schema: public; Owner: nakama
--

COPY public.savings (id, user_id, balance, target, start_date, end_date) FROM stdin;
42	1	15000000	35000000	2019-08-31 06:40:29	2022-01-01 00:00:00
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: nakama
--

COPY public.users (id, email, name) FROM stdin;
1	admin@admin.com	admin admin
2	toped@tokopedia.com	Toped Tokopedia
3	test@test.com	tester
4	test2@toko.com	dua two
5	toko@tokped.com	toko toko
\.


--
-- Name: savings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: nakama
--

SELECT pg_catalog.setval('public.savings_id_seq', 43, true);


--
-- Name: savings savings_pkey; Type: CONSTRAINT; Schema: public; Owner: nakama
--

ALTER TABLE ONLY public.savings
    ADD CONSTRAINT savings_pkey PRIMARY KEY (id);


--
-- Name: savings savings_user_id_key; Type: CONSTRAINT; Schema: public; Owner: nakama
--

ALTER TABLE ONLY public.savings
    ADD CONSTRAINT savings_user_id_key UNIQUE (user_id);


--
-- Name: users user_email_key; Type: CONSTRAINT; Schema: public; Owner: nakama
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT user_email_key UNIQUE (email);


--
-- Name: users user_pkey; Type: CONSTRAINT; Schema: public; Owner: nakama
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- Name: savings savings_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: nakama
--

ALTER TABLE ONLY public.savings
    ADD CONSTRAINT savings_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: TABLE savings; Type: ACL; Schema: public; Owner: nakama
--

GRANT ALL ON TABLE public.savings TO postgres;


--
-- Name: SEQUENCE savings_id_seq; Type: ACL; Schema: public; Owner: nakama
--

GRANT SELECT,USAGE ON SEQUENCE public.savings_id_seq TO postgres;


--
-- Name: TABLE users; Type: ACL; Schema: public; Owner: nakama
--

GRANT ALL ON TABLE public.users TO postgres;


--
-- PostgreSQL database dump complete
--

