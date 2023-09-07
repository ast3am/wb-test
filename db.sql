--
-- PostgreSQL database dump
--

-- Dumped from database version 15.4
-- Dumped by pg_dump version 15.3

-- Started on 2023-09-05 19:33:52 MSK

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
-- TOC entry 215 (class 1259 OID 16396)
-- Name: delivery; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.delivery (
    order_uid character varying(20) NOT NULL,
    name character varying(50) NOT NULL,
    phone character varying(20) NOT NULL,
    zip character varying(10) NOT NULL,
    city character varying(50) NOT NULL,
    address character varying(100) NOT NULL,
    region character varying(50) NOT NULL,
    email character varying(50) NOT NULL
);


ALTER TABLE public.delivery OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 16406)
-- Name: items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.items (
    order_uid character varying(20) NOT NULL,
    chrt_id integer NOT NULL,
    track_number character varying(50) NOT NULL,
    price integer NOT NULL,
    rid character varying(50) NOT NULL,
    name character varying(50) NOT NULL,
    sale integer NOT NULL,
    size character varying(20) NOT NULL,
    total_price integer NOT NULL,
    nm_id integer NOT NULL,
    brand character varying(50) NOT NULL,
    status integer NOT NULL
);


ALTER TABLE public.items OWNER TO postgres;

--
-- TOC entry 214 (class 1259 OID 16391)
-- Name: orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orders (
    order_uid character varying(20) NOT NULL,
    track_number character varying(20) NOT NULL,
    entry character varying(10) NOT NULL,
    locale character varying(2) NOT NULL,
    internal_signature character varying(100) NOT NULL,
    customer_id character varying(20) NOT NULL,
    delvery_service character varying(10) NOT NULL,
    shardkey character varying(10) NOT NULL,
    sm_id integer NOT NULL,
    data_created timestamp without time zone NOT NULL,
    oof_shard character varying(10) NOT NULL
);


ALTER TABLE public.orders OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 16401)
-- Name: payment; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.payment (
    order_uid character varying(20) NOT NULL,
    transaction character varying(50) NOT NULL,
    request_id character varying(50) NOT NULL,
    currency character varying(3) NOT NULL,
    provider character varying(10) NOT NULL,
    amount integer NOT NULL,
    payment_dt integer NOT NULL,
    bank character varying(50) NOT NULL,
    delivery_cost integer NOT NULL,
    goods_total integer NOT NULL,
    custom_fee integer NOT NULL
);


ALTER TABLE public.payment OWNER TO postgres;

--
-- TOC entry 3623 (class 0 OID 16396)
-- Dependencies: 215
-- Data for Name: delivery; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.delivery (order_uid, name, phone, zip, city, address, region, email) FROM stdin;
\.


--
-- TOC entry 3625 (class 0 OID 16406)
-- Dependencies: 217
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) FROM stdin;
\.


--
-- TOC entry 3622 (class 0 OID 16391)
-- Dependencies: 214
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delvery_service, shardkey, sm_id, data_created, oof_shard) FROM stdin;
\.


--
-- TOC entry 3624 (class 0 OID 16401)
-- Dependencies: 216
-- Data for Name: payment; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.payment (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) FROM stdin;
\.


--
-- TOC entry 3475 (class 2606 OID 16400)
-- Name: delivery delivery_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.delivery
    ADD CONSTRAINT delivery_pkey PRIMARY KEY (order_uid);


--
-- TOC entry 3479 (class 2606 OID 16410)
-- Name: items items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (order_uid);


--
-- TOC entry 3473 (class 2606 OID 16395)
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (order_uid);


--
-- TOC entry 3477 (class 2606 OID 16405)
-- Name: payment payment_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment
    ADD CONSTRAINT payment_pkey PRIMARY KEY (order_uid);


-- Completed on 2023-09-05 19:33:52 MSK

--
-- PostgreSQL database dump complete
--

