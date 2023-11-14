# Coding Challenge

Thanks for applying for a backend role at Outdoorsy. We've put together this code challenge, which should take around 3-4 hours to complete.

## Running
### Running the db
1. Go to the project home directory
2. Access the directory migrations via `cd migrations`
3. Run the db with `docker-compose up`

### Running the server
1. Go to the project home directory
2. Run the server with `go run main.go`

### Running the tests
1. Go to the project home directory
2. Install the necessary tools with `make tools-install`
3. Run the unit tests with `make unit-test`

## Functionality
The task is to develop a rentals JSON API that returns a list of rentals that can be filtered, sorted, and paginated. We have included files to create a database of rentals.

Your application should support the following endpoints.

- `/rentals/<RENTAL_ID>` Read one rental endpoint
- `/rentals` Read many (list) rentals endpoint
    - Supported query parameters
        - price_min (number)
        - price_max (number)
        - limit (number)
        - offset (number)
        - ids (comma separated list of rental ids)
        - near (comma separated pair [lat,lng])
        - sort (string)
    - Examples:
        - `rentals?price_min=9000&price_max=75000`
        - `rentals?limit=3&offset=6`
        - `rentals?ids=3,4,5`
        - `rentals?near=33.64,-117.93` // within 100 miles
        - `rentals?sort=price`
        - `rentals?near=33.64,-117.93&price_min=9000&price_max=75000&limit=3&offset=6&sort=price`

The rental object JSON in the response should have the following structure:
```json
{
  "id": "int",
  "name": "string",
  "description": "string",
  "type": "string",
  "make": "string",
  "model": "string",
  "year": "int",
  "length": "decimal",
  "sleeps": "int",
  "primary_image_url": "string",
  "price": {
    "day": "int"
  },
  "location": {
    "city": "string",
    "state": "string",
    "zip": "string",
    "country": "string",
    "lat": "decimal",
    "lng": "decimal"
  },
  "user": {
    "id": "int",
    "first_name": "string",
    "last_name": "string"
  }
}
```

## Notes
- Running `docker-compose up` will automatically generate a postgres database and some data to work with. Connect and use this database.
- Write production ready code.
- Please make frequent, and descriptive git commits.
- Use third-party libraries or not; your choice.
- Please use Golang to complete this task.
- Feel free to add functionality as you have time, but the feature described above is the priority.
- Please add tests

## What we're looking for
- The functionality of the project matches the description above
- An ability to think through all potential states
- In the README of the project, describe exactly how to run the application and execute the tests

When complete, please push your code to Github to your own account and send the link to the project or zip the project (including the `.git` directory) and send it back.

Thank you and please ask if you have any questions!
