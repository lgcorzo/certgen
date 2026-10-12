# lgcorzo/certgen

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Dark Gravity](https://img.shields.io/badge/Autonomous%20Factory-Dark%20Gravity-blueviolet?style=flat-square)](https://github.com/lgcorzo/rust_CACD_autonomous_factory)

> [!NOTE]
> **UPSTREAM STATUS & SOVEREIGN FORK PURPOSE:**
> Following upstream repository archiving and licensing shifts across the MinIO ecosystem, this repository is **actively maintained and hardened** under **@lgcorzo** as an essential infrastructure component of the **[Dark Gravity Autonomous CA/CD Factory](https://github.com/lgcorzo/rust_CACD_autonomous_factory)** ([Architecture Knowledge Base](https://notebook.google.com/notebook/116d5f03-2fb0-4fa7-bd7c-1c9849920abc)). Continuous security auditing, zero-CVE verification, and hardware SIMD optimizations are independently maintained to guarantee supply-chain sovereignty and operational continuity.

---

## 🛡️ Dark Gravity Factory: Technological Sovereignty & Security Rationale

### 1. Architectural Role in Dark Gravity
In the **Dark Gravity Autonomous CA/CD Software Factory (V7.2 / V7.3)**, Generates self-signed TLS/mTLS X.509 certificates for air-gapped, zero-trust storage deployment in Dark Gravity Kubernetes namespaces.

### 2. Supply-Chain Sovereignty & Independent Maintenance
- **Insulation from Upstream License Closures:** When upstream open-source projects close distributions or alter licensing terms, Dark Gravity enforces **semantic mirroring and sovereign forks** to guarantee total independence from commercial vendors.
- **Autonomous Security Verification & Zero-CVE Policy:** This repository is continuously audited for CVEs and vulnerabilities, maintaining a 100% clean secret scanning posture. Regressions and upstream bugs are resolved directly within the autonomous factory ecosystem.
- **Compliance & Provenance:** Ensures cryptographic integrity and auditable lineage in compliance with the **EU AI Act**, **SOC 2 Type II**, and **Hazitek / SPRI** regulatory frameworks.

---

# certgen

`certgen` is a simple tool to generate self-signed certificates, and provides SAN certificates with DNS and IP entries.

## Install

<details open="true"><summary><b><a name="binary-releases">Binary Releases</a></b></summary>

| OS       | ARCH    | Binary                                                                                               |
|:--------:|:-------:|:----------------------------------------------------------------------------------------------------:|
| Linux    | amd64   | [linux-amd64](https://github.com/minio/certgen/releases/latest/download/certgen-linux-amd64)         |
| Linux    | arm64   | [linux-arm64](https://github.com/minio/certgen/releases/latest/download/certgen-linux-arm64)         |
| Linux    | ppc64le | [linux-ppc64le](https://github.com/minio/certgen/releases/latest/download/certgen-linux-ppc64le)     |
| Linux    | s390x   | [linux-s390x](https://github.com/minio/certgen/releases/latest/download/certgen-linux-s390x)         |
| Apple M1 | arm64   | [darwin-arm64](https://github.com/minio/certgen/releases/latest/download/certgen-darwin-arm64)       |
| Apple    | amd64   | [darwin-amd64](https://github.com/minio/certgen/releases/latest/download/certgen-darwin-amd64)       |
| Windows  | amd64   | [windows-amd64](https://github.com/minio/certgen/releases/latest/download/certgen-windows-amd64.exe) |
| Windows  | arm64   | [windows-amd64](https://github.com/minio/certgen/releases/latest/download/certgen-windows-arm64.exe) |
| FreeBSD  | amd64   | [freebsd-amd64](https://github.com/minio/certgen/releases/latest/download/certgen-freebsd-amd64)     |
| FreeBSD  | arm64   | [freebsd-amd64](https://github.com/minio/certgen/releases/latest/download/certgen-freebsd-arm64)     |

Download [`certgen`](https://github.com/minio/certgen/releases/latest) for your specific operating system and platform.

## Example (server)

```sh
certgen -host "127.0.0.1,localhost"

Created a new certificate 'public.crt', 'private.key' valid for the following names 📜
 - "127.0.0.1"
 - "localhost"
```

## Example (client)

```sh
certgen -client -host "localhost"

Created a new certificate 'client.crt', 'client.key' valid for the following names 📜
 - "localhost"
```
