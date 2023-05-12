CREATE TABLE "sessions" (
	"id" UUID PRIMARY KEY,
	"username" VARCHAR NOT NULL,
	"refresh_token" VARCHAR NOT NULL,
	"user_agent" VARCHAR NOT NULL,
	"client_ip" VARCHAR NOT NULL,
	"is_blocked" BOOLEAN NOT NULL DEFAULT FALSE,
	"expires_at" timestamptz NOT NULL,
	"created_at" timestamptz NOT NULL DEFAULT (now())
);

-- add foreign key
ALTER TABLE
	"sessions"
ADD
	FOREIGN KEY ("username") REFERENCES "users" ("username");