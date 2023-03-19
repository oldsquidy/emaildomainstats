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

# Assumptions made

- Sorting alphabetically, you would want to group the letters together Capital first. e.g. AaBbCc... 
- A customer with the same name cannot be assumed to be a unique customer, the email address is the only unique item so I will discount any information other than email
- You could assume someone with the same name and ip addresses is the same person and therefor treat them as a customer with multiple ip addresses. However there are exceptions to this, such as father and son could have the same name and live in the same property, thus also sharing an IP address. Once more we can only guarantee that the email address is the only unique item