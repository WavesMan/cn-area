import { areas, AreaNode } from "./data.js";

export type { AreaNode };

export interface AreaRecord {
  provinceCode: string;
  provinceName: string;
  provinceType: string;
  cityCode?: string;
  cityName?: string;
  cityType?: string;
  districtCode?: string;
  districtName?: string;
  districtType?: string;
}

/** 获取全部 34 个省级行政区 */
export function provinces(): AreaNode[] {
  return areas;
}

/** 按省查询地级市（直辖市无城市层，返回空） */
export function cities(provinceCode: string): AreaNode[] {
  const prov = areas.find((a) => a.code === provinceCode);
  if (!prov?.children) return [];
  // 直辖市的 children 是 district，不是 city
  const first = prov.children[0];
  if (first && !first.children) return [];
  return prov.children;
}

/** 按市查询区县 */
export function districts(cityCode: string): AreaNode[] {
  for (const prov of areas) {
    if (!prov.children) continue;
    for (const city of prov.children) {
      if (city.code === cityCode) {
        return city.children ?? [];
      }
    }
  }
  return [];
}

/** 按行政区划码精确定位 */
export function lookup(code: string): AreaRecord | undefined {
  for (const prov of areas) {
    if (prov.code === code) {
      return {
        provinceCode: prov.code,
        provinceName: prov.name,
        provinceType: prov.type,
      };
    }
    if (!prov.children) continue;

    // 直辖市：children 直接是 district（无 city 层）
    const isDirectCity = prov.children.length > 0 && !prov.children[0].children;
    if (isDirectCity) {
      for (const dist of prov.children) {
        if (dist.code === code) {
          return {
            provinceCode: prov.code,
            provinceName: prov.name,
            provinceType: prov.type,
            districtCode: dist.code,
            districtName: dist.name,
            districtType: dist.type,
          };
        }
      }
      continue;
    }

    // 普通省份：province → city → district
    for (const city of prov.children) {
      if (city.code === code) {
        return {
          provinceCode: prov.code,
          provinceName: prov.name,
          provinceType: prov.type,
          cityCode: city.code,
          cityName: city.name,
          cityType: city.type,
        };
      }
      if (!city.children) continue;
      for (const dist of city.children) {
        if (dist.code === code) {
          return {
            provinceCode: prov.code,
            provinceName: prov.name,
            provinceType: prov.type,
            cityCode: city.code,
            cityName: city.name,
            cityType: city.type,
            districtCode: dist.code,
            districtName: dist.name,
            districtType: dist.type,
          };
        }
      }
    }
  }
  return undefined;
}

/** 返回全部扁平记录 */
export function flatten(): AreaRecord[] {
  const result: AreaRecord[] = [];
  for (const prov of areas) {
    if (!prov.children || prov.children.length === 0) {
      result.push({
        provinceCode: prov.code,
        provinceName: prov.name,
        provinceType: prov.type,
      });
      continue;
    }
    for (const city of prov.children) {
      if (!city.children || city.children.length === 0) {
        // 直辖市的 district 或无 district 的 city
        result.push({
          provinceCode: prov.code,
          provinceName: prov.name,
          provinceType: prov.type,
          districtCode: city.code,
          districtName: city.name,
          districtType: city.type,
        });
        continue;
      }
      for (const dist of city.children) {
        result.push({
          provinceCode: prov.code,
          provinceName: prov.name,
          provinceType: prov.type,
          cityCode: city.code,
          cityName: city.name,
          cityType: city.type,
          districtCode: dist.code,
          districtName: dist.name,
          districtType: dist.type,
        });
      }
    }
  }
  return result;
}
