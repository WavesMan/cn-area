package cnarea

import "testing"

func TestProvinces(t *testing.T) {
	if got := Provinces(); len(got) != 34 {
		t.Errorf("Provinces() = %d items, want 34", len(got))
	}
}

func TestFlatten(t *testing.T) {
	if got := Flatten(); len(got) != 2852 {
		t.Errorf("Flatten() = %d items, want 2852", len(got))
	}
}

func TestLookupDongcheng(t *testing.T) {
	r, ok := Lookup("110101")
	if !ok {
		t.Fatal("Lookup('110101') not found")
	}
	if r.DistrictName != "东城区" {
		t.Errorf("DistrictName = %q, want %q", r.DistrictName, "东城区")
	}
	if r.ProvinceName != "北京市" {
		t.Errorf("ProvinceName = %q, want %q", r.ProvinceName, "北京市")
	}
}

func TestLookupHongkong(t *testing.T) {
	r, ok := Lookup("81")
	if !ok {
		t.Fatal("Lookup('81') not found")
	}
	if r.ProvinceName != "香港特别行政区" {
		t.Errorf("ProvinceName = %q, want %q", r.ProvinceName, "香港特别行政区")
	}
	if r.CityCode != "" {
		t.Errorf("CityCode = %q, want empty", r.CityCode)
	}
}

func TestCitiesInnerMongolia(t *testing.T) {
	c := Cities("15")
	if len(c) == 0 {
		t.Fatal("Cities('15') is empty")
	}
	if c[0].Name != "呼和浩特市" {
		t.Errorf("first city = %q, want %q", c[0].Name, "呼和浩特市")
	}
}

func TestCitiesBeijingEmpty(t *testing.T) {
	if got := Cities("11"); len(got) != 0 {
		t.Errorf("Cities('11') = %d items, want 0", len(got))
	}
}

func TestDistrictsHohhot(t *testing.T) {
	d := Districts("1501")
	if len(d) == 0 {
		t.Fatal("Districts('1501') is empty")
	}
}

func TestSearchExact(t *testing.T) {
	r := Search("东城区")
	if len(r) == 0 {
		t.Fatal("Search('东城区') is empty")
	}
	if r[0].DistrictName != "东城区" {
		t.Errorf("DistrictName = %q, want %q", r[0].DistrictName, "东城区")
	}
}

func TestSearchMultiple(t *testing.T) {
	r := Search("朝阳区")
	if len(r) <= 1 {
		t.Errorf("Search('朝阳区') = %d items, want > 1", len(r))
	}
}

func TestSearchEmpty(t *testing.T) {
	if got := Search(""); len(got) != 0 {
		t.Errorf("Search('') = %d items, want 0", len(got))
	}
}

func TestSearchFuzzy(t *testing.T) {
	r := Search("朝阳")
	if len(r) == 0 {
		t.Fatal("Search('朝阳') is empty")
	}
}
