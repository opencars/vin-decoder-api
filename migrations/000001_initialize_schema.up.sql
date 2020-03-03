CREATE TABLE manufacturers
(
    "wmi" VARCHAR NOT NULL,
    "name" TEXT NOT NULL
);

CREATE UNIQUE INDEX manufacturers_wmi_idx ON manufacturers("wmi");
