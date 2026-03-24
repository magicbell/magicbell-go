# SmtpConfig

**Properties**

| Name     | Type                           | Required | Description                     |
| :------- | :----------------------------- | :------- | :------------------------------ |
| From     | integrations.SmtpConfigFrom    | ✅       | Default sender email address    |
| Host     | string                         | ✅       | SMTP server hostname            |
| Password | string                         | ✅       | SMTP authentication password    |
| Port     | int64                          | ✅       | SMTP server port                |
| Username | string                         | ✅       | SMTP authentication username    |
| ReplyTo  | integrations.SmtpConfigReplyTo | ❌       | Reply-to email address          |
| Security | integrations.Security          | ❌       | SMTP security/encryption method |

# SmtpConfigFrom

Default sender email address

**Properties**

| Name  | Type   | Required | Description          |
| :---- | :----- | :------- | :------------------- |
| Email | string | ✅       | Sender email address |
| Name  | string | ❌       | Sender name          |

# SmtpConfigReplyTo

Reply-to email address

**Properties**

| Name  | Type   | Required | Description            |
| :---- | :----- | :------- | :--------------------- |
| Email | string | ✅       | Reply-to email address |
| Name  | string | ❌       | Reply-to name          |

# Security

SMTP security/encryption method

**Properties**

| Name     | Type   | Required | Description |
| :------- | :----- | :------- | :---------- |
| None     | string | ✅       | "none"      |
| Ssl      | string | ✅       | "ssl"       |
| Starttls | string | ✅       | "starttls"  |
