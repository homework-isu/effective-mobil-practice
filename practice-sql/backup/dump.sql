--
-- PostgreSQL database dump
--

-- Dumped from database version 15.4 (Debian 15.4-1.pgdg120+1)
-- Dumped by pg_dump version 15.4 (Debian 15.4-1.pgdg120+1)

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

ALTER TABLE ONLY public.employees_projects DROP CONSTRAINT employees_projects_project_id_fkey;
ALTER TABLE ONLY public.employees_projects DROP CONSTRAINT employees_projects_employee_id_fkey;
ALTER TABLE ONLY public.employees DROP CONSTRAINT employees_department_id_fkey;
ALTER TABLE ONLY public.projects DROP CONSTRAINT projects_pkey;
ALTER TABLE ONLY public.employees_projects DROP CONSTRAINT employees_projects_pkey;
ALTER TABLE ONLY public.employees DROP CONSTRAINT employees_pkey;
ALTER TABLE ONLY public.departments DROP CONSTRAINT departments_pkey;
ALTER TABLE public.projects ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.employees_projects ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.employees ALTER COLUMN id DROP DEFAULT;
ALTER TABLE public.departments ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE public.projects_id_seq;
DROP TABLE public.projects;
DROP SEQUENCE public.employees_projects_id_seq;
DROP TABLE public.employees_projects;
DROP SEQUENCE public.employees_id_seq;
DROP TABLE public.employees;
DROP SEQUENCE public.departments_id_seq;
DROP TABLE public.departments;
SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: departments; Type: TABLE; Schema: public; Owner: supervisor
--

CREATE TABLE public.departments (
    id bigint NOT NULL,
    title character varying(255)
);


ALTER TABLE public.departments OWNER TO supervisor;

--
-- Name: departments_id_seq; Type: SEQUENCE; Schema: public; Owner: supervisor
--

CREATE SEQUENCE public.departments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.departments_id_seq OWNER TO supervisor;

--
-- Name: departments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: supervisor
--

ALTER SEQUENCE public.departments_id_seq OWNED BY public.departments.id;


--
-- Name: employees; Type: TABLE; Schema: public; Owner: supervisor
--

CREATE TABLE public.employees (
    id bigint NOT NULL,
    name character varying(255),
    surname character varying(255),
    department_id bigint
);


ALTER TABLE public.employees OWNER TO supervisor;

--
-- Name: employees_id_seq; Type: SEQUENCE; Schema: public; Owner: supervisor
--

CREATE SEQUENCE public.employees_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.employees_id_seq OWNER TO supervisor;

--
-- Name: employees_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: supervisor
--

ALTER SEQUENCE public.employees_id_seq OWNED BY public.employees.id;


--
-- Name: employees_projects; Type: TABLE; Schema: public; Owner: supervisor
--

CREATE TABLE public.employees_projects (
    id bigint NOT NULL,
    employee_id bigint,
    project_id bigint
);


ALTER TABLE public.employees_projects OWNER TO supervisor;

--
-- Name: employees_projects_id_seq; Type: SEQUENCE; Schema: public; Owner: supervisor
--

CREATE SEQUENCE public.employees_projects_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.employees_projects_id_seq OWNER TO supervisor;

--
-- Name: employees_projects_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: supervisor
--

ALTER SEQUENCE public.employees_projects_id_seq OWNED BY public.employees_projects.id;


--
-- Name: projects; Type: TABLE; Schema: public; Owner: supervisor
--

CREATE TABLE public.projects (
    id bigint NOT NULL,
    title character varying(255)
);


ALTER TABLE public.projects OWNER TO supervisor;

--
-- Name: projects_id_seq; Type: SEQUENCE; Schema: public; Owner: supervisor
--

CREATE SEQUENCE public.projects_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.projects_id_seq OWNER TO supervisor;

--
-- Name: projects_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: supervisor
--

ALTER SEQUENCE public.projects_id_seq OWNED BY public.projects.id;


--
-- Name: departments id; Type: DEFAULT; Schema: public; Owner: supervisor
--

ALTER TABLE ONLY public.departments ALTER COLUMN id SET DEFAULT nextval('public.departments_id_seq'::regclass);


--
-- Name: employees id; Type: DEFAULT; Schema: public; Owner: supervisor
--

ALTER TABLE ONLY public.employees ALTER COLUMN id SET DEFAULT nextval('public.employees_id_seq'::regclass);


--
-- Name: employees_projects id; Type: DEFAULT; Schema: public; Owner: supervisor
--

ALTER TABLE ONLY public.employees_projects ALTER COLUMN id SET DEFAULT nextval('public.employees_projects_id_seq'::regclass);


--
-- Name: projects id; Type: DEFAULT; Schema: public; Owner: supervisor
--

ALTER TABLE ONLY public.projects ALTER COLUMN id SET DEFAULT nextval('public.projects_id_seq'::regclass);


--
-- Data for Name: departments; Type: TABLE DATA; Schema: public; Owner: supervisor
--

COPY public.departments (id, title) FROM stdin;
1	IT
2	Marketing
3	Finance
\.


--
-- Data for Name: employees; Type: TABLE DATA; Schema: public; Owner: supervisor
--

COPY public.employees (id, name, surname, department_id) FROM stdin;
1	Danila	Ivashenko	1
2	Fedor	Kuznetsov	2
3	Egor	Egorov	1
4	Ilia	Iliin	3
\.


--
-- Data for Name: employees_projects; Type: TABLE DATA; Schema: public; Owner: supervisor
--

COPY public.employees_projects (id, employee_id, project_id) FROM stdin;
1	1	1
2	2	1
3	3	2
4	4	3
\.


--
-- Data for Name: projects; Type: TABLE DATA; Schema: public; Owner: supervisor
--

COPY public.projects (id, title) FROM stdin;
1	Project A
2	Project B
3	Project C
\.


--
-- Name: departments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: supervisor
--

SELECT pg_catalog.setval('public.departments_id_seq', 3, true);


--
-- Name: employees_id_seq; Type: SEQUENCE SET; Schema: public; Owner: supervisor
--

SELECT pg_catalog.setval('public.employees_id_seq', 4, true);


--
-- Name: employees_projects_id_seq; Type: SEQUENCE SET; Schema: public; Owner: supervisor
--

SELECT pg_catalog.setval('public.employees_projects_id_seq', 4, true);


--
-- Name: projects_id_seq; Type: SEQUENCE SET; Schema: public; Owner: supervisor
--

SELECT pg_catalog.setval('public.projects_id_seq', 3, true);


--
-- Name: departments departments_pkey; Type: CONSTRAINT; Schema: public; Owner: supervisor
--

ALTER TABLE ONLY public.departments
    ADD CONSTRAINT departments_pkey PRIMARY KEY (id);


--
-- Name: employees employees_pkey; Type: CONSTRAINT; Schema: public; Owner: supervisor
--

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT employees_pkey PRIMARY KEY (id);


--
-- Name: employees_projects employees_projects_pkey; Type: CONSTRAINT; Schema: public; Owner: supervisor
--

ALTER TABLE ONLY public.employees_projects
    ADD CONSTRAINT employees_projects_pkey PRIMARY KEY (id);


--
-- Name: projects projects_pkey; Type: CONSTRAINT; Schema: public; Owner: supervisor
--

ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_pkey PRIMARY KEY (id);


--
-- Name: employees employees_department_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: supervisor
--

ALTER TABLE ONLY public.employees
    ADD CONSTRAINT employees_department_id_fkey FOREIGN KEY (department_id) REFERENCES public.departments(id);


--
-- Name: employees_projects employees_projects_employee_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: supervisor
--

ALTER TABLE ONLY public.employees_projects
    ADD CONSTRAINT employees_projects_employee_id_fkey FOREIGN KEY (employee_id) REFERENCES public.employees(id) ON DELETE CASCADE;


--
-- Name: employees_projects employees_projects_project_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: supervisor
--

ALTER TABLE ONLY public.employees_projects
    ADD CONSTRAINT employees_projects_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.projects(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

