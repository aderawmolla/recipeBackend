CREATE TABLE "public"."instruction" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "step" integer NOT NULL, "instruction" text NOT NULL, "recipeId" uuid NOT NULL, PRIMARY KEY ("id") , FOREIGN KEY ("recipeId") REFERENCES "public"."recipes"("id") ON UPDATE cascade ON DELETE cascade);
CREATE EXTENSION IF NOT EXISTS pgcrypto;
