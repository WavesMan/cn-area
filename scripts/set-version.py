#!/usr/bin/env python3
"""从 VERSION 文件同步版本号到各语言包"""

import json
import re
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
VERSION_FILE = ROOT / "VERSION"
TS_PKG = ROOT / "packages" / "typescript" / "package.json"
PY_PKG = ROOT / "packages" / "python" / "pyproject.toml"


def get_version() -> str:
    version = VERSION_FILE.read_text().strip()
    if not re.match(r"^\d+\.\d+\.\d+$", version):
        print(f"Invalid version format: {version}", file=sys.stderr)
        sys.exit(1)
    return version


def update_json(path: Path, version: str):
    data = json.loads(path.read_text(encoding="utf-8"))
    data["version"] = version
    path.write_text(json.dumps(data, indent=2, ensure_ascii=False) + "\n", encoding="utf-8")
    print(f"  Updated {path.relative_to(ROOT)}: {version}")


def update_toml(path: Path, version: str):
    content = path.read_text(encoding="utf-8")
    content = re.sub(r'^version = ".*"$', f'version = "{version}"', content, count=1, flags=re.MULTILINE)
    path.write_text(content, encoding="utf-8")
    print(f"  Updated {path.relative_to(ROOT)}: {version}")


def main():
    version = get_version()
    print(f"Setting version: {version}")
    update_json(TS_PKG, version)
    update_toml(PY_PKG, version)
    print("Done.")


if __name__ == "__main__":
    main()
