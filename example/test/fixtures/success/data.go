package success

import "github.com/cjwcjswo/dbunit"

func Before() dbunit.FixtureData {
	return dbunit.FixtureData{
		"TB_USER": {
			{
				"name":    "cjwoov",
				"age":     28,
				"country": "korea",
			},
			{
				"name":    "battlecook",
				"age":     34,
				"country": "korea",
			},
		},
	}
}
func After() dbunit.FixtureData {
	return dbunit.FixtureData{
		"TB_USER": {
			{
				"name": "cjwoov",
				"age":  29,
			},
			{
				"name":    "battlecook",
				"age":     34,
				"country": "korea",
			},
		},
	}
}
