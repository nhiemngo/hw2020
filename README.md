# Seller Identity Package

## Getting the server running
Set up the database (whenever changes to db are made):
- sudo mysql
- source "path to database.sql"
- source "path to test_data.sql" (optional to fill database with data)

Install missing packages (if any) 

cd \<name of folder>

go build \<name of folder> (or go run *.go)

navigate to localhost:2000

## Current Paths:
localhost:2000/create/ - create new seller. after creating, users will be redirected to /option/

localhost:2000/option/id/ - from here can go to either /view/ or /order/

localhost:2000/view/id/ - display info of seller with id

localhost:2000/order/id/ - order offline package

## todo:
- polish
- put generated qr code to /order/
- seed with fake data