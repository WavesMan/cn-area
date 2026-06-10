"""中国行政区划数据包"""

from __future__ import annotations

from dataclasses import dataclass
from typing import Optional

from ._data import AREAS


@dataclass(frozen=True)
class Area:
    code: str
    name: str
    type: str


@dataclass(frozen=True)
class AreaRecord:
    province_code: str
    province_name: str
    province_type: str
    city_code: Optional[str] = None
    city_name: Optional[str] = None
    city_type: Optional[str] = None
    district_code: Optional[str] = None
    district_name: Optional[str] = None
    district_type: Optional[str] = None


def provinces() -> list[Area]:
    """获取全部 34 个省级行政区"""
    return [Area(code=p["code"], name=p["name"], type=p["type"]) for p in AREAS]


def cities(province_code: str) -> list[Area]:
    """按省查询地级市（直辖市无城市层，返回空）"""
    for prov in AREAS:
        if prov["code"] == province_code:
            children = prov.get("children", [])
            if not children:
                return []
            # 直辖市的 children 是 district，不是 city
            if "children" not in children[0]:
                return []
            return [Area(code=c["code"], name=c["name"], type=c["type"]) for c in children]
    return []


def districts(city_code: str) -> list[Area]:
    """按市查询区县"""
    for prov in AREAS:
        for city in prov.get("children", []):
            if city["code"] == city_code:
                return [
                    Area(code=d["code"], name=d["name"], type=d["type"])
                    for d in city.get("children", [])
                ]
    return []


def lookup(code: str) -> Optional[AreaRecord]:
    """按行政区划码精确定位"""
    for prov in AREAS:
        if prov["code"] == code:
            return AreaRecord(
                province_code=prov["code"],
                province_name=prov["name"],
                province_type=prov["type"],
            )

        children = prov.get("children", [])
        if not children:
            continue

        # 直辖市：children 直接是 district（无 city 层）
        is_direct = "children" not in children[0]
        if is_direct:
            for dist in children:
                if dist["code"] == code:
                    return AreaRecord(
                        province_code=prov["code"],
                        province_name=prov["name"],
                        province_type=prov["type"],
                        district_code=dist["code"],
                        district_name=dist["name"],
                        district_type=dist["type"],
                    )
            continue

        # 普通省份：province → city → district
        for city in children:
            if city["code"] == code:
                return AreaRecord(
                    province_code=prov["code"],
                    province_name=prov["name"],
                    province_type=prov["type"],
                    city_code=city["code"],
                    city_name=city["name"],
                    city_type=city["type"],
                )
            for dist in city.get("children", []):
                if dist["code"] == code:
                    return AreaRecord(
                        province_code=prov["code"],
                        province_name=prov["name"],
                        province_type=prov["type"],
                        city_code=city["code"],
                        city_name=city["name"],
                        city_type=city["type"],
                        district_code=dist["code"],
                        district_name=dist["name"],
                        district_type=dist["type"],
                    )
    return None


def flatten() -> list[AreaRecord]:
    """返回全部扁平记录"""
    result: list[AreaRecord] = []
    for prov in AREAS:
        children = prov.get("children", [])
        if not children:
            result.append(AreaRecord(
                province_code=prov["code"],
                province_name=prov["name"],
                province_type=prov["type"],
            ))
            continue
        for city in children:
            if "children" not in city:
                # 直辖市的 district
                result.append(AreaRecord(
                    province_code=prov["code"],
                    province_name=prov["name"],
                    province_type=prov["type"],
                    district_code=city["code"],
                    district_name=city["name"],
                    district_type=city["type"],
                ))
                continue
            for dist in city.get("children", []):
                result.append(AreaRecord(
                    province_code=prov["code"],
                    province_name=prov["name"],
                    province_type=prov["type"],
                    city_code=city["code"],
                    city_name=city["name"],
                    city_type=city["type"],
                    district_code=dist["code"],
                    district_name=dist["name"],
                    district_type=dist["type"],
                ))
    return result


def search(name: str) -> list[AreaRecord]:
    """按地区名称反查（精确匹配优先，模糊匹配兜底）"""
    if not name:
        return []
    all_records = flatten()
    exact = [
        r for r in all_records
        if r.province_name == name or r.city_name == name or r.district_name == name
    ]
    if exact:
        return exact
    return [
        r for r in all_records
        if name in r.province_name
        or (r.city_name and name in r.city_name)
        or (r.district_name and name in r.district_name)
    ]
