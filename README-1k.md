# Awesome Rust [![lint badge](https://github.com/rust-unofficial/awesome-rust/actions/workflows/lint.yml/badge.svg)](https://github.com/rust-unofficial/awesome-rust/actions/workflows/lint.yml) [![build badge](https://github.com/rust-unofficial/awesome-rust/actions/workflows/rust.yml/badge.svg?branch=main)](https://github.com/rust-unofficial/awesome-rust/actions/workflows/rust.yml) [![Track Awesome List](https://www.trackawesomelist.com/badge.svg)](https://www.trackawesomelist.com/rust-unofficial/awesome-rust/)

A curated list of Rust code and resources.

If you want to contribute, please read [this](CONTRIBUTING.md).

## Table of contents

<!-- toc -->

- [Applications](#applications)
- [Development tools](#development-tools)
- [Libraries](#libraries)
    + [Genetic algorithms](#genetic-algorithms)
    + [Machine learning](#machine-learning)
    + [OpenAI](#openai)
- [Registries](#registries)
- [Resources](#resources)
- [License](#license)

<!-- tocstop -->

## Applications

See also [Rust - Production](https://www.rust-lang.org/production) organizations running Rust in production.

* [denoland/deno](https://github.com/denoland/deno) - A secure JavaScript/TypeScript runtime built with V8 and Tokio [![Build Status](https://github.com/denoland/deno/actions/workflows/ci.yml/badge.svg)](https://github.com/denoland/deno/actions) Stars:`102.5K`.
* [alacritty](https://github.com/alacritty/alacritty) - A cross-platform, GPU enhanced terminal emulator Stars:`58.2K`.
* [SWC](https://github.com/swc-project/swc) - super-fast TypeScript / JavaScript compiler Stars:`32.0K`.
* [Servo](https://github.com/servo/servo) - A prototype web browser engine Stars:`30.1K`.
* [zellij](https://github.com/zellij-org/zellij) - A terminal multiplexer (workspace) with batteries included Stars:`23.8K`.
* [Sniffnet](https://github.com/GyulyVGC/sniffnet) - Cross-platform application to monitor your network traffic with ease [![build badge](https://img.shields.io/github/actions/workflow/status/gyulyvgc/sniffnet/rust.yml?logo=github)](https://github.com/GyulyVGC/sniffnet/blob/main/.github/workflows/rust.yml) [![crate](https://img.shields.io/crates/v/sniffnet?logo=rust)](https://crates.io/crates/sniffnet) Stars:`23.2K`.
* [wezterm](https://github.com/wezterm/wezterm) - A GPU-accelerated cross-platform terminal emulator and multiplexer Stars:`19.8K`.
* [wasmer](https://github.com/wasmerio/wasmer) - A safe and fast WebAssembly runtime supporting WASI and Emscripten [![Build Status](https://github.com/wasmerio/wasmer/actions/workflows/build.yml/badge.svg)](https://github.com/wasmerio/wasmer/actions) Stars:`19.6K`.
* [mdBook](https://github.com/rust-lang/mdBook) - A command line utility to create books from markdown files [![Build Status](https://github.com/rust-lang/mdBook/actions/workflows/main.yml/badge.svg)](https://github.com/rust-lang/mdBook/actions) Stars:`19.4K`.
* [shuttle](https://github.com/shuttle-hq/shuttle) - A serverless platform. Stars:`6.4K`.
* [cloudflare/boringtun](https://github.com/cloudflare/boringtun) - A Userspace WireGuard VPN Implementation [![build badge](https://img.shields.io/crates/v/boringtun.svg)](https://crates.io/crates/boringtun) Stars:`6.3K`.
* [innernet](https://github.com/tonarino/innernet) - An overlay or private mesh network that uses Wireguard under the hood Stars:`5.1K`.
* [Rio](https://github.com/raphamorim/rio) - A hardware-accelerated GPU terminal emulator powered by WebGPU, focusing to run in desktops and browsers. Stars:`4.9K`.
* [hickory-dns](https://crates.io/crates/hickory-dns) - A DNS-server [![Build Status](https://github.com/hickory-dns/hickory-dns/actions/workflows/test.yml/badge.svg)](https://github.com/hickory-dns/hickory-dns/actions?query=workflow%3Atest) Stars:`4.4K`.
* [mirrord](https://github.com/metalbear-co/mirrord) - Connect your local process and your cloud environment, and run local code in cloud conditions Stars:`4.0K`.
* [EasyTier](https://github.com/EasyTier/EasyTier) - A simple, full-featured and decentralized mesh VPN with WireGuard support. [![crates.io](https://img.shields.io/crates/v/easytier)](https://crates.io/crates/easytier) [![GitHub actions](https://github.com/EasyTier/EasyTier/actions/workflows/core.yml/badge.svg)](https://github.com/EasyTier/EasyTier/actions/)[![GitHub actions](https://github.com/EasyTier/EasyTier/actions/workflows/gui.yml/badge.svg)](https://github.com/EasyTier/EasyTier/actions/) Stars:`3.6K`.
* [rx](https://github.com/cloudhead/rx) - Vi inspired Modern Pixel Art Editor Stars:`3.2K`.
* [WinterJS](https://github.com/wasmerio/winterjs) - A secure JavaScript runtime built with SpiderMonkey and Axum Stars:`3.1K`.
* [habitat](https://github.com/habitat-sh/habitat) - A tool created by Chef to build, deploy, and manage applications. Stars:`2.6K`.
* [Ryot](https://github.com/ignisda/ryot) - A self hosted application to track media consumption, fitness, etc. Stars:`2.4K`.
* [fcsonline/drill](https://github.com/fcsonline/drill) - A HTTP load testing application inspired by Ansible syntax Stars:`2.2K`.
* [linkerd/linkerd2-proxy](https://github.com/linkerd/linkerd2-proxy) - Ultralight service mesh for Kubernetes. Stars:`2.0K`.
* [defguard](https://github.com/defguard/defguard) - Enterprise Open Source SSO & WireGuard VPN with real 2FA/MFA Stars:`1.9K`.
* [kalker](https://github.com/PaddiM8/kalker) - A scientific calculator that supports math-like syntax with user-defined variables, functions, derivation, integration, and complex numbers. Cross-platform + WASM support [![Build Status](https://github.com/PaddiM8/kalker/workflows/Release/badge.svg)](https://github.com/PaddiM8/kalker/actions) Stars:`1.7K`.
* [tiny](https://github.com/osa1/tiny) - A terminal IRC client Stars:`1.1K`.

### Audio and Music

* [Spotifyd](https://github.com/Spotifyd/spotifyd) - An open source Spotify client running as a UNIX daemon. [![Continuous Integration](https://github.com/Spotifyd/spotifyd/actions/workflows/ci.yml/badge.svg)](https://github.com/Spotifyd/spotifyd/actions/workflows/ci.yml) Stars:`10.1K`.
* [ncspot](https://github.com/hrkfdn/ncspot) - Cross-platform ncurses Spotify client, inspired by ncmpc and the likes. [![build badge](https://github.com/hrkfdn/ncspot/actions/workflows/ci.yml/badge.svg)](https://github.com/hrkfdn/ncspot/actions?query=workflow%3ABuild) Stars:`5.4K`.
* [Spotify Player](https://github.com/aome510/spotify-player) - A Spotify player in the terminal with full feature parity. Stars:`4.2K`.
* [Glicol](https://github.com/chaosprint/glicol) - Graph-oriented live coding language, for collaborative musicking in browsers. Stars:`2.6K`.
* [Polaris](https://github.com/agersant/polaris) - A music streaming application. Stars:`2.2K`.
* [termusic](https://github.com/tramhao/termusic) - Music Player TUI written Stars:`1.3K`.

### Blockchain

* [Diem](https://github.com/diem/diem) - Diem’s mission is to enable a simple global currency and financial infrastructure that empowers billions of people. Stars:`16.7K`.
* [Foundry](https://github.com/foundry-rs/foundry) - Foundry is a blazing fast, portable and modular toolkit for Ethereum application development. ![Build Status](https://img.shields.io/github/workflow/status/foundry-rs/foundry/test?style=flat-square) Stars:`8.8K`.
* [Sui](https://github.com/MystenLabs/sui) - A next-generation smart contract platform with high throughput, low latency, and an asset-oriented programming model powered by the Move programming language. Stars:`6.8K`.
* [Grin](https://github.com/mimblewimble/grin/) - Evolution of the MimbleWimble protocol Stars:`5.1K`.
* [zcash](https://github.com/zcash/zcash) - Zcash is an implementation of the "Zerocash" protocol. Stars:`5.0K`.
* [reth](https://github.com/paradigmxyz/reth) - Modular, contributor-friendly and blazing-fast implementation of the Ethereum protocol. Stars:`4.5K`.
* [Lighthouse](https://github.com/sigp/lighthouse) - Ethereum Consensus Layer (CL) Client [![Build Status](https://github.com/sigp/lighthouse/actions/workflows/test-suite.yml/badge.svg)](https://github.com/sigp/lighthouse/actions) Stars:`3.1K`.
* [artemis](https://github.com/paradigmxyz/artemis) - A simple, modular, and fast framework for writing MEV bots. Stars:`2.5K`.
* [near/nearcore](https://github.com/near/nearcore) - decentralized smart-contract platform for low-end mobile devices. Stars:`2.4K`.
* [rust-bitcoin](https://github.com/rust-bitcoin/rust-bitcoin) - Library with support for de/serialization, parsing and executing on data structures and network messages related to Bitcoin. Stars:`2.3K`.
* [polkadot-sdk](https://github.com/paritytech/polkadot-sdk) - The Parity Polkadot Blockchain SDK Stars:`2.2K`.
* [revm](https://github.com/bluealloy/revm) - Revolutionary Machine (revm) is a fast Ethereum virtual machine. Stars:`1.8K`.
* [cairo](https://github.com/starkware-libs/cairo) - Cairo is the first Turing-complete language for creating provable programs for general computation. This is also the native language of [StarkNet](https://www.starknet.io), a ZK-Rollup using STARK proofs ![GitHub Workflow Status](https://img.shields.io/github/workflow/status/starkware-libs/cairo/CI?style=flat-square&logo=github) Stars:`1.7K`.
* [Joystream](https://github.com/Joystream/joystream) - A user governed video platform Stars:`1.4K`.
* [CITA](https://github.com/citahub/cita) - A high performance blockchain kernel for enterprise users. Stars:`1.3K`.
* [rust-lightning](https://github.com/lightningdevkit/rust-lightning) [![Crate](https://img.shields.io/crates/v/lightning.svg?logo=rust)](https://crates.io/crates/lightning) - Bitcoin Lightning library. The main crate,`lightning`, does not handle networking, persistence, or any other I/O. Thus,it is runtime-agnostic, but users must implement basic networking logic, chain interactions, and disk storage.po on linking crate. Stars:`1.3K`.
* [Holochain](https://github.com/holochain/holochain) - Scalable P2P alternative to blockchain for all those distributed apps you always wanted to build. [![detect critical check failures](https://github.com/holochain/holochain/actions/workflows/autorebase.yml/badge.svg)](https://github.com/holochain/holochain/actions/) Stars:`1.2K`.
* [Nervos CKB](https://github.com/nervosnetwork/ckb) - Nervos CKB is a public permissionless blockchain, the common knowledge layer of Nervos network. Stars:`1.2K`.
* [electrumrs](https://github.com/romanz/electrs) - An efficient re-implementation of Electrum Server. Stars:`1.2K`.

### Database

* [SurrealDB](https://github.com/surrealdb/surrealdb) - A scalable, distributed, document-graph database [![Build Status](https://img.shields.io/github/workflow/status/surrealdb/surrealdb/Continuous%20integration/main)](https://github.com/surrealdb/surrealdb/actions) Stars:`29.0K`.
* [Qdrant](https://github.com/qdrant/qdrant) - An open source vector similarity search engine with extended filtering support [![Tests](https://github.com/qdrant/qdrant/actions/workflows/rust.yml/badge.svg)](https://github.com/qdrant/qdrant/actions) Stars:`22.9K`.
* [Neon](https://github.com/neondatabase/neon) - Serverless Postgres. We separated storage and compute to offer autoscaling, branching, and bottomless storage. Stars:`16.7K`.
* [tikv](https://github.com/tikv/tikv) - A distributed KV database in Rust [![Build Status](https://ci.pingcap.net/job/tikv_ghpr_test/badge/icon)](https://ci.pingcap.net/job/tikv_ghpr_test/) Stars:`15.7K`.
* [Limbo](https://github.com/tursodatabase/limbo) - Limbo is a work-in-progress, in-process OLTP database management system, compatible with SQLite. Stars:`10.0K`.
* [sled](https://crates.io/crates/sled) - A (beta) modern embedded database [![Build Status](https://github.com/spacejam/sled/actions/workflows/test.yml/badge.svg)](https://github.com/spacejam/sled/actions?workflow=Rust) Stars:`8.4K`.
* [Databend](https://github.com/databendlabs/databend) - A Modern Real-Time Data Processing & Analytics DBMS with Cloud-Native Architecture [![Release](https://github.com/databendlabs/databend/actions/workflows/release.yml/badge.svg)](https://github.com/databendlabs/databend/actions) Stars:`8.3K`.
* [RisingWaveLabs/RisingWave](https://github.com/RisingWaveLabs/risingwave) - the next-generation streaming database in the cloud [![CI](https://github.com/RisingWaveLabs/risingwave/actions/workflows/main.yml/badge.svg)](https://github.com/RisingWaveLabs/risingwave/actions/workflows/main.yml/badge.svg?branch=main) Stars:`7.6K`.
* [ParadeDB](https://github.com/paradedb/paradedb/) - ParadeDB is an Elasticsearch alternative built on Postgres, designed for real-time search and analytics. Stars:`6.9K`.
* [erikgrinaker/toydb](https://github.com/erikgrinaker/toydb) - Distributed SQL database, written as a learning project. Stars:`6.4K`.
* [lancedb](https://github.com/lancedb/lancedb) [[vectordb](https://crates.io/crates/vectordb)] - A serverless, low-latency vector database for AI applications Stars:`6.0K`.
* [Materialize](https://github.com/MaterializeInc/materialize) - Streaming SQL database powered by Timely Dataflow :heavy_dollar_sign: [![Build status](https://badge.buildkite.com/97d6604e015bf633d1c2a12d166bb46f3b43a927d3952c999a.svg?branch=main)](https://buildkite.com/materialize/test) Stars:`5.9K`.
* [noria](https://github.com/mit-pdos/noria) [[noria](https://crates.io/crates/noria)] - Dynamically changing, partially-stateful data-flow for web application backends Stars:`5.1K`.
* [GreptimeDB](https://github.com/grepTimeTeam/greptimedb/) - An open-source, cloud-native, distributed time-series database with PromQL/SQL/Python supported.[![CI](https://github.com/greptimeTeam/greptimedb/actions/workflows/develop.yml/badge.svg)](https://github.com/greptimeTeam/greptimedb/actions/workflows/develop.yml) Stars:`4.8K`.
* [CozoDB](https://github.com/cozodb/cozo) - A transactional, relational database that uses Datalog and focuses on graph data and algorithms. Time-travel-capable, and fast! [![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/cozodb/cozo/build.yml?branch=main)](https://github.com/cozodb/cozo/actions/workflows/build.yml) Stars:`3.6K`.
* [USearch](https://github.com/unum-cloud/usearch) - Similarity Search Engine for Vectors and Strings [![crates.io](https://img.shields.io/crates/v/usearch.svg)](https://crates.io/crates/usearch) Stars:`2.6K`.
* [SQLSync](https://github.com/orbitinghail/sqlsync) - Multiplayer offline-first SQLite [![GitHub Workflow Status](https://github.com/orbitinghail/sqlsync/actions/workflows/actions.yaml/badge.svg?branch=main)](https://github.com/orbitinghail/sqlsync/actions?query=branch%3Amain) Stars:`2.6K`.
* [Skytable](https://github.com/skytable/skytable) - A multi-model NoSQL database ![GitHub Workflow Status](https://img.shields.io/github/workflow/status/skytable/skytable/Tests?style=flat-square) Stars:`2.5K`.
* [TrailBase](https://github.com/trailbaseio/trailbase) - A fast, lightweight, single-file FireBase alternative with type-safe APIs, built-in V8 JS/ES6/TS engine, auth and admin dashboard [![GitHub Workflow Status](https://github.com/trailbaseio/trailbase/workflows/test/badge.svg)](https://github.com/trailbaseio/trailbase/actions?workflow=test) Stars:`1.9K`.
* [seppo0010/rsedis](https://github.com/seppo0010/rsedis) - A Redis reimplementation. Stars:`1.8K`.
* [PumpkinDB](https://github.com/PumpkinDB/PumpkinDB) - an event sourcing database engine Stars:`1.4K`.
* [Atomic-Server](https://github.com/atomicdata-dev/atomic-server/) [[atomic-server](https://crates.io/crates/atomic_server)] - NoSQL graph database with realtime updates, dynamic indexing and easy-to-use GUI for CMS purposes. [![Release](https://github.com/atomicdata-dev/atomic-server/actions/workflows/release_please.yml/badge.svg)](https://github.com/atomicdata-dev/atomic-server/actions) Stars:`1.2K`.
* [oxigraph/oxigraph](https://github.com/oxigraph/oxigraph) [[oxigraph](https://crates.io/crates/oxigraph)] - graph database implementing the [SPARQL](https://www.w3.org/TR/sparql11-overview/) standard ![Crates.io Version](https://img.shields.io/crates/v/oxigraph?logo=Rust) Stars:`1.2K`.

### Embedded


### Emulators

See also [crates matching keyword 'emulator'](https://crates.io/keywords/emulator).

* CHIP-8
* Commodore 64
* Flash Player
  * [Ruffle](https://github.com/ruffle-rs/ruffle) - Ruffle is an Adobe Flash Player emulator. Ruffle targets both the desktop and the web using WebAssembly. [![CI](https://github.com/ruffle-rs/ruffle/actions/workflows/test_rust.yml/badge.svg)](https://github.com/ruffle-rs/ruffle/actions/workflows/test_rust.yml)[![CI](https://github.com/ruffle-rs/ruffle/actions/workflows/test_web.yml/badge.svg)](https://github.com/ruffle-rs/ruffle/actions/workflows/test_web.yml) Stars:`16.4K`.
* Gameboy
  * [mohanson/gameboy](https://github.com/mohanson/gameboy) - Full featured Cross-platform GameBoy emulator. Forever boys!. Stars:`1.4K`.
* Gameboy Advance
* GameMaker
* IBM PC
* Intel 8080 CPU
* iOS
  * [touchHLE](https://github.com/touchHLE/touchHLE) - High-level emulator for iPhone OS apps Stars:`3.0K`.
* iPod
* NES
* Nintendo 64
* Nintendo DS
* PlayStation 4
* Shockwave Player
* ZX Spectrum

### File manager

* [yazi](https://github.com/sxyazi/yazi) - Blazing fast terminal file manager, based on async I/O. Stars:`23.8K`.
* [broot](https://github.com/Canop/broot) - A new way to see and navigate directory trees (get an overview of a directory, even a big one; find a directory then `cd` to it; never lose track of file hierarchy while you search; manipulate your files, ...), further reading [dystroy.org/broot](https://dystroy.org/broot/) [![Latest Version](https://img.shields.io/crates/v/broot.svg)](https://crates.io/crates/broot) Stars:`11.2K`.
* [xplr](https://github.com/sayanarijit/xplr) - A hackable, minimal, fast TUI file explorer Stars:`4.3K`.
* [joshuto](https://github.com/kamiyaa/joshuto) - ranger-like terminal file manager Stars:`3.5K`.

### Finance

See also [Payments](#payments) applications.

* [tarkah/tickrs](https://github.com/tarkah/tickrs) - Realtime ticker data in your terminal Stars:`1.3K`.

### Games

See also [Games Made With Piston](https://github.com/PistonDevelopers/piston/wiki/Games-Made-With-Piston).

* [citybound](https://github.com/citybound/citybound) - The city sim you deserve Stars:`7.8K`.
* [mtkennerly/ludusavi](https://github.com/mtkennerly/ludusavi) - Backup tool for PC game saves [![build badge](https://img.shields.io/github/actions/workflow/status/mtkennerly/ludusavi/main.yaml?logo=github)](https://github.com/mtkennerly/ludusavi/actions/workflows/main.yaml) [![crate](https://img.shields.io/crates/v/ludusavi?logo=rust)](https://crates.io/crates/ludusavi) Stars:`3.5K`.
* [cristicbz/rust-doom](https://github.com/cristicbz/rust-doom) - A renderer for Doom, may progress to being a playable game Stars:`2.4K`.
* [ozkriff/zemeroth](https://github.com/ozkriff/zemeroth) - A small 2D turn-based hexagonal strategy game Stars:`1.4K`.
* [gorilla-devs/ferium](https://github.com/gorilla-devs/ferium) - Ferium is a fast and feature rich CLI program for downloading and updating Minecraft mods from Modrinth, CurseForge, and GitHub Releases, and modpacks from Modrinth and CurseForge ![ferium build](https://github.com/gorilla-devs/ferium/actions/workflows/build.yml/badge.svg?branch=main) Stars:`1.2K`.
* [doukutsu-rs](https://github.com/doukutsu-rs/doukutsu-rs) - Reimplementation of Cave Story engine with some enhancements. Stars:`1.0K`.

### Graphics

* [flxzt/rnote](https://github.com/flxzt/rnote) - Sketch and take handwritten notes. Stars:`9.2K`.
* [ivanceras/svgbob](https://github.com/ivanceras/svgbob) - converts ASCII diagrams into SVG graphics Stars:`4.0K`.
* [linebender/resvg](https://github.com/linebender/resvg) - An SVG rendering library. Stars:`3.1K`.

### Image processing

* [shssoichiro/oxipng](https://github.com/shssoichiro/oxipng) [[oxipng](https://crates.io/crates/oxipng)] - Multithreaded PNG optimizer written in Rust. [![Build Status](https://github.com/shssoichiro/oxipng/workflows/oxipng/badge.svg)](https://github.com/shssoichiro/oxipng/actions?query=branch%3Amaster) [![Version](https://img.shields.io/crates/v/oxipng.svg)](https://crates.io/crates/oxipng) Stars:`3.2K`.

### Industrial automation


### Message Queue

* [Rocketmq-Rust](https://github.com/mxsm/rocketmq-rust) - 🚀Apache RocketMQ build in Rust🦀. Faster, safer, and with lower memory usage. Stars:`1.0K`.

### MLOps

* [TensorZero](https://github.com/tensorzero/tensorzero) - data & learning flywheel for LLMs that unifies inference, observability, optimization, and experimentation ![TensorZero Build Status](https://img.shields.io/github/check-runs/tensorzero/tensorzero/main) Stars:`3.4K`.

### Observability

* [vectordotdev/vector](https://github.com/vectordotdev/vector) - A High-Performance, Logs, Metrics, & Events Router. Stars:`19.1K`.
* [openobserve](https://github.com/openobserve/openobserve) - 10x easier, 140x lower storage cost, high performance, petabyte scale - Elasticsearch/Splunk/Datadog alternative. Stars:`14.9K`.
* [Quickwit-oss/quickwit](https://github.com/quickwit-oss/quickwit) - Cloud-native and highly cost-efficient search engine for log management. [![CI](https://github.com/quickwit-oss/quickwit/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/quickwit-oss/quickwit/actions?query=workflow%3ACI) Stars:`9.9K`.
* [OpenTelemetry](https://crates.io/crates/opentelemetry) - OpenTelemetry provides a single set of APIs, libraries, agents, and collector services to capture distributed traces and metrics from your application. You can analyze them using Prometheus, Jaeger, and other observability tools. [![GitHub Actions CI](https://github.com/open-telemetry/opentelemetry-rust/actions/workflows/ci.yml/badge.svg)](https://github.com/open-telemetry/opentelemetry-rust/actions/workflows/ci.yml) Stars:`2.1K`.
* [Scaphandre](https://github.com/hubblo-org/scaphandre) - A power consumption monitoring agent, to track host and each service power consumption and enable designing systems and applications for more sustainability. Designed to fit any monitoring toolchain (already supports prometheus, warp10, riemann...). Stars:`1.7K`.

### Operating systems

See also [A comparison of operating systems written in Rust](https://github.com/flosse/rust-os-comparison).

* [tock/tock](https://github.com/tock/tock) - A secure embedded operating system for Cortex-M based microcontrollers Stars:`5.7K`.
* [theseus-os/Theseus](https://github.com/theseus-os/Theseus) - A safe-language, single address space and single privilege level OS written from scratch - [![build badge](https://img.shields.io/github/workflow/status/theseus-os/Theseus/Documentation?label=docs%20build)](https://www.theseus-os.com/Theseus/book/index.html) Stars:`3.0K`.
* [Andy-Python-Programmer/aero](https://github.com/Andy-Python-Programmer/aero) - A modern, unix-like operating system following the monolithic kernel design. Stars:`1.2K`.

### Package Managers


### Payments

* [hyperswitch](https://github.com/juspay/hyperswitch) - An open source payments orchestrator that lets you connect with multiple payment processors and route payment traffic effortlessly, all with a single API integration ![GitHub last commit](https://img.shields.io/github/last-commit/juspay/hyperswitch?style=flat-square) Stars:`15.6K`.

### Productivity

* [espanso](https://github.com/espanso/espanso) - A cross-platform Text Expander. [![CI](https://github.com/espanso/espanso/actions/workflows/ci.yml/badge.svg?branch=dev&event=push)](https://github.com/espanso/espanso/actions/workflows/ci.yml) Stars:`10.8K`.
* [ast-grep](https://github.com/ast-grep/ast-grep) - A CLI tool for code structural search, lint and rewriting. Stars:`8.4K`.
* [aichat](https://github.com/sigoden/aichat) - All-in-one LLM CLI tool featuring Shell Assistant, Chat-REPL, RAG, AI Tools & Agents, with access to OpenAI, Claude, Gemini, Ollama, Groq, and more. Stars:`6.3K`.
* [LLDAP](https://github.com/lldap/lldap) - Simplified LDAP interface for authentication. Stars:`4.9K`.

### Routing protocols


### Security tools

* [rustscan](https://github.com/bee-san/RustScan) - Make Nmap faster with this port scanning tool [![build badge](https://github.com/bee-san/RustScan/actions/workflows/test.yml/badge.svg)](https://github.com/bee-san/RustScan/actions) Stars:`16.1K`.
* [epi052/feroxbuster](https://github.com/epi052/feroxbuster) - A simple, fast, recursive content discovery tool. Stars:`6.4K`.
* [AFLplusplus/LibAFL](https://github.com/AFLplusplus/LibAFL) - Advanced Fuzzing Library - Slot your Fuzzer together in Rust! Scales across cores and machines. For Windows, Android, MacOS, Linux, no_std, etc. [![build and test](https://github.com/AFLplusplus/LibAFL/actions/workflows/build_and_test.yml/badge.svg)](https://github.com/AFLplusplus/LibAFL/actions/workflows/build_and_test.yml) Stars:`2.2K`.
* [kpcyrd/sn0int](https://github.com/kpcyrd/sn0int) - A semi-automatic OSINT framework and package manager Stars:`2.2K`.
* [observer_ward](https://github.com/emo-crab/observer_ward) - Web application and service fingerprint identification tool Stars:`1.4K`.
* [kpcyrd/sniffglue](https://github.com/kpcyrd/sniffglue) - A secure multithreaded packet sniffer Stars:`1.2K`.
* [Cherrybomb](https://github.com/blst-security/cherrybomb) - Stop half-done API specifications with a CLI tool that helps you avoid undefined user behaviour by validating your API specifications. Stars:`1.2K`.
* [AdGuardian-Term](https://github.com/Lissy93/AdGuardian-Term) [[adguardian](https://crates.io/crates/adguardian)] - Terminal-based, real-time traffic monitoring and statistics for your AdGuard Home instance Stars:`1.1K`.

### Social networks

* Mastodon
* Telegram

### System tools

* [sharkdp/bat](https://github.com/sharkdp/bat) - A cat(1) clone with wings. [![CICD](https://github.com/sharkdp/bat/actions/workflows/CICD.yml/badge.svg?branch=master)](https://github.com/sharkdp/bat/actions/workflows/CICD.yml) Stars:`52.1K`.
* [sharkdp/fd](https://github.com/sharkdp/fd) - A simple, fast and user-friendly alternative to find. [![CICD](https://github.com/sharkdp/fd/actions/workflows/CICD.yml/badge.svg)](https://github.com/sharkdp/fd/actions/workflows/CICD.yml) Stars:`37.4K`.
* [nushell/nushell](https://github.com/nushell/nushell) - A new type of shell Stars:`34.6K`.
* [fish-shell/fish-shell](https://github.com/fish-shell/fish-shell) - The user-friendly command line shell Stars:`29.4K`.
* [ajeetdsouza/zoxide](https://github.com/ajeetdsouza/zoxide/) - A fast alternative to `cd` that learns your habits [![release](https://github.com/ajeetdsouza/zoxide/actions/workflows/release.yml/badge.svg)](https://github.com/ajeetdsouza/zoxide/actions) Stars:`25.8K`.
* [atuin](https://github.com/atuinsh/atuin) [[atuin](https://crates.io/crates/atuin)] - Atuin replaces your existing shell history with a SQLite database, and records additional context for your commands. Additionally, it provides optional and fully encrypted synchronisation of your history between machines, via an Atuin server. Stars:`23.2K`.
* [qarmin/czkawka](https://github.com/qarmin/czkawka) - Multi-functional app to find duplicates, empty folders, similar images, etc. [![GitHub Actions Workflow](https://github.com/qarmin/czkawka/actions/workflows/pages/pages-build-deployment/badge.svg?branch=master)](https://github.com/qarmin/czkawka/actions) Stars:`22.9K`.
* [uutils/coreutils](https://github.com/uutils/coreutils) - A cross-platform rewrite of the GNU coreutils [![CICD](https://github.com/uutils/coreutils/actions/workflows/CICD.yml/badge.svg)](https://github.com/uutils/coreutils/actions/workflows/CICD.yml) Stars:`19.8K`.
* [gitui](https://github.com/gitui-org/gitui) - Blazing fast terminal client for git. [![build](https://github.com/gitui-org/gitui/actions/workflows/ci.yml/badge.svg)](https://github.com/gitui-org/gitui/actions) Stars:`19.4K`.
* [eza-community/eza](https://github.com/eza-community/eza) - A replacement for 'ls' Stars:`14.8K`.
* [lsd](https://github.com/lsd-rs/lsd) - An ls with a lot of pretty colors and awesome icons [![build](https://github.com/lsd-rs/lsd/actions/workflows/CICD.yml/badge.svg)](https://github.com/lsd-rs/lsd/actions) Stars:`14.2K`.
* [XAMPPRocky/tokei](https://github.com/XAMPPRocky/tokei) - counts the lines of code Stars:`12.2K`.
* [bottom](https://github.com/ClementTsang/bottom) - Yet another cross-platform graphical process/system monitor. [![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/ClementTsang/bottom/ci/master)](https://github.com/ClementTsang/bottom/actions?query=branch%3Amaster) Stars:`11.1K`.
* [bandwhich](https://github.com/imsnif/bandwhich) - Terminal bandwidth utilization tool Stars:`10.4K`.
* [dust](https://github.com/bootandy/dust) - A more intuitive version of du Stars:`9.6K`.
* [cantino/mcfly](https://github.com/cantino/mcfly) - Fly through your shell history. Great Scott! Stars:`7.2K`.
* [watchexec](https://github.com/watchexec/watchexec) - Executes commands in response to file modifications Stars:`5.9K`.
* [skim](https://github.com/skim-rs/skim) - A fuzzy finder Stars:`5.6K`.
* [pueue](https://github.com/nukesor/pueue) - Manage your long running shell commands. [![GitHub Actions Workflow](https://github.com/Nukesor/pueue/actions/workflows/test.yml/badge.svg)](https://github.com/nukesor/pueue/actions) Stars:`5.5K`.
* [ynqa/jnv](https://github.com/ynqa/jnv) - Interactive JSON filter using jq [![ci](https://github.com/ynqa/jnv/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/ynqa/jnv/actions/workflows/ci.yml) Stars:`5.5K`.
* [dalance/procs](https://github.com/dalance/procs) - A modern replacement for 'ps' [![Regression](https://github.com/dalance/procs/actions/workflows/regression.yml/badge.svg)](https://github.com/dalance/procs/actions/workflows/regression.yml) Stars:`5.4K`.
* [trippy](https://github.com/fujiapple852/trippy) - A network diagnostic tool [![build badge](https://github.com/fujiapple852/trippy/workflows/CI/badge.svg)](https://github.com/fujiapple852/trippy/actions/workflows/ci.yml) Stars:`4.4K`.
* [GQL](https://github.com/amrdeveloper/gql) - A SQL like query language to run on .git files. Stars:`3.4K`.
* [ouch](https://github.com/ouch-org/ouch) - Painless compression and decompression on the command-line [![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/ouch-org/ouch/build-and-test)](https://github.com/ouch-org/ouch/actions?query=branch%3Amaster) Stars:`2.8K`.
* [orhun/kmon](https://github.com/orhun/kmon) - Linux Kernel Manager and Activity Monitor ![https://github.com/orhun/kmon/actions](https://img.shields.io/github/actions/workflow/status/orhun/kmon/ci.yml?branch=master&label=build) Stars:`2.7K`.
* [diskonaut](https://github.com/imsnif/diskonaut) - Terminal visual disk space navigator Stars:`2.6K`.
* [pkolaczk/fclones](https://github.com/pkolaczk/fclones) - Efficient duplicate file finder and remover Stars:`2.2K`.
* [LACT](https://github.com/ilya-zlobintsev/LACT) - Linux AMDGPU Controller Stars:`2.1K`.
* [Kondo](https://github.com/tbillington/kondo) - CLI & GUI tool for deleting software project artifacts and reclaiming disk space Stars:`1.9K`.
* [m4b/bingrep](https://github.com/m4b/bingrep) - Greps through binaries from various OSs and architectures, and colors them. Stars:`1.7K`.
* [nivekuil/rip](https://github.com/nivekuil/rip) - A safe and ergonomic alternative to `rm` Stars:`1.5K`.
* [redox-os/ion](https://github.com/redox-os/ion) - Next-generation system shell Stars:`1.5K`.
* [httm](https://github.com/kimono-koans/httm) - Interactive, file-level Time Machine-like tool for ZFS/btrfs/nilfs2 (and even actual Time Machine backups!) Stars:`1.5K`.
* [orhun/systeroid](https://github.com/orhun/systeroid) - A more powerful alternative to sysctl(8) with a terminal user interface ![https://github.com/orhun/systeroid/actions](https://img.shields.io/github/actions/workflow/status/orhun/systeroid/ci.yml?branch=main&label=build) Stars:`1.3K`.
* [netscanner](https://github.com/Chleba/netscanner) - TUI Network Scanner Stars:`1.1K`.

### Task scheduling


### Text editors

* [zed](https://github.com/zed-industries/zed) - A high-performance, multiplayer code editor from the creators of Atom and Tree-sitter. Stars:`56.6K`.
* [helix](https://github.com/helix-editor/helix) - A post-modern modal text editor inspired by Neovim/Kakoune. [![build badge](https://github.com/helix-editor/helix/actions/workflows/build.yml/badge.svg)](https://github.com/helix-editor/helix/actions) Stars:`36.8K`.
* [Lapce](https://github.com/lapce/lapce) - A modern editor with a backend. Taking inspiration from the discontinued [xi-editor](https://github.com/xi-editor/xi-editor). Stars:`35.1K`.
* [ox](https://github.com/curlpipe/ox) - An independent Rust text editor that runs in your terminal! Stars:`3.5K`.
* [emacs-ng](https://github.com/emacs-ng/emacs-ng) - Complementing the C codebase with rust code to introduce new features. Stars:`1.7K`.
* [ilai-deutel/kibi](https://github.com/ilai-deutel/kibi) - A tiny (≤1024 LOC) text editor with syntax highlighting, incremental search and more. [![build badge](https://github.com/ilai-deutel/kibi/workflows/CI/badge.svg?branch=master)](https://github.com/ilai-deutel/kibi/actions?query=branch%3Amaster) Stars:`1.7K`.
* [gchp/iota](https://github.com/gchp/iota) - A simple text editor Stars:`1.6K`.

### Text processing

* [phiresky/ripgrep-all](https://github.com/phiresky/ripgrep-all) - ripgrep, but also search in PDFs, E-Books, Office documents, zip, tar.gz, etc. Stars:`8.7K`.
* [grex](https://github.com/pemistahl/grex) - A command-line tool and library for generating regular expressions from user-provided test cases Stars:`7.4K`.
* [Melody](https://github.com/yoav-lavi/melody) - A language that compiles to regular expressions and aims to be more easily readable and maintainable [![build badge](https://github.com/yoav-lavi/melody/actions/workflows/rust.yml/badge.svg)](https://github.com/yoav-lavi/melody/actions/workflows/rust.yml) [![crates.io](https://img.shields.io/crates/v/melody_compiler?label=compiler)](https://crates.io/crates/melody_compiler) Stars:`4.7K`.
* [dathere/qsv](https://github.com/dathere/qsv) [[qsv](https://crates.io/crates/qsv)] - A high performance CSV data-wrangling toolkit. Forked from xsv, with 34+ additional commands & more. [![Linux build status](https://github.com/dathere/qsv/actions/workflows/rust.yml/badge.svg)](https://github.com/dathere/qsv/actions/workflows/rust.yml) [![Windows build status](https://github.com/dathere/qsv/actions/workflows/rust-windows.yml/badge.svg)](https://github.com/dathere/qsv/actions/workflows/rust-windows.yml) [![macOS build status](https://github.com/dathere/qsv/actions/workflows/rust-macos.yml/badge.svg)](https://github.com/dathere/qsv/actions/workflows/rust-macos.yml) Stars:`2.7K`.
* [ashvardanian/stringzilla](https://github.com/ashvardanian/StringZilla) - SIMD-accelerated string search, sort, edit distances, alignments, and generators for x86 AVX2 & AVX-512, and Arm NEON [![crates.io](https://img.shields.io/crates/v/stringzilla.svg)](https://crates.io/crates/stringzilla) Stars:`2.5K`.
* [dominikwilkowski/cfonts](https://github.com/dominikwilkowski/cfonts) [[cfonts](https://crates.io/crates/cfonts)] - Sexy ANSI fonts for the console ![build badge](https://github.com/dominikwilkowski/cfonts/actions/workflows/testing.yml/badge.svg) Stars:`1.6K`.

### Utilities

* [rustdesk/rustdesk](https://github.com/rustdesk/rustdesk) - A remote desktop software, great alternative to TeamViewer and AnyDesk. Stars:`86.2K`.
* [vaultwarden](https://github.com/dani-garcia/vaultwarden#readme) [![Build](https://github.com/dani-garcia/vaultwarden/actions/workflows/build.yml/badge.svg)](https://github.com/dani-garcia/vaultwarden/actions/workflows/build.yml) - Alternative implementation of the Bitwarden server API written in Rust Stars:`43.2K`.
* [warpdotdev/Warp](https://github.com/warpdotdev/Warp) - :heavy_dollar_sign: Warp is a blazingly-fast modern GPU-accelerated terminal built to make you and your team more productive. Stars:`23.0K`.
* [str4d/rage](https://github.com/str4d/rage) [[rage](https://crates.io/crates/rage)] - Rust implementation of [age](https://github.com/FiloSottile/age). Stars:`2.9K`.
* [television](https://github.com/alexpasmantier/television) - A blazing fast general purpose fuzzy finder TUI ![GitHub branch check runs](https://img.shields.io/github/check-runs/alexpasmantier/television/main) Stars:`2.6K`.
* [rustic-rs/rustic](https://github.com/rustic-rs/rustic) [[rustic-rs](https://crates.io/crates/rustic-rs)] - Fast, encrypted, deduplicated backups powered by Rust. [![Version](https://img.shields.io/crates/v/rustic-rs.svg)](https://crates.io/crates/rustic-rs) Stars:`2.3K`.
* [Vibe](https://github.com/thewh1teagle/vibe) - Transcribe audio or video in every language on every platform. Stars:`2.2K`.
* [mprocs](https://github.com/pvolok/mprocs) - TUI for running multiple processes Stars:`1.7K`.
* [nix-community/nix-init](https://github.com/nix-community/nix-init) - Generate Nix packages from URLs with hash prefetching, dependency inference, license detection, and more [![build-badge](https://github.com/nix-community/nix-init/actions/workflows/ci.yml/badge.svg)](https://github.com/nix-community/nix-init/actions/workflows/ci.yml) Stars:`1.0K`.

### Video

* [gyroflow/gyroflow](https://github.com/gyroflow/gyroflow) - Video stabilization application using gyroscope data Stars:`7.2K`.
* [xiph/rav1e](https://github.com/xiph/rav1e) - The fastest and safest AV1 encoder. Stars:`3.8K`.
* [harlanc/xiu](https://github.com/harlanc/xiu) - A powerful and secure live server (rtmp/httpflv/hls/relay). [![crates.io](https://img.shields.io/crates/v/xiu.svg)](https://crates.io/crates/xiu) Stars:`2.0K`.

### Virtualization

* [firecracker-microvm/firecracker](https://github.com/firecracker-microvm/firecracker) - A lightweight virtual machine for container workload [Firecracker Microvm](https://firecracker-microvm.github.io/) Stars:`27.5K`.
* [youki-dev/youki](https://github.com/youki-dev/youki) - A container runtime [![build badge](https://github.com/youki-dev/youki/actions/workflows/basic.yml/badge.svg)](https://github.com/youki-dev/youki/actions) Stars:`6.6K`.
* [kata-containers/kata-containers](https://github.com/kata-containers/kata-containers) - A implementation of lightweight Virtual Machines (VMs) that feel and perform like containers, but provide the workload isolation and security advantages of VMs. Stars:`6.0K`.
* [tailhook/vagga](https://github.com/tailhook/vagga) - A containerization tool without daemons Stars:`1.9K`.

### Web

* [LemmyNet/lemmy](https://github.com/LemmyNet/lemmy) - A link aggregator / reddit clone for the fediverse [![Build Status](https://cloud.drone.io/api/badges/LemmyNet/lemmy/status.svg)](https://cloud.drone.io/LemmyNet/lemmy) Stars:`13.7K`.
* [Plume-org/Plume](https://github.com/Plume-org/Plume) - ActivityPub federating blogging application Stars:`2.2K`.
* [Redlib](https://github.com/redlib-org/redlib) - An alternative private front-end to Reddit, with its origins in [Libreddit](https://github.com/libreddit/libreddit) Stars:`1.9K`.
* [Revolt/backend](https://github.com/revoltchat/backend) - User-first chat platform built with modern web technologies. Stars:`1.5K`.

### Web Servers

* [cloudflare/pingora](https://github.com/cloudflare/pingora) - A library for building fast, reliable and evolvable network services. Stars:`23.7K`.
* [svenstaro/miniserve](https://github.com/svenstaro/miniserve) - A small, self-contained cross-platform CLI tool that allows you to just grab the binary and serve some file(s) via HTTP [![build badge](https://github.com/svenstaro/miniserve/workflows/CI/badge.svg?branch=master)](https://github.com/svenstaro/miniserve/actions) Stars:`6.5K`.
* [TheWaWaR/simple-http-server](https://github.com/TheWaWaR/simple-http-server) - simple static http server Stars:`3.1K`.
* [static-web-server](https://github.com/static-web-server/static-web-server) - A blazing fast and asynchronous web server for static files-serving. ⚡ [![CI](https://github.com/static-web-server/static-web-server/actions/workflows/devel.yml/badge.svg)](https://github.com/static-web-server/static-web-server/actions/workflows/devel.yml?query=branch%3Amaster) Stars:`1.7K`.
* [mufeedvh/binserve](https://github.com/mufeedvh/binserve) - A blazingly fast static web server with routing, templating, and security in a single binary you can set up with zero code [![build badge](https://github.com/mufeedvh/binserve/actions/workflows/build.yml/badge.svg)](https://github.com/mufeedvh/binserve/actions) Stars:`1.1K`.

## Development tools

* [delta](https://crates.io/crates/git-delta) - A syntax-highlighter for git and diff output[![build badge](https://github.com/dandavison/delta/actions/workflows/ci.yml/badge.svg)](https://github.com/dandavison/delta//actions) Stars:`25.8K`.
* [just](https://github.com/casey/just) - A handy command runner for project-specific tasks Stars:`24.7K`.
* [git-cliff](https://github.com/orhun/git-cliff) - A highly customizable Changelog Generator that follows Conventional Commit specifications ![https://github.com/orhun/git-cliff/actions](https://img.shields.io/github/actions/workflow/status/orhun/git-cliff/ci.yml?branch=main&label=build) Stars:`9.9K`.
* [Rustup](https://github.com/rust-lang/rustup) - the Rust toolchain installer [![build badge](https://github.com/rust-lang/rustup/actions/workflows/ci.yaml/badge.svg)](https://github.com/rust-lang/rustup/actions) Stars:`6.4K`.
* [Racer](https://github.com/racer-rust/racer) - code completion for Rust Stars:`3.4K`.
* [Flox](https://github.com/flox/flox) - Flox is a virtual environment and package manager all in one. Stars:`3.1K`.
* [typos](https://github.com/crate-ci/typos) [[typos-cli](https://crates.io/crates/typos-cli)] - Source code spell checker Stars:`3.0K`.
* [ATAC](https://github.com/Julien-cpsn/ATAC) - A feature-full TUI API client made in Rust. ATAC is free, open-source, offline and account-less. Stars:`2.6K`.
* [bacon](https://github.com/Canop/bacon) - background rust code checker, similar to cargo-watch Stars:`2.4K`.
* [dotenv-linter](https://github.com/dotenv-linter/dotenv-linter) - Linter for `.env` files [![build badge](https://github.com/dotenv-linter/dotenv-linter/actions/workflows/ci.yml/badge.svg)](https://github.com/dotenv-linter/dotenv-linter/actions?query=workflow%3ACI+branch%3Amaster) Stars:`1.9K`.
* [create-rust-app](https://github.com/Wulf/create-rust-app) - Set up a modern rust+react web app by running one command. [![crate](https://img.shields.io/crates/v/create-rust-app.svg)](https://crates.io/crates/create-rust-app) Stars:`1.6K`.
* [geiger](https://github.com/geiger-rs/cargo-geiger) - A program that list statistics related to usage of unsafe code in a crate and all its dependencies [![Build Status](https://dev.azure.com/cargo-geiger/cargo-geiger/_apis/build/status/geiger-rs.cargo-geiger?branchName=master)](https://dev.azure.com/cargo-geiger/cargo-geiger/_build/latest?definitionId=1&branchName=master) Stars:`1.4K`.
* [cloudflare/foundations](https://github.com/cloudflare/foundations) - Foundations is a modular Rust library, designed to help scale programs for distributed, production-grade systems. Stars:`1.4K`.
* [Rust Search Extension](https://github.com/huhu/rust-search-extension) - A handy browser extension to search crates and docs in address bar (omnibox). [![Build Status](https://github.com/huhu/rust-search-extension/workflows/build/badge.svg?branch=master)](https://github.com/huhu/rust-search-extension/actions) Stars:`1.2K`.
* [mask](https://github.com/jacobdeichert/mask) - A CLI task runner defined by a simple markdown file [![build badge](https://github.com/jacobdeichert/mask/workflows/CI/badge.svg?branch=master)](https://github.com/jacobdeichert/mask/actions?query=workflow%3ACI) Stars:`1.1K`.
* [scriptisto](https://github.com/igor-petruk/scriptisto) - A language-agnostic "shebang interpreter" that enables you to write one file scripts in compiled languages. [![Build Status](https://cloud.drone.io/api/badges/igor-petruk/scriptisto/status.svg)](https://cloud.drone.io/igor-petruk/scriptisto) Stars:`1.0K`.

### Build system

  * [dtolnay/cargo-expand](https://github.com/dtolnay/cargo-expand) - Expand macros in your source code Stars:`2.8K`.
  * [cargo-make](https://crates.io/crates/cargo-make) - Task runner and build tool. [![build badge](https://github.com/sagiegurari/cargo-make/workflows/CI/badge.svg?branch=master)](https://github.com/sagiegurari/cargo-make/actions) Stars:`2.7K`.
  * [cargo-generate](https://github.com/cargo-generate/cargo-generate) - A generator of a rust project by leveraging a pre-existing git repository as a template. Stars:`2.1K`.
  * [cargo-udeps](https://github.com/est31/cargo-udeps) [[cargo-udeps](https://crates.io/crates/cargo-udeps)] - find unused dependencies Stars:`1.9K`.
  * [cargo-release](https://crates.io/crates/cargo-release) - tool for releasing git-managed cargo project, build, tag, publish, doc and push [![Rust](https://github.com/crate-ci/cargo-release/actions/workflows/ci.yml/badge.svg)](https://github.com/crate-ci/cargo-release/actions/workflows/rust.yml) Stars:`1.4K`.
* CMake
* [facebook/buck2](https://github.com/facebook/buck2) - [Buck2](https://buck2.build/) is a large-scale build tool built in Rust Stars:`3.8K`.
* [Fleet](https://github.com/dimensionhq/fleet) [[fleet-rs](https://crates.io/crates/fleet-rs)] - The blazing fast build tool for Rust. Stars:`2.4K`.
* GitHub actions
* [pantsbuild/pants](https://github.com/pantsbuild/pants) - [Pants](https://www.pantsbuild.org/) is a fast, scalable, user-friendly build system for codebases of all sizes built in Rust. Stars:`3.5K`.
* [tracemachina/nativelink](https://github.com/tracemachina/nativelink) - [NativeLink](https://www.nativelink.com) is a Backend Remote Execution platform written in rust for client build systems such as [Buck2](https://buck2.build/), [Bazel](https://bazel.build/), [Pants](https://www.pantsbuild.org/), etc.. [![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/TraceMachina/nativelink/badge)](https://securityscorecards.dev/viewer/?uri=github.com/TraceMachina/nativelink) [![OpenSSF Best Practices](https://www.bestpractices.dev/projects/8050/badge)](https://www.bestpractices.dev/projects/8050) [![Slack](https://img.shields.io/badge/slack--channel-blue?logo=slack)](https://nativelink.slack.com/join/shared_invite/zt-281qk1ho0-krT7HfTUIYfQMdwflRuq7A#/shared-invite/email) Stars:`1.3K`.

### Debugging

* GDB
  * [gdbgui](https://github.com/cs01/gdbgui) - Browser based frontend for gdb to debug C, C++, Rust, and go. Stars:`10.0K`.
* LLDB

### Deployment

* Docker
  * [LukeMathWalker/cargo-chef](https://github.com/LukeMathWalker/cargo-chef) - A tool and pre-built images for caching compiling remote dependencies between Docker builds. Stars:`2.0K`.
  * [emk/rust-musl-builder](https://github.com/emk/rust-musl-builder) - Docker images for compiling static Rust binaries using musl-libc and musl-gcc, with static versions of useful C libraries Stars:`1.6K`.
* Heroku
* [release-plz](https://github.com/release-plz/release-plz) [[release-plz](https://crates.io/crates/release-plz)] - Release crates from CI, with changelog generation and semver check. [![build badge](https://github.com/release-plz/release-plz/workflows/CI/badge.svg)](https://github.com/release-plz/release-plz/actions) Stars:`1.0K`.

### Embedded

[Rust Embedded](https://rust-embedded.org/) focuses on improving the end-to-end experience of using Rust in resource-constrained environments and non-traditional platforms. See [awesome-embedded-rust](https://github.com/rust-embedded/awesome-embedded-rust) for a curated, and more extended list of embedded Rust resources.

* Arduino
* Cross compiling
  * [japaric/rust-cross](https://github.com/japaric/rust-cross) - everything you need to know about cross compiling Rust programs Stars:`2.5K`.
  * [japaric/xargo](https://github.com/japaric/xargo) - effortless cross compilation of Rust programs to custom bare-metal targets like ARM Cortex-M Stars:`1.1K`.
* Espressif
* Firmware
  * [oreboot/oreboot](https://github.com/oreboot/oreboot) - oreboot is a fork of coreboot, with C removed, written in Rust Stars:`1.7K`.
* nRF

### FFI

See also [Foreign Function Interface](https://doc.rust-lang.org/book/first-edition/ffi.html), [The Rust FFI Omnibus](http://jakegoulding.com/rust-ffi-omnibus/) (a collection of examples of using code written in Rust from other languages) and [FFI examples written in Rust](https://github.com/alexcrichton/rust-ffi-examples).

* C
  * [mozilla/cbindgen](https://github.com/mozilla/cbindgen) - generates C header files from Rust source files. Used in Gecko for WebRender Stars:`2.6K`.
* C#
* C++
  * [dtolnay/cxx](https://github.com/dtolnay/cxx) - Safe interop between Rust and C++ [![build badge](https://img.shields.io/badge/github-dtolnay/cxx-8da0cb?style=for-the-badge&labelColor=555555&logo=github)](https://github.com/dtolnay/cxx) Stars:`6.2K`.
  * [rust-lang/rust-bindgen](https://github.com/rust-lang/rust-bindgen) - A Rust bindings generator Stars:`4.7K`.
* Erlang
  * [rusterlium/rustler](https://github.com/rusterlium/rustler) - safe Rust bridge for creating Erlang NIF functions Stars:`4.5K`.
* Java
* Lua
  * [mlua-rs/mlua](https://github.com/mlua-rs/mlua) - High level Lua 5.4/5.3/5.2/5.1 (including LuaJIT) and Roblox Luau bindings to Rust with async/await support [![build badge](https://github.com/mlua-rs/mlua/workflows/CI/badge.svg)](https://github.com/mlua-rs/mlua/actions) Stars:`2.0K`.
* mruby
* Node.js
  * [neon-bindings/neon](https://github.com/neon-bindings/neon) - Rust bindings for writing safe and fast native Node.js modules Stars:`8.2K`.
* Objective-C
* PHP
* Prolog
  * [mthom/scryer-prolog](https://github.com/mthom/scryer-prolog/) - Scryer Prolog is a free software ISO Prolog system written in Rust Stars:`2.2K`.
* Python
  * [RustPython](https://github.com/RustPython/RustPython) - A Python Interpreter written in Rust [![Build Status](https://github.com/RustPython/RustPython/workflows/CI/badge.svg)](https://github.com/RustPython/RustPython/actions?query=workflow%3ACI) Stars:`19.9K`.
  * [PyO3/PyO3](https://github.com/PyO3/PyO3) - Rust bindings for the Python interpreter Stars:`13.4K`.
  * [dgrunwald/rust-cpython](https://github.com/dgrunwald/rust-cpython) - Python bindings Stars:`1.8K`.
* Ruby
* Web Assembly
  * [rustwasm/wasm-bindgen](https://github.com/rustwasm/wasm-bindgen) - A project for facilitating high-level interactions between wasm modules and JS. Stars:`8.1K`.
  * [rustwasm/wasm-pack](https://github.com/rustwasm/wasm-pack) - :package: :sparkles: pack up the wasm and publish it to npm! Stars:`6.6K`.

### Formatters

* [rustfmt](https://github.com/rust-lang/rustfmt) - Rust code formatter maintained by the Rust team and included in cargo Stars:`6.2K`.
* [dprint](https://github.com/dprint/dprint) - A pluggable and configurable code formatting platform [![build badge](https://github.com/dprint/dprint/workflows/CI/badge.svg)](https://github.com/dprint/dprint/actions?query=workflow%3ACI) Stars:`3.4K`.

### IDEs

See also [Are we (I)DE yet?](https://areweideyet.com/) and [Rust Tools](https://www.rust-lang.org/tools).

  * [lapce](https://github.com/lapce/lapce) - Lightning-fast and Powerful Code Editor written in Rust. [![build badge](https://github.com/lapce/lapce/actions/workflows/release.yml/badge.svg)](https://github.com/lapce/lapce/actions/workflows/release.yml) Stars:`35.1K`.
    * [intellij-rust/intellij-rust](https://github.com/intellij-rust/intellij-rust) - Rust plugin for the IntelliJ Platform Stars:`4.5K`.
    * [rust.vim](https://github.com/rust-lang/rust.vim) - provides file detection, syntax highlighting, formatting, Syntastic integration, and more. Stars:`4.0K`.
    * [autozimu/LanguageClient-neovim](https://github.com/autozimu/LanguageClient-neovim) - [LSP](https://microsoft.github.io/language-server-protocol/) client. Implemented in Rust and supports rls out of the box. Stars:`3.6K`.
    * [rust-mode](https://github.com/rust-lang/rust-mode) - Rust Major Mode Stars:`1.2K`.
  * Visual Studio

### Profiling

* [bheisler/criterion.rs](https://github.com/bheisler/criterion.rs) - Statistics-driven benchmarking library Stars:`4.9K`.
* [Bytehound](https://github.com/koute/bytehound) - A memory profiler for Linux Stars:`4.6K`.
* [Divan](https://github.com/nvzqz/divan) - Simple yet powerful benchmarking library with allocation profiling Stars:`1.1K`.
* FlameGraphs
* [sharkdp/hyperfine](https://github.com/sharkdp/hyperfine) - A command-line benchmarking tool Stars:`24.7K`.

### Services


### Static analysis

[[assert](https://crates.io/keywords/assert), [static](https://crates.io/keywords/static)]

* [verus-lang/verus](https://github.com/verus-lang/verus) - Verified Rust for low-level systems code Stars:`1.5K`.

### Testing

[[test](https://crates.io/keywords/test), [testing](https://crates.io/keywords/testing)]

* Code Coverage
* Continuous Integration
  * [trust](https://github.com/japaric/trust) - A Travis CI and AppVeyor template to test your Rust crate on 5 architectures and publish binary releases of it for Linux, macOS and Windows Stars:`1.3K`.
* Frameworks and Runners
  * [rstest](https://crates.io/crates/rstest) - Fixture-based test framework [![Build Status](https://github.com/la10736/rstest/workflows/Test/badge.svg?branch=master)](https://github.com/la10736/rstest/actions) Stars:`1.3K`.
* Mocking and Test Data
  * [asomers/mockall](https://github.com/asomers/mockall) [[mockall](https://crates.io/crates/mockall)] - A powerful mock object library. [![Cirrus Build Status](https://api.cirrus-ci.com/github/asomers/mockall.svg)](https://cirrus-ci.com/github/asomers/mockall) Stars:`1.6K`.
  * [synth](https://github.com/shuttle-hq/synth/) - Generate database data declaratively. [![build](https://github.com/shuttle-hq/synth/actions/workflows/synth-test.yml/badge.svg)](https://github.com/shuttle-hq/synth) Stars:`1.4K`.
  * [fake-rs](https://github.com/cksac/fake-rs) - A library for generating fake data Stars:`1.1K`.
* Mutation Testing
* Property Testing and Fuzzing
  * [rust-fuzz/afl.rs](https://github.com/rust-fuzz/afl.rs) - A Rust fuzzer, using [AFL](https://lcamtuf.coredump.cx/afl/) Stars:`1.7K`.

### Transpiling

* [immunant/c2rust](https://github.com/immunant/c2rust) - C to Rust translator and cross checker built atop Clang/LLVM. Stars:`4.2K`.
* [BayesWitnesses/m2cgen](https://github.com/BayesWitnesses/m2cgen) - A CLI tool to transpile trained classic machine learning models into a native Rust code with zero dependencies. [![GitHub Actions Status](https://github.com/BayesWitnesses/m2cgen/workflows/GitHub%20Actions/badge.svg?branch=master)](https://github.com/BayesWitnesses/m2cgen/actions) Stars:`2.9K`.
* [jameysharp/corrode](https://github.com/jameysharp/corrode) - A C to Rust translator written in Haskell. Stars:`2.2K`.

## Libraries


### Artificial Intelligence

#### Genetic algorithms


#### Machine learning

See [[Machine learning](https://crates.io/keywords/machine-learning)]

See also [About Rust’s Machine Learning Community](https://medium.com/@autumn_eng/about-rust-s-machine-learning-community-4cda5ec8a790#.hvkp56j3f) and [Are we learning yet?](https://www.arewelearningyet.com).

* [huggingface/candle](https://github.com/huggingface/candle) [[candle-core](https://crates.io/crates/candle-core)] - a minimalist ML framework with a focus on easiness of use and on performance (including GPU support) Stars:`17.0K`.
* [burn](https://github.com/tracel-ai/burn) - A Flexible and Comprehensive Deep Learning Framework. Stars:`10.2K`.
* [huggingface/tokenizers](https://github.com/huggingface/tokenizers) - Hugging Face's tokenizers for modern NLP pipelines (original implementation) with bindings for Python. [![Build Status](https://github.com/huggingface/tokenizers/workflows/Rust/badge.svg?branch=master)](https://github.com/huggingface/tokenizers/actions) Stars:`9.6K`.
* [autumnai/leaf](https://github.com/autumnai/leaf) - Open Machine Intelligence framework.. Abandoned project. The most updated fork is [juice](https://github.com/fff-rs/juice). Stars:`5.6K`.
* [tensorflow/rust](https://github.com/tensorflow/rust) - Bindings for TensorFlow. Stars:`5.3K`.
* [LaurentMazare/tch-rs](https://github.com/LaurentMazare/tch-rs) - Bindings for PyTorch. Stars:`4.7K`.
* [rust-ml/linfa](https://github.com/rust-ml/linfa) - Machine learning framework. Stars:`4.1K`.
* [guillaume-be/rust-bert](https://github.com/guillaume-be/rust-bert) [[rust_bert](https://crates.io/crates/rust_bert)] - Ready-to-use NLP pipelines and language models Stars:`2.8K`.
* [coreylowman/dfdx](https://github.com/coreylowman/dfdx) - CUDA accelerated machine learning framework that leverages many of Rust's unique features. ![Crates.io](https://img.shields.io/crates/v/dfdx) Stars:`1.8K`.

#### OpenAI

* [64bit/async-openai](https://github.com/64bit/async-openai) [[async-openai](https://crates.io/crates/async-openai)] - Ergonomic Rust bindings for OpenAI API based on OpenAPI spec. Stars:`1.4K`.

### Astronomy

[[astronomy](https://crates.io/keywords/astronomy)]


### Asynchronous

* [tokio-rs/tokio](https://github.com/tokio-rs/tokio) - A runtime for writing reliable, asynchronous, and slim applications with the Rust programming language. Stars:`28.3K`.
* [mio](https://github.com/tokio-rs/mio) - MIO is a lightweight IO library, with a focus on adding as little overhead as possible over the OS abstractions Stars:`6.6K`.
* [rust-lang/futures-rs](https://github.com/rust-lang/futures-rs) - Zero-cost futures Stars:`5.6K`.
* [async-std](https://async.rs/) [[async-std](https://crates.io/crates/async-std)] - Async version of the Rust standard library [![CI](https://github.com/async-rs/async-std/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/async-rs/async-std/actions/workflows/ci.yml) Stars:`4.0K`.
* [Xudong-Huang/may](https://github.com/Xudong-Huang/may) - Stackful coroutine library Stars:`2.1K`.

### Audio and Music

[[audio](https://crates.io/keywords/audio)]

  * [RustAudio/cpal](https://github.com/RustAudio/cpal) - Low-level cross-platform audio I/O library. [![Actions Status](https://github.com/RustAudio/cpal/workflows/cpal/badge.svg?branch=master)](https://github.com/RustAudio/cpal/actions) Stars:`3.0K`.
* [pdeljanov/Symphonia](https://github.com/pdeljanov/Symphonia) - Audio decoding and media demuxing library supporting AAC, FLAC, MP3, MP4, OGG, Vorbis, and WAV. Stars:`2.6K`.
  * [RustAudio/rodio](https://github.com/RustAudio/rodio) - Audio playback library Stars:`1.9K`.

### Authentication

* [Keats/jsonwebtoken](https://github.com/Keats/jsonwebtoken) - [JSON Web Token](https://en.wikipedia.org/wiki/JSON_Web_Token) library Stars:`1.8K`.

### Automotive


### Bioinformatics


### Caching

* [mozilla/sccache](https://github.com/mozilla/sccache/) - Shared Compilation Cache, great compilation Stars:`6.2K`.
* [moka-rs/moka](https://github.com/moka-rs/moka) - A high performance concurrent caching library inspired by the Caffeine library for Java [![build badge](https://github.com/moka-rs/moka/workflows/CI/badge.svg)](https://github.com/moka-rs/moka/actions/workflows/CI.yml) Stars:`1.9K`.
* [jaemk/cached](https://github.com/jaemk/cached) - Simple function caching/memoization Stars:`1.7K`.

### Cloud

* AWS [[aws](https://crates.io/keywords/aws)]
  * [awslabs/aws-lambda-rust-runtime](https://github.com/awslabs/aws-lambda-rust-runtime) [[lambda_runtime](https://crates.io/crates/lambda_runtime)] - Runtime for AWS Lambda [![build badge](https://github.com/awslabs/aws-lambda-rust-runtime/workflows/Rust/badge.svg)](https://github.com/awslabs/aws-lambda-rust-runtime/actions) Stars:`3.4K`.
  * [awslabs/aws-sdk-rust](https://github.com/awslabs/aws-sdk-rust) - The new AWS SDK Stars:`3.1K`.
  * [rusoto/rusoto](https://github.com/rusoto/rusoto) - An AWS SDK for Rust Stars:`2.7K`.
* Azure
* Load Balancer
* Multi Cloud
  * [Qovery/engine](https://github.com/Qovery/engine) - Abstraction layer library that turns easy application deployment on Cloud providers in just a few minutes Stars:`2.3K`.

### Command-line

* Argument parsing
  * [clap-rs](https://github.com/clap-rs/clap) [[clap](https://crates.io/crates/clap)] - A simple to use, full featured command-line argument parser Stars:`14.9K`.
  * [TeXitoi/structopt](https://github.com/TeXitoi/structopt) [[structopt](https://crates.io/crates/structopt)] - parse command line argument by defining a struct Stars:`2.7K`.
  * [google/argh](https://github.com/google/argh) [[argh](https://crates.io/crates/argh)] - An opinionated Derive-based argument parser optimized for code size [![build badge](https://github.com/google/argh/workflows/Argh/badge.svg?branch=master)](https://github.com/google/argh/actions) Stars:`1.8K`.
* Data visualization
  * [zhiburt/tabled](https://github.com/zhiburt/tabled) [[tabled](https://crates.io/crates/tabled)] - An easy to use library for pretty print tables of structs and enums. [![Build Status](https://github.com/zhiburt/tabled/actions/workflows/ci.yml/badge.svg)](https://github.com/zhiburt/tabled/actions) Stars:`2.1K`.
  * [nukesor/comfy-table](https://github.com/nukesor/comfy-table) [[comfy-table](https://crates.io/crates/comfy-table)] - Beautiful dynamic tables for your cli tools. [![Build status](https://github.com/Nukesor/comfy-table/workflows/Tests/badge.svg?branch=master)](https://github.com/nukesor/comfy-table/actions) Stars:`1.1K`.
* Human-centered design
  * [rust-cli/human-panic](https://github.com/rust-cli/human-panic) [[human-panic](https://crates.io/crates/human-panic)] - panic messages for humans Stars:`1.7K`.
* Line editor
  * [kkawakam/rustyline](https://github.com/kkawakam/rustyline) [[rustyline](https://crates.io/crates/rustyline)] - readline implementation Stars:`1.6K`.
* Other
* Pipeline
* Progress
  * [console-rs/indicatif](https://github.com/console-rs/indicatif) [[indicatif](https://crates.io/crates/indicatif)] - indicate progress to users Stars:`4.7K`.
* Prompt
  * [starship/starship](https://starship.rs/) [[starship](https://crates.io/crates/starship)] - A minimal, blazing fast, and extremely customizable prompt for any shell [![Build status](https://github.com/starship/starship/actions/workflows/workflow.yml/badge.svg)](https://github.com/starship/starship/actions) Stars:`47.9K`.
  * [mikaelmello/inquire](https://github.com/mikaelmello/inquire) [[inquire](https://crates.io/crates/inquire)] - A library for building interactive prompts on terminals. [![Build status](https://github.com/mikaelmello/inquire/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/mikaelmello/inquire/actions) Stars:`2.2K`.
* Style
  * [colored](https://github.com/colored-rs/colored) [[colored](https://crates.io/crates/colored)] - Coloring terminal so simple, you already know how to do it! Stars:`1.8K`.
  * [console-rs/dialoguer](https://github.com/console-rs/dialoguer) [[dialoguer](https://crates.io/crates/dialoguer)] - Library for command line prompts and similar things. Stars:`1.4K`.
* TUI
  * BearLibTerminal
  * [gyscos/Cursive](https://github.com/gyscos/Cursive) [[cursive](https://crates.io/crates/cursive)] - build rich TUI applications Stars:`4.5K`.
  * ncurses
  * [ratatui-org/ratatui](https://github.com/ratatui/ratatui) [[ratatui](https://crates.io/crates/ratatui)] - Library that's all about cooking up terminal user interfaces (TUIs) Stars:`12.5K`.
  * [redox-os/termion](https://github.com/redox-os/termion) [[termion](https://crates.io/crates/termion)] - bindless library for controlling terminals/TTY Stars:`2.1K`.
  * Termbox
  * [TimonPost/crossterm](https://github.com/crossterm-rs/crossterm) [[crossterm](https://crates.io/crates/crossterm)] - crossplatform terminal library Stars:`3.5K`.

### Compression

* bzip2
* gzip
* gzp
* miniz
* tar
* zip
* zstd

### Computation

* [dimforge/nalgebra](https://github.com/dimforge/nalgebra) - low-dimensional linear algebra library Stars:`4.2K`.
* [faer-rs](https://github.com/sarah-quinones/faer-rs) [[faer](https://crates.io/crates/faer)] - Linear algebra foundation for Rust Stars:`2.1K`.
* [calebwin/emu](https://github.com/calebwin/emu) - A language for GPGPU numerical computing Stars:`1.6K`.
* [argmin-rs/argmin](https://github.com/argmin-rs/argmin) [[argmin](https://crates.io/crates/argmin)] - Optimization library Stars:`1.1K`.
* Parallel
* Science
* Statrs

### Concurrency

* [Rayon](https://github.com/rayon-rs/rayon) - A data parallelism library Stars:`11.6K`.
* [crossbeam-rs/crossbeam](https://github.com/crossbeam-rs/crossbeam) - Support for parallelism and low-level concurrency Stars:`7.8K`.

### Configuration

* [rust-cli/config-rs](https://github.com/rust-cli/config-rs) [[config](https://crates.io/crates/config)] - Layered configuration system (with strong support for 12-factor applications). Stars:`2.8K`.

### Cryptography

[[crypto](https://crates.io/keywords/crypto), [cryptography](https://crates.io/keywords/cryptography)]

* [rustls/rustls](https://github.com/rustls/rustls) - Implementation of TLS Stars:`6.6K`.
* [briansmith/ring](https://github.com/briansmith/ring) - Safe, fast, small crypto using Rust and BoringSSL's cryptography primitives. Stars:`3.9K`.
* [RustCrypto/hashes](https://github.com/RustCrypto/hashes) - Collection of cryptographic hash functions Stars:`2.0K`.
* [cossacklabs/themis](https://github.com/cossacklabs/themis) [[themis](https://crates.io/crates/themis)] - a high-level cryptographic library for solving typical data security tasks, best fit for multi-platform apps. [![build badge](https://circleci.com/gh/cossacklabs/themis/tree/master.svg?style=shield)](https://app.circleci.com/pipelines/github/cossacklabs/themis) Stars:`1.9K`.
* [sfackler/rust-openssl](https://github.com/sfackler/rust-openssl) - [OpenSSL](https://www.openssl.org/) bindings Stars:`1.5K`.
* [DaGenix/rust-crypto](https://github.com/DaGenix/rust-crypto) - cryptographic algorithms Stars:`1.4K`.
* [exonum/exonum](https://github.com/exonum/exonum) [[exonum](https://crates.io/crates/exonum)] - extensible framework for blockchain projects Stars:`1.2K`.

### Data processing

* [pola-rs/polars](https://github.com/pola-rs/polars) - Fast feature complete DataFrame library ![Build and test](https://github.com/pola-rs/polars/workflows/Build%20and%20test/badge.svg?branch=master) Stars:`32.9K`.
* [datafusion](https://github.com/apache/datafusion) - DataFusion is a very fast, extensible query engine for building high-quality data-centric systems in Rust, using the Apache Arrow in-memory format. Stars:`7.0K`.
* [pg_analytics](https://github.com/paradedb/paradedb/tree/dev/pg_analytics) - PostgreSQL extension that accelerates analytical query processing inside Postgres to a performance level comparable to dedicated OLAP databases. Stars:`6.9K`.
* [pg_lakehouse](https://github.com/paradedb/paradedb/tree/dev/pg_lakehouse) - PostgreSQL extension that transforms Postgres into an analytical query engine over object stores like AWS S3/GCS and table formats like Delta Lake/Iceberg. Stars:`6.9K`.
* [bluss/ndarray](https://github.com/rust-ndarray/ndarray) - N-dimensional array with array views, multidimensional slicing, and efficient operations Stars:`3.8K`.
* [weld-project/weld](https://github.com/weld-project/weld) - High-performance runtime for data analytics applications Stars:`3.0K`.

### Data streaming

* [infinyon/fluvio](https://github.com/infinyon/fluvio) - Programmable data streaming platform [![CI](https://github.com/infinyon/fluvio/workflows/CI/badge.svg?branch=stable)](https://github.com/infinyon/fluvio/actions) Stars:`4.4K`.
* [ArroyoSystems/arroyo](https://github.com/ArroyoSystems/arroyo) - High-performance real-time analytics in Rust and SQL [![CI](https://github.com/ArroyoSystems/arroyo/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/ArroyoSystems/arroyo/actions) Stars:`4.1K`.
* [iggy](https://github.com/apache/iggy) [[iggy](https://crates.io/crates/iggy)] - Persistent message streaming platform, supporting QUIC, TCP and HTTP transport protocols [![CI](https://github.com/apache/iggy/actions/workflows/test.yml/badge.svg)](https://github.com/apache/iggy/actions/workflows/test.yml) Stars:`2.4K`.

### Data structures

* [rust-itertools/itertools](https://github.com/rust-itertools/itertools) - Extra iterator adaptors, functions and macros Stars:`2.9K`.
* [greyblake/nutype](https://github.com/greyblake/nutype) [[nutype](https://crates.io/crates/nutype)] - define newtype structures with validation constraints. [![build status](https://github.com/greyblake/nutype/actions/workflows/ci.yml/badge.svg)](https://github.com/greyblake/nutype/actions) Stars:`1.5K`.
* [orium/rpds](https://github.com/orium/rpds) [[rpds](https://crates.io/crates/rpds)] - Persistent data structures. [![build badge](https://github.com/orium/rpds/workflows/CI/badge.svg)](https://github.com/orium/rpds/actions?query=workflow%3ACI) Stars:`1.4K`.
* [ashvardanian/simsimd](https://github.com/ashvardanian/SimSIMD) - SIMD-accelerated vector distances and similarity functions for x86 AVX2 & AVX-512, and Arm NEON [![crates.io](https://img.shields.io/crates/v/simsimd.svg)](https://crates.io/crates/simsimd) Stars:`1.3K`.

### Data visualization

* [rerun](https://github.com/rerun-io/rerun) - [[rerun](https://crates.io/crates/rerun)] - An SDK for logging computer vision and robotics data (tensors, point clouds, etc) paired with a visualizer for exploring that data over time. Stars:`8.1K`.
* [plotters](https://github.com/plotters-rs/plotters) - [![build badge](https://github.com/plotters-rs/plotters/workflows/CI/badge.svg)](https://github.com/plotters-rs/plotters/actions) Stars:`4.1K`.
* [plotly](https://github.com/plotly/plotly.rs) - Plotly for Rust Stars:`1.3K`.

### Database

[[database](https://crates.io/keywords/database)]

* NoSQL [[nosql](https://crates.io/keywords/nosql)]

  * CouchDB [[couchdb](https://crates.io/keywords/couchdb)]
  * Elasticsearch [[elasticsearch](https://crates.io/keywords/elasticsearch)]
  * etcd
  * LevelDB
  * LMDB [[lmdb](https://crates.io/keywords/lmdb)]
  * MongoDB [[mongodb](https://crates.io/keywords/mongodb)]
    * [Redb](https://github.com/cberner/redb) - An embedded key-value database. It provides a similar interface to other embedded key-value stores such as rocksdb and lmdb. ![GitHub Workflow Status](https://github.com/cberner/redb/actions/workflows/ci.yml/badge.svg) Stars:`3.6K`.
    * [mongodb/mongo-rust-driver](https://github.com/mongodb/mongo-rust-driver) [[mongodb](https://crates.io/crates/mongodb)] - [MongoDB](https://www.mongodb.com/) bindings Stars:`1.5K`.
    * [PoloDB](https://github.com/PoloDB/PoloDB) - An embedded JSON-based database has API similar to MongoDB. ![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/PoloDB/PoloDB/rust.yml) Stars:`1.0K`.
  * Redis [[redis](https://crates.io/keywords/redis)]
    * [surrealdb/surrealdb](https://github.com/surrealdb/surrealdb) - SurrealDB embedded document-graph database Stars:`29.0K`.
    * [redis-rs](https://github.com/redis-rs/redis-rs) - [Redis](https://redis.io/) library [![Rust](https://github.com/redis-rs/redis-rs/actions/workflows/rust.yml/badge.svg)](https://github.com/redis-rs/redis-rs/actions/workflows/rust.yml) Stars:`3.8K`.
  * [UnQLite](https://github.com/symisc/unqlite) Stars:`2.2K`.
    * [rust-rocksdb/rust-rocksdb](https://github.com/rust-rocksdb/rust-rocksdb) - RocksDB bindings [![RocksDB CI](https://github.com/rust-rocksdb/rust-rocksdb/actions/workflows/rust.yml/badge.svg?branch=master)](https://github.com/rust-rocksdb/rust-rocksdb/actions/workflows/rust.yml) Stars:`2.0K`.
* OGM [[ogm](https://crates.io/keywords/ogm)]
* ORM [[orm](https://crates.io/keywords/orm)]
  * [diesel-rs/diesel](https://github.com/diesel-rs/diesel) - an ORM and Query builder Stars:`13.2K`.
  * [SeaQL/sea-orm](https://github.com/SeaQL/sea-orm) - 🐚 An async & dynamic ORM  [![crate](https://img.shields.io/crates/v/sea-orm.svg)](https://crates.io/crates/sea-orm) [![docs](https://img.shields.io/docsrs/sea-orm/latest)](https://docs.rs/sea-orm) [![build status](https://github.com/SeaQL/sea-orm/actions/workflows/rust.yml/badge.svg)](https://github.com/SeaQL/sea-orm/actions/workflows/rust.yml) Stars:`8.0K`.
  * [rbatis/rbatis](https://github.com/rbatis/rbatis) - ORM Framework High Performance(JSON based) Stars:`2.3K`.
  * [Brendonovich/prisma-client-rust](https://github.com/Brendonovich/prisma-client-rust) - An autogenerated query builder that provides simple and fully type-safe database access using the Prisma ecosystem. [![Test Status](https://img.shields.io/github/workflow/status/Brendonovich/prisma-client-rust/CI?label=tests&style=flat-square)](https://github.com/Brendonovich/prisma-client-rust/actions) Stars:`1.9K`.
* [sfackler/r2d2](https://github.com/sfackler/r2d2) - generic connection pool Stars:`1.6K`.
* SQL [[sql](https://crates.io/keywords/sql)]
  * Generic
    * [launchbadge/sqlx](https://github.com/launchbadge/sqlx) - async PostgreSQL/MySQL/SQLite connection pool with strong typing support [![build badge](https://img.shields.io/github/workflow/status/launchbadge/sqlx/Rust/master?style=flat-square)](https://github.com/launchbadge/sqlx) Stars:`14.5K`.
    * [SeaQL/sea-query](https://github.com/SeaQL/sea-query) - 🔱 A dynamic SQL query builder for MySQL, Postgres and SQLite [![crate](https://img.shields.io/crates/v/sea-query.svg)](https://crates.io/crates/sea-query) [![docs](https://img.shields.io/docsrs/sea-query/latest)](https://docs.rs/sea-query) [![build status](https://github.com/SeaQL/sea-query/actions/workflows/rust.yml/badge.svg)](https://github.com/SeaQL/sea-query/actions/workflows/rust.yml) Stars:`1.3K`.
  * Microsoft SQL
  * MySql [[mysql](https://crates.io/keywords/mysql)]
  * Oracle
  * PostgreSql [[postgres](https://crates.io/keywords/postgres), [postgresql](https://crates.io/keywords/postgresql)]
    * [sfackler/rust-postgres](https://github.com/sfackler/rust-postgres) [[postgres](https://crates.io/crates/postgres)] - A native [PostgreSQL](https://www.postgresql.org/) client Stars:`3.7K`.
  * Sqlite [[sqlite](https://crates.io/keywords/sqlite)]
    * [rusqlite](https://github.com/rusqlite/rusqlite) - [Sqlite3](https://www.sqlite.org/index.html) bindings Stars:`3.5K`.

### Date and time

[[date](https://crates.io/keywords/date), [time](https://crates.io/keywords/time)]

* [chronotope/chrono](https://github.com/chronotope/chrono) - Date and time library Stars:`3.5K`.
* [burntSushi/jiff](https://github.com/BurntSushi/jiff) - A date-time library for Rust that encourages you to jump into the pit of success. [![Build status](https://github.com/BurntSushi/jiff/workflows/ci/badge.svg)](https://github.com/BurntSushi/jiff/actions) Stars:`2.1K`.
* [time-rs/time](https://github.com/time-rs/time) - [![build badge](https://github.com/time-rs/time/workflows/Build/badge.svg)](https://github.com/time-rs/time/actions) Stars:`1.2K`.

### Distributed systems

* Antimony
* Apache Kafka
  * [fede1024/rust-rdkafka](https://github.com/fede1024/rust-rdkafka) [[rdkafka](https://crates.io/crates/rdkafka)] - [librdkafka](https://github.com/confluentinc/librdkafka) bindings Stars:`1.7K`.
  * [kafka-rust/kafka-rust](https://github.com/kafka-rust/kafka-rust) - Rust client for Apache Kafka Stars:`1.3K`.
* HDFS
* Other
  * [build-trust/ockam](https://github.com/build-trust/ockam) [[ockam](https://crates.io/crates/ockam)] - End-to-End Encryption, Mutual Authentication, and ABAC for distributed applications [![build badge](https://github.com/build-trust/ockam/workflows/Rust/badge.svg)](https://github.com/build-trust/ockam) Stars:`4.5K`.

### Domain driven design


### eBPF

* [aya/aya-rs](https://github.com/aya-rs/aya) - Built with a focus on developer experience and operability. Stars:`3.5K`.

### Email

[[email](https://crates.io/keywords/email), [imap](https://crates.io/keywords/imap), [smtp](https://crates.io/keywords/smtp)]

* [lettre/lettre](https://github.com/lettre/lettre) - an SMTP-library [![CI](https://github.com/lettre/lettre/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/lettre/lettre/actions/workflows/test.yml) Stars:`2.0K`.

### Encoding

[[encoding](https://crates.io/keywords/encoding)]

* ASN.1
* Binary
  * [bincode-org/bincode](https://github.com/bincode-org/bincode) - A binary encoder/decoder [![CI](https://github.com/bincode-org/bincode/actions/workflows/rust.yml/badge.svg?branch=trunk)](https://github.com/bincode-org/bincode/actions/workflows/rust.yml) Stars:`2.9K`.
  * [m4b/goblin](https://github.com/m4b/goblin) [[goblin](https://crates.io/crates/goblin)] - cross-platform, zero-copy, and endian-aware binary parsing Stars:`1.3K`.
  * [jamesmunns/postcard](https://github.com/jamesmunns/postcard) [[postcard](https://crates.io/crates/postcard)] - Postcard is a #![no_std] focused serializer and deserializer for Serde. Stars:`1.1K`.
* BSON
* Byte swapping
  * [BurntSushi/byteorder](https://github.com/BurntSushi/byteorder) - Supports big-endian, little-endian and native byte orders Stars:`1.0K`.
* Cap'n Proto
  * [capnproto/capnproto-rust](https://github.com/capnproto/capnproto-rust) - Cap'n Proto is a type system for distributed systems Stars:`2.2K`.
* CBOR
* Character Encoding
* CRC
* CSV
  * [BurntSushi/rust-csv](https://github.com/BurntSushi/rust-csv) - A fast and flexible CSV reader and writer, with support for Serde Stars:`1.8K`.
* EDN
* HAR
* HTML
  * [servo/html5ever](https://github.com/servo/html5ever) - High-performance browser-grade HTML5 parser Stars:`2.2K`.
* JSON
  * [serde-rs/json](https://github.com/serde-rs/json) [[serde\_json](https://crates.io/crates/serde_json)] - JSON support for [Serde](https://github.com/serde-rs/serde) framework Stars:`5.1K`.
  * [simd-lite/simd-json](https://github.com/simd-lite/simd-json) [[simd-json](https://crates.io/crates/simd-json)] - High performance JSON parser based on a port of simdjson Stars:`1.2K`.
* MsgPack
  * [3Hren/msgpack-rust](https://github.com/3Hren/msgpack-rust) - Low/high level MessagePack implementation Stars:`1.2K`.
* NetCDF
* PEM
* ProtocolBuffers
  * [tokio-rs/prost](https://github.com/tokio-rs/prost) - [![continuous integration](https://github.com/tokio-rs/prost/workflows/continuous%20integration/badge.svg?branch=master)](https://github.com/tokio-rs/prost/actions) Stars:`4.2K`.
  * [stepancheg/rust-protobuf](https://github.com/stepancheg/rust-protobuf) - Rust implementation of Google protocol buffers Stars:`2.9K`.
* rkyv
  * [rkyv/rkyv](https://github.com/rkyv/rkyv) [[rkyv](https://crates.io/crates/rkyv)] - rkyv (archive) is a zero-copy deserialization framework Stars:`3.2K`.
* RON (Rusty Object Notation)
  * [https://github.com/ron-rs/ron](https://github.com/ron-rs/ron) - Rusty Object Notation Stars:`3.5K`.
* Serde
* TOML
  * [tamasfe/taplo](https://github.com/tamasfe/taplo) [[taplo](https://crates.io/crates/taplo)] - A TOML toolkit [![CI](https://github.com/tamasfe/taplo/workflows/Continuous%20integration/badge.svg)](https://github.com/tamasfe/taplo/actions?query=workflow%3A%22Continuous+integration%22) Stars:`1.7K`.
* XML
  * [tafia/quick-xml](https://github.com/tafia/quick-xml) - High performance XML pull reader/writer Stars:`1.3K`.
* YAML

### Filesystem

[[filesystem](https://crates.io/keywords/filesystem)]
* Operations
  * [OpenDAL](https://github.com/apache/opendal) [[opendal](https://crates.io/crates/opendal)] - A unified data access layer, empowering users to seamlessly and efficiently retrieve data from diverse storage services. [![build](https://img.shields.io/github/actions/workflow/status/apache/opendal/ci_core.yml?branch=main)](https://github.com/apache/opendal/actions?query=branch%3Amain) Stars:`4.0K`.
* Temporary Files
  * [zboxfs/zbox](https://github.com/zboxfs/zbox) [[zbox](https://crates.io/crates/zbox)] - Zero-details, privacy-focused embeddable file system. Stars:`1.5K`.
  * [Stebalien/tempfile](https://github.com/Stebalien/tempfile) - temporary file library Stars:`1.3K`.

### Finance

* [avhz/RustQuant](https://github.com/avhz/RustQuant) [[RustQuant](https://crates.io/crates/RustQuant)] - A quantitative finance library. ![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/avhz/RustQuant/build.yml) Stars:`1.3K`.

### Functional Programming

[[functional programming](https://crates.io/keywords/fp)]
* Prelude
  * [JasonShin/fp-core.rs](https://github.com/JasonShin/fp-core.rs) - A library for functional programming Stars:`1.4K`.

### Game development

See also [Are we game yet?](https://arewegameyet.rs)
* Allegro
* bracket-lib (previously RLTK)
  * [bracket-lib](https://github.com/amethyst/bracket-lib) [[bracket-lib](https://crates.io/crates/bracket-lib)] - The Roguelike Toolkit (RLTK). [![Rust](https://github.com/amethyst/bracket-lib/actions/workflows/rust.yml/badge.svg)](https://github.com/amethyst/bracket-lib/actions/workflows/rust.yml) Stars:`1.6K`.
* Challonge
* Entity-Component Systems (ECS)
  * [amethyst/specs](https://github.com/amethyst/specs) - Specs Parallel ECS Stars:`2.6K`.
  * [legion](https://github.com/amethyst/legion) - A feature rich high performance ECS library with minimal boilerplate [![build badge](https://github.com/amethyst/legion/workflows/CI/badge.svg?branch=master)](https://github.com/amethyst/legion/actions) Stars:`1.7K`.
* Game Engines
  * [Bevy](https://github.com/bevyengine/bevy) - is a refreshingly simple data-driven game engine. - [![Crates.io](https://img.shields.io/crates/v/bevy.svg)](https://crates.io/crates/bevy) [![Crates.io](https://img.shields.io/crates/d/bevy.svg)](https://crates.io/crates/bevy) Stars:`38.8K`.
  * [Fyrox](https://fyrox.rs/) - Game engine 3D [![Crates.io](https://img.shields.io/crates/v/fyrox.svg)](https://crates.io/crates/fyrox) [![license](https://img.shields.io/crates/l/fyrox.svg)](https://github.com/FyroxEngine/Fyrox/blob/master/LICENSE.md) [![Crates.io](https://img.shields.io/crates/d/fyrox.svg)](https://crates.io/crates/fyrox) Stars:`8.2K`.
  * [Piston](https://www.piston.rs/) - [![Crates.io](https://img.shields.io/crates/v/piston.svg?style=flat-square)](https://crates.io/crates/piston) [![Crates.io](https://img.shields.io/crates/l/piston.svg)](https://github.com/PistonDevelopers/piston/blob/master/LICENSE) [![Crates.io](https://img.shields.io/crates/d/piston.svg)](https://crates.io/crates/piston) Stars:`4.7K`.
  * [ggez](https://github.com/ggez/ggez) - A lightweight game framework for making 2D games with minimum friction - [![Crates.io](https://img.shields.io/crates/v/ggez.svg)](https://crates.io/crates/ggez) [![license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/ggez/ggez/blob/master/LICENSE) [![Crates.io](https://img.shields.io/crates/d/ggez.svg)](https://crates.io/crates/ggez) Stars:`4.4K`.
* Game Servers
  * [godot-rust/gdnative](https://github.com/godot-rust/gdnative) [[gdnative](https://crates.io/crates/gdnative)] - Bindings to the Godot 3+ game engine [![CI](https://github.com/godot-rust/gdnative/actions/workflows/full-ci.yml/badge.svg)](https://github.com/godot-rust/gdnative/actions/workflows/full-ci.yml) Stars:`3.6K`.
  * [godot-rust/gdext](https://github.com/godot-rust/gdext) [[gdext](https://crates.io/crates/gdext)] - Bindings to the Godot 4+ game engine [![CI](https://github.com/godot-rust/gdext/actions/workflows/full-ci.yml/badge.svg)](https://github.com/godot-rust/gdext/actions/workflows/full-ci.yml) Stars:`3.5K`.
* Minecraft
  * [Pumpkin](https://github.com/pumpkin-mc/pumpkin) - A high-performance Minecraft server Software fully written in Rust Stars:`4.3K`.
  * [Rust-SDL2/rust-sdl2](https://github.com/Rust-SDL2/rust-sdl2) - SDL2 bindings Stars:`2.8K`.
* SFML
* Skillratings
* Toornament-rs
* Victorem

### Geospatial

[[geo](https://crates.io/keywords/geo), [gis](https://crates.io/keywords/gis)]

* [MapLibre/Martin](https://github.com/maplibre/martin) - Map tile server with PostGIS, MBTiles, PMTiles, and sprites support. [![CI build](https://github.com/maplibre/martin/actions/workflows/ci.yml/badge.svg)](https://github.com/maplibre/martin/actions)[![crates.io version](https://img.shields.io/crates/v/martin.svg)](https://crates.io/crates/martin)[![Book](https://img.shields.io/badge/docs-Book-informational)](https://maplibre.org/martin/) Stars:`2.6K`.

### Graph algorithms

* [petgraph/petgraph](https://github.com/petgraph/petgraph) - Graph data structure library. [![graph CI status](https://github.com/petgraph/petgraph/workflows/Continuous%20integration/badge.svg?branch=master)](https://github.com/petgraph/petgraph/actions/workflows/ci.yml) Stars:`3.2K`.

### Graphics

[[graphics](https://crates.io/keywords/graphics)]

* Fonts
* [gfx-rs/wgpu](https://github.com/gfx-rs/wgpu) - Native WebGPU implementation based on gfx-hal. [![build badge](https://github.com/gfx-rs/wgpu/workflows/CI/badge.svg?branch=master)](https://github.com/gfx-rs/wgpu/actions) Stars:`13.8K`.
* [gfx-rs/gfx](https://github.com/gfx-rs/gfx) - A high-performance, bindless graphics API. Stars:`5.4K`.
* OpenGL [[opengl](https://crates.io/keywords/opengl)]
  * [glium/glium](https://github.com/glium/glium) - safe OpenGL wrapper. Stars:`3.6K`.
* PDF
  * [vulkano](https://github.com/vulkano-rs/vulkano) [[vulkano](https://crates.io/crates/vulkano)] - Safe and rich Rust wrapper around the Vulkan API Stars:`4.7K`.
  * [J-F-Liu/lopdf](https://github.com/J-F-Liu/lopdf) - PDF document manipulation Stars:`1.8K`.

### GUI

[[gui](https://crates.io/keywords/gui)]

* Cocoa
* [tauri-apps/tauri](https://github.com/tauri-apps/tauri) - Build smaller, faster, and more secure desktop applications with a web frontend, powered by [WRY](https://github.com/tauri-apps/wry). [![test library](https://img.shields.io/github/workflow/status/tauri-apps/tauri/test%20library?label=test%20library)](https://github.com/tauri-apps/tauri/actions?query=workflow%3A%22test+library%22) Stars:`91.2K`.
* [ImGui](https://github.com/ocornut/imgui) Stars:`64.4K`.
* [DioxusLabs/dioxus](https://github.com/dioxuslabs/dioxus) - a portable, performant, and ergonomic framework for building cross-platform user interfaces in Rust. ![rust ci](https://github.com/dioxuslabs/dioxus/actions/workflows/main.yml/badge.svg) Stars:`26.4K`.
* [iced-rs/iced](https://github.com/iced-rs/iced) [[iced](https://crates.io/crates/iced)] - A cross-platform GUI library, focused on simplicity and type-safety. Inspired by Elm. Stars:`26.2K`.
* [emilk/egui](https://github.com/emilk/egui) - Simple, fast, and highly portable immediate mode GUI library. egui runs on the web, natively, and in your favorite game engine. [![Build Status](https://github.com/emilk/egui/workflows/CI/badge.svg)](https://github.com/emilk/egui/actions?workflow=CI) Stars:`24.3K`.
* [slint-ui/slint](https://github.com/slint-ui/slint) [slint](https://crates.io/crates/slint) - [Slint](https://slint.dev/) is a toolkit to efficiently develop fluid graphical user interfaces for embedded devices and desktop applications. [![Build Status](https://github.com/slint-ui/slint/workflows/CI/badge.svg?branch=master)](https://github.com/slint-ui/slint/actions?query=workflow%3ACI) Stars:`18.8K`.
* [libui](https://github.com/andlabs/libui) Stars:`10.8K`.
* [Nuklear](https://github.com/Immediate-Mode-UI/Nuklear) Stars:`9.8K`.
* [fschutt/azul](https://github.com/fschutt/azul) - A free, functional, IMGUI-oriented GUI framework for rapid development of desktop applications written in Rust, supported by the Mozilla WebRender rendering engine. Stars:`6.0K`.
* [makepad/makepad](https://github.com/makepad/makepad) [[makepad-widgets](https://crates.io/crates/makepad-widgets)] - Makepad is a creative software development platform that compiles to wasm/webGL, osx/metal, windows/dx11 linux/opengl. Stars:`5.3K`.
  * [fzyzcjy/flutter_rust_bridge](https://github.com/fzyzcjy/flutter_rust_bridge) - High-level memory-safe binding generator for Flutter/Dart <-> Rust Stars:`4.6K`.
* [xilem](https://github.com/linebender/xilem) - Successor of the data-first UI design toolkit [druid](https://github.com/linebender/druid). Stars:`4.1K`.
* [tauri-apps/wry](https://github.com/tauri-apps/wry) - Webview Rendering librarY. Stars:`4.0K`.
* [OrbTk](https://github.com/redox-os/orbtk) - The Orbital Widget Toolkit is a multi platform (G)UI toolkit using SDL2 [![Build and test](https://github.com/redox-os/orbtk/workflows/build/badge.svg?branch=develop)](https://github.com/redox-os/orbtk/actions) Stars:`3.8K`.
* [PistonDevelopers/conrod](https://github.com/PistonDevelopers/conrod/) - An easy-to-use, immediate-mode, 2D GUI library Stars:`3.4K`.
  * [imgui-rs](https://github.com/imgui-rs/imgui-rs) - Bindings for ImGui [![Build Status](https://github.com/imgui-rs/imgui-rs/workflows/ci/badge.svg?branch=master)](https://github.com/imgui-rs/imgui-rs/actions) Stars:`2.8K`.
  * [relm](https://github.com/antoyo/relm) - Asynchronous, GTK+-based, GUI library, inspired by Elm Stars:`2.4K`.
  * [cunarist/rinf](https://github.com/cunarist/rinf) - Rust as your Flutter backend, Flutter as your Rust frontend [![Build Test](https://github.com/cunarist/rinf/actions/workflows/build_test.yaml/badge.svg)](https://github.com/cunarist/rinf/actions/workflows/build_test.yaml?query=branch%3Amain) Stars:`2.3K`.
  * [flutter-rs](https://github.com/flutter-rs/flutter-rs) - Build flutter desktop app in dart & rust. Stars:`2.1K`.
  * [gtk-rs/gtk4-rs](https://github.com/gtk-rs/gtk4-rs) - GTK4 binding ![CI](https://github.com/gtk-rs/gtk4-rs/workflows/CI/badge.svg) Stars:`2.0K`.
  * [fltk-rs](https://github.com/fltk-rs/fltk-rs) - FLTK bindings [![Build](https://github.com/fltk-rs/fltk-rs/workflows/Build/badge.svg?branch=master)](https://github.com/fltk-rs/fltk-rs/actions) Stars:`1.7K`.
  * [servo/core-foundation-rs](https://github.com/servo/core-foundation-rs) - Rust bindings to Core Foundation and other low level libraries on Mac OS X and iOS Stars:`1.1K`.
* [emoon/rust_minifb](https://github.com/emoon/rust_minifb) - minifb is a cross-platform window setup with optional bitmap rendering. It also comes with easy mouse and keyboard input. Primarily designed for prototyping Stars:`1.1K`.
* [Ribir](https://github.com/RibirX/Ribir) - Ribir is a Rust GUI framework that helps you build beautiful and native multi-platform applications from a single codebase. Stars:`1.1K`.

### Image processing

* [image-rs/image](https://github.com/image-rs/image) - Basic imaging processing functions and methods for converting to and from image formats Stars:`5.2K`.
* [twistedfall/opencv-rust](https://github.com/twistedfall/opencv-rust) - Bindings for OpenCV Stars:`2.1K`.

### Language specification


### Logging

[[log](https://crates.io/keywords/log)]

* [donnie4w/tklog](https://github.com/donnie4w/tklog "donnie4w/tklog") - lightweight and efficient rust structured log library with support for log levels, file segmentation, compressed archiving.
* [tokio-rs/tracing](https://github.com/tokio-rs/tracing) - An application level tracing framework for async-aware structured logging, error handling, metrics, and more [![Build Status](https://github.com/tokio-rs/tracing/workflows/CI/badge.svg?branch=master)](https://github.com/tokio-rs/tracing/actions?query=workflow%3ACI) Stars:`5.8K`.
* [rust-lang/log](https://github.com/rust-lang/log) - Logging implementation Stars:`2.3K`.
* [slog-rs/slog](https://github.com/slog-rs/slog) - Structured, composable logging Stars:`1.6K`.
* [estk/log4rs](https://github.com/estk/log4rs) - highly configurable logging framework modeled after Java's Logback and log4j libraries [![CircleCI](https://circleci.com/gh/estk/log4rs.svg?style=shield)](https://app.circleci.com/pipelines/github/estk/log4rs) Stars:`1.0K`.

### Macro

* cute
* [elastio/bon](https://github.com/elastio/bon) [[bon](https://crates.io/crates/bon)] - generate compile-time-checked builders for structs and functions, provides partial application, optional and named parameters for functions and methods. [![build status](https://github.com/elastio/bon/actions/workflows/ci.yml/badge.svg)](https://github.com/elastio/bon/actions) Stars:`1.5K`.

### Markup language

* CommonMark
  * [pulldown-cmark/pulldown-cmark](https://github.com/pulldown-cmark/pulldown-cmark) - [CommonMark](https://commonmark.org/) parser Stars:`2.2K`.

### Mobile

* Android / iOS
* Generic
  * [redbadger/crux](https://github.com/redbadger/crux) [[crux_core](https://crates.io/crates/crux_core)] - Cross-platform app development. Crux helps you share your app's business logic and behavior across mobile (iOS/Android) and web - as a single reusable core. [![Build status](https://img.shields.io/github/actions/workflow/status/redbadger/crux/build.yaml)](https://github.com/redbadger/crux/actions) Stars:`1.9K`.
* iOS

### Network programming

* Bluetooth
* CoAP
* Docker
  * [fussybeaver/bollard](https://github.com/fussybeaver/bollard) - Docker daemon API Stars:`1.0K`.
* FTP
* gRPC
  * [hyperium/tonic](https://github.com/hyperium/tonic) - A native gRPC client & server implementation with async/await support [![Crates.io](https://img.shields.io/crates/v/tonic)](https://crates.io/crates/tonic) Stars:`10.7K`.
  * [tikv/grpc-rs](https://github.com/tikv/grpc-rs) - The gRPC library built on C Core library and futures Stars:`1.8K`.
* HTTP
  * [Hurl](https://github.com/Orange-OpenSource/hurl) - Run and test HTTP requests with plain text and libcurl [![CI](https://github.com/Orange-OpenSource/hurl/workflows/CI/badge.svg)](https://github.com/Orange-OpenSource/hurl/actions) Stars:`14.4K`.
* IPNetwork
* Low level
  * [actix/actix](https://github.com/actix/actix) - Actor library Stars:`8.8K`.
  * [smoltcp-rs/smoltcp](https://github.com/smoltcp-rs/smoltcp) - A standalone, event-driven TCP/IP stack that is designed for bare-metal, real-time systems Stars:`4.0K`.
  * [libpnet/libpnet](https://github.com/libpnet/libpnet) - A cross-platform, low level networking Stars:`2.4K`.
* message-io
  * [lemunozm/message-io](https://github.com/lemunozm/message-io) - Event-driven message library to build network applications easy and fast. Supports TCP, UDP and WebSockets. [![build badge](https://img.shields.io/github/workflow/status/lemunozm/message-io/message-io%20ci)](https://github.com/lemunozm/message-io/actions?query=workflow%3A%22message-io+ci%22) Stars:`1.2K`.
* MQTT
  * [bytebeamio/rumqtt](https://github.com/bytebeamio/rumqtt) - A library for developers to build applications that communicate with the [MQTT protocol](https://mqtt.org) over TCP and WebSockets, with or without TLS. [![Build and Test](https://github.com/bytebeamio/rumqtt/actions/workflows/build.yml/badge.svg)](https://github.com/bytebeamio/rumqtt/actions/workflows/build.yml) Stars:`1.8K`.
* NanoMsg
* NATS
  * [nats-io/nats.rs](https://github.com/nats-io/nats.rs) - Client for NATS, the cloud native messaging system. [![Build Status](https://github.com/nats-io/nats.rs/workflows/Rust/badge.svg?branch=master)](https://github.com/nats-io/nats.rs/actions) Stars:`1.2K`.
* Nng
* NNTP
* P2P
  * [libp2p/rust-libp2p](https://github.com/libp2p/rust-libp2p) - Implementation of libp2p networking stack. [![Circle CI](https://circleci.com/gh/libp2p/rust-libp2p.svg?style=svg)](https://app.circleci.com/pipelines/github/libp2p/rust-libp2p) Stars:`4.9K`.
  * [n0-computer/iroh](https://github.com/n0-computer/iroh) [[iroh](https://crates.io/crates/iroh)] - crate for building on direct connections between devices [![CI](https://github.com/n0-computer/iroh/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/n0-computer/iroh/actions/workflows/ci.yml) Stars:`4.4K`.
* POP3
* QUIC
  * [cloudflare/quiche](https://github.com/cloudflare/quiche) - cloudflare implementation of the QUIC transport protocol and HTTP/3 ![build](https://img.shields.io/github/actions/workflow/status/cloudflare/quiche/stable.yml?branch=master) Stars:`10.0K`.
  * [quinn-rs/quinn](https://github.com/quinn-rs/quinn) - Futures-based QUIC implementation [![build badge](https://dev.azure.com/dochtman/Projects/_apis/build/status/Quinn?branchName=master)](https://dev.azure.com/dochtman/Projects/_build) Stars:`4.2K`.
  * [mozilla/neqo](https://github.com/mozilla/neqo) - an Implementation of QUIC Stars:`1.9K`.
  * [tencent/tquic](https://github.com/Tencent/tquic) - A high-performance, lightweight, and cross-platform QUIC library [![Build Status](https://img.shields.io/github/actions/workflow/status/tencent/tquic/rust.yml)](https://github.com/Tencent/tquic/actions/workflows/rust.yml) Stars:`1.2K`.
  * [aws/s2n-quic](https://github.com/aws/s2n-quic) - An implementation of the IETF QUIC protocol ![ci](https://img.shields.io/github/actions/workflow/status/aws/s2n-quic/ci.yml?branch=main) Stars:`1.2K`.
* Raknet
* RPC
* Socket.io
* SSH
* Stomp
* VPN
* Zenoh
  * [eclipse-zenoh/zenoh](https://github.com/eclipse-zenoh/zenoh) - Zero Overhead Network Protocol Stars:`1.7K`.
* ZeroMQ

### Parsing

  * [tree-sitter/tree-sitter](https://github.com/tree-sitter/tree-sitter) - A parser generator tool and an incremental parsing library geared towards programming tools Stars:`20.1K`.
  * [rust-bakery/nom](https://github.com/rust-bakery/nom) - parser combinator library Stars:`9.8K`.
  * [pest-parser/pest](https://github.com/pest-parser/pest) - The Elegant Parser Stars:`4.9K`.
  * [lalrpop/lalrpop](https://github.com/lalrpop/lalrpop) - LR(1) parser generator Stars:`3.2K`.
  * [kevinmehall/rust-peg](https://github.com/kevinmehall/rust-peg) - Parsing Expression Grammar (PEG) parser generator Stars:`1.5K`.
  * [Marwes/combine](https://github.com/Marwes/combine) - parser combinator library Stars:`1.3K`.

### Peripherals

* Fingerprint reader
* Serial Port

### Platform specific

* Cross-platform
* FreeBSD
* Linux
* Unix-like
  * [nix-rust/nix](https://github.com/nix-rust/nix) - Unix-like API bindings [![Cirrus Build Status](https://api.cirrus-ci.com/github/nix-rust/nix.svg)](https://cirrus-ci.com/github/nix-rust/nix) Stars:`2.8K`.
  * [rustix](https://github.com/bytecodealliance/rustix) - Safe bindings to POSIX/Unix/Linux/Winsock2 syscalls [![Actions Status](https://github.com/bytecodealliance/rustix/workflows/CI/badge.svg)](https://github.com/bytecodealliance/rustix/actions?query=workflow%3ACI) Stars:`1.6K`.
  * [zargony/fuse-rs](https://github.com/zargony/fuse-rs) - [FUSE](https://github.com/libfuse/libfuse) bindings Stars:`1.1K`.
* Windows
  * [microsoft/windows-rs](https://github.com/microsoft/windows-rs) - Rust for Windows [![Actions Status](https://github.com/microsoft/windows-rs/workflows/CI/badge.svg)](https://github.com/microsoft/windows-rs/actions) Stars:`11.1K`.
  * [retep998/winapi-rs](https://github.com/retep998/winapi-rs) - Windows API bindings [![Rust](https://github.com/retep998/winapi-rs/actions/workflows/rust.yml/badge.svg?branch=dev)](https://github.com/retep998/winapi-rs/actions/workflows/rust.yml) Stars:`1.9K`.

### Reverse engineering


### Scripting

[[scripting](https://crates.io/keywords/scripting)]

* [rhaiscript/rhai](https://github.com/rhaiscript/rhai) - A tiny and fast embedded scripting language resembling a combination of JavaScript and Rust [![build badge](https://github.com/rhaiscript/rhai/workflows/Build/badge.svg)](https://github.com/rhaiscript/rhai/actions) Stars:`4.5K`.
* [gluon-lang/gluon](https://github.com/gluon-lang/gluon) - A small, statically-typed, functional programming language Stars:`3.3K`.
* [mun](https://github.com/mun-lang/mun) - A compiled, statically-typed scripting language with first class hot reloading support Stars:`2.0K`.
* [kcl](https://github.com/kcl-lang/kcl) - A constraint-based record & functional language mainly used in configuration and policy scenarios. Stars:`1.9K`.
* [rune-rs/rune](https://github.com/rune-rs/rune) - An embeddable dynamic programming language Stars:`1.9K`.
* [PistonDevelopers/dyon](https://github.com/PistonDevelopers/dyon) - A rusty dynamically typed scripting language Stars:`1.8K`.
* [metacall/core](https://github.com/metacall/core) [[metacall](https://crates.io/crates/metacall)] - Cross-platform Polyglot Runtime which supports NodeJS, JavaScript, TypeScript, Python, Ruby, C#, Wasm, Java, Cobol and more. [![build badge](https://gitlab.com/metacall/core/badges/master/pipeline.svg)](https://gitlab.com/metacall/core) Stars:`1.6K`.
* [trynova/nova](https://github.com/trynova/nova) - JavaScript engine written entirely in Rust Stars:`1.2K`.

### Simulation

[[simulation](https://crates.io/keywords/simulation)]


### Social networks

* Telegram

### System

* [GuillaumeGomez/sysinfo](https://github.com/GuillaumeGomez/sysinfo) [[sysinfo](https://crates.io/crates/sysinfo)] - Cross-platform library to fetch system information [![build badge](https://github.com/GuillaumeGomez/sysinfo/actions/workflows/CI.yml/badge.svg?branch=master)](https://github.com/GuillaumeGomez/sysinfo/actions/workflows/CI.yml) Stars:`2.4K`.

### Task scheduling


### Template engine

* Handlebars
  * [sunng87/handlebars-rust](https://github.com/sunng87/handlebars-rust) - Handlebars template engine with inheritance, custom helper support. Stars:`1.4K`.
* HTML
  * [Keats/tera](https://github.com/Keats/tera) - template engine based on Jinja2 and the Django template language. [![Actions Status](https://github.com/Keats/tera/workflows/ci/badge.svg?branch=master)](https://github.com/Keats/tera/actions) Stars:`3.8K`.
  * [lambda-fairy/maud](https://github.com/lambda-fairy/maud) - compile-time HTML templates Stars:`2.3K`.
* Mustache

### Text processing

* [rust-lang/regex](https://github.com/rust-lang/regex) - Regular expressions (RE2 style) Stars:`3.7K`.
* [greyblake/whatlang-rs](https://github.com/greyblake/whatlang-rs) - Natural language detection library based on trigrams Stars:`1.0K`.

### Text search

* [meilisearch/MeiliSearch](https://github.com/meilisearch/MeiliSearch) - Ultra relevant, instant and typo-tolerant full-text search API. [![Build Status](https://github.com/meilisearch/MeiliSearch/workflows/Cargo%20test/badge.svg?branch=master)](https://github.com/meilisearch/MeiliSearch/actions) Stars:`50.1K`.
* [tantivy](https://github.com/quickwit-oss/tantivy) [[tantivy](https://crates.io/crates/tantivy)] - A horse-speed full-text search engine library written in Rust. [![Build Status](https://github.com/quickwit-oss/tantivy/actions/workflows/test.yml/badge.svg)](https://github.com/quickwit-oss/tantivy/actions/workflows/test.yml) Stars:`12.9K`.
* [pg_search](https://github.com/paradedb/paradedb/tree/dev/pg_search) - PostgreSQL extension that enables full-text search over SQL tables using the BM25 algorithm, the state-of-the-art ranking function for full-text search. Stars:`6.9K`.
* [BurntSushi/fst](https://github.com/BurntSushi/fst) [[fst](https://crates.io/crates/fst)] - a fast implementation of ordered sets and maps using finite state machines Stars:`1.9K`.
* [SeekStorm](https://github.com/SeekStorm/SeekStorm) [[SeekStorm](https://crates.io/crates/seekstorm)] - sub-millisecond full-text search library & multi-tenancy server in Rust Stars:`1.7K`.

### Unsafe


### Video


### Virtualization

* [bytecodealliance/wasmtime](https://github.com/bytecodealliance/wasmtime) - A standalone runtime for WebAssembly [![Build Status](https://github.com/bytecodealliance/wasmtime/workflows/CI/badge.svg)](https://github.com/bytecodealliance/wasmtime/actions?query=workflow%3ACI) Stars:`16.1K`.

### Web programming

See also [Are we web yet?](https://www.arewewebyet.org) and [Rust web framework comparison](https://github.com/flosse/rust-web-framework-comparison).

* Client-side / WASM
  * [leptos](https://github.com/leptos-rs/leptos) - Leptos is a full-stack, isomorphic web framework leveraging fine-grained reactivity to build declarative user interfaces.[![crate](https://img.shields.io/crates/v/create-rust-app.svg)](https://crates.io/crates/leptos) Stars:`18.1K`.
  * [seed](https://github.com/seed-rs/seed) - A framework for creating web apps Stars:`3.8K`.
  * [sauron](https://github.com/ivanceras/sauron) - Client side web framework which closely adheres to The Elm Architecture. Stars:`2.0K`.
* HTTP Client
  * [hyperium/hyper](https://github.com/hyperium/hyper) - an HTTP implementation [![CI](https://github.com/hyperium/hyper/workflows/CI/badge.svg?branch=master)](https://github.com/hyperium/hyper/actions?query=workflow%3ACI) Stars:`15.1K`.
  * [seanmonstar/reqwest](https://github.com/seanmonstar/reqwest) - an ergonomic HTTP Client. Stars:`10.5K`.
  * [ducaale/xh](https://github.com/ducaale/xh) - Friendly and fast tool for sending HTTP requests [![crate](https://img.shields.io/crates/v/create-rust-app.svg)](https://crates.io/crates/xh) [![GitHub actions Status](https://github.com/ducaale/xh/workflows/CI/badge.svg?branch=master)](https://github.com/ducaale/xh/actions) Stars:`6.2K`.
  * [async-graphql](https://github.com/async-graphql/async-graphql) - A GraphQL server library [![Build Status](https://dev.azure.com/graphql-rust/GraphQL%20Rust/_apis/build/status/graphql-rust.juniper)](https://dev.azure.com/graphql-rust/GraphQL%20Rust/_build/latest?definitionId=1) Stars:`3.5K`.
  * [graphql-client](https://github.com/graphql-rust/graphql-client) - Typed, correct GraphQL requests and responses. [![GitHub actions Status](https://github.com/graphql-rust/graphql-client/workflows/CI/badge.svg?branch=master)](https://github.com/graphql-rust/graphql-client/actions) Stars:`1.2K`.
  * [alexcrichton/curl-rust](https://github.com/alexcrichton/curl-rust) - [libcurl](https://curl.se/libcurl/) bindings Stars:`1.0K`.
* HTTP Server
  * [Rocket](https://github.com/rwf2/Rocket) - Rocket is a web framework with a focus on ease-of-use, expressability, and speed Stars:`25.0K`.
  * [actix/actix-web](https://github.com/actix/actix-web) - A lightweight async web framework with websocket support Stars:`22.7K`.
  * [tokio/axum](https://github.com/tokio-rs/axum) - Ergonomic and modular web framework built with Tokio, Tower, and Hyper [![Build badge](https://github.com/tokio-rs/axum/actions/workflows/CI.yml/badge.svg?branch=main)](https://github.com/tokio-rs/axum/actions/workflows/CI.yml) Stars:`21.1K`.
  * [hyperium/hyper](https://github.com/hyperium/hyper) - an HTTP implementation [![CI](https://github.com/hyperium/hyper/workflows/CI/badge.svg?branch=master)](https://github.com/hyperium/hyper/actions?query=workflow%3ACI) Stars:`15.1K`.
  * [seanmonstar/warp](https://github.com/seanmonstar/warp) - A super-easy, composable, web server framework for warp speeds. [![crate](https://img.shields.io/crates/v/create-rust-app.svg)](https://crates.io/crates/warp) Stars:`9.8K`.
  * [Iron](https://github.com/iron/iron) - A middleware-based server framework Stars:`6.1K`.
  * [Juniper](https://github.com/graphql-rust/juniper) - GraphQL server library Stars:`5.8K`.
  * [poem-web/poem](https://github.com/poem-web/poem) - A full-featured and easy-to-use web framework. [![CI](https://github.com/poem-web/poem/actions/workflows/ci.yml/badge.svg)](https://github.com/poem-web/poem/actions/workflows/ci.yml) Stars:`3.9K`.
  * [Salvo](https://github.com/salvo-rs/salvo) - an easy to use webframework base on hyper and tokio. [![build build](https://github.com/salvo-rs/salvo/actions/workflows/release.yml/badge.svg)](https://github.com/salvo-rs/salvo/actions) Stars:`3.6K`.
  * [Nickel](https://github.com/nickel-org/nickel.rs/) - inspired by [Express](https://expressjs.com/) Stars:`3.0K`.
  * [Gotham](https://github.com/gotham-rs/gotham) - A flexible web framework that does not sacrifice safety, security or speed. Stars:`2.3K`.
  * [handlebars-rust](https://github.com/sunng87/handlebars-rust) - an Iron web framework middleware. Stars:`1.4K`.
  * [tomaka/rouille](https://github.com/tomaka/rouille) - Web framework Stars:`1.2K`.
  * [tiny-http](https://github.com/tiny-http/tiny-http) - Low level HTTP server library Stars:`1.1K`.
* Miscellaneous
  * [serenity-rs/serenity](https://github.com/serenity-rs/serenity) [[serenity](https://crates.io/crates/serenity)] - A library for the Discord API Stars:`5.1K`.
  * [teloxide/teloxide](https://github.com/teloxide/teloxide/) - An elegant Telegram bots framework [![Build Status](https://github.com/teloxide/teloxide/actions/workflows/ci.yml/badge.svg)](https://github.com/teloxide/teloxide/actions) Stars:`3.5K`.
  * [osohq/oso](https://github.com/osohq/oso) [[oso](https://crates.io/crates/oso)] - A policy engine for authorization that's embedded in your application. [![Build Status](https://github.com/osohq/oso/workflows/Development/badge.svg?branch=main)](https://github.com/osohq/oso/actions?query=branch%3Amain+workflow%3ADevelopment) Stars:`3.5K`.
  * [Utoipa](https://github.com/juhaku/utoipa) - Simple, Fast, Code first and Compile time generated OpenAPI documentation [![crates.io](https://img.shields.io/crates/v/utoipa.svg?label=crates.io&color=orange&logo=rust)](https://crates.io/crates/utoipa) [![Utoipa build](https://github.com/juhaku/utoipa/actions/workflows/build.yaml/badge.svg)](https://github.com/juhaku/utoipa/actions/workflows/build.yaml) Stars:`3.0K`.
  * [svix/svix-webhooks](https://github.com/svix/svix-webhooks) [[svix](https://crates.io/crates/svix)] - A library for sending webhooks and verifying signatures. Stars:`2.6K`.
  * [rust-scraper/scraper](https://github.com/rust-scraper/scraper) [[scraper](https://crates.io/crates/scraper)] - HTML parsing and querying with CSS selectors. [![Build Status](https://github.com/rust-scraper/scraper/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/rust-scraper/scraper/actions) Stars:`2.1K`.
  * [pyrossh/rust-embed](https://github.com/pyrossh/rust-embed) - A macro to embed static assets into the rust binary Stars:`1.8K`.
* Reverse Proxy
  * [sozu-proxy/sozu](https://github.com/sozu-proxy/sozu) [[sozu](https://crates.io/crates/sozu)] - A HTTP reverse proxy. [![CI](https://github.com/sozu-proxy/sozu/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/sozu-proxy/sozu/actions/workflows/ci.yml) Stars:`3.3K`.
* Static Site Generators
  * [getzola/zola](https://github.com/getzola/zola) [[zola](https://www.getzola.org/)] - An opinionated static site generator with everything built-in. [![Build Status](https://dev.azure.com/getzola/zola/_apis/build/status/getzola.zola?branchName=master)](https://dev.azure.com/getzola/zola/_build) Stars:`15.0K`.
  * [vi/websocat](https://github.com/vi/websocat) - CLI for interacting with WebSockets, with functionality of Netcat, Curl and Socat. Stars:`7.5K`.
  * [snapview/tungstenite-rs](https://github.com/snapview/tungstenite-rs) - Lightweight stream-based WebSocket implementation. Stars:`2.1K`.
  * [rust-websocket](https://github.com/websockets-rs/rust-websocket) - A framework for dealing with WebSocket connections (both clients and servers) Stars:`1.6K`.
  * [housleyjk/ws-rs](https://github.com/housleyjk/ws-rs) - lightweight, event-driven WebSockets Stars:`1.5K`.
  * [cobalt-org/cobalt.rs](https://github.com/cobalt-org/cobalt.rs) - Static site generator [![Build Status](https://dev.azure.com/cobalt-org/cobalt-org/_apis/build/status/cobalt.rs?branchName=master)](https://dev.azure.com/cobalt-org/cobalt-org/_build?definitionId=2) Stars:`1.5K`.

## Registries

A registry allows you to publish your Rust libraries as crate packages, to share them with others publicly and privately.


## Resources

* Benchmarks
* Decks & Presentations
* Learning
  * [Rustlings](https://github.com/rust-lang/rustlings) - small exercises to get you used to reading and writing Rust code Stars:`57.4K`.
  * [rust-learning](https://github.com/ctjhoa/rust-learning) - A collection of useful resources to learn Rust Stars:`11.8K`.
  * [Easy Rust](https://github.com/Dhghomon/easy_rust) - Learn Rust in easy English. Stars:`8.2K`.
  * [Idiomatic Rust](https://github.com/mre/idiomatic-rust) - A peer-reviewed collection of articles/talks/repos which teach idiomatic Rust. Stars:`7.0K`.
  * [Aquascope](https://github.com/cognitive-engineering-lab/aquascope) - Interactive visualizations of Rust at compile-time and run-time Stars:`2.5K`.
  * [stdx](https://github.com/brson/stdx) - Learn these crates first as an extension to std Stars:`2.0K`.
  * [rust-how-do-i-start](https://github.com/jondot/rust-how-do-i-start) - A repo dedicated to answering the question: "So, Rust. How do I *start*?". A beginner only hand-picked resources and learning track. Stars:`1.1K`.
* Podcasts
* [Rust Design Patterns](https://github.com/rust-unofficial/patterns) - A catalogue of Rust design patterns, anti-patterns and idioms Stars:`8.3K`.
* [RustBooks](https://github.com/sger/RustBooks) - list of RustBooks Stars:`4.9K`.
* [RustViz](https://github.com/rustviz/rustviz) - generates visualizations from simple Rust programs to assist users in better understanding the Rust Lifetime and Borrowing mechanism. Stars:`2.8K`.

## License

[![CC0](https://licensebuttons.net/p/zero/1.0/88x31.png)](https://creativecommons.org/publicdomain/zero/1.0/)
