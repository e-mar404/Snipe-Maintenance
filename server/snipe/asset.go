package snipe

type Asset struct {
  AssetTag int
  SerialNumber string 
  Name string
  Company string
  Title string
  Notes string
  StartDate string
  CompletionDate string
}

func NewAsset (tag int, sn, name, company, title, notes, start, completion string) Asset {
  return Asset {
    AssetTag: tag,
    SerialNumber: sn,
    Name: name,
    Company: company,
    Title: title,
    Notes: notes,
    StartDate: start,
    CompletionDate: completion,
  }
}
