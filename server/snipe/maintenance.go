package snipe

type MaintenanceTableData struct {
  Assets []Asset 
}

func CreateMaintenanceTableData() MaintenanceTableData {
  return MaintenanceTableData {
    Assets: InRepairAssets(),
  }
}
