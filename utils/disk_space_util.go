package utils

import (
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
)

func CalculateTotalSizeByClient(client models.Client, clinicId uint) uint {
	var size uint = 0
	totalSize := uint(0)

	// get images size
	database.GormDB.Table("images").
		Where("client_id = ?", client.ID).
		Select("IF(size IS NULL, 0, SUM(size)) as size").
		Scan(&size)
	totalSize += size

	//get videos size
	size = 0
	database.GormDB.Table("videos").
		Where("client_id = ?", client.ID).
		Select("IF(size IS NULL, 0, SUM(size)) as size").
		Scan(&size)
	totalSize += size

	//get heartbeat size
	size = 0
	database.GormDB.Table("heartbeats").
		Where("client_id = ?", client.ID).
		Select("IF(size IS NULL, 0, SUM(size)) as size").
		Scan(&size)
	totalSize += size


	return totalSize
}
