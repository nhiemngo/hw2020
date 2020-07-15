# Seller Identity Package

## Starting the server
cd sellercom

go build sellercom

sudo mysql

source "path to database.sql"
source "path to test_data.sql" (optional to fill database with data)

navigate to localhost:2000 (that 2000 is a random port number btw)

## Currently supported:
localhost:2000/ - landing page

localhost:2000/create/ - create new seller

localhost:2000/view/?id="id of seller" - display seller info

## todo:
- improve the seller's page
- save seller info to database instead of .txt files (and retrieve as well)
- QR code
