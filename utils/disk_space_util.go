package utils

import (
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
)

func CalculateTotalSizeByClient(client models.Client, clinicId uint) uint {
	var size uint = 0
	totalSize := uint(0)

	// get images size
	database.GormDB.Model(models.Image{}).
		Where("client_id = ? AND clinic_id = ?", client.ID, clinicId).
		Select("IF(size IS NULL, 0, SUM(size)) as size").
		Scan(&size)
	totalSize += size

	//get videos size
	size = 0
	database.GormDB.Model(models.Video{}).
		Where("client_id = ? AND clinic_id = ?", client.ID, clinicId).
		Select("IF(size IS NULL, 0, SUM(size)) as size").
		Scan(&size)
	totalSize += size

	//get heartbeat size
	size = 0
	database.GormDB.Model(models.Heartbeat{}).
		Where("client_id = ? AND clinic_id = ?", client.ID, clinicId).
		Select("IF(size IS NULL, 0, SUM(size)) as size").
		Scan(&size)
	totalSize += size


	return totalSize
}
