create
database ojire;

create sequence user_id_sequence;

create sequence product_id_sequence;

create sequence r_user_product_id_sequence;

create table "user"
(
    id          integer                  default nextval('user_id_sequence'::regclass) not null
        primary key,
    name        text                                                                   not null,
    password    text                                                                   not null,
    "phoneNo"   text                                                                   not null
        constraint phoneno_key
            unique,
    email       text                                                                   not null
        constraint email_key
            unique,
    "createdAt" timestamp with time zone default now()                                 not null,
    "updatedAt" timestamp with time zone default now()                                 not null
);

create table product
(
    id          integer                  default nextval('product_id_sequence'::regclass) not null
        primary key,
    name        text                                                                      not null,
    sku         text                                                                      not null
        constraint sku_key
            unique,
    quantity    integer                                                                   not null,
    "createdAt" timestamp with time zone default now()                                    not null,
    "updatedAt" timestamp with time zone default now()                                    not null
);

create table r_user_product
(
    id          integer                  default nextval('r_user_product_id_sequence'::regclass) not null
        primary key,
    "userId"    integer                                                                          not null
        references "user"
            on update cascade on delete cascade,
    "productId" integer                                                                          not null
        references product
            on update cascade on delete cascade,
    "createdAt" timestamp with time zone default now()                                           not null,
    "updatedAt" timestamp with time zone default now()                                           not null
);
