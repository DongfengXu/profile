package main

import (
	"madaoQT/utils"
	"profile/tools"
)

var profile_cio map[string] *tools.Profile_CIO



func main()  {
	appcode_cio := tools.LoadAppcodeCIO()
	profile := tools.LoadProfile()
	profile_cio = make(map[string] *tools.Profile_CIO)
	for appcode, prof := range profile{
		if _,ok := appcode_cio[appcode];ok{
			cio := appcode_cio[appcode].CIO
			inform := &tools.Profile_CIO{
				appcode,
				prof.CONEXT,
				cio,
			}
			profile_cio[appcode] = inform
		}else{
			utils.Logger.Info("There is no the CIO information: " + appcode)
		}
	}

	tools.CreateCSV("result.csv",profile_cio)

}
