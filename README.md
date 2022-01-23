# GoWithPostgreSQL
 Go code for PostgreSQL. A Go language code which connects to PostgreSQL database for CRUD operation.

## Softwares required
Go compiler --> https://go.dev/doc/install  
Visual Studio code --> https://code.visualstudio.com/download  
PostgreSQL --> https://www.postgresql.org/download/  

## Step 1
Open a command prompt and cd to your home directory.  
  
  On Linux or Mac:  
    
    cd

  On Windows:  

    cd %HOMEPATH%

## Step 2
Create a GoWithPostgreSQL directory.
    For example, from your home directory use the following commands:

    mkdir GoWithPostgreSQL
    cd GoWithPostgreSQL

## Step 3
Start your module using the **go mod init** command.

    go mod init gowithpostgresql

## Step 4
In your text editor, create **main.go** and paste the content and save.

## Step 5
Build the current module's packages and dependencies.

    go mod tidy

## Step 6
Prepare the database for the repo.

### Step 6a
Create a database named **customer**

    CREATE DATABASE customer
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1;

### Step 6b
Create a schema named **customer** in customer database.

    CREATE SCHEMA customer
    AUTHORIZATION postgres;
   
### Step 6c
Create a table called **customer**.

    CREATE TABLE IF NOT EXISTS customer.customer
    (
        customerid bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
        firstname character varying(200) COLLATE pg_catalog."default",
        lastname character varying(200) COLLATE pg_catalog."default",
        CONSTRAINT customer_pkey PRIMARY KEY (customerid)
    )

    TABLESPACE pg_default;

    ALTER TABLE IF EXISTS customer.customer
    OWNER to postgres;

### Step 6d
Insert few reords to the **customer** table.

    INSERT INTO customer.customer(
	firstname, lastname)
	VALUES ('John', 'Doe'),
	('Richard', 'Roe'),
	('Mark', 'Moe');

### Step 6e
Create the stored procedure **insertcustomer**

    CREATE OR REPLACE PROCEDURE customer.insertcustomer(
	    IN firstname character varying,
	    IN lastname character varying)
    LANGUAGE 'sql'
    AS $BODY$
    INSERT INTO customer.customer(
	 firstname, lastname)
		VALUES (firstname,lastname);
    $BODY$;

## Step 7
Build the repo and execute the code in the terminal

    go build .\main.go
    .\main.exe

## References
https://go.dev/blog/using-go-modules  
https://go.dev/doc/tutorial/create-module  
https://go.dev/doc/database/
https://github.com/lib/pq