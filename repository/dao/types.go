package dao

import (
	"database/sql"
	"encoding/json"
	"github.com/webook-project-go/webook-user/domain"
)

type UserEntity struct {
	ID       int64          `gorm:"primaryKey;autoIncrement"`
	Email    sql.NullString `gorm:"uniqueIndex;type:varchar(50)"`
	Password string         `gorm:"type:varchar(200)"`
	Phone    sql.NullString `gorm:"uniqueIndex;type:varchar(20)"`
	Name     string         `gorm:"type:varchar(10)"`
	Birthday string         `gorm:"type:date;default:NULL"`
	AboutMe  string         `gorm:"varchar(50)"`
	Ctime    int64
	Utime    int64
}

func (UserEntity) TableName() string {
	return "users"
}

func userDomainToEntity(dm *domain.User) UserEntity {
	return UserEntity{
		ID:       dm.Id,
		Name:     dm.Name,
		Password: dm.Password,
		Phone: sql.NullString{
			String: dm.Phone,
			Valid:  dm.Phone != "",
		},
		Email: sql.NullString{
			String: dm.Email,
			Valid:  dm.Email != "",
		},
		AboutMe:  dm.AboutMe,
		Birthday: dm.Birthday,
	}
}

func userEntityToDomain(ud *UserEntity) domain.User {
	return domain.User{
		Id:       ud.ID,
		Password: ud.Password,
		Name:     ud.Name,
		Phone:    ud.Phone.String,
		Email:    ud.Email.String,
		AboutMe:  ud.AboutMe,
		Birthday: ud.Birthday,
	}
}

type Oauth2BindingEntity struct {
	ID     int64      `gorm:"primary key;autoIncrement"`
	UserID int64      `gorm:"index"`
	User   UserEntity `gorm:"foreignKey:UserID;references:ID;OnDelete:CASCADE"`

	Provider        sql.NullString `gorm:"type:varchar(20);uniqueIndex:idx_provider_external_id"`
	ExternalID      sql.NullString `gorm:"type:varchar(50);uniqueIndex:idx_provider_external_id"`
	ProviderUnionID sql.NullString

	AccessToken  sql.NullString
	RefreshToken sql.NullString

	Ctime     int64
	Utime     int64
	ExtraInfo json.RawMessage `gorm:"type:json"`
}

func (Oauth2BindingEntity) TableName() string {
	return "oauth2BindingEntities"
}
func oauth2BindingDomainToEntity(dm *domain.Oauth2Binding) Oauth2BindingEntity {
	return Oauth2BindingEntity{
		ID:     dm.ID,
		UserID: dm.UserID,

		Provider: sql.NullString{
			String: dm.Provider,
			Valid:  dm.Provider != "",
		},
		ExternalID: sql.NullString{
			String: dm.ExternalID,
			Valid:  dm.ExternalID != "",
		},
		AccessToken: sql.NullString{
			String: dm.AccessToken,
			Valid:  dm.AccessToken != "",
		},
	}
}

func oauth2BindingEntityToDomain(entity *Oauth2BindingEntity) domain.Oauth2Binding {
	return domain.Oauth2Binding{
		ID:          entity.ID,
		UserID:      entity.UserID,
		ExternalID:  entity.ExternalID.String,
		Provider:    entity.Provider.String,
		AccessToken: entity.AccessToken.String,
	}
}
