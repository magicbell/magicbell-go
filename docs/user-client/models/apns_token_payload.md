# ApnsTokenPayload

**Properties**

| Name           | Type                                    | Required | Description                                                                                                                                      |
| :------------- | :-------------------------------------- | :------- | :----------------------------------------------------------------------------------------------------------------------------------------------- |
| DeviceToken    | string                                  | ✅       | The APNs device token to register with MagicBell.                                                                                                |
| AppId          | string                                  | ❌       | The bundle identifier of the application registering this token. Use this to override the default identifier configured on the APNs integration. |
| InstallationId | channels.ApnsTokenPayloadInstallationId | ❌       | The APNs environment this token belongs to. If omitted we assume it targets `production`.                                                        |

# ApnsTokenPayloadInstallationId

The APNs environment this token belongs to. If omitted we assume it targets `production`.

**Properties**

| Name        | Type   | Required | Description   |
| :---------- | :----- | :------- | :------------ |
| Development | string | ✅       | "development" |
| Production  | string | ✅       | "production"  |
