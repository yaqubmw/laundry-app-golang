CREATE USER enigmalaundry WITH CREATEDB NOSUPERUSER INHERIT PASSWORD 'password';
CREATE DATABASE enigmalaundry OWNER enigmalaundry;

\c enigmalaundry enigmalaundry

create table customer (
    id varchar(100) primary key,
    name varchar(100),
    phone_number varchar(15) unique,
    address text
);

create table uom (
    id varchar(100) primary key,
    name varchar(30) not null
);

create table product (
    id varchar(100) primary key,
    name varchar(50) not null,
    price bigint,
    uom_id varchar(100),
    foreign key(uom_id) references uom(id)
);

create table employee (
    id varchar(100) primary key,
    name varchar(100),
    phone_number varchar(15) unique,
    address text
);

create table bill (
    id varchar(100) primary key,
    bill_date date,
    entry_date date,
    finish_date date,
    employee_id varchar(100),
    customer_id varchar(100),
    foreign key(employee_id) references employee(id),
    foreign key(customer_id) references customer(id)
);

create table bill_detail (
    id varchar(100) primary key,
    bill_id varchar(100),
    product_id varchar(100),
    product_price bigint,
    qty int
);