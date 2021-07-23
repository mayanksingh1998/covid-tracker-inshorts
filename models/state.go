package models
import (
"github.com/google/uuid"
"time"
)

type State struct {
	ID        uuid.UUID `gorm:"type:uuid;size:36;primaryKey:true;index;not null;default:gen_random_uuid();" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null;default:now();" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null;default:now();" json:"updated_at"`
	Name string `gorm:"type:varchar(100);size:100;not null;default:null;" json:"name"`
	StateCode string `gorm:"type:varchar(5);size:5;not null;default:null;index;unique;" json:"state_code"`
	AbbreviationCode string `gorm:"type:varchar(5);size:5;not null;default:null;index;unique;" json:"abbreviation_code"`
}

func (l State) TableName() string {
	return "state"
}
