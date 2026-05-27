# Changelog

## 0.20.0 (2026-05-27)

Full Changelog: [v0.19.0...v0.20.0](https://github.com/boltz-bio/boltz-api-go/compare/v0.19.0...v0.20.0)

### ⚠ BREAKING CHANGES

* **compute-api:** rename pipeline route params to id

### Features

* [codex] Add ADME scoring for small molecule pipelines ([205f228](https://github.com/boltz-bio/boltz-api-go/commit/205f228930859100d0220aed79f24672351e2a30))
* [codex] Add Baseten ADME model and public prediction API ([caf0e4d](https://github.com/boltz-bio/boltz-api-go/commit/caf0e4d66aa3241f2bf4cafdc8b67730b74bf26d))
* [codex] Add Benchling workspace spending limits ([f063b0c](https://github.com/boltz-bio/boltz-api-go/commit/f063b0cb168276a5efeae6740b70eab4a1cfee59))
* [codex] Add custom MSA support to structure predictions ([fecafd7](https://github.com/boltz-bio/boltz-api-go/commit/fecafd7e02f2362dc5d0905de70c1c7573d511fb))
* Add Boltz2 templates to compute API ([b1bf5b7](https://github.com/boltz-bio/boltz-api-go/commit/b1bf5b789b4979d0b865b1bf1ae026e9509ac213))
* Address ADME API review followups ([ed0c95e](https://github.com/boltz-bio/boltz-api-go/commit/ed0c95efe0cc1511a95dfb795e78fa8b0237069a))
* **billing:** attribute usage by product category ([b6747ce](https://github.com/boltz-bio/boltz-api-go/commit/b6747ce8b23c206ee02198d60f5766b9b19c7fcc))
* **client:** optimize json encoder for internal types ([f014a10](https://github.com/boltz-bio/boltz-api-go/commit/f014a1039ee16b4c95de7832920292366f7833ad))
* **compute-api:** add auth context endpoint ([0921572](https://github.com/boltz-bio/boltz-api-go/commit/0921572068177df2e2e9029529a89dcc156ef06c))
* **compute-api:** add CLI version metadata endpoint ([c7cb0ce](https://github.com/boltz-bio/boltz-api-go/commit/c7cb0ceb58a5834a34ebcb7abe5dcadf442186cb))
* **compute-api:** add protein redesign endpoint ([86ecb23](https://github.com/boltz-bio/boltz-api-go/commit/86ecb23ffb0e8212297bdbd8e2c43b7781a0d69a))
* **compute-api:** add small molecule structure templates ([f39dc9d](https://github.com/boltz-bio/boltz-api-go/commit/f39dc9d83d411da1b9aa72f8271f2011b4c86bf8))
* **compute-api:** prefix structure-binding ids with sab_pred ([65e677f](https://github.com/boltz-bio/boltz-api-go/commit/65e677f3c7dd88ff5d25d83ae3fbbfb89848d344))
* **compute-api:** support non-binding target residues ([b1c05cc](https://github.com/boltz-bio/boltz-api-go/commit/b1c05cc97f2eb387cf80de4609c2dd58fb492998))
* **compute:** accept user OAuth bearer tokens ([c0f9879](https://github.com/boltz-bio/boltz-api-go/commit/c0f9879e1cf319e3e51f69a665880258266625fa))
* **compute:** filter workspaces by name prefix ([7a9df49](https://github.com/boltz-bio/boltz-api-go/commit/7a9df49e7b00886e5a936f8d674793814b75c520))
* **compute:** gate console backend by compute roles ([b60b172](https://github.com/boltz-bio/boltz-api-go/commit/b60b172b5323070f22e30717ab2293d866b760ea))
* **compute:** support curated binder specs ([3931d6a](https://github.com/boltz-bio/boltz-api-go/commit/3931d6a787e5b1e0adb57e447d866facb0831e3c))
* **go:** add default http client with timeout ([89409f6](https://github.com/boltz-bio/boltz-api-go/commit/89409f63c1bf48fc8ef658a96acb53ae4e7d5f54))
* rename cli command ([8b05561](https://github.com/boltz-bio/boltz-api-go/commit/8b055612f34bb1f4c06c4ba6b3bbf6a38b0b826b))
* support setting headers via env ([0100a89](https://github.com/boltz-bio/boltz-api-go/commit/0100a89c2ef270616495cd2ae0aceea6caf87398))


### Bug Fixes

* **backend:** narrow hydrophobic set to "very hydrophobic" residues ([ca068d3](https://github.com/boltz-bio/boltz-api-go/commit/ca068d38856bd5c6a9305bc38f4f2a9f6f8d5bcb))
* **compute-api:** default omitted polymer modifications ([994bfd2](https://github.com/boltz-bio/boltz-api-go/commit/994bfd2bfcaa7d0e3cba20ba183d022a6f849e04))
* **compute-api:** display requested estimate units ([a4e9fe6](https://github.com/boltz-bio/boltz-api-go/commit/a4e9fe620f7d26c2df0d3e73b009e8d62b26d12b))
* **compute-api:** gate pocket conditioning schemas ([bfa741c](https://github.com/boltz-bio/boltz-api-go/commit/bfa741cb06253807289fbd701f8d95c6c8ab358e))
* **compute-api:** hide ADME from small-molecule results ([fc35230](https://github.com/boltz-bio/boltz-api-go/commit/fc35230fd3b2274b692012320b657c6008db75ca))
* **compute-api:** make design motifs optional ([60aba58](https://github.com/boltz-bio/boltz-api-go/commit/60aba58bd5da101dcbb02b5c5d5be7ad8d4c0078))
* **compute-api:** polish reference docs for CLI and auth ([d352953](https://github.com/boltz-bio/boltz-api-go/commit/d352953e7299887410935531aef64102a3eef2a6))
* **compute-api:** reject idempotency key input mismatches ([cc42427](https://github.com/boltz-bio/boltz-api-go/commit/cc424279e6ee96bc28e33d504a86479a27b82900))
* **compute-api:** replace unsafe application schema ([71f52bd](https://github.com/boltz-bio/boltz-api-go/commit/71f52bdd0bb889a712cdde0a4fe473b81683e3c6))
* **compute-api:** support SM target bonds and constraints ([2b33f46](https://github.com/boltz-bio/boltz-api-go/commit/2b33f46dd870c3cc4da57cfe56f6ed2bd00bdda3))
* **go:** avoid panic when http.DefaultTransport is wrapped ([cb6d66c](https://github.com/boltz-bio/boltz-api-go/commit/cb6d66c6bc063fed67380ae6e126b7ba31cca3c4))
* **sdk:** serialize usage arrays as repeated params ([dfd269d](https://github.com/boltz-bio/boltz-api-go/commit/dfd269d182c3b7fe42b4c2228248a64b64532340))
* stainless changes ([efde456](https://github.com/boltz-bio/boltz-api-go/commit/efde456a842a9989d0b587821fa0019c347ec3a5))
* **stainless:** expose auth context as auth me ([1354d93](https://github.com/boltz-bio/boltz-api-go/commit/1354d93aa9840c38637dd96b18e092f517ed773c))
* use friendly Go SDK package name ([c04b501](https://github.com/boltz-bio/boltz-api-go/commit/c04b50111928ada106ea9492965a89b5a9e8cdc1))


### Chores

* avoid embedding reflect.Type for dead code elimination ([a50f0e6](https://github.com/boltz-bio/boltz-api-go/commit/a50f0e618d961567656c3c2b3cb03ff7411813eb))
* configure new SDK language ([2270974](https://github.com/boltz-bio/boltz-api-go/commit/22709743fc68132e1f53e4b9a22f8343a0d36dbf))
* **internal:** more robust bootstrap script ([58962d2](https://github.com/boltz-bio/boltz-api-go/commit/58962d21c88ec5f9ecc83a7a2478ac992c42f137))
* redact api-key headers in debug logs ([6d94030](https://github.com/boltz-bio/boltz-api-go/commit/6d9403072095bf7671925f6c06f6ddb910b64512))
* **typebox:** upgrade backend schemas to v1 ([f357a7d](https://github.com/boltz-bio/boltz-api-go/commit/f357a7dda08584dcfa7d40813d4df07c794d3ac3))
* update SDK settings ([b7d10e6](https://github.com/boltz-bio/boltz-api-go/commit/b7d10e621a02df387122c5f69f6e0975bcc88d4f))
* update SDK settings ([ed32489](https://github.com/boltz-bio/boltz-api-go/commit/ed324898f9ba836625a19dc8b514d81c2d7203e5))


### Refactors

* **compute-api:** rename pipeline route params to id ([f10b31e](https://github.com/boltz-bio/boltz-api-go/commit/f10b31ea983f2b50804ce84e1d1a1c04dbe789cd))

## 0.19.0 (2026-05-27)

Full Changelog: [v0.18.0...v0.19.0](https://github.com/boltz-bio/boltz-api-go/compare/v0.18.0...v0.19.0)

### Features

* [codex] Add custom MSA support to structure predictions ([fecafd7](https://github.com/boltz-bio/boltz-api-go/commit/fecafd7e02f2362dc5d0905de70c1c7573d511fb))

## 0.18.0 (2026-05-26)

Full Changelog: [v0.17.1...v0.18.0](https://github.com/boltz-bio/boltz-api-go/compare/v0.17.1...v0.18.0)

### Features

* Add Boltz2 templates to compute API ([ddb1515](https://github.com/boltz-bio/boltz-api-go/commit/ddb1515e0609151532146b8b5b2b1b82e36d1fb2))

## 0.17.1 (2026-05-26)

Full Changelog: [v0.17.0...v0.17.1](https://github.com/boltz-bio/boltz-api-go/compare/v0.17.0...v0.17.1)

### Bug Fixes

* **compute-api:** gate pocket conditioning schemas ([99d8795](https://github.com/boltz-bio/boltz-api-go/commit/99d8795879ff2e4fa6e2b9fb29afcb44c1fa7584))

## 0.17.0 (2026-05-26)

Full Changelog: [v0.16.0...v0.17.0](https://github.com/boltz-bio/boltz-api-go/compare/v0.16.0...v0.17.0)

### Features

* **compute-api:** add protein redesign endpoint ([be2a559](https://github.com/boltz-bio/boltz-api-go/commit/be2a5595e8aa8d7bb13895dbbb7e9fa8c18f1931))

## 0.16.0 (2026-05-25)

Full Changelog: [v0.15.2...v0.16.0](https://github.com/boltz-bio/boltz-api-go/compare/v0.15.2...v0.16.0)

### Features

* **compute-api:** add small molecule structure templates ([15040a8](https://github.com/boltz-bio/boltz-api-go/commit/15040a8ef39b602082ab6085dd9481cc5c06974c))

## 0.15.2 (2026-05-25)

Full Changelog: [v0.15.1...v0.15.2](https://github.com/boltz-bio/boltz-api-go/compare/v0.15.1...v0.15.2)

### Chores

* **typebox:** upgrade backend schemas to v1 ([aa697d1](https://github.com/boltz-bio/boltz-api-go/commit/aa697d106ab682cfea097709c1dd52a988c6e619))

## 0.15.1 (2026-05-22)

Full Changelog: [v0.15.0...v0.15.1](https://github.com/boltz-bio/boltz-api-go/compare/v0.15.0...v0.15.1)

### Bug Fixes

* **compute-api:** reject idempotency key input mismatches ([1341344](https://github.com/boltz-bio/boltz-api-go/commit/1341344ca0778575093a37ec56ee133c0548294b))

## 0.15.0 (2026-05-15)

Full Changelog: [v0.14.0...v0.15.0](https://github.com/boltz-bio/boltz-api-go/compare/v0.14.0...v0.15.0)

### Features

* [codex] Add Benchling workspace spending limits ([fb1fb3e](https://github.com/boltz-bio/boltz-api-go/commit/fb1fb3e576c0f8759ad15c94d647509a8af0e534))

## 0.14.0 (2026-05-14)

Full Changelog: [v0.13.0...v0.14.0](https://github.com/boltz-bio/boltz-api-go/compare/v0.13.0...v0.14.0)

### Features

* **client:** optimize json encoder for internal types ([a5ea7a8](https://github.com/boltz-bio/boltz-api-go/commit/a5ea7a80fd0daccedbbcd2bfa2d879e7ce3b74b3))

## 0.13.0 (2026-05-13)

Full Changelog: [v0.12.1...v0.13.0](https://github.com/boltz-bio/boltz-api-go/compare/v0.12.1...v0.13.0)

### Features

* **compute:** filter workspaces by name prefix ([9b31722](https://github.com/boltz-bio/boltz-api-go/commit/9b31722585c64904f6c7c3c93309575af4b5a19f))

## 0.12.1 (2026-05-13)

Full Changelog: [v0.12.0...v0.12.1](https://github.com/boltz-bio/boltz-api-go/compare/v0.12.0...v0.12.1)

## 0.12.0 (2026-05-11)

Full Changelog: [v0.11.0...v0.12.0](https://github.com/boltz-bio/boltz-api-go/compare/v0.11.0...v0.12.0)

### Features

* **compute-api:** support non-binding target residues ([2da90d9](https://github.com/boltz-bio/boltz-api-go/commit/2da90d983e07cbb25de25e0f36994c9753caa54b))

## 0.11.0 (2026-05-11)

Full Changelog: [v0.10.4...v0.11.0](https://github.com/boltz-bio/boltz-api-go/compare/v0.10.4...v0.11.0)

### Features

* [codex] Add ADME scoring for small molecule pipelines ([e49dd61](https://github.com/boltz-bio/boltz-api-go/commit/e49dd611358263f8e103c2b1bca62f520a0e7a9f))
* Address ADME API review followups ([664fd80](https://github.com/boltz-bio/boltz-api-go/commit/664fd8044ab5ae5ce27532a7dcbf77d06f9fa8c0))

## 0.10.4 (2026-05-08)

Full Changelog: [v0.10.3...v0.10.4](https://github.com/boltz-bio/boltz-api-go/compare/v0.10.3...v0.10.4)

### Bug Fixes

* **go:** avoid panic when http.DefaultTransport is wrapped ([7a0abf1](https://github.com/boltz-bio/boltz-api-go/commit/7a0abf1c1bdf00e16972c9ba10a831a68f017572))


### Chores

* redact api-key headers in debug logs ([3c8aff8](https://github.com/boltz-bio/boltz-api-go/commit/3c8aff8e21d896b63d11305a245bc131851f1be7))

## 0.10.3 (2026-05-02)

Full Changelog: [v0.10.2...v0.10.3](https://github.com/boltz-bio/boltz-api-go/compare/v0.10.2...v0.10.3)

### Bug Fixes

* stainless changes ([efde456](https://github.com/boltz-bio/boltz-api-go/commit/efde456a842a9989d0b587821fa0019c347ec3a5))


### Chores

* update SDK settings ([b7d10e6](https://github.com/boltz-bio/boltz-api-go/commit/b7d10e621a02df387122c5f69f6e0975bcc88d4f))

## 0.10.2 (2026-05-01)

Full Changelog: [v0.10.1...v0.10.2](https://github.com/boltz-bio/boltz-compute-api-go/compare/v0.10.1...v0.10.2)

### Bug Fixes

* **compute-api:** polish reference docs for CLI and auth ([424bbe4](https://github.com/boltz-bio/boltz-compute-api-go/commit/424bbe4c38802be699e5b33a915e2114c75bef34))

## 0.10.1 (2026-05-01)

Full Changelog: [v0.10.0...v0.10.1](https://github.com/boltz-bio/boltz-compute-api-go/compare/v0.10.0...v0.10.1)

### Chores

* avoid embedding reflect.Type for dead code elimination ([e1f9b86](https://github.com/boltz-bio/boltz-compute-api-go/commit/e1f9b860cc167f586c1bdfb0d32fcac66044bfb9))

## 0.10.0 (2026-04-30)

Full Changelog: [v0.9.1...v0.10.0](https://github.com/boltz-bio/boltz-compute-api-go/compare/v0.9.1...v0.10.0)

### Features

* **compute:** gate console backend by compute roles ([a6d97b2](https://github.com/boltz-bio/boltz-compute-api-go/commit/a6d97b242127f7009d17122bc30867954931ac02))

## 0.9.1 (2026-04-30)

Full Changelog: [v0.9.0...v0.9.1](https://github.com/boltz-bio/boltz-compute-api-go/compare/v0.9.0...v0.9.1)

### Bug Fixes

* **compute-api:** display requested estimate units ([ae8169a](https://github.com/boltz-bio/boltz-compute-api-go/commit/ae8169a183802ccd39894bd5605bd2b616d220de))
* **sdk:** serialize usage arrays as repeated params ([6ebd332](https://github.com/boltz-bio/boltz-compute-api-go/commit/6ebd332f02e728488e0582eb21320f3b369cb027))

## 0.9.0 (2026-04-29)

Full Changelog: [v0.8.0...v0.9.0](https://github.com/boltz-bio/boltz-compute-api-go/compare/v0.8.0...v0.9.0)

### Features

* **compute:** support curated binder specs ([2e23387](https://github.com/boltz-bio/boltz-compute-api-go/commit/2e23387a4a3a5722ca47aedfa8f90034f21d8de9))

## 0.8.0 (2026-04-28)

Full Changelog: [v0.7.0...v0.8.0](https://github.com/boltz-bio/boltz-compute-api-go/compare/v0.7.0...v0.8.0)

### Features

* **go:** add default http client with timeout ([1c3b90e](https://github.com/boltz-bio/boltz-compute-api-go/commit/1c3b90edee209ec69ea6e0994916b40a98d8dc59))
* support setting headers via env ([db9a09c](https://github.com/boltz-bio/boltz-compute-api-go/commit/db9a09cd268a3e790fc6473adeadde3364548361))

## 0.7.0 (2026-04-28)

Full Changelog: [v0.6.1...v0.7.0](https://github.com/boltz-bio/boltz-compute-api-go/compare/v0.6.1...v0.7.0)

### Features

* **compute-api:** add CLI version metadata endpoint ([ac8b4f7](https://github.com/boltz-bio/boltz-compute-api-go/commit/ac8b4f735a3f09cdc0fbe36b6f74e3f66837eb03))

## 0.6.1 (2026-04-27)

Full Changelog: [v0.6.0...v0.6.1](https://github.com/boltz-bio/boltz-compute-api-go/compare/v0.6.0...v0.6.1)

### Bug Fixes

* **stainless:** expose auth context as auth me ([1354d93](https://github.com/boltz-bio/boltz-compute-api-go/commit/1354d93aa9840c38637dd96b18e092f517ed773c))

## 0.6.0 (2026-04-27)

Full Changelog: [v0.5.2...v0.6.0](https://github.com/boltz-bio/boltz-compute-api-go/compare/v0.5.2...v0.6.0)

### Features

* **compute-api:** add auth context endpoint ([9dedfa5](https://github.com/boltz-bio/boltz-compute-api-go/commit/9dedfa5b38425d603bd3471d0379a7b20e3e2b28))


### Bug Fixes

* **compute-api:** make design motifs optional ([8c7a79e](https://github.com/boltz-bio/boltz-compute-api-go/commit/8c7a79e5b81c80ab65d29d5ababf883c1348d43b))

## 0.5.2 (2026-04-24)

Full Changelog: [v0.5.1...v0.5.2](https://github.com/boltz-bio/boltz-compute-api-go/compare/v0.5.1...v0.5.2)

### Bug Fixes

* **backend:** narrow hydrophobic set to "very hydrophobic" residues ([0ef6acd](https://github.com/boltz-bio/boltz-compute-api-go/commit/0ef6acdd204391bfe971ee5c5d2f205581a6390c))

## 0.5.1 (2026-04-23)

Full Changelog: [v0.5.0...v0.5.1](https://github.com/boltz-bio/boltz-compute-api-go/compare/v0.5.0...v0.5.1)

### Chores

* **internal:** more robust bootstrap script ([b0260a6](https://github.com/boltz-bio/boltz-compute-api-go/commit/b0260a685a2eefb186ca242282f9a2eef12f1774))

## 0.5.0 (2026-04-23)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/boltz-bio/boltz-compute-api-go/compare/v0.4.0...v0.5.0)

### Features

* **compute-api:** prefix structure-binding ids with sab_pred ([65e677f](https://github.com/boltz-bio/boltz-compute-api-go/commit/65e677f3c7dd88ff5d25d83ae3fbbfb89848d344))


### Bug Fixes

* **compute-api:** hide ADME from small-molecule results ([fc35230](https://github.com/boltz-bio/boltz-compute-api-go/commit/fc35230fd3b2274b692012320b657c6008db75ca))

## 0.4.0 (2026-04-22)

Full Changelog: [v0.3.2...v0.4.0](https://github.com/boltz-bio/boltz-compute-api-go/compare/v0.3.2...v0.4.0)

### Features

* [codex] Add Baseten ADME model and public prediction API ([09b2516](https://github.com/boltz-bio/boltz-compute-api-go/commit/09b251663c6ed1181dee6ab651f17f8529b50528))

## 0.3.2 (2026-04-22)

Full Changelog: [v0.3.1...v0.3.2](https://github.com/boltz-bio/boltz-compute-api-go/compare/v0.3.1...v0.3.2)

### Bug Fixes

* **compute-api:** replace unsafe application schema ([71f52bd](https://github.com/boltz-bio/boltz-compute-api-go/commit/71f52bdd0bb889a712cdde0a4fe473b81683e3c6))

## 0.3.1 (2026-04-22)

Full Changelog: [v0.3.0...v0.3.1](https://github.com/boltz-bio/boltz-compute-api-go/compare/v0.3.0...v0.3.1)

### Bug Fixes

* **compute-api:** default omitted polymer modifications ([994bfd2](https://github.com/boltz-bio/boltz-compute-api-go/commit/994bfd2bfcaa7d0e3cba20ba183d022a6f849e04))

## 0.3.0 (2026-04-21)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/boltz-bio/boltz-compute-api-go/compare/v0.2.0...v0.3.0)

### ⚠ BREAKING CHANGES

* **compute-api:** rename pipeline route params to id

### Features

* **compute:** accept user OAuth bearer tokens ([c0f9879](https://github.com/boltz-bio/boltz-compute-api-go/commit/c0f9879e1cf319e3e51f69a665880258266625fa))


### Refactors

* **compute-api:** rename pipeline route params to id ([f10b31e](https://github.com/boltz-bio/boltz-compute-api-go/commit/f10b31ea983f2b50804ce84e1d1a1c04dbe789cd))

## 0.2.0 (2026-04-21)

Full Changelog: [v0.1.1...v0.2.0](https://github.com/boltz-bio/boltz-compute-api-go/compare/v0.1.1...v0.2.0)

### Features

* **billing:** attribute usage by product category ([687d333](https://github.com/boltz-bio/boltz-compute-api-go/commit/687d333495b081984bf70549db8b5bd06238e437))

## 0.1.1 (2026-04-21)

Full Changelog: [v0.1.0...v0.1.1](https://github.com/boltz-bio/boltz-compute-api-go/compare/v0.1.0...v0.1.1)

### Bug Fixes

* **compute-api:** support SM target bonds and constraints ([2b33f46](https://github.com/boltz-bio/boltz-compute-api-go/commit/2b33f46dd870c3cc4da57cfe56f6ed2bd00bdda3))
* use friendly Go SDK package name ([c04b501](https://github.com/boltz-bio/boltz-compute-api-go/commit/c04b50111928ada106ea9492965a89b5a9e8cdc1))

## 0.1.0 (2026-04-20)

Full Changelog: [v0.0.2...v0.1.0](https://github.com/boltz-bio/boltz-compute-api-go/compare/v0.0.2...v0.1.0)

### Features

* rename cli command ([8b05561](https://github.com/boltz-bio/boltz-compute-api-go/commit/8b055612f34bb1f4c06c4ba6b3bbf6a38b0b826b))

## 0.0.2 (2026-04-20)

Full Changelog: [v0.0.1...v0.0.2](https://github.com/boltz-bio/boltz-compute-api-go/compare/v0.0.1...v0.0.2)

### Chores

* configure new SDK language ([2270974](https://github.com/boltz-bio/boltz-compute-api-go/commit/22709743fc68132e1f53e4b9a22f8343a0d36dbf))
* update SDK settings ([96e2f2f](https://github.com/boltz-bio/boltz-compute-api-go/commit/96e2f2f3ef7eb6c4c56ad3c5b4b52b47c18933ff))
