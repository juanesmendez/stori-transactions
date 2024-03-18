# stori-transactions

## How to execute code?
Execute the following commands standing in the current directory:

1. `docker build -t stori-transactions .`
2. `docker run -it -e EMAIL=<any_gmail_mail> -e EMAIL_PASSWORD=<gmail_generated_app_password> -e TO_EMAIL=<recipient_email>  stori-transactions`

## Code Interface Description

### Introduction

The stori-transactions project provides functionality to read from a `.csv` file contained in the root directory of the project, process it, and send a summary of its content via email. Each record of the .csv file represents a bank transaction. It has an `ID`, a `date`, and an `value` (which can be a debit or credit movement).

### Components

* FileReader: Responsible from reading the data from the `.csv` file and parsing it to convert it to a `Transaction` struct.
* TransactionsSummary: Responsible for traversing the `Transaction` structs slice, and calculating a summary of them.  
* Email: Responsible for injecting the data calculated by the TransactionSummary component, into an html template, and sending it via email using the smtp protocol.

### Inputs and Outputs

* FileReader:
  * Input: File path to csv
  * Output: Data read from the csv file
* TransactionsSummary:
  * Input: Transactions array
  * Output: MonthlyTransactionSummary object with detailed data about the transactions.
* Email
  * Input: Array of `Transaction` objects.
  * Output: None

### Functions or Methods

* FileReader:
  * `ReadTransactionsFromCsv(fileName string) -> []model.Transaction`: Reads csv files, parses records and returns an array of `Transaction` objects.
  * `getRecords(fileName string) -> ([][]string, error)`: Opens the csv file and gets the array of records.
  * `getTransactionId(record []string, rowNum int) -> (uint64, error)`: Gets transaction ID from each csv record.
  * `getDate(record []string, rowNum int) -> (time.Time, error)`: Gets date from each csv record.
  * `getTransactionType(record []string, rowNum int) -> (model.TransactionType, error)`: Gets transaction type of each csv record, based on if it is negative or positive (debit or credit respectively).
  * `getTransactionValue(record []string, rowNum int) -> (float64, error)`: Gets the transaction value from the csv record.
* TransactionsSummary:
  * `GetBalance(transactions []model.Transaction) -> float64`: Returns balance  of the bank account by calculating the average of all the transaction values.
  * `GroupTransactionsByMonth(transactions []model.Transaction) -> map[time.Month]MonthlyTransactionSummary`: Calculates a transactions summary grouped by each month of the year.
  * `AverageByType(transactions []model.Transaction) -> map[model.TransactionType]float64`: Calculates the average of the transactions value by transaction type (debit or credit)
  * `calculateMonthlyAvgValueByTransType(transSummaryByMonth map[time.Month]MonthlyTransactionSummary) -> map[time.Month]MonthlyTransactionSummary`: Calculates transactions monthly average value.
  * `getTransactionsSummaryByMonth(transactions []model.Transaction) -> map[time.Month]MonthlyTransactionSummary`: Gets general transactions summary grpuped by month, without segregating by type of transaction.

* Email
  * `SendEmail(body string) -> error`: Received the email body as a string and sends it via email using a gmail account with the use of the smtp protocol.
  * `StartEmail(transactions []model.Transaction)`: Received the array of transactions, calls all of the TransactionsSummary component functions, injects the data in an html template, parses it to a string, and calls the `SendEmail` function.

### Communication Protocol

* Components communicate through method calls.
* FileReader reads data from a file.
* Email component sends email using the SMTP protocol.