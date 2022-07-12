create table country
(
    id   bigserial
        constraint country_pk
            primary key,
    name varchar(128) not null
        constraint country_ak_1
            unique
);

create table currency
(
    id               bigserial
        constraint currency_pk
            primary key,
    code             varchar(8)   not null
        constraint currency_ak_1
            unique,
    name             varchar(128) not null
        constraint currency_ak_2
            unique,
    is_active        boolean      not null,
    is_base_currency boolean      not null
);

create table currency_rate
(
    id               bigserial
        constraint currency_rate_pk
            primary key,
    currency_id      integer        not null
        constraint currency_rate_currency
            references currency,
    base_currency_id integer        not null
        constraint currency_rate_base_currency
            references currency,
    rate             numeric(16, 6) not null,
    ts               timestamp      not null
);

create table currency_used
(
    id          bigserial
        constraint currency_used_pk
            primary key,
    country_id  integer not null
        constraint currency_used_country
            references country,
    currency_id integer not null
        constraint currency_used_currency
            references currency,
    date_from   date    not null,
    date_to     date
);

create table item
(
    id          bigserial
        constraint item_pk
            primary key,
    code        varchar(64)  not null
        constraint item_ak_1
            unique,
    name        varchar(255) not null,
    is_active   boolean      not null,
    currency_id integer      not null
        constraint item_currency
            references currency,
    details     text
);

create table price
(
    id          bigserial
        constraint price_pk
            primary key,
    item_id     integer        not null
        constraint price_item
            references item,
    currency_id integer        not null
        constraint price_currency
            references currency,
    buy         numeric(16, 6) not null,
    sell        numeric(16, 6) not null,
    ts          timestamp      not null
);

create table report
(
    id           bigserial
        constraint report_pk
            primary key,
    trading_date date    not null,
    item_id      integer not null
        constraint report_item
            references item,
    currency_id  integer not null
        constraint report_currency
            references currency,
    first_price  numeric(16, 6),
    last_price   numeric(16, 6),
    min_price    numeric(16, 6),
    max_price    numeric(16, 6),
    avg_price    numeric(16, 6),
    total_amount numeric(16, 6),
    quantity     numeric(16, 6),
    constraint report_ak_1
        unique (trading_date, item_id, currency_id)
);

create table trader
(
    id                    bigserial
        constraint trader_pk
            primary key,
    first_name            varchar(64)  not null,
    last_name             varchar(64)  not null,
    user_name             varchar(64)  not null,
    password              varchar(64)  not null,
    email                 varchar(128) not null,
    confirmation_code     varchar(128) not null,
    time_registered       timestamp    not null,
    time_confirmed        timestamp    not null,
    country_id            integer
        constraint trader_country
            references country,
    preferred_currency_id integer
        constraint trader_currency
            references currency,
    constraint trader_ak_1
        unique (user_name, email)
);

create table current_inventory
(
    id        bigserial
        constraint current_inventory_pk
            primary key,
    trader_id integer        not null
        constraint current_inventory_trader
            references trader,
    item_id   integer        not null
        constraint current_inventory_item
            references item,
    quantity  numeric(16, 6) not null,
    constraint current_inventory_ak_1
        unique (trader_id, item_id)
);

create table offer
(
    id        bigserial
        constraint offer_pk
            primary key,
    trader_id integer        not null
        constraint offer_trader
            references trader,
    item_id   integer        not null
        constraint offer_item
            references item,
    quantity  numeric(16, 6) not null,
    buy       boolean        not null,
    sell      boolean        not null,
    price     numeric(16, 6),
    ts        timestamp      not null,
    is_active boolean        not null
);

create table trade
(
    id          bigserial
        constraint trade_pk
            primary key,
    item_id     integer        not null
        constraint trade_item
            references item,
    seller_id   integer
        constraint trade_seller
            references trader,
    buyer_id    integer        not null
        constraint trade_buyer
            references trader,
    quantity    numeric(16, 6) not null,
    unit_price  numeric(16, 6) not null,
    description text           not null,
    offer_id    integer        not null
        constraint trade_offer
            references offer
);

UPDATE public.trader
SET first_name            = 'Luis',
    last_name             = 'Gimenez',
    user_name             = 'lgimen',
    password              = '122345',
    email                 = 'dunztar@gmail.com',
    confirmation_code     = '23445',
    time_registered       = '2022-07-07 20:18:26.000000',
    time_confirmed        = '2022-07-07 20:18:27.000000',
    country_id            = null,
    preferred_currency_id = null
WHERE id = 1;