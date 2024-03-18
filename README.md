# stori-transactions

For detailed instructions, read the README.md file in each of the folders inside this directory.

## Directories
* `/lambda-function`: Contains the code that the AWS lambda function uses to read an event from an S3 bucket that stores `.csv` files, and sends an email with its content summary.
* `/local-project`: Contains the code that reads a local `.csv` file with transactions, and sends an email via SMTP protocol with a summary of them.