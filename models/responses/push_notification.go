package models

type PushNotificationRange struct {
	AppID            string                  `json:"app_id"`
	IncludePlayerIds []string                `json:"include_player_ids"`
	Contents         PushNotificationContent `json:"contents"`
}

type PushNotificationAll struct {
	AppID            string                  `json:"app_id"`
	IncludedSegments []string                `json:"included_segments"`
	Contents         PushNotificationContent `json:"contents"`
}

type PushNotificationContent struct {
	Text      string `json:"en"`
}

type PushNotificationResult struct {
	ID         string `json:"id"`
	Recipients int    `json:"recipients"`
	ExternalId int    `json:"external_id"`
}

type OneSignalInsertUserResult struct {
	Success bool   `json:"success"`
	ID      string `json:"id"`
}

type OneSignalUser struct {
	AppID      string `json:"app_id"`
	DeviceType int    `json:"device_type"`
	Identifier string `json:"identifier"`
}
