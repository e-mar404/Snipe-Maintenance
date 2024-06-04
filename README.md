# Snipe-Maintenance
Web app to manage maintenance of assets in snipe

## Requirements:

### maintenace management 
- when creating maintenance also set asset's status to be in repair
- update completion date on maintenance

---

### maintenance report
- pull up with these cols:
    - sn
    - asset tag
    - name
    - company/location
    - title
    - notes
    - start date
    - completion date

- option to download the reports + filter by:
    - active/inactive 
    - date range (of start time)

---

## Testing:

- [x] Env file exists in root (Not sure of if to keep env test)
- [x] Env file has API_KEY and URL
- [ ] CheckStatus() returns error message if existing (this might have to be done since snipe api always returns status code of 200 if connection was successful)
- [ ] MaintenanceList() takes in json response from api and returns parsed []Asset
- [ ] InRepairDevicesList() takes in []Asset and returns only assets without a completion date
