CREATE SCHEMA IF NOT EXISTS "public";

CREATE TYPE "carrier_user_type" AS ENUM ('owner', 'admin', 'viewer');

CREATE TYPE "fwd_aft_type" AS ENUM ('fwd', 'aft');

CREATE TYPE "weight_type" AS ENUM (
	'take_off',
	'zero_fuel',
	'inflight',
	'taxi',
	'landing'
);

CREATE TABLE
	"public"."maximum_weight" (
		"uuid" uuid NOT NULL,
		"subtype" uuid NOT NULL,
		"name" varchar(16) NOT NULL,
		PRIMARY KEY ("uuid")
	);

CREATE TABLE
	"public"."ground_stability_map" (
		"registration" uuid NOT NULL,
		"ground_stability" uuid NOT NULL,
		"effective_date" date NOT NULL,
		PRIMARY KEY ("registration", "ground_stability")
	);

CREATE TABLE
	"public"."aircraft_weight_map" (
		"registration" uuid NOT NULL,
		"aircraft_weight" uuid NOT NULL,
		"effective_date" date NOT NULL,
		PRIMARY KEY ("registration", "aircraft_weight")
	);

CREATE TABLE
	"public"."maximum_weight_map" (
		"registration" uuid NOT NULL,
		"maximum_weights" uuid NOT NULL,
		"effective_date" date NOT NULL,
		PRIMARY KEY ("registration", "maximum_weights")
	);

CREATE TABLE
	"public"."registration" (
		"uuid" uuid NOT NULL,
		"subtype" uuid NOT NULL,
		"mark" varchar(16) NOT NULL,
		"name" varchar(16),
		"created_at" timestamp NOT NULL,
		PRIMARY KEY ("uuid")
	);

CREATE TABLE
	"public"."ground_stability" (
		"uuid" uuid NOT NULL,
		"subtype" uuid NOT NULL,
		"name" varchar(16) NOT NULL,
		PRIMARY KEY ("uuid")
	);

CREATE TABLE
	"public"."cabin" (
		"uuid" uuid NOT NULL,
		"subtype" uuid NOT NULL,
		"name" varchar(16) NOT NULL,
		"weight" bigint NOT NULL,
		"index" bigint NOT NULL,
		"created_at" timestamp NOT NULL,
		PRIMARY KEY ("uuid")
	);

CREATE TABLE
	"public"."centre_of_gravity_limit" (
		"uuid" uuid NOT NULL,
		"subtype" uuid NOT NULL,
		"name" varchar(16) NOT NULL,
		PRIMARY KEY ("uuid")
	);

CREATE TABLE
	"public"."cargo_hold_map" (
		"registration" uuid NOT NULL,
		"cargo_holds" uuid NOT NULL,
		"effective_date" date NOT NULL,
		PRIMARY KEY ("registration", "cargo_holds")
	);

CREATE TABLE
	"public"."fuel_tank_map" (
		"registration" uuid NOT NULL,
		"fuel_tank" uuid NOT NULL,
		"effective_date" date NOT NULL,
		PRIMARY KEY ("registration", "fuel_tank")
	);

CREATE TABLE
	"public"."centre_of_gravity_limit_map" (
		"registration" uuid NOT NULL,
		"centre_of_gravity_limits" uuid NOT NULL,
		"effective_date" date NOT NULL,
		PRIMARY KEY ("registration", "centre_of_gravity_limits")
	);

CREATE TABLE
	"public"."aircraft_constants_map" (
		"registration" uuid NOT NULL,
		"aircraft_constants" uuid NOT NULL,
		"effective_date" date NOT NULL,
		PRIMARY KEY ("registration", "aircraft_constants")
	);

CREATE TABLE
	"public"."subtype" (
		"uuid" uuid NOT NULL,
		"fleet" uuid NOT NULL,
		"name" varchar(16) NOT NULL,
		"created_at" timestamp NOT NULL,
		PRIMARY KEY ("uuid")
	);

CREATE TABLE
	"public"."fuel_tank_table" (
		"weight" bigint NOT NULL,
		"fuel_tank" uuid NOT NULL,
		"balance_arm" bigint NOT NULL,
		PRIMARY KEY ("weight", "fuel_tank")
	);

CREATE TABLE
	"public"."cargo_hold" (
		"uuid" uuid NOT NULL,
		"subtype" uuid NOT NULL,
		"name" varchar(16) NOT NULL,
		"created_at" timestamp NOT NULL,
		PRIMARY KEY ("uuid")
	);

CREATE TABLE
	"public"."aircraft_weight" (
		"uuid" uuid NOT NULL,
		"name" varchar(16) NOT NULL,
		"weight" bigint NOT NULL,
		"index" bigint NOT NULL,
		PRIMARY KEY ("uuid")
	);

CREATE TABLE
	"public"."cabin_map" (
		"registration" uuid NOT NULL,
		"cabin" uuid NOT NULL,
		"effective_date" date NOT NULL,
		PRIMARY KEY ("registration", "cabin")
	);

CREATE TABLE
	"public"."aircraft_constants" (
		"uuid" uuid NOT NULL,
		"subtype" uuid NOT NULL,
		"name" varchar(16) NOT NULL,
		"c" bigint NOT NULL,
		"k" bigint NOT NULL,
		"reference_station" bigint NOT NULL,
		"length_of_mac" bigint NOT NULL,
		"leading_edge_mac" bigint NOT NULL,
		PRIMARY KEY ("uuid")
	);

CREATE TABLE
	"public"."fuel_tank" (
		"uuid" uuid NOT NULL,
		"subtype" uuid NOT NULL,
		"name" varchar(16) NOT NULL,
		PRIMARY KEY ("uuid")
	);

CREATE TABLE
	"public"."maximum_weight_table" (
		"weight_type" weight_type NOT NULL,
		"maximum_weight" uuid NOT NULL,
		"weight" bigint NOT NULL,
		PRIMARY KEY ("weight_type", "maximum_weight")
	);

CREATE TABLE
	"public"."carrier" (
		"uuid" uuid NOT NULL,
		"name" varchar(16) NOT NULL,
		PRIMARY KEY ("uuid")
	);

CREATE TABLE
	"public"."user" (
		"uuid" uuid NOT NULL,
		"name" varchar(32) NOT NULL,
		PRIMARY KEY ("uuid")
	);

CREATE TABLE
	"public"."fleet" (
		"uuid" uuid NOT NULL,
		"carrier" uuid NOT NULL,
		"name" varchar(16) NOT NULL,
		PRIMARY KEY ("uuid")
	);

CREATE TABLE
	"public"."carrier_map" (
		"carrier" uuid NOT NULL,
		"user" uuid NOT NULL,
		"permission" carrier_user_type NOT NULL,
		PRIMARY KEY ("carrier", "user")
	);

-- for a given centre_of_gravity_limit point there must be a unique combination of weight_type, centre_of_gravity_limit reference, fwd_aft type and the weight value. mac is the stored value here.
CREATE TABLE
	"public"."centre_of_gravity_limit_table" (
		"weight_type" weight_type NOT NULL,
		"centre_of_gravity_limit" uuid NOT NULL,
		"fwd_aft" fwd_aft_type NOT NULL,
		"weight" bigint NOT NULL,
		"mac" bigint NOT NULL,
		PRIMARY KEY (
			"weight_type",
			"centre_of_gravity_limit",
			"fwd_aft",
			"weight"
		)
	);

COMMENT ON TABLE "public"."centre_of_gravity_limit_table" IS 'for a given centre_of_gravity_limit point there must be a unique combination of weight_type, centre_of_gravity_limit reference, fwd_aft type and the weight value. mac is the stored value here.';

CREATE TABLE
	"public"."ground_stability_table" (
		"ground_stability" uuid NOT NULL,
		"weight" bigint NOT NULL,
		"mac" bigint NOT NULL,
		PRIMARY KEY ("ground_stability", "weight")
	);

CREATE TABLE
	"public"."deck" (
		"uuid" uuid NOT NULL,
		-- for example deck name could be 'lower', 'main', etc.
		"name" varchar(16) NOT NULL,
		PRIMARY KEY ("uuid")
	);

COMMENT ON COLUMN "public"."deck"."name" IS 'for example deck name could be ''lower'', ''main'', etc.';

CREATE TABLE
	"public"."hold" (
		"uuid" uuid NOT NULL,
		"deck" uuid NOT NULL,
		-- for example hold name could be 'fwd', 'aft', 'bulk', etc.
		"name" varchar(16) NOT NULL,
		PRIMARY KEY ("uuid")
	);

COMMENT ON COLUMN "public"."hold"."name" IS 'for example hold name could be ''fwd'', ''aft'', ''bulk'', etc.';

CREATE TABLE
	"public"."compartment" (
		"uuid" uuid NOT NULL,
		"deck" uuid NOT NULL,
		"hold" uuid NOT NULL,
		-- for example compartment could be '1', '2'
		"name" varchar(16) NOT NULL,
		PRIMARY KEY ("uuid")
	);

COMMENT ON COLUMN "public"."compartment"."name" IS 'for example compartment could be ''1'', ''2''';

CREATE TABLE
	"public"."deck_map" (
		"deck" uuid NOT NULL,
		"cargo_hold" uuid NOT NULL,
		"effective_data" date NOT NULL,
		PRIMARY KEY ("deck", "cargo_hold")
	);

ALTER TABLE "public"."aircraft_constants_map" ADD CONSTRAINT "fk_aircraft_constants_map_aircraft_constants_aircraft_consta" FOREIGN KEY ("aircraft_constants") REFERENCES "public"."aircraft_constants" ("uuid");

ALTER TABLE "public"."aircraft_constants_map" ADD CONSTRAINT "fk_aircraft_constants_map_registration_registration_uuid" FOREIGN KEY ("registration") REFERENCES "public"."registration" ("uuid");

ALTER TABLE "public"."aircraft_constants" ADD CONSTRAINT "fk_aircraft_constants_subtype_subtype_uuid" FOREIGN KEY ("subtype") REFERENCES "public"."subtype" ("uuid");

ALTER TABLE "public"."cabin_map" ADD CONSTRAINT "fk_cabin_map_cabin_cabin_uuid" FOREIGN KEY ("cabin") REFERENCES "public"."cabin" ("uuid");

ALTER TABLE "public"."cabin_map" ADD CONSTRAINT "fk_cabin_map_registration_registration_uuid" FOREIGN KEY ("registration") REFERENCES "public"."registration" ("uuid");

ALTER TABLE "public"."cabin" ADD CONSTRAINT "fk_cabin_subtype_subtype_uuid" FOREIGN KEY ("subtype") REFERENCES "public"."subtype" ("uuid");

ALTER TABLE "public"."cargo_hold_map" ADD CONSTRAINT "fk_cargo_hold_map_cargo_holds_cargo_hold_uuid" FOREIGN KEY ("cargo_holds") REFERENCES "public"."cargo_hold" ("uuid");

ALTER TABLE "public"."cargo_hold_map" ADD CONSTRAINT "fk_cargo_hold_map_registration_registration_uuid" FOREIGN KEY ("registration") REFERENCES "public"."registration" ("uuid");

ALTER TABLE "public"."cargo_hold" ADD CONSTRAINT "fk_cargo_hold_subtype_subtype_uuid" FOREIGN KEY ("subtype") REFERENCES "public"."subtype" ("uuid");

ALTER TABLE "public"."centre_of_gravity_limit" ADD CONSTRAINT "fk_centre_of_gravity_limit_subtype_subtype_uuid" FOREIGN KEY ("subtype") REFERENCES "public"."subtype" ("uuid");

ALTER TABLE "public"."aircraft_weight_map" ADD CONSTRAINT "fk_aircraft_weight_map_aircraft_weight_aircraft_weight_uuid" FOREIGN KEY ("aircraft_weight") REFERENCES "public"."aircraft_weight" ("uuid");

ALTER TABLE "public"."aircraft_weight_map" ADD CONSTRAINT "fk_aircraft_weight_map_registration_registration_uuid" FOREIGN KEY ("registration") REFERENCES "public"."registration" ("uuid");

ALTER TABLE "public"."centre_of_gravity_limit_map" ADD CONSTRAINT "fk_centre_of_gravity_limit_map_centre_of_gravity_limits_cent" FOREIGN KEY ("centre_of_gravity_limits") REFERENCES "public"."centre_of_gravity_limit" ("uuid");

ALTER TABLE "public"."centre_of_gravity_limit_map" ADD CONSTRAINT "fk_centre_of_gravity_limit_map_registration_registration_uui" FOREIGN KEY ("registration") REFERENCES "public"."registration" ("uuid");

ALTER TABLE "public"."ground_stability_map" ADD CONSTRAINT "fk_ground_stability_map_registration_registration_uuid" FOREIGN KEY ("registration") REFERENCES "public"."registration" ("uuid");

ALTER TABLE "public"."maximum_weight_map" ADD CONSTRAINT "fk_maximum_weight_map_maximum_weights_maximum_weight_uuid" FOREIGN KEY ("maximum_weights") REFERENCES "public"."maximum_weight" ("uuid");

ALTER TABLE "public"."maximum_weight_map" ADD CONSTRAINT "fk_maximum_weight_map_registration_registration_uuid" FOREIGN KEY ("registration") REFERENCES "public"."registration" ("uuid");

ALTER TABLE "public"."registration" ADD CONSTRAINT "fk_registration_subtype_subtype_uuid" FOREIGN KEY ("subtype") REFERENCES "public"."subtype" ("uuid");

ALTER TABLE "public"."fuel_tank_map" ADD CONSTRAINT "fk_fuel_tank_map_fuel_tank_fuel_tank_uuid" FOREIGN KEY ("fuel_tank") REFERENCES "public"."fuel_tank" ("uuid");

ALTER TABLE "public"."fuel_tank_map" ADD CONSTRAINT "fk_fuel_tank_map_registration_registration_uuid" FOREIGN KEY ("registration") REFERENCES "public"."registration" ("uuid");

ALTER TABLE "public"."fuel_tank" ADD CONSTRAINT "fk_fuel_tank_subtype_subtype_uuid" FOREIGN KEY ("subtype") REFERENCES "public"."subtype" ("uuid");

ALTER TABLE "public"."fuel_tank_table" ADD CONSTRAINT "fk_fuel_tank_table_fuel_tank_fuel_tank_uuid" FOREIGN KEY ("fuel_tank") REFERENCES "public"."fuel_tank" ("uuid");

ALTER TABLE "public"."ground_stability" ADD CONSTRAINT "fk_ground_stability_subtype_subtype_uuid" FOREIGN KEY ("subtype") REFERENCES "public"."subtype" ("uuid");

ALTER TABLE "public"."ground_stability" ADD CONSTRAINT "fk_ground_stability_uuid_ground_stability_map_ground_stabili" FOREIGN KEY ("uuid") REFERENCES "public"."ground_stability_map" ("ground_stability");

ALTER TABLE "public"."maximum_weight" ADD CONSTRAINT "fk_maximum_weight_subtype_subtype_uuid" FOREIGN KEY ("subtype") REFERENCES "public"."subtype" ("uuid");

ALTER TABLE "public"."maximum_weight_table" ADD CONSTRAINT "fk_maximum_weight_table_maximum_weight_maximum_weight_uuid" FOREIGN KEY ("maximum_weight") REFERENCES "public"."maximum_weight" ("uuid");

ALTER TABLE "public"."fleet" ADD CONSTRAINT "fk_fleet_carrier_carrier_uuid" FOREIGN KEY ("carrier") REFERENCES "public"."carrier" ("uuid");

ALTER TABLE "public"."subtype" ADD CONSTRAINT "fk_subtype_fleet_fleet_uuid" FOREIGN KEY ("fleet") REFERENCES "public"."fleet" ("uuid");

ALTER TABLE "public"."user" ADD CONSTRAINT "fk_user_uuid_carrier_map_user" FOREIGN KEY ("uuid") REFERENCES "public"."carrier_map" ("user");

ALTER TABLE "public"."carrier_map" ADD CONSTRAINT "fk_carrier_map_carrier_carrier_uuid" FOREIGN KEY ("carrier") REFERENCES "public"."carrier" ("uuid");

ALTER TABLE "public"."centre_of_gravity_limit_table" ADD CONSTRAINT "fk_centre_of_gravity_limit_table_centre_of_gravity_limit_cen" FOREIGN KEY ("centre_of_gravity_limit") REFERENCES "public"."centre_of_gravity_limit" ("uuid");

ALTER TABLE "public"."ground_stability_table" ADD CONSTRAINT "fk_ground_stability_table_ground_stability_ground_stability_" FOREIGN KEY ("ground_stability") REFERENCES "public"."ground_stability" ("uuid");

ALTER TABLE "public"."hold" ADD CONSTRAINT "fk_hold_deck_deck_uuid" FOREIGN KEY ("deck") REFERENCES "public"."deck" ("uuid");

ALTER TABLE "public"."compartment" ADD CONSTRAINT "fk_compartment_hold_hold_uuid" FOREIGN KEY ("hold") REFERENCES "public"."hold" ("uuid");

ALTER TABLE "public"."deck_map" ADD CONSTRAINT "fk_deck_map_deck_deck_uuid" FOREIGN KEY ("deck") REFERENCES "public"."deck" ("uuid");

ALTER TABLE "public"."deck_map" ADD CONSTRAINT "fk_deck_map_cargo_hold_cargo_hold_uuid" FOREIGN KEY ("cargo_hold") REFERENCES "public"."cargo_hold" ("uuid");

ALTER TABLE "public"."compartment" ADD CONSTRAINT "fk_compartment_deck_deck_uuid" FOREIGN KEY ("deck") REFERENCES "public"."deck" ("uuid");
