from cn_area import provinces, cities, districts, lookup, flatten


def test_provinces_count():
    assert len(provinces()) == 34


def test_flatten_count():
    assert len(flatten()) == 2852


def test_lookup_dongcheng():
    r = lookup("110101")
    assert r is not None
    assert r.district_name == "东城区"
    assert r.province_name == "北京市"


def test_lookup_hongkong():
    r = lookup("81")
    assert r is not None
    assert r.province_name == "香港特别行政区"
    assert r.city_code is None


def test_cities_inner_mongolia():
    c = cities("15")
    assert len(c) > 0
    assert c[0].name == "呼和浩特市"


def test_cities_beijing_empty():
    """直辖市无城市层"""
    assert len(cities("11")) == 0


def test_districts_hohhot():
    d = districts("1501")
    assert len(d) > 0
