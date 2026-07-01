# Changelog

## 0.32.1 (2026-07-01)

Full Changelog: [v0.32.0...v0.32.1](https://github.com/boltz-bio/boltz-api-go/compare/v0.32.0...v0.32.1)

### Documentation

* **api:** remove em-dashes from public-facing compute API docs ([165a4a5](https://github.com/boltz-bio/boltz-api-go/commit/165a4a54dc753c3ea60dd85299a3a9524e4a4f6c))

## 0.32.0 (2026-06-26)

Full Changelog: [v0.31.0...v0.32.0](https://github.com/boltz-bio/boltz-api-go/compare/v0.31.0...v0.32.0)

### Features

* **share-links:** map share_links resource for SDK + CLI generation ([fb4d092](https://github.com/boltz-bio/boltz-api-go/commit/fb4d092e1a3a64461690d9a40df6e483ff02af00))

## 0.31.0 (2026-06-25)

Full Changelog: [v0.30.0...v0.31.0](https://github.com/boltz-bio/boltz-api-go/compare/v0.30.0...v0.31.0)

### Features

* **affinity:** raise ligand heavy atom cap to 100 ([565e0a4](https://github.com/boltz-bio/boltz-api-go/commit/565e0a4770b8702986da4cdab4af042eb9a78d8b))

## 0.30.0 (2026-06-24)

Full Changelog: [v0.29.0...v0.30.0](https://github.com/boltz-bio/boltz-api-go/compare/v0.29.0...v0.30.0)

### Features

* **compute-api:** add public screen resume endpoints ([6c4668a](https://github.com/boltz-bio/boltz-api-go/commit/6c4668a20475d758181bf54a2de69698ffaa145e))

## 0.29.0 (2026-06-24)

Full Changelog: [v0.28.0...v0.29.0](https://github.com/boltz-bio/boltz-api-go/compare/v0.28.0...v0.29.0)

### Features

* **compute-api:** expose design resume endpoints ([86baeb6](https://github.com/boltz-bio/boltz-api-go/commit/86baeb6e3f1b52ef372a391441ad6d20dc593db4))

## 0.28.0 (2026-06-22)

Full Changelog: [v0.27.0...v0.28.0](https://github.com/boltz-bio/boltz-api-go/compare/v0.27.0...v0.28.0)

### Features

* **compute-api:** filter list-results by id ([096a48f](https://github.com/boltz-bio/boltz-api-go/commit/096a48fc05b893d8824c812dc23da2c15de18d17))

## 0.27.0 (2026-06-22)

Full Changelog: [v0.26.1...v0.27.0](https://github.com/boltz-bio/boltz-api-go/compare/v0.26.1...v0.27.0)

### Features

* **compute-api:** include organization_name in /v1/auth/me ([0835ff9](https://github.com/boltz-bio/boltz-api-go/commit/0835ff95c7be67c355a9537420ea57e7463f3422))

## 0.26.1 (2026-06-16)

Full Changelog: [v0.26.0...v0.26.1](https://github.com/boltz-bio/boltz-api-go/compare/v0.26.0...v0.26.1)

### Documentation

* **api:** fix runnable predictions example + document request limits ([10f50d7](https://github.com/boltz-bio/boltz-api-go/commit/10f50d72af0a45c60b5dc74c2f3b547731b5fc06))
* **api:** use ASCII hyphens in limit descriptions ([217be1b](https://github.com/boltz-bio/boltz-api-go/commit/217be1be8fecb14a79a51529be5c68b94daf54c2))

## 0.26.0 (2026-06-15)

Full Changelog: [v0.25.1...v0.26.0](https://github.com/boltz-bio/boltz-api-go/compare/v0.25.1...v0.26.0)

### Features

* **compute-api:** add ADME estimate-cost endpoint ([847726f](https://github.com/boltz-bio/boltz-api-go/commit/847726f63531c7a117b24859287e4aadb481818a))

## 0.25.1 (2026-06-15)

Full Changelog: [v0.25.0...v0.25.1](https://github.com/boltz-bio/boltz-api-go/compare/v0.25.0...v0.25.1)

### Bug Fixes

* **compute-api:** remove public model codenames ([5d692a3](https://github.com/boltz-bio/boltz-api-go/commit/5d692a3270e10fc220abeacf58eda8824262f4db))
* **sm-screen:** batch molecule filtering to avoid 60s Python timeout ([3345f47](https://github.com/boltz-bio/boltz-api-go/commit/3345f4777b60bdb61f18a4e9f17f64f8d1195cae))

## 0.25.0 (2026-06-15)

Full Changelog: [v0.24.1...v0.25.0](https://github.com/boltz-bio/boltz-api-go/compare/v0.24.1...v0.25.0)

### Features

* **compute-api:** rename public pipeline identifiers ([9629dfe](https://github.com/boltz-bio/boltz-api-go/commit/9629dfe72d1037deb8373d81cebd8848c18956b2))

## 0.24.1 (2026-06-15)

Full Changelog: [v0.24.0...v0.24.1](https://github.com/boltz-bio/boltz-api-go/compare/v0.24.0...v0.24.1)

### Bug Fixes

* **boltz2:** route to H100 based on token count, not polymer residues ([0937b34](https://github.com/boltz-bio/boltz-api-go/commit/0937b345bf688bb2064dc5b9c2359e4d917fea3f))
* **compute-api:** raise Boltz2 sampling steps minimum ([cc2a83a](https://github.com/boltz-bio/boltz-api-go/commit/cc2a83a9f4e15a67d00166da509d504f19d0c984))

## 0.24.0 (2026-06-12)

Full Changelog: [v0.23.0...v0.24.0](https://github.com/boltz-bio/boltz-api-go/compare/v0.23.0...v0.24.0)

### Features

* **predictions:** expose predicted ligand SDF URL in compute_api + platform ([66d98d7](https://github.com/boltz-bio/boltz-api-go/commit/66d98d7e2cefeeaf128e692e8ac51da07feaab16))

## 0.23.0 (2026-06-09)

Full Changelog: [v0.22.0...v0.23.0](https://github.com/boltz-bio/boltz-api-go/compare/v0.22.0...v0.23.0)

### Features

* **compute:** accept legacy Boltz template inputs ([b4be1d6](https://github.com/boltz-bio/boltz-api-go/commit/b4be1d6caa2c7f37976aea72ab98d288f0c19508))

## 0.22.0 (2026-06-09)

Full Changelog: [v0.21.0...v0.22.0](https://github.com/boltz-bio/boltz-api-go/compare/v0.21.0...v0.22.0)

### Features

* **compute:** canonicalize Boltz template inputs ([30d59ea](https://github.com/boltz-bio/boltz-api-go/commit/30d59eac760d7f74dd762589f7f8e7f596e435aa))

## 0.21.0 (2026-06-05)

Full Changelog: [v0.20.0...v0.21.0](https://github.com/boltz-bio/boltz-api-go/compare/v0.20.0...v0.21.0)

### Features

* **lab:** reuse precomputed small molecule pockets ([251e950](https://github.com/boltz-bio/boltz-api-go/commit/251e950d005dd7b8de4482a92e8241b504a85612))

## 0.20.0 (2026-06-04)

Full Changelog: [v0.19.3...v0.20.0](https://github.com/boltz-bio/boltz-api-go/compare/v0.19.3...v0.20.0)

### Features

* **compute-api:** abort generative pipelines when filter+dedup rate exceeds 95% ([5fecf72](https://github.com/boltz-bio/boltz-api-go/commit/5fecf72b69dc68ce78e05bd57b0229d9b21f5dd5))

## 0.19.3 (2026-06-03)

Full Changelog: [v0.19.2...v0.19.3](https://github.com/boltz-bio/boltz-api-go/compare/v0.19.2...v0.19.3)

### Documentation

* **api:** document CCD-only polymer modifications ([2190dea](https://github.com/boltz-bio/boltz-api-go/commit/2190deabf4d0992906979c38ffe8ceeca06d48d1))

## 0.19.2 (2026-05-28)

Full Changelog: [v0.19.1...v0.19.2](https://github.com/boltz-bio/boltz-api-go/compare/v0.19.1...v0.19.2)

### Bug Fixes

* **protein-design:** count tyrosine in hydrophobic filter ([3b9594e](https://github.com/boltz-bio/boltz-api-go/commit/3b9594e6e2380028846abcc38cf4542b9318e644))

## 0.19.1 (2026-05-27)

Full Changelog: [v0.19.0...v0.19.1](https://github.com/boltz-bio/boltz-api-go/compare/v0.19.0...v0.19.1)

### Documentation

* document structure templates and custom MSA ([673a9ab](https://github.com/boltz-bio/boltz-api-go/commit/673a9ab7ea2db2c50b7bcc6d8a33a96010042361))

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
