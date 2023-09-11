
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

INSERT INTO public.delivery (order_uid, name, phone, zip, city, address, region, email) VALUES
    ('1', 'Test Testov', '+9721111111', '2639809', 'Kiryat Mozkin', 'Ploshad Mira 15', 'Kraiot', 'test@gmail.com');

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

INSERT INTO public.items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES
    ('1', 1, 'WBILMTESTTRACK', 453, 'ab4219087a764ae0btest', 'Mascaras', 30, '0', 317, 2389212, 'Vivienne Sabo', 202);

INSERT INTO public.items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES
    ('1', 1, 'WBILMTESTTRACK', 1453, 'ab4219087a764ae0btest', 'Sdas', 50, '0', 317, 2389212, 'Vivienne Sabo', 202);

CREATE TABLE public.orders (
    order_uid character varying(20) NOT NULL,
    track_number character varying(20) NOT NULL,
    entry character varying(10) NOT NULL,
    locale character varying(2) NOT NULL,
    internal_signature character varying(100) NOT NULL,
    customer_id character varying(20) NOT NULL,
    delivery_service character varying(10) NOT NULL,
    shardkey character varying(10) NOT NULL,
    sm_id integer NOT NULL,
    date_created timestamp without time zone NOT NULL,
    oof_shard character varying(10) NOT NULL
);

INSERT INTO public.orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES
    ('1', 'WBILMTESTTRACK', 'WBIL', 'en', 'sad', 'test', 'meest', '9', 99, '2021-11-26 06:22:19', '1');

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

INSERT INTO public.payment (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) VALUES
    ('1', 'b563feb7b2b84b6test', '12345', 'USD', 'wbpay', 1817, 1637907727, 'alpha', 1500, 317, 0);

ALTER TABLE ONLY public.delivery
    ADD CONSTRAINT delivery_pkey PRIMARY KEY (order_uid);

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (order_uid);

ALTER TABLE ONLY public.payment
    ADD CONSTRAINT payment_pkey PRIMARY KEY (order_uid);


