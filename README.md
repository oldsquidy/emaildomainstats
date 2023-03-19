# Email Domain Stats

This application takes a CSV data file location and outputs CSV data containing the amount of customers per domain sorted alphabetically to Stdout.

## Running the application

Provide the application with a CSV file location as an argument

### Example application run
`go run main.go test_data/customer_data.csv` 

## Example output
```
A-domain.com,1
a-domain.com,2
B-domain.com,3
b-domain.com,4
Z-domain.com,5
z-domain.com,6
```