# try-mailer
An attempt to try abstracts mailing framework

### Usage and Examples

 Route                | Description
---------------------|---------------------------------------------
`POST /send`         | send mail using json object
`GET /history/:mail` | list sending history from given emai address

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
