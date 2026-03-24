# ExpoToken

**Properties**

| Name        | Type   | Required | Description                                                |
| :---------- | :----- | :------- | :--------------------------------------------------------- |
| CreatedAt   | string | ✅       | The timestamp when the token was created.                  |
| DeviceToken | string | ✅       | The Expo push token returned by the Expo client.           |
| Id          | string | ✅       | The unique identifier for the token.                       |
| DiscardedAt | string | ❌       | The timestamp when the token was discarded, if applicable. |
| UpdatedAt   | string | ❌       | The timestamp when the token metadata last changed.        |
