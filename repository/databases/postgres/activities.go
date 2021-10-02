package postgres

import (
	"aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/activities"
	"aprian1337/thukul-service/repository"
	"aprian1337/thukul-service/repository/databases/records"
	"context"
	"gorm.io/gorm"
)

type ActivitiesRepository struct {
	ConnPostgres *gorm.DB
}

func NewPostgresActivitiesRepository(conn *gorm.DB) *ActivitiesRepository {
	return &ActivitiesRepository{
		ConnPostgres: conn,
	}
}
func (repo *ActivitiesRepository) ActivitiesGetAll(ctx context.Context, pocketId int) ([]activities.Domain, error) {
	var data []records.Activities
	if pocketId > 0 {
		err := repo.ConnPostgres.Find(&data, "pocket_id=?", pocketId)
		if err.Error != nil {
			return []activities.Domain{}, err.Error
		}
	} else {
		err := repo.ConnPostgres.Find(&data)
		if err.Error != nil {
			return []activities.Domain{}, err.Error
		}
	}

	return records.ActivitiesToListDomain(data), nil
}

func (repo *ActivitiesRepository) ActivitiesGetById(ctx context.Context, pocketId int, id int) (activities.Domain, error) {
	var data records.Activities
	err := repo.ConnPostgres.First(&data, "id=? AND pocket_id=?", id, pocketId)
	if err.Error != nil {
		return activities.Domain{}, err.Error
	}
	return data.ActivitiesToDomain(), nil
}

func (repo *ActivitiesRepository) ActivitiesCreate(ctx context.Context, domain activities.Domain, pocketId int) (activities.Domain, error) {
	data := records.ActivitiesFromDomain(domain)
	var pocket records.Pockets
	err := repo.ConnPostgres.First(&pocket, "id=?", pocketId)
	if err.Error != nil {
		return activities.Domain{}, businesses.ErrUserIdNotFound
	}
	err = repo.ConnPostgres.Create(&data)
	if err.Error != nil {
		return activities.Domain{}, err.Error
	}
	return data.ActivitiesToDomain(), nil
}

func (repo *ActivitiesRepository) ActivitiesUpdate(ctx context.Context, domain activities.Domain, pocketId int, id int) (activities.Domain, error) {
	data := records.Activities{}
	dataTemp := records.ActivitiesFromDomain(domain)
	err := repo.ConnPostgres.First(&data, "id=?", id)
	if err.Error != nil {
		return activities.Domain{}, err.Error
	}
	repo.ConnPostgres.Save(&dataTemp)
	return data.ActivitiesToDomain(), nil
}

func (repo *ActivitiesRepository) ActivitiesDelete(ctx context.Context, pocketId int, id int) (int64, error) {
	data := records.Activities{}
	err := repo.ConnPostgres.First(&data, "id=? AND pocket_id=?", id, pocketId)
	if err.Error != nil {
		return 0, err.Error
	}
	if err.RowsAffected == 0 {
		return 0, repository.ErrDataNotFound
	}

	err = repo.ConnPostgres.Delete(&data, id)
	if err.Error != nil {
		return 0, err.Error
	}
	return err.RowsAffected, nil
}

func (repo *ActivitiesRepository) ActivitiesGetTotal(ctx context.Context, userId int, pocketId int, kind string) (int64, error) {
	total := records.ActivitiesTotal{}
	activity := records.Activities{}
	err := repo.ConnPostgres.Model(activity).Select("id, sum(nominal) as total").Where("pocket_id = ? ", pocketId).Group("id").Having("type=?", kind).First(&total)
	if err.Error != nil {
		return 0, err.Error
	}
	return total.Total, nil
}
