package models

import (
"github.com/google/uuid"
"time"
)

type Cases struct {
	ID        uuid.UUID `gorm:"type:uuid;size:36;primaryKey:true;index;not null;default:gen_random_uuid();" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null;default:now();" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null;default:now();" json:"updated_at"`
	Status string `gorm:"type:varchar(100);size:100;not null;default:null;" json:"name"`
	State   *State `gorm:"foreignKey:StateID;not null;default:null;" json:"raw_video,omitempty"`
	Count int `gorm:"type:int;size:100;not null;default:null;" json:"count"`
}

func (l Cases) TableName() string {
	return "cases"
}
