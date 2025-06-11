# Broadcast

**Properties**

| Name             | Type                       | Required | Description                                   |
| :--------------- | :------------------------- | :------- | :-------------------------------------------- |
| Recipients       | []shared.User              | ✅       |                                               |
| Title            | string                     | ✅       |                                               |
| ActionUrl        | string                     | ❌       |                                               |
| Category         | string                     | ❌       |                                               |
| Content          | string                     | ❌       |                                               |
| CreatedAt        | string                     | ❌       | The timestamp when the broadcast was created. |
| CustomAttributes | any                        | ❌       |                                               |
| Id               | string                     | ❌       | The unique id for this broadcast.             |
| Overrides        | broadcasts.Overrides       | ❌       |                                               |
| Status           | broadcasts.BroadcastStatus | ❌       |                                               |
| Topic            | string                     | ❌       |                                               |

# Overrides

**Properties**

| Name      | Type                         | Required | Description |
| :-------- | :--------------------------- | :------- | :---------- |
| Channels  | broadcasts.OverridesChannels | ❌       |             |
| Providers | broadcasts.Providers         | ❌       |             |

# OverridesChannels

**Properties**

| Name       | Type                  | Required | Description |
| :--------- | :-------------------- | :------- | :---------- |
| Email      | broadcasts.Email      | ❌       |             |
| InApp      | broadcasts.InApp      | ❌       |             |
| MobilePush | broadcasts.MobilePush | ❌       |             |
| Sms        | broadcasts.Sms        | ❌       |             |

# Email

**Properties**

| Name      | Type   | Required | Description |
| :-------- | :----- | :------- | :---------- |
| ActionUrl | string | ❌       |             |
| Content   | string | ❌       |             |
| Title     | string | ❌       |             |

# InApp

**Properties**

| Name      | Type   | Required | Description |
| :-------- | :----- | :------- | :---------- |
| ActionUrl | string | ❌       |             |
| Content   | string | ❌       |             |
| Title     | string | ❌       |             |

# MobilePush

**Properties**

| Name      | Type   | Required | Description |
| :-------- | :----- | :------- | :---------- |
| ActionUrl | string | ❌       |             |
| Content   | string | ❌       |             |
| Title     | string | ❌       |             |

# Sms

**Properties**

| Name      | Type   | Required | Description |
| :-------- | :----- | :------- | :---------- |
| ActionUrl | string | ❌       |             |
| Content   | string | ❌       |             |
| Title     | string | ❌       |             |

# Providers

**Properties**

| Name     | Type | Required | Description |
| :------- | :--- | :------- | :---------- |
| Apns     | any  | ❌       |             |
| Expo     | any  | ❌       |             |
| Fcm      | any  | ❌       |             |
| Mailgun  | any  | ❌       |             |
| Sendgrid | any  | ❌       |             |
| Ses      | any  | ❌       |             |
| Slack    | any  | ❌       |             |
| Teams    | any  | ❌       |             |
| Twilio   | any  | ❌       |             |
| WebPush  | any  | ❌       |             |

# BroadcastStatus

**Properties**

| Name    | Type                    | Required | Description |
| :------ | :---------------------- | :------- | :---------- |
| Errors  | []broadcasts.Errors     | ✅       |             |
| Status  | broadcasts.StatusStatus | ✅       |             |
| Summary | broadcasts.Summary      | ✅       |             |

# Errors

**Properties**

| Name    | Type   | Required | Description |
| :------ | :----- | :------- | :---------- |
| Message | string | ❌       |             |

# StatusStatus

**Properties**

| Name       | Type   | Required | Description  |
| :--------- | :----- | :------- | :----------- |
| enqueued   | string | ✅       | "enqueued"   |
| processing | string | ✅       | "processing" |
| processed  | string | ✅       | "processed"  |

# Summary

**Properties**

| Name     | Type  | Required | Description                                              |
| :------- | :---- | :------- | :------------------------------------------------------- |
| Failures | int64 | ✅       | The number of failures while processing the broadcast.   |
| Total    | int64 | ✅       | The number of recipients that the broadcast was sent to. |
