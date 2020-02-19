package tools

var appcode_cio map[string] *Appcode_CIO

func LoadAppcodeCIO() map[string] *Appcode_CIO  {
	var result = LoadCsvCfg("Appcode_CIO.csv",1)
	if result == nil{
		return  nil
	}
	appcode_cio = make(map[string] *Appcode_CIO)
	for _,record := range result.Records{
		APP_CODE := record.GetString("APP_CODE")
		inform := &Appcode_CIO{
			APP_CODE,
			record.GetString("CIO"),
		}
		appcode_cio[APP_CODE] = inform
	}
	return appcode_cio
}
