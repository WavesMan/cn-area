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
