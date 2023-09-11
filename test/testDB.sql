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

ALTER TABLE ONLY public.delivery
    ADD CONSTRAINT delivery_pkey PRIMARY KEY (order_uid);


ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (order_uid);


ALTER TABLE ONLY public.payment
    ADD CONSTRAINT payment_pkey PRIMARY KEY (order_uid);



