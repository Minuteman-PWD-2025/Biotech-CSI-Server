--
-- PostgreSQL database dump
--

-- Dumped from database version 14.9 (Ubuntu 14.9-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.9 (Ubuntu 14.9-0ubuntu0.22.04.1)

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
-- Name: adminpack; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS adminpack WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION adminpack; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION adminpack IS 'administrative functions for PostgreSQL';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: accounts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.accounts (
    id integer NOT NULL,
    address character varying(320),
    password character varying(50),
    level integer DEFAULT 1 NOT NULL
);


ALTER TABLE public.accounts OWNER TO postgres;

--
-- Name: accounts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.accounts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.accounts_id_seq OWNER TO postgres;

--
-- Name: accounts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.accounts_id_seq OWNED BY public.accounts.id;


--
-- Name: checked_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.checked_items (
    person_id integer NOT NULL,
    item_id integer NOT NULL,
    id bigint NOT NULL
);


ALTER TABLE public.checked_items OWNER TO postgres;

--
-- Name: checked_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.checked_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.checked_items_id_seq OWNER TO postgres;

--
-- Name: checked_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.checked_items_id_seq OWNED BY public.checked_items.id;


--
-- Name: items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.items (
    id integer NOT NULL,
    name character varying(50)
);


ALTER TABLE public.items OWNER TO postgres;

--
-- Name: items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.items_id_seq OWNER TO postgres;

--
-- Name: items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.items_id_seq OWNED BY public.items.id;


--
-- Name: people; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.people (
    id integer NOT NULL,
    name character varying(50) NOT NULL
);


ALTER TABLE public.people OWNER TO postgres;

--
-- Name: people_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.people_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.people_id_seq OWNER TO postgres;

--
-- Name: people_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.people_id_seq OWNED BY public.people.id;


--
-- Name: tokens; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tokens (
    created date DEFAULT CURRENT_DATE,
    updated date DEFAULT CURRENT_DATE,
    expires date DEFAULT (CURRENT_DATE + '01:00:00'::interval),
    token character varying(20) NOT NULL
);


ALTER TABLE public.tokens OWNER TO postgres;

--
-- Name: accounts id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accounts ALTER COLUMN id SET DEFAULT nextval('public.accounts_id_seq'::regclass);


--
-- Name: checked_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checked_items ALTER COLUMN id SET DEFAULT nextval('public.checked_items_id_seq'::regclass);


--
-- Name: items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items ALTER COLUMN id SET DEFAULT nextval('public.items_id_seq'::regclass);


--
-- Name: people id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.people ALTER COLUMN id SET DEFAULT nextval('public.people_id_seq'::regclass);


--
-- Data for Name: accounts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.accounts (id, address, password, level) FROM stdin;
\.


--
-- Data for Name: checked_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.checked_items (person_id, item_id, id) FROM stdin;
8	30	1
8	29	2
7	28	3
7	27	4
6	26	5
6	25	6
6	24	7
5	23	8
5	22	9
5	21	10
4	20	11
4	19	12
3	18	13
3	17	14
3	16	15
3	15	16
3	14	17
3	13	18
3	12	19
2	11	20
2	10	21
1	9	22
1	8	23
1	7	24
1	6	25
6	5	26
5	4	27
5	3	28
5	2	29
5	1	30
\.


--
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.items (id, name) FROM stdin;
1	Cell Sample - Blue
2	Cell Sample - Blue
3	Cell Sample - Blue
4	Cell Sample - Blue
5	Cell Sample - Blue
6	Cell Sample - Blue
7	Cell Sample - Blue
8	Cell Sample - Blue
9	Cell Sample - Yellow
10	Cell Sample - Yellow
11	Cell Sample - Yellow
12	Cell Sample - Yellow
13	Cell Sample - Yellow
14	Cell Sample - Yellow
15	Cell Sample - Yellow
16	Cell Sample - Yellow
17	Cell Sample - Red
18	Cell Sample - Red
19	Cell Sample - Red
20	Cell Sample - Red
21	Cell Sample - Red
22	Cell Sample - Red
23	Cell Sample - Red
24	Cell Sample - Red
25	Cell Sample - Red
26	Cell Sample - Red
27	Cell Sample - Green
28	Cell Sample - Green
29	Cell Sample - Green
30	Cell Sample - Green
\.


--
-- Data for Name: people; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.people (id, name) FROM stdin;
1	Ben
2	Iain
3	Cormac
4	Andrew
5	Robert
6	Drew
7	Jun
8	Patrick
9	beingtwo
\.


--
-- Data for Name: tokens; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tokens (created, updated, expires, token) FROM stdin;
\.


--
-- Name: accounts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.accounts_id_seq', 1, false);


--
-- Name: checked_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.checked_items_id_seq', 30, true);


--
-- Name: items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.items_id_seq', 30, true);


--
-- Name: people_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.people_id_seq', 9, true);


--
-- Name: accounts accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);


--
-- Name: checked_items checked_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checked_items
    ADD CONSTRAINT checked_items_pkey PRIMARY KEY (id);


--
-- Name: items items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);


--
-- Name: people people_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.people
    ADD CONSTRAINT people_pkey PRIMARY KEY (id);


--
-- Name: tokens tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tokens
    ADD CONSTRAINT tokens_pkey PRIMARY KEY (token);


--
-- Name: checked_items checked_items_item_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checked_items
    ADD CONSTRAINT checked_items_item_id_fk FOREIGN KEY (item_id) REFERENCES public.items(id);


--
-- Name: checked_items checked_items_person_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.checked_items
    ADD CONSTRAINT checked_items_person_id_fk FOREIGN KEY (person_id) REFERENCES public.people(id);


--
-- PostgreSQL database dump complete
--

