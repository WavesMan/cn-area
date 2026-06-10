#!/usr/bin/env python3
"""将 area-code.csv 转换为 data/areas.json（三级树形结构）"""

import csv
import json
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
CSV_PATH = ROOT / "area-code.csv"
JSON_PATH = ROOT / "data" / "areas.json"


def convert():
    provinces: dict[str, dict] = {}

    with open(CSV_PATH, encoding="utf-8-sig") as f:
        reader = csv.DictReader(f)
        for row in reader:
            p_code = row["province_code"].strip()
            p_name = row["province_name"].strip()
            p_type = row["province_type"].strip()
            c_code = row["city_code"].strip()
            c_name = row["city_name"].strip()
            c_type = row["city_type"].strip()
            d_code = row["district_code"].strip()
            d_name = row["district_name"].strip()
            d_type = row["district_type"].strip()

            # 省级节点
            if p_code not in provinces:
                provinces[p_code] = {
                    "code": p_code,
                    "name": p_name,
                    "type": p_type,
                    "children": [],
                }
            prov = provinces[p_code]

            # 直辖市 / 特别行政区：无 city 层，district 直接挂省级
            if not c_code:
                if d_code:
                    prov["children"].append(
                        {"code": d_code, "name": d_name, "type": d_type}
                    )
                continue

            # 普通省份：查找或创建城市节点
            city = None
            for child in prov["children"]:
                if child["code"] == c_code:
                    city = child
                    break
            if city is None:
                city = {"code": c_code, "name": c_name, "type": c_type, "children": []}
                prov["children"].append(city)

            # 区县级
            if d_code:
                city["children"].append(
                    {"code": d_code, "name": d_name, "type": d_type}
                )

    result = list(provinces.values())
    JSON_PATH.parent.mkdir(parents=True, exist_ok=True)
    with open(JSON_PATH, "w", encoding="utf-8") as f:
        json.dump(result, f, ensure_ascii=False, indent=2)

    # 统计
    total_districts = sum(
        len(city.get("children", []))
        for prov in result
        for city in prov.get("children", [])
    ) + sum(
        len([c for c in prov["children"] if "children" not in c])
        for prov in result
    )
    print(f"生成完成: {len(result)} 个省级 → {JSON_PATH}")


if __name__ == "__main__":
    convert()
