package streaming

type CreateStreamingRequest struct {
	ClientID uint   `gorm:"type:INTEGER not null" json:"client_id"`
	ClinicID uint   `gorm:"type:INTEGER not null" json:"clinic_id"`
	Url      string `gorm:"type:varchar(256) not null" json:"url"`
}
