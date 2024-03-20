# stori-transactions

For detailed instructions, read the README.md file in each of the folders inside this directory.

## Directories
* `/lambda-function`: Contains the code that the AWS lambda function uses to read an event from an S3 bucket that stores `.csv` files, and sends an email with its content summary.
* `/local-project`: Contains the code that reads a local `.csv` file with transactions, and sends an email via SMTP protocol with a summary of them.

## How to upload files to S3 bucket that contains transactions .csv files?

Use the following curl to upload `.csv` files representing transactions in the S3 bucket named `stori-cvs-transactions`:

```
curl --location --request PUT 'https://yf58uztgna.execute-api.us-east-2.amazonaws.com/test/stori-cvs-bucket/transactions' \
--header 'Content-Type: text/csv' \
--data '@/Users/juanestebanmendez/Downloads/test-api.csv'
```
If you use Postman, select the binary option in dropdown when selecting the body tab, and choose a .csv file from your computer.