CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.stores (
id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
"name" varchar NOT NULL,
PRIMARY KEY(id));