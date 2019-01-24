# try-mailer

An attempt to try abstracts mailing framework

### development dependencies
- PostgreSQL >= 9.5
- go version >= 1.11

### Usage and Examples

 Route                | Description
----------------------|---------------------------------------------
`POST /send`          | send mail using json object
`GET /history/:mail`  | list sending history from given email address

JSON request example, note that attachment content is base64 encoded

```json
{
  "sender": "mailer@domain.com",
  "recipients": ["someone@gmail.com"],
  "bcc": ["somehiddenone@gmail.com"],
  "cc": ["somerecipient@gmail.com"],
  "subject": "send mail testing try-mailer",
  "html_message": "<p>hello html from try-mailer</p>",
  "text_message": "hello text from try-mailer",
  "attachments": [
    {
      "name": "text.txt",
      "body": "cG9zdG1hbiBmaWxlIGJvZHkK",
      "content_type": "text/plain"
    }
  ]
}
```

### Authentication
This project use key authentication make a request with following format:
`Authorization: Bearer YOURCREDENTIALSHERE`
on both `GET` and `POST` method

### Host it yourself
this project use go module to handle dependencies, run `go mod vendor` first then
compile and build binary with
```
export AUTH_KEY=$(pwgen -s 96 1) # generate strong credential here
go build -ldflags "-X main.authKey=$AUTH_KEY"
```
then export environmental variabables which may written inside .env file then run `source .env` to load configuration

Example `.env` config file
```
export MAILER_API_HOST="0.0.0.0"
export MAILER_API_PORT=40001
export MAILER_DEBUG="true"
export MAILER_LOG_LEVEL="DEBUG"

export MAILER_DB_HOST="localhost"
export MAILER_DB_PORT="5432"
export MAILER_DB_USER="postgres"
export MAILER_DB_PASSWORD=""
export MAILER_DB_NAME="mailer"

export SENDGRID_API_KEY="API_KEY_HERE"
export MAILGUN_API_KEY="API_KEY_HERE"
export MAILGUN_DOMAIN="yourprefix.example.com"
export MAILGUN_SENDING_TIMEOUT="10s"
```


## Development Aspect
This project uses echo as http framework because it handles request headers nicely and handler composing is quite nice
with flexible middleware support, PostgreSQL because it is only sane choice, see more here -> http://howfuckedismydatabase.com/

### Not A Fullstack Solution
This project comes with backend service only

### Project structure
This project layout supports feature growing by grouping most implementation into `/pkg/` and separate major executor into `/pkg/service`. with interface layout at the root of each service (e.g. `pkg/mailer` has mailer.go with defines interface and `pkg/mailer/*` is its interface implementations

The main executable will handle main configuration from `config.go` then passes to service in need.

There is migrations folder for holding database migrations at the root of project, with both up and down migrations.

### Trade-off and decisions
This project uses a lot of higher order function to handle configuration which maybe overkill in small project, but it ensure that code structure will remain readable and stateless

This project run service in almost stateless manner which seems inefficient, but it is less likely to be surprised by harmful state management, also, stateless improves scalability. you can run service as many as you desire.


### FAQs.

- Why there is no frontend?
  
  because the author have time restricted from his resident environment, he needs to work at his office only.
  so he choose to complete most of the core features instead. 
  also, his frontend knowledge is technically none.
  
  Also, there is much more superior frontend ready to use out there, it called "Postman".
  
 - Test does not run, seems like something is missing?
   
   because some of constants and variables contains personal infomation, so he can't pushes into his repository.
   
### Wishlists
 - [ ] Statsd integration
 - [ ] Graylog integration
 - [ ] Better auth service
 - [ ] RPC interface
 - [ ] Message Queue interface
 - [ ] docker-compose example
 - [ ] kubernetes deployment example
  
