# UserPreferences

**Properties**

| Name       | Type                  | Required | Description |
| :--------- | :-------------------- | :------- | :---------- |
| Categories | []channels.Categories | ❌       |             |

# Categories

**Properties**

| Name     | Type                | Required | Description |
| :------- | :------------------ | :------- | :---------- |
| Channels | []channels.Channels | ❌       |             |
| Key      | string              | ❌       |             |
| Label    | string              | ❌       |             |

# Channels

**Properties**

| Name    | Type   | Required | Description |
| :------ | :----- | :------- | :---------- |
| Enabled | bool   | ❌       |             |
| Name    | string | ❌       |             |
