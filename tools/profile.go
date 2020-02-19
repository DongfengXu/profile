package tools

var profile map[string] *Profile

func LoadProfile() map[string] *Profile  {
	var result = LoadCsvCfg("profile.csv",1)
	if result == nil{
		return  nil
	}
	profile = make(map[string] *Profile)
	for _,record := range result.Records{
		APP_CODE := record.GetString("APP_CODE")
		inform := &Profile{
			APP_CODE,
			record.GetString("CONTEXT"),
		}
		profile[APP_CODE] = inform
	}

	return profile
}