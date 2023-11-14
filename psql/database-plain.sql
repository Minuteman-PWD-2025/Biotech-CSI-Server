--
-- PostgreSQL database dump
--

-- Dumped from database version 16.0
-- Dumped by pg_dump version 16.0

-- Started on 2023-11-14 11:16:31

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
-- TOC entry 2 (class 3079 OID 16384)
-- Name: adminpack; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS adminpack WITH SCHEMA pg_catalog;


--
-- TOC entry 4885 (class 0 OID 0)
-- Dependencies: 2
-- Name: EXTENSION adminpack; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION adminpack IS 'administrative functions for PostgreSQL';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 221 (class 1259 OID 19326)
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
-- TOC entry 220 (class 1259 OID 19325)
-- Name: accounts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.accounts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.accounts_id_seq OWNER TO postgres;

--
-- TOC entry 4886 (class 0 OID 0)
-- Dependencies: 220
-- Name: accounts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.accounts_id_seq OWNED BY public.accounts.id;


--
-- TOC entry 223 (class 1259 OID 19349)
-- Name: people_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.people_items (
    person_id integer NOT NULL,
    item_id integer NOT NULL,
    id bigint NOT NULL
);


ALTER TABLE public.people_items OWNER TO postgres;

--
-- TOC entry 224 (class 1259 OID 19363)
-- Name: checked_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.checked_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.checked_items_id_seq OWNER TO postgres;

--
-- TOC entry 4887 (class 0 OID 0)
-- Dependencies: 224
-- Name: checked_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.checked_items_id_seq OWNED BY public.people_items.id;


--
-- TOC entry 219 (class 1259 OID 19319)
-- Name: items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.items (
    id integer NOT NULL,
    name character varying(50)
);


ALTER TABLE public.items OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 19318)
-- Name: items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.items_id_seq OWNER TO postgres;

--
-- TOC entry 4888 (class 0 OID 0)
-- Dependencies: 218
-- Name: items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.items_id_seq OWNED BY public.items.id;


--
-- TOC entry 217 (class 1259 OID 19312)
-- Name: people; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.people (
    id integer NOT NULL,
    name character varying(50) NOT NULL
);


ALTER TABLE public.people OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 19311)
-- Name: people_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.people_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.people_id_seq OWNER TO postgres;

--
-- TOC entry 4889 (class 0 OID 0)
-- Dependencies: 216
-- Name: people_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.people_id_seq OWNED BY public.people.id;


--
-- TOC entry 222 (class 1259 OID 19334)
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
-- TOC entry 4710 (class 2604 OID 19374)
-- Name: accounts id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accounts ALTER COLUMN id SET DEFAULT nextval('public.accounts_id_seq'::regclass);


--
-- TOC entry 4709 (class 2604 OID 19376)
-- Name: items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items ALTER COLUMN id SET DEFAULT nextval('public.items_id_seq'::regclass);


--
-- TOC entry 4708 (class 2604 OID 19377)
-- Name: people id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.people ALTER COLUMN id SET DEFAULT nextval('public.people_id_seq'::regclass);


--
-- TOC entry 4715 (class 2604 OID 19375)
-- Name: people_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.people_items ALTER COLUMN id SET DEFAULT nextval('public.checked_items_id_seq'::regclass);


--
-- TOC entry 4876 (class 0 OID 19326)
-- Dependencies: 221
-- Data for Name: accounts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.accounts (id, address, password, level) FROM stdin;
\.


--
-- TOC entry 4874 (class 0 OID 19319)
-- Dependencies: 219
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
-- TOC entry 4872 (class 0 OID 19312)
-- Dependencies: 217
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
\.


--
-- TOC entry 4878 (class 0 OID 19349)
-- Dependencies: 223
-- Data for Name: people_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.people_items (person_id, item_id, id) FROM stdin;
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
2	30	31
\.


--
-- TOC entry 4877 (class 0 OID 19334)
-- Dependencies: 222
-- Data for Name: tokens; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tokens (created, updated, expires, token) FROM stdin;
\.


--
-- TOC entry 4890 (class 0 OID 0)
-- Dependencies: 220
-- Name: accounts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.accounts_id_seq', 1, false);


--
-- TOC entry 4891 (class 0 OID 0)
-- Dependencies: 224
-- Name: checked_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.checked_items_id_seq', 31, true);


--
-- TOC entry 4892 (class 0 OID 0)
-- Dependencies: 218
-- Name: items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.items_id_seq', 30, true);


--
-- TOC entry 4893 (class 0 OID 0)
-- Dependencies: 216
-- Name: people_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.people_id_seq', 8, true);


--
-- TOC entry 4721 (class 2606 OID 19331)
-- Name: accounts accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);


--
-- TOC entry 4719 (class 2606 OID 19324)
-- Name: items items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);


--
-- TOC entry 4725 (class 2606 OID 19369)
-- Name: people_items people_items_id_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.people_items
    ADD CONSTRAINT people_items_id_pkey PRIMARY KEY (id);


--
-- TOC entry 4717 (class 2606 OID 19317)
-- Name: people people_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.people
    ADD CONSTRAINT people_pkey PRIMARY KEY (id);


--
-- TOC entry 4723 (class 2606 OID 19348)
-- Name: tokens tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tokens
    ADD CONSTRAINT tokens_pkey PRIMARY KEY (token);


--
-- TOC entry 4726 (class 2606 OID 19357)
-- Name: people_items people_items_item_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.people_items
    ADD CONSTRAINT people_items_item_id_fk FOREIGN KEY (item_id) REFERENCES public.items(id);


--
-- TOC entry 4727 (class 2606 OID 19352)
-- Name: people_items people_items_person_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.people_items
    ADD CONSTRAINT people_items_person_id_fk FOREIGN KEY (person_id) REFERENCES public.people(id);


-- Completed on 2023-11-14 11:16:31

--
-- PostgreSQL database dump complete
--

