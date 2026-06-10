import { describe, it } from "node:test";
import assert from "node:assert";
import { provinces, cities, districts, lookup, flatten } from "./index.js";

describe("cn-area", () => {
  it("provinces() returns 34 items", () => {
    assert.strictEqual(provinces().length, 34);
  });

  it("flatten() returns 2852 items", () => {
    assert.strictEqual(flatten().length, 2852);
  });

  it("lookup('110101') returns 东城区", () => {
    const r = lookup("110101");
    assert.ok(r);
    assert.strictEqual(r!.districtName, "东城区");
    assert.strictEqual(r!.provinceName, "北京市");
  });

  it("lookup('81') returns 香港特别行政区", () => {
    const r = lookup("81");
    assert.ok(r);
    assert.strictEqual(r!.provinceName, "香港特别行政区");
    assert.strictEqual(r!.cityCode, undefined);
  });

  it("cities('15') returns Inner Mongolia cities", () => {
    const c = cities("15");
    assert.ok(c.length > 0);
    assert.strictEqual(c[0].name, "呼和浩特市");
  });

  it("cities('11') returns empty (直辖市)", () => {
    assert.strictEqual(cities("11").length, 0);
  });

  it("districts('1501') returns Hohhot districts", () => {
    const d = districts("1501");
    assert.ok(d.length > 0);
  });
});
