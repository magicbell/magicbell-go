# InboxConfigPayload

**Properties**

| Name   | Type                | Required | Description                                                       |
| :----- | :------------------ | :------- | :---------------------------------------------------------------- |
| Images | integrations.Images | ✅       | Image overrides for assets used in the inbox UI.                  |
| Locale | string              | ✅       | Locale code (ISO language tag) used to localize built-in strings. |
| Theme  | integrations.Theme  | ✅       | Visual customization options for the hosted inbox widget.         |

# Images

Image overrides for assets used in the inbox UI.

**Properties**

| Name          | Type   | Required | Description                                             |
| :------------ | :----- | :------- | :------------------------------------------------------ |
| EmptyInboxUrl | string | ✅       | URL for the illustration shown when the inbox is empty. |

# Theme

Visual customization options for the hosted inbox widget.

**Properties**

| Name         | Type                           | Required | Description                                    |
| :----------- | :----------------------------- | :------- | :--------------------------------------------- |
| Banner       | integrations.Banner            | ❌       | Top banner styling options.                    |
| Dialog       | integrations.Dialog            | ❌       | Styling for confirmation and action dialogs.   |
| Footer       | integrations.Footer            | ❌       | Footer styling for the inbox modal.            |
| Header       | integrations.Header            | ❌       | Header styling for the inbox modal.            |
| Icon         | integrations.Icon              | ❌       | Launcher icon styling overrides.               |
| Notification | integrations.ThemeNotification | ❌       | Styling overrides for notification list items. |
| UnseenBadge  | integrations.UnseenBadge       | ❌       | Badge styling for unseen notification counts.  |

# Banner

Top banner styling options.

**Properties**

| Name              | Type    | Required | Description                               |
| :---------------- | :------ | :------- | :---------------------------------------- |
| BackgroundColor   | string  | ✅       | Banner background color.                  |
| FontSize          | string  | ✅       | Font size for banner text.                |
| TextColor         | string  | ✅       | Banner text color.                        |
| BackgroundOpacity | float64 | ❌       | Opacity applied to the banner background. |

# Dialog

Styling for confirmation and action dialogs.

**Properties**

| Name            | Type   | Required | Description                                     |
| :-------------- | :----- | :------- | :---------------------------------------------- |
| AccentColor     | string | ✅       | Accent color for dialog buttons and highlights. |
| BackgroundColor | string | ✅       | Dialog background color.                        |
| TextColor       | string | ✅       | Dialog text color.                              |

# Footer

Footer styling for the inbox modal.

**Properties**

| Name            | Type   | Required | Description                                    |
| :-------------- | :----- | :------- | :--------------------------------------------- |
| BackgroundColor | string | ✅       | Footer background color.                       |
| BorderRadius    | string | ✅       | Border radius applied to the footer container. |
| FontSize        | string | ✅       | Font size used in the footer.                  |
| TextColor       | string | ✅       | Footer text color.                             |

# Header

Header styling for the inbox modal.

**Properties**

| Name            | Type   | Required | Description                                    |
| :-------------- | :----- | :------- | :--------------------------------------------- |
| BackgroundColor | string | ✅       | Header background color.                       |
| BorderRadius    | string | ✅       | Border radius applied to the header container. |
| FontFamily      | string | ✅       | CSS font family for the header title.          |
| FontSize        | string | ✅       | Font size used in the header.                  |
| TextColor       | string | ✅       | Header text color.                             |

# Icon

Launcher icon styling overrides.

**Properties**

| Name        | Type   | Required | Description                                  |
| :---------- | :----- | :------- | :------------------------------------------- |
| BorderColor | string | ✅       | CSS color used for the icon border.          |
| Width       | string | ✅       | Width of the launcher icon (any CSS length). |

# ThemeNotification

Styling overrides for notification list items.

**Properties**

| Name      | Type                   | Required | Description                                     |
| :-------- | :--------------------- | :------- | :---------------------------------------------- |
| Default\_ | integrations.Default\_ | ✅       | Base styles applied to every notification item. |
| Unread    | integrations.Unread    | ✅       | Overrides for unread notifications.             |
| Unseen    | integrations.Unseen    | ✅       | Overrides for unseen notifications.             |

# Default\_

Base styles applied to every notification item.

**Properties**

| Name            | Type                      | Required | Description                                                |
| :-------------- | :------------------------ | :------- | :--------------------------------------------------------- |
| BackgroundColor | string                    | ✅       | Background color for notifications in their default state. |
| BorderRadius    | string                    | ✅       | Border radius applied to each notification card.           |
| FontFamily      | string                    | ✅       | Font family for notification text.                         |
| FontSize        | string                    | ✅       | Font size for notification text.                           |
| Margin          | string                    | ✅       | CSS margin applied around each notification card.          |
| TextColor       | string                    | ✅       | Default text color for notifications.                      |
| Hover           | integrations.DefaultHover | ❌       | Styles applied when a notification is hovered.             |
| State           | integrations.DefaultState | ❌       | Accent colors for notification state indicators.           |

# DefaultHover

Styles applied when a notification is hovered.

**Properties**

| Name            | Type   | Required | Description                |
| :-------------- | :----- | :------- | :------------------------- |
| BackgroundColor | string | ✅       | Background color on hover. |

# DefaultState

Accent colors for notification state indicators.

**Properties**

| Name  | Type   | Required | Description                         |
| :---- | :----- | :------- | :---------------------------------- |
| Color | string | ✅       | Color used for the state indicator. |

# Unread

Overrides for unread notifications.

**Properties**

| Name            | Type                     | Required | Description                                       |
| :-------------- | :----------------------- | :------- | :------------------------------------------------ |
| BackgroundColor | string                   | ✅       | Background color applied to unread notifications. |
| TextColor       | string                   | ✅       | Text color used when a notification is unread.    |
| Hover           | integrations.UnreadHover | ❌       | Hover styles for unread notifications.            |
| State           | integrations.UnreadState | ❌       | State indicator styling for unread notifications. |

# UnreadHover

Hover styles for unread notifications.

**Properties**

| Name            | Type   | Required | Description                                         |
| :-------------- | :----- | :------- | :-------------------------------------------------- |
| BackgroundColor | string | ✅       | Background color on hover for unread notifications. |

# UnreadState

State indicator styling for unread notifications.

**Properties**

| Name  | Type   | Required | Description                           |
| :---- | :----- | :------- | :------------------------------------ |
| Color | string | ✅       | Color for the unread state indicator. |

# Unseen

Overrides for unseen notifications.

**Properties**

| Name            | Type                     | Required | Description                                       |
| :-------------- | :----------------------- | :------- | :------------------------------------------------ |
| BackgroundColor | string                   | ✅       | Background color applied to unseen notifications. |
| TextColor       | string                   | ✅       | Text color used when a notification is unseen.    |
| Hover           | integrations.UnseenHover | ❌       | Hover styles for unseen notifications.            |
| State           | integrations.UnseenState | ❌       | State indicator styling for unseen notifications. |

# UnseenHover

Hover styles for unseen notifications.

**Properties**

| Name            | Type   | Required | Description                                         |
| :-------------- | :----- | :------- | :-------------------------------------------------- |
| BackgroundColor | string | ✅       | Background color on hover for unseen notifications. |

# UnseenState

State indicator styling for unseen notifications.

**Properties**

| Name  | Type   | Required | Description                           |
| :---- | :----- | :------- | :------------------------------------ |
| Color | string | ✅       | Color for the unseen state indicator. |

# UnseenBadge

Badge styling for unseen notification counts.

**Properties**

| Name            | Type   | Required | Description             |
| :-------------- | :----- | :------- | :---------------------- |
| BackgroundColor | string | ✅       | Badge background color. |

<!-- This file was generated by liblab | https://liblab.com/ -->
