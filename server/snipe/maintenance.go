package snipe

type MaintenanceTableData struct {
  Assets []Asset 
}

func NewMaintenanceTableData() MaintenanceTableData {
  return MaintenanceTableData {
    Assets: []Asset {
      NewAsset(
        1234,
        "123asd0p",
        "Test CB",
        "Campus 001",
        "Broken camera",
        "camera is blurry", 
        "2024-06-01",
        ""),

      NewAsset(
        5678, 
        "567hjk0p",
        "Test Laptop",
        "Campus 002",
        "Broken bottom cover",
        "right corner broke from hinge issue",
        "2024-06-01",
        ""),
    },
  }
}
