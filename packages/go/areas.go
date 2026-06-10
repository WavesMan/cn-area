package cnarea

// Province 省级行政区
type Province struct {
	Code  string
	Name  string
	Type  string
	Cities []City
}

// City 地级市
type City struct {
	Code      string
	Name      string
	Type      string
	Districts []District
}

// District 区县级
type District struct {
	Code string
	Name string
	Type string
}

// Record 扁平化记录
type Record struct {
	ProvinceCode string
	ProvinceName string
	ProvinceType string
	CityCode     string
	CityName     string
	CityType     string
	DistrictCode string
	DistrictName string
	DistrictType string
}

// Provinces 获取全部 34 个省级行政区
func Provinces() []Province {
	return areas
}

// Cities 按省查询地级市（直辖市无城市层，返回 nil）
func Cities(provinceCode string) []City {
	for _, prov := range areas {
		if prov.Code == provinceCode {
			if len(prov.Cities) == 0 {
				return nil
			}
			// 直辖市的 Cities 实际是 District
			if len(prov.Cities[0].Districts) == 0 {
				return nil
			}
			return prov.Cities
		}
	}
	return nil
}

// Districts 按市查询区县
func Districts(cityCode string) []District {
	for _, prov := range areas {
		for _, city := range prov.Cities {
			if city.Code == cityCode {
				return city.Districts
			}
		}
	}
	return nil
}

// Lookup 按行政区划码精确定位
func Lookup(code string) (*Record, bool) {
	for _, prov := range areas {
		if prov.Code == code {
			return &Record{
				ProvinceCode: prov.Code,
				ProvinceName: prov.Name,
				ProvinceType: prov.Type,
			}, true
		}

		if len(prov.Cities) == 0 {
			continue
		}

		// 直辖市：Cities 实际是 District（无 Districts）
		isDirect := len(prov.Cities[0].Districts) == 0
		if isDirect {
			for _, dist := range prov.Cities {
				if dist.Code == code {
					return &Record{
						ProvinceCode: prov.Code,
						ProvinceName: prov.Name,
						ProvinceType: prov.Type,
						DistrictCode: dist.Code,
						DistrictName: dist.Name,
						DistrictType: dist.Type,
					}, true
				}
			}
			continue
		}

		// 普通省份：province → city → district
		for _, city := range prov.Cities {
			if city.Code == code {
				return &Record{
					ProvinceCode: prov.Code,
					ProvinceName: prov.Name,
					ProvinceType: prov.Type,
					CityCode:     city.Code,
					CityName:     city.Name,
					CityType:     city.Type,
				}, true
			}
			for _, dist := range city.Districts {
				if dist.Code == code {
					return &Record{
						ProvinceCode: prov.Code,
						ProvinceName: prov.Name,
						ProvinceType: prov.Type,
						CityCode:     city.Code,
						CityName:     city.Name,
						CityType:     city.Type,
						DistrictCode: dist.Code,
						DistrictName: dist.Name,
						DistrictType: dist.Type,
					}, true
				}
			}
		}
	}
	return nil, false
}

// Flatten 返回全部扁平记录
func Flatten() []Record {
	var result []Record
	for _, prov := range areas {
		if len(prov.Cities) == 0 {
			result = append(result, Record{
				ProvinceCode: prov.Code,
				ProvinceName: prov.Name,
				ProvinceType: prov.Type,
			})
			continue
		}
		for _, city := range prov.Cities {
			if len(city.Districts) == 0 {
				// 直辖市的 district
				result = append(result, Record{
					ProvinceCode: prov.Code,
					ProvinceName: prov.Name,
					ProvinceType: prov.Type,
					DistrictCode: city.Code,
					DistrictName: city.Name,
					DistrictType: city.Type,
				})
				continue
			}
			for _, dist := range city.Districts {
				result = append(result, Record{
					ProvinceCode: prov.Code,
					ProvinceName: prov.Name,
					ProvinceType: prov.Type,
					CityCode:     city.Code,
					CityName:     city.Name,
					CityType:     city.Type,
					DistrictCode: dist.Code,
					DistrictName: dist.Name,
					DistrictType: dist.Type,
				})
			}
		}
	}
	return result
}

// Search 按地区名称反查（精确匹配优先，模糊匹配兜底）
func Search(name string) []Record {
	if name == "" {
		return nil
	}
	all := Flatten()
	var exact []Record
	for _, r := range all {
		if r.ProvinceName == name || r.CityName == name || r.DistrictName == name {
			exact = append(exact, r)
		}
	}
	if len(exact) > 0 {
		return exact
	}
	var fuzzy []Record
	for _, r := range all {
		if contains(r.ProvinceName, name) || contains(r.CityName, name) || contains(r.DistrictName, name) {
			fuzzy = append(fuzzy, r)
		}
	}
	return fuzzy
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && len(substr) > 0 && findSubstring(s, substr))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
