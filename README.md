# Seller Identity Package

## Getting the server running
Set up the database:
- sudo mysql
- source "path to database.sql"
- source "path to test_data.sql" (optional to fill database with data)

Install missing packages (if any) 

cd \<name of folder>

go build \<name of folder> (or go build *.go)

navigate to localhost:2000

## Currently supported:
localhost:2000/ - landing page

localhost:2000/create/ - create new seller

localhost:2000/view/?id="id of seller" - display seller info

## todo:
- improve the seller's page
- save seller info to database instead of .txt files (and retrieve as well)
- QR code