# WebPushStartInstallationResponse

**Properties**

| Name      | Type   | Required | Description                                                |
| :-------- | :----- | :------- | :--------------------------------------------------------- |
| AuthToken | string | ✅       | Auth secret returned from PushSubscription.getKey('auth'). |
| PublicKey | string | ✅       | VAPID public key generated for this web push installation. |
