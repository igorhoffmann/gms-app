--
-- PostgreSQL database dump
--

-- Dumped from database version 14.2 (Debian 14.2-1.pgdg110+1)
-- Dumped by pg_dump version 14.2 (Debian 14.2-1.pgdg110+1)

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
-- Name: info_main; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.info_main (
    id integer NOT NULL,
    first_name character varying(255) NOT NULL,
    last_name character varying(255) NOT NULL,
    middle_name character varying(255),
    relationship character varying(20) NOT NULL,
    phone character varying(255) NOT NULL,
    date_of_birth date NOT NULL,
    date_of_registry date DEFAULT now() NOT NULL
);


ALTER TABLE public.info_main OWNER TO postgres;

--
-- Name: info_main_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.info_main_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.info_main_id_seq OWNER TO postgres;

--
-- Name: info_main_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.info_main_id_seq OWNED BY public.info_main.id;


--
-- Name: instructors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.instructors (
    info_id integer NOT NULL,
    hire_date date DEFAULT now() NOT NULL,
    salary integer NOT NULL
);


ALTER TABLE public.instructors OWNER TO postgres;

--
-- Name: members; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.members (
    info_id integer NOT NULL,
    membership_id integer NOT NULL,
    expires_at timestamp without time zone NOT NULL
);


ALTER TABLE public.members OWNER TO postgres;

--
-- Name: memberships; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.memberships (
    id integer NOT NULL,
    title character varying(255) NOT NULL,
    price integer NOT NULL,
    duration integer NOT NULL,
    instructor_id integer
);


ALTER TABLE public.memberships OWNER TO postgres;

--
-- Name: memberships_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.memberships_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.memberships_id_seq OWNER TO postgres;

--
-- Name: memberships_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.memberships_id_seq OWNED BY public.memberships.id;


--
-- Name: system_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.system_users (
    id integer NOT NULL,
    username character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    password_hash character varying(255) NOT NULL
);


ALTER TABLE public.system_users OWNER TO postgres;

--
-- Name: system_users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.system_users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.system_users_id_seq OWNER TO postgres;

--
-- Name: system_users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.system_users_id_seq OWNED BY public.system_users.id;


--
-- Name: visits; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.visits (
    id integer NOT NULL,
    visitor_id integer NOT NULL,
    came_at timestamp without time zone DEFAULT now() NOT NULL,
    left_at timestamp without time zone
);


ALTER TABLE public.visits OWNER TO postgres;

--
-- Name: visits_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.visits_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.visits_id_seq OWNER TO postgres;

--
-- Name: visits_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.visits_id_seq OWNED BY public.visits.id;


--
-- Name: info_main id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.info_main ALTER COLUMN id SET DEFAULT nextval('public.info_main_id_seq'::regclass);


--
-- Name: memberships id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.memberships ALTER COLUMN id SET DEFAULT nextval('public.memberships_id_seq'::regclass);


--
-- Name: system_users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.system_users ALTER COLUMN id SET DEFAULT nextval('public.system_users_id_seq'::regclass);


--
-- Name: visits id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.visits ALTER COLUMN id SET DEFAULT nextval('public.visits_id_seq'::regclass);


-- --
-- -- Data for Name: info_main; Type: TABLE DATA; Schema: public; Owner: postgres
-- --

-- COPY public.info_main (id, first_name, last_name, middle_name, relationship, phone, date_of_birth, date_of_registry) FROM stdin;
-- 3	Semen	Smirnov	Valeriyevich	instructor	+380957823211	1995-03-05	2022-03-05
-- 1	Vasya	Pupkin	Petrovich	member	+380983423738	1999-03-05	2022-03-05
-- 2	Petya	Pupkin	Petrovich	member	+380983429345	1998-02-04	2022-03-05
-- 4	Evgeniy	Shashkin	\N	user	+380932314771	2001-12-08	2022-03-06
-- 5	Nikita	Pasichnik	Vasilievich	member	+380679614369	1993-01-28	2022-05-03
-- 6	Dmitry	Bobrov	Leonidovich	member	+380505875952	2003-12-08	2022-05-03
-- 7	Gleb	Tolochko	Vadimovich	instructor	+380975384955	1978-12-12	2022-05-03
-- 8	Stepan	Osipov	Andreevich	member	+380662223148	1997-04-30	2022-05-03
-- 9	Kristina	Tsvetkova	Romanovna	user	+380987818823	1988-03-15	2022-05-03
-- 10	Eva	Kulakova	Stanislavovna	member	+380508312425	1974-11-02	2022-05-03
-- \.


-- --
-- -- Data for Name: instructors; Type: TABLE DATA; Schema: public; Owner: postgres
-- --

-- COPY public.instructors (info_id, hire_date, salary) FROM stdin;
-- 3	2022-03-05	15000
-- 7	2022-05-03	20000
-- \.


-- --
-- -- Data for Name: members; Type: TABLE DATA; Schema: public; Owner: postgres
-- --

-- COPY public.members (info_id, membership_id, expires_at) FROM stdin;
-- 1	1	2022-03-08 00:00:00
-- 2	5	2022-03-11 00:00:00
-- 5	3	2023-03-05 00:00:00
-- 6	7	2022-03-06 00:00:00
-- 8	2	2022-03-11 00:00:00
-- 10	4	2022-03-08 00:00:00
-- \.


-- --
-- -- Data for Name: memberships; Type: TABLE DATA; Schema: public; Owner: postgres
-- --

-- COPY public.memberships (id, title, price, duration, instructor_id) FROM stdin;
-- 1	Basic 3 month	2999	3	\N
-- 2	Basic 6 month	5999	6	\N
-- 3	Basic 12 month	8999	12	\N
-- 4	Advanced 3 month	3599	3	3
-- 5	Advanced 6 month	7599	6	3
-- 6	Advanced 12 month	11999	12	7
-- 7	Promo 1 month	899	1	\N
-- \.


-- --
-- -- Data for Name: system_users; Type: TABLE DATA; Schema: public; Owner: postgres
-- --

-- COPY public.system_users (id, username, email, password_hash) FROM stdin;
-- 1	Romashka	jenya_shash21@gmail.com	 05e232f70dddfc2f3163220ab688f96ceb919239
-- 2	Berezka	tsvet_proz11@gmail.com	6b8b570c56e29d5bd77edfacb7ae807072b46b76
-- \.


-- --
-- -- Data for Name: visits; Type: TABLE DATA; Schema: public; Owner: postgres
-- --

-- COPY public.visits (id, visitor_id, came_at, left_at) FROM stdin;
-- 1	2	2022-03-06 15:55:52.495514	2022-03-06 16:09:45.625473
-- 2	3	2022-05-03 08:38:43.187857	\N
-- 4	5	2022-05-03 09:26:27.321932	\N
-- 6	10	2022-05-03 09:26:27.321932	\N
-- 9	2	2022-05-03 09:26:27.321932	\N
-- 3	1	2022-05-03 08:38:43.187857	2022-05-03 09:31:48.89969
-- 5	6	2022-05-03 09:26:27.321932	2022-05-03 09:31:48.89969
-- 7	7	2022-05-03 09:26:27.321932	2022-05-03 09:31:48.89969
-- 8	9	2022-05-03 09:26:27.321932	2022-05-03 09:31:48.89969
-- 10	8	2022-05-03 09:26:27.321932	2022-05-03 09:31:48.89969
-- \.


--
-- Name: info_main_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.info_main_id_seq', 10, true);


--
-- Name: memberships_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.memberships_id_seq', 7, true);


--
-- Name: system_users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.system_users_id_seq', 2, true);


--
-- Name: visits_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.visits_id_seq', 10, true);


--
-- Name: info_main info_main_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.info_main
    ADD CONSTRAINT info_main_id_key UNIQUE (id);


--
-- Name: info_main info_main_phone_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.info_main
    ADD CONSTRAINT info_main_phone_key UNIQUE (phone);


--
-- Name: memberships memberships_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.memberships
    ADD CONSTRAINT memberships_id_key UNIQUE (id);


--
-- Name: system_users system_users_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.system_users
    ADD CONSTRAINT system_users_id_key UNIQUE (id);


--
-- Name: system_users users_login_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.system_users
    ADD CONSTRAINT users_login_key UNIQUE (username);


--
-- Name: visits visits_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.visits
    ADD CONSTRAINT visits_id_key UNIQUE (id);


--
-- Name: instructors instructors_info_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.instructors
    ADD CONSTRAINT instructors_info_id_fkey FOREIGN KEY (info_id) REFERENCES public.info_main(id);


--
-- Name: members members_info_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.members
    ADD CONSTRAINT members_info_id_fkey FOREIGN KEY (info_id) REFERENCES public.info_main(id);


--
-- Name: members members_membership_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.members
    ADD CONSTRAINT members_membership_id_fkey FOREIGN KEY (membership_id) REFERENCES public.memberships(id);


--
-- Name: memberships memberships_instructor_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.memberships
    ADD CONSTRAINT memberships_instructor_id_fkey FOREIGN KEY (instructor_id) REFERENCES public.info_main(id);


--
-- Name: visits visits_person_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.visits
    ADD CONSTRAINT visits_person_id_fkey FOREIGN KEY (visitor_id) REFERENCES public.info_main(id);


--
-- PostgreSQL database dump complete
--

