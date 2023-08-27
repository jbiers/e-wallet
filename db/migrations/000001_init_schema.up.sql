CREATE TABLE "accounts" (
  "id" integer PRIMARY KEY,
  "username" varchar NOT NULL,
  "type" varchar NOT NULL,
  "document" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "balance" integer DEFAULT 0 NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" integer PRIMARY KEY,
  "amount" integer NOT NULL,
  "account_id" integer NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" integer PRIMARY KEY,
  "amount" integer NOT NULL,
  "sender_id" integer NOT NULL,
  "receiver_id" integer NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE INDEX ON "accounts" ("username");

CREATE INDEX ON "accounts" ("email");

CREATE INDEX ON "accounts" ("document");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transactions" ("sender_id");

CREATE INDEX ON "transactions" ("receiver_id");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("sender_id") REFERENCES "accounts" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("receiver_id") REFERENCES "accounts" ("id");
