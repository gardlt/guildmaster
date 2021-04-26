CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE SCHEMA IF NOT EXISTS "genbo";

COMMENT ON SCHEMA public IS 'standard public schema';

/* Table 'stores' */
CREATE TABLE genbo.stores (
id uuid DEFAULT uuid_generate_v4() NOT NULL,
"name" character varying NOT NULL,
owner_id uuid NOT NULL,
store_information_id uuid NOT NULL,
PRIMARY KEY(id));

/* Table 'owner' */
CREATE TABLE genbo."owner" (
id uuid DEFAULT uuid_generate_v4() NOT NULL,
"name" varchar NOT NULL,
created_at timestamp with time zone DEFAULT now() NOT NULL,
PRIMARY KEY(id));
CREATE UNIQUE INDEX owner_uniqe_name ON genbo."owner" (name);


/* Table 'products' */
CREATE TABLE genbo.products (
id uuid DEFAULT uuid_generate_v4() NOT NULL,
"name" varchar NOT NULL,
sku varchar NOT NULL,
PRIMARY KEY(id));

/* Table 'inventory' */
CREATE TABLE genbo.inventory (
id uuid DEFAULT uuid_generate_v4() NOT NULL,
stores_id uuid DEFAULT uuid_generate_v4() NOT NULL,
quantity integer NOT NULL,
products_id uuid NOT NULL,
PRIMARY KEY(id));

/* Table 'languages' */
CREATE TABLE genbo.languages (
id uuid DEFAULT uuid_generate_v4() NOT NULL,
"name" varchar NOT NULL,
code varchar NOT NULL,
PRIMARY KEY(id));

/* Table 'store_information' */
CREATE TABLE genbo.store_information (
id uuid DEFAULT uuid_generate_v4() NOT NULL,
address varchar NOT NULL,
phone_number varchar,
timezone varchar NOT NULL,
support_email varchar NOT NULL,
PRIMARY KEY(id));

/* Table 'settings' */
CREATE TABLE genbo.settings (
id uuid DEFAULT uuid_generate_v4() NOT NULL,
stores_id uuid DEFAULT uuid_generate_v4() NOT NULL,
"key" varchar NOT NULL,
"value" varchar NOT NULL,
PRIMARY KEY(id));

/* Table 'currency' */
CREATE TABLE genbo.currency (
id uuid NOT NULL,
"name" varchar NOT NULL,
code varchar NOT NULL,
PRIMARY KEY(id));

/* Relation 'owner-stores' */
ALTER TABLE genbo.stores ADD CONSTRAINT "owner-stores"
FOREIGN KEY (owner_id)
REFERENCES genbo."owner"(id);

/* Relation 'stores-inventory' */
ALTER TABLE genbo.inventory ADD CONSTRAINT "stores-inventory"
FOREIGN KEY (stores_id)
REFERENCES genbo.stores(id);

/* Relation 'stores-configurations' */
ALTER TABLE genbo.settings ADD CONSTRAINT "stores-configurations"
FOREIGN KEY (stores_id)
REFERENCES genbo.stores(id);

/* Relation 'store_information-stores' */
ALTER TABLE genbo.stores ADD CONSTRAINT "store_information-stores"
FOREIGN KEY (store_information_id)
REFERENCES genbo.store_information(id);

/* Relation 'products-inventory' */
ALTER TABLE genbo.inventory ADD CONSTRAINT "products-inventory"
FOREIGN KEY (products_id)
REFERENCES genbo.products(id);