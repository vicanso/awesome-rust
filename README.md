# Awesome Rust [![build badge](https://github.com/rust-unofficial/awesome-rust/actions/workflows/rust.yml/badge.svg?branch=main)](https://github.com/rust-unofficial/awesome-rust/actions/workflows/rust.yml) [![Track Awesome List](https://www.trackawesomelist.com/badge.svg)](https://www.trackawesomelist.com/rust-unofficial/awesome-rust/)

A curated list of Rust code and resources.

If you want to contribute, please read [this](CONTRIBUTING.md).

## Table of contents

<!-- toc -->

- [Applications](#applications)
  * [Audio and Music](#audio-and-music)
  * [Cryptocurrencies](#cryptocurrencies)
  * [Database](#database)
  * [Emulators](#emulators)
  * [Games](#games)
  * [Graphics](#graphics)
  * [Image processing](#image-processing)
  * [Industrial automation](#industrial-automation)
  * [Observability](#observability)
  * [Operating systems](#operating-systems)
  * [Productivity](#productivity)
  * [Security tools](#security-tools)
  * [Simulation](#simulation)
  * [System tools](#system-tools)
  * [Task scheduling](#task-scheduling)
  * [Text editors](#text-editors)
  * [Text processing](#text-processing)
  * [Utilities](#utilities)
  * [Video](#video)
  * [Virtualization](#virtualization)
  * [Web](#web)
  * [Web Servers](#web-servers)
- [Development tools](#development-tools)
  * [Build system](#build-system)
  * [Debugging](#debugging)
  * [Deployment](#deployment)
  * [Embedded](#embedded)
  * [FFI](#ffi)
  * [Formatters](#formatters)
  * [IDEs](#ides)
  * [Profiling](#profiling)
  * [Services](#services)
  * [Static analysis](#static-analysis)
  * [Testing](#testing)
  * [Transpiling](#transpiling)
- [Libraries](#libraries)
  * [Artificial Intelligence](#artificial-intelligence)
    + [Genetic algorithms](#genetic-algorithms)
    + [Machine learning](#machine-learning)
  * [Astronomy](#astronomy)
  * [Asynchronous](#asynchronous)
  * [Audio and Music](#audio-and-music-1)
  * [Authentication](#authentication)
  * [Automotive](#automotive)
  * [Bioinformatics](#bioinformatics)
  * [Caching](#caching)
  * [Cloud](#cloud)
  * [Command-line](#command-line)
  * [Compression](#compression)
  * [Computation](#computation)
  * [Concurrency](#concurrency)
  * [Configuration](#configuration)
  * [Cryptography](#cryptography)
  * [Data processing](#data-processing)
  * [Data streaming](#data-streaming)
  * [Data structures](#data-structures)
  * [Data visualization](#data-visualization)
  * [Database](#database-1)
  * [Date and time](#date-and-time)
  * [Distributed systems](#distributed-systems)
  * [Domain driven design](#domain-driven-design)
  * [Email](#email)
  * [Encoding](#encoding)
  * [Filesystem](#filesystem)
  * [Functional Programming](#functional-programming)
  * [Game development](#game-development)
  * [Geospatial](#geospatial)
  * [Graphics](#graphics-1)
  * [GUI](#gui)
  * [Image processing](#image-processing-1)
  * [Language specification](#language-specification)
  * [Logging](#logging)
  * [Macro](#macro)
  * [Markup language](#markup-language)
  * [Mobile](#mobile)
  * [Network programming](#network-programming)
  * [Parsing](#parsing)
  * [Peripherals](#peripherals)
  * [Platform specific](#platform-specific)
  * [Scripting](#scripting)
  * [Simulation](#simulation-1)
  * [Task scheduling](#task-scheduling-1)
  * [Template engine](#template-engine)
  * [Text processing](#text-processing-1)
  * [Text search](#text-search)
  * [Unsafe](#unsafe)
  * [Virtualization](#virtualization-1)
  * [Web programming](#web-programming)
- [Registries](#registries)
- [Resources](#resources)
- [License](#license)

<!-- tocstop -->

## Applications

See also [Rust — Production](https://www.rust-lang.org/production) organizations running Rust in production.

* [denoland/deno](https://github.com/denoland/deno) — A secure JavaScript/TypeScript runtime built with V8, Rust, and Tokio [![Build Status](https://github.com/denoland/deno/workflows/ci/badge.svg?branch=master&event=push)](https://github.com/denoland/deno/actions) Stars:`86.2K`.
* [alacritty](https://github.com/alacritty/alacritty) — A cross-platform, GPU enhanced terminal emulator Stars:`42.8K`.
* [SWC](https://github.com/swc-project/swc) — super-fast TypeScript / JavaScript compiler Stars:`24.7K`.
* [Servo](https://github.com/servo/servo) — A prototype web browser engine Stars:`21.8K`.
* [wasmer](https://github.com/wasmerio/wasmer) — A safe and fast WebAssembly runtime supporting WASI and Emscripten [![Build Status](https://github.com/wasmerio/wasmer/workflows/build/badge.svg?style=flat-square)](https://github.com/wasmerio/wasmer/actions) Stars:`13.3K`.
* [zellij](https://github.com/zellij-org/zellij) — A terminal multiplexer (workspace) with batteries included Stars:`8.8K`.
* [wezterm](https://github.com/wez/wezterm) — A GPU-accelerated cross-platform terminal emulator and multiplexer Stars:`6.2K`.
* [cloudflare/boringtun](https://github.com/cloudflare/boringtun) — A Userspace WireGuard VPN Implementation [![build badge](https://img.shields.io/badge/crates.io-v0.2.0-orange.svg)](https://crates.io/crates/boringtun) Stars:`4.7K`.
* [datafusion](https://github.com/apache/arrow-datafusion) — Apache Arrow DataFusion and Ballista query engines Stars:`2.7K`.
* [rx](https://github.com/cloudhead/rx) — Vi inspired Modern Pixel Art Editor Stars:`2.5K`.
* [habitat](https://github.com/habitat-sh/habitat) — A tool created by Chef to build, deploy, and manage applications. Stars:`2.4K`.
* [notty](https://github.com/withoutboats/notty) — A new kind of terminal Stars:`2.2K`.
* [linkerd/linkerd2-proxy](https://github.com/linkerd/linkerd2-proxy) — Ultralight service mesh for Kubernetes. Stars:`1.6K`.
* [fcsonline/drill](https://github.com/fcsonline/drill) — A HTTP load testing application inspired by Ansible syntax [![build badge](https://api.travis-ci.org/fcsonline/drill.svg?branch=master)](https://travis-ci.org/fcsonline/drill) Stars:`1.4K`.
* [shuttle](https://github.com/shuttle-hq/shuttle) — A serverless platform built for Rust Stars:`1.3K`.
* [kalker](https://github.com/PaddiM8/kalker) - A scientific calculator that supports math-like syntax with user-defined variables, functions, derivation, integration, and complex numbers. Cross platform + WASM support [![Build Status](https://github.com/PaddiM8/kalker/workflows/Release/badge.svg)](https://github.com/PaddiM8/kalker/actions) Stars:`1.2K`.
* [tiny](https://github.com/osa1/tiny) — A terminal IRC client Stars:`791`.
* [Fractalide](https://github.com/fractalide/fractalide) — Simple Rust Microservices Stars:`778`.
* [jedisct1/flowgger](https://github.com/awslabs/flowgger) — A fast, simple and lightweight data collector Stars:`722`.
* [kytan](https://github.com/changlan/kytan) — High Performance Peer-to-Peer VPN Stars:`356`.
* [Weld](https://github.com/serayuzgur/weld) — Full fake REST API generator [![build badge](https://api.travis-ci.org/serayuzgur/weld.svg?branch=master)](https://travis-ci.org/serayuzgur/weld) Stars:`234`.
* [asm-cli-rust](https://github.com/cch123/asm-cli-rust) — An interactive assembly shell written in rust. Stars:`216`.
* [Factotum](https://github.com/snowplow/factotum) — [A system to programmatically run data pipelines](https://snowplow.io/blog/introducing-factotum-data-pipeline-runner/) [![build badge](https://api.travis-ci.org/snowplow/factotum.svg?branch=master)](https://travis-ci.org/snowplow/factotum) Stars:`194`.
* [nicohman/eidolon](https://github.com/nicohman/eidolon) — A steam and drm-free game registry and launcher for linux and macosx [![build badge](https://api.travis-ci.org/nicohman/eidolon.svg?branch=master)](https://travis-ci.org/nicohman/eidolon) Stars:`110`.
* [Herd](https://github.com/imjacobclark/Herd) — an experimental HTTP load testing application Stars:`96`.
* [fend](https://github.com/printfn/fend) - Arbitrary-precision unit-aware calculator [![build](https://github.com/printfn/fend/workflows/build/badge.svg)](https://github.com/printfn/fend) Stars:`94`.
* [trust-dns](https://crates.io/crates/trust-dns) — A DNS-server [![Build Status](https://github.com/bluejekyll/trust-dns/workflows/test/badge.svg?branch=main)](https://github.com/bluejekyll/trust-dns/actions?query=workflow%3Atest)
* [Pijul](https://pijul.org) — A patch-based distributed version control system
* [mdBook](https://crates.io/crates/mdbook) — A command line utility to create books from markdown files [![Build Status](https://github.com/rust-lang/mdBook/workflows/CI/badge.svg?branch=master)](https://github.com/rust-lang/mdBook/actions)
* [MaidSafe](https://github.com/maidsafe) — A decentralized platform.

### Audio and Music

* [Spotify TUI](https://github.com/Rigellute/spotify-tui) — A Spotify client for the terminal written in Rust. ![Continuous Integration](https://github.com/Rigellute/spotify-tui/workflows/Continuous%20Integration/badge.svg?branch=master) Stars:`13.6K`.
* [Spotifyd](https://github.com/Spotifyd/spotifyd) — An open source Spotify client running as a UNIX daemon. ![Continuous Integration](https://github.com/Spotifyd/spotifyd/workflows/Continuous%20Integration/badge.svg?branch=master) Stars:`6.5K`.
* [ncspot](https://github.com/hrkfdn/ncspot) - Cross-platform ncurses Spotify client, inspired by ncmpc and the likes. [![build badge](https://github.com/hrkfdn/ncspot/workflows/Build/badge.svg)](https://github.com/hrkfdn/ncspot/actions?query=workflow%3ABuild) Stars:`3.3K`.
* [Glicol](https://github.com/chaosprint/glicol) — Graph-oriented live coding language written in Rust for collaborative musicking in browsers. Stars:`1.1K`.
* [Polaris](https://github.com/agersant/polaris) — A music streaming application.  [![build badge](https://api.travis-ci.org/agersant/polaris.svg?branch=master)](https://travis-ci.org/agersant/polaris) Stars:`985`.
* [enginesound](https://github.com/DasEtwas/enginesound) — A GUI and command line application used to procedurally generate semi-realistic engine sounds. Featuring in-depth configuration, variable sample rate and a frequency analysis window. Stars:`187`.

### Cryptocurrencies

* [Diem](https://github.com/diem/diem) — Diem’s mission is to enable a simple global currency and financial infrastructure that empowers billions of people. Stars:`16.7K`.
* [Solana](https://github.com/solana-labs/solana) — Incredibly fast, highly scalable blockchain using Proof-of-History. Stars:`9.6K`.
* [Substrate](https://github.com/paritytech/substrate) — Generic modular blockchain template written in Rust Stars:`7.6K`.
* [Polkadot](https://github.com/paritytech/polkadot) — Heterogeneous multi‑chain technology with pooled security Stars:`6.4K`.
* [zcash](https://github.com/zcash/zcash) — Zcash is an implementation of the "Zerocash" protocol. Stars:`4.7K`.
* [Foundry](https://github.com/foundry-rs/foundry) - Foundry is a blazing fast, portable and modular toolkit for Ethereum application development written in Rust. ![Build Status](https://img.shields.io/github/workflow/status/foundry-rs/foundry/test?style=flat-square) Stars:`4.6K`.
* [Lighthouse](https://github.com/sigp/lighthouse) — Rust Ethereum Consensus Layer (CL) Client [![Build Status](https://github.com/sigp/lighthouse/workflows/test-suite/badge.svg?branch=master)](https://github.com/sigp/lighthouse/actions) Stars:`2.1K`.
* [near/nearcore](https://github.com/near/nearcore) — decentralized smart-contract platform for low-end mobile devices. Stars:`1.9K`.
* [ethers-rs](https://github.com/gakonst/ethers-rs) - Complete Ethereum & Celo library and wallet implementation in Rust. ![Build Status](https://github.com/gakonst/ethers-rs/workflows/Tests/badge.svg) Stars:`1.3K`.
* [CITA](https://github.com/citahub/cita) — A high performance blockchain kernel for enterprise users. Stars:`1.3K`.
* [rust-bitcoin](https://github.com/rust-bitcoin/rust-bitcoin) — Library with support for de/serialization, parsing and executing on data structures and network messages related to Bitcoin. Stars:`1.2K`.
* [Nervos CKB](https://github.com/nervosnetwork/ckb) — Nervos CKB is a public permissionless blockchain, the common knowledge layer of Nervos network. Stars:`1.0K`.
* [rust-lightning](https://github.com/lightningdevkit/rust-lightning) [![Crate](https://img.shields.io/crates/v/lightning.svg?logo=rust)](https://crates.io/crates/lightning) — Bitcoin Lightning library written in Rust. The main crate,`lightning`, does not handle networking, persistence, or any other I/O. Thus,it is runtime-agnostic, but users must implement basic networking logic, chain interactions, and disk storage.po on linking crate. Stars:`814`.
* [Holochain](https://github.com/holochain/holochain) — Scalable P2P alternative to blockchain for all those distributed apps you always wanted to build. [![detect critical check failures](https://github.com/holochain/holochain/actions/workflows/check_run_detect_release_pr_failure.yml/badge.svg)](https://github.com/holochain/holochain/actions/workflows/check_run_detect_release_pr_failure.yml) Stars:`723`.
* [Akula](https://github.com/akula-bft/akula) - Rust Ethereum Execution Layer (EL) Client (WIP) Stars:`721`.
* [Parity-Bitcoin](https://github.com/paritytech/parity-bitcoin) — The Parity Bitcoin client [![build badge](https://api.travis-ci.org/paritytech/parity-bitcoin.svg?branch=master)](https://app.travis-ci.com/github/paritytech/parity-bitcoin) Stars:`714`.
* [electrumrs](https://github.com/romanz/electrs) — An efficient re-implementation of Electrum Server in Rust. Stars:`681`.
* [wagyu](https://github.com/AleoHQ/wagyu) [[wagyu](https://crates.io/crates/wagyu)] — Rust library for generating cryptocurrency wallets [![build badge](https://api.travis-ci.com/AleoHQ/wagyu.svg?branch=master)](https://api.travis-ci.com/AleoHQ/wagyu.svg?branch=master) Stars:`526`.
* [Forest](https://github.com/ChainSafe/forest) - Rust Filecoin implementation [![Build Status](https://img.shields.io/circleci/build/gh/ChainSafe/forest/main?branch=master)](https://app.circleci.com/pipelines/github/ChainSafe/forest?branch=main) Stars:`487`.
* [tendermint-rs](https://github.com/informalsystems/tendermint-rs) - Rust implementation of Tendermint blockchain data structures and clients Stars:`427`.
* [revm](https://github.com/bluealloy/revm) - Revolutionary Machine (revm) is a fast Ethereum virtual machine written in rust. Stars:`420`.
* [mev-inspect-rs](https://github.com/flashbots/mev-inspect-rs) - Ethereum MEV Inspector in Rust Stars:`407`.
* [ethabi](https://github.com/rust-ethereum/ethabi) - Encode and decode smart contract invocations. Stars:`346`.
* [Phala-Network/phala-blockchain](https://github.com/Phala-Network/phala-blockchain) — Confidential smart contract blockchain based on Intel SGX and Substrate Stars:`305`.
* [ibc-rs](https://github.com/informalsystems/hermes) - Rust implementation of the [Interblockchain Communication](https://ibcprotocol.org/) protocol Stars:`280`.
* [ChainX](https://github.com/chainx-org/ChainX) — Fully Decentralized Interchain Crypto Asset Management on Polkadot. Stars:`258`.
* [etk](https://github.com/quilt/etk) - etk is a collection of tools for writing, reading, and analyzing EVM bytecode. Stars:`238`.
* [opensea-rs](https://github.com/gakonst/opensea-rs) - Rust bindings & CLI to the Opensea API and Contracts. Stars:`217`.
* [interBTC](https://github.com/interlay/interbtc) — Trustless and fully decentralized Bitcoin bridge to Polkadot and Kusama. Stars:`179`.
* [Joystream](https://github.com/Joystream/joystream) — A user governed video platform [![Build Status](https://api.travis-ci.org/Joystream/joystream.svg?branch=master)](https://travis-ci.org/Joystream/joystream) Stars:`151`.
* [ethaddrgen](https://github.com/Limeth/ethaddrgen) — Custom Ethereum vanity address generator made in Rust [![build badge](https://api.travis-ci.org/Limeth/ethaddrgen.svg?branch=master)](https://travis-ci.org/Limeth/ethaddrgen) Stars:`150`.
* [coinbase-pro-rs](https://github.com/inv2004/coinbase-pro-rs) — Coinbase pro client in Rust, supports sync/async/websocket [![build badge](https://api.travis-ci.org/inv2004/coinbase-pro-rs.svg?branch=master)](https://travis-ci.org/inv2004/coinbase-pro-rs) Stars:`128`.
* [svm-rs](https://github.com/roynalnaruto/svm-rs) - Solidity-Compiler Version Manager. Stars:`107`.
* [Nimiq](https://github.com/nimiq/core-rs) — Rust implementation of Nimiq node Stars:`72`.
* [Bitcoin Satoshi's Vision](https://github.com/brentongunning/rust-sv) [[sv](https://crates.io/crates/sv)] — A Rust library for working with Bitcoin SV . Stars:`51`.
* [infincia/bip39-rs](https://github.com/infincia/bip39-rs) [[bip39](https://crates.io/crates/bip39)] — Rust implementation of BIP39. Stars:`48`.
* [hdwallet](https://github.com/jjyr/hdwallet) [[hdwallet](https://crates.io/crates/hdwallet)] — BIP-32 HD wallet related key derivation utilities. Stars:`24`.
* [Grin](https://github.com/mimblewimble/grin/) — Evolution of the MimbleWimble protocol

### Database

* [SurrealDB](https://github.com/surrealdb/surrealdb) — A scalable, distributed, document-graph database [![Build Status](https://img.shields.io/github/workflow/status/surrealdb/surrealdb/Continuous%20integration/main)](https://github.com/surrealdb/surrealdb/actions) Stars:`15.1K`.
* [tikv](https://github.com/tikv/tikv) — A distributed KV database in Rust [![Build Status](https://ci.pingcap.net/job/tikv_ghpr_test/badge/icon)](https://ci.pingcap.net/job/tikv_ghpr_test/) Stars:`12.2K`.
* [Databend](https://github.com/datafuselabs/databend) - A Modern Real-Time Data Processing & Analytics DBMS with Cloud-Native Architecture [![Release](https://github.com/datafuselabs/databend/actions/workflows/databend-release.yml/badge.svg)](https://github.com/datafuselabs/databend/actions/workflows/databend-release.yml) Stars:`4.7K`.
* [Materialize](https://github.com/MaterializeInc/materialize) - Streaming SQL database powered by Timely Dataflow :heavy_dollar_sign: [![Build status](https://badge.buildkite.com/97d6604e015bf633d1c2a12d166bb46f3b43a927d3952c999a.svg?branch=main)](https://buildkite.com/materialize/tests) Stars:`4.6K`.
* [noria](https://github.com/mit-pdos/noria) [[noria](https://crates.io/crates/noria)] — Dynamically changing, partially-stateful data-flow for web application backends [![build badge](https://api.travis-ci.org/mit-pdos/noria.svg?branch=master)](https://travis-ci.org/mit-pdos/noria) Stars:`4.4K`.
* [RisingWaveLabs/RisingWave](https://github.com/RisingWaveLabs/risingwave) - the next-generation streaming database in the cloud [![CI](https://github.com/RisingWaveLabs/risingwave/actions/workflows/main.yml/badge.svg)](https://github.com/RisingWaveLabs/risingwave/actions/workflows/main.yml/badge.svg?branch=main) Stars:`3.4K`.
* [Qdrant](https://github.com/qdrant/qdrant) - An open source vector similarity search engine with extended filtering support [![Tests](https://github.com/qdrant/qdrant/workflows/Tests/badge.svg)](https://github.com/qdrant/qdrant/actions) Stars:`2.9K`.
* [seppo0010/rsedis](https://github.com/seppo0010/rsedis) — A Redis reimplementation in Rust [![build badge](https://api.travis-ci.org/seppo0010/rsedis.svg?branch=master)](https://travis-ci.org/seppo0010/rsedis) Stars:`1.6K`.
* [Skytable](https://github.com/skytable/skytable) — A multi-model NoSQL database ![GitHub Workflow Status](https://img.shields.io/github/workflow/status/skytable/skytable/Tests?style=flat-square) Stars:`1.4K`.
* [PumpkinDB](https://github.com/PumpkinDB/PumpkinDB) — an event sourcing database engine Stars:`1.3K`.
* [TerminusDB](https://github.com/terminusdb/terminusdb-store) - open source graph database and document store [![Build Status](https://github.com/terminusdb/terminusdb-store/workflows/Build/badge.svg?branch=master)](https://github.com/terminusdb/terminusdb-store/actions) Stars:`298`.
* [Lucid](https://github.com/lucid-kv/lucid) — High performance and distributed KV store accessible through a HTTP API. [![Build Status](https://github.com/lucid-kv/lucid/workflows/Lucid/badge.svg?branch=master)](https://github.com/lucid-kv/lucid/actions?workflow=Lucid) Stars:`294`.
* [ParityDB](https://github.com/paritytech/parity-db) — Fast and reliable database, optimised for read operation Stars:`161`.
* [WooriDB](https://github.com/naomijub/wooridb) - General purpose time serial database inspired by Crux and Datomic. Stars:`115`.
* [vorot93/libmdbx-rs](https://github.com/vorot93/libmdbx-rs) [[mdbx-sys](https://crates.io/crates/mdbx-sys)] — Rust bindings for MDBX, a "fast, compact, powerful, embedded, transactional key-value database, with permissive license". This is a fork of mozilla/lmdb-rs with patches to make it work with libmdbx. Stars:`36`.
* [sled](https://crates.io/crates/sled) — A (beta) modern embedded database [![Build Status](https://github.com/spacejam/sled/workflows/Rust/badge.svg?branch=master)](https://github.com/spacejam/sled/actions?workflow=Rust)
* [indradb](https://crates.io/crates/indradb) — Rust based graph database [![build badge](https://api.travis-ci.org/indradb/indradb.svg?branch=master)](https://travis-ci.org/indradb/indradb)

### Emulators

See also [crates matching keyword 'emulator'](https://crates.io/keywords/emulator).

* CHIP-8
  * [ColinEberhardt/wasm-rust-chip8](https://github.com/ColinEberhardt/wasm-rust-chip8) — A WebAssembly CHIP-8 emulator written with Rust Stars:`235`.
  * [starrhorne/chip8-rust](https://github.com/starrhorne/chip8-rust) — Yet another rust chip8 emulator Stars:`99`.
* Commodore 64
  * [kondrak/rust64](https://github.com/kondrak/rust64) — [![build badge](https://api.travis-ci.org/kondrak/rust64.svg?branch=master)](https://travis-ci.org/kondrak/rust64) Stars:`211`.
* Flash Player
  * [Ruffle](https://github.com/ruffle-rs/ruffle) — Ruffle is an Adobe Flash Player emulator written in the Rust programming language. Ruffle targets both the desktop and the web using WebAssembly. [![Build Status Rust](https://img.shields.io/github/workflow/status/ruffle-rs/ruffle/Test%20Rust?label=Rust%20Build&logo=github)](https://github.com/ruffle-rs/ruffle/actions?workflow=Test%20Rust)[![Build Status Web](https://img.shields.io/github/workflow/status/ruffle-rs/ruffle/Test%20Web?label=Web%20Build&logo=github)](https://github.com/ruffle-rs/ruffle/actions?workflow=Test%20Web) Stars:`10.9K`.
* Gameboy
  * [mohanson/gameboy](https://github.com/mohanson/gameboy) — Full featured Cross-platform GameBoy emulator. Forever boys!. Stars:`1.2K`.
  * [Gekkio/mooneye-gb](https://github.com/Gekkio/mooneye-gb) — [![build badge](https://api.travis-ci.org/Gekkio/mooneye-gb.svg?branch=master)](https://travis-ci.org/Gekkio/mooneye-gb) Stars:`788`.
  * [mvdnes/rboy](https://github.com/mvdnes/rboy) — [![build badge](https://api.travis-ci.org/mvdnes/rboy.svg?branch=master)](https://travis-ci.org/mvdnes/rboy) Stars:`511`.
* Gameboy Advance
  * [michelhe/rustboyadvance-ng](https://github.com/michelhe/rustboyadvance-ng) - RustboyAdvance-ng is a Gameboy Advance emulator with desktop, android and [WebAssembly](https://michelhe.github.io/rustboyadvance-ng/) support. [![build badge](https://github.com/michelhe/rustboyadvance-ng/workflows/Deploy/badge.svg?branch=master)](https://github.com/michelhe/rustboyadvance-ng/actions?query=workflow%3ADeploy) Stars:`493`.
* GameMaker
  * [OpenGMK](https://github.com/OpenGMK/OpenGMK) — OpenGMK is a modern rewrite of the proprietary GameMaker Classic engines, providing a full sourceport of the runner, a decompiler, a TASing framework, and libraries for working with gamedata yourself. Stars:`139`.
* Intel 8080 CPU
  * [mohanson/i8080](https://github.com/mohanson/i8080) — Intel 8080 cpu emulator by Rust Stars:`80`.
* NES
  * [pcwalton/sprocketnes](https://github.com/pcwalton/sprocketnes) Stars:`722`.
  * [koute/pinky](https://github.com/koute/pinky) — [![build badge](https://api.travis-ci.org/koute/pinky.svg?branch=master)](https://travis-ci.org/koute/pinky) Stars:`700`.
* ZX Spectrum
  * [rustzx/rustzx](https://github.com/rustzx/rustzx) — [![RustZX CI](https://github.com/rustzx/rustzx/actions/workflows/ci.yml/badge.svg)](https://github.com/rustzx/rustzx/actions/workflows/ci.yml) Stars:`157`.

### Games

See also [Games Made With Piston](https://github.com/PistonDevelopers/piston/wiki/Games-Made-With-Piston).

* [citybound](https://github.com/citybound/citybound) — The city sim you deserve Stars:`7.2K`.
* [cristicbz/rust-doom](https://github.com/cristicbz/rust-doom) — A renderer for Doom, may progress to being a playable game [![build badge](https://api.travis-ci.org/cristicbz/rust-doom.svg?branch=master)](https://travis-ci.org/cristicbz/rust-doom) Stars:`2.2K`.
* [ozkriff/zemeroth](https://github.com/ozkriff/zemeroth) — A small 2D turn-based hexagonal strategy game [![build badge](https://api.travis-ci.org/ozkriff/zemeroth.svg?branch=master)](https://travis-ci.org/ozkriff/zemeroth) Stars:`1.2K`.
* [gorilla-devs/ferium](https://github.com/gorilla-devs/ferium) — Ferium is a fast and feature rich CLI program for downloading and updating Minecraft mods from Modrinth, CurseForge, and GitHub Releases, and modpacks from Modrinth and CurseForge ![ferium build](https://github.com/gorilla-devs/ferium/actions/workflows/build.yml/badge.svg?branch=main) Stars:`559`.
* [doukutsu-rs](https://github.com/doukutsu-rs/doukutsu-rs) — A Rust reimplementation of Cave Story engine with some enhancements. Stars:`532`.
* [garkimasera/rusted-ruins](https://github.com/garkimasera/rusted-ruins) — Extensible open world rogue like game with pixel art [![build badge](https://api.travis-ci.org/garkimasera/rusted-ruins.svg?branch=master)](https://travis-ci.org/garkimasera/rusted-ruins) Stars:`427`.
* [Zone of Control](https://github.com/ozkriff/zoc) — A turn-based hexagonal strategy game [![build badge](https://api.travis-ci.org/ozkriff/zoc.svg?branch=master)](https://travis-ci.org/ozkriff/zoc) Stars:`354`.
* [rsaarelm/magog](https://github.com/rsaarelm/magog) — A roguelike game in Rust Stars:`346`.
* [thetawavegame/thetawave-legacy](https://github.com/thetawavegame/thetawave-legacy) - A space shooter game that strives to be an entry point for new game developers to make their first contributions. ![build badge](https://github.com/thetawavegame/thetawave-legacy/actions/workflows/ci.yml/badge.svg?branch=master) Stars:`192`.
* [SoftbearStudios/mk48](https://github.com/SoftbearStudios/mk48) — Mk48.io is an online multiplayer naval combat game Stars:`156`.
* [rhex](https://github.com/dpc/rhex) — hexagonal ascii roguelike Stars:`134`.
* [swatteau/sokoban-rs](https://github.com/swatteau/sokoban-rs) — A Sokoban implementation Stars:`134`.
* [lifthrasiir/angolmois-rust](https://github.com/lifthrasiir/angolmois-rust) — A minimalistic music video game which supports the BMS format [![build badge](https://api.travis-ci.org/lifthrasiir/angolmois-rust.svg?branch=master)](https://travis-ci.org/lifthrasiir/angolmois-rust) Stars:`95`.
* [mara214/rsnake](https://github.com/mara214/rsnake) — Snake written in Rust. Stars:`86`.
* [Thinkofname/rust-quake](https://github.com/Thinkofname/rust-quake) — Quake map renderer in Rust Stars:`56`.
* [ttyperacer/terminal-typeracer](https://gitlab.com/ttyperacer/terminal-typeracer) - Single player typing test game written for the terminal
* [Veloren](https://gitlab.com/veloren/veloren) — An open world, open source multiplayer voxel RPG game currently in alpha development [![build badge](https://gitlab.com/veloren/veloren/badges/master/pipeline.svg)](https://gitlab.com/veloren/veloren/-/pipelines)

### Graphics

* [ivanceras/svgbob](https://github.com/ivanceras/svgbob) — converts ASCII diagrams into SVG graphics [![build badge](https://api.travis-ci.org/ivanceras/svgbob.svg?branch=master)](https://travis-ci.org/ivanceras/svgbob) Stars:`3.3K`.
* [RazrFalcon/resvg](https://github.com/RazrFalcon/resvg) — An SVG rendering library. [![build badge](https://api.travis-ci.org/RazrFalcon/resvg.svg?branch=master)](https://travis-ci.org/RazrFalcon/resvg) Stars:`1.7K`.
* [Twinklebear/tray_rust](https://github.com/Twinklebear/tray_rust) — A ray tracer [![build badge](https://api.travis-ci.org/Twinklebear/tray_rust.svg?branch=master)](https://travis-ci.org/Twinklebear/tray_rust) Stars:`484`.
* [Limeth/euclider](https://github.com/Limeth/euclider) — A real-time 4D CPU ray tracer [![build badge](https://api.travis-ci.org/Limeth/euclider.svg?branch=master)](https://travis-ci.org/Limeth/euclider) Stars:`186`.
* [turnage/valora](https://crates.io/crates/valora) — A library for generative fine art ![Rust](https://github.com/turnage/valora/workflows/Rust/badge.svg?branch=master)

### Image processing

* [shssoichiro/oxipng](https://github.com/shssoichiro/oxipng) [[oxipng](https://crates.io/crates/oxipng)] — Multithreaded PNG optimizer written in Rust. [![Build Status](https://github.com/shssoichiro/oxipng/workflows/oxipng/badge.svg)](https://github.com/shssoichiro/oxipng/actions?query=branch%3Amaster) [![Version](https://img.shields.io/crates/v/oxipng.svg)](https://crates.io/crates/oxipng) Stars:`1.7K`.
* [Imager](https://github.com/imager-io/imager) — Automated image optimization. Stars:`469`.

### Industrial automation

* [locka99/opcua](https://github.com/locka99/opcua) — A pure rust [OPC UA](https://opcfoundation.org/about/opc-technologies/opc-ua/) library. Stars:`349`.
* [slowtec/tokio-modbus](https://github.com/slowtec/tokio-modbus) — A [tokio](https://tokio.rs)-based [modbus](https://modbus.org) library. [![Build Status](https://api.travis-ci.org/slowtec/tokio-modbus.svg?branch=master)](https://travis-ci.org/slowtec/tokio-modbus) Stars:`214`.

### Observability

* [vectordotdev/vector](https://github.com/vectordotdev/vector) — A High-Performance, Logs, Metrics, & Events Router. Stars:`11.6K`.
* [Quickwit-oss/quickwit](https://github.com/quickwit-oss/quickwit) - Cloud-native and highly cost-efficient search engine for log management. [![CI](https://github.com/quickwit-oss/quickwit/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/quickwit-oss/quickwit/actions?query=workflow%3ACI) Stars:`2.6K`.
* [Scaphandre](https://github.com/hubblo-org/scaphandre) - A power consumption monitoring agent, to track host and each service power consumption and enable designing systems and applications for more sustainability. Designed to fit any monitoring toolchain (already supports prometheus, warp10, riemann...). Stars:`839`.
* [avito-tech/bioyino](https://github.com/avito-tech/bioyino) — A high-performance scalable StatsD compatible server. Stars:`208`.
* [OpenTelemetry](https://crates.io/crates/opentelemetry) — OpenTelemetry provides a single set of APIs, libraries, agents, and collector services to capture distributed traces and metrics from your application. You can analyze them using Prometheus, Jaeger, and other observability tools. [![GitHub Actions CI](https://github.com/open-telemetry/opentelemetry-rust/workflows/CI/badge.svg?branch=master)](https://github.com/open-telemetry/opentelemetry-rust/actions?query=workflow%3ACI+branch%3Amaster)

### Operating systems

See also [A comparison of operating systems written in Rust](https://github.com/flosse/rust-os-comparison).
* [tock/tock](https://github.com/tock/tock) — A secure embedded operating system for Cortex-M based microcontrollers Stars:`3.9K`.
* [theseus-os/Theseus](https://github.com/theseus-os/Theseus) — A safe-language, single address space and single privilege level OS written from scratch in pure Rust - [![build badge](https://img.shields.io/github/workflow/status/theseus-os/Theseus/Documentation?label=docs%20build)](https://www.theseus-os.com/Theseus/book/index.html) Stars:`2.1K`.
* [thepowersgang/rust_os](https://github.com/thepowersgang/rust_os) — [![build badge](https://api.travis-ci.org/thepowersgang/rust_os.svg?branch=master)](https://travis-ci.org/thepowersgang/rust_os) Stars:`597`.
* [0x59616e/SteinsOS](https://github.com/0x59616e/SteinsOS)  — An OS for armv8-a architecture. Stars:`84`.
* [redox-os/redox](https://gitlab.redox-os.org/redox-os/redox) — [![build badge](https://api.travis-ci.org/redox-os/redox.svg?branch=master)](https://travis-ci.org/redox-os/redox)

### Productivity

* [espanso](https://github.com/espanso/espanso) — A cross-platform Text Expander written in Rust[![CI](https://github.com/espanso/espanso/actions/workflows/ci.yml/badge.svg?branch=dev&event=push)](https://github.com/espanso/espanso/actions/workflows/ci.yml) Stars:`6.2K`.
* [pier-cli/pier](https://github.com/pier-cli/pier) — A central repository to manage (add, search metadata, etc.) all your one-liners, scripts, tools, and CLIs Stars:`398`.
* [Bartib](https://github.com/nikolassv/bartib) [[Bartib](https://crates.io/crates/bartib)] - A simple timetracker for the command line [![Tests](https://github.com/nikolassv/bartib/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/nikolassv/bartib/actions/workflows/test.yml) Stars:`316`.
* [eureka](https://crates.io/crates/eureka) — A CLI tool to input and store your ideas without leaving the terminal

### Security tools

* [rustscan/rustscan](https://github.com/RustScan/RustScan) — Make Nmap faster with this port scanning tool [![build badge](https://github.com/RustScan/RustScan/workflows/Continuous%20integration/badge.svg?branch=master)](https://github.com/RustScan/RustScan/actions?query=workflow%3A%22Continuous+integration%22) Stars:`7.5K`.
* [epi052/feroxbuster](https://github.com/epi052/feroxbuster) - A simple, fast, recursive content discovery tool written in Rust ( Stars:`3.3K`.
* [kpcyrd/sn0int](https://github.com/kpcyrd/sn0int) — A semi-automatic OSINT framework and package manager [![build badge](https://api.travis-ci.org/kpcyrd/sn0int.svg?branch=master)](https://travis-ci.org/kpcyrd/sn0int) Stars:`1.4K`.
* [AFLplusplus/LibAFL](https://github.com/AFLplusplus/LibAFL) - Advanced Fuzzing Library - Slot your Fuzzer together in Rust! Scales across cores and machines. For Windows, Android, MacOS, Linux, no_std, etc. [![build and test](https://github.com/AFLplusplus/LibAFL/actions/workflows/build_and_test.yml/badge.svg)](https://github.com/AFLplusplus/LibAFL/actions/workflows/build_and_test.yml) Stars:`1.1K`.
* [kpcyrd/sniffglue](https://github.com/kpcyrd/sniffglue) — A secure multithreaded packet sniffer [![build badge](https://api.travis-ci.org/kpcyrd/sniffglue.svg?branch=master)](https://travis-ci.org/kpcyrd/sniffglue) Stars:`904`.
* [Cherrybomb](https://github.com/blst-security/cherrybomb) - Stop half-done API specifications with a CLI tool that helps you avoid undefined user behaviour by validating your API specifications. Stars:`629`.
* [ObserverWard](https://github.com/0x727/ObserverWard) — Community based web technologies analysis tool. Stars:`492`.
* [phra/rustbuster](https://github.com/phra/rustbuster) — A Comprehensive Web Fuzzer and Content Discovery Tool Stars:`448`.
* [kpcyrd/rshijack](https://github.com/kpcyrd/rshijack) — A TCP connection hijacker, rust rewrite of shijack [![build badge](https://api.travis-ci.org/kpcyrd/rshijack.svg?branch=master)](https://travis-ci.org/kpcyrd/rshijack) Stars:`371`.
* [kpcyrd/authoscope](https://github.com/kpcyrd/authoscope) — A scriptable network authentication cracker [![build badge](https://api.travis-ci.org/kpcyrd/authoscope.svg?branch=master)](https://travis-ci.org/kpcyrd/authoscope) Stars:`337`.
* [Inspektor](https://github.com/inspektor-dev/inspektor) - A database protocol-aware proxy that is used to enforce access policies 👮 Stars:`255`.
* [arvancloud/libinjection-rs](https://github.com/arvancloud/libinjection-rs) — Rust bindings for [libinjection](https://github.com/client9/libinjection) [![build badge](https://api.travis-ci.org/arvancloud/libinjection-rs.svg?branch=master)](https://travis-ci.org/arvancloud/libinjection-rs) Stars:`35`.
* [ripasso](https://github.com/cortex/ripasso/) — A password manager, filesystem compatible with pass

### Simulation

* [hEngine](https://github.com/hashintel/hash/tree/main/packages/engine) - A Rust-implemented computational simulation engine, supporting large-scale agent-based modelling, with simulation logic written in JavaScript and Python.

### System tools

* [sharkdp/bat](https://github.com/sharkdp/bat) — A cat(1) clone with wings. [![CICD](https://github.com/sharkdp/bat/actions/workflows/CICD.yml/badge.svg?branch=master)](https://github.com/sharkdp/bat/actions/workflows/CICD.yml) Stars:`38.0K`.
* [sharkdp/fd](https://github.com/sharkdp/fd) — A simple, fast and user-friendly alternative to find. [![CICD](https://github.com/sharkdp/fd/actions/workflows/CICD.yml/badge.svg)](https://github.com/sharkdp/fd/actions/workflows/CICD.yml) Stars:`25.2K`.
* [ogham/exa](https://github.com/ogham/exa) — A replacement for 'ls' [![build badge](https://api.travis-ci.org/ogham/exa.svg?branch=master)](https://travis-ci.org/ogham/exa) Stars:`19.7K`.
* [uutils/coreutils](https://github.com/uutils/coreutils) — A cross-platform Rust rewrite of the GNU coreutils [[![CICD](https://github.com/uutils/coreutils/actions/workflows/CICD.yml/badge.svg)](https://github.com/uutils/coreutils/actions/workflows/CICD.yml) Stars:`12.8K`.
* [gitui](https://github.com/extrawurst/gitui) - Blazing fast terminal client for git written in Rust. [![build](https://github.com/extrawurst/gitui/workflows/CI/badge.svg?branch=master)](https://github.com/extrawurst/gitui/actions) Stars:`11.3K`.
* [Peltoche/lsd](https://github.com/Peltoche/lsd) — An ls with a lot of pretty colors and awesome icons [![build](https://github.com/Peltoche/lsd/workflows/CICD/badge.svg?branch=master)](https://github.com/Peltoche/lsd/actions) Stars:`8.7K`.
* [bandwhich](https://github.com/imsnif/bandwhich) — Terminal bandwidth utilization tool [![build badge](https://api.travis-ci.com/imsnif/bandwhich.svg?branch=master)](https://app.travis-ci.com/github/imsnif/bandwhich) Stars:`7.4K`.
* [XAMPPRocky/tokei](https://github.com/XAMPPRocky/tokei) — counts the lines of code [![build badge](https://api.travis-ci.org/XAMPPRocky/tokei.svg?branch=master)](https://travis-ci.org/XAMPPRocky/tokei) Stars:`7.3K`.
* [bottom](https://github.com/ClementTsang/bottom) - Yet another cross-platform graphical process/system monitor. [![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/ClementTsang/bottom/ci/master)](https://github.com/ClementTsang/bottom/actions?query=branch%3Amaster) Stars:`5.3K`.
* [dust](https://github.com/bootandy/dust) — A more intuitive version of du Stars:`5.2K`.
* [cantino/mcfly](https://github.com/cantino/mcfly) - Fly through your shell history. Great Scott! [![build badge](https://api.travis-ci.org/cantino/mcfly.svg?branch=master)](https://travis-ci.org/cantino/mcfly) Stars:`4.6K`.
* [lotabout/skim](https://github.com/lotabout/skim) — A fuzzy finder in pure rust [![build badge](https://api.travis-ci.org/lotabout/skim.svg?branch=master)](https://travis-ci.org/lotabout/skim) Stars:`3.6K`.
* [watchexec](https://github.com/watchexec/watchexec) — Executes commands in response to file modifications [![build badge](https://api.travis-ci.org/watchexec/watchexec.svg?branch=master)](https://travis-ci.org/watchexec/watchexec) Stars:`3.5K`.
* [dalance/procs](https://github.com/dalance/procs) — A modern replacement for 'ps' written by Rust [![Regression](https://github.com/dalance/procs/actions/workflows/regression.yml/badge.svg)](https://github.com/dalance/procs/actions/workflows/regression.yml) Stars:`3.4K`.
* [pueue](https://github.com/nukesor/pueue) — Manage your long running shell commands. [![GitHub Actions Workflow](https://github.com/nukesor/pueue/workflows/Test%20build/badge.svg?branch=master)](https://github.com/nukesor/pueue/actions) Stars:`3.2K`.
* [orhun/kmon](https://github.com/orhun/kmon) — Linux Kernel Manager and Activity Monitor ![https://github.com/orhun/kmon/actions](https://img.shields.io/github/workflow/status/orhun/kmon/Continuous%20Integration/master?label=build) Stars:`1.7K`.
* [m4b/bingrep](https://github.com/m4b/bingrep) — Greps through binaries from various OSs and architectures, and colors them. [![build badge](https://api.travis-ci.org/m4b/bingrep.svg?branch=master)](https://travis-ci.org/m4b/bingrep) Stars:`1.6K`.
* [diskonaut](https://github.com/imsnif/diskonaut) — Terminal visual disk space navigator [![build badge](https://api.travis-ci.com/imsnif/diskonaut.svg?branch=main)](https://app.travis-ci.com/github/imsnif/diskonaut) Stars:`1.5K`.
* [redox-os/ion](https://github.com/redox-os/ion) — Next-generation system shell [![build badge](https://api.travis-ci.org/redox-os/ion.svg?branch=master)](https://travis-ci.org/redox-os/ion) Stars:`1.3K`.
* [mitnk/cicada](https://github.com/mitnk/cicada) — A bash-like Unix shell [![build badge](https://api.travis-ci.org/mitnk/cicada.svg?branch=master)](https://travis-ci.org/mitnk/cicada) Stars:`907`.
* [Kondo](https://github.com/tbillington/kondo) - CLI & GUI tool for deleting software project artifacts and reclaiming disk space Stars:`888`.
* [Luminarys/synapse](https://github.com/Luminarys/synapse) — Flexible and fast BitTorrent daemon. [![Build Status](https://api.travis-ci.org/Luminarys/synapse.svg?branch=master)](https://travis-ci.org/Luminarys/synapse) Stars:`802`.
* [orhun/systeroid](https://github.com/orhun/systeroid) — A more powerful alternative to sysctl(8) with a terminal user interface ![https://github.com/orhun/systeroid/actions](https://img.shields.io/github/workflow/status/orhun/systeroid/Continuous%20Integration/main?label=build) Stars:`717`.
* [nivekuil/rip](https://github.com/nivekuil/rip) - A safe and ergonomic alternative to `rm` [![build badge](https://api.travis-ci.org/nivekuil/rip.svg?branch=master)](https://travis-ci.org/nivekuil/rip) Stars:`685`.
* [ouch](https://github.com/ouch-org/ouch) - Painless compression and decompression on the command-line [![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/ouch-org/ouch/build-and-test)](https://github.com/ouch-org/ouch/actions?query=branch%3Amaster) Stars:`669`.
* [pop-os/popsicle](https://github.com/pop-os/popsicle) — GTK3 & CLI utility for flashing multiple USB devices in parallel Stars:`464`.
* [sitkevij/hex](https://github.com/sitkevij/hex) — A colorized hexdump terminal utility. [![build badge](https://api.travis-ci.org/sitkevij/hex.svg?branch=master)](https://travis-ci.org/sitkevij/hex) Stars:`382`.
* [lotabout/rargs](https://github.com/lotabout/rargs) [[rargs](https://crates.io/crates/rargs)] — xargs + awk with pattern matching support [![build badge](https://api.travis-ci.org/lotabout/rargs.svg?branch=master)](https://travis-ci.org/lotabout/rargs) Stars:`373`.
* [ddh](https://github.com/darakian/ddh) — Fast duplicate file finder [![build badge](https://api.travis-ci.org/darakian/ddh.svg?branch=master)](https://travis-ci.org/darakian/ddh) Stars:`371`.
* [mmstick/fontfinder](https://github.com/mmstick/fontfinder) — GTK3 application for previewing and installing Google's fonts Stars:`247`.
* [brocode/fblog](https://github.com/brocode/fblog) — Small command-line JSON Log viewer [![build badge](https://api.travis-ci.org/brocode/fblog.svg?branch=master)](https://travis-ci.org/brocode/fblog) Stars:`243`.
* [crabz](https://github.com/sstadick/crabz) - Multi-threaded compression and decompression CLI tool [![Build Status](https://github.com/sstadick/crabz/workflows/Check/badge.svg)](https://github.com/sstadick/crabz/actions?query=workflow%3ACheck) Stars:`226`.
* [nickgerace/gfold](https://github.com/nickgerace/gfold) [[gfold](https://crates.io/crates/gfold)] - CLI tool to help keep track of multiple Git repositories [![build](https://img.shields.io/github/workflow/status/nickgerace/gfold/merge/main)](https://github.com/nickgerace/gfold/actions?query=workflow%3Amerge+branch%3Amain) Stars:`195`.
* [cristianoliveira/funzzy](https://github.com/cristianoliveira/funzzy) — A configurable filesystem watcher inspired by [entr](http://eradman.com/entrproject/) [![build badge](https://api.travis-ci.org/cristianoliveira/funzzy.svg?branch=master)](https://travis-ci.org/cristianoliveira/funzzy) Stars:`185`.
* [bustd](https://github.com/vrmiguel/bustd) - Lightweight process killer daemon to handle out-of-memory scenarios on Linux. [![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/vrmiguel/bustd/build-and-test)](https://github.com/vrmiguel/bustd/actions?query=branch%3Amaster) Stars:`182`.
* [mmstick/tv-renamer](https://github.com/mmstick/tv-renamer) — A tv series renaming application with an optional GTK3 frontend. [![build badge](https://api.travis-ci.org/mmstick/tv-renamer.svg?branch=master)](https://travis-ci.org/mmstick/tv-renamer) Stars:`143`.
* [buster/rrun](https://github.com/buster/rrun) — A command launcher for Linux, similar to gmrun [![build badge](https://api.travis-ci.org/buster/rrun.svg?branch=master)](https://travis-ci.org/buster/rrun) Stars:`104`.
* [trippy](https://github.com/fujiapple852/trippy) - A network diagnostic tool [![build badge](https://github.com/fujiapple852/trippy/workflows/CI/badge.svg)](https://github.com/fujiapple852/trippy/actions/workflows/ci.yml) Stars:`103`.
* [mmstick/concurr](https://github.com/mmstick/concurr) — Alternative to GNU Parallel w/ a client-server architecture Stars:`92`.
* [mxseev/logram](https://github.com/mxseev/logram) — Push log files' updates to Telegram Stars:`83`.
* [Alonely0/Voila](https://github.com/Alonely0/Voila) — Voila is a domain-specific language launched through CLI tool for operating with files and directories in massive amounts in a fast & reliable way. [![Linux build](https://github.com/Alonely0/Voila/actions/workflows/linux-ci.yml/badge.svg)](https://github.com/Alonely0/Voila/actions/workflows/linux-ci.yml) [![macOS build](https://github.com/Alonely0/Voila/actions/workflows/mac-ci.yml/badge.svg)](https://github.com/Alonely0/Voila/actions/workflows/mac-ci.yml) [![Windows build](https://github.com/Alonely0/Voila/actions/workflows/windows-ci.yml/badge.svg)](https://github.com/Alonely0/Voila/actions/workflows/windows-ci.yml) Stars:`83`.
* [pop-os/system76-power](https://github.com/pop-os/system76-power/) — Linux power management daemon (DBus-interface) with CLI tool.
* [fselect](https://crates.io/crates/fselect) — Find files with SQL-like queries [![build badge](https://api.travis-ci.org/jhspetersson/fselect.svg?branch=master)](https://travis-ci.org/jhspetersson/fselect)
* [ajeetdsouza/zoxide](https://github.com/ajeetdsouza/zoxide/) — A fast alternative to `cd` that learns your habits [![release](https://github.com/ajeetdsouza/zoxide/workflows/.github/workflows/release.yml/badge.svg)](https://github.com/ajeetdsouza/zoxide/actions)

### Task scheduling

* [delicate](https://github.com/BinChengZhao/delicate) — A lightweight and distributed task scheduling platform written in rust. [![Build Status](https://github.com/BinChengZhao/delicate/workflows/CI/badge.svg)](https://github.com/BinChengZhao/delicate/actions) Stars:`500`.

### Text editors

* [lapce](https://github.com/lapce/lapce) — Lightning-fast and Powerful Code Editor written in Rust. [![build badge](https://github.com/lapce/lapce/actions/workflows/release.yml/badge.svg)](https://github.com/lapce/lapce/actions/workflows/release.yml) Stars:`20.1K`.
* [xi-editor](https://github.com/xi-editor/xi-editor) — A modern editor with a backend written in Rust. Stars:`19.7K`.
* [helix](https://github.com/helix-editor/helix) — A post-modern modal text editor inspired by Neovim/Kakoune. [![build badge](https://github.com/helix-editor/helix/actions/workflows/build.yml/badge.svg)](https://github.com/helix-editor/helix/actions) Stars:`15.9K`.
* [Remacs](https://github.com/remacs/remacs) — A community-driven port of Emacs to Rust. [![build badge](https://api.travis-ci.org/remacs/remacs.svg?branch=master)](https://travis-ci.org/remacs/remacs) Stars:`4.5K`.
* [ox](https://github.com/curlpipe/ox) — An independent Rust text editor that runs in your terminal! Stars:`2.8K`.
* [gchp/iota](https://github.com/gchp/iota) — A simple text editor [![build badge](https://api.travis-ci.org/gchp/iota.svg?branch=master)](https://travis-ci.org/gchp/iota) Stars:`1.5K`.
* [ilai-deutel/kibi](https://github.com/ilai-deutel/kibi) — A tiny (≤1024 LOC) text editor with syntax highlighting, incremental search and more. [![build badge](https://github.com/ilai-deutel/kibi/workflows/CI/badge.svg?branch=master)](https://github.com/ilai-deutel/kibi/actions?query=branch%3Amaster) Stars:`849`.
* [mathall/rim](https://github.com/mathall/rim) — Vim-like text editor written in Rust Stars:`553`.
* [vamolessa/pepper](https://github.com/vamolessa/pepper) [[pepper](https://crates.io/crates/pepper)] — An opinionated modal editor to simplify code editing from the terminal [![build badge](https://github.com/vamolessa/pepper/workflows/rust/badge.svg?branch=master)](https://github.com/vamolessa/pepper) Stars:`281`.
* [amp](https://amp.rs) — Inspired by Vi/Vim. [![build badge](https://api.travis-ci.org/jmacdonald/amp.svg?branch=master)](https://travis-ci.org/jmacdonald/amp)

### Text processing

* [grex](https://github.com/pemistahl/grex) — A command-line tool and library for generating regular expressions from user-provided test cases [![build badge](https://api.travis-ci.org/pemistahl/grex.svg?branch=master)](https://travis-ci.org/pemistahl/grex) Stars:`5.7K`.
* [phiresky/ripgrep-all](https://github.com/phiresky/ripgrep-all) — ripgrep, but also search in PDFs, E-Books, Office documents, zip, tar.gz, etc. [![Build Status](https://api.travis-ci.org/phiresky/ripgrep-all.svg?branch=master)](https://travis-ci.org/phiresky/ripgrep-all) Stars:`5.1K`.
* [Melody](https://github.com/yoav-lavi/melody) - A language that compiles to regular expressions and aims to be more easily readable and maintainable [![build badge](https://github.com/yoav-lavi/melody/actions/workflows/rust.yml/badge.svg)](https://github.com/yoav-lavi/melody/actions/workflows/rust.yml) [![crates.io](https://img.shields.io/crates/v/melody_compiler?label=compiler)](https://crates.io/crates/melody_compiler) Stars:`4.0K`.
* [dominikwilkowski/cfonts](https://github.com/dominikwilkowski/cfonts) [[cfonts](https://crates.io/crates/cfonts)] — Sexy ANSI fonts for the console ![build badge](https://github.com/dominikwilkowski/cfonts/actions/workflows/testing.yml/badge.svg) Stars:`1.2K`.
* [sstadick/hck](https://github.com/sstadick/hck) - A faster and more featureful drop in replacement for `cut`  [![build badge](https://github.com/sstadick/hck/workflows/Check/badge.svg?branch=master)](https://github.com/sstadick/hck) Stars:`593`.
* [ruplacer](https://github.com/your-tools/ruplacer) — Find and replace text in source files [![Run tests](https://github.com/your-tools/ruplacer/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/your-tools/ruplacer/actions/workflows/test.yml) Stars:`324`.
* [vishaltelangre/ff](https://github.com/vishaltelangre/ff) — Find files (ff) by name! [![build badge](https://api.travis-ci.org/vishaltelangre/ff.svg?branch=master)](https://travis-ci.org/vishaltelangre/ff) Stars:`307`.
* [whitfin/runiq](https://github.com/whitfin/runiq) — an efficient way to filter duplicate lines from unsorted input. Stars:`167`.
* [Lisprez/so_stupid_search](https://github.com/Lisprez/so_stupid_search) — A simple and fast string search tool for human beings Stars:`135`.
* [whitfin/bytelines](https://github.com/whitfin/bytelines) [[bytelines](https://crates.io/crates/bytelines)] — Read input lines as byte slices for high efficiency. Stars:`53`.
* [replicadse/complate](https://github.com/replicadse/complate) — An in-terminal text templating tool designed for standardizing messages (like for GIT commits). [![crates.io](https://img.shields.io/crates/v/complate.svg)](https://crates.io/crates/complate) [![crates.io](https://img.shields.io/crates/d/complate?label=crates.io%20downloads)](https://crates.io/crates/complate) [![build badge](https://github.com/replicadse/complate/workflows/pipeline/badge.svg?branch=master)](https://github.com/replicadse/complate/actions) Stars:`20`.
* [ripgrep](https://crates.io/crates/ripgrep) — combines the usability of The Silver Searcher with the raw speed of grep [![build badge](https://api.travis-ci.org/BurntSushi/ripgrep.svg?branch=master)](https://travis-ci.org/BurntSushi/ripgrep)
* [sd](https://crates.io/crates/sd) — Intuitive find & replace CLI
* [xsv](https://crates.io/crates/xsv) — A fast CSV command line tool (slicing, indexing, selecting, searching, sampling, etc.) [![build badge](https://api.travis-ci.org/BurntSushi/xsv.svg?branch=master)](https://travis-ci.org/BurntSushi/xsv)

### Utilities

* [1History](https://github.com/1History/1History) — Command line interface to backup Firefox/Chrome/Safari history to one SQLite file [![Build Status](https://github.com/1History/1History/actions/workflows/CI.yml/badge.svg)](https://github.com/1History/1History/actions/workflows/CI.yml) Stars:`330`.
* [evansmurithi/cloak](https://github.com/evansmurithi/cloak) — A Command Line OTP (One Time Password) Authenticator application. Stars:`195`.
* [brycx/checkpwn](https://github.com/brycx/checkpwn) — A Have I Been Pwned (HIBP) command-line utility tool that lets you easily check for compromised accounts and passwords. Stars:`91`.
![CI](https://github.com/evansmurithi/cloak/workflows/CI/badge.svg) [![build badge](https://ci.appveyor.com/api/projects/status/9mlfpfru3ng4c689/branch/master?svg=true)](https://ci.appveyor.com/project/evansmurithi/cloak)
* [rustdesk/rustdesk](https://github.com/rustdesk/rustdesk) — A remote desktop software, great alternative to TeamViewer and AnyDesk. Stars:`32.0K`.
* [vaultwarden](https://github.com/dani-garcia/vaultwarden#readme) [![Build](https://github.com/dani-garcia/vaultwarden/actions/workflows/build.yml/badge.svg)](https://github.com/dani-garcia/vaultwarden/actions/workflows/build.yml) — Alternative implementation of the Bitwarden server API written in Rust Stars:`19.7K`.
* [fcsonline/tmux-thumbs](https://github.com/fcsonline/tmux-thumbs) — A lightning fast version of tmux-fingers written in Rust, copy/pasting tmux like vimium/vimperator. Stars:`576`.
* [mprocs](https://github.com/pvolok/mprocs) — TUI for running multiple processes Stars:`565`.
* [nomino](https://github.com/yaa110/nomino) — Batch rename utility for developers [![Build Status](https://api.travis-ci.org/yaa110/nomino.svg?branch=master)](https://travis-ci.org/yaa110/nomino) Stars:`440`.
* [guoxbin/dtool](https://github.com/guoxbin/dtool) — A useful command-line tool collection to assist development including conversion, codec, hashing, encryption, etc. [![Build Status](https://api.travis-ci.org/guoxbin/dtool.svg?branch=master)](https://travis-ci.org/guoxbin/dtool) Stars:`306`.
* [vamolessa/verco](https://github.com/vamolessa/verco) [[verco](https://crates.io/crates/verco)] — A simple Git/Hg tui client focused on keyboard shortcuts Stars:`208`.
* [tversteeg/emplace](https://github.com/tversteeg/emplace) — Synchronize installed packages on multiple machines Stars:`198`.
* [raftario/licensor](https://github.com/raftario/licensor) — write licenses to stdout [![GitHub Actions](https://github.com/raftario/licensor/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/raftario/licensor/actions/workflows/build.yml) Stars:`159`.
* [yaa110/cb](https://github.com/yaa110/cb) — Command line interface to manage clipboard [![Build Status](https://api.travis-ci.org/yaa110/cb.svg?branch=master)](https://travis-ci.org/yaa110/cb) Stars:`73`.

### Video

* [xiph/rav1e](https://github.com/xiph/rav1e) — The fastest and safest AV1 encoder. [![build badge](https://api.travis-ci.org/xiph/rav1e.svg?branch=master)](https://travis-ci.org/xiph/rav1e) Stars:`3.0K`.
* [harlanc/xiu](https://github.com/harlanc/xiu) — A powerful and secure live server by pure rust (rtmp/httpflv/hls/relay). [![Build Status](https://api.travis-ci.com/harlanc/xiu.svg?branch=master)](https://app.travis-ci.com/github/harlanc/xiu) [![crates.io](https://img.shields.io/crates/v/xiu.svg)](https://crates.io/crates/xiu) Stars:`543`.
* [dertuxmalwieder/yaydl](https://github.com/dertuxmalwieder/yaydl) [[yaydl](https://crates.io/crates/yaydl)] - A simple video downloader Stars:`199`.

### Virtualization

* [firecracker-microvm/firecracker](https://github.com/firecracker-microvm/firecracker) — A lightweight virtual machine for container workload  [Firecracker Microvm](https://firecracker-microvm.github.io/) Stars:`19.8K`.
* [containers/youki](https://github.com/containers/youki) — A container runtime in Rust [![build badge](https://github.com/containers/youki/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/containers/youki/actions) Stars:`3.9K`.
* [tailhook/vagga](https://github.com/tailhook/vagga) — A containerization tool without daemons [![build badge](https://api.travis-ci.org/tailhook/vagga.svg?branch=master)](https://travis-ci.org/tailhook/vagga) Stars:`1.8K`.

### Web

* [LemmyNet/lemmy](https://github.com/LemmyNet/lemmy) — A link aggregator / reddit clone for the fediverse [![Build Status](https://cloud.drone.io/api/badges/LemmyNet/lemmy/status.svg)](https://cloud.drone.io/LemmyNet/lemmy) Stars:`6.9K`.
* [libreddit](https://github.com/libreddit/libreddit) - An alternative private front-end to Reddit Stars:`3.6K`.
* [Plume-org/Plume](https://github.com/Plume-org/Plume) — ActivityPub federating blogging application [![build badge](https://api.travis-ci.org/Plume-org/Plume.svg?branch=master)](https://travis-ci.org/Plume-org/Plume) Stars:`1.8K`.
* [Revolt/backend](https://github.com/revoltchat/backend) - User-first chat platform built with modern web technologies. Stars:`717`.
* [MASQ-Project/Node](https://github.com/MASQ-Project/Node) — MASQ Node software provides a decentralized mesh-network of nodes for global users to access normal internet content - next evolution of tech beyond Tor & VPN [![build badge](https://github.com/MASQ-Project/Node/actions/workflows/ci-matrix.yml/badge.svg)](https://github.com/MASQ-Project/Node/actions) Stars:`122`.
* [cfal/tobaru](https://github.com/cfal/tobaru) - Port forwarder with allowlists, IP and TLS SNI/ALPN rule-based routing, iptables support, round-robin forwarding (load balancing), and hot reloading. Stars:`100`.

### Web Servers

* [svenstaro/miniserve](https://github.com/svenstaro/miniserve) — A small, self-contained cross-platform CLI tool that allows you to just grab the binary and serve some file(s) via HTTP [![build badge](https://github.com/svenstaro/miniserve/workflows/CI/badge.svg?branch=master)](https://github.com/svenstaro/miniserve/actions) Stars:`3.8K`.
* [TheWaWaR/simple-http-server](https://github.com/TheWaWaR/simple-http-server) — simple static http server Stars:`757`.
* [mufeedvh/binserve](https://github.com/mufeedvh/binserve) — A blazingly fast static web server with routing, templating, and security in a single binary you can set up with zero code [![build badge](https://github.com/mufeedvh/binserve/workflows/CICD/badge.svg?branch=master)](https://github.com/mufeedvh/binserve/actions) Stars:`695`.
* [static-web-server](https://github.com/static-web-server/static-web-server) — A blazing fast and asynchronous web server for static files-serving. ⚡ [![CI](https://github.com/static-web-server/static-web-server/actions/workflows/devel.yml/badge.svg)](https://github.com/static-web-server/static-web-server/actions/workflows/devel.yml?query=branch%3Amaster) Stars:`451`.
* [thecoshman/http](https://github.com/thecoshman/http) — Host These Things Please — A basic http server for hosting a folder fast and simply [![build badge](https://api.travis-ci.org/thecoshman/http.svg?branch=master)](https://travis-ci.org/thecoshman/http) Stars:`363`.
* [wyhaya/see](https://github.com/wyhaya/see) — Static HTTP file server [![Build Status](https://api.travis-ci.org/wyhaya/see.svg?branch=master)](https://travis-ci.org/wyhaya/see) Stars:`171`.
* [orhun/rustypaste](https://github.com/orhun/rustypaste) — A minimal file upload/pastebin service ![https://github.com/orhun/rustypaste/actions](https://img.shields.io/github/workflow/status/orhun/rustypaste/Continuous%20Integration/master?label=build) Stars:`147`.
* [ronanyeah/rust-hasura](https://github.com/ronanyeah/rust-hasura) — A demonstration of how a Rust GraphQL server can be used as a remote schema with [Hasura](https://hasura.io/) ![Rust](https://github.com/ronanyeah/rust-hasura/workflows/Rust/badge.svg?branch=master) Stars:`126`.

## Development tools

* [just](https://github.com/casey/just) — A handy command runner for project-specific tasks [![build badge](https://api.travis-ci.org/casey/just.svg?branch=master)](https://travis-ci.org/casey/just) Stars:`7.5K`.
* [Rustup](https://github.com/rust-lang/rustup) — the Rust toolchain installer [![build badge](https://github.com/rust-lang/rustup/workflows/Linux%20(master)/badge.svg?branch=master)](https://github.com/rust-lang/rustup/actions) Stars:`5.0K`.
* [git-cliff](https://github.com/orhun/git-cliff) — A highly customizable Changelog Generator that follows Conventional Commit specifications ![https://github.com/orhun/git-cliff/actions](https://img.shields.io/github/workflow/status/orhun/git-cliff/Continuous%20Integration/main?label=build) Stars:`4.7K`.
* [Racer](https://github.com/racer-rust/racer) — code completion for Rust [![build badge](https://api.travis-ci.org/racer-rust/racer.svg?branch=master)](https://travis-ci.org/racer-rust/racer) Stars:`3.4K`.
* [dotenv-linter](https://github.com/dotenv-linter/dotenv-linter) — Linter for `.env` files [![build badge](https://github.com/dotenv-linter/dotenv-linter/workflows/CI/badge.svg?branch=master)](https://github.com/dotenv-linter/dotenv-linter/actions?query=workflow%3ACI+branch%3Amaster) Stars:`1.5K`.
* [geiger](https://github.com/rust-secure-code/cargo-geiger) — A program that list statistics related to usage of unsafe Rust code in a Rust crate and all its dependencies [![Build Status](https://dev.azure.com/cargo-geiger/cargo-geiger/_apis/build/status/rust-secure-code.cargo-geiger?branchName=master)](https://dev.azure.com/cargo-geiger/cargo-geiger/_build/latest?definitionId=1&branchName=master) Stars:`1.0K`.
* [create-rust-app](https://github.com/Wulf/create-rust-app) — Set up a modern rust+react web app by running one command. [![crate](https://img.shields.io/crates/v/create-rust-app.svg)](https://crates.io/crates/create-rust-app) Stars:`968`.
* [Rust Search Extension](https://github.com/huhu/rust-search-extension) — A handy browser extension to search crates and docs in address bar (omnibox). [![Build Status](https://github.com/huhu/rust-search-extension/workflows/build/badge.svg?branch=master)](https://github.com/huhu/rust-search-extension/actions) Stars:`912`.
* [rust-lang/rustfix](https://github.com/rust-lang/rustfix)  — automatically applies the suggestions made by rustc Stars:`783`.
* [clog-tool/clog-cli](https://github.com/clog-tool/clog-cli) — generates a changelog from git metadata ([conventional changelog](https://blog.thoughtram.io/announcements/tools/2014/09/18/announcing-clog-a-conventional-changelog-generator-for-the-rest-of-us.html)) [![build badge](https://api.travis-ci.org/clog-tool/clog-cli.svg?branch=master)](https://travis-ci.org/clog-tool/clog-cli) Stars:`772`.
* [mask](https://github.com/jacobdeichert/mask) — A CLI task runner defined by a simple markdown file [![build badge](https://github.com/jacobdeichert/mask/workflows/CI/badge.svg?branch=master)](https://github.com/jacobdeichert/mask/actions?query=workflow%3ACI) Stars:`740`.
* [bacon](https://github.com/Canop/bacon) — background rust code checker, similar to cargo-watch Stars:`557`.
* [fw](https://github.com/brocode/fw) — workspace productivity booster [![Rust](https://github.com/brocode/fw/actions/workflows/rust.yml/badge.svg)](https://github.com/brocode/fw/actions/workflows/rust.yml) Stars:`404`.
* [datanymizer/datanymizer](https://github.com/datanymizer/datanymizer) - Powerful database anonymizer with flexible rules [![build badge](https://github.com/datanymizer/datanymizer/workflows/CI/badge.svg?branch=main)](https://github.com/datanymizer/datanymizer/actions?query=workflow%3ACI+branch%3Amain) Stars:`376`.
* [dan-t/rusty-tags](https://github.com/dan-t/rusty-tags) — create ctags/etags for a cargo project and all of its dependencies [![build badge](https://api.travis-ci.org/dan-t/rusty-tags.svg?branch=master)](https://travis-ci.org/dan-t/rusty-tags) Stars:`366`.
* [comtrya](https://github.com/comtrya/comtrya) — A configuration management tool for localhost / dotfiles [![build badge](https://github.com/comtrya/comtrya/actions/workflows/main.yaml/badge.svg)](https://github.com/comtrya/comtrya/actions) Stars:`255`.
* [Module Linker](https://github.com/fiatjaf/module-linker) — Extension that adds `<a>` links to references in `mod`, `use` and `extern crate` statements at GitHub. Stars:`247`.
* [scriptisto](https://github.com/igor-petruk/scriptisto) A language-agnostic "shebang interpreter" that enables you to write one file scripts in compiled languages. [![Build Status](https://cloud.drone.io/api/badges/igor-petruk/scriptisto/status.svg)](https://cloud.drone.io/igor-petruk/scriptisto) Stars:`230`.
* [hot-lib-reloader](https://github.com/rksm/hot-lib-reloader-rs) — Hot reload Rust code [![build badge](https://github.com/rksm/hot-lib-reloader-rs/actions/workflows/ci.yml/badge.svg)](https://github.com/rksm/hot-lib-reloader-rs/actions/workflows/ci.yml) Stars:`229`.
* [ptags](https://github.com/dalance/ptags) — A parallel universal-ctags wrapper for git repository [![Build Status](https://api.travis-ci.org/dalance/ptags.svg?branch=master)](https://travis-ci.org/dalance/ptags) Stars:`105`.
* [git-journal](https://github.com/saschagrunert/git-journal/) — The Git Commit Message and Changelog Generation Framework [![build badge](https://api.travis-ci.org/saschagrunert/git-journal.svg?branch=master)](https://travis-ci.org/saschagrunert/git-journal)
* [delta](https://crates.io/crates/git-delta) — A syntax-highlighter for git and diff output[![build badge](https://github.com/dandavison/delta/workflows/Continuous%20Integration/badge.svg)](https://github.com/dandavison/delta//actions)
* [clippy](https://crates.io/crates/clippy) — Rust lints [![build badge](https://api.travis-ci.com/rust-lang/rust-clippy.svg?branch=master)](https://travis-ci.org/rust-lang/rust-clippy)

### Build system

  * [dtolnay/cargo-expand](https://github.com/dtolnay/cargo-expand) — Expand macros in your source code Stars:`1.7K`.
  * [cargo-generate](https://github.com/cargo-generate/cargo-generate) A generator of a rust project by leveraging a pre-existing git repository as a template. Stars:`1.1K`.
  * [cargo-all-features](https://github.com/frewsxcv/cargo-all-features) - A configurable subcommand to simplify testing, building and much more for all combinations of features [![CI](https://github.com/frewsxcv/cargo-all-features/actions/workflows/ci.yml/badge.svg)](https://github.com/frewsxcv/cargo-all-features/actions/workflows/ci.yml) Stars:`82`.
  * [cargo-count](https://crates.io/crates/cargo-count) — lists source code counts and details about cargo projects, including unsafe statistics [![build badge](https://api.travis-ci.org/kbknapp/cargo-count.svg?branch=master)](https://travis-ci.org/kbknapp/cargo-count)
  * [cargo-info](https://crates.io/crates/cargo-info) — queries crates.io for crates details from command line [![build badge](https://api.travis-ci.org/imp/cargo-info.svg?branch=master)](https://travis-ci.org/imp/cargo-info)
  * [cargo-check](https://crates.io/crates/cargo-check) — A wrapper around `cargo rustc -- -Zno-trans` which can be helpful for running a faster compile if you only need correctness checks [![build badge](https://api.travis-ci.org/rsolomo/cargo-check.svg?branch=master)](https://travis-ci.org/rsolomo/cargo-check)
* [Cargo](https://crates.io/) — the Rust package manager
  * [cargo-deb](https://crates.io/crates/cargo-deb) — Generates binary Debian packages [![build badge](https://api.travis-ci.org/mmstick/cargo-deb.svg?branch=master)](https://travis-ci.org/mmstick/cargo-deb)
  * [cargo-deps](https://crates.io/crates/cargo-deps) — build dependency graphs of Rust projects [![build badge](https://api.travis-ci.com/m-cat/cargo-deps.svg?branch=master)](https://travis-ci.org/m-cat/cargo-deps)
  * [cargo-do](https://crates.io/crates/cargo-do) — run multiple cargo commands in a row [![build badge](https://api.travis-ci.org/pwoolcoc/cargo-do.svg?branch=master)](https://travis-ci.org/pwoolcoc/cargo-do)
  * [cargo-ebuild](https://crates.io/crates/cargo-ebuild) — cargo extension that can generate ebuilds using the in-tree eclasses [![build badge](https://api.travis-ci.org/cardoe/cargo-ebuild.svg?branch=master)](https://travis-ci.org/cardoe/cargo-ebuild)
  * [cargo-edit](https://crates.io/crates/cargo-edit) — allows you to add and list dependencies by reading/writing to your Cargo.toml file from the command line [![build badge](https://api.travis-ci.org/killercup/cargo-edit.svg?branch=master)](https://travis-ci.org/killercup/cargo-edit)
  * [cargo-bitbake](https://crates.io/crates/cargo-bitbake) — A cargo extension that can generate BitBake recipes utilizing the classes from meta-rust [![build badge](https://api.travis-ci.org/cardoe/cargo-bitbake.svg?branch=master)](https://travis-ci.org/cardoe/cargo-bitbake)
  * [cargo-graph](https://crates.io/crates/cargo-graph) — updated fork of `cargo-dot` with additional features. Unmaintained, see `cargo-deps` [![build badge](https://api.travis-ci.org/kbknapp/cargo-graph.svg?branch=master)](https://travis-ci.org/kbknapp/cargo-graph)
  * [cargo-cache](https://crates.io/crates/cargo-cache) — inspect/manage/clean your cargo cache (`~/.cargo/`/`${CARGO_HOME}`), print sizes etc [![Build Status](https://github.com/matthiaskrgr/cargo-cache/workflows/ci/badge.svg?branch=master)](https://github.com/matthiaskrgr/cargo-cache/actions)
  * [cargo-license](https://crates.io/crates/cargo-license) — A cargo subcommand to quickly view the licenses of all dependencies. [![build badge](https://api.travis-ci.org/onur/cargo-license.svg?branch=master)](https://travis-ci.org/onur/cargo-license)
  * [cargo-make](https://crates.io/crates/cargo-make) — Rust task runner and build tool. [![build badge](https://github.com/sagiegurari/cargo-make/workflows/CI/badge.svg?branch=master)](https://github.com/sagiegurari/cargo-make/actions)
  * [cargo-modules](https://crates.io/crates/cargo-modules) — A cargo plugin for showing a tree-like overview of a crate's modules. [![build badge](https://api.travis-ci.org/regexident/cargo-modules.svg?branch=master)](https://travis-ci.org/regexident/cargo-modules)
  * [cargo-multi](https://crates.io/crates/cargo-multi) — runs specified cargo command on multiple crates [![build badge](https://api.travis-ci.org/imp/cargo-multi.svg?branch=master)](https://travis-ci.org/imp/cargo-multi)
  * [cargo-outdated](https://crates.io/crates/cargo-outdated) — displays when newer versions of Rust dependencies are available, or out of date [![build badge](https://api.travis-ci.org/kbknapp/cargo-outdated.svg?branch=master)](https://travis-ci.org/kbknapp/cargo-outdated)
  * [cargo-release](https://crates.io/crates/cargo-release) — tool for releasing git-managed cargo project, build, tag, publish, doc and push [![Rust](https://github.com/crate-ci/cargo-release/actions/workflows/ci.yml/badge.svg)](https://github.com/crate-ci/cargo-release/actions/workflows/rust.yml)
  * [cargo-script](https://crates.io/crates/cargo-script) — lets people quickly and easily run Rust "scripts" which can make use of Cargo's package ecosystem [![build badge](https://api.travis-ci.org/DanielKeep/cargo-script.svg?branch=master)](https://travis-ci.org/DanielKeep/cargo-script)
  * [cargo-update](https://crates.io/crates/cargo-update) — cargo subcommand for checking and applying updates to installed executables [![build badge](https://api.travis-ci.org/nabijaczleweli/cargo-update.svg?branch=master)](https://travis-ci.org/nabijaczleweli/cargo-update)
  * [cargo-watch](https://crates.io/crates/cargo-watch) — utility for cargo to compile projects when sources change [![build badge](https://api.travis-ci.org/passcod/cargo-watch.svg?branch=master)](https://travis-ci.org/passcod/cargo-watch)
  * [cargo-benchcmp](https://crates.io/crates/cargo-benchcmp) — A utility to compare Rust micro-benchmarks [![build badge](https://api.travis-ci.org/BurntSushi/cargo-benchcmp.svg?branch=master)](https://travis-ci.org/BurntSushi/cargo-benchcmp)
* CMake
* [Fleet](https://github.com/dimensionhq/fleet) [[fleet-rs](https://crates.io/crates/fleet-rs)] - The blazing fast build tool for Rust. Stars:`2.2K`.
  * [Devolutions/CMakeRust](https://github.com/Devolutions/CMakeRust) — useful for integrating a Rust library into a CMake project Stars:`134`.
  * [SiegeLord/RustCMake](https://github.com/SiegeLord/RustCMake) — an example project showing usage of CMake with Rust [![build badge](https://api.travis-ci.org/SiegeLord/RustCMake.svg?branch=master)](https://travis-ci.org/SiegeLord/RustCMake) Stars:`103`.
* Github actions
  * [peaceiris/actions-mdbook](https://github.com/peaceiris/actions-mdbook) — GitHub Actions for mdBook Stars:`220`.
  * [icepuma/rust-action](https://github.com/icepuma/rust-action) — rust github action Stars:`73`.

### Debugging

* GDB
  * [gdbgui](https://github.com/cs01/gdbgui) — Browser based frontend for gdb to debug C, C++, Rust, and go. [![build badge](https://api.travis-ci.org/cs01/gdbgui.svg?branch=master)](https://travis-ci.org/cs01/gdbgui) Stars:`8.9K`.
* LLDB
  * [CodeLLDB](https://marketplace.visualstudio.com/items?itemName=vadimcn.vscode-lldb) — A LLDB extension for [Visual Studio Code](https://code.visualstudio.com/).

### Deployment

* Docker
  * [emk/rust-musl-builder](https://github.com/emk/rust-musl-builder) — Docker images for compiling static Rust binaries using musl-libc and musl-gcc, with static versions of useful C libraries Stars:`1.3K`.
  * [LukeMathWalker/cargo-chef](https://github.com/LukeMathWalker/cargo-chef) - A tool and pre-built images for caching compiling remote dependencies between Docker builds. Stars:`781`.
  * [messense/rust-musl-cross](https://github.com/messense/rust-musl-cross) — Docker images for compiling static Rust binaries using musl-cross [![build badge](https://api.travis-ci.org/messense/rust-musl-cross.svg?branch=master)](https://travis-ci.org/messense/rust-musl-cross) Stars:`342`.
  * [rust-lang-nursery/docker-rust](https://github.com/rust-lang/docker-rust) — the official Rust Docker image Stars:`312`.
  * [kpcyrd/mini-docker-rust](https://github.com/kpcyrd/mini-docker-rust) — An example project for very small rust docker images [![build badge](https://api.travis-ci.org/kpcyrd/mini-docker-rust.svg?branch=master)](https://travis-ci.org/kpcyrd/mini-docker-rust) Stars:`148`.
  * [liuchong/docker-rustup](https://github.com/liuchong/docker-rustup) — A multiple version (with musl tools) Rust Docker image Stars:`80`.
* Github Pages
  * [wasm-template-rust](https://github.com/sn99/wasm-template-rust) — A wasm template for Rust to publish to gh-pages without npm-deploy Stars:`100`.
* Heroku
  * [emk/heroku-buildpack-rust](https://github.com/emk/heroku-buildpack-rust) — A buildpack for Rust applications on Heroku Stars:`501`.

### Embedded

[Rust Embedded](https://rust-embedded.org/) focuses on improving the end-to-end experience of using Rust in resource-constrained environments and non-traditional platforms. See [awesome-embedded-rust](https://github.com/rust-embedded/awesome-embedded-rust) for a curated, and more extended list of embedded Rust resources.

* Arduino
  * [avr-rust/ruduino](https://github.com/avr-rust/ruduino) Reusable components for the Arduino Uno. Stars:`599`.
* Cross compiling
  * [japaric/rust-cross](https://github.com/japaric/rust-cross) — everything you need to know about cross compiling Rust programs [![build badge](https://api.travis-ci.org/japaric/rust-cross.svg?branch=master)](https://travis-ci.org/japaric/rust-cross) Stars:`2.2K`.
  * [japaric/xargo](https://github.com/japaric/xargo) — effortless cross compilation of Rust programs to custom bare-metal targets like ARM Cortex-M [![build badge](https://api.travis-ci.org/japaric/xargo.svg?branch=master)](https://travis-ci.org/japaric/xargo) Stars:`986`.
* Espressif
  * [esp-rs](https://github.com/esp-rs) home to a number of community projects enabling the use of the Rust programming language on various SoCs and modules produced by Espressif Systems.
* nRF
  * [nrf-rs/nrf-hal](https://github.com/nrf-rs/nrf-hal) — A Rust HAL for the nRF family of devices Stars:`334`.
[![build badge](https://api.travis-ci.org/nrf-rs/nrf-hal.svg?branch=master)](https://travis-ci.org/nrf-rs/nrf-hal)

### FFI

See also [Foreign Function Interface](https://doc.rust-lang.org/book/first-edition/ffi.html),  [The Rust FFI Omnibus](http://jakegoulding.com/rust-ffi-omnibus/) (a collection of examples of using code written in Rust from other languages) and [FFI examples written in Rust](https://github.com/alexcrichton/rust-ffi-examples).

* C
  * [rlhunt/cbindgen](https://github.com/eqrion/cbindgen) — generates C header files from Rust source files. Used in Gecko for WebRender [![build badge](https://api.travis-ci.org/eqrion/cbindgen.svg?branch=master)](https://travis-ci.org/eqrion/cbindgen) Stars:`1.6K`.
  * [Sean1708/rusty-cheddar](https://github.com/Sean1708/rusty-cheddar) — generates C header files from Rust source files [![build badge](https://api.travis-ci.org/Sean1708/rusty-cheddar.svg?branch=master)](https://travis-ci.org/Sean1708/rusty-cheddar) Stars:`190`.
* C++
  * [dtolnay/cxx](https://github.com/dtolnay/cxx) — Safe interop between Rust and C++ [![build badge](https://img.shields.io/badge/github-dtolnay/cxx-8da0cb?style=for-the-badge&labelColor=555555&logo=github)](https://github.com/dtolnay/cxx) Stars:`4.3K`.
  * [rust-lang/rust-bindgen](https://github.com/rust-lang/rust-bindgen) — A Rust bindings generator Stars:`3.1K`.
  * [rust-cpp](https://crates.io/crates/cpp) - Embed C++ code directly in Rust. [![Build Status](https://api.travis-ci.org/mystor/rust-cpp.svg?branch=master)](https://travis-ci.org/mystor/rust-cpp) [![Build status](https://ci.appveyor.com/api/projects/status/uu76vmcrwnjqra0u/branch/master?svg=true)](https://ci.appveyor.com/project/mystor/rust-cpp/branch/master)
* Erlang
  * [rusterlium/rustler](https://github.com/rusterlium/rustler) — safe Rust bridge for creating Erlang NIF functions [![build badge](https://api.travis-ci.org/rusterlium/rustler.svg?branch=master)](https://travis-ci.org/rusterlium/rustler) Stars:`3.5K`.
* Java
  * [drrb/java-rust-example](https://github.com/drrb/java-rust-example) — use Rust from Java [![build badge](https://api.travis-ci.org/drrb/java-rust-example.svg?branch=master)](https://travis-ci.org/drrb/java-rust-example) Stars:`314`.
  * [bennettanderson/rjni](https://github.com/benanders/rjni) — use Java from Rust Stars:`65`.
  * [j4rs](https://crates.io/crates/j4rs) — use Java from Rust [![build badge](https://api.travis-ci.org/astonbitecode/j4rs.svg?branch=master)](https://travis-ci.org/astonbitecode/j4rs)
  * [jni](https://crates.io/crates/jni) — use Rust from Java [![build badge](https://api.travis-ci.org/jni-rs/jni-rs.svg?branch=master)](https://travis-ci.org/jni-rs/jni-rs)
  * [jni-sys](https://crates.io/crates/jni-sys) — Rust definitions corresponding to jni.h [![build badge](https://api.travis-ci.org/sfackler/rust-jni-sys.svg?branch=master)](https://travis-ci.org/sfackler/rust-jni-sys)
  * [rucaja](https://crates.io/crates/rucaja) — use Java from Rust [![build badge](https://api.travis-ci.org/kud1ing/rucaja.svg?branch=master)](https://travis-ci.org/kud1ing/rucaja)
* Lua
  * [tomaka/hlua](https://github.com/tomaka/hlua) — Rust library to interface with Lua [![build badge](https://api.travis-ci.org/tomaka/hlua.svg?branch=master)](https://travis-ci.org/tomaka/hlua) Stars:`489`.
  * [jcmoyer/rust-lua53](https://github.com/jcmoyer/rust-lua53) — Lua 5.3 bindings for Rust [![build badge](https://api.travis-ci.org/jcmoyer/rust-lua53.svg?branch=master)](https://travis-ci.org/jcmoyer/rust-lua53) Stars:`150`.
  * [lilyball/rust-lua](https://github.com/lilyball/rust-lua) — Safe Rust bindings to Lua 5.1 [![build badge](https://api.travis-ci.org/lilyball/rust-lua.svg?branch=master)](https://travis-ci.org/lilyball/rust-lua) Stars:`121`.
  * [tickbh/td_rlua](https://github.com/tickbh/td_rlua) [[td_rlua](https://crates.io/crates/td_rlua)] — Zero-cost high-level lua 5.3 wrapper for Rust [![build badge](https://api.travis-ci.org/tickbh/td_rlua.svg?branch=master)](https://travis-ci.org/tickbh/td_rlua) Stars:`47`.
* mruby
  * [anima-engine/mrusty](https://github.com/anima-engine/mrusty) — mruby safe bindings for Rust [![build badge](https://api.travis-ci.org/anima-engine/mrusty.svg?branch=master)](https://travis-ci.org/anima-engine/mrusty) Stars:`200`.
* Node.js
  * [neon-bindings/neon](https://github.com/neon-bindings/neon) — Rust bindings for writing safe and fast native Node.js modules [![build badge](https://api.travis-ci.org/neon-bindings/neon.svg?branch=master)](https://travis-ci.org/neon-bindings/neon) Stars:`6.9K`.
  * [infinyon/node-bindgen](https://github.com/infinyon/node-bindgen) - Easy way to generate nodejs module using Rust Stars:`335`.
* Objective-C
  * [SSheldon/rust-objc](https://github.com/SSheldon/rust-objc) — Objective-C Runtime bindings and wrapper for Rust Stars:`333`.
* Python
  * [PyO3/PyO3](https://github.com/PyO3/PyO3) — Rust bindings for the Python interpreter [![build badge](https://api.travis-ci.org/PyO3/pyo3.svg?branch=master)](https://travis-ci.org/PyO3/pyo3) Stars:`6.8K`.
  * [dgrunwald/rust-cpython](https://github.com/dgrunwald/rust-cpython) — Python bindings [![build badge](https://api.travis-ci.org/dgrunwald/rust-cpython.svg?branch=master)](https://travis-ci.org/dgrunwald/rust-cpython) Stars:`1.7K`.
  * [getsentry/milksnake](https://github.com/getsentry/milksnake) — extension for python setuptools that allows you to distribute dynamic linked libraries in Python wheels in the most portable way imaginable. Stars:`746`.
* Ruby
  * [tildeio/helix](https://github.com/tildeio/helix) — write Ruby classes in Rust [![build badge](https://api.travis-ci.org/tildeio/helix.svg?branch=master)](https://travis-ci.org/tildeio/helix) Stars:`2.0K`.
  * [d-unseductable/ruru](https://github.com/d-unseductable/ruru) — native Ruby extensions written in Rust [![build badge](https://api.travis-ci.org/d-unseductable/ruru.svg?branch=master)](https://travis-ci.org/d-unseductable/ruru) Stars:`807`.
  * [danielpclark/rutie](https://github.com/danielpclark/rutie) — native Ruby extensions written in Rust and vice versa [![Build Status](https://api.travis-ci.org/danielpclark/rutie.svg?branch=master)](https://travis-ci.org/danielpclark/rutie) Stars:`691`.
* Web Assembly
  * [rustwasm/wasm-bindgen](https://github.com/rustwasm/wasm-bindgen) — A project for facilitating high-level interactions between wasm modules and JS. [![build badge](https://api.travis-ci.com/rustwasm/wasm-bindgen.svg?branch=master)](https://travis-ci.org/rustwasm/wasm-bindgen) Stars:`5.7K`.
  * [rustwasm/wasm-pack](https://github.com/rustwasm/wasm-pack) — :package: :sparkles: pack up the wasm and publish it to npm! [![build badge](https://api.travis-ci.com/rustwasm/wasm-pack.svg?branch=master)](https://travis-ci.org/rustwasm/wasm-pack) Stars:`4.6K`.
  * [rhysd/wain](https://github.com/rhysd/wain) - wain: WebAssembly INterpreter from scratch in Safe Rust with zero dependency [![build badge](https://github.com/rhysd/wain/workflows/CI/badge.svg?branch=master&event=push)](https://github.com/rhysd/wain/actions?query=workflow%3ACI+branch%3Amaster+event%3Apush) Stars:`315`.

### Formatters

* [rustfmt](https://github.com/rust-lang/rustfmt) — Rust code formatter maintained by the Rust team and included in cargo [![build badge](https://api.travis-ci.com/rust-lang/rustfmt.svg?branch=master)](https://app.travis-ci.com/github/rust-lang/rustfmt) Stars:`4.7K`.
* [dprint](https://github.com/dprint/dprint) — A pluggable and configurable code formatting platform [![build badge](https://github.com/dprint/dprint/workflows/CI/badge.svg)](https://github.com/dprint/dprint/actions?query=workflow%3ACI) Stars:`1.6K`.
* [Prettier Rust](https://github.com/jinxdash/prettier-plugin-rust) — An opinionated Rust code formatter that autofixes bad syntax ([Prettier](https://prettier.io/) community plugin) Stars:`79`.

### IDEs

See also [Are we (I)DE yet?](https://areweideyet.com/) and [Rust Tools](https://www.rust-lang.org/tools).

    * [intellij-rust/intellij-rust](https://github.com/intellij-rust/intellij-rust) — [![build badge](https://api.travis-ci.org/intellij-rust/intellij-rust.svg?branch=master)](https://travis-ci.org/intellij-rust/intellij-rust) Stars:`4.2K`.
    * [autozimu/LanguageClient-neovim](https://github.com/autozimu/LanguageClient-neovim) — [LSP](https://microsoft.github.io/language-server-protocol/) client. Implemented in Rust and supports rls out of the box. Stars:`3.4K`.
    * [rust.vim](https://github.com/rust-lang/rust.vim) — provides file detection, syntax highlighting, formatting, Syntastic integration, and more. Stars:`3.3K`.
    * [rust-mode](https://github.com/rust-lang/rust-mode) — Rust Major Mode Stars:`895`.
    * [rust-lang/rust-enhanced](https://github.com/rust-lang/rust-enhanced) — official Rust package Stars:`693`.
    * [vim-racer](https://github.com/racer-rust/vim-racer) — allows vim to use [Racer](https://github.com/racer-rust/racer) for Rust code completion and navigation. Stars:`623`.
    * [rustic](https://github.com/brotzeit/rustic) - Rust development environment for Emacs [![build badge](https://github.com/brotzeit/rustic/workflows/CI/badge.svg)](https://github.com/brotzeit/rustic/actions?query=workflow%3ACI) Stars:`590`.
    * [emacs-racer](https://github.com/racer-rust/emacs-racer) — Autocompletion (see also [company](https://company-mode.github.io) and [auto-complete](https://github.com/auto-complete/auto-complete)) Stars:`396`.
    * [crates.nvim](https://github.com/Saecki/crates.nvim) - plugin that helps to managing crates.io dependencies. Stars:`320`.
    * [rust-lang/atom-ide-rust](https://github.com/rust-lang/atom-ide-rust) — Rust IDE support for Atom, powered by the Rust Language Server (RLS) [![Build Status](https://api.travis-ci.com/rust-lang/atom-ide-rust.svg?branch=master)](https://app.travis-ci.com/grust-lang/atom-ide-rust) Stars:`236`.
    * [Eclipse Corrosion](https://github.com/eclipse/corrosion) Stars:`190`.
  * [Ride](https://github.com/madeso/ride) — [![build badge](https://api.travis-ci.org/madeso/ride.svg?branch=master)](https://travis-ci.org/madeso/ride) Stars:`169`.
    * [flycheck-rust](https://github.com/flycheck/flycheck-rust) — Rust support for [Flycheck](https://github.com/flycheck/flycheck) Stars:`112`.
  * [IntelliJ](https://www.jetbrains.com/idea/)
  * [Kakoune](http://kakoune.org/)
    * [kak-lsp/kak-lsp](https://github.com/kak-lsp/kak-lsp/) — [LSP](https://microsoft.github.io/language-server-protocol/) client. Implemented in Rust and supports rls out of the box.
  * [Atom](https://atom.io/)
  * [Sublime Text](https://www.sublimetext.com/)
  * [gnome-builder](https://wiki.gnome.org/Apps/Builder) native support for rust and cargo since Version 3.22.2
  * [Vim](https://vim.sourceforge.io/) — the ubiquitous text editor
  * [gitpod.io](https://gitpod.io) — Online IDE with full Rust support based on Rust Language Server
  * [Emacs](https://www.gnu.org/software/emacs/)
  * [Eclipse](https://www.eclipse.org/)
  * Visual Studio
    * [PistonDevelopers/VisualRust](https://github.com/PistonDevelopers/VisualRust) — A Visual Studio extension for Rust [![Build status](https://ci.appveyor.com/api/projects/status/5nw5no10jj0y4p3f?svg=true)](https://ci.appveyor.com/project/vosen/visualrust) Stars:`696`.
    * [crates](https://github.com/serayuzgur/crates) — crates is an extension for crates.io dependencies. [![build badge](https://img.shields.io/vscode-marketplace/v/serayuzgur.crates.svg)](https://github.com/serayuzgur/crates) [![build badge](https://api.travis-ci.org/serayuzgur/crates.svg?branch=master)](https://travis-ci.org/serayuzgur/crates) Stars:`158`.
    * [dgriffen/rls-vs2017](https://github.com/ZoeyR/rls-vs2017) — Rust support for Visual Studio 2017 Preview [![build badge](https://ci.appveyor.com/api/projects/status/d2lxlincwninhsng?svg=true)](https://ci.appveyor.com/project/dgriffen/rls-vs2017) Stars:`111`.
  * [Visual Studio Code](https://code.visualstudio.com/)
    * [Better TOML](https://marketplace.visualstudio.com/items?itemName=bungcip.better-toml) - TOML support in vscode
    * [CodeLLDB](https://marketplace.visualstudio.com/items?itemName=vadimcn.vscode-lldb) — A LLDB extension
    * [Prettier - Code formatter (Rust)](https://marketplace.visualstudio.com/items?itemName=jinxdash.prettier-rust) — Opinionated Rust code formatter that autofixes bad syntax ([Prettier](https://prettier.io/) community plugin)
    * [rust-analyzer](https://marketplace.visualstudio.com/items?itemName=rust-lang.rust-analyzer) — An alternative rust language server to the RLS
    * [rust-lang/rls-vscode](https://marketplace.visualstudio.com/items?itemName=rust-lang.rust) — Rust support for Visual Studio Code (supports both RLS and rust-analyzer)

### Profiling

* [Bytehound](https://github.com/koute/bytehound) — A memory profiler for Linux Stars:`3.2K`.
* [bheisler/criterion.rs](https://github.com/bheisler/criterion.rs) — Statistics-driven benchmarking library for Rust [![Build Status](https://api.travis-ci.org/bheisler/criterion.rs.svg?branch=master)](https://travis-ci.org/bheisler/criterion.rs) Stars:`2.9K`.
* [ellisonch/rust-stopwatch](https://github.com/ellisonch/rust-stopwatch) — A stopwatch library [![build badge](https://api.travis-ci.org/ellisonch/rust-stopwatch.svg?branch=master)](https://travis-ci.org/ellisonch/rust-stopwatch) Stars:`75`.
* FlameGraphs
* [sharkdp/hyperfine](https://github.com/sharkdp/hyperfine) — A command-line benchmarking tool [![Version info](https://img.shields.io/crates/v/hyperfine.svg)](https://crates.io/crates/hyperfine) [![Build Status](https://api.travis-ci.org/sharkdp/hyperfine.svg?branch=master)](https://travis-ci.org/sharkdp/hyperfine) Stars:`13.3K`.
  * [llogiq/flame](https://github.com/llogiq/flame) — [![build badge](https://api.travis-ci.org/llogiq/flame.svg?branch=master)](https://travis-ci.org/llogiq/flame) Stars:`626`.
  * [mrhooray/torch](https://github.com/mrhooray/torch) — generates FlameGraphs based on DWARF Debug Info Stars:`125`.

### Services

* [deps.rs](https://github.com/deps-rs/deps.rs) — Detect outdated or insecure dependencies Stars:`363`.
* [docs.rs](https://docs.rs) — Automatic documentation generation of crates

### Static analysis

[[assert](https://crates.io/keywords/assert), [static](https://crates.io/keywords/static)]

* [facebookexperimental/MIRAI](https://github.com/facebookexperimental/mirai) — an abstract interpreter operating on Rust's mid-level intermediate representation (MIR) [![Continuous Integration](https://github.com/facebookexperimental/MIRAI/actions/workflows/rust.yml/badge.svg)](https://github.com/facebookexperimental/MIRAI/actions/workflows/rust.yml) Stars:`774`.
* [static_assertions](https://crates.io/crates/static_assertions) — Compile-time assertions to ensure that invariants are met [![Build Status](https://api.travis-ci.org/nvzqz/static-assertions-rs.svg?branch=master)](https://travis-ci.org/nvzqz/static-assertions-rs/)

### Testing

[[test](https://crates.io/keywords/test), [testing](https://crates.io/keywords/testing)]

* Code Coverage
  * [tarpaulin](https://crates.io/crates/cargo-tarpaulin) — A code coverage tool designed for Rust [![build badge](https://api.travis-ci.org/repositories/xd009642/tarpaulin.svg?branch=master)](https://travis-ci.org/xd009642/tarpaulin)
* Continuous Integration
  * [trust](https://github.com/japaric/trust) — A Travis CI and AppVeyor template to test your Rust crate on 5 architectures and publish binary releases of it for Linux, macOS and Windows Stars:`1.2K`.
* Frameworks and Runners
  * [AlKass/polish](https://github.com/AlKass/polish) — Mini Testing/Test-Driven Framework [![Build Status](https://api.travis-ci.org/AlKass/polish.svg?branch=master)](https://travis-ci.org/AlKass/polish) [![Crates Package Status](https://img.shields.io/crates/v/polish.svg)](https://crates.io/crates/polish) Stars:`47`.
  * [cargo-dinghy](https://crates.io/crates/cargo-dinghy/) - A cargo extension to simplify running library tests and benches on smartphones and other small processor devices.
  * [cucumber](https://crates.io/crates/cucumber) [![Latest Version](https://img.shields.io/crates/v/cucumber.svg)](https://crates.io/crates/cucumber) — An implementation of the Cucumber testing framework for Rust. Fully native, no external test runners or dependencies. [![Build Status](https://github.com/cucumber-rs/cucumber/workflows/CI/badge.svg?branch=master)](https://github.com/cucumber-rs/cucumber)
  * [demonstrate](https://crates.io/crates/demonstrate) — Declarative Testing Framework [![Build Status](https://github.com/aubaugh/demonstrate/workflows/Continuous%20Integration/badge.svg?branch=master)](https://github.com/aubaugh/demonstrate)
  * [rstest](https://crates.io/crates/rstest) — Fixture-based test framework for Rust [![Build Status](https://github.com/la10736/rstest/workflows/Test/badge.svg?branch=master)](https://github.com/la10736/rstest/actions)
  * [speculate](https://crates.io/crates/speculate) — An RSpec inspired minimal testing framework for Rust
* Mocking and Test Data
  * [asomers/mockall](https://github.com/asomers/mockall) [[mockall](https://crates.io/crates/mockall)] — A powerful mock object library for Rust. [![Cirrus Build Status](https://api.cirrus-ci.com/github/asomers/mockall.svg)](https://cirrus-ci.com/github/asomers/mockall) Stars:`871`.
  * [fake-rs](https://github.com/cksac/fake-rs) —  A library for generating fake data [![build badge](https://api.travis-ci.org/repositories/cksac/fake-rs.svg?branch=master)](https://travis-ci.org/cksac/fake-rs) Stars:`510`.
  * [httpmock](https://github.com/alexliesenfeld/httpmock) — HTTP mocking [![build badge](https://dev.azure.com/alexliesenfeld/httpmock/_apis/build/status/alexliesenfeld.httpmock?branchName=master)](https://dev.azure.com/alexliesenfeld/httpmock/_build/latest?definitionId=2&branchName=master) Stars:`313`.
  * [goldenfile](https://github.com/calder/rust-goldenfile) [[goldenfile](https://crates.io/crates/goldenfile)] - A library providing a simple API for goldenfile testing. [![build badge](https://api.travis-ci.org/calder/rust-goldenfile.svg?branch=master)](https://travis-ci.org/calder/rust-goldenfile) Stars:`21`.
  * [mockiato](https://crates.io/crates/mockiato) — A strict, yet friendly mocking library for Rust 2018 [![build badge](https://api.travis-ci.com/mockiato/mockiato.svg?branch=master)](https://app.travis-ci.com/github/mockiato/mockiato)
  * [mockito](https://crates.io/crates/mockito) — HTTP mocking [![build badge](https://api.travis-ci.org/lipanski/mockito.svg?branch=master)](https://travis-ci.org/lipanski/mockito)
  * [nrxus/faux](https://github.com/nrxus/faux/) [![Latest Version](https://img.shields.io/crates/v/faux.svg)](https://crates.io/crates/faux) — A library to create mocks out of structs. ![build](https://github.com/nrxus/faux/workflows/test/badge.svg?branch=master)
* Mutation Testing
  * [mutagen](https://github.com/llogiq/mutagen) [[mutagen](https://crates.io/crates/mutagen)] — A source-level mutation testing framework (nightly only) [![build badge](https://api.travis-ci.org/llogiq/mutagen.svg?branch=master)](https://travis-ci.org/llogiq/mutagen) Stars:`524`.
  * [cargo-mutants](https://github.com/sourcefrog/cargo-mutants) [[cargo-mutants](https://crates.io/crates/cargo-mutants)] - Finds inadequately tested code by injecting mutations, no source changes required. [![build badge](https://github.com/sourcefrog/cargo-mutants/actions/workflows/tests.yml/badge.svg?branch=main&event=push)](https://github.com/sourcefrog/cargo-mutants/actions/workflows/tests.yml?query=branch%3Amain) Stars:`177`.
* Property Testing and Fuzzing
  * [rust-fuzz/afl.rs](https://github.com/rust-fuzz/afl.rs) — A Rust fuzzer, using [AFL](https://lcamtuf.coredump.cx/afl/) [![build badge](https://api.travis-ci.org/rust-fuzz/afl.rs.svg?branch=master)](https://travis-ci.org/rust-fuzz/afl.rs) Stars:`1.3K`.
  * [proptest](https://crates.io/crates/proptest) — property testing framework inspired by the [Hypothesis](https://hypothesis.works/) framework for Python [![build badge](https://api.travis-ci.org/altsysrq/proptest.svg?branch=master)](https://travis-ci.org/altsysrq/proptest)
  * [quickcheck](https://crates.io/crates/quickcheck) — A Rust implementation of [QuickCheck](https://wiki.haskell.org/Introduction_to_QuickCheck1) [![build badge](https://api.travis-ci.org/BurntSushi/quickcheck.svg?branch=master)](https://travis-ci.org/BurntSushi/quickcheck)

### Transpiling

* [immunant/c2rust](https://github.com/immunant/c2rust) — C to Rust translator and cross checker built atop Clang/LLVM. [![Build Status](https://api.travis-ci.org/immunant/c2rust.svg?branch=master)](https://travis-ci.org/immunant/c2rust) Stars:`2.9K`.
* [BayesWitnesses/m2cgen](https://github.com/BayesWitnesses/m2cgen) — A CLI tool to transpile trained classic machine learning models into a native Rust code with zero dependencies. [![GitHub Actions Status](https://github.com/BayesWitnesses/m2cgen/workflows/GitHub%20Actions/badge.svg?branch=master)](https://github.com/BayesWitnesses/m2cgen/actions) Stars:`2.3K`.
* [jameysharp/corrode](https://github.com/jameysharp/corrode) — A C to Rust translator written in Haskell. Stars:`2.1K`.

## Libraries

* [perf-monitor-rs](https://github.com/larksuite/perf-monitor-rs) — A toolkit designed to be a foundation for applications to monitor their performance. [![crates.io](https://img.shields.io/crates/v/perf_monitor.svg)](https://crates.io/crates/perf_monitor) Stars:`153`.
* [Phate6660/nixinfo](https://github.com/Phate6660/nixinfo) [[crate](https://crates.io/crates/nixinfo)] — A lib crate for gathering system info such as cpu, distro, environment, kernel, etc. Stars:`35`.

### Artificial Intelligence

#### Genetic algorithms

* [Martin1887/oxigen](https://github.com/Martin1887/oxigen) — Fast, parallel, extensible and adaptable genetic algorithm library. A example using this library solves the N Queens problem for N = 255 in only few seconds and using less than 1 MB of RAM. Stars:`139`.
* [pkalivas/radiate](https://github.com/pkalivas/radiate) — A customizable parallel genetic programming engine capable of evolving solutions for supervised, unsupervised, and reinforcement learning problems. Comes with complete and customizable implementation of NEAT and Evtree. [![Build Status](https://api.travis-ci.com/pkalivas/radiate.svg?branch=master)](https://app.travis-ci.com/github/pkalivas/radiate)![Crates.io](https://img.shields.io/crates/v/radiate) Stars:`103`.
* [innoave/genevo](https://github.com/innoave/genevo) — Execute genetic algorithm (GA) simulations in a customizable and extensible way. Stars:`102`.
* [willi-kappler/darwin-rs](https://github.com/willi-kappler/darwin-rs) — Evolutionary algorithms with Rust [![Build Status](https://api.travis-ci.org/willi-kappler/darwin-rs.svg?branch=master)](https://travis-ci.org/willi-kappler/darwin-rs) Stars:`92`.
* [m-decoster/RsGenetic](https://github.com/m-decoster/RsGenetic) — Genetic Algorithm library in Rust. In maintenance mode. Stars:`72`.

#### Machine learning

See [[Machine learning](https://crates.io/keywords/machine-learning)]

See also [About Rust’s Machine Learning Community](https://medium.com/@autumn_eng/about-rust-s-machine-learning-community-4cda5ec8a790#.hvkp56j3f) and [Are we learning yet?](https://www.arewelearningyet.com).

* [huggingface/tokenizers](https://github.com/huggingface/tokenizers) - Hugging Face's tokenizers for modern NLP pipelines written in Rust (original implementation) with bindings for Python. [![Build Status](https://github.com/huggingface/tokenizers/workflows/Rust/badge.svg?branch=master)](https://github.com/huggingface/tokenizers/actions) Stars:`6.0K`.
* [autumnai/leaf](https://github.com/autumnai/leaf) — Open Machine Intelligence framework. [![Build Status](https://api.travis-ci.org/autumnai/leaf.svg?branch=master)](https://travis-ci.org/autumnai/leaf). Abandoned project. The most updated fork is [spearow/juice]( https://github.com/spearow/juice). Stars:`5.5K`.
* [tensorflow/rust](https://github.com/tensorflow/rust) — Rust language bindings for TensorFlow. [![Build Status](https://api.travis-ci.org/tensorflow/rust.svg?branch=master)](https://travis-ci.org/tensorflow/rust) Stars:`4.0K`.
* [LaurentMazare/tch-rs](https://github.com/LaurentMazare/tch-rs) — Rust language bindings for PyTorch. [![Build Status](https://api.travis-ci.org/LaurentMazare/tch-rs.svg?branch=master)](https://travis-ci.org/LaurentMazare/tch-rs) Stars:`2.1K`.
* [rust-ml/linfa](https://github.com/rust-ml/linfa) — Machine learning framework. Stars:`2.0K`.
* [maciejkula/rustlearn](https://github.com/maciejkula/rustlearn) — Machine learning crate for Rust. [![Circle CI](https://circleci.com/gh/maciejkula/rustlearn.svg?style=svg)](https://app.circleci.com/pipelines/github/maciejkula/rustlearn) Stars:`541`.
* [smartcorelib/smartcore](https://github.com/smartcorelib/smartcore) — Machine Learning Library In Rust [![Build Status](https://img.shields.io/circleci/build/github/smartcorelib/smartcore)](https://smartcorelib.org/) Stars:`377`.

### Astronomy

[[astronomy](https://crates.io/keywords/astronomy)]

* [saurvs/astro-rust](https://github.com/saurvs/astro-rust) — astronomy for Rust [![build badge](https://api.travis-ci.org/saurvs/astro-rust.svg?branch=master)](https://travis-ci.org/saurvs/astro-rust) Stars:`206`.
* [flosse/rust-sun](https://github.com/flosse/rust-sun) [[sun](https://crates.io/crates/sun)] — A rust port of the JS library suncalc [![build badge](https://api.travis-ci.org/flosse/rust-sun.svg?branch=master)](https://travis-ci.org/flosse/rust-sun) Stars:`32`.
* [fitsio](https://crates.io/crates/fitsio) — fits interface library wrapping cfitsio [![build badge](https://api.travis-ci.org/mindriot101/rust-fitsio.svg?branch=master)](https://travis-ci.org/mindriot101/rust-fitsio)

### Asynchronous

* [mio](https://github.com/tokio-rs/mio) — MIO is a lightweight IO library for Rust with a focus on adding as little overhead as possible over the OS abstractions [![build badge](https://api.travis-ci.org/tokio-rs/mio.svg?branch=master)](https://travis-ci.org/tokio-rs/mio) Stars:`5.2K`.
* [rust-lang/futures-rs](https://github.com/rust-lang/futures-rs) — Zero-cost futures in Rust [![build badge](https://api.travis-ci.com/rust-lang/futures-rs.svg?branch=master)](https://travis-ci.org/rust-lang/futures-rs) Stars:`4.6K`.
* [Xudong-Huang/may](https://github.com/Xudong-Huang/may) — rust stackful coroutine library [![build badge](https://api.travis-ci.org/Xudong-Huang/may.svg?branch=master)](https://travis-ci.org/Xudong-Huang/may) Stars:`1.2K`.
* [zonyitoo/coio-rs](https://github.com/zonyitoo/coio-rs) — A coroutine I/O library with a working-stealing scheduler [![build badge](https://api.travis-ci.org/zonyitoo/coio-rs.svg?branch=master)](https://travis-ci.org/zonyitoo/coio-rs) Stars:`453`.
* [dpc/mioco](https://github.com/dpc/mioco) — Scalable, coroutine-based, asynchronous IO handling library [![build badge](https://api.travis-ci.org/dpc/mioco.svg?branch=master)](https://travis-ci.org/dpc/mioco) Stars:`135`.
* [TeaEntityLab/fpRust](https://github.com/TeaEntityLab/fpRust) — Monad/MonadIO, Handler, Coroutine/doNotation, Functional Programming features for Rust [![build badge](https://api.travis-ci.org/TeaEntityLab/fpRust.svg?branch=master)](https://travis-ci.org/TeaEntityLab/fpRust) Stars:`95`.
* [async-std](https://async.rs/) [[async-std]](https://crates.io/crates/async-std) - Async version of the Rust standard library [![CI](https://github.com/async-rs/async-std/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/async-rs/async-std/actions/workflows/ci.yml)

### Audio and Music

[[audio](https://crates.io/keywords/audio)]

  * [RustAudio/cpal](https://github.com/RustAudio/cpal) - Low-level cross-platform audio I/O library in pure Rust. [![Actions Status](https://github.com/RustAudio/cpal/workflows/cpal/badge.svg?branch=master)](https://github.com/RustAudio/cpal/actions) Stars:`1.7K`.
  * [RustAudio/rodio](https://github.com/RustAudio/rodio) — A Rust audio playback library [![Build Status](https://api.travis-ci.org/RustAudio/rodio.svg?branch=master)](https://travis-ci.org/RustAudio/rodio) Stars:`1.1K`.
* [pdeljanov/Symphonia](https://github.com/pdeljanov/Symphonia) — A pure Rust audio decoding and media demuxing library supporting AAC, FLAC, MP3, MP4, OGG, Vorbis, and WAV. Stars:`950`.
* [ozankasikci/rust-music-theory](https://github.com/ozankasikci/rust-music-theory) — A Rust music theory library [![Build Status](https://api.travis-ci.com/ozankasikci/rust-music-theory.svg?branch=master)](https://travis-ci.org/ozankasikci/rust-music-theory) Stars:`531`.
  * [RustAudio/rust-portaudio](https://github.com/RustAudio/rust-portaudio) — PortAudio bindings [![build badge](https://api.travis-ci.org/RustAudio/rust-portaudio.svg?branch=master)](https://travis-ci.org/RustAudio/rust-portaudio) Stars:`325`.
* [jhasse/ears](https://github.com/jhasse/ears) — A simple library to play Sounds and Musics, on top of OpenAL and libsndfile [![build badge](https://api.travis-ci.org/jhasse/ears.svg?branch=master)](https://travis-ci.org/jhasse/ears) Stars:`81`.
* [musitdev/portmidi-rs](https://github.com/musitdev/portmidi-rs) — [PortMidi](https://portmedia.sourceforge.net/portmidi/) bindings [![build badge](https://api.travis-ci.org/musitdev/portmidi-rs.svg?branch=master)](https://travis-ci.org/musitdev/portmidi-rs) Stars:`68`.
* [Serial-ATA/lofty-rs](https://github.com/Serial-ATA/lofty-rs) [[lofty](https://crates.io/crates/lofty)] — A library for reading and editing the metadata of various audio formats [![build badge](https://github.com/Serial-ATA/lofty-rs/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/Serial-ATA/lofty-rs/actions) Stars:`56`.
* [insomnimus/nodi](https://github.com/insomnimus/nodi) [[nodi](https://crates.io/crates/nodi)] — A library for playback and abstraction of MIDI files. [![build badge](https://github.com/insomnimus/nodi/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/insomnimus/nodi/actions) Stars:`3`.
* [hound](https://crates.io/crates/hound) — A WAV encoding and decoding library [![build badge](https://api.travis-ci.org/ruuda/hound.svg?branch=master)](https://travis-ci.org/ruuda/hound)
* [RustAudio](https://github.com/RustAudio)

### Authentication

* [Keats/jsonwebtoken](https://github.com/Keats/jsonwebtoken) — [JSON Web Token](https://en.wikipedia.org/wiki/JSON_Web_Token) lib in rust  [![Build Status](https://api.travis-ci.org/Keats/jsonwebtoken.svg?branch=master)](https://travis-ci.org/Keats/jsonwebtoken) Stars:`1.0K`.
* [oauth2](https://github.com/ramosbugs/oauth2-rs) — Extensible, strongly-typed Rust OAuth2 client library [![Build Status](https://api.travis-ci.org/ramosbugs/oauth2-rs.svg?branch=main)](https://travis-ci.org/ramosbugs/oauth2-rs) Stars:`571`.
* [oxide-auth](https://github.com/HeroicKatora/oxide-auth) — A OAuth2 server library, for use in combination with actix or other frontends, featuring a set of configurable and pluggable backends [![Build Status](https://api.cirrus-ci.com/github/HeroicKatora/oxide-auth.svg?branch=master)](https://cirrus-ci.com/github/HeroicKatora/oxide-auth) Stars:`455`.
* [yup-oauth2](https://github.com/dermesser/yup-oauth2) — An oauth2 client implementation providing the Device, Installed and Service Account flows [![Build Status](https://api.travis-ci.org/dermesser/yup-oauth2.svg?branch=master)](https://travis-ci.org/dermesser/yup-oauth2) Stars:`166`.
* [sgrust01/jwtvault](https://github.com/sgrust01/jwtvault) — Async library to manage and orchestrate JWT workflow  [![Build Status](https://api.travis-ci.org/sgrust01/jwtvault.svg?branch=master)](https://travis-ci.org/sgrust01/jwtvault) Stars:`64`.
* [constantoine/totp-rs](https://github.com/constantoine/totp-rs) [[totp-rs](https://crates.io/crates/totp-rs)] — 2fa library to generate and verify TOTP-based tokens ![Build Status](https://github.com/constantoine/totp-rs/workflows/Rust/badge.svg) Stars:`39`.

### Automotive

* [mbr/socketcan](https://github.com/socketcan-rs/socketcan-rs) [[socketcan](https://crates.io/crates/socketcan)] — Linux SocketCAN library Stars:`52`.
* [marcelbuesing/can-dbc](https://github.com/marcelbuesing/can-dbc) [[can-dbc](https://crates.io/crates/can-dbc)] — A parser for the DBC format [![build badge](https://api.travis-ci.org/marcelbuesing/can-dbc.svg?branch=dev)](https://travis-ci.org/marcelbuesing/can-dbc) Stars:`34`.
* [oefd/tokio-socketcan](https://github.com/oefd/tokio-socketcan) [[tokio-socketcan]](https://crates.io/crates/tokio-socketcan)] — Linux SocketCAN support for tokio based on the socketcan crate Stars:`28`.
* [Sensirion/lin-bus](https://github.com/Sensirion/lin-bus-rs) [[lin-bus](https://crates.io/crates/lin-bus)] — LIN bus driver traits and protocol implementation [![build badge](https://circleci.com/gh/Sensirion/lin-bus-rs.svg?style=svg)](https://app.circleci.com/pipelines/github/Sensirion/lin-bus-rs) Stars:`10`.
* [marcelbuesing/tokio-socketcan-bcm](https://github.com/marcelbuesing/tokio-socketcan-bcm) [[tokio-socketcan-bcm](https://crates.io/crates/tokio-socketcan-bcm)] — Linux SocketCAN BCM support for tokio [![build badge](https://api.travis-ci.org/marcelbuesing/tokio-socketcan-bcm.svg?branch=master)](https://travis-ci.org/marcelbuesing/tokio-socketcan-bcm) Stars:`3`.

### Bioinformatics

* [Rust-Bio](https://github.com/rust-bio) — bioinformatics libraries in Rust.

### Caching

* [jaemk/cached](https://github.com/jaemk/cached) — Simple function caching/memoization Stars:`954`.
* [al8n/stretto](https://github.com/al8n/stretto) - A high performance thread-safe memory-bound Rust cache [![build badge](https://github.com/al8n/stretto/actions/workflows/ci.yml/badge.svg)](https://github.com/al8n/stretto/actions/workflows/ci.yml) Stars:`302`.
* [aisk/rust-memcache](https://github.com/aisk/rust-memcache) — Memcached client library [![build badge](https://api.travis-ci.org/aisk/rust-memcache.svg?branch=master)](https://travis-ci.org/aisk/rust-memcache) Stars:`102`.
* [mozilla/sccache](https://github.com/mozilla/sccache/) - Shared Compilation Cache, great for Rust compilation [![build badge](https://api.travis-ci.org/mozilla/sccache.svg?branch=master)](https://travis-ci.org/mozilla/sccache)

### Cloud

* AWS [[aws](https://crates.io/keywords/aws)]
  * [rusoto/rusoto](https://github.com/rusoto/rusoto) — [![build badge](https://api.travis-ci.org/rusoto/rusoto.svg?branch=master)](https://travis-ci.org/rusoto/rusoto) Stars:`2.6K`.
  * [awslabs/aws-lambda-rust-runtime](https://github.com/awslabs/aws-lambda-rust-runtime) [[lambda_runtime](https://crates.io/crates/lambda_runtime)] — A Rust runtime for AWS Lambda  [![build badge](https://github.com/awslabs/aws-lambda-rust-runtime/workflows/Rust/badge.svg)](https://github.com/awslabs/aws-lambda-rust-runtime/actions) Stars:`2.3K`.
  * [awslabs/aws-sdk-rust](https://github.com/awslabs/aws-sdk-rust) - The new AWS SDK for Rust Stars:`1.9K`.
* Load Balancer
  * [Convey](https://github.com/bparli/convey) - Layer 4 Load Balancer with dynamic configuration loading. Stars:`275`.
* Multi Cloud
  * [Qovery/engine](https://github.com/Qovery/engine) - Abstraction layer library that turns easy application deployment on Cloud providers in just a few minutes Stars:`1.8K`.

### Command-line

* Argument parsing
  * [clap-rs](https://github.com/clap-rs/clap) [[clap](https://crates.io/crates/clap)] — A simple to use, full featured command-line argument parser [![build badge](https://api.travis-ci.org/clap-rs/clap.svg?branch=master)](https://travis-ci.org/clap-rs/clap) Stars:`9.9K`.
  * [TeXitoi/structopt](https://github.com/TeXitoi/structopt) [[structopt](https://crates.io/crates/structopt)] — parse command line argument by defining a struct [![build badge](https://api.travis-ci.org/TeXitoi/structopt.svg?branch=master)](https://travis-ci.org/TeXitoi/structopt) Stars:`2.5K`.
  * [google/argh](https://github.com/google/argh) [[argh](https://crates.io/crates/argh)] — An opinionated Derive-based argument parser optimized for code size [![build badge](https://github.com/google/argh/workflows/Argh/badge.svg?branch=master)](https://github.com/google/argh/actions) Stars:`1.2K`.
  * [docopt/docopt.rs](https://github.com/docopt/docopt.rs) [[docopt](https://crates.io/crates/docopt)] — A Rust implementation of [DocOpt](http://docopt.org) [![build badge](https://api.travis-ci.org/docopt/docopt.rs.svg?branch=master)](https://travis-ci.org/docopt/docopt.rs) Stars:`742`.
  * [killercup/quicli](https://github.com/killercup/quicli) [[quicli](https://crates.io/crates/quicli)] — quickly build cool CLI apps in Rust [![build badge](https://api.travis-ci.org/killercup/quicli.svg?branch=master)](https://travis-ci.org/killercup/quicli) Stars:`537`.
  * [ksk001100/seahorse](https://github.com/ksk001100/seahorse) [[seahorse](https://crates.io/crates/seahorse)] — A minimal CLI framework written in Rust [![Build status](https://github.com/ksk001100/seahorse/workflows/CI/badge.svg?branch=master)](https://github.com/ksk001100/seahorse/actions) Stars:`213`.
  * [cliparser](https://crates.io/crates/cliparser) — Simple command line parser. [![build badge](https://github.com/sagiegurari/cliparser/workflows/CI/badge.svg?branch=master)](https://github.com/sagiegurari/cliparser/actions)
* Data visualization
  * [zhiburt/tabled](https://github.com/zhiburt/tabled) [[tabled](https://crates.io/crates/tabled)] — An easy to use library for pretty print tables of Rust structs and enums.  [![Build Status](https://github.com/zhiburt/tabled/actions/workflows/ci.yml/badge.svg)](https://github.com/zhiburt/tabled/actions) Stars:`1.2K`.
  * [nukesor/comfy-table](https://github.com/nukesor/comfy-table) [[comfy-table](https://crates.io/crates/comfy-table)] — Beautiful dynamic tables for your cli tools. [![Build status](https://github.com/Nukesor/comfy-table/workflows/Tests/badge.svg?branch=master)](https://github.com/nukesor/comfy-table/actions) Stars:`486`.
* Human-centered design
  * [rust-cli/human-panic](https://github.com/rust-cli/human-panic) [[human-panic](https://crates.io/crates/human-panic)] — panic messages for humans [![build badge](https://api.travis-ci.org/rust-cli/human-panic.svg?branch=master)](https://travis-ci.org/rust-cli/human-panic) Stars:`983`.
* Line editor
  * [kkawakam/rustyline](https://github.com/kkawakam/rustyline) [[rustyline](https://crates.io/crates/rustyline)] — readline implementation in Rust [![build badge](https://api.travis-ci.org/kkawakam/rustyline.svg?branch=master)](https://travis-ci.org/kkawakam/rustyline) Stars:`1.1K`.
  * [murarth/linefeed](https://github.com/murarth/linefeed) [[linefeed](https://crates.io/crates/linefeed)] — Configurable, extensible, interactive line reader [![build badge](https://api.travis-ci.org/murarth/linefeed.svg?branch=master)](https://travis-ci.org/murarth/linefeed) Stars:`176`.
  * [MovingtoMars/liner](https://github.com/MovingtoMars/liner) [[liner](https://crates.io/crates/liner)] — A library offering readline-like functionality [![build badge](https://api.travis-ci.org/MovingtoMars/liner.svg?branch=master)](https://travis-ci.org/MovingtoMars/liner) Stars:`70`.
  * [srijs/rust-copperline](https://github.com/srijs/rust-copperline) [[copperline](https://crates.io/crates/copperline)] — pure-Rust command line editing library Stars:`25`.
* Other
  * [mgrachev/update-informer](https://github.com/mgrachev/update-informer) [[update-informer](https://crates.io/crates/update-informer)] — Update informer for CLI applications. It checks for a new version on Crates.io and GitHub [![build badge](https://github.com/mgrachev/update-informer/workflows/CI/badge.svg)](https://github.com/mgrachev/update-informer/actions) Stars:`161`.
* Pipeline
  * [oconnor663/duct.rs](https://github.com/oconnor663/duct.rs) [[duct](https://crates.io/crates/duct)] — A builder for subprocess pipelines and IO redirection [![build badge](https://api.travis-ci.org/oconnor663/duct.rs.svg?branch=master)](https://travis-ci.org/oconnor663/duct.rs) Stars:`610`.
  * [hniksic/rust-subprocess](https://github.com/hniksic/rust-subprocess) [[subprocess](https://crates.io/crates/subprocess)] — facilities for interaction with external pipelines [![build badge](https://api.travis-ci.org/hniksic/rust-subprocess.svg?branch=master)](https://travis-ci.org/hniksic/rust-subprocess) Stars:`363`.
  * [rust-cli/rexpect](https://github.com/rust-cli/rexpect) [[rexpect](https://crates.io/crates/rexpect)] — automate interactive applications such as ssh, ftp, passwd, etc [![CI](https://github.com/rust-cli/rexpect/actions/workflows/ci.yml/badge.svg)](https://github.com/rust-cli/rexpect/actions/workflows/ci.yml) Stars:`167`.
  * [zhiburt/expectrl](https://github.com/zhiburt/expectrl) [[expectrl](https://crates.io/crates/expectrl)] — A library for controlling interactive programs in a pseudo-terminal [![build badge](https://github.com/zhiburt/expectrl/actions/workflows/ci.yml/badge.svg)](https://github.com/zhiburt/expectrl/actions/workflows/ci.yml) Stars:`129`.
  * [imp/pager-rs](https://gitlab.com/imp/pager-rs) [[pager](https://crates.io/crates/pager)] — pipe your output through an external pager
* Progress
  * [console-rs/indicatif](https://github.com/console-rs/indicatif) [[indicatif](https://crates.io/crates/indicatif)] — indicate progress to users Stars:`3.1K`.
  * [a8m/pb](https://github.com/a8m/pb) [[pbr](https://crates.io/crates/pbr)] — console progress bar for Rust Stars:`496`.
  * [FGRibreau/spinners](https://github.com/FGRibreau/spinners) [[spinners](https://crates.io/crates/spinners)] — 60+ elegant terminal spinners Stars:`423`.
  * [etienne-napoleone/spinach](https://github.com/etienne-napoleone/spinach) [[spinach](https://crates.io/crates/spinach)] — Practical spinner for Rust. [![CI](https://github.com/etienne-napoleone/spinach/actions/workflows/ci.yml/badge.svg)](https://github.com/etienne-napoleone/spinach/actions/workflows/ci.yml) Stars:`61`.
* Prompt
  * [ynqa/promkit](https://github.com/ynqa/promkit) [[promkit](https://crates.io/crates/promkit)]  — A toolkit for building interactive command-line tools [![Build status](https://github.com/ynqa/promkit/workflows/promkit/badge.svg?branch=master)](https://github.com/ynqa/promkit/actions) Stars:`69`.
  * [hashmismatch/terminal_cli.rs](https://github.com/hashmismatch/terminal_cli.rs) [[terminal_cli](https://crates.io/crates/terminal_cli)]  — build an interactive command prompt [![build badge](https://api.travis-ci.org/hashmismatch/terminal_cli.rs.svg?branch=master)](https://travis-ci.org/hashmismatch/terminal_cli.rs) Stars:`48`.
  * [starship/starship](https://starship.rs/) [[starship](https://crates.io/crates/starship)]  — A minimal, blazing fast, and extremely customizable prompt for any shell [![Build status](https://github.com/starship/starship/workflows/Main%20workflow/badge.svg?branch=master)](https://github.com/starship/starship/actions)
* Style
  * [mackwic/colored](https://github.com/mackwic/colored) [[colored](https://crates.io/crates/colored)] — Coloring terminal so simple, you already know how to do it! Stars:`1.1K`.
  * [ogham/rust-ansi-term](https://github.com/ogham/rust-ansi-term) [[ansi_term](https://crates.io/crates/ansi_term)] — control colours and formatting on ANSI terminals [![build badge](https://api.travis-ci.org/ogham/rust-ansi-term.svg?branch=master)](https://travis-ci.org/ogham/rust-ansi-term) Stars:`400`.
  * [LukasKalbertodt/bunt](https://github.com/LukasKalbertodt/bunt) [[bunt](https://crates.io/crates/bunt)] — cross-platform terminal colors and styling with macros [![Build status](https://github.com/LukasKalbertodt/bunt/actions/workflows/ci.yml/badge.svg)](https://github.com/LukasKalbertodt/bunt/actions?query=workflow%3ACI+branch%3Amaster) Stars:`199`.
  * [SergioBenitez/yansi](https://github.com/SergioBenitez/yansi) [[yansi](https://crates.io/crates/yansi)] — A dead simple ANSI terminal color painting library Stars:`165`.
  * [LukasKalbertodt/term-painter](https://github.com/LukasKalbertodt/term-painter) [[term-painter](https://crates.io/crates/term-painter)] — cross-platform styled terminal output [![build badge](https://api.travis-ci.org/LukasKalbertodt/term-painter.svg?branch=master)](https://travis-ci.org/LukasKalbertodt/term-painter) Stars:`75`.
* TUI
  * BearLibTerminal
  * [fdehau/tui-rs](https://github.com/fdehau/tui-rs) [[tui](https://crates.io/crates/tui)] — A TUI library inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib) and [termui](https://github.com/gizak/termui) [![build badge](https://api.travis-ci.org/fdehau/tui-rs.svg?branch=master)](https://travis-ci.org/fdehau/tui-rs) Stars:`8.9K`.
  * [gyscos/Cursive](https://github.com/gyscos/Cursive) [[cursive](https://crates.io/crates/cursive)] — build rich TUI applications [![build badge](https://api.travis-ci.org/gyscos/Cursive.svg?branch=master)](https://travis-ci.org/gyscos/Cursive) Stars:`3.2K`.
  * [ivanceras/titik](https://github.com/ivanceras/titik) - a crossplatform TUI widget library with the goal of providing interactive widgets [![Build Status](https://api.travis-ci.com/ivanceras/titik.svg?branch=master)](https://app.travis-ci.com/github/ivanceras/titik) Stars:`110`.
    * [cfyzium/bearlibterminal](https://github.com/nabijaczleweli/BearLibTerminal.rs) [[bear-lib-terminal](https://crates.io/crates/bear-lib-terminal)] — [BearLibTerminal](https://github.com/tommyettinger/BearLibTerminal) bindings [![build badge](https://api.travis-ci.org/nabijaczleweli/BearLibTerminal.rs.svg?branch=master)](https://travis-ci.org/nabijaczleweli/BearLibTerminal.rs) Stars:`30`.
  * ncurses
  * [redox-os/termion](https://github.com/redox-os/termion) [[termion](https://crates.io/crates/termion)] — bindless library for controlling terminals/TTY [![build badge](https://api.travis-ci.org/redox-os/termion.svg?branch=master)](https://travis-ci.org/redox-os/termion) Stars:`1.8K`.
    * [jeaye/ncurses-rs](https://github.com/jeaye/ncurses-rs) [[ncurses](https://crates.io/crates/ncurses)] — [ncurses](https://www.gnu.org/software/ncurses/) bindings [![build badge](https://api.travis-ci.org/jeaye/ncurses-rs.svg?branch=master)](https://travis-ci.org/jeaye/ncurses-rs) Stars:`624`.
    * [ihalila/pancurses](https://github.com/ihalila/pancurses) [[pancurses](https://crates.io/crates/pancurses)] — curses library, supports linux and windows [![build badge](https://api.travis-ci.org/ihalila/pancurses.svg?branch=master)](https://travis-ci.org/ihalila/pancurses) Stars:`353`.
  * [ogham/rust-term-grid](https://github.com/ogham/rust-term-grid) [[term_grid](https://crates.io/crates/term_grid)] — Rust library for putting things in a grid [![build badge](https://api.travis-ci.org/ogham/rust-term-grid.svg?branch=master)](https://travis-ci.org/ogham/rust-term-grid) Stars:`61`.
  * Termbox
  * [TimonPost/crossterm](https://github.com/crossterm-rs/crossterm) [[crossterm](https://crates.io/crates/crossterm)] — crossplatform terminal library Stars:`2.0K`.
    * [gchp/rustbox](https://github.com/gchp/rustbox) [[rustbox](https://crates.io/crates/rustbox)] — bindings to [Termbox](https://github.com/nsf/termbox) [![build badge](https://api.travis-ci.org/gchp/rustbox.svg?branch=master)](https://travis-ci.org/gchp/rustbox) Stars:`459`.

### Compression

  * [dropbox/rust-brotli](https://github.com/dropbox/rust-brotli) — Brotli decompressor in Rust that optionally avoids the stdlib Stars:`647`.
  * [ende76/brotli-rs](https://github.com/ende76/brotli-rs) — implementation of Brotli compression Stars:`59`.
* [Brotli](https://opensource.googleblog.com/2015/09/introducing-brotli-new-compression.html)
* bzip2
  * [alexcrichton/bzip2-rs](https://github.com/alexcrichton/bzip2-rs) — [libbz2](https://www.sourceware.org/bzip2/) bindings [![build badge](https://api.travis-ci.com/alexcrichton/bzip2-rs.svg?branch=master)](https://travis-ci.org/alexcrichton/bzip2-rs) Stars:`67`.
* gzip
  * [zopfli](https://github.com/zopfli-rs/zopfli) [[zopfli](https://crates.io/crates/zopfli)] — implementation of the Zopfli compression algorithm for higher quality deflate or zlib compression Stars:`10`.
* gzp
  * [sstadick/gzp](https://github.com/sstadick/gzp/) - multi-threaded encoding and decoding of deflate formats and snappy
* miniz
  * [rust-lang/flate2-rs](https://github.com/rust-lang/flate2-rs) — [miniz](https://code.google.com/archive/p/miniz) bindings [![build badge](https://github.com/rust-lang/flate2-rs/workflows/CI/badge.svg?branch=master)](https://github.com/rust-lang/flate2-rs/actions) Stars:`587`.
* snappy
  * [JeffBelgum/rust-snappy](https://github.com/JeffBelgum/rust-snappy) — [snappy](https://github.com/google/snappy) bindings [![build badge](https://api.travis-ci.org/JeffBelgum/rust-snappy.svg?branch=master)](https://travis-ci.org/JeffBelgum/rust-snappy) Stars:`14`.
* tar
  * [alexcrichton/tar-rs](https://github.com/alexcrichton/tar-rs) — tar archive reading/writing in Rust [![build badge](https://api.travis-ci.com/alexcrichton/tar-rs.svg?branch=master)](https://travis-ci.org/alexcrichton/tar-rs) Stars:`478`.
* zip
  * [zip-rs/zip](https://github.com/zip-rs/zip) — read and write ZIP archives [![Build Status](https://api.travis-ci.org/mvdnes/zip-rs.svg?branch=master)](https://travis-ci.org/mvdnes/zip-rs) Stars:`526`.

### Computation

* [dimforge/nalgebra](https://github.com/dimforge/nalgebra) — low-dimensional linear algebra library [![build badge](https://api.travis-ci.org/dimforge/nalgebra.svg?branch=dev)](https://travis-ci.org/dimforge/nalgebra) Stars:`2.9K`.
* [calebwin/emu](https://github.com/calebwin/emu) — A language for GPGPU numerical computing from a Rust macro Stars:`1.5K`.
* [argmin-rs/argmin](https://github.com/argmin-rs/argmin) [[argmin](https://crates.io/crates/argmin)] — A pure Rust optimization library [![build badge](https://api.travis-ci.org/argmin-rs/argmin.svg?branch=master)](https://travis-ci.org/argmin-rs/argmin) Stars:`530`.
  * [GuillaumeGomez/rust-GSL](https://github.com/GuillaumeGomez/rust-GSL) — GSL bindings [![build badge](https://api.travis-ci.org/GuillaumeGomez/rust-GSL.svg?branch=master)](https://travis-ci.org/GuillaumeGomez/rust-GSL) Stars:`156`.
  * [mikkyang/rust-blas](https://github.com/mikkyang/rust-blas) — BLAS bindings Stars:`77`.
  * [stainless-steel/lapack](https://github.com/blas-lapack-rs/lapack) — LAPACK bindings [![build badge](https://api.travis-ci.org/blas-lapack-rs/lapack.svg?branch=master)](https://travis-ci.org/blas-lapack-rs/lapack) Stars:`67`.
* [BLAS](https://en.wikipedia.org/wiki/Basic_Linear_Algebra_Subprograms) [[blas](https://crates.io/keywords/blas)]
* [GSL](http://www.gnu.org/software/gsl/)
* [LAPACK](https://en.wikipedia.org/wiki/LAPACK)
* Parallel
  * [arrayfire/arrayfire-rust](https://github.com/arrayfire/arrayfire-rust) — [Arrayfire](https://github.com/arrayfire) bindings Stars:`683`.
  * [autumnai/collenchyma](https://github.com/autumnai/collenchyma) — An extensible, pluggable, backend-agnostic framework for parallel, high-performance computations on CUDA, OpenCL and common host CPU. [![build badge](https://api.travis-ci.org/autumnai/collenchyma.svg?branch=master)](https://travis-ci.org/autumnai/collenchyma) Stars:`460`.
  * [luqmana/rust-opencl](https://github.com/luqmana/rust-opencl) — [OpenCL](https://www.khronos.org/opencl/) bindings Stars:`170`.
* Scirust
  * [indigits/scirust](https://github.com/indigits/scirust) — scientific computing library in Rust [![Build Status](https://api.travis-ci.org/indigits/scirust.svg?branch=master)](https://travis-ci.org/indigits/scirust) Stars:`240`.
* Statrs
  * [statrs-dev/statrs](https://github.com/statrs-dev/statrs) — Robust statistical computation library in Rust [![Build Status](https://api.travis-ci.org/boxtown/statrs.svg?branch=master)](https://travis-ci.org/boxtown/statrs) Stars:`377`.

### Concurrency

* [Rayon](https://github.com/rayon-rs/rayon) – A data parallelism library for Rust [![build badge](https://api.travis-ci.org/rayon-rs/rayon.svg?branch=master)](https://travis-ci.org/rayon-rs/rayon) Stars:`7.5K`.
* [crossbeam-rs/crossbeam](https://github.com/crossbeam-rs/crossbeam) – Support for parallelism and low-level concurrency in Rust [![build badge](https://api.travis-ci.org/crossbeam-rs/crossbeam.svg?branch=master)](https://travis-ci.org/crossbeam-rs/crossbeam) Stars:`5.5K`.
* [zonyitoo/coio-rs](https://github.com/zonyitoo/coio-rs) – Coroutine I/O for Rust [![build badge](https://api.travis-ci.org/zonyitoo/coio-rs.svg?branch=master)](https://travis-ci.org/zonyitoo/coio-rs) Stars:`453`.
* [rustcc/coroutine-rs](https://github.com/rustcc/coroutine-rs) – Coroutine Library in Rust [![build badge](https://api.travis-ci.org/rustcc/coroutine-rs.svg?branch=master)](https://travis-ci.org/rustcc/coroutine-rs) Stars:`401`.
* [orium/archery](https://github.com/orium/archery) [[archery](https://crates.io/crates/archery)] — Library to abstract from `Rc`/`Arc` pointer types. [![build badge](https://api.travis-ci.org/orium/archery.svg?branch=master)](https://travis-ci.org/orium/archery) Stars:`101`.

### Configuration

* [mehcode/config-rs](https://github.com/mehcode/config-rs) [[config](https://crates.io/crates/config)] — Layered configuration system for Rust applications (with strong support for 12-factor applications). [![build badge](https://api.travis-ci.org/mehcode/config-rs.svg?branch=master)](https://travis-ci.org/mehcode/config-rs) Stars:`1.6K`.
* [Kixunil/configure_me](https://github.com/Kixunil/configure_me) [[configure_me](https://crates.io/crates/configure_me)] — library for processing application configuration easily Stars:`48`.
* [andoriyu/uclicious](https://github.com/andoriyu/uclicious) [[uclicious](https://crates.io/crates/uclicious)] — [libUCL](https://github.com/vstakhov/libucl) based feature-rich configuration library. [![CircleCI](https://circleci.com/gh/vstakhov/libucl.svg?style=svg)](https://app.circleci.com/pipelines/github/vstakhov/libucl) Stars:`14`.

### Cryptography

[[crypto](https://crates.io/keywords/crypto), [cryptography](https://crates.io/keywords/cryptography)]

* [rustls/rustls](https://github.com/rustls/rustls) — A Rust implementation of TLS Stars:`3.9K`.
* [briansmith/ring](https://github.com/briansmith/ring) — Safe, fast, small crypto using Rust and BoringSSL's cryptography primitives. [![build badge](https://api.travis-ci.org/briansmith/ring.svg?branch=master)](https://travis-ci.org/briansmith/ring) Stars:`2.9K`.
* [cossacklabs/themis](https://github.com/cossacklabs/themis) [[themis](https://crates.io/crates/themis)] — a high-level cryptographic library for solving typical data security tasks, best fit for multi-platform apps. [![build badge](https://circleci.com/gh/cossacklabs/themis/tree/master.svg?style=shield)](https://app.circleci.com/pipelines/github/cossacklabs/themis) Stars:`1.6K`.
* [DaGenix/rust-crypto](https://github.com/DaGenix/rust-crypto) — cryptographic algorithms in Rust [![build badge](https://api.travis-ci.org/DaGenix/rust-crypto.svg?branch=master)](https://travis-ci.org/DaGenix/rust-crypto) Stars:`1.2K`.
* [exonum/exonum](https://github.com/exonum/exonum) [[exonum](https://crates.io/crates/exonum)] — extensible framework for blockchain projects [![build badge](https://api.travis-ci.com/exonum/exonum.svg?branch=master)](https://travis-ci.org/exonum/exonum) Stars:`1.2K`.
* [RustCrypto/hashes](https://github.com/RustCrypto/hashes) — Collection of cryptographic hash functions written in pure Rust [![build badge](https://api.travis-ci.org/RustCrypto/hashes.svg?branch=master)](https://travis-ci.org/RustCrypto/hashes) Stars:`1.1K`.
* [sfackler/rust-openssl](https://github.com/sfackler/rust-openssl) — [OpenSSL](https://www.openssl.org/) bindings [![build badge](https://api.travis-ci.org/sfackler/rust-openssl.svg?branch=master)](https://travis-ci.org/sfackler/rust-openssl) Stars:`1.0K`.
* [sodiumoxide/sodiumoxide](https://github.com/sodiumoxide/sodiumoxide) — [libsodium](https://github.com/jedisct1/libsodium) bindings [![build badge](https://api.travis-ci.org/sodiumoxide/sodiumoxide.svg?branch=master)](https://travis-ci.org/sodiumoxide/sodiumoxide) Stars:`641`.
* [dalek-cryptography/curve25519-dalek](https://github.com/dalek-cryptography/curve25519-dalek) — Pure Rust implementation of Curve25519 operations Stars:`597`.
* [dalek-cryptography/ed25519-dalek](https://github.com/dalek-cryptography/ed25519-dalek) — Pure Rust implementation of Ed25519 digital signatures Stars:`547`.
* [orion-rs/orion](https://github.com/orion-rs/orion) — This library aims to provide easy and usable crypto. 'Usable' meaning exposing high-level API's that are easy to use and hard to misuse. [![Tests](https://github.com/orion-rs/orion/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/orion-rs/orion/actions/workflows/test.yml) Stars:`457`.
* [briansmith/webpki](https://github.com/briansmith/webpki) — Web PKI TLS X.509 certificate validation in Rust. [![build badge](https://api.travis-ci.org/briansmith/webpki.svg?branch=master)](https://travis-ci.org/briansmith/webpki) Stars:`400`.
* [sfackler/rust-native-tls](https://github.com/sfackler/rust-native-tls) — Bindings for native TLS libraries Stars:`359`.
* [dalek-cryptography/x25519-dalek](https://github.com/dalek-cryptography/x25519-dalek) — Pure Rust implementation of X25519 key exchange Stars:`246`.
* [w3f/schnorrkel](https://github.com/w3f/schnorrkel) - Schnorr VRFs and signatures on the Ristretto group Stars:`246`.
* [debris/tiny-keccak](https://github.com/debris/tiny-keccak) — Pure Rust implementation of the Keccak family (SHA3) Stars:`168`.
* [kornelski/rust-security-framework](https://github.com/kornelski/rust-security-framework) — Bindings for Security Framework (OSX native) Stars:`164`.
* [novifinancial/opaque-ke](https://github.com/novifinancial/opaque-ke) — Pure Rust implementation of the recent [OPAQUE](https://datatracker.ietf.org/doc/draft-krawczyk-cfrg-opaque/) password-authenticated key exchange. [![build badge](https://github.com/novifinancial/opaque-ke/workflows/Rust%20CI/badge.svg?branch=master)](https://github.com/novifinancial/opaque-ke) Stars:`161`.
* [libOctavo/octavo](https://github.com/libOctavo/octavo) — Modular hash and crypto library in Rust [![build badge](https://api.travis-ci.org/libOctavo/octavo.svg?branch=master)](https://travis-ci.org/libOctavo/octavo) Stars:`138`.
* [conradkleinespel/rooster](https://github.com/conradkleinespel/rooster) [[rooster](https://crates.io/crates/rooster)] — Simple password manager to use in your terminal Stars:`130`.
* [klutzy/suruga](https://github.com/klutzy/suruga) — A Rust implementation of [TLS 1.2](https://datatracker.ietf.org/doc/html/rfc5246) Stars:`123`.
* [gakonst/ark-circom](https://github.com/gakonst/ark-circom) - Arkworks bindings to Circom's R1CS, for Groth16 Proof and Witness generation in Rust. Stars:`119`.
* [racum/rust-djangohashers](https://github.com/racum/rust-djangohashers) [[djangohashers](https://crates.io/crates/djangohashers)] — A Rust port of the password primitives used in the Django Project. It doesn't require Django, only hashes and validates passwords according to its style. [![build badge](https://api.travis-ci.org/Racum/rust-djangohashers.svg?branch=master)](https://travis-ci.org/Racum/rust-djangohashers) Stars:`51`.
* [vityafx/randomorg](https://github.com/vityafx/randomorg) - A random.org client library. [![Crates badge](https://img.shields.io/crates/v/randomorg.svg)](https://crates.io/crates/randomorg) Stars:`7`.

### Data processing

* [pola-rs/polars](https://github.com/pola-rs/polars) - Fast feature complete DataFrame library ![Build and test](https://github.com/pola-rs/polars/workflows/Build%20and%20test/badge.svg?branch=master) Stars:`9.0K`.
* [weld-project/weld](https://github.com/weld-project/weld) — High-performance runtime for data analytics applications Stars:`2.8K`.
* [bluss/ndarray](https://github.com/rust-ndarray/ndarray) — N-dimensional array with array views, multidimensional slicing, and efficient operations Stars:`2.4K`.
* [amv-dev/yata](https://github.com/amv-dev/yata) — high perfomance technical analysis library [![Build Status](https://img.shields.io/github/workflow/status/amv-dev/yata/Rust?branch=master)](https://github.com/amv-dev/yata/actions?query=workflow%3ARust) Stars:`190`.
* [kernelmachine/utah](https://github.com/kernelmachine/utah) — Dataframe structure and operations in Rust Stars:`139`.

### Data streaming

* [infinyon/fluvio](https://github.com/infinyon/fluvio) - Programmable data streaming platform [![CI](https://github.com/infinyon/fluvio/workflows/CI/badge.svg?branch=stable)](https://github.com/infinyon/fluvio/actions) Stars:`1.5K`.

### Data structures

* [rust-itertools/itertools](https://github.com/rust-itertools/itertools) — [![build badge](https://api.travis-ci.org/rust-itertools/itertools.svg?branch=master)](https://travis-ci.org/rust-itertools/itertools) Stars:`1.9K`.
* [orium/rpds](https://github.com/orium/rpds) [[rpds](https://crates.io/crates/rpds)] — Persistent data structures in Rust. [![build badge](https://github.com/orium/rpds/workflows/CI/badge.svg)](https://github.com/orium/rpds/actions?query=workflow%3ACI) Stars:`864`.
* [RoaringBitmap/roaring-rs](https://github.com/RoaringBitmap/roaring-rs) – Roaring Bitmaps in Rust Stars:`536`.
* [fizyk20/generic-array](https://github.com/fizyk20/generic-array) – a hack to allow for arrays sized by typenums [![build badge](https://api.travis-ci.org/fizyk20/generic-array.svg?branch=master)](https://travis-ci.org/fizyk20/generic-array) Stars:`318`.
* [yamafaktory/hypergraph](https://github.com/yamafaktory/hypergraph) [[hypergraph](https://crates.io/crates/hypergraph)] — Hypergraph is a data structure library to generate directed hypergraphs. [![ci](https://github.com/yamafaktory/hypergraph/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/yamafaktory/hypergraph/actions/workflows/ci.yml) Stars:`214`.
* [tnballo/scapegoat](https://github.com/tnballo/scapegoat) [[scapegoat](https://crates.io/crates/scapegoat)] — Safe, fallible, stack-only alternative to `BTreeSet` and `BTreeMap`. [![GitHub Actions](https://github.com/tnballo/scapegoat/workflows/test/badge.svg?branch=master)](https://github.com/tnballo/scapegoat/actions) Stars:`214`.
* [mrhooray/kdtree-rs](https://github.com/mrhooray/kdtree-rs) — K-dimensional tree in Rust for fast geospatial indexing and nearest neighbors lookup Stars:`140`.
* [danielpclark/array_tool](https://github.com/danielpclark/array_tool) — Array helpers for Rust. Some of the most common methods you would use on Arrays made available on Vectors. Polymorphic implementations for handling most of your use cases. [![build badge](https://api.travis-ci.org/danielpclark/array_tool.svg?branch=master)](https://travis-ci.org/danielpclark/array_tool) Stars:`65`.
* [xfix/enum-map](https://github.com/xfix/enum-map) [[enum-map](https://crates.io/crates/enum-map)] — An optimized map implementation for enums using an array to store values. [![build badge](https://api.travis-ci.org/xfix/enum-map.svg?branch=master)](https://travis-ci.org/xfix/enum-map) Stars:`52`.
* [becheran/grid](https://github.com/becheran/grid) [[grid](https://crates.io/crates/grid)] —  Provide a two dimensional data structure for rust that is easy to use and fast. [![build status](https://github.com/becheran/grid/actions/workflows/rust.yml/badge.svg)](https://github.com/becheran/grid/actions) Stars:`42`.
* [billyevans/tst](https://github.com/billyevans/tst) [[tst](https://crates.io/crates/tst)] — Ternary search tree collection [![build badge](https://api.travis-ci.org/billyevans/tst.svg?branch=master)](https://travis-ci.org/billyevans/tst) Stars:`20`.
* [garro95/priority-queue](https://github.com/garro95/priority-queue)[[priority-queue](https://crates.io/crates/priority-queue)] — A priority queue that implements priority changes. [![build badge](https://api.travis-ci.org/garro95/priority-queue.svg?branch=master)](https://travis-ci.org/garro95/priority-queue)
* [contain-rs](https://github.com/contain-rs) — Extension of Rust's std::collections

### Data visualization

* [plotters](https://github.com/plotters-rs/plotters) — [![build badge](https://github.com/plotters-rs/plotters/workflows/CI/badge.svg)](https://github.com/plotters-rs/plotters/actions) Stars:`2.5K`.
* [igiagkiozis/plotly](https://github.com/igiagkiozis/plotly) — Plotly for Rust. Stars:`625`.
* [milliams/plotlib](https://github.com/milliams/plotlib) — [![build badge](https://api.travis-ci.org/milliams/plotlib.svg?branch=master)](https://travis-ci.org/milliams/plotlib) Stars:`407`.
* [saresend/gust](https://github.com/saresend/Gust) — [![build badge](https://api.travis-ci.org/saresend/Gust.svg?branch=master)](https://travis-ci.org/saresend/Gust) Stars:`128`.
* [djduque/pgfplots](https://github.com/djduque/pgfplots) [[pgfplots](https://crates.io/crates/pgfplots)] — A Rust library to generate publication-quality figures. [![build](https://github.com/DJDuque/pgfplots/actions/workflows/rust.yml/badge.svg)](https://github.com/DJDuque/pgfplots/actions/workflows/rust.yml) Stars:`62`.

### Database

[[database](https://crates.io/keywords/database)]

* NoSQL [[nosql](https://crates.io/keywords/nosql)]

    * [AlexPikalov/cdrs](https://github.com/AlexPikalov/cdrs) [[cdrs](https://crates.io/crates/cdrs)] — native client written in Rust [![build badge](https://api.travis-ci.org/AlexPikalov/cdrs.svg?branch=master)](https://travis-ci.org/AlexPikalov/cdrs) Stars:`338`.
    * [Arangors](https://github.com/fMeow/arangors) [[arangors](https://crates.io/crates/arangors)] - An ArangoDB driver for Rust Stars:`110`.
    * [Metaswitch/cassandra-rs](https://github.com/Metaswitch/cassandra-rs) —  bindings to the DataStax C/C++ client [![build badge](https://api.travis-ci.org/Metaswitch/cassandra-rs.svg?branch=master)](https://travis-ci.org/Metaswitch/cassandra-rs) Stars:`87`.
    * [krojew/cdrs-tokio](https://github.com/krojew/cdrs-tokio) [[cdrs-tokio](https://crates.io/crates/cdrs-tokio)] - production-ready async Apache Cassandra driver written in pure Rust [![build badge](https://github.com/krojew/cdrs-tokio/actions/workflows/rust.yml/badge.svg)](https://github.com/krojew/cdrs-tokio/actions) Stars:`68`.
  * [ArangoDB](https://www.arangodb.com)
    * [Aragog](https://gitlab.com/qonfucius/aragog) [[aragog](https://crates.io/crates/aragog)] - A Lightweight ArangoDB Object document, relational and graph mapper [![pipeline status](https://gitlab.com/qonfucius/aragog/badges/master/pipeline.svg)](https://gitlab.com/qonfucius/aragog/-/commits/master)
  * [Cassandra](https://cassandra.apache.org/_/index.html) [[cassandra](https://crates.io/keywords/cassandra), [cql](https://crates.io/keywords/cql)]
  * CouchDB [[couchdb](https://crates.io/keywords/couchdb)]
    * [softprops/dynomite](https://github.com/softprops/dynomite) - A library for strongly-typed and convenient interaction with `rusoto_dynamodb` [![build badge](https://github.com/softprops/dynomite/workflows/Main/badge.svg?branch=master)](https://github.com/softprops/dynomite/actions) Stars:`196`.
    * [chill-rs/chill](https://github.com/chill-rs/chill) [[couchdb](https://crates.io/crates/chill)] — A Rust client for the CouchDB REST API [![build badge](https://api.travis-ci.org/chill-rs/chill.svg?branch=master)](https://travis-ci.org/chill-rs/chill) Stars:`35`.
  * [DynamoDB](https://aws.amazon.com/dynamodb/) [[dynamodb](https://crates.io/keywords/dynamodb)]
  * Elasticsearch [[elasticsearch](https://crates.io/keywords/elasticsearch)]
    * [elastic-rs/elastic](https://github.com/elastic-rs/elastic) [[elastic](https://crates.io/crates/elastic)] — elastic is an efficient, modular API client for Elasticsearch written in Rust [![build badge](https://ci.appveyor.com/api/projects/status/csa78tcumdpnbur2?svg=true)](https://ci.appveyor.com/project/KodrAus/elastic) Stars:`249`.
    * [benashford/rs-es](https://github.com/benashford/rs-es) [[rs-es](https://crates.io/crates/rs-es)] — A Rust client for the [Elastic](https://www.elastic.co/) REST API [![build badge](https://api.travis-ci.org/benashford/rs-es.svg?branch=master)](https://travis-ci.org/benashford/rs-es) Stars:`216`.
  * etcd
    * [lodrem/etcd-rs](https://github.com/lodrem/etcd-rs) — An asynchronous etcd client for rust [![CI](https://github.com/lodrem/etcd-rs/actions/workflows/ci.yml/badge.svg)](https://github.com/lodrem/etcd-rs/actions/workflows/ci.yml) Stars:`154`.
    * [jimmycuadra/rust-etcd](https://github.com/jimmycuadra/rust-etcd) [[etcd](https://crates.io/crates/etcd)] — A client library for CoreOS's etcd. [![build badge](https://api.travis-ci.org/jimmycuadra/rust-etcd.svg?branch=master)](https://travis-ci.org/jimmycuadra/rust-etcd) Stars:`135`.
  * ForestDB
    * [driftluo/InfluxDBClient-rs](https://github.com/driftluo/InfluxDBClient-rs) — Synchronization interface [![build badge](https://api.travis-ci.org/driftluo/InfluxDBClient-rs.svg?branch=master)](https://travis-ci.org/driftluo/InfluxDBClient-rs) Stars:`75`.
    * [vhbit/sherwood](https://github.com/vhbit/sherwood) — [ForestDB](https://github.com/couchbase/forestdb) bindings [![build badge](https://api.travis-ci.org/vhbit/sherwood.svg?branch=master)](https://travis-ci.org/vhbit/sherwood) Stars:`8`.
  * [InfluxDB](https://www.influxdata.com/)
  * LevelDB
    * [skade/leveldb](https://github.com/skade/leveldb) — [LevelDB](https://github.com/google/leveldb) bindings [![build badge](https://api.travis-ci.org/skade/leveldb.svg?branch=master)](https://travis-ci.org/skade/leveldb) Stars:`160`.
  * LMDB [[lmdb](https://crates.io/keywords/lmdb)]
    * [vhbit/lmdb-rs](https://github.com/vhbit/lmdb-rs) [[lmdb-rs](https://crates.io/crates/lmdb-rs)] — [LMDB](https://www.symas.com/symas-embedded-database-lmdb) bindings [![build badge](https://api.travis-ci.org/vhbit/lmdb-rs.svg?branch=master)](https://travis-ci.org/vhbit/lmdb-rs) Stars:`103`.
  * MongoDB [[mongodb](https://crates.io/keywords/mongodb)]
    * [mongodb/mongo-rust-driver](https://github.com/mongodb/mongo-rust-driver) [[mongodb](https://crates.io/crates/mongodb)] — [MongoDB](https://www.mongodb.com/) bindings Stars:`1.1K`.
    * [seladb/pickledb-rs](https://github.com/seladb/pickledb-rs) — a lightweight and simple key-value store, heavily inspired by Python's PickleDB. [![build badge](https://api.travis-ci.org/seladb/pickledb-rs.svg?branch=master)](https://travis-ci.org/seladb/pickledb-rs) Stars:`148`.
  * [PickleDB](https://pythonhosted.org/pickleDB/)
  * Redis [[redis](https://crates.io/keywords/redis)]
    * [surrealdb/surrealdb](https://github.com/surrealdb/surrealdb) — SurrealDB embedded document-graph database Stars:`15.1K`.
    * [redis-rs](https://github.com/redis-rs/redis-rs) — [Redis](https://redis.io/) library in Rust [![Rust](https://github.com/redis-rs/redis-rs/actions/workflows/rust.yml/badge.svg)](https://github.com/redis-rs/redis-rs/actions/workflows/rust.yml) Stars:`2.7K`.
    * [rust-rocksdb/rust-rocksdb](https://github.com/rust-rocksdb/rust-rocksdb) — RocksDB bindings [![RocksDB CI](https://github.com/rust-rocksdb/rust-rocksdb/actions/workflows/rust.yml/badge.svg?branch=master)](https://github.com/rust-rocksdb/rust-rocksdb/actions/workflows/rust.yml) Stars:`1.3K`.
    * [bonifaido/rust-zookeeper](https://github.com/bonifaido/rust-zookeeper) [[zookeeper](https://crates.io/crates/zookeeper)] — A client library for Apache ZooKeeper. [![build badge](https://api.travis-ci.org/bonifaido/rust-zookeeper.svg?branch=master)](https://travis-ci.org/bonifaido/rust-zookeeper) Stars:`166`.
    * [zitsen/unqlite.rs](https://github.com/zitsen/unqlite.rs) — UnQLite bindings [![build badge](https://api.travis-ci.org/zitsen/unqlite.rs.svg?branch=master)](https://travis-ci.org/zitsen/unqlite.rs) Stars:`99`.
    * [aembke/fred](https://github.com/aembke/fred.rs) [[fred](https://crates.io/crates/fred)] - A high level async [Redis](https://redis.io/) client for Rust with Tokio. [![CircleCI](https://circleci.com/gh/aembke/fred.rs/tree/main.svg?style=svg)]([https://circleci.com/gh/aembke/fred.rs/tree/main](https://app.circleci.com/pipelines/github/aembke/fred.rs?branch=main)) Stars:`95`.
    * [krojew/rust-zookeeper](https://github.com/krojew/rust-zookeeper) [[zookeeper-async](https://crates.io/crates/zookeeper-async)] - Async Zookeeper client written 100% in Rust, based on tokio.  ![build status](https://github.com/krojew/rust-zookeeper/actions/workflows/rust.yml/badge.svg) Stars:`15`.
  * [RocksDB](https://rocksdb.org/)
  * [SurrealDB](https://surrealdb.com/)
  * [UnQLite](https://unqlite.org/)
  * [ZooKeeper](https://zookeeper.apache.org/)
* OGM [[ogm](https://crates.io/keywords/ogm)]
    * [Aragog](https://gitlab.com/qonfucius/aragog) [[aragog](https://crates.io/crates/aragog)] - A Lightweight ArangoDB Object document, relational and graph mapper [![pipeline status](https://gitlab.com/qonfucius/aragog/badges/master/pipeline.svg)](https://gitlab.com/qonfucius/aragog/-/commits/master)
* ORM [[orm](https://crates.io/keywords/orm)]
  * [diesel-rs/diesel](https://github.com/diesel-rs/diesel) — an ORM and Query builder for Rust [![Build Status](https://api.travis-ci.org/diesel-rs/diesel.svg?branch=master)](https://travis-ci.org/diesel-rs/diesel) Stars:`9.4K`.
  * [SeaQL/sea-orm](https://github.com/SeaQL/sea-orm) — 🐚 An async & dynamic ORM for Rust [![Build Status](https://github.com/SeaQL/sea-orm/actions/workflows/rust.yml/badge.svg)](https://github.com/SeaQL/sea-orm/actions/workflows/rust.yml) Stars:`3.1K`.
  * [rbatis/rbatis](https://github.com/rbatis/rbatis) — Rust ORM Framework High Performance(JSON based) [![Build Status](https://api.travis-ci.org/zhuxiujia/rbatis.svg?branch=master)](https://travis-ci.org/zhuxiujia/rbatis) Stars:`1.6K`.
* [sfackler/r2d2](https://github.com/sfackler/r2d2) — generic connection pool [![build badge](https://api.travis-ci.org/sfackler/r2d2.svg?branch=master)](https://travis-ci.org/sfackler/r2d2) Stars:`1.1K`.
  * [Brendonovich/prisma-client-rust](https://github.com/Brendonovich/prisma-client-rust) — An autogenerated query builder that provides simple and fully type-safe database access using the Prisma ecosystem. [![Test Status](https://img.shields.io/github/workflow/status/Brendonovich/prisma-client-rust/CI?label=tests&style=flat-square)](https://github.com/Brendonovich/prisma-client-rust/actions) Stars:`712`.
  * [ivanceras/rustorm](https://github.com/ivanceras/rustorm) — an ORM for Rust [![Build Status](https://api.travis-ci.org/ivanceras/rustorm.svg?branch=master)](https://travis-ci.org/ivanceras/rustorm) Stars:`235`.
* SQL [[sql](https://crates.io/keywords/sql)]
  * Generic
    * [launchbadge/sqlx](https://github.com/launchbadge/sqlx) - async PostgreSQL/MySQL/SQLite connection pool with strong typing support [![build badge](https://img.shields.io/github/workflow/status/launchbadge/sqlx/Rust/master?style=flat-square)](https://github.com/launchbadge/sqlx) Stars:`7.2K`.
    * [SeaQL/sea-query](https://github.com/SeaQL/sea-query) - 🔱 A dynamic SQL query builder for MySQL, Postgres and SQLite [![build status](https://github.com/SeaQL/sea-query/actions/workflows/rust.yml/badge.svg)](https://github.com/SeaQL/sea-query/actions/workflows/rust.yml) Stars:`594`.
    * [SeaQL/sea-schema](https://github.com/SeaQL/sea-schema) - 🌿 SQL schema management suite [![build status](https://github.com/SeaQL/sea-schema/actions/workflows/rust.yml/badge.svg)](https://github.com/SeaQL/sea-schema/actions/workflows/rust.yml) Stars:`85`.
  * Microsoft SQL
    * [prisma/tiberius](https://github.com/prisma/tiberius) — [![Cargo tests](https://github.com/prisma/tiberius/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/prisma/tiberius/actions/workflows/test.yml) Stars:`176`.
  * MySql [[mysql](https://crates.io/keywords/mysql)]
    * [blackbeam/rust-mysql-simple](https://github.com/blackbeam/rust-mysql-simple) [[mysql](https://crates.io/crates/mysql)] — A native MySql client [![build badge](https://api.travis-ci.org/blackbeam/rust-mysql-simple.svg?branch=master)](https://travis-ci.org/blackbeam/rust-mysql-simple) Stars:`533`.
    * [blackbeam/mysql_async](https://github.com/blackbeam/mysql_async) [[mysql_async](https://crates.io/crates/mysql_async)] — asyncronous Rust Mysql driver based on Tokio. [![CircleCI](https://circleci.com/gh/blackbeam/mysql_async/tree/master.svg?style=shield)](https://app.circleci.com/pipelines/github/blackbeam/mysql_async?branch=master) Stars:`284`.
    * [AgilData/mysql-proxy-rs](https://github.com/AgilData/mysql-proxy-rs) — A MySQL Proxy [![CircleCI](https://circleci.com/gh/AgilData/mysql-proxy-rs/tree/master.svg?style=svg)](https://app.circleci.com/pipelines/github/AgilData/mysql-proxy-rs?branch=master) Stars:`173`.
  * PostgreSql [[postgres](https://crates.io/keywords/postgres), [postgresql](https://crates.io/keywords/postgresql)]
    * [sfackler/rust-postgres](https://github.com/sfackler/rust-postgres) [[postgres](https://crates.io/crates/postgres)] — A native [PostgreSQL](https://www.postgresql.org/) client [![build badge](https://api.travis-ci.org/sfackler/rust-postgres.svg?branch=master)](https://travis-ci.org/sfackler/rust-postgres) Stars:`2.7K`.
  * Sqlite [[sqlite](https://crates.io/keywords/sqlite)]
    * [rusqlite](https://github.com/rusqlite/rusqlite) — [Sqlite3](https://www.sqlite.org/index.html) bindings [![build badge](https://api.travis-ci.org/rusqlite/rusqlite.svg?branch=master)](https://travis-ci.org/rusqlite/rusqlite) Stars:`1.8K`.

### Date and time

[[date](https://crates.io/keywords/date), [time](https://crates.io/keywords/time)]

* [chronotope/chrono](https://github.com/chronotope/chrono) — [![build badge](https://api.travis-ci.org/chronotope/chrono.svg?branch=master)](https://travis-ci.org/chronotope/chrono) Stars:`2.4K`.
* [time-rs/time](https://github.com/time-rs/time) — [![build badge](https://github.com/time-rs/time/workflows/Build/badge.svg)](https://github.com/time-rs/time/actions) Stars:`651`.
* [Mnwa/ms](https://github.com/Mnwa/ms) [[ms-converter](https://crates.io/crates/ms-converter)] — it's a library for converting human-like times to milliseconds [![build badge](https://github.com/Mnwa/ms/workflows/build/badge.svg?branch=master)](https://github.com/Mnwa/ms/actions?query=workflow%3Abuild) Stars:`30`.

### Distributed systems

* Antimony
  * [antimonyproject/antimony](https://github.com/antimonyproject/antimony) [[antimony](https://crates.io/crates/antimony)] — stream processing / distributed computation platform [![build badge](https://api.travis-ci.org/antimonyproject/antimony.svg?branch=master)](https://travis-ci.org/antimonyproject/antimony) Stars:`61`.
* Apache Kafka
  * [fede1024/rust-rdkafka](https://github.com/fede1024/rust-rdkafka) [[rdkafka](https://crates.io/crates/rdkafka)] — [librdkafka](https://github.com/edenhill/librdkafka) bindings [![build badge](https://api.travis-ci.org/fede1024/rust-rdkafka.svg?branch=master)](https://travis-ci.org/fede1024/rust-rdkafka) Stars:`1.0K`.
  * [kafka-rust/kafka-rust](https://github.com/kafka-rust/kafka-rust) — [![build badge](https://api.travis-ci.org/kafka-rust/kafka-rust.svg?branch=master)](https://travis-ci.org/kafka-rust/kafka-rust) Stars:`868`.
  * [gklijs/schema_registry_converter](https://github.com/gklijs/schema_registry_converter) [[schema_registry_converter](https://crates.io/crates/schema_registry_converter)] — to integrate with [confluent schema registry](https://www.confluent.io/product/confluent-platform/data-compatibility/) [![build badge](https://api.travis-ci.org/gklijs/schema_registry_converter.svg?branch=master)](https://travis-ci.org/gklijs/schema_registry_converter) Stars:`65`.
* Beanstalkd
  * [schickling/rust-beanstalkd](https://github.com/schickling/rust-beanstalkd) — [Beanstalkd](https://github.com/beanstalkd/beanstalkd) bindings [![build badge](https://api.travis-ci.org/schickling/rust-beanstalkd.svg?branch=master)](https://travis-ci.org/schickling/rust-beanstalkd) Stars:`44`.
* HDFS
  * [hyunsik/hdfs-rs](https://github.com/hyunsik/hdfs-rs) [[hdfs](https://crates.io/crates/hdfs)] — libhdfs bindings Stars:`31`.

### Domain driven design

  * [serverlesstechnology/cqrs](https://github.com/serverlesstechnology/cqrs) [[cqrs-es](https://crates.io/crates/cqrs-es)] — A framework for CQRS and event sourcing with [user guide](https://doc.rust-cqrs.org/) Stars:`138`.

### Email

[[email](https://crates.io/keywords/email), [imap](https://crates.io/keywords/imap), [smtp](https://crates.io/keywords/smtp)]

* [lettre/lettre](https://github.com/lettre/lettre) — an SMTP-library for Rust [![CI](https://github.com/lettre/lettre/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/lettre/lettre/actions/workflows/test.yml) Stars:`1.2K`.
* [jdrouet/mrml](https://github.com/jdrouet/mrml) [![build badge](https://api.travis-ci.com/jdrouet/mrml.svg?branch=master)](https://travis-ci.org/jdrouet/mrml) - A library to generate nice email templates working on any mail client. Stars:`210`.
* [stalwartlabs/mail-parser](https://github.com/stalwartlabs/mail-parser) [[mail-parser](https://crates.io/crates/mail-parser)] - A fast and robust e-mail parsing library with full MIME support [![build badge](https://github.com/stalwartlabs/mail-parser/actions/workflows/rust.yml/badge.svg)](https://github.com/stalwartlabs/mail-parser/actions/workflows/rust.yml) Stars:`150`.
* [staktrace/mailparse](https://github.com/staktrace/mailparse) [[mailparse](https://crates.io/crates/mailparse)] — A library for parsing real-world email files [![build badge](https://api.travis-ci.org/staktrace/mailparse.svg?branch=master)](https://travis-ci.org/staktrace/mailparse) Stars:`148`.
* [stalwartlabs/mail-send](https://github.com/stalwartlabs/mail-send) [[mail-send](https://crates.io/crates/mail-send)] - E-mail builder and SMTP client library with DKIM support [![build badge](https://github.com/stalwartlabs/mail-send/actions/workflows/rust.yml/badge.svg)](https://github.com/stalwartlabs/mail-send/actions/workflows/rust.yml) Stars:`115`.
* [jdrouet/catapulte](https://github.com/jdrouet/catapulte) [![build badge](https://api.travis-ci.com/jdrouet/catapulte.svg?branch=main)](https://travis-ci.org/jdrouet/catapulte) - A microservice to send emails using [MRML](https://github.com/jdrouet/mrml) templates. Stars:`102`.
* [jdrouet/jolimail](https://github.com/jdrouet/jolimail) [![pipeline status](https://gitlab.com/jeremie.drouet/jolimail/badges/main/pipeline.svg)](https://gitlab.com/jeremie.drouet/jolimail/-/commits/main) - A web application to build [MRML](https://github.com/jdrouet/mrml) templates. Stars:`97`.
* [gsquire/sendgrid-rs](https://github.com/gsquire/sendgrid-rs) — unofficial Rust library for SendGrid API [![build badge](https://api.travis-ci.org/gsquire/sendgrid-rs.svg?branch=master)](https://travis-ci.org/gsquire/sendgrid-rs) Stars:`85`.

### Encoding

[[encoding](https://crates.io/keywords/encoding)]

* ASN.1
  * [alex/rust-asn1](https://github.com/alex/rust-asn1) — A Rust ASN.1 (DER) serializer [![build badge](https://api.travis-ci.org/alex/rust-asn1.svg?branch=master)](https://travis-ci.org/alex/rust-asn1) Stars:`83`.
* Binary
  * [bincode-org/bincode](https://github.com/bincode-org/bincode) — A binary encoder/decoder in Rust [![CI](https://github.com/bincode-org/bincode/actions/workflows/rust.yml/badge.svg?branch=trunk)](https://github.com/bincode-org/bincode/actions/workflows/rust.yml) Stars:`1.8K`.
  * [m4b/goblin](https://github.com/m4b/goblin) [[goblin](https://crates.io/crates/goblin)] —  cross-platform, zero-copy, and endian-aware binary parsing [![build badge](https://api.travis-ci.org/m4b/goblin.svg?branch=master)](https://travis-ci.org/m4b/goblin) Stars:`871`.
* BSON
  * [mongodb/bson-rust](https://github.com/mongodb/bson-rust) — Encoding and decoding support for BSON in Rust Stars:`296`.
* Byte swapping
  * [BurntSushi/byteorder](https://github.com/BurntSushi/byteorder) — Supports big-endian, little-endian and native byte orders [![build badge](https://api.travis-ci.org/BurntSushi/byteorder.svg?branch=master)](https://travis-ci.org/BurntSushi/byteorder) Stars:`794`.
* Cap'n Proto
  * [capnproto/capnproto-rust](https://github.com/capnproto/capnproto-rust) — [![build badge](https://api.travis-ci.org/capnproto/capnproto-rust.svg?branch=master)](https://travis-ci.org/capnproto/capnproto-rust) Stars:`1.5K`.
* CBOR
  * [serde_cbor](https://crates.io/crates/serde_cbor) — CBOR support for serde [![build badge](https://api.travis-ci.org/pyfisch/cbor.svg?branch=master)](https://travis-ci.org/pyfisch/cbor)
* Character Encoding
  * [hsivonen/encoding_rs](https://github.com/hsivonen/encoding_rs) [[encoding_rs](https://crates.io/crates/encoding_rs)] — A Gecko-oriented implementation of the Encoding Standard in Rust [![build badge](https://api.travis-ci.org/hsivonen/encoding_rs.svg?branch=master)](https://travis-ci.org/hsivonen/encoding_rs) Stars:`277`.
  * [lifthrasiir/rust-encoding](https://github.com/lifthrasiir/rust-encoding) — [![build badge](https://api.travis-ci.org/lifthrasiir/rust-encoding.svg?branch=master)](https://travis-ci.org/lifthrasiir/rust-encoding) Stars:`261`.
* CRC
  * [mrhooray/crc-rs](https://github.com/mrhooray/crc-rs) — [![build badge](https://api.travis-ci.org/mrhooray/crc-rs.svg?branch=master)](https://travis-ci.org/mrhooray/crc-rs) Stars:`119`.
* CSV
  * [BurntSushi/rust-csv](https://github.com/BurntSushi/rust-csv) — A fast and flexible CSV reader and writer, with support for Serde [![build badge](https://api.travis-ci.org/BurntSushi/rust-csv.svg?branch=master)](https://travis-ci.org/BurntSushi/rust-csv) Stars:`1.3K`.
* EDN
  * [frol/flatc-rust](https://github.com/frol/flatc-rust) — FlatBuffers compiler (flatc) integration for Cargo build scripts [![build badge](https://api.travis-ci.org/frol/flatc-rust.svg?branch=master)](https://travis-ci.org/frol/flatc-rust) Stars:`84`.
  * [naomijub/edn-rs](https://github.com/naomijub/edn-rs) [[edn-rs](https://crates.io/crates/edn-rs)] — crate to parse and emit EDN format into Rust types. [![Build Status]( https://api.travis-ci.org/naomijub/edn-rs.svg?branch=master)](https://travis-ci.org/naomijub/edn-rs) Stars:`61`.
* [FlatBuffers](https://google.github.io/flatbuffers/)
* HAR
  * [mandrean/har-rs](https://github.com/mandrean/har-rs) [[har](https://crates.io/crates/har)] — A HTTP Archive Format (HAR) serialization & deserialization library [![Build Status](https://api.travis-ci.org/mandrean/har-rs.svg?branch=master)](https://travis-ci.org/mandrean/har-rs) Stars:`23`.
* HTML
  * [servo/html5ever](https://github.com/servo/html5ever) — High-performance browser-grade HTML5 parser [![build badge](https://api.travis-ci.com/servo/html5ever.svg?branch=master)](https://travis-ci.org/servo/html5ever) Stars:`1.6K`.
* JSON
  * [serde-rs/json](https://github.com/serde-rs/json) [[serde\_json](https://crates.io/crates/serde_json)] — JSON support for [Serde](https://github.com/serde-rs/serde) framework [![build badge](https://api.travis-ci.org/serde-rs/json.svg?branch=master)](https://travis-ci.org/serde-rs/json) Stars:`3.4K`.
  * [simd-lite/simd-json](https://github.com/simd-lite/simd-json) [[simd-json](https://crates.io/crates/simd-json)] — High performance JSON parser based on a port of simdjson Stars:`711`.
  * [pikkr/pikkr](https://github.com/pikkr/pikkr) [[pikkr](https://crates.io/crates/pikkr)] — JSON parser which picks up values directly without performing tokenization in Rust Stars:`609`.
  * [maciejhirsz/json-rust](https://github.com/maciejhirsz/json-rust) [[json](https://crates.io/crates/json)] — JSON implementation in Rust [![build badge](https://api.travis-ci.org/maciejhirsz/json-rust.svg?branch=master)](https://travis-ci.org/maciejhirsz/json-rust) Stars:`493`.
  * [importcjj/rust-ajson](https://github.com/importcjj/rust-ajson) [[ajson]](https://crates.io/crates/ajson) —  Get JSON values quickly [![build badge](https://api.travis-ci.com/importcjj/rust-ajson.svg?branch=master)](https://app.travis-ci.com/github/importcjj/rust-ajson) Stars:`87`.
* MsgPack
  * [3Hren/msgpack-rust](https://github.com/3Hren/msgpack-rust) — A pure Rust low/high level MessagePack implementation [![build badge](https://api.travis-ci.org/3Hren/msgpack-rust.svg?branch=master)](https://travis-ci.org/3Hren/msgpack-rust) Stars:`817`.
* PEM
  * [jcreekmore/pem-rs](https://github.com/jcreekmore/pem-rs) [[pem](https://crates.io/crates/pem)] — A Rust based way to parse and encode PEM-encoded data [![build badge](https://api.travis-ci.org/jcreekmore/pem-rs.svg?branch=master)](https://travis-ci.org/jcreekmore/pem-rs) Stars:`28`.
* ProtocolBuffers
  * [tokio-rs/prost](https://github.com/tokio-rs/prost) — [![continuous integration](https://github.com/tokio-rs/prost/workflows/continuous%20integration/badge.svg?branch=master)](https://github.com/tokio-rs/prost/actions) Stars:`2.4K`.
  * [stepancheg/rust-protobuf](https://github.com/stepancheg/rust-protobuf) — [![build badge](https://api.travis-ci.org/stepancheg/rust-protobuf.svg?branch=master)](https://travis-ci.org/stepancheg/rust-protobuf) Stars:`2.2K`.
* RON (Rusty Object Notation)
  * [https://github.com/ron-rs/ron](https://github.com/ron-rs/ron) — [![build badge](https://api.travis-ci.org/ron-rs/ron.svg?branch=master)](https://travis-ci.org/https://github.com/ron-rs/ron) Stars:`2.3K`.
* Serde
  * [vityafx/serde-aux](https://github.com/vityafx/serde-aux/) - additional tools for using with the serde library. [![CI](https://github.com/vityafx/serde-aux/actions/workflows/ci.yml/badge.svg)](https://github.com/vityafx/serde-aux/actions/workflows/ci.yml) [![Crates badge](https://img.shields.io/crates/v/serde-aux.svg)](https://crates.io/crates/serde-aux)
* TOML
  * [toml-rs/toml](https://github.com/toml-rs/toml) — [![CI](https://github.com/toml-rs/toml/actions/workflows/ci.yml/badge.svg)](https://github.com/toml-rs/toml/actions/workflows/ci.yml) Stars:`219`.
* XML
  * [tafia/quick-xml](https://github.com/tafia/quick-xml) — High performance XML pull reader/writer [![build badge](https://api.travis-ci.org/tafia/quick-xml.svg?branch=master)](https://travis-ci.org/tafia/quick-xml) Stars:`782`.
  * [netvl/xml-rs](https://github.com/netvl/xml-rs) — A streaming XML library [![build badge](https://api.travis-ci.org/netvl/xml-rs.svg?branch=master)](https://travis-ci.org/netvl/xml-rs) Stars:`408`.
  * [shepmaster/sxd-document](https://github.com/shepmaster/sxd-document) — An XML library in Rust [![build badge](https://api.travis-ci.org/shepmaster/sxd-document.svg?branch=master)](https://travis-ci.org/shepmaster/sxd-document) Stars:`145`.
  * [media-io/yaserde](https://github.com/media-io/yaserde) — Yet Another Serializer/Deserializer specialized for XML [![build badge](https://api.travis-ci.org/media-io/yaserde.svg?branch=master)](https://travis-ci.org/media-io/yaserde) Stars:`133`.
  * [shepmaster/sxd-xpath](https://github.com/shepmaster/sxd-xpath) — An XPath library in Rust [![build badge](https://api.travis-ci.org/shepmaster/sxd-xpath.svg?branch=master)](https://travis-ci.org/shepmaster/sxd-xpath) Stars:`106`.
  * [Florob/RustyXML](https://github.com/Florob/RustyXML) — an XML parser written in Rust [![build badge](https://api.travis-ci.org/Florob/RustyXML.svg?branch=master)](https://travis-ci.org/Florob/RustyXML) Stars:`95`.
* YAML
  * [dtolnay/serde-yaml](https://github.com/dtolnay/serde-yaml) [[serde\_yaml](https://crates.io/crates/serde_yaml)] — YAML support for [Serde](https://github.com/serde-rs/serde) framework [![build](https://img.shields.io/github/workflow/status/dtolnay/serde-yaml/CI/master)](https://github.com/dtolnay/serde-yaml/actions?query=branch%3Amaster) Stars:`636`.
  * [chyh1990/yaml-rust](https://github.com/chyh1990/yaml-rust) — The missing YAML 1.2 implementation for Rust. [![build badge](https://api.travis-ci.org/chyh1990/yaml-rust.svg?branch=master)](https://travis-ci.org/chyh1990/yaml-rust) Stars:`532`.
  * [vitiral/stfu8](https://github.com/vitiral/stfu8) [[stfu8](https://crates.io/crates/stfu8)] — Sorta Text Format in UTF-8 [![build badge](https://api.travis-ci.org/vitiral/stfu8.svg?branch=master)](https://travis-ci.org/vitiral/stfu8) Stars:`18`.

### Filesystem

[[filesystem](https://crates.io/keywords/filesystem)]
* Operations
  * [webdesus/fs_extra](https://github.com/webdesus/fs_extra) — expanding opportunities standard library std::fs and std::io [![build badge](https://api.travis-ci.org/webdesus/fs_extra.svg?branch=master)](https://travis-ci.org/webdesus/fs_extra) Stars:`158`.
  * [vitiral/path_abs](https://github.com/vitiral/path_abs) [[path_abs](https://crates.io/crates/path_abs)] — Absolute serializable path types and associated methods. [![build badge](https://api.travis-ci.org/vitiral/path_abs.svg?branch=master)](https://travis-ci.org/webdesus/fs_extr://travis-ci.org/vitiral/path_abs) Stars:`45`.
  * [pop-os/sys-mount](https://github.com/pop-os/sys-mount) [[sys-mount](https://crates.io/crates/sys-mount)] — High level abstraction for the `mount` / `umount2` system calls. Stars:`28`.
  * [pop-os/dbus-udisks2](https://github.com/pop-os/dbus-udisks2) [[dbus-udisks2](https://crates.io/crates/dbus-udisks2)] - UDisks2 DBus API Stars:`13`.
* Temporary Files
  * [zboxfs/zbox](https://github.com/zboxfs/zbox) [[zbox](https://crates.io/crates/zbox)] — Zero-details, privacy-focused embeddable file system. [![build badge](https://api.travis-ci.org/zboxfs/zbox.svg?branch=master)](https://travis-ci.org/zboxfs/zbox) Stars:`1.3K`.
  * [Stebalien/tempfile](https://github.com/Stebalien/tempfile) — temporary file library [![build badge](https://api.travis-ci.org/Stebalien/tempfile.svg?branch=master)](https://travis-ci.org/Stebalien/tempfile) Stars:`734`.
  * [Stebalien/xattr](https://github.com/Stebalien/xattr) [[xattr](https://crates.io/crates/xattr)] — list and manipulate unix extended file attributes [![build badge](https://api.travis-ci.org/Stebalien/xattr.svg?branch=master)](https://travis-ci.org/Stebalien/xattr) Stars:`31`.

### Functional Programming

[[functional programming](https://crates.io/keywords/fp)]
* Prelude
  * [JasonShin/fp-core.rs](https://github.com/JasonShin/fp-core.rs) — A library for functional programming in Rust Stars:`1.1K`.
  * [myrrlyn/tap](https://github.com/myrrlyn/tap) - Suffix-Position Pipeline Behavior Stars:`203`.

### Game development

See also [Are we game yet?](https://arewegameyet.rs)
* Allegro
* [Awesome wgpu](https://github.com/rofrol/awesome-wgpu) — A curated list of wgpu code and resources Stars:`269`.
* [Awesome Quads](https://github.com/ozkriff/awesome-quads) — A curated list of links to miniquad/macroquad-related code & resources Stars:`103`.
  * [SiegeLord/RustAllegro](https://github.com/SiegeLord/RustAllegro) — [Allegro 5](https://liballeg.org/) bindings [![build badge](https://api.travis-ci.org/SiegeLord/RustAllegro.svg?branch=master)](https://travis-ci.org/SiegeLord/RustAllegro) Stars:`79`.
* bracket-lib (previously RLTK)
  * [bracket-lib](https://github.com/amethyst/bracket-lib) [[bracket-lib](https://crates.io/crates/bracket-lib)] - The Roguelike Toolkit (RLTK), implemented for Rust. [![Rust](https://github.com/amethyst/bracket-lib/actions/workflows/rust.yml/badge.svg)](https://github.com/amethyst/bracket-lib/actions/workflows/rust.yml) Stars:`1.1K`.
* Challonge
  * [vityafx/challonge-rs](https://github.com/vityafx/challonge-rs) [[challonge](https://crates.io/crates/challonge)] — Client library for the Challonge REST API. Helps to organize tournaments. [![CI](https://github.com/vityafx/challonge-rs/actions/workflows/ci.yml/badge.svg)](https://github.com/vityafx/challonge-rs/actions/workflows/ci.yml) Stars:`2`.
* Corange
  * [lucidscape/corange-rs](https://github.com/lucidscape/corange-rs) — [Corange](https://github.com/orangeduck/Corange) bindings Stars:`43`.
* Entity-Component Systems (ECS)
  * [amethyst/specs](https://github.com/amethyst/specs) — Specs Parallel ECS [![build badge](https://api.travis-ci.org/amethyst/specs.svg?branch=master)](https://travis-ci.org/amethyst/specs) Stars:`2.1K`.
  * [legion](https://github.com/amethyst/legion) — A feature rich high performance ECS library with minimal boilerplate [![build badge](https://github.com/amethyst/legion/workflows/CI/badge.svg?branch=master)](https://github.com/amethyst/legion/actions) Stars:`1.4K`.
* Game Engines
  * [Bevy](https://github.com/bevyengine/bevy) is a refreshingly simple data-driven game engine built in Rust. - [![Crates.io](https://img.shields.io/crates/v/bevy.svg)](https://crates.io/crates/bevy) Stars:`19.6K`.
[![Crates.io](https://img.shields.io/crates/d/bevy.svg)](https://crates.io/crates/bevy)
  * [ggez](https://github.com/ggez/ggez) — A lightweight game framework for making 2D games with minimum friction - [![Crates.io](https://img.shields.io/crates/v/ggez.svg)](https://crates.io/crates/ggez) [![license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/ggez/ggez/blob/master/LICENSE) [![Crates.io](https://img.shields.io/crates/d/ggez.svg)](https://crates.io/crates/ggez) Stars:`3.5K`.
  * [godot-rust/godot-rust](https://github.com/godot-rust/godot-rust) [[gdnative](https://crates.io/crates/gdnative)] - Rust bindings to the Godot game engine [![CI](https://github.com/godot-rust/godot-rust/actions/workflows/full-ci.yml/badge.svg)](https://github.com/godot-rust/godot-rust/actions/workflows/full-ci.yml) Stars:`3.0K`.
  * [Rust-SDL2/rust-sdl2](https://github.com/Rust-SDL2/rust-sdl2) — SDL2 bindings [![build badge](https://api.travis-ci.org/Rust-SDL2/rust-sdl2.svg?branch=master)](https://travis-ci.org/Rust-SDL2/rust-sdl2) Stars:`2.2K`.
  * [deltaphc/raylib-rs](https://github.com/deltaphc/raylib-rs) [[raylib](https://crates.io/crates/raylib)] — Rust bindings for raylib Stars:`466`.
  * [Unrust](https://github.com/unrust/unrust) — unrust — A pure rust based (webgl 2.0 / native) game engine Stars:`366`.
  * [oxidator](https://github.com/Ruddle/oxidator) — A real time strategy game/engine written with Rust and WebGPU Stars:`255`.
  * [brson/rust-sdl](https://github.com/brson/rust-sdl) — SDL1 bindings [![build badge](https://api.travis-ci.org/brson/rust-sdl.svg?branch=master)](https://travis-ci.org/brson/rust-sdl) Stars:`173`.
* [Godot](https://godotengine.org/)
  * [Piston](https://www.piston.rs/) — [![Crates.io](https://img.shields.io/crates/v/piston.svg?style=flat-square)](https://crates.io/crates/piston) [![Crates.io](https://img.shields.io/crates/l/piston.svg)](https://github.com/PistonDevelopers/piston/blob/master/LICENSE) [![Crates.io](https://img.shields.io/crates/d/piston.svg)](https://crates.io/crates/piston)
* [Raylib](https://www.raylib.com/)
  * [Fyrox](https://fyrox.rs/) — Rust Game engine 3D [![Crates.io](https://img.shields.io/crates/v/fyrox.svg)](https://crates.io/crates/fyrox) [![license](https://img.shields.io/crates/l/fyrox.svg)](https://github.com/FyroxEngine/Fyrox/blob/master/LICENSE.md) [![Crates.io](https://img.shields.io/crates/d/fyrox.svg)](https://crates.io/crates/fyrox)
* [SDL](http://www.libsdl.org/) [[sdl](https://crates.io/keywords/sdl)]
  * [Kiss3d](http://kiss3d.org) — A Keep It Simple, Stupid 3d graphics engine written with Rust [![Crates.io](https://img.shields.io/crates/d/kiss3d.svg)](https://crates.io/crates/kiss3d)
* SFML
  * [jeremyletang/rust-sfml](https://github.com/jeremyletang/rust-sfml) — [SFML](https://www.sfml-dev.org/) bindings Stars:`558`.
* Tcod-rs
  * [tomassedovic/tcod-rs](https://github.com/tomassedovic/tcod-rs) — Libtcod bindings for Rust. Stars:`227`.
  * Warning: Not maintained anymore
* Toornament-rs
  * [vityafx/toornament-rs](https://github.com/vityafx/toornament-rs) - Toornament.com API bindings for Rust. [![CI](https://github.com/vityafx/toornament-rs/actions/workflows/ci.yml/badge.svg)](https://github.com/vityafx/toornament-rs/actions/workflows/ci.yml) [![Crates badge](https://img.shields.io/crates/v/toornament.svg)](https://crates.io/crates/toornament) Stars:`3`.
* Victorem
  * [VictoremWinbringer/Victorem](https://github.com/VictoremWinbringer/Victorem) [[Victorem](https://crates.io/crates/Victorem)] — Easy UDP Game Server and UDP Client framework for creating simple 2D and 3D online game prototype [![build badge](https://api.travis-ci.org/VictoremWinbringer/Victorem.svg?branch=master)](https://travis-ci.org/VictoremWinbringer/Victorem) Stars:`26`.

### Geospatial

[[geo](https://crates.io/keywords/geo), [gis](https://crates.io/keywords/gis)]

* [rust-reverse-geocoder](https://github.com/gx0r/rrgeo) — A fast, offline reverse geocoder in Rust, inspired by [thampiman/reverse-geocoder](https://github.com/thampiman/reverse-geocoder) Stars:`80`.
* [DaveKram/coord_transforms](https://github.com/DaveKram/coord_transforms) [[coord_transforms](https://crates.io/crates/coord_transforms)] — coordinate transformations (2-d, 3-d, and geospatial) [![build badge](https://api.travis-ci.org/DaveKram/coord_transforms.svg?branch=master)](https://travis-ci.org/DaveKram/coord_transforms) Stars:`25`.
* [vlopes11/geomorph](https://github.com/vlopes11/geomorph) [[geomorph](https://crates.io/crates/geomorph)] — conversion between UTM, LatLon and MGRS coordinates [![build badge](https://api.travis-ci.org/vlopes11/geomorph.svg?branch=master)](https://travis-ci.org/vlopes11/geomorph) Stars:`9`.
* [Georust](https://github.com/georust) — geospatial tools and libraries written in Rust

### Graphics

[[graphics](https://crates.io/keywords/graphics)]

* Font
* [gfx-rs/wgpu](https://github.com/gfx-rs/wgpu) - Native WebGPU implementation based on gfx-hal. [![build badge](https://github.com/gfx-rs/wgpu/workflows/CI/badge.svg?branch=master)](https://github.com/gfx-rs/wgpu/actions) Stars:`6.0K`.
* [gfx-rs/gfx](https://github.com/gfx-rs/gfx) — A high-performance, bindless graphics API for Rust. [![build badge](https://api.travis-ci.org/gfx-rs/gfx.svg?branch=master)](https://travis-ci.org/gfx-rs/gfx) Stars:`5.2K`.
  * [redox-os/rusttype](https://github.com/redox-os/rusttype) — A pure Rust alternative to libraries like FreeType [![build badge](https://api.travis-ci.org/redox-os/rusttype.svg?branch=master)](https://travis-ci.org/redox-os/rusttype) Stars:`565`.
  * [RazrFalcon/rustybuzz](https://github.com/RazrFalcon/rustybuzz) - An incremental harfbuzz port to Rust [![build badge](https://api.travis-ci.org/RazrFalcon/rustybuzz.svg?branch=master)](https://travis-ci.org/RazrFalcon/rustybuzz) Stars:`299`.
* OpenGL [[opengl](https://crates.io/keywords/opengl)]
  * [glium/glium](https://github.com/glium/glium) — safe OpenGL wrapper for the Rust language. [![build badge](https://api.travis-ci.org/glium/glium.svg?branch=master)](https://travis-ci.org/glium/glium) Stars:`3.1K`.
  * [brendanzab/gl-rs](https://github.com/brendanzab/gl-rs) — [![build badge](https://api.travis-ci.org/brendanzab/gl-rs.svg?branch=master)](https://travis-ci.org/brendanzab/gl-rs) Stars:`619`.
  * [PistonDevelopers/glfw-rs](https://github.com/PistonDevelopers/glfw-rs) — [![build badge](https://api.travis-ci.org/PistonDevelopers/glfw-rs.svg?branch=master)](https://travis-ci.org/PistonDevelopers/glfw-rs) Stars:`530`.
  * [glutin](https://crates.io/crates/glutin) — Rust alternative to [GLFW](https://www.glfw.org/) [![build badge](https://api.travis-ci.org/rust-windowing/glutin.svg?branch=master)](https://travis-ci.org/rust-windowing/glutin)
  * [Kiss3d](http://kiss3d.org) — draw simple geometric figures and play with them with one-liners [![build badge](https://api.travis-ci.org/sebcrozet/kiss3d.svg?branch=master)](https://api.travis-ci.org/repositories/sebcrozet/kiss3d)
* PDF
  * [vulkano](https://github.com/vulkano-rs/vulkano) [[vulkano](https://crates.io/crates/vulkano)] — [![build badge](https://api.travis-ci.org/vulkano-rs/vulkano.svg?branch=master)](https://travis-ci.org/vulkano-rs/vulkano) Stars:`3.5K`.
  * [J-F-Liu/lopdf](https://github.com/J-F-Liu/lopdf) — PDF document manipulation [![build badge](https://api.travis-ci.org/J-F-Liu/lopdf.svg?branch=master)](https://travis-ci.org/J-F-Liu/lopdf) Stars:`1.0K`.
  * [fschutt/printpdf](https://github.com/fschutt/printpdf) — PDF writing library [![build badge](https://api.travis-ci.org/fschutt/printpdf.svg?branch=master)](https://travis-ci.org/fschutt/printpdf) Stars:`570`.
  * [WASM-PDF](https://github.com/jussiniinikoski/wasm-pdf) – Generates PDF files with JavaScript and WASM (WebAssembly) [![build badge](https://api.travis-ci.org/jussiniinikoski/wasm-pdf.svg?branch=master)](https://travis-ci.org/jussiniinikoski/wasm-pdf) Stars:`349`.
  * [kaj/rust-pdf](https://github.com/kaj/rust-pdf) — [![build badge](https://api.travis-ci.org/kaj/rust-pdf.svg?branch=master)](https://travis-ci.org/kaj/rust-pdf) Stars:`123`.
* [Vulkan](https://www.vulkan.org/) [[vulkan](https://crates.io/keywords/vulkan)]
  * [erupt](https://gitlab.com/Friz64/erupt) [[erupt](https://crates.io/crates/erupt)] — [![build badge](https://gitlab.com/Friz64/erupt/badges/main/pipeline.svg)](https://gitlab.com/Friz64/erupt/-/pipelines)

### GUI

[[gui](https://crates.io/keywords/gui)]

* [autopilot-rs/autopilot-rs](https://github.com/autopilot-rs/autopilot-rs) — A simple, cross-platform GUI automation library for Rust. Stars:`265`.
* Cocoa
* [tauri-apps/tauri](https://github.com/tauri-apps/tauri) — Build smaller, faster, and more secure desktop applications with a web frontend, powered by [WRY](https://github.com/tauri-apps/wry). [![test library](https://img.shields.io/github/workflow/status/tauri-apps/tauri/test%20library?label=test%20library)](https://github.com/tauri-apps/tauri/actions?query=workflow%3A%22test+library%22) Stars:`53.7K`.
* [ImGui](https://github.com/ocornut/imgui) Stars:`43.0K`.
* [iced-rs/iced](https://github.com/iced-rs/iced) [[iced](https://crates.io/crates/iced)] — A cross-platform GUI library for Rust focused on simplicity and type-safety. Inspired by Elm. Stars:`16.8K`.
* [emilk/egui](https://github.com/emilk/egui) - Simple, fast, and highly portable immediate mode GUI library for Rust. egui runs on the web, natively, and in your favorite game engine. [![Build Status](https://github.com/emilk/egui/workflows/CI/badge.svg)](https://github.com/emilk/egui/actions?workflow=CI) Stars:`11.6K`.
* [libui](https://github.com/andlabs/libui) Stars:`10.3K`.
* [Druid](https://github.com/linebender/druid) [[druid](https://crates.io/crates/druid)] — [Druid](https://linebender.org/druid/), a data-first Rust-native UI design toolkit. [![build badge](https://github.com/linebender/druid/workflows/.github/workflows/ci.yml/badge.svg)](https://github.com/linebender/druid/actions) Stars:`7.9K`.
* [Nuklear](https://github.com/Immediate-Mode-UI/Nuklear) Stars:`6.2K`.
* [fschutt/azul](https://github.com/fschutt/azul) — A free, functional, IMGUI-oriented GUI framework for rapid development of desktop applications written in Rust, supported by the Mozilla WebRender rendering engine. [![build badge](https://api.travis-ci.org/fschutt/azul.svg?branch=master)](https://travis-ci.org/fschutt/azul) Stars:`5.3K`.
* [DioxusLabs/dioxus](https://github.com/dioxuslabs/dioxus) - a portable, performant, and ergonomic framework for building cross-platform user interfaces in Rust. ![rust ci](https://github.com/dioxuslabs/dioxus/actions/workflows/main.yml/badge.svg) Stars:`5.2K`.
* [slint-ui/slint](https://github.com/slint-ui/slint) [[slint](https://crates.io/crates/slint)] — [Slint](https://slint-ui.com) is a toolkit to efficiently develop fluid graphical user interfaces for embedded devices and desktop applications. [![Build Status](https://github.com/slint-ui/slint/workflows/CI/badge.svg?branch=master)](https://github.com/slint-ui/slint/actions?query=workflow%3ACI) Stars:`5.1K`.
* [OrbTk](https://github.com/redox-os/orbtk) — The Orbital Widget Toolkit is a multi platform (G)UI toolkit using SDL2 [![Build and test](https://github.com/redox-os/orbtk/workflows/build/badge.svg?branch=develop)](https://github.com/redox-os/orbtk/actions) Stars:`3.7K`.
  * [relm](https://github.com/antoyo/relm) — Asynchronous, GTK+-based, GUI library, inspired by Elm [![build badge](https://api.travis-ci.org/antoyo/relm.svg?branch=master)](https://travis-ci.org/antoyo/relm) Stars:`2.2K`.
* [tauri-apps/wry](https://github.com/tauri-apps/wry) - Webview Rendering librarY. Stars:`2.1K`.
  * [flutter-rs](https://github.com/flutter-rs/flutter-rs) — Build flutter desktop app in dart & rust. Stars:`2.0K`.
  * [imgui-rs](https://github.com/imgui-rs/imgui-rs) — Rust bindings for ImGui [![Build Status](https://github.com/imgui-rs/imgui-rs/workflows/ci/badge.svg?branch=master)](https://github.com/imgui-rs/imgui-rs/actions) Stars:`2.0K`.
  * [fltk-rs](https://github.com/fltk-rs/fltk-rs) — FLTK Rust bindings [![Build](https://github.com/fltk-rs/fltk-rs/workflows/Build/badge.svg?branch=master)](https://github.com/fltk-rs/fltk-rs/actions) Stars:`966`.
  * [rust-native-ui/libui-rs](https://github.com/rust-native-ui/libui-rs) — libui bindings [![build badge](https://api.travis-ci.org/rust-native-ui/libui-rs.svg?branch=master)](https://travis-ci.org/rust-native-ui/libui-rs). Stars:`863`.
  * [sciter-sdk/rust-sciter](https://github.com/sciter-sdk/rust-sciter) — Sciter bindings [![build badge](https://ci.appveyor.com/api/projects/status/github/sciter-sdk/rust-sciter?svg=true)](https://ci.appveyor.com/project/sciter-sdk/rust-sciter) Stars:`755`.
* [emoon/rust_minifb](https://github.com/emoon/rust_minifb) — minifb is a cross-platform window setup with optional bitmap rendering. It also comes with easy mouse and keyboard input. Primarily designed for prototyping Stars:`746`.
  * [servo/core-foundation-rs](https://github.com/servo/core-foundation-rs) — [![build badge](https://api.travis-ci.com/servo/core-foundation-rs.svg?branch=master)](https://travis-ci.org/servo/core-foundation-rs) Stars:`662`.
* [ivanceras/sauron-native](https://github.com/ivanceras/sauron-native) - A truly native and cross platform GUI library. One unified code can be run as native GUI, Html Web and TUI. [![Build Status](https://api.travis-ci.com/ivanceras/sauron-native.svg?branch=master)](https://app.travis-ci.com/github/ivanceras/sauron-native) Stars:`626`.
  * [woboq/qmetaobject-rs](https://github.com/woboq/qmetaobject-rs) — Integrate Qml and Rust by building the QMetaObject at compile time. [![build badge](https://api.travis-ci.org/woboq/qmetaobject-rs.svg?branch=master)](https://travis-ci.org/woboq/qmetaobject-rs) Stars:`479`.
  * [cyndis/qmlrs](https://github.com/cyndis/qmlrs) — QtQuick bindings [![build badge](https://api.travis-ci.org/cyndis/qmlrs.svg?branch=master)](https://travis-ci.org/cyndis/qmlrs) Stars:`432`.
  * [gtk-rs/gtk3-rs](https://github.com/gtk-rs/gtk3-rs) - GTK3 binding for rust ![CI](https://github.com/gtk-rs/gtk3-rs/workflows/CI/badge.svg) Stars:`415`.
  * [Kiss-ui](https://github.com/KISS-UI/kiss-ui) — A simple UI framework built on IUP [![Build Status](https://api.travis-ci.org/cybergeek94/kiss-ui.svg?branch=master)](https://travis-ci.org/cybergeek94/kiss-ui) Stars:`342`.
  * [nuklear-rust](https://github.com/snuk182/nuklear-rust) — Rust bindings for Nuklear Stars:`326`.
* [saurvs/nfd-rs](https://github.com/saurvs/nfd-rs) — [nativefiledialog](https://github.com/mlabbe/nativefiledialog) bindings Stars:`151`.
* [rise-ui](https://github.com/rise-ui/rise) — Simple component-based cross-Platform GUI Toolkit for developing beautiful and user-friendly interfaces. Stars:`71`.
* [PistonDevelopers/conrod](https://github.com/PistonDevelopers/conrod/) — An easy-to-use, immediate-mode, 2D GUI library written entirely in Rust [![build badge](https://api.travis-ci.org/PistonDevelopers/conrod.svg?branch=master)](https://travis-ci.org/PistonDevelopers/conrod)
  * [rust-qt](https://github.com/rust-qt)
* [Qt](https://doc.qt.io)
* [Sciter](https://sciter.com/)
* [IUP](http://webserver2.tecgraf.puc-rio.br/iup/)
* [GTK+](https://www.gtk.org/) [[gtk](https://crates.io/keywords/gtk)]
* [Flutter](https://flutter.dev/)
* [FLTK](https://www.fltk.org/)

### Image processing

* [image-rs/image](https://github.com/image-rs/image) — Basic imaging processing functions and methods for converting to and from image formats [![build badge](https://api.travis-ci.org/image-rs/image.svg?branch=master)](https://travis-ci.org/image-rs/image) Stars:`3.4K`.
* [twistedfall/opencv-rust](https://github.com/twistedfall/opencv-rust) — Rust bindings for OpenCV [![build badge](https://api.travis-ci.org/twistedfall/opencv-rust.svg?branch=cv2)](https://travis-ci.org/twistedfall/opencv-rust) Stars:`1.1K`.
* [image-rs/imageproc](https://github.com/image-rs/imageproc) — An image processing library, based on the `image` library. [![Build Status](https://api.travis-ci.org/image-rs/imageproc.svg?branch=master)](https://travis-ci.org/image-rs/imageproc) Stars:`497`.
* [rust-cv/cv](https://github.com/rust-cv/cv) — Rust CV is a project to implement computer vision algorithms, abstractions, and systems in Rust. #[no_std] is supported where possible. ![build badge](https://github.com/rust-cv/cv/workflows/tests/badge.svg) Stars:`397`.
* [abonander/img_hash](https://github.com/abonander/img_hash) — Perceptual image hashing and comparison for equality and similarity. [![Build Status](https://api.travis-ci.org/abonander/img_hash.svg?branch=master)](https://travis-ci.org/abonander/img_hash) Stars:`261`.
* [teovoinea/steganography](https://github.com/teovoinea/steganography) [[steganography](https://crates.io/crates/steganography)] — A simple steganography library [![build badge](https://api.travis-ci.org/teovoinea/steganography.svg?branch=master)](https://travis-ci.org/teovoinea/steganography) Stars:`78`.

### Language specification

* [shnewto/bnf](https://github.com/shnewto/bnf) — A library for parsing Backus–Naur form context-free grammars. [![build badge](https://api.travis-ci.org/shnewto/bnf.svg?branch=master)](https://travis-ci.org/shnewto/bnf) Stars:`179`.

### Logging

[[log](https://crates.io/keywords/log)]

* [tokio-rs/tracing](https://github.com/tokio-rs/tracing) — An application level tracing framework for async-aware structured logging, error handling, metrics, and more [![Build Status](https://github.com/tokio-rs/tracing/workflows/CI/badge.svg?branch=master)](https://github.com/tokio-rs/tracing/actions?query=workflow%3ACI) Stars:`3.1K`.
* [rust-lang/log](https://github.com/rust-lang/log) — Logging implementation for Rust [![Build Status](https://api.travis-ci.org/rust-lang/log.svg?branch=master)](https://travis-ci.org/rust-lang/log) Stars:`1.5K`.
* [slog-rs/slog](https://github.com/slog-rs/slog) — Structured, composable logging for Rust [![Build Status](https://api.travis-ci.org/slog-rs/slog.svg?branch=master)](https://travis-ci.org/slog-rs/slog) Stars:`1.3K`.
* [estk/log4rs](https://github.com/estk/log4rs) — highly configurable logging framework modeled after Java's Logback and log4j libraries [![CircleCI](https://circleci.com/gh/estk/log4rs.svg?style=shield)](https://app.circleci.com/pipelines/github/estk/log4rs) Stars:`727`.
* [seanmonstar/pretty-env-logger](https://github.com/seanmonstar/pretty-env-logger) — A pretty, easy-to-use logger for Rust. [![Build Status](https://api.travis-ci.org/seanmonstar/pretty-env-logger.svg?branch=master)](https://travis-ci.org/seanmonstar/pretty-env-logger) Stars:`373`.
* [jesusprubio/leg](https://github.com/jesusprubio/leg) — Elegant print for lazy devs. Make your CLIs nicer with minimal effort. [![Build Status](https://github.com/jesusprubio/leg/workflows/CI/badge.svg)](https://github.com/jesusprubio/leg/actions/workflows/ci.yml) Stars:`202`.
* [rbatis/fast_log](https://github.com/rbatis/fast_log) — Rust async log High-performance asynchronous logging [![Build Status](https://api.travis-ci.com/rbatis/fast_log.svg?branch=master)](https://travis-ci.org/rbatis/fast_log) Stars:`89`.

### Macro

* cute
  * [mattgathu/cute](https://github.com/mattgathu/cute) — Macro for Python-esque list comprehensions in Rust. [![Build Status](https://api.travis-ci.org/mattgathu/cute.svg?branch=master)](https://travis-ci.org/tensorflow/rust) Stars:`303`.
* [Linq-in-Rust](https://github.com/StardustDL/Linq-in-Rust) - Macro and methods for C#-LINQ-like expressions. [![CI](https://github.com/StardustDL/Linq-in-Rust/workflows/CI/badge.svg?branch=master)](https://github.com/StardustDL/Linq-in-Rust/actions?query=workflow%3ACI) Stars:`88`.

### Markup language

* CommonMark
  * [raphlinus/pulldown-cmark](https://github.com/raphlinus/pulldown-cmark) — [CommonMark](https://commonmark.org/) parser in Rust Stars:`1.5K`.

### Mobile

* Android / iOS
  * [ivanschuetz/rust_android_ios](https://github.com/ivanschuetz/rust_android_ios) — An example of using a shared Rust lib for Android and iOS using rust-swig and cbindgen respectively. Stars:`189`.
* Generic
  * [Geal/rust_on_mobile](https://github.com/Geal/rust_on_mobile) Stars:`158`.
* iOS
  * [TimNN/cargo-lipo](https://github.com/TimNN/cargo-lipo) — A cargo lipo subcommand which automatically creates a universal library for use with your iOS application. [![build badge](https://api.travis-ci.org/TimNN/cargo-lipo.svg?branch=master)](https://travis-ci.org/TimNN/cargo-lipo) Stars:`424`.

### Network programming

* Bluetooth
  * [bluez/bluer](https://github.com/bluez/bluer) [[bluer](https://crates.io/crates/bluer)] — Official BlueZ bindings for Rust. [![build badge](https://github.com/bluez/bluer/actions/workflows/rust.yml/badge.svg?branch=master)](https://github.com/bluez/bluer/actions/workflows/rust.yml) Stars:`105`.
* CoAP
  * [Covertness/coap-rs](https://github.com/Covertness/coap-rs) — A [Constrained Application Protocol(CoAP)](https://datatracker.ietf.org/doc/html/rfc7252) library for Rust. [![build badge](https://api.travis-ci.org/Covertness/coap-rs.svg?branch=master)](https://travis-ci.org/Covertness/coap-rs) Stars:`167`.
* Docker
  * [fussybeaver/bollard](https://github.com/fussybeaver/bollard) — Docker daemon API in Rust Stars:`406`.
* FTP
  * [mattnenterprise/rust-ftp](https://github.com/mattnenterprise/rust-ftp) — an [FTP](https://en.wikipedia.org/wiki/File_Transfer_Protocol) client for Rust [![build badge](https://api.travis-ci.org/mattnenterprise/rust-ftp.svg?branch=master)](https://travis-ci.org/mattnenterprise/rust-ftp) Stars:`153`.
* gRPC
  * [tikv/grpc-rs](https://github.com/tikv/grpc-rs) — The gRPC library for Rust built on C Core library and futures [![Build Status](https://api.travis-ci.org/tikv/grpc-rs.svg?branch=master)](https://travis-ci.org/tikv/grpc-rs) Stars:`1.6K`.
* HTTP
  * [Hurl](https://github.com/Orange-OpenSource/hurl) — Run and test HTTP requests with plain text and libcurl [![CI](https://github.com/Orange-OpenSource/hurl/workflows/CI/badge.svg)](https://github.com/Orange-OpenSource/hurl/actions) Stars:`2.0K`.
* IPNetwork
  * [jesusprubio/online](https://github.com/jesusprubio/online) — Library to check your Internet connectivity [![CI](https://github.com/jesusprubio/online/actions/workflows/ci.yml/badge.svg)](https://github.com/jesusprubio/online/actions/workflows/ci.yml) Stars:`114`.
  * [candrew/netsim](https://github.com/canndrew/netsim) — A Rust library for network simulation and testing [![build badge](https://api.travis-ci.org/canndrew/netsim.svg?branch=master)](https://travis-ci.org/canndrew/netsim) Stars:`111`.
  * [achanda/ipnetwork](https://github.com/achanda/ipnetwork) — A library to work with IP networks in pure Rust [![build badge](https://api.travis-ci.org/achanda/ipnetwork.svg?branch=master)](https://travis-ci.org/achanda/ipnetwork) Stars:`95`.
* Low level
  * [tokio-rs/tokio](https://github.com/tokio-rs/tokio) — A network application framework for rapid development and highly scalable production deployments of clients and servers. Stars:`18.1K`.
  * [actix/actix](https://github.com/actix/actix) — Actor library for Rust [![build badge](https://api.travis-ci.org/actix/actix.svg?branch=master)](https://travis-ci.org/actix/actix) Stars:`7.5K`.
  * [smoltcp-rs/smoltcp](https://github.com/smoltcp-rs/smoltcp) — A standalone, event-driven TCP/IP stack that is designed for bare-metal, real-time systems [![build badge](https://api.travis-ci.org/smoltcp-rs/smoltcp.svg?branch=master)](https://travis-ci.org/smoltcp-rs/smoltcp) Stars:`2.7K`.
  * [libpnet/libpnet](https://github.com/libpnet/libpnet) — A cross-platform, low level networking [![build badge](https://api.travis-ci.org/libpnet/libpnet.svg?branch=master)](https://travis-ci.org/libpnet/libpnet) Stars:`1.7K`.
  * [dylanmckay/protocol](https://github.com/dylanmckay/protocol) — Custom TCP/UDP protocol definitions Stars:`150`.
* message-io
  * [lemunozm/message-io](https://github.com/lemunozm/message-io) — Event-driven message library to build network applications easy and fast. Supports TCP, UDP and WebSockets. [![build badge](https://img.shields.io/github/workflow/status/lemunozm/message-io/message-io%20ci)](https://github.com/lemunozm/message-io/actions?query=workflow%3A%22message-io+ci%22) Stars:`806`.
* MQTT
  * [bytebeamio/rumqtt](https://github.com/bytebeamio/rumqtt) - A library for developers to build applications that communicate with the [MQTT protocol](https://mqtt.org) over TCP and WebSockets, with or without TLS. [![Build and Test](https://github.com/bytebeamio/rumqtt/actions/workflows/build.yml/badge.svg)](https://github.com/bytebeamio/rumqtt/actions/workflows/build.yml) Stars:`728`.
* NanoMsg
  * [thehydroimpulse/nanomsg.rs](https://github.com/thehydroimpulse/nanomsg.rs) — [nanomsg](https://nanomsg.org/) bindings [![build badge](https://api.travis-ci.org/thehydroimpulse/nanomsg.rs.svg?branch=master)](https://travis-ci.org/thehydroimpulse/nanomsg.rs) Stars:`372`.
* NATS
  * [nats-io/nats.rs](https://github.com/nats-io/nats.rs) — Rust client for NATS, the cloud native messaging system. [![Build Status](https://github.com/nats-io/nats.rs/workflows/Rust/badge.svg?branch=master)](https://github.com/nats-io/nats.rs/actions) Stars:`615`.
* Nng
  * [neachdainn/nng-rs](https://gitlab.com/neachdainn/nng-rs) [[Nng](https://crates.io/crates/nng)] — [Nng (nanomsg v2)](https://nng.nanomsg.org/index.html) bindings [![build badge](https://gitlab.com/neachdainn/nng-rs/badges/master/pipeline.svg)](https://gitlab.com/neachdainn/nng-rs/-/pipelines)
* NNTP
  * [mattnenterprise/rust-nntp](https://github.com/mattnenterprise/rust-nntp) [[nntp](https://crates.io/crates/nntp)] — an [NNTP](https://en.wikipedia.org/wiki/Network_News_Transfer_Protocol) client for Rust [![build badge](https://api.travis-ci.org/mattnenterprise/rust-nntp.svg?branch=master)](https://travis-ci.org/mattnenterprise/rust-nntp) Stars:`13`.
* P2P
  * [libp2p/rust-libp2p](https://github.com/libp2p/rust-libp2p) — The Rust Implementation of libp2p networking stack. [![Circle CI](https://circleci.com/gh/libp2p/rust-libp2p.svg?style=svg)](https://app.circleci.com/pipelines/github/libp2p/rust-libp2p) Stars:`2.8K`.
* POP3
  * [mattnenterprise/rust-pop3](https://github.com/mattnenterprise/rust-pop3) [[pop3](https://crates.io/crates/pop3)] — A [POP3](https://en.wikipedia.org/wiki/Post_Office_Protocol) client for Rust [![build badge](https://api.travis-ci.org/mattnenterprise/rust-pop3.svg?branch=master)](https://travis-ci.org/mattnenterprise/rust-pop3) Stars:`25`.
* QUIC
  * [cloudflare/quiche](https://github.com/cloudflare/quiche) — cloudflare implementation of the QUIC transport protocol and HTTP/3 ![build](https://img.shields.io/github/workflow/status/cloudflare/quiche/Stable) Stars:`6.8K`.
  * [quinn-rs/quinn](https://github.com/quinn-rs/quinn) — Futures-based QUIC implementation in Rust [![build badge](https://dev.azure.com/dochtman/Projects/_apis/build/status/Quinn?branchName=master)](https://dev.azure.com/dochtman/Projects/_build) Stars:`2.4K`.
  * [mozilla/neqo](https://github.com/mozilla/neqo) — an Implementation of QUIC written in Rust Stars:`1.5K`.
* Raknet
  * [b23r0/rust-raknet](https://github.com/b23r0/rust-raknet) — RakNet Protocol implementation by Rust [![Build Status](https://img.shields.io/github/workflow/status/b23r0/rust-raknet/Rust)](https://github.com/b23r0/rust-raknet/actions/workflows/rust.yml) Stars:`150`.
* RPC
  * [smallnest/rpcx-rs](https://github.com/smallnest/rpcx-rs) — A RPC library for Rust for developing microservices in easy and simple way. [![build badge](https://api.travis-ci.org/smallnest/rpcx-rs.svg?branch=master)](https://travis-ci.org/smallnest/rpcx-rs) Stars:`102`.
  * [ENQT-GmbH/remoc](https://github.com/ENQT-GmbH/remoc) [[remoc](https://crates.io/crates/remoc)] - Remoc provides channels (broadcast, mpsc, oneshot, watch) similar to Tokio's and trait calling over any remote transport. [![build badge](https://github.com/ENQT-GmbH/remoc/actions/workflows/rust.yml/badge.svg?branch=master)](https://github.com/ENQT-GmbH/remoc/actions/workflows/rust.yml) Stars:`94`.
* Socket.io
  * [1c3t3a/rust-socketio](https://github.com/1c3t3a/rust-socketio) [[rust_socketio](https://crates.io/crates/rust_socketio)] — an implementation of a [socket.io](https://socket.io) client written in Rust. [![build badge](https://github.com/1c3t3a/rust-socketio/actions/workflows/build.yml/badge.svg)](https://github.com/1c3t3a/rust-socketio/actions/workflows/build.yml) Stars:`169`.
* SSH
  * [alexcrichton/ssh2-rs](https://github.com/alexcrichton/ssh2-rs) — [libssh2](https://www.libssh2.org/) bindings [![build badge](https://api.travis-ci.com/alexcrichton/ssh2-rs.svg?branch=master)](https://travis-ci.org/alexcrichton/ssh2-rs) Stars:`350`.
  * [Thrussh](https://pijul.org/thrussh) [[thrussh](https://crates.io/crates/thrussh)] — an SSH library written from scratch in Rust, backed by [libsodium](https://doc.libsodium.org/)
* Stomp
  * [zslayton/stomp-rs](https://github.com/zslayton/stomp-rs) — A [STOMP 1.2](http://stomp.github.io/stomp-specification-1.2.html) client implementation in Rust [![build badge](https://api.travis-ci.org/zslayton/stomp-rs.svg?branch=master)](https://travis-ci.org/zslayton/stomp-rs) Stars:`83`.
* ZeroMQ
  * [erickt/rust-zmq](https://github.com/erickt/rust-zmq) — [ZeroMQ](https://zeromq.org/) bindings [![build badge](https://api.travis-ci.org/erickt/rust-zmq.svg?branch=master)](https://travis-ci.org/erickt/rust-zmq) Stars:`710`.

### Parsing

  * [Geal/nom](https://github.com/Geal/nom) — parser combinator library [![build badge](https://api.travis-ci.org/Geal/nom.svg?branch=master)](https://travis-ci.org/Geal/nom) Stars:`7.2K`.
  * [pest-parser/pest](https://github.com/pest-parser/pest) — The Elegant Parser [![Build Status](https://api.travis-ci.org/pest-parser/pest.svg?branch=master)](https://travis-ci.org/pest-parser/pest) Stars:`3.4K`.
  * [lalrpop/lalrpop](https://github.com/lalrpop/lalrpop) — LR(1) parser generator for Rust [![Build status](https://api.travis-ci.org/lalrpop/lalrpop.svg?branch=master)](https://travis-ci.org/lalrpop/lalrpop) Stars:`2.3K`.
  * [kevinmehall/rust-peg](https://github.com/kevinmehall/rust-peg) — Parsing Expression Grammar (PEG) parser generator Stars:`1.2K`.
  * [Marwes/combine](https://github.com/Marwes/combine) — parser combinator library [![build badge](https://api.travis-ci.org/Marwes/combine.svg?branch=master)](https://travis-ci.org/Marwes/combine) Stars:`1.1K`.
  * [m4rw3r/chomp](https://github.com/m4rw3r/chomp) – A fast monadic-style parser combinator [![build badge](https://api.travis-ci.org/m4rw3r/chomp.svg?branch=master)](https://travis-ci.org/m4rw3r/chomp) Stars:`228`.
  * [ptal/oak](https://github.com/ptal/oak) — A typed PEG parser generator (compiler plugin) Stars:`138`.
  * [freestrings/jsonpath](https://github.com/freestrings/jsonpath) — [JsonPath](https://goessner.net/articles/JsonPath/) engine written in Rust. Webassembly and Javascript support too [![Build Status](https://api.travis-ci.org/freestrings/jsonpath.svg?branch=master)](https://travis-ci.org/freestrings/jsonpath) Stars:`90`.
  * [Folyd/robotstxt](https://github.com/Folyd/robotstxt) - A native Rust port of Google's robots.txt parser and matcher C++ library Stars:`72`.
  * [s-panferov/queryst](https://github.com/s-panferov/queryst) — A query string parsing library for Rust inspired by [gs](https://github.com/ljharb/qs#readme) Stars:`63`.
  * [nrc/zero](https://github.com/nrc/zero) [[zero](https://crates.io/crates/zero/)] — zero-allocation parsing of binary data Stars:`45`.
  * [replicadse/wavefront_rs](https://github.com/replicadse/wavefront_rs) — A parser for the Wavefront OBJ format. [![crates.io](https://img.shields.io/crates/v/wavefront_rs.svg)](https://crates.io/crates/wavefront_rs) [![crates.io](https://img.shields.io/crates/d/wavefront_rs?label=crates.io%20downloads)](https://crates.io/crates/wavefront_rs) [![build badge](https://github.com/replicadse/wavefront_rs/workflows/pipeline/badge.svg?branch=master)](https://github.com/replicadse/wavefront_rs/actions) Stars:`2`.
  * [softdevteam/grmtools](https://github.com/softdevteam/grmtools/) - A LR parser with better error correction

### Peripherals

* Serial Port
  * [Susurrus/serialport-rs](https://gitlab.com/susurrus/serialport-rs) [[serialport](https://crates.io/crates/serialport)] — A cross-platform library that provides access to a serial port

### Platform specific

* Cross-platform
  * [svartalf/rust-battery](https://crates.io/crates/battery) — Cross-platform information about the notebook batteries [![build badge](https://api.travis-ci.org/svartalf/rust-battery.svg?branch=master)](https://travis-ci.org/svartalf/rust-battery)
  * [vityafx/thread-priority](https://github.com/vityafx/thread-priority/) - Simple, crossplatform thread priority management. [![CI](https://github.com/vityafx/thread-priority/actions/workflows/ci.yml/badge.svg)](https://github.com/vityafx/thread-priority/actions/workflows/ci.yml) [![Crates badge](https://img.shields.io/crates/v/thread-priority.svg)](https://crates.io/crates/thread-priority)
* FreeBSD
  * [fubarnetes/libjail-rs](https://github.com/fubarnetes/libjail-rs/) [[jail](https://crates.io/crates/jail)] — Rust implementation of a FreeBSD jail library
* Linux
  * [hannobraun/inotify-rs](https://github.com/hannobraun/inotify-rs) — [inotify](https://en.wikipedia.org/wiki/Inotify) bindings [![Rust](https://github.com/hannobraun/inotify-rs/actions/workflows/rust.yml/badge.svg)](https://github.com/hannobraun/inotify-rs/actions/workflows/rust.yml) Stars:`219`.
  * [arvancloud/nginx-rs](https://github.com/arvancloud/nginx-rs) — [Nginx](https://www.nginx.com) bindings [![build badge](https://api.travis-ci.org/arvancloud/nginx-rs.svg?branch=master)](https://travis-ci.org/arvancloud/nginx-rs) Stars:`103`.
  * [yaa110/rust-iptables](https://github.com/yaa110/rust-iptables) [[iptables](https://crates.io/crates/iptables)] — [iptables](https://www.netfilter.org/projects/iptables/index.html) bindings [![build badge](https://api.travis-ci.org/yaa110/rust-iptables.svg?branch=master)](https://travis-ci.org/yaa110/rust-iptables) Stars:`62`.
  * [pop-os/distinst](https://github.com/pop-os/distinst/) — Linux distribution installer
* Unix-like
  * [nix-rust/nix](https://github.com/nix-rust/nix) — Unix-like API bindings [![Cirrus Build Status](https://api.cirrus-ci.com/github/nix-rust/nix.svg)](https://cirrus-ci.com/github/nix-rust/nix) Stars:`2.0K`.
  * [zargony/fuse-rs](https://github.com/zargony/fuse-rs) — [FUSE](https://github.com/libfuse/libfuse) bindings Stars:`893`.
  * [rustix](https://github.com/bytecodealliance/rustix) — Safe Rust bindings to POSIX/Unix/Linux/Winsock2 syscalls [![Actions Status](https://github.com/bytecodealliance/rustix/workflows/CI/badge.svg)](https://github.com/bytecodealliance/rustix/actions?query=workflow%3ACI) Stars:`656`.
* Windows
  * [retep998/winapi-rs](https://github.com/retep998/winapi-rs) — Windows API bindings [![Rust](https://github.com/retep998/winapi-rs/actions/workflows/rust.yml/badge.svg?branch=dev)](https://github.com/retep998/winapi-rs/actions/workflows/rust.yml) Stars:`1.5K`.

### Scripting

[[scripting](https://crates.io/keywords/scripting)]

* [gluon-lang/gluon](https://github.com/gluon-lang/gluon) —  A small, statically-typed, functional programming language Stars:`2.7K`.
* [rhaiscript/rhai](https://github.com/rhaiscript/rhai) — A tiny and fast embedded scripting language resembling a combination of JavaScript and Rust [![build badge](https://github.com/rhaiscript/rhai/workflows/Build/badge.svg)](https://github.com/rhaiscript/rhai/actions) Stars:`2.3K`.
* [PistonDevelopers/dyon](https://github.com/PistonDevelopers/dyon) — A rusty dynamically typed scripting language Stars:`1.5K`.
* [mun](https://github.com/mun-lang/mun) — A compiled, statically-typed scripting language with first class hot reloading support [![build badge](https://api.travis-ci.org/mun-lang/mun.svg?branch=master)](https://travis-ci.org/mun-lang/mun) Stars:`1.4K`.
* [metacall/core](https://github.com/metacall/core) [[metacall](https://crates.io/crates/metacall)] — Cross-platform Polyglot Runtime which supports NodeJS, JavaScript, TypeScript, Python, Ruby, C#, Wasm, Java, Cobol and more. [![build badge](https://gitlab.com/metacall/core/badges/master/pipeline.svg)](https://gitlab.com/metacall/core) Stars:`1.1K`.
* [rune-rs/rune](https://github.com/rune-rs/rune) — An embeddable dynamic programming language for Rust Stars:`1.1K`.
* [murarth/ketos](https://github.com/murarth/ketos) — A Lisp dialect functional programming language serving as a scripting and extension language for rust Stars:`720`.
* [fleabitdev/gamelisp](https://github.com/fleabitdev/glsp) — A LISP-lisk scripting language for Rust game development Stars:`347`.
* [duckscript](https://crates.io/crates/duckscript) — [Simple, extendable and embeddable scripting language.](https://github.com/sagiegurari/duckscript) [![build badge](https://github.com/sagiegurari/duckscript/workflows/CI/badge.svg?branch=master)](https://github.com/sagiegurari/duckscript/actions) Stars:`343`.
* [KusionStack/KCLVM](https://github.com/KusionStack/KCLVM) - A constraint-based record & functional language mainly used in configuration and policy scenarios. Stars:`126`.

### Simulation

[[simulation](https://crates.io/keywords/simulation)]

* [nyx-space](https://crates.io/crates/nyx-space) - High fidelity, fast, reliable and validated astrodynamical toolkit library, used for spacecraft mission design and orbit determination [![Build Status](https://gitlab.com/nyx-space/nyx/badges/master/pipeline.svg)](https://gitlab.com/nyx-space/nyx/-/pipelines)

### Task scheduling

* [delay-timer](https://github.com/BinChengZhao/delay-timer) — Time-manager of delayed tasks. Like crontab, but asynchronous tasks are possible. Stars:`254`.
[![Build](https://github.com/BinChengZhao/delay-timer/actions/workflows/rust.yml/badge.svg)](
https://github.com/BinChengZhao/delay-timer/actions)

### Template engine

* Handlebars
  * [sunng87/handlebars-rust](https://github.com/sunng87/handlebars-rust) — Handlebars template engine with inheritance, custom helper support. [![build badge](https://api.travis-ci.org/sunng87/handlebars-rust.svg?branch=master)](https://travis-ci.org/sunng87/handlebars-rust) Stars:`893`.
  * [botika/yarte](https://github.com/botika/yarte) — Yarte stands for **Y**et **A**nother **R**ust **T**emplate **E**ngine, is the fastest template engine. [![Build Status](https://api.travis-ci.org/botika/yarte.svg?branch=master)](https://travis-ci.org/botika/yarte) Stars:`241`.
* HTML
  * [Keats/tera](https://github.com/Keats/tera) — template engine based on Jinja2 and the Django template language. [![Actions Status](https://github.com/Keats/tera/workflows/ci/badge.svg?branch=master)](https://github.com/Keats/tera/actions) Stars:`2.3K`.
  * [djc/askama](https://github.com/djc/askama) — template rendering engine based on Jinja [![build badge](https://api.travis-ci.org/djc/askama.svg?branch=master)](https://travis-ci.org/djc/askama) Stars:`1.9K`.
  * [lambda-fairy/maud](https://github.com/lambda-fairy/maud) — compile-time HTML templates [![build badge](https://api.travis-ci.org/lambda-fairy/maud.svg?branch=master)](https://travis-ci.org/lambda-fairy/maud) Stars:`1.4K`.
  * [kaj/ructe](https://github.com/kaj/ructe) — HTML template system for Rust [![build badge](https://api.travis-ci.org/kaj/ructe.svg?branch=master)](https://travis-ci.org/kaj/ructe) Stars:`319`.
  * [Stebalien/horrorshow-rs](https://github.com/Stebalien/horrorshow-rs) — compile-time HTML templates [![build badge](https://api.travis-ci.org/Stebalien/horrorshow-rs.svg?branch=master)](https://travis-ci.org/Stebalien/horrorshow-rs) Stars:`267`.
* Mustache
  * [rustache/rustache](https://github.com/rustache/rustache) — [![build badge](https://api.travis-ci.org/rustache/rustache.svg?branch=master)](https://travis-ci.org/rustache/rustache) Stars:`208`.

### Text processing

* [rust-lang/regex](https://github.com/rust-lang/regex) — Regular expressions (RE2 style) [![build badge](https://api.travis-ci.com/rust-lang/regex.svg?branch=master)](https://travis-ci.org/rust-lang/regex) Stars:`2.5K`.
* [greyblake/whatlang-rs](https://github.com/greyblake/whatlang-rs) — Natural language detection library based on trigrams [![build badge](https://api.travis-ci.org/greyblake/whatlang-rs.svg?branch=master)](https://travis-ci.org/greyblake/whatlang-rs) Stars:`784`.
* [mgeisler/textwrap](https://github.com/mgeisler/textwrap) [[textwrap](https://crates.io/crates/textwrap)] — Word wrap text (with support for hyphenation) [![build badge](https://api.travis-ci.org/mgeisler/textwrap.svg?branch=master)](https://travis-ci.org/mgeisler/textwrap) Stars:`309`.
* [fancy-regex/fancy-regex](https://github.com/fancy-regex/fancy-regex) [[fancy-regex](https://crates.io/crates/fancy-regex)] - Regular expressions implementation designed to support a relatively rich set of features such as look-around and backtracking. [![crates](https://img.shields.io/crates/v/fancy-regex.svg)](https://crates.io/crates/fancy-regex) [![build badge](https://github.com/fancy-regex/fancy-regex/workflows/ci/badge.svg)](https://github.com/fancy-regex/fancy-regex/actions/workflows/ci.yml) Stars:`289`.
* [BurntSushi/tabwriter](https://github.com/BurntSushi/tabwriter) — Elastic tab stops (i.e., text column alignment) [![build badge](https://api.travis-ci.org/BurntSushi/tabwriter.svg?branch=master)](https://travis-ci.org/BurntSushi/tabwriter) Stars:`211`.
* [BurntSushi/suffix](https://github.com/BurntSushi/suffix) — Linear time suffix array construction (with Unicode support) [![build badge](https://api.travis-ci.org/BurntSushi/suffix.svg?branch=master)](https://travis-ci.org/BurntSushi/suffix) Stars:`204`.
* [ps1dr3x/easy_reader](https://github.com/ps1dr3x/easy_reader) — A reader that allows forwards, backwards and random navigations through the lines of huge files without consuming iterators [![build badge](https://api.travis-ci.org/ps1dr3x/easy_reader.svg?branch=master)](https://travis-ci.org/ps1dr3x/easy_reader) Stars:`80`.
* [cpc](https://github.com/probablykasper/cpc) - Parses and calculates strings of math with support for units and unit conversion, from `1+2` to `1% of round(1 lightyear / 14!s to km/h)`. Stars:`76`.
* [Daniel-Liu-c0deb0t/triple_accel](https://github.com/Daniel-Liu-c0deb0t/triple_accel) [[triple_accel](https://crates.io/crates/triple_accel)] - Rust edit distance routines accelerated using SIMD; supports fast Hamming, Levenshtein, restricted Damerau-Levenshtein, etc. distance calculations and string search [![build badge](https://github.com/Daniel-Liu-c0deb0t/triple_accel/workflows/Test/badge.svg?branch=master)](https://github.com/Daniel-Liu-c0deb0t/triple_accel/actions) Stars:`71`.
* [Lucretiel/joinery](https://github.com/Lucretiel/joinery) [[joinery](https://crates.io/crates/joinery)] – Generic string + iterable joining [![build badge](https://api.travis-ci.org/Lucretiel/joinery.svg?branch=master)](https://travis-ci.org/Lucretiel/joinery) Stars:`69`.
* [becheran/wildmatch](https://github.com/becheran/wildmatch) [[wildmatch](https://crates.io/crates/wildmatch)] — Simple string matching with questionmark- and star-wildcard operator [![Actions Status](https://github.com/becheran/wildmatch/workflows/Build/badge.svg?branch=master)](https://github.com/becheran/wildmatch/actions) Stars:`37`.
* [pwoolcoc/ngrams](https://github.com/pwoolcoc/ngrams) [[ngrams](https://crates.io/crates/ngrams)] — Construct [n-grams](https://en.wikipedia.org/wiki/N-gram) from arbitrary iterators [![build badge](https://api.travis-ci.org/pwoolcoc/ngrams.svg?branch=master)](https://travis-ci.org/pwoolcoc/ngrams) Stars:`27`.
* [yaa110/rake-rs](https://github.com/yaa110/rake-rs) [[rake](https://crates.io/crates/rake)] — Multilingual implementation of RAKE algorithm for Rust [![build badge](https://api.travis-ci.org/yaa110/rake-rs.svg?branch=master)](https://travis-ci.org/yaa110/rake-rs) Stars:`25`.
* [strsim-rs](https://crates.io/crates/strsim) — String similarity metrics [![build badge](https://api.travis-ci.org/dguo/strsim-rs.svg?branch=master)](https://travis-ci.org/dguo/strsim-rs)

### Text search

* [meilisearch/MeiliSearch](https://github.com/meilisearch/MeiliSearch) — Ultra relevant, instant and typo-tolerant full-text search API. [![Build Status](https://github.com/meilisearch/MeiliSearch/workflows/Cargo%20test/badge.svg?branch=master)](https://github.com/meilisearch/MeiliSearch/actions) Stars:`30.7K`.
* [tantivy](https://github.com/quickwit-oss/tantivy) [[tantivy](https://crates.io/crates/tantivy)] — A horse-speed full-text search engine library written in Rust. [![Build Status](https://github.com/quickwit-oss/tantivy/actions/workflows/test.yml/badge.svg)](https://github.com/quickwit-oss/tantivy/actions/workflows/test.yml) Stars:`7.2K`.
* [BurntSushi/fst](https://github.com/BurntSushi/fst) [[fst](https://crates.io/crates/fst)] — [![build badge](https://api.travis-ci.org/BurntSushi/fst.svg?branch=master)](https://travis-ci.org/BurntSushi/fst) Stars:`1.5K`.
* [andylokandy/simsearch-rs](https://github.com/andylokandy/simsearch-rs) [[simsearch](https://crates.io/crates/simsearch)] — A simple and lightweight fuzzy search engine that works in memory, searching for similar strings Stars:`113`.
* [CurrySoftware/perlin](https://github.com/CurrySoftware/perlin) [[perlin](https://crates.io/crates/perlin)] Stars:`70`.

### Unsafe

* [zerocopy](https://crates.io/crates/zerocopy) — Utilities for safely reinterpreting arbitrary byte sequences as native Rust types

### Virtualization

* [bytecodealliance/wasmtime](https://github.com/bytecodealliance/wasmtime) — A standalone runtime for WebAssembly [![Build Status](https://github.com/bytecodealliance/wasmtime/workflows/CI/badge.svg)](https://github.com/bytecodealliance/wasmtime/actions?query=workflow%3ACI) Stars:`10.5K`.
* [beneills/quantum](https://github.com/beneills/quantum) — Advanced Rust quantum computer simulator [![build badge](https://api.travis-ci.org/beneills/quantum.svg?branch=master)](https://travis-ci.org/beneills/quantum) Stars:`212`.
* [unicorn-rs/unicorn-rs](https://github.com/unicorn-rs/unicorn-rs) — Rust bindings for the unicorn CPU emulator [![Build Status](https://api.travis-ci.org/ekse/unicorn-rs.svg?branch=master)](https://travis-ci.org/ekse/unicorn-rs) Stars:`129`.
* [oxidecomputer/propolis](https://github.com/oxidecomputer/propolis) - Rust-based userspace program for illumos bhyve kernel modules Stars:`102`.
* [saurvs/hypervisor-rs](https://github.com/saurvs/hypervisor-rs) — Hardware-accelerated virtualization on OS X Stars:`56`.
* [chromium/chromiumos/platform/crosvm](https://chromium.googlesource.com/chromiumos/platform/crosvm/) CrOSVM — Enables Chrome OS to run Linux apps inside a fast, secure virtualized environment

### Web programming

See also [Are we web yet?](https://www.arewewebyet.org) and [Rust web framework comparison](https://github.com/flosse/rust-web-framework-comparison).

* Client-side / WASM
  * [sauron](https://github.com/ivanceras/sauron) - Client side web framework which closely adheres to The Elm Architecture. [![Build Status](https://api.travis-ci.org/ivanceras/sauron.svg?branch=master)](https://travis-ci.org/ivanceras/sauron) Stars:`1.7K`.
  * [cargo-web](https://crates.io/crates/cargo-web) — A Cargo subcommand for the client-side Web [![Build Status](https://api.travis-ci.org/koute/cargo-web.svg?branch=master)](https://travis-ci.org/koute/cargo-web)
  * [seed](https://seed-rs.org/) — A Rust framework for creating web apps
  * [stdweb](https://crates.io/crates/stdweb) — A standard library for the client-side Web [![Build Status](https://api.travis-ci.org/koute/stdweb.svg?branch=master)](https://travis-ci.org/koute/stdweb)
  * [yew](https://crates.io/crates/yew) — Rust framework for making client web apps
* HTTP Client
  * [hyperium/hyper](https://github.com/hyperium/hyper) — an HTTP implementation [![CI](https://github.com/hyperium/hyper/workflows/CI/badge.svg?branch=master)](https://github.com/hyperium/hyper/actions?query=workflow%3ACI) Stars:`10.7K`.
  * [seanmonstar/reqwest](https://github.com/seanmonstar/reqwest) — an ergonomic HTTP Client for Rust. [![build badge](https://api.travis-ci.org/seanmonstar/reqwest.svg?branch=master)](https://travis-ci.org/seanmonstar/reqwest) Stars:`6.5K`.
  * [async-graphql](https://github.com/async-graphql/async-graphql) - A GraphQL server library implemented in Rust [![Build Status](https://dev.azure.com/graphql-rust/GraphQL%20Rust/_apis/build/status/graphql-rust.juniper)](https://dev.azure.com/graphql-rust/GraphQL%20Rust/_build/latest?definitionId=1) Stars:`2.5K`.
  * [graphql-client](https://github.com/graphql-rust/graphql-client) — Typed, correct GraphQL requests and responses in Rust. [![Github actions Status](https://github.com/graphql-rust/graphql-client/workflows/CI/badge.svg?branch=master)](https://github.com/graphql-rust/graphql-client/actions) Stars:`896`.
  * [alexcrichton/curl-rust](https://github.com/alexcrichton/curl-rust) — [libcurl](https://curl.se/libcurl/) bindings [![build badge](https://api.travis-ci.com/alexcrichton/curl-rust.svg?branch=master)](https://travis-ci.org/alexcrichton/curl-rust) Stars:`870`.
  * [DoumanAsh/yukikaze](https://gitlab.com/Douman/yukikaze) [[yukikaze](https://crates.io/crates/yukikaze)] — Beautiful and elegant Yukikaze is little HTTP client library based on hyper. [![build badge](https://gitlab.com/Douman/yukikaze/badges/master/pipeline.svg)](https://gitlab.com/Douman/yukikaze)
* HTTP Server
  * [Rocket](https://github.com/SergioBenitez/Rocket) — Rocket is web framework for Rust (nightly) with a focus on ease-of-use, expressability, and speed [![build badge](https://api.travis-ci.org/SergioBenitez/Rocket.svg?branch=master)](https://travis-ci.org/SergioBenitez/Rocket) Stars:`18.9K`.
  * [actix/actix-web](https://github.com/actix/actix-web) — A lightweight async web framework for Rust with websocket support [![build badge](https://api.travis-ci.org/actix/actix-web.svg?branch=master)](https://travis-ci.org/actix/actix-web) Stars:`15.7K`.
  * [hyperium/hyper](https://github.com/hyperium/hyper) — an HTTP implementation [![CI](https://github.com/hyperium/hyper/workflows/CI/badge.svg?branch=master)](https://github.com/hyperium/hyper/actions?query=workflow%3ACI) Stars:`10.7K`.
  * [seanmonstar/warp](https://github.com/seanmonstar/warp) — A super-easy, composable, web server framework for warp speeds. [![crate](https://img.shields.io/crates/v/create-rust-app.svg)](https://crates.io/crates/warp) Stars:`7.3K`.
  * [tokio/axum](https://github.com/tokio-rs/axum) - Ergonomic and modular web framework built with Tokio, Tower, and Hyper [![Build badge](https://github.com/tokio-rs/axum/actions/workflows/CI.yml/badge.svg?branch=main)](https://github.com/tokio-rs/axum/actions/workflows/CI.yml) Stars:`6.9K`.
  * [Iron](https://github.com/iron/iron) — A middleware-based server framework [![build badge](https://api.travis-ci.org/GildedHonour/frank_jwt.svg?branch=master)](https://travis-ci.org/GildedHonour/frank_jwt) Stars:`6.1K`.
  * [Juniper](https://github.com/graphql-rust/juniper) — GraphQL server library for Rust [![build badge](https://api.travis-ci.org/graphql-rust/juniper.svg?branch=master)](https://travis-ci.org/graphql-rust/juniper) Stars:`4.8K`.
  * [Gotham](https://github.com/gotham-rs/gotham) — A flexible web framework that does not sacrifice safety, security or speed. [![build badge](https://api.travis-ci.org/gotham-rs/gotham.svg?branch=master)](https://travis-ci.org/gotham-rs/gotham) Stars:`2.0K`.
  * [Salvo](https://github.com/salvo-rs/salvo) — an easy to use webframework base on hyper and tokio. [![build build](https://github.com/salvo-rs/salvo/workflows/CI%20(Linux)/badge.svg?branch=master&event=push)](https://github.com/salvo-rs/salvo/actions) Stars:`969`.
  * [carllerche/tower-web](https://github.com/carllerche/tower-web) [[tower-web](https://crates.io/crates/tower-web)] — A fast, boilerplate free, web framework for Rust [![build badge](https://api.travis-ci.org/carllerche/tower-web.svg?branch=master)](https://travis-ci.org/carllerche/tower-web) Stars:`966`.
  * [handlebars-rust](https://github.com/sunng87/handlebars-rust) — an Iron web framework middleware. [![build badge](https://api.travis-ci.org/sunng87/handlebars-iron.svg?branch=master)](https://travis-ci.org/sunng87/handlebars-iron) Stars:`893`.
  * [Ogeon/rustful](https://github.com/Ogeon/rustful) — A RESTful web framework for Rust  [![build badge](https://api.travis-ci.org/Ogeon/rustful.svg?branch=master)](https://travis-ci.org/Ogeon/rustful) Stars:`872`.
  * [tomaka/rouille](https://github.com/tomaka/rouille) — Web framework in Rust [![build badge](https://api.travis-ci.org/tomaka/rouille.svg?branch=master)](https://travis-ci.org/tomaka/rouille) Stars:`825`.
  * [tiny-http](https://github.com/tiny-http/tiny-http) — Low level HTTP server library [![build badge](https://api.travis-ci.org/tiny-http/tiny-http.svg?branch=master)](https://travis-ci.org/tiny-http/tiny-http) Stars:`768`.
  * [miketang84/sapper](https://github.com/miketang84/sapper) — A lightweight web framework built on async hyper, implemented in Rust language. Stars:`622`.
  * [Rustless](https://github.com/rustless/rustless) — A REST-like API micro-framework inspired by [Grape](https://github.com/ruby-grape/grape) and [Hyper](https://github.com/hyperium/hyper) [![build badge](https://api.travis-ci.org/rustless/rustless.svg?branch=master)](https://travis-ci.org/rustless/rustless) Stars:`609`.
  * [GildedHonour/frank_jwt](https://github.com/GildedHonour/frank_jwt) — JSON Web Token implementation in Rust. [![build badge](https://api.travis-ci.org/GildedHonour/frank_jwt.svg?branch=master)](https://travis-ci.org/GildedHonour/frank_jwt) Stars:`241`.
  * [Graphul](https://github.com/graphul-rs/graphul) — An Express-inspired web framework written in Rust. [![crate](https://img.shields.io/crates/v/create-rust-app.svg)](https://crates.io/crates/graphul) Stars:`219`.
  * [danclive/sincere](https://github.com/danclive/sincere) — A micro web framework for Rust(stable) based on hyper and multithreading. [![build badge](https://api.travis-ci.org/danclive/sincere.svg?branch=master)](https://travis-ci.org/danclive/sincere) Stars:`94`.
  * [Saphir](https://github.com/richerarc/saphir) — A progressive web framework with low-level control, without the pain. Stars:`82`.
  * [branca](https://crates.io/crates/branca) — A Pure Rust implementation of Branca for Authenticated and Encrypted API tokens. [![build badge](https://api.travis-ci.org/return/branca.svg?branch=master)](https://travis-ci.org/return/branca)
  * [Nickel](https://github.com/nickel-org/nickel.rs/) — inspired by [Express](http://expressjs.com/) [![build badge](https://api.travis-ci.org/nickel-org/nickel.rs.svg?branch=master)](https://travis-ci.org/nickel-org/nickel.rs)
* Miscellaneous
  * [serenity-rs/serenity](https://github.com/serenity-rs/serenity) [[serenity](https://crates.io/crates/serenity)] - A Rust library for the Discord API Stars:`3.1K`.
  * [osohq/oso](https://github.com/osohq/oso) [[oso](https://crates.io/crates/oso)] - A policy engine for authorization that's embedded in your application. [![Build Status](https://github.com/osohq/oso/workflows/Development/badge.svg?branch=main)](https://github.com/osohq/oso/actions?query=branch%3Amain+workflow%3ADevelopment) Stars:`2.8K`.
  * [causal-agent/scraper](https://github.com/causal-agent/scraper) [[scraper](https://crates.io/crates/scraper)] - HTML parsing and querying with CSS selectors. [![Build Status](https://github.com/causal-agent/scraper/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/causal-agent/scraper/actions) Stars:`1.1K`.
  * [pyrossh/rust-embed](https://github.com/pyrossh/rust-embed) — A macro to embed static assets into the rust binary Stars:`947`.
  * [utkarshkukreti/select.rs](https://github.com/utkarshkukreti/select.rs) [[select](https://crates.io/crates/select)] — A library to extract useful data from HTML documents, suitable for web scraping. [![Build Status](https://api.travis-ci.org/utkarshkukreti/select.rs.svg?branch=master)](https://travis-ci.org/utkarshkukreti/select.rs) Stars:`813`.
  * [cargonauts](https://github.com/cargonauts-rs/cargonauts) — A web framework intended for building maintainable, well-factored web apps. Stars:`178`.
  * [hominee/dyer](https://github.com/hominee/dyer) [[dyer](https://crates.io/crates/dyer)] - dyer is designed for reliable, flexible and fast Request-Response based service, including data processing, web-crawling and so on, providing some friendly, flexible, comprehensive features without compromising speed. Stars:`116`.
  * [softprops/openapi](https://github.com/softprops/openapi) — A library for processing openapi spec files Stars:`104`.
  * [pwoolcoc/soup](https://gitlab.com/pwoolcoc/soup) [[soup](https://crates.io/crates/soup)] — A library similar to Python's BeautifulSoup, designed to enable quick and easy manipulation and querying of HTML documents. [![Build Status](https://gitlab.com/pwoolcoc/soup/badges/master/pipeline.svg)](https://gitlab.com/pwoolcoc/soup/badges/master/pipeline.svg)
  * [tbot](https://gitlab.com/SnejUgal/tbot) [[tbot](https://crates.io/crates/tbot)] - Make cool Telegram bots with Rust easily [![pipeline status](https://gitlab.com/SnejUgal/tbot/badges/master/pipeline.svg)](https://gitlab.com/SnejUgal/tbot/-/commits/master)
  * [teloxide/teloxide](https://github.com/teloxide/teloxide/) - An elegant Telegram bots framework for Rust [![Build Status](https://github.com/teloxide/teloxide/workflows/Continuous%20integration/badge.svg?branch=master)](https://github.com/teloxide/teloxide/actions)
* Reverse Proxy
  * [sozu-proxy/sozu](https://github.com/sozu-proxy/sozu) [[sozu](https://crates.io/crates/sozu)] — A HTTP reverse proxy. [![CI](https://github.com/sozu-proxy/sozu/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/sozu-proxy/sozu/actions/workflows/ci.yml) Stars:`1.9K`.
* Static Site Generators
  * [getzola/zola](https://github.com/getzola/zola) [[zola](https://www.getzola.org/)] — An opinionated static site generator with everything built-in. [![Build Status](https://dev.azure.com/getzola/zola/_apis/build/status/getzola.zola?branchName=master)](https://dev.azure.com/getzola/zola/_build) Stars:`9.7K`.
  * [vi/websocat](https://github.com/vi/websocat) — CLI for interacting with WebSockets, with functionality of Netcat, Curl and Socat. [![build badge](https://api.travis-ci.org/vi/websocat.svg?branch=master)](https://travis-ci.org/vi/websocat) Stars:`4.7K`.
  * [housleyjk/ws-rs](https://github.com/housleyjk/ws-rs) — lightweight, event-driven WebSockets for Rust [![build badge](https://api.travis-ci.org/housleyjk/ws-rs.svg?branch=stable)](https://travis-ci.org/housleyjk/ws-rs) Stars:`1.3K`.
  * [rust-websocket](https://github.com/websockets-rs/rust-websocket) — A framework for dealing with WebSocket connections (both clients and servers) [![build badge](https://api.travis-ci.org/websockets-rs/rust-websocket.svg?branch=master)](https://travis-ci.org/websockets-rs/rust-websocket) Stars:`1.2K`.
  * [snapview/tungstenite-rs](https://github.com/snapview/tungstenite-rs) — Lightweight stream-based WebSocket implementation for Rust. Stars:`1.2K`.
  * [cobalt-org/cobalt.rs](https://github.com/cobalt-org/cobalt.rs) — Static site generator written in Rust [![Build Status](https://dev.azure.com/cobalt-org/cobalt-org/_apis/build/status/cobalt.rs?branchName=master)](https://dev.azure.com/cobalt-org/cobalt-org/_build?definitionId=2) Stars:`1.2K`.
  * [leven-the-blog/leven](https://github.com/leven-the-blog/leven) [[leven](https://crates.io/crates/leven)] — A simple, parallelized blog generator. [![build badge](https://api.travis-ci.org/quadrupleslap/leven.svg?branch=master)](https://travis-ci.org/quadrupleslap/leven) Stars:`55`.
  * [FuGangqiang/mdblog.rs](https://github.com/FuGangqiang/mdblog.rs) [[mdblog](https://crates.io/crates/mdblog)] — Static site generator from markdown files. Stars:`48`.
  * [vityafx/urlshortener-rs](https://github.com/vityafx/urlshortener-rs) — A very simple urlshortener library for Rust. [![CI](https://github.com/vityafx/urlshortener-rs/actions/workflows/ci.yml/badge.svg)](https://github.com/vityafx/urlshortener-rs/actions/workflows/ci.yml) [![Crates badge](https://img.shields.io/crates/v/urlshortener.svg)](https://crates.io/crates/urlshortener) Stars:`38`.
* [WebSocket](https://datatracker.ietf.org/doc/rfc6455/)

## Registries

A registry allows you to publish your Rust libraries as crate packages, to share them with others publicly and privately.

* [w4/chartered](https://github.com/w4/chartered) - A private, authenticated, permissioned Cargo registry [![CI](https://github.com/w4/chartered/actions/workflows/ci.yml/badge.svg)](https://github.com/w4/chartered/actions/workflows/ci.yml) Stars:`117`.
* [Cloudsmith :heavy_dollar_sign:](https://cloudsmith.com/cargo-registry/) — A fully managed package management SaaS, with first-class support for public and private Cargo/Rust registries (plus many others). Has a generous free-tier and is also completely free for open-source.
* [Crates](https://crates.io) — The official public registry for Rust/Cargo.

## Resources

* Benchmarks
  * [TeXitoi/benchmarksgame-rs](https://github.com/TeXitoi/benchmarksgame-rs) — Rust implementations for the [The Computer Language Benchmarks Game](https://benchmarksgame-team.pages.debian.net/benchmarksgame/) [![build badge](https://api.travis-ci.org/TeXitoi/benchmarksgame-rs.svg?branch=master)](https://travis-ci.org/TeXitoi/benchmarksgame-rs) Stars:`69`.
* Decks & Presentations
  * [Learning systems programming with Rust](https://speakerdeck.com/jvns/learning-systems-programming-with-rust) — Presented by [Julia Evans](https://twitter.com/@b0rk) @ Rustconf 2016.
  * [Rust: Hack Without Fear!](https://www.youtube.com/watch?v=lO1z-7cuRYI) — Presented by [Nicholas Matsakis](https://github.com/nikomatsakis) @ C++Now 2018
  * [Shipping a Solid Rust Crate](https://www.youtube.com/watch?v=t4CyEKb-ywA) — Presented by [Michael Gattozzi](https://github.com/mgattozzi) @ RustConf 2017
* [Discover Rust Libraries & Code Snippets](https://kandi.openweaver.com/explore/rust) - A curated list of Rust libraries, authors, kits, tutorials & learning resources on kandi
* Learning
  * [Rustlings](https://github.com/rust-lang/rustlings) — small exercises to get you used to reading and writing Rust code Stars:`31.4K`.
  * [rust-learning](https://github.com/ctjhoa/rust-learning) — A collection of useful resources to learn Rust Stars:`8.7K`.
  * [Easy Rust](https://github.com/Dhghomon/easy_rust) - Learn Rust in easy English. Stars:`7.2K`.
  * [Idiomatic Rust](https://github.com/mre/idiomatic-rust) —  A peer-reviewed collection of articles/talks/repos which teach idiomatic Rust. Stars:`4.0K`.
  * [stdx](https://github.com/brson/stdx) — Learn these crates first as an extension to std Stars:`1.8K`.
  * [rust-how-do-i-start](https://github.com/jondot/rust-how-do-i-start) - A repo dedicated to answering the question: "So, Rust. How do I _start_?". A beginner only hand-picked resources and learning track. Stars:`640`.
  * [Rust Gym](https://github.com/warycat/rustgym) - A big collection of coding interview problems solved in Rust. Stars:`587`.
  * [Awesome Rust Streaming](https://github.com/jamesmunns/awesome-rust-streaming) - A community curated list of livestreams about Rust. Stars:`578`.
  * [Rusty CS](https://github.com/AbdesamedBendjeddou/Rusty-CS) - A Computer Science Curriculum that helps practice the acquired academic knowledge in Rust Stars:`572`.
  * [Refactoring to Rust](https://www.manning.com/books/refactoring-to-rust) - A book that introduces to Rust language.
  * [Rust Language Cheat Sheet](https://cheats.rs/)
  * [Hands-on Rust](https://pragprog.com/titles/hwrust/hands-on-rust/) - A hands-on guide to learning Rust by making games - by [Herbert Wolverson](https://github.com/thebracket/) (paid)
  * [Rust by Example](https://doc.rust-lang.org/rust-by-example/)
  * [Rust Cookbook](https://rust-lang-nursery.github.io/rust-cookbook/) — A collection of simple examples that demonstrate good practices to accomplish common programming tasks, using the crates of the Rust ecosystem.
  * [Rust for professionals](https://overexact.com/rust-for-professionals/) — A quick introduction to Rust for experienced software developers.
  * [Little Book of Rust Books](https://lborb.github.io/book/) - Curated list of rust books and how-tos.
  * [Rust in Action](https://www.manning.com/books/rust-in-action) — A hands-on guide to systems programming with Rust by [Tim McNamara](https://github.com/timClicks) (paid)
  * [Rust in Motion](https://www.manning.com/livevideo/rust-in-motion?a_aid=cnichols&a_bid=6a993c2e) — A video series by [Carol Nichols](https://github.com/carols10cents) and [Jake Goulding](https://github.com/shepmaster) (paid)
  * [Programming Community Curated Resources for Learning Rust](https://hackr.io/tutorials/learn-rust) — A list of recommended resources voted by the programming community.
  * [Rust Online Courses at Classpert](https://classpert.com/search/rust) — A list of Rust online courses (paid) from Classpert Online Course Search
  * [Rust Tiếng Việt](https://rust-tieng-viet.github.io/) - Learn Rust in Vietnamese.
  * [Learning Rust With Entirely Too Many Linked Lists](https://rust-unofficial.github.io/too-many-lists/) — in-depth exploration of Rust's memory management rules, through implementing a few different types of list structures.
  * [CodeCrafters.io](https://app.codecrafters.io/tracks/rust) — Build your own Redis, Git, Docker, or SQLite in Rust
  * [Build a language VM](https://blog.subnetzero.io/post/building-language-vm-part-00/)
  * [exercism.org](https://exercism.org/tracks/rust) — programming exercises that help you learn new concepts in Rust.
  * [awesome-rust-mentors](https://rustbeginners.github.io/awesome-rust-mentors/) — A list of helpful Rust mentors willing to take mentees and eductate them about Rust and programming.
  * [Take your first steps with Rust](https://learn.microsoft.com/en-us/training/paths/rust-first-steps/) - Lay the foundation of knowledge you need to build fast and effective programs in Rust.
  * [University of Pennsylvania's Comp Sci Rust Programming Course](http://cis198-2016s.github.io/schedule/)
* Podcasts
* [Rust Design Patterns](https://github.com/rust-unofficial/patterns) Stars:`6.3K`.
* [RustViz](https://github.com/rustviz/rustviz) — generates visualizations from simple Rust programs to assist users in better understanding the Rust Lifetime and Borrowing mechanism. Stars:`2.1K`.
* [RustBooks](https://github.com/sger/RustBooks) — list of RustBooks Stars:`2.1K`.
  * [New Rustacean](https://newrustacean.com) — A podcast about learning Rust
  * [Rustacean Station](https://rustacean-station.org/) — A community project for creating podcast content for Rust
* [Rust Guidelines](http://aturon.github.io/)
* [Rust Servers, Services and Apps - MEAP](https://www.manning.com/books/rust-servers-services-and-apps) - Build backend servers, services, and front-ends in Rust to get fast, reliable, and maintainable applications.
* [Rust Subreddit](https://www.reddit.com/r/rust/) — A subreddit(forum) where rust related questions, articles and resources are posted and discussed
* [RustCamp 2015 Talks](https://www.youtube.com/playlist?list=PLE7tQUdRKcybdIw61JpCoo89i4pWU5f_t)

## License

[![CC0](https://licensebuttons.net/p/zero/1.0/88x31.png)](https://creativecommons.org/publicdomain/zero/1.0/)
