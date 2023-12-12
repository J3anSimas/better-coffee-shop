create table user (
	id varchar(36) primary key,
    name varchar(255) not null,
    auth0_id varchar(30) not null
    
);
create table product (
	id varchar(36) primary key,
    description varchar(255) primary key,
    price integer not null,
    created_at timestamp default CURRENT_TIMESTAMP
);
create table cart (
	id varchar(36) primary key,
    user_id varchar(36) not null,
    created_at timestamp default CURRENT_TIMESTAMP
    );
create table cart_item (
	id varchar(36) primary key,
    cart_id varchar(36) not null,
    product_id varchar(36) not null,
    quantity integer not null,
    created_at timestamp default CURRENT_TIMESTAMP

);
create table payment_method (
    id varchar(36) primary key,
    description varchar(255) not null

);
)
create table order (
    id varchar(36) primary key,
    user_id varchar(36) not null,
    payment_id varchar(36) not null,
    created_at timestamp default CURRENT_TIMESTAMP
);
create table order_item (
    id varchar(36) primary key,
    order_id varchar(36) not null,
    product_id varchar(36) not null,
    quantity integer not null,
    price integer not null
);
alter table cart add constraint fk_cart_user foreign key (user_id) references user(id);
alter table cart_item add constraint fk_cart_item_cart foreign key (cart_id) references cart(id);
alter table cart_item add constraint fk_cart_item_product foreign key (product_id) references product(id);
alter table order add constraint fk_order_user foreign key (user_id) references user(id);
alter table order add constraint fk_order_payment foreign key (payment_id) references payment_method(id);
alter table order_item add constraint fk_order_item_order foreign key (order_id) references order(id);
alter table order_item add constraint fk_order_item_product foreign key (product_id) references product(id);
