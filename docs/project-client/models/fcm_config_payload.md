# FcmConfigPayload

**Properties**

| Name                    | Type                | Required | Description                                                                                      |
| :---------------------- | :------------------ | :------- | :----------------------------------------------------------------------------------------------- |
| AuthProviderX509CertUrl | string              | ✅       | URL for Google's OAuth provider x509 certificates used to validate tokens.                       |
| AuthUri                 | string              | ✅       | OAuth authorization endpoint used when exchanging Firebase credentials.                          |
| ClientEmail             | string              | ✅       | The client email address from the Firebase service account.                                      |
| ClientId                | string              | ✅       | The numeric client identifier for the Firebase service account.                                  |
| ClientX509CertUrl       | string              | ✅       | URL to the public x509 certificate for this service account.                                     |
| PrivateKey              | string              | ✅       | The PEM encoded service account private key used to sign Firebase credentials.                   |
| PrivateKeyId            | string              | ✅       | Identifier of the private key inside the downloaded service account JSON.                        |
| ProjectId               | string              | ✅       | The Firebase project ID associated with this service account.                                    |
| TokenUri                | string              | ✅       | OAuth token endpoint used to mint access tokens for FCM.                                         |
| Type\_                  | integrations.Type\_ | ✅       | Indicates the kind of Google credential. Service accounts always use the `service_account` type. |
| UniverseDomain          | string              | ✅       | The Google Cloud universe domain hosting the Firebase APIs.                                      |

# Type\_

Indicates the kind of Google credential. Service accounts always use the `service_account` type.

**Properties**

| Name           | Type   | Required | Description       |
| :------------- | :----- | :------- | :---------------- |
| ServiceAccount | string | ✅       | "service_account" |
