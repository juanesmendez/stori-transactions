# stori-transactions

## How to run?
Execute the following commands standing in the current directory:

1. `git checkout main`
2. `docker build -t stori-transactions .`
3. `docker run -it -e EMAIL=<any_gmail_mail> -e EMAIL_PASSWORD=<gmail_generated_app_password> -e TO_EMAIL=<recipient_email>  stori-transactions`
