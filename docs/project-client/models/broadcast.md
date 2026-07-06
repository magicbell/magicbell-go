# Broadcast

**Properties**

| Name             | Type                       | Required | Description                                                                 |
| :--------------- | :------------------------- | :------- | :-------------------------------------------------------------------------- |
| Recipients       | [][shared.User](user.md)   | ✅       | A collection of users or filters that determine who receives the broadcast. |
| Title            | string                     | ✅       | The subject or headline that will be shown to recipients.                   |
| ActionUrl        | string                     | ❌       | The URL recipients will be directed to when interacting with the broadcast. |
| Category         | string                     | ❌       | The label used to group broadcasts.                                         |
| Content          | string                     | ❌       | The body content delivered with the broadcast.                              |
| CreatedAt        | string                     | ❌       | The timestamp when the broadcast was created.                               |
| CustomAttributes | any                        | ❌       | Arbitrary custom data associated with the broadcast.                        |
| Id               | string                     | ❌       | The unique id for this broadcast.                                           |
| Overrides        | broadcasts.Overrides       | ❌       | Channel- or provider-specific values that override the defaults.            |
| Status           | broadcasts.BroadcastStatus | ❌       | The runtime state of the broadcast execution.                               |
| Topic            | string                     | ❌       | The topic that further classifies the broadcast.                            |

# Overrides

Channel- or provider-specific values that override the defaults.

**Properties**

| Name      | Type                         | Required | Description                                                    |
| :-------- | :--------------------------- | :------- | :------------------------------------------------------------- |
| Channels  | broadcasts.OverridesChannels | ❌       | Overrides that are scoped to individual delivery channels.     |
| Providers | broadcasts.Providers         | ❌       | Overrides that are scoped to specific providers for a channel. |

# OverridesChannels

Overrides that are scoped to individual delivery channels.

**Properties**

| Name       | Type                  | Required | Description                              |
| :--------- | :-------------------- | :------- | :--------------------------------------- |
| Email      | broadcasts.Email      | ❌       | Overrides for email notifications.       |
| InApp      | broadcasts.InApp      | ❌       | Overrides for in-app notifications.      |
| MobilePush | broadcasts.MobilePush | ❌       | Overrides for mobile push notifications. |
| Sms        | broadcasts.Sms        | ❌       | Overrides for SMS notifications.         |

# Email

Overrides for email notifications.

**Properties**

| Name      | Type   | Required | Description                                                 |
| :-------- | :----- | :------- | :---------------------------------------------------------- |
| ActionUrl | string | ❌       | The link associated with the channel-specific notification. |
| Content   | string | ❌       | The channel-specific content.                               |
| Title     | string | ❌       | The channel-specific title.                                 |

# InApp

Overrides for in-app notifications.

**Properties**

| Name      | Type   | Required | Description                                                 |
| :-------- | :----- | :------- | :---------------------------------------------------------- |
| ActionUrl | string | ❌       | The link associated with the channel-specific notification. |
| Content   | string | ❌       | The channel-specific content.                               |
| Title     | string | ❌       | The channel-specific title.                                 |

# MobilePush

Overrides for mobile push notifications.

**Properties**

| Name      | Type   | Required | Description                                                 |
| :-------- | :----- | :------- | :---------------------------------------------------------- |
| ActionUrl | string | ❌       | The link associated with the channel-specific notification. |
| Content   | string | ❌       | The channel-specific content.                               |
| Title     | string | ❌       | The channel-specific title.                                 |

# Sms

Overrides for SMS notifications.

**Properties**

| Name      | Type   | Required | Description                                                 |
| :-------- | :----- | :------- | :---------------------------------------------------------- |
| ActionUrl | string | ❌       | The link associated with the channel-specific notification. |
| Content   | string | ❌       | The channel-specific content.                               |
| Title     | string | ❌       | The channel-specific title.                                 |

# Providers

Overrides that are scoped to specific providers for a channel.

**Properties**

| Name     | Type | Required | Description                                                      |
| :------- | :--- | :------- | :--------------------------------------------------------------- |
| Apns     | any  | ❌       | Provider-specific overrides for Apple Push Notification service. |
| Expo     | any  | ❌       | Provider-specific overrides for Expo push notifications.         |
| Fcm      | any  | ❌       | Provider-specific overrides for Firebase Cloud Messaging.        |
| Mailgun  | any  | ❌       | Provider-specific overrides for Mailgun.                         |
| Sendgrid | any  | ❌       | Provider-specific overrides for Sendgrid.                        |
| Ses      | any  | ❌       | Provider-specific overrides for AWS SES.                         |
| Slack    | any  | ❌       | Provider-specific overrides for Slack.                           |
| Teams    | any  | ❌       | Provider-specific overrides for Microsoft Teams.                 |
| Twilio   | any  | ❌       | Provider-specific overrides for Twilio.                          |
| WebPush  | any  | ❌       | Provider-specific overrides for the web push provider.           |

# BroadcastStatus

The runtime state of the broadcast execution.

**Properties**

| Name    | Type                    | Required | Description                                                  |
| :------ | :---------------------- | :------- | :----------------------------------------------------------- |
| Errors  | []broadcasts.Errors     | ✅       | A list of errors encountered while processing the broadcast. |
| Status  | broadcasts.StatusStatus | ✅       | The overall processing status of the broadcast.              |
| Summary | broadcasts.Summary      | ✅       | The summary counts for total recipients and failures.        |

# Errors

**Properties**

| Name    | Type   | Required | Description                             |
| :------ | :----- | :------- | :-------------------------------------- |
| Message | string | ❌       | The details about the processing error. |

# StatusStatus

The overall processing status of the broadcast.

**Properties**

| Name       | Type   | Required | Description  |
| :--------- | :----- | :------- | :----------- |
| Enqueued   | string | ✅       | "enqueued"   |
| Processing | string | ✅       | "processing" |
| Processed  | string | ✅       | "processed"  |

# Summary

The summary counts for total recipients and failures.

**Properties**

| Name     | Type  | Required | Description                                              |
| :------- | :---- | :------- | :------------------------------------------------------- |
| Failures | int64 | ✅       | The number of failures while processing the broadcast.   |
| Total    | int64 | ✅       | The number of recipients that the broadcast was sent to. |

<!-- This file was generated by liblab | https://liblab.com/ -->
