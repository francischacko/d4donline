# d4donline-test
This project is written in go using fiber framework and have used mysql as database

## Features

- Implemented indexing to database for optimized data retrieval.
- The code uses goroutines to fetch data from the database concurrently. For each row in the result set, a new goroutine is launched to process the data.
- The code effectively retrieves data from the database based on the requested country and stores it in RAM using a map called countryDataMap.
- Implemented Rate limiting
- Implemented Graceful Shutdown
- Made use of compressiion middleware for api responses
- Implemented unit testng
