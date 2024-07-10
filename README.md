# Awesome Rust [![build badge](https://github.com/rust-unofficial/awesome-rust/actions/workflows/rust.yml/badge.svg?branch=main)](https://github.com/rust-unofficial/awesome-rust/actions/workflows/rust.yml) [![Track Awesome List](https://www.trackawesomelist.com/badge.svg)](https://www.trackawesomelist.com/rust-unofficial/awesome-rust/)

A curated list of Rust code and resources.

If you want to contribute, please read [this](CONTRIBUTING.md).

## Table of contents

<!-- toc -->

- [Applications](#applications)
  * [Audio and Music](#audio-and-music)
  * [Blockchain](#blockchain)
  * [Database](#database)
  * [Emulators](#emulators)
  * [File manager](#file-manager)
  * [Games](#games)
  * [Graphics](#graphics)
  * [Image processing](#image-processing)
  * [Industrial automation](#industrial-automation)
  * [Observability](#observability)
  * [Operating systems](#operating-systems)
  * [Package Managers](#package-managers)
  * [Payments](#payments)
  * [Productivity](#productivity)
  * [Routing protocols](#routing-protocols)
  * [Security tools](#security-tools)
  * [Social networks](#social-networks)
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
    + [OpenAI](#openai)
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
  * [eBPF](#ebpf)
  * [Email](#email)
  * [Encoding](#encoding)
  * [Filesystem](#filesystem)
  * [Finance](#finance)
  * [Functional Programming](#functional-programming)
  * [Game development](#game-development)
  * [Geospatial](#geospatial)
  * [Graph algorithms](#graph-algorithms)
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
  * [Simulation](#simulation)
  * [System](#system)
  * [Task scheduling](#task-scheduling-1)
  * [Template engine](#template-engine)
  * [Text processing](#text-processing-1)
  * [Text search](#text-search)
  * [Unsafe](#unsafe)
  * [Video](#video-1)
  * [Virtualization](#virtualization-1)
  * [Web programming](#web-programming)
- [Registries](#registries)
- [Resources](#resources)
- [License](#license)

<!-- tocstop -->

## Applications

See also [Rust - Production](https://www.rust-lang.org/production) organizations running Rust in production.

* [denoland/deno](https://github.com/denoland/deno) - A secure JavaScript/TypeScript runtime built with V8 and Tokio [![Build Status](https://github.com/denoland/deno/workflows/ci/badge.svg?branch=master&event=push)](https://github.com/denoland/deno/actions) Stars:`93.4K`.
* [alacritty](https://github.com/alacritty/alacritty) - A cross-platform, GPU enhanced terminal emulator Stars:`54.6K`.
* [SWC](https://github.com/swc-project/swc) - super-fast TypeScript / JavaScript compiler Stars:`30.5K`.
* [Servo](https://github.com/servo/servo) - A prototype web browser engine Stars:`26.7K`.
* [zellij](https://github.com/zellij-org/zellij) - A terminal multiplexer (workspace) with batteries included Stars:`19.5K`.
* [wasmer](https://github.com/wasmerio/wasmer) - A safe and fast WebAssembly runtime supporting WASI and Emscripten [![Build Status](https://github.com/wasmerio/wasmer/workflows/build/badge.svg?style=flat-square)](https://github.com/wasmerio/wasmer/actions) Stars:`18.2K`.
* [mdBook](https://github.com/rust-lang/mdBook) - A command line utility to create books from markdown files [![Build Status](https://github.com/rust-lang/mdBook/workflows/CI/badge.svg?branch=master)](https://github.com/rust-lang/mdBook/actions) Stars:`17.3K`.
* [wezterm](https://github.com/wez/wezterm) - A GPU-accelerated cross-platform terminal emulator and multiplexer Stars:`14.9K`.
* [Sniffnet](https://github.com/GyulyVGC/sniffnet) - Cross-platform application to monitor your network traffic with ease [![build badge](https://img.shields.io/github/actions/workflow/status/gyulyvgc/sniffnet/rust.yml?logo=github)](https://github.com/GyulyVGC/sniffnet/blob/main/.github/workflows/rust.yml) [![crate](https://img.shields.io/crates/v/sniffnet?logo=rust)](https://crates.io/crates/sniffnet) Stars:`14.4K`.
* [cloudflare/boringtun](https://github.com/cloudflare/boringtun) - A Userspace WireGuard VPN Implementation [![build badge](https://img.shields.io/badge/crates.io-v0.2.0-orange.svg)](https://crates.io/crates/boringtun) Stars:`5.9K`.
* [shuttle](https://github.com/shuttle-hq/shuttle) - A serverless platform. Stars:`5.8K`.
* [datafusion](https://github.com/apache/datafusion) - Apache Arrow DataFusion and Ballista query engines Stars:`5.5K`.
* [innernet](https://github.com/tonarino/innernet) - An overlay or private mesh network that uses Wireguard under the hood Stars:`4.9K`.
* [hickory-dns](https://crates.io/crates/trust-dns) - A DNS-server [![Build Status](https://github.com/hickory-dns/hickory-dns/workflows/test/badge.svg?branch=main)](https://github.com/hickory-dns/hickory-dns/actions?query=workflow%3Atest) Stars:`3.8K`.
* [mirrord](https://github.com/metalbear-co/mirrord) - Connect your local process and your cloud environment, and run local code in cloud conditions Stars:`3.6K`.
* [Rio](https://github.com/raphamorim/rio) - A hardware-accelerated GPU terminal emulator powered by WebGPU, focusing to run in desktops and browsers. Stars:`3.2K`.
* [rx](https://github.com/cloudhead/rx) - Vi inspired Modern Pixel Art Editor Stars:`3.1K`.
* [WinterJS](https://github.com/wasmerio/winterjs) - A secure JavaScript runtime built with SpiderMonkey and Axum Stars:`2.9K`.
* [habitat](https://github.com/habitat-sh/habitat) - A tool created by Chef to build, deploy, and manage applications. Stars:`2.6K`.
* [notty](https://github.com/withoutboats/notty) - A new kind of terminal Stars:`2.3K`.
* [fcsonline/drill](https://github.com/fcsonline/drill) - A HTTP load testing application inspired by Ansible syntax Stars:`2.0K`.
* [linkerd/linkerd2-proxy](https://github.com/linkerd/linkerd2-proxy) - Ultralight service mesh for Kubernetes. Stars:`2.0K`.
* [Ryot](https://github.com/ignisda/ryot) - A self hosted application to track media consumption, fitness, etc. Stars:`1.7K`.
* [kalker](https://github.com/PaddiM8/kalker) - A scientific calculator that supports math-like syntax with user-defined variables, functions, derivation, integration, and complex numbers. Cross-platform + WASM support [![Build Status](https://github.com/PaddiM8/kalker/workflows/Release/badge.svg)](https://github.com/PaddiM8/kalker/actions) Stars:`1.6K`.
* [tiny](https://github.com/osa1/tiny) - A terminal IRC client Stars:`989`.
* [Fractalide](https://github.com/fractalide/fractalide) - Simple microservices Stars:`862`.
* [jedisct1/flowgger](https://github.com/awslabs/flowgger) - A fast, simple and lightweight data collector Stars:`821`.
* [defguard](https://github.com/defguard/defguard) - Enterprise Open Source SSO & WireGuard VPN with real 2FA/MFA Stars:`763`.
* [kftray](https://github.com/hcavarsan/kftray) - A cross-platform system tray app for managing and sharing multiple kubectl port-forward configurations. [![Build Status](https://github.com/hcavarsan/kftray/workflows/Release/badge.svg)](https://github.com/hcavarsan/kftray/actions) Stars:`657`.
* [fend](https://github.com/printfn/fend) - Arbitrary-precision unit-aware calculator [![build](https://github.com/printfn/fend/workflows/build/badge.svg)](https://github.com/printfn/fend) Stars:`591`.
* [EasyTier](https://github.com/EasyTier/EasyTier) - A simple, full-featured and decentralized mesh VPN with WireGuard support. [![crates.io](https://img.shields.io/crates/v/easytier)](https://crates.io/crates/easytier) [![GitHub actions](https://github.com/EasyTier/EasyTier/actions/workflows/core.yml/badge.svg)](https://github.com/EasyTier/EasyTier/actions/)[![GitHub actions](https://github.com/EasyTier/EasyTier/actions/workflows/gui.yml/badge.svg)](https://github.com/EasyTier/EasyTier/actions/) Stars:`586`.
* [kytan](https://github.com/changlan/kytan) - High Performance Peer-to-Peer VPN Stars:`472`.
* [asm-cli-rust](https://github.com/cch123/asm-cli-rust) - An interactive assembly shell. Stars:`306`.
* [Weld](https://github.com/serayuzgur/weld) - Full fake REST API generator Stars:`297`.
* [Rauthy](https://github.com/sebadob/rauthy) - OpenID Connect Single Sign-On Identity & Access Management Stars:`234`.
* [Factotum](https://github.com/snowplow/factotum) - A system to programmatically run data pipelines Stars:`221`.
* [shoes](https://github.com/cfal/shoes) - A multi-protocol proxy server Stars:`183`.
* [doprz/dipc](https://github.com/doprz/dipc) - Convert your favorite images and wallpapers with your favorite color palettes/themes [![crates.io](https://img.shields.io/crates/v/dipc)](https://crates.io/crates/dipc) Stars:`181`.
* [UpVPN](https://github.com/upvpn/upvpn-app) - WireGuard VPN client for macOS, Linux, and Windows built on Tauri. Stars:`162`.
* [nicohman/eidolon](https://github.com/nicohman/eidolon) - A steam and drm-free game registry and launcher for linux and macosx Stars:`125`.
* [Herd](https://github.com/imjacobclark/Herd) - an experimental HTTP load testing application Stars:`109`.
* [Pijul](https://pijul.org) - A patch-based distributed version control system
* [MaidSafe](https://github.com/maidsafe) - A decentralized platform.
* [Arti](https://gitlab.torproject.org/tpo/core/arti) - An implementation of Tor. (So far, it's a not-very-complete client. But watch this space!) [![Crates.io](https://img.shields.io/crates/v/arti.svg)](https://crates.io/crates/arti)

### Audio and Music

* [Spotifyd](https://github.com/Spotifyd/spotifyd) - An open source Spotify client running as a UNIX daemon. ![Continuous Integration](https://github.com/Spotifyd/spotifyd/workflows/Continuous%20Integration/badge.svg?branch=master) Stars:`9.7K`.
* [ncspot](https://github.com/hrkfdn/ncspot) - Cross-platform ncurses Spotify client, inspired by ncmpc and the likes. [![build badge](https://github.com/hrkfdn/ncspot/workflows/Build/badge.svg)](https://github.com/hrkfdn/ncspot/actions?query=workflow%3ABuild) Stars:`4.8K`.
* [Spotify Player](https://github.com/aome510/spotify-player) - A Spotify player in the terminal with full feature parity. Stars:`2.9K`.
* [Glicol](https://github.com/chaosprint/glicol) - Graph-oriented live coding language, for collaborative musicking in browsers. Stars:`2.1K`.
* [Polaris](https://github.com/agersant/polaris) - A music streaming application. Stars:`1.4K`.
* [termusic](https://github.com/tramhao/termusic) - Music Player TUI written Stars:`968`.
* [enginesound](https://github.com/DasEtwas/enginesound) - A GUI and command line application used to procedurally generate semi-realistic engine sounds. Featuring in-depth configuration, variable sample rate and a frequency analysis window. Stars:`293`.
* [Festival](https://github.com/hinto-janai/festival) - A local music player/server/client [![build-badge](https://github.com/hinto-janai/festival/actions/workflows/ci.yml/badge.svg)](https://github.com/hinto-janai/festival/actions/workflows/ci.yml) Stars:`262`.
* [Pinepods](https://github.com/madeofpendletonwool/PinePods) - A rust based podcast management system with multi-user support. Pinepods utilizes a central database so aspects like listen time and themes follow from device to device. With clients built using Tauri, it's a full cross-platform listening solution! ![Docker Container Build](https://github.com/madeofpendletonwool/PinePods/actions/workflows/docker-publish.yml/badge.svg) Stars:`165`.
* [dano](https://github.com/kimono-koans/dano) - A hashdeep/md5tree (but much more) for media files Stars:`136`.
* [figsoda/mmtc](https://github.com/figsoda/mmtc) [[mmtc](https://crates.io/crates/mmtc)] - Minimal mpd terminal client that aims to be simple yet highly configurable [![build-badge](https://github.com/figsoda/mmtc/actions/workflows/ci.yml/badge.svg)](https://github.com/figsoda/mmtc/actions/workflows/ci.yml) Stars:`88`.
* [WhatBPM](https://github.com/sergree/whatbpm) - A daily statically generated information resource for electronic dance music producers. Provides daily analytics on the most frequently used values for each EDM genre: tempos, keys, root notes, and so on, using publicly available data such as Beatport and Spotify. ![Continuous Integration](https://github.com/sergree/whatbpm/actions/workflows/website_build_deploy.yml/badge.svg?branch=main) Stars:`67`.

### Blockchain

* [Diem](https://github.com/diem/diem) - Diem’s mission is to enable a simple global currency and financial infrastructure that empowers billions of people. Stars:`16.7K`.
* [Solana](https://github.com/solana-labs/solana) - Incredibly fast, highly scalable blockchain using Proof-of-History. Stars:`12.7K`.
* [Foundry](https://github.com/foundry-rs/foundry) - Foundry is a blazing fast, portable and modular toolkit for Ethereum application development. ![Build Status](https://img.shields.io/github/workflow/status/foundry-rs/foundry/test?style=flat-square) Stars:`7.9K`.
* [Sui](https://github.com/MystenLabs/sui) - A next-generation smart contract platform with high throughput, low latency, and an asset-oriented programming model powered by the Move programming language. Stars:`5.8K`.
* [Grin](https://github.com/mimblewimble/grin/) - Evolution of the MimbleWimble protocol Stars:`5.0K`.
* [zcash](https://github.com/zcash/zcash) - Zcash is an implementation of the "Zerocash" protocol. Stars:`4.9K`.
* [Lighthouse](https://github.com/sigp/lighthouse) - Ethereum Consensus Layer (CL) Client [![Build Status](https://github.com/sigp/lighthouse/workflows/test-suite/badge.svg?branch=master)](https://github.com/sigp/lighthouse/actions) Stars:`2.8K`.
* [ethers-rs](https://github.com/gakonst/ethers-rs) - Complete Ethereum & Celo library and wallet implementation. ![Build Status](https://github.com/gakonst/ethers-rs/workflows/Tests/badge.svg) Stars:`2.4K`.
* [near/nearcore](https://github.com/near/nearcore) - decentralized smart-contract platform for low-end mobile devices. Stars:`2.3K`.
* [artemis](https://github.com/paradigmxyz/artemis) - A simple, modular, and fast framework for writing MEV bots. Stars:`2.2K`.
* [rust-bitcoin](https://github.com/rust-bitcoin/rust-bitcoin) - Library with support for de/serialization, parsing and executing on data structures and network messages related to Bitcoin. Stars:`2.0K`.
* [polkadot-sdk](https://github.com/paritytech/polkadot-sdk) - The Parity Polkadot Blockchain SDK Stars:`1.6K`.
* [revm](https://github.com/bluealloy/revm) - Revolutionary Machine (revm) is a fast Ethereum virtual machine. Stars:`1.5K`.
* [cairo](https://github.com/starkware-libs/cairo) - Cairo is the first Turing-complete language for creating provable programs for general computation. This is also the native language of [StarkNet](https://www.starknet.io), a ZK-Rollup using STARK proofs ![GitHub Workflow Status](https://img.shields.io/github/workflow/status/starkware-libs/cairo/CI?style=flat-square&logo=github) Stars:`1.5K`.
* [Joystream](https://github.com/Joystream/joystream) - A user governed video platform Stars:`1.4K`.
* [CITA](https://github.com/citahub/cita) - A high performance blockchain kernel for enterprise users. Stars:`1.3K`.
* [Holochain](https://github.com/holochain/holochain) - Scalable P2P alternative to blockchain for all those distributed apps you always wanted to build. [![detect critical check failures](https://github.com/holochain/holochain/actions/workflows/check_run_detect_release_pr_failure.yml/badge.svg)](https://github.com/holochain/holochain/actions/workflows/check_run_detect_release_pr_failure.yml) Stars:`1.1K`.
* [Nervos CKB](https://github.com/nervosnetwork/ckb) - Nervos CKB is a public permissionless blockchain, the common knowledge layer of Nervos network. Stars:`1.1K`.
* [rust-lightning](https://github.com/lightningdevkit/rust-lightning) [![Crate](https://img.shields.io/crates/v/lightning.svg?logo=rust)](https://crates.io/crates/lightning) - Bitcoin Lightning library. The main crate,`lightning`, does not handle networking, persistence, or any other I/O. Thus,it is runtime-agnostic, but users must implement basic networking logic, chain interactions, and disk storage.po on linking crate. Stars:`1.1K`.
* [electrumrs](https://github.com/romanz/electrs) - An efficient re-implementation of Electrum Server. Stars:`1.0K`.
* [Parity-Bitcoin](https://github.com/paritytech/parity-bitcoin) - The Parity Bitcoin client Stars:`727`.
* [Forest](https://github.com/ChainSafe/forest) - Filecoin implementation [![Build Status](https://img.shields.io/circleci/build/gh/ChainSafe/forest/main?branch=master)](https://app.circleci.com/pipelines/github/ChainSafe/forest?branch=main) Stars:`618`.
* [wagyu](https://github.com/howardwu/wagyu) [[wagyu](https://crates.io/crates/wagyu)] - Library for generating cryptocurrency wallets Stars:`606`.
* [tendermint-rs](https://github.com/informalsystems/tendermint-rs) - Tendermint blockchain data structures and clients Stars:`591`.
* [mev-inspect-rs](https://github.com/flashbots/mev-inspect-rs) - Ethereum MEV Inspector. Stars:`544`.
* [madara](https://github.com/keep-starknet-strange/madara) - Kaioshin is a ⚡ blazing fast ⚡ Starknet sequencer, based on substrate. [![GitHub Workflow Status](https://github.com/keep-starknet-strange/madara/actions/workflows/test.yml/badge.svg)](https://github.com/keep-starknet-strange/madara/actions/workflows/test.yml) Stars:`527`.
* [ethabi](https://github.com/rust-ethereum/ethabi) - Encode and decode smart contract invocations. Stars:`511`.
* [cairo-vm](https://github.com/lambdaclass/cairo-vm) - Implementation of the Cairo VM [![rust](https://github.com/lambdaclass/cairo-vm/actions/workflows/rust.yml/badge.svg)](https://github.com/lambdaclass/cairo-vm/actions/workflows/rust.yml) Stars:`485`.
* [ibc-rs](https://github.com/informalsystems/hermes) - Implementation of the [Interblockchain Communication](https://ibc.cosmos.network/) protocol Stars:`434`.
* [Subspace](https://github.com/subspace/subspace) - The first layer-one blockchain that can fully resolve the blockchain trilemma by simultaneously achieving scalability, security, and decentralization. Stars:`366`.
* [etk](https://github.com/quilt/etk) - etk is a collection of tools for writing, reading, and analyzing EVM bytecode. Stars:`350`.
* [Phala-Network/phala-blockchain](https://github.com/Phala-Network/phala-blockchain) - Confidential smart contract blockchain based on Intel SGX and Substrate Stars:`330`.
* [ChainX](https://github.com/chainx-org/ChainX) - Fully Decentralized Interchain Crypto Asset Management on Polkadot. Stars:`321`.
* [interBTC](https://github.com/interlay/interbtc) - Trustless and fully decentralized Bitcoin bridge to Polkadot and Kusama. Stars:`245`.
* [opensea-rs](https://github.com/gakonst/opensea-rs) - Bindings & CLI to the Opensea API and Contracts. Stars:`242`.
* [beerus](https://github.com/eigerco/beerus) - Beerus is a trustless StarkNet Light Client, ⚡blazing fast ⚡ [![GitHub Workflow Status](https://github.com/eigerco/beerus/actions/workflows/test.yml/badge.svg)](https://github.com/eigerco/beerus/actions/workflows/test.yml) Stars:`241`.
* [Hyperlane](https://github.com/hyperlane-xyz/hyperlane-monorepo) - Framework for permissionless, modular interoperability. The offchain clients are written in Rust, as well as the smart contracts for Solana VM and CosmWasm. Stars:`240`.
* [svm-rs](https://github.com/alloy-rs/svm-rs) - Solidity-Compiler Version Manager. Stars:`239`.
* [ethaddrgen](https://github.com/Limeth/ethaddrgen) - Custom Ethereum vanity address generator Stars:`175`.
* [coinbase-pro-rs](https://github.com/inv2004/coinbase-pro-rs) - Coinbase pro client, supports sync/async/websocket Stars:`145`.
* [sigma-rust](https://github.com/ergoplatform/sigma-rust) - ErgoTree interpreter and wallet-related features. Stars:`70`.
* [Bitcoin Satoshi's Vision](https://github.com/brentongunning/rust-sv) [[sv](https://crates.io/crates/sv)] - A library for working with Bitcoin SV. Stars:`60`.
* [infincia/bip39-rs](https://github.com/infincia/bip39-rs) [[bip39](https://crates.io/crates/bip39)] - Implementation of BIP39. Stars:`54`.
* [hdwallet](https://github.com/jjyr/hdwallet) [[hdwallet](https://crates.io/crates/hdwallet)] - BIP-32 HD wallet related key derivation utilities. Stars:`33`.

### Database

* [SurrealDB](https://github.com/surrealdb/surrealdb) - A scalable, distributed, document-graph database [![Build Status](https://img.shields.io/github/workflow/status/surrealdb/surrealdb/Continuous%20integration/main)](https://github.com/surrealdb/surrealdb/actions) Stars:`26.2K`.
* [Qdrant](https://github.com/qdrant/qdrant) - An open source vector similarity search engine with extended filtering support [![Tests](https://github.com/qdrant/qdrant/workflows/Tests/badge.svg)](https://github.com/qdrant/qdrant/actions) Stars:`18.8K`.
* [tikv](https://github.com/tikv/tikv) - A distributed KV database in Rust [![Build Status](https://ci.pingcap.net/job/tikv_ghpr_test/badge/icon)](https://ci.pingcap.net/job/tikv_ghpr_test/) Stars:`14.8K`.
* [Neon](https://github.com/neondatabase/neon) - Serverless Postgres. We separated storage and compute to offer autoscaling, branching, and bottomless storage. Stars:`13.3K`.
* [sled](https://crates.io/crates/sled) - A (beta) modern embedded database [![Build Status](https://github.com/spacejam/sled/workflows/Rust/badge.svg?branch=master)](https://github.com/spacejam/sled/actions?workflow=Rust) Stars:`7.9K`.
* [Databend](https://github.com/datafuselabs/databend) - A Modern Real-Time Data Processing & Analytics DBMS with Cloud-Native Architecture [![Release](https://github.com/datafuselabs/databend/actions/workflows/databend-release.yml/badge.svg)](https://github.com/datafuselabs/databend/actions/workflows/databend-release.yml) Stars:`7.4K`.
* [RisingWaveLabs/RisingWave](https://github.com/RisingWaveLabs/risingwave) - the next-generation streaming database in the cloud [![CI](https://github.com/RisingWaveLabs/risingwave/actions/workflows/main.yml/badge.svg)](https://github.com/RisingWaveLabs/risingwave/actions/workflows/main.yml/badge.svg?branch=main) Stars:`6.6K`.
* [erikgrinaker/toydb](https://github.com/erikgrinaker/toydb) - Distributed SQL database, written as a learning project. Stars:`6.0K`.
* [Materialize](https://github.com/MaterializeInc/materialize) - Streaming SQL database powered by Timely Dataflow :heavy_dollar_sign: [![Build status](https://badge.buildkite.com/97d6604e015bf633d1c2a12d166bb46f3b43a927d3952c999a.svg?branch=main)](https://buildkite.com/materialize/test) Stars:`5.7K`.
* [noria](https://github.com/mit-pdos/noria) [[noria](https://crates.io/crates/noria)] - Dynamically changing, partially-stateful data-flow for web application backends Stars:`5.0K`.
* [ParadeDB](https://github.com/paradedb/paradedb/) - ParadeDB is an Elasticsearch alternative built on Postgres, designed for real-time search and analytics. Stars:`4.7K`.
* [GreptimeDB](https://github.com/grepTimeTeam/greptimedb/) - An open-source, cloud-native, distributed time-series database with PromQL/SQL/Python supported.[![CI](https://github.com/greptimeTeam/greptimedb/actions/workflows/develop.yml/badge.svg)](https://github.com/greptimeTeam/greptimedb/actions/workflows/develop.yml) Stars:`4.0K`.
* [lancedb](https://github.com/lancedb/lancedb) [[vectordb](https://crates.io/crates/vectordb)] - A serverless, low-latency vector database for AI applications Stars:`3.5K`.
* [CozoDB](https://github.com/cozodb/cozo) - A transactional, relational database that uses Datalog and focuses on graph data and algorithms. Time-travel-capable, and fast! [![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/cozodb/cozo/build.yml?branch=main)](https://github.com/cozodb/cozo/actions/workflows/build.yml) Stars:`3.2K`.
* [Skytable](https://github.com/skytable/skytable) - A multi-model NoSQL database ![GitHub Workflow Status](https://img.shields.io/github/workflow/status/skytable/skytable/Tests?style=flat-square) Stars:`2.3K`.
* [SQLSync](https://github.com/orbitinghail/sqlsync) - Multiplayer offline-first SQLite [![GitHub Workflow Status](https://github.com/orbitinghail/sqlsync/actions/workflows/actions.yaml/badge.svg?branch=main)](https://github.com/orbitinghail/sqlsync/actions?query=branch%3Amain) Stars:`2.2K`.
* [USearch](https://github.com/unum-cloud/usearch) - Similarity Search Engine for Vectors and Strings [![crates.io](https://img.shields.io/crates/v/usearch.svg)](https://crates.io/crates/usearch) Stars:`1.9K`.
* [seppo0010/rsedis](https://github.com/seppo0010/rsedis) - A Redis reimplementation. Stars:`1.8K`.
* [PumpkinDB](https://github.com/PumpkinDB/PumpkinDB) - an event sourcing database engine Stars:`1.4K`.
* [Atomic-Server](https://github.com/atomicdata-dev/atomic-server/) [[atomic-server](https://crates.io/crates/atomic_server)] - NoSQL graph database with realtime updates, dynamic indexing and easy-to-use GUI for CMS purposes. [![Release](https://github.com/atomicdata-dev/atomic-server/actions/workflows/docker.yml/badge.svg)](https://github.com/atomicdata-dev/atomic-server/actions/workflows/docker.yml) Stars:`866`.
* [FnckSQL](https://github.com/KipData/FnckSQL) - SQL as a Function for Rust Stars:`512`.
* [darkbird](https://github.com/Rustixir/darkbird) [[darkbird](https://crates.io/crates/darkbird)] - HighConcurrency, RealTime, InMemory storage inspired by erlang mnesia Stars:`417`.
* [Garage](https://github.com/deuxfleurs-org/garage) [[garage](https://crates.io/crates/garage)] - S3-compatible distributed object storage service designed for self-hosting at a small-to-medium scale. [![status-badge](https://woodpecker.deuxfleurs.fr/api/badges/1/status.svg)](https://woodpecker.deuxfleurs.fr/repos/1) Stars:`402`.
* [Lucid](https://github.com/lucid-kv/lucid) - High performance and distributed KV store accessible through a HTTP API. [![Build Status](https://github.com/lucid-kv/lucid/workflows/Lucid/badge.svg?branch=master)](https://github.com/lucid-kv/lucid/actions?workflow=Lucid) Stars:`371`.
* [TerminusDB](https://github.com/terminusdb/terminusdb-store) - open source graph database and document store [![Build Status](https://github.com/terminusdb/terminusdb-store/workflows/Build/badge.svg?branch=master)](https://github.com/terminusdb/terminusdb-store/actions) Stars:`360`.
* [native_db](https://github.com/vincent-herlemont/native_db) [[native_db](https://crates.io/crates/native_db)] - Drop-in, embedded database for multi-platform apps (server, desktop, mobile). Sync Rust types effortlessly Stars:`352`.
* [DB3 Network](https://github.com/dbpunk-labs/db3) - DB3 is a community-driven blockchain layer2 decentralized database network ![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/dbpunk-labs/db3/ci.yml?branch=main&style=flat-square) Stars:`347`.
* [ParityDB](https://github.com/paritytech/parity-db) - Fast and reliable database, optimised for read operation Stars:`263`.
* [WooriDB](https://github.com/naomijub/wooridb) - General purpose time serial database inspired by Crux and Datomic. Stars:`131`.
* [vorot93/libmdbx-rs](https://github.com/vorot93/libmdbx-rs) [[mdbx-sys](https://crates.io/crates/mdbx-sys)] - Bindings for MDBX, a "fast, compact, powerful, embedded, transactional key-value database, with permissive license". This is a fork of mozilla/lmdb-rs with patches to make it work with libmdbx. Stars:`76`.
* [Qrlew/qrlew](https://github.com/Qrlew/qrlew) [[qrlew](https://crates.io/crates/qrlew)] - The SQL-to-SQL Differential Privacy layer [![Qrlew](https://github.com/Qrlew/qrlew/actions/workflows/ci.yml/badge.svg)](https://github.com/Qrlew/qrlew/actions) ![Crates.io Version](https://img.shields.io/crates/v/qrlew?logo=Rust) Stars:`37`.
* [indradb](https://crates.io/crates/indradb) - Graph database

### Emulators

See also [crates matching keyword 'emulator'](https://crates.io/keywords/emulator).

* CHIP-8
  * [ColinEberhardt/wasm-rust-chip8](https://github.com/ColinEberhardt/wasm-rust-chip8) - A WebAssembly CHIP-8 emulator. Stars:`256`.
  * [starrhorne/chip8-rust](https://github.com/starrhorne/chip8-rust) - chip8 emulator Stars:`141`.
* Commodore 64
  * [kondrak/rust64](https://github.com/kondrak/rust64) - Commodore 64 emulator Stars:`256`.
* Flash Player
  * [Ruffle](https://github.com/ruffle-rs/ruffle) - Ruffle is an Adobe Flash Player emulator. Ruffle targets both the desktop and the web using WebAssembly. [![CI](https://github.com/ruffle-rs/ruffle/actions/workflows/test_rust.yml/badge.svg)](https://github.com/ruffle-rs/ruffle/actions/workflows/test_rust.yml)[![CI](https://github.com/ruffle-rs/ruffle/actions/workflows/test_web.yml/badge.svg)](https://github.com/ruffle-rs/ruffle/actions/workflows/test_web.yml) Stars:`15.1K`.
* Gameboy
  * [mohanson/gameboy](https://github.com/mohanson/gameboy) - Full featured Cross-platform GameBoy emulator. Forever boys!. Stars:`1.3K`.
  * [Gekkio/mooneye-gb](https://github.com/Gekkio/mooneye-gb) - A Game Boy research project and emulator Stars:`886`.
  * [joamag/boytacean](https://github.com/joamag/boytacean) - GameBoy Color emulator that runs on the Web using WebAssembly. Stars:`614`.
  * [mvdnes/rboy](https://github.com/mvdnes/rboy) - A Gameboy Emulator Stars:`579`.
* Gameboy Advance
  * [michelhe/rustboyadvance-ng](https://github.com/michelhe/rustboyadvance-ng) - RustboyAdvance-ng is a Gameboy Advance emulator with desktop, android and [WebAssembly](https://michelhe.github.io/rustboyadvance-ng/) support. [![build badge](https://github.com/michelhe/rustboyadvance-ng/workflows/Deploy/badge.svg?branch=master)](https://github.com/michelhe/rustboyadvance-ng/actions?query=workflow%3ADeploy) Stars:`586`.
* GameMaker
  * [OpenGMK](https://github.com/OpenGMK/OpenGMK) - OpenGMK is a modern rewrite of the proprietary GameMaker Classic engines, providing a full sourceport of the runner, a decompiler, a TASing framework, and libraries for working with gamedata yourself. Stars:`265`.
* IBM PC
  * [MartyPC](https://github.com/dbalsom/martypc) - An IBM PC/XT emulator written in Rust. Stars:`504`.
* Intel 8080 CPU
  * [mohanson/i8080](https://github.com/mohanson/i8080) - Intel 8080 CPU emulator Stars:`109`.
* iOS
  * [touchHLE](https://github.com/touchHLE/touchHLE) - High-level emulator for iPhone OS apps Stars:`2.6K`.
* iPod
  * [clicky](https://github.com/daniel5151/clicky) - A clickwheel iPod emulator (WIP) Stars:`157`.
* NES
  * [koute/pinky](https://github.com/koute/pinky) - A NES emulator Stars:`766`.
  * [pcwalton/sprocketnes](https://github.com/pcwalton/sprocketnes) - A NES emulator Stars:`749`.
* Nintendo 64
  * [gopher64](https://github.com/gopher64/gopher64) - N64 emulator written in Rust Stars:`81`.
* Nintendo DS
  * [dust](https://github.com/kelpsyberry/dust) - A Nintendo DS emulator Stars:`192`.
* PlayStation 4
  * [Obliteration](https://github.com/obhq/obliteration) - Experimental PS4 emulator for Windows, macOS and Linux [![CI](https://github.com/obhq/obliteration/actions/workflows/main.yml/badge.svg)](https://github.com/obhq/obliteration/actions/workflows/main.yml) Stars:`585`.
* ZX Spectrum
  * [rustzx/rustzx](https://github.com/rustzx/rustzx) - [![RustZX CI](https://github.com/rustzx/rustzx/actions/workflows/ci.yml/badge.svg)](https://github.com/rustzx/rustzx/actions/workflows/ci.yml) Stars:`195`.

### File manager

* [yazi](https://github.com/sxyazi/yazi) - Blazing fast terminal file manager, based on async I/O. Stars:`11.9K`.
* [broot](https://github.com/Canop/broot) - A new way to see and navigate directory trees (get an overview of a directory, even a big one; find a directory then `cd` to it; never lose track of file hierarchy while you search; manipulate your files, ...), further reading [dystroy.org/broot](https://dystroy.org/broot/) [![Latest Version](https://img.shields.io/crates/v/broot.svg)](https://crates.io/crates/broot) Stars:`10.3K`.
* [xplr](https://github.com/sayanarijit/xplr) - A hackable, minimal, fast TUI file explorer Stars:`4.1K`.
* [joshuto](https://github.com/kamiyaa/joshuto) - ranger-like terminal file manager Stars:`3.3K`.

### Games

See also [Games Made With Piston](https://github.com/PistonDevelopers/piston/wiki/Games-Made-With-Piston).

* [citybound](https://github.com/citybound/citybound) - The city sim you deserve Stars:`7.7K`.
* [cristicbz/rust-doom](https://github.com/cristicbz/rust-doom) - A renderer for Doom, may progress to being a playable game Stars:`2.3K`.
* [mtkennerly/ludusavi](https://github.com/mtkennerly/ludusavi) - Backup tool for PC game saves [![build badge](https://img.shields.io/github/actions/workflow/status/mtkennerly/ludusavi/main.yaml?logo=github)](https://github.com/mtkennerly/ludusavi/actions/workflows/main.yaml) [![crate](https://img.shields.io/crates/v/ludusavi?logo=rust)](https://crates.io/crates/ludusavi) Stars:`2.1K`.
* [ozkriff/zemeroth](https://github.com/ozkriff/zemeroth) - A small 2D turn-based hexagonal strategy game Stars:`1.4K`.
* [gorilla-devs/ferium](https://github.com/gorilla-devs/ferium) - Ferium is a fast and feature rich CLI program for downloading and updating Minecraft mods from Modrinth, CurseForge, and GitHub Releases, and modpacks from Modrinth and CurseForge ![ferium build](https://github.com/gorilla-devs/ferium/actions/workflows/build.yml/badge.svg?branch=main) Stars:`1.1K`.
* [doukutsu-rs](https://github.com/doukutsu-rs/doukutsu-rs) - Reimplementation of Cave Story engine with some enhancements. Stars:`845`.
* [garkimasera/rusted-ruins](https://github.com/garkimasera/rusted-ruins) - Extensible open world rogue like game with pixel art Stars:`484`.
* [mcthesw/game-save-manager](https://github.com/mcthesw/game-save-manager) - A user-friendly tool for managing game saves [![build badge](https://github.com/mcthesw/game-save-manager/actions/workflows/tauri.yml/badge.svg)](https://github.com/mcthesw/game-save-manager/actions/workflows/tauri.yml) Stars:`448`.
* [Zone of Control](https://github.com/ozkriff/zoc) - A turn-based hexagonal strategy game Stars:`374`.
* [rsaarelm/magog](https://github.com/rsaarelm/magog) - A roguelike game. Stars:`368`.
* [chess-tui](https://github.com/thomas-mauran/chess-tui) - A Chess TUI implementation ♟️ Stars:`342`.
* [SoftbearStudios/mk48](https://github.com/SoftbearStudios/mk48) - Mk48.io is an online multiplayer naval combat game Stars:`315`.
* [thetawavegame/thetawave-legacy](https://github.com/thetawavegame/thetawave-legacy) - A space shooter game that strives to be an entry point for new game developers to make their first contributions. ![build badge](https://github.com/thetawavegame/thetawave-legacy/actions/workflows/ci.yml/badge.svg?branch=master) Stars:`193`.
* [rhex](https://github.com/dpc/rhex) - hexagonal ascii roguelike Stars:`153`.
* [swatteau/sokoban-rs](https://github.com/swatteau/sokoban-rs) - A Sokoban implementation Stars:`151`.
* [maras-archive/rsnake](https://github.com/maras-archive/rsnake) - Snake. Stars:`129`.
* [lifthrasiir/angolmois-rust](https://github.com/lifthrasiir/angolmois-rust) - A minimalistic music video game which supports the BMS format Stars:`101`.
* [buxx/OpenCombat](https://github.com/buxx/OpenCombat) - A realtime 2nd world war tactical game Stars:`78`.
* [Thinkofname/rust-quake](https://github.com/Thinkofname/rust-quake) - Quake map renderer. Stars:`70`.
* [HactarCE/Hyperspeedcube](https://github.com/HactarCE/Hyperspeedcube) - A modern, beginner-friendly 3D and 4D Rubik's cube simulator with customizable mouse and keyboard controls and advanced features for speedsolving Stars:`57`.
* [ttyperacer/terminal-typeracer](https://gitlab.com/ttyperacer/terminal-typeracer) - Single player typing test game written for the terminal
* [Veloren](https://gitlab.com/veloren/veloren) - An open world, open source multiplayer voxel RPG game currently in alpha development [![build badge](https://gitlab.com/veloren/veloren/badges/master/pipeline.svg)](https://gitlab.com/veloren/veloren/-/pipelines)

### Graphics

* [flxzt/rnote](https://github.com/flxzt/rnote) - Sketch and take handwritten notes. Stars:`6.5K`.
* [ivanceras/svgbob](https://github.com/ivanceras/svgbob) - converts ASCII diagrams into SVG graphics Stars:`3.8K`.
* [RazrFalcon/resvg](https://github.com/RazrFalcon/resvg) - An SVG rendering library. Stars:`2.6K`.
* [wahn/rs_pbrt](https://github.com/wahn/rs_pbrt) - Implements a counterpart to the PBRT book's (3rd edition) C++ code. Stars:`804`.
* [turnage/valora](https://crates.io/crates/valora) - A library for generative fine art ![Rust](https://github.com/turnage/valora/workflows/Rust/badge.svg?branch=master) Stars:`677`.
* [Twinklebear/tray_rust](https://github.com/Twinklebear/tray_rust) - A ray tracer Stars:`514`.
* [rustq/vue-skia](https://github.com/rustq/vue-skia) - Skia based 2d graphics vue rendering library. It is based on Rust to implement software rasterization to perform rendering. Stars:`263`.
* [dps/rust-raytracer](https://github.com/dps/rust-raytracer) - An implementation of a very simple raytracer based on Ray Tracing in One Weekend by Peter Shirley. Stars:`234`.
* [Limeth/euclider](https://github.com/Limeth/euclider) - A real-time 4D CPU ray tracer Stars:`212`.
* [rodrigorc/papercraft](https://github.com/rodrigorc/papercraft) - A tool to unwrap 3D models and create them in paper with scissors and glue. Stars:`112`.
* [KaminariOS/rustracer](https://github.com/KaminariOS/rustracer) - A PBR glTF 2.0 renderer based on Vulkan ray-tracing. Stars:`64`.

### Image processing

* [shssoichiro/oxipng](https://github.com/shssoichiro/oxipng) [[oxipng](https://crates.io/crates/oxipng)] - Multithreaded PNG optimizer written in Rust. [![Build Status](https://github.com/shssoichiro/oxipng/workflows/oxipng/badge.svg)](https://github.com/shssoichiro/oxipng/actions?query=branch%3Amaster) [![Version](https://img.shields.io/crates/v/oxipng.svg)](https://crates.io/crates/oxipng) Stars:`2.7K`.
* [Imager](https://github.com/imager-io/imager) - Automated image optimization. Stars:`627`.

### Industrial automation

* [locka99/opcua](https://github.com/locka99/opcua) - A [OPC UA](https://opcfoundation.org/about/opc-technologies/opc-ua/) library. Stars:`476`.
* [slowtec/tokio-modbus](https://github.com/slowtec/tokio-modbus) - A [tokio](https://tokio.rs)-based [modbus](https://modbus.org) library. Stars:`379`.

### Observability

* [vectordotdev/vector](https://github.com/vectordotdev/vector) - A High-Performance, Logs, Metrics, & Events Router. Stars:`17.0K`.
* [openobserve](https://github.com/openobserve/openobserve) - 10x easier, 140x lower storage cost, high performance, petabyte scale - Elasticsearch/Splunk/Datadog alternative. Stars:`10.2K`.
* [Quickwit-oss/quickwit](https://github.com/quickwit-oss/quickwit) - Cloud-native and highly cost-efficient search engine for log management. [![CI](https://github.com/quickwit-oss/quickwit/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/quickwit-oss/quickwit/actions?query=workflow%3ACI) Stars:`7.0K`.
* [OpenTelemetry](https://crates.io/crates/opentelemetry) - OpenTelemetry provides a single set of APIs, libraries, agents, and collector services to capture distributed traces and metrics from your application. You can analyze them using Prometheus, Jaeger, and other observability tools. [![GitHub Actions CI](https://github.com/open-telemetry/opentelemetry-rust/workflows/CI/badge.svg?branch=master)](https://github.com/open-telemetry/opentelemetry-rust/actions?query=workflow%3ACI+branch%3Amaster) Stars:`1.7K`.
* [Scaphandre](https://github.com/hubblo-org/scaphandre) - A power consumption monitoring agent, to track host and each service power consumption and enable designing systems and applications for more sustainability. Designed to fit any monitoring toolchain (already supports prometheus, warp10, riemann...). Stars:`1.5K`.
* [avito-tech/bioyino](https://github.com/avito-tech/bioyino) - A high-performance scalable StatsD compatible server. Stars:`226`.

### Operating systems

See also [A comparison of operating systems written in Rust](https://github.com/flosse/rust-os-comparison).

* [tock/tock](https://github.com/tock/tock) - A secure embedded operating system for Cortex-M based microcontrollers Stars:`5.2K`.
* [theseus-os/Theseus](https://github.com/theseus-os/Theseus) - A safe-language, single address space and single privilege level OS written from scratch - [![build badge](https://img.shields.io/github/workflow/status/theseus-os/Theseus/Documentation?label=docs%20build)](https://www.theseus-os.com/Theseus/book/index.html) Stars:`2.8K`.
* [Andy-Python-Programmer/aero](https://github.com/Andy-Python-Programmer/aero) - A modern, unix-like operating system following the monolithic kernel design. Stars:`1.1K`.
* [DragonOS-Community/DragonOS](https://github.com/DragonOS-Community/DragonOS) - An operating system with a self-developed kernel from scratch and Linux compatibility. Stars:`820`.
* [thepowersgang/rust_os](https://github.com/thepowersgang/rust_os) - An OS kernel written in rust. Non POSIX Stars:`715`.
* [0x59616e/SteinsOS](https://github.com/0x59616e/SteinsOS) - An OS for armv8-a architecture. Stars:`115`.
* [redox-os/redox](https://gitlab.redox-os.org/redox-os/redox) - A rust operating system

### Package Managers

* [helsing-ai/buffrs](https://github.com/helsing-ai/buffrs) [[buffrs](https://crates.io/crates/buffrs)] - A modern package manager for protocol buffers and gRPC architectures. Stars:`160`.

### Payments

* [hyperswitch](https://github.com/juspay/hyperswitch) - An open source payments orchestrator that lets you connect with multiple payment processors and route payment traffic effortlessly, all with a single API integration ![GitHub last commit](https://img.shields.io/github/last-commit/juspay/hyperswitch?style=flat-square) Stars:`11.1K`.

### Productivity

* [espanso](https://github.com/espanso/espanso) - A cross-platform Text Expander. [![CI](https://github.com/espanso/espanso/actions/workflows/ci.yml/badge.svg?branch=dev&event=push)](https://github.com/espanso/espanso/actions/workflows/ci.yml) Stars:`9.5K`.
* [ast-grep](https://github.com/ast-grep/ast-grep) - A CLI tool for code structural search, lint and rewriting. Stars:`6.4K`.
* [LLDAP](https://github.com/lldap/lldap) - Simplified LDAP interface for authentication. Stars:`3.7K`.
* [Bartib](https://github.com/nikolassv/bartib) [[Bartib](https://crates.io/crates/bartib)] - A simple timetracker for the command line [![Tests](https://github.com/nikolassv/bartib/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/nikolassv/bartib/actions/workflows/test.yml) Stars:`635`.
* [pier-cli/pier](https://github.com/pier-cli/pier) - A central repository to manage (add, search metadata, etc.) all your one-liners, scripts, tools, and CLIs Stars:`518`.
* [Furtherance](https://github.com/lakoliu/Furtherance) - Time tracking app built with GTK4 Stars:`242`.
* [illacloud/illa](https://github.com/illacloud/illa) [[ILLA Cloud](https://illacloud.com)] - Low-code internal tool builder. Stars:`225`.
* [yashs662/rust_kanban](https://github.com/yashs662/rust_kanban) [[rust-kanban](https://crates.io/crates/rust-kanban)] [![Build](https://github.com/yashs662/rust_kanban/actions/workflows/build.yml/badge.svg)](https://github.com/yashs662/rust_kanban/releases) - A Kanban App for the terminal Stars:`153`.
* [ShadoySV/work-break](https://github.com/ShadoySV/work-break) [[work-break](https://crates.io/crates/work-break)] - Work and rest time balancer taking into account your current and today strain [![Build](https://github.com/shadoysv/work-break/actions/workflows/release.yml/badge.svg)](https://github.com/ShadoySV/work-break/releases) Stars:`5`.
* [eureka](https://crates.io/crates/eureka) - A CLI tool to input and store your ideas without leaving the terminal

### Routing protocols

* [RustyBGP](https://github.com/osrg/rustybgp) - BGP Stars:`468`.
* [Holo](https://github.com/holo-routing/holo) - Holo is a suite of routing protocols designed to support high-scale and automation-driven networks Stars:`239`.

### Security tools

* [rustscan/rustscan](https://github.com/RustScan/RustScan) - Make Nmap faster with this port scanning tool [![build badge](https://github.com/RustScan/RustScan/workflows/Continuous%20integration/badge.svg?branch=master)](https://github.com/RustScan/RustScan/actions?query=workflow%3A%22Continuous+integration%22) Stars:`13.3K`.
* [epi052/feroxbuster](https://github.com/epi052/feroxbuster) - A simple, fast, recursive content discovery tool. Stars:`5.6K`.
* [kpcyrd/sn0int](https://github.com/kpcyrd/sn0int) - A semi-automatic OSINT framework and package manager Stars:`1.9K`.
* [AFLplusplus/LibAFL](https://github.com/AFLplusplus/LibAFL) - Advanced Fuzzing Library - Slot your Fuzzer together in Rust! Scales across cores and machines. For Windows, Android, MacOS, Linux, no_std, etc. [![build and test](https://github.com/AFLplusplus/LibAFL/actions/workflows/build_and_test.yml/badge.svg)](https://github.com/AFLplusplus/LibAFL/actions/workflows/build_and_test.yml) Stars:`1.9K`.
* [Cherrybomb](https://github.com/blst-security/cherrybomb) - Stop half-done API specifications with a CLI tool that helps you avoid undefined user behaviour by validating your API specifications. Stars:`1.1K`.
* [kpcyrd/sniffglue](https://github.com/kpcyrd/sniffglue) - A secure multithreaded packet sniffer Stars:`1.1K`.
* [ObserverWard](https://github.com/0x727/ObserverWard) - Community based web technologies analysis tool. Stars:`1.1K`.
* [ripasso](https://github.com/cortex/ripasso/) - A password manager, filesystem compatible with pass Stars:`692`.
* [kpcyrd/rshijack](https://github.com/kpcyrd/rshijack) - A TCP connection hijacker; rewrite of shijack Stars:`449`.
* [kpcyrd/authoscope](https://github.com/kpcyrd/authoscope) - A scriptable network authentication cracker Stars:`388`.
* [Inspektor](https://github.com/inspektor-dev/inspektor) - A database protocol-aware proxy that is used to enforce access policies 👮 Stars:`279`.
* [cotp](https://github.com/replydev/cotp) - Trustworthy, encrypted, command-line TOTP/HOTP authenticator app with import functionality. Stars:`220`.
* [Raspirus](https://github.com/Raspirus/Raspirus) - User- and resources-friendly signatures-based malware scanner [![status](https://github.com/Raspirus/Raspirus/actions/workflows/testproject.yml/badge.svg)](https://github.com/Raspirus/Raspirus/actions/workflows/testproject.yml) Stars:`99`.
* [arp-scan-rs](https://github.com/kongbytes/arp-scan-rs) - A minimalistic ARP scan tool for fast local network scans Stars:`81`.
* [entropic-security/xgadget](https://github.com/entropic-security/xgadget) [[xgadget](https://crates.io/crates/xgadget)] - Fast, parallel, cross-variant ROP/JOP gadget search [![GitHub Actions](https://github.com/entropic-security/xgadget/workflows/test/badge.svg)](https://github.com/entropic-security/xgadget/actions) Stars:`79`.
* [cargo-deny](https://crates.io/crates/cargo-deny) - Cargo plugin to help you manage large dependency graphs
* [cargo-crev](https://crates.io/crates/cargo-crev) - A cryptographically verifiable code review system for the cargo package manager.
* [cargo-auditable](https://crates.io/crates/cargo-auditable) - Make production Rust binaries auditable
* [cargo-audit](https://crates.io/crates/cargo-audit) - Audit Cargo.lock for crates with security vulnerabilities

### Social networks

* Mastodon
  * [Rustodon](https://github.com/rustodon/rustodon) - A Mastodon-compatible, ActivityPub-speaking server. Stars:`866`.

### System tools

* [sharkdp/bat](https://github.com/sharkdp/bat) - A cat(1) clone with wings. [![CICD](https://github.com/sharkdp/bat/actions/workflows/CICD.yml/badge.svg?branch=master)](https://github.com/sharkdp/bat/actions/workflows/CICD.yml) Stars:`47.7K`.
* [sharkdp/fd](https://github.com/sharkdp/fd) - A simple, fast and user-friendly alternative to find. [![CICD](https://github.com/sharkdp/fd/actions/workflows/CICD.yml/badge.svg)](https://github.com/sharkdp/fd/actions/workflows/CICD.yml) Stars:`32.6K`.
* [nushell/nushell](https://github.com/nushell/nushell) - A new type of shell Stars:`30.9K`.
* [ajeetdsouza/zoxide](https://github.com/ajeetdsouza/zoxide/) - A fast alternative to `cd` that learns your habits [![release](https://github.com/ajeetdsouza/zoxide/workflows/.github/workflows/release.yml/badge.svg)](https://github.com/ajeetdsouza/zoxide/actions) Stars:`20.4K`.
* [atuin](https://github.com/atuinsh/atuin) [[atuin](https://crates.io/crates/atuin)] - Atuin replaces your existing shell history with a SQLite database, and records additional context for your commands. Additionally, it provides optional and fully encrypted synchronisation of your history between machines, via an Atuin server. Stars:`18.7K`.
* [qarmin/czkawka](https://github.com/qarmin/czkawka) - Multi-functional app to find duplicates, empty folders, similar images, etc. [![GitHub Actions Workflow](https://github.com/qarmin/czkawka/actions/workflows/pages/pages-build-deployment/badge.svg?branch=master)](https://github.com/qarmin/czkawka/actions) Stars:`18.4K`.
* [gitui](https://github.com/extrawurst/gitui) - Blazing fast terminal client for git. [![build](https://github.com/extrawurst/gitui/workflows/CI/badge.svg?branch=master)](https://github.com/extrawurst/gitui/actions) Stars:`17.8K`.
* [uutils/coreutils](https://github.com/uutils/coreutils) - A cross-platform rewrite of the GNU coreutils [![CICD](https://github.com/uutils/coreutils/actions/workflows/CICD.yml/badge.svg)](https://github.com/uutils/coreutils/actions/workflows/CICD.yml) Stars:`17.3K`.
* [lsd](https://github.com/lsd-rs/lsd) - An ls with a lot of pretty colors and awesome icons [![build](https://github.com/lsd-rs/lsd/workflows/CICD/badge.svg?branch=master)](https://github.com/lsd-rs/lsd/actions) Stars:`12.8K`.
* [XAMPPRocky/tokei](https://github.com/XAMPPRocky/tokei) - counts the lines of code Stars:`10.6K`.
* [bottom](https://github.com/ClementTsang/bottom) - Yet another cross-platform graphical process/system monitor. [![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/ClementTsang/bottom/ci/master)](https://github.com/ClementTsang/bottom/actions?query=branch%3Amaster) Stars:`9.4K`.
* [bandwhich](https://github.com/imsnif/bandwhich) - Terminal bandwidth utilization tool Stars:`9.1K`.
* [eza-community/eza](https://github.com/eza-community/eza) - A replacement for 'ls' Stars:`8.9K`.
* [dust](https://github.com/bootandy/dust) - A more intuitive version of du Stars:`8.1K`.
* [cantino/mcfly](https://github.com/cantino/mcfly) - Fly through your shell history. Great Scott! Stars:`6.8K`.
* [watchexec](https://github.com/watchexec/watchexec) - Executes commands in response to file modifications Stars:`5.1K`.
* [lotabout/skim](https://github.com/lotabout/skim) - A fuzzy finder Stars:`5.0K`.
* [dalance/procs](https://github.com/dalance/procs) - A modern replacement for 'ps' [![Regression](https://github.com/dalance/procs/actions/workflows/regression.yml/badge.svg)](https://github.com/dalance/procs/actions/workflows/regression.yml) Stars:`4.9K`.
* [pueue](https://github.com/nukesor/pueue) - Manage your long running shell commands. [![GitHub Actions Workflow](https://github.com/nukesor/pueue/workflows/Test%20build/badge.svg?branch=master)](https://github.com/nukesor/pueue/actions) Stars:`4.7K`.
* [ynqa/jnv](https://github.com/ynqa/jnv) - Interactive JSON filter using jq [![ci](https://github.com/ynqa/jnv/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/ynqa/jnv/actions/workflows/ci.yml) Stars:`4.6K`.
* [trippy](https://github.com/fujiapple852/trippy) - A network diagnostic tool [![build badge](https://github.com/fujiapple852/trippy/workflows/CI/badge.svg)](https://github.com/fujiapple852/trippy/actions/workflows/ci.yml) Stars:`3.2K`.
* [GQL](https://github.com/amrdeveloper/gql) - A SQL like query language to run on .git files. Stars:`3.1K`.
* [orhun/kmon](https://github.com/orhun/kmon) - Linux Kernel Manager and Activity Monitor ![https://github.com/orhun/kmon/actions](https://img.shields.io/github/actions/workflow/status/orhun/kmon/ci.yml?branch=master&label=build) Stars:`2.5K`.
* [diskonaut](https://github.com/imsnif/diskonaut) - Terminal visual disk space navigator Stars:`2.3K`.
* [ouch](https://github.com/ouch-org/ouch) - Painless compression and decompression on the command-line [![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/ouch-org/ouch/build-and-test)](https://github.com/ouch-org/ouch/actions?query=branch%3Amaster) Stars:`2.1K`.
* [pkolaczk/fclones](https://github.com/pkolaczk/fclones) - Efficient duplicate file finder and remover Stars:`1.8K`.
* [Kondo](https://github.com/tbillington/kondo) - CLI & GUI tool for deleting software project artifacts and reclaiming disk space Stars:`1.7K`.
* [m4b/bingrep](https://github.com/m4b/bingrep) - Greps through binaries from various OSs and architectures, and colors them. Stars:`1.7K`.
* [redox-os/ion](https://github.com/redox-os/ion) - Next-generation system shell Stars:`1.4K`.
* [nivekuil/rip](https://github.com/nivekuil/rip) - A safe and ergonomic alternative to `rm` Stars:`1.3K`.
* [orhun/systeroid](https://github.com/orhun/systeroid) - A more powerful alternative to sysctl(8) with a terminal user interface ![https://github.com/orhun/systeroid/actions](https://img.shields.io/github/actions/workflow/status/orhun/systeroid/ci.yml?branch=main&label=build) Stars:`1.2K`.
* [httm](https://github.com/kimono-koans/httm) - Interactive, file-level Time Machine-like tool for ZFS/btrfs/nilfs2 (and even actual Time Machine backups!) Stars:`1.2K`.
* [LACT](https://github.com/ilya-zlobintsev/LACT) - Linux AMDGPU Controller Stars:`1.0K`.
* [mitnk/cicada](https://github.com/mitnk/cicada) - A bash-like Unix shell Stars:`970`.
* [Luminarys/synapse](https://github.com/Luminarys/synapse) - Flexible and fast BitTorrent daemon. Stars:`851`.
* [netscanner](https://github.com/Chleba/netscanner) - TUI Network Scanner Stars:`742`.
* [pop-os/popsicle](https://github.com/pop-os/popsicle) - GTK3 & CLI utility for flashing multiple USB devices in parallel Stars:`624`.
* [pop-os/system76-power](https://github.com/pop-os/system76-power/) - Linux power management daemon (DBus-interface) with CLI tool. Stars:`564`.
* [mdgaziur/findex](https://github.com/mdgaziur/findex) - Findex is a highly customizable application finder using GTK3 Stars:`548`.
* [ynqa/sig](https://github.com/ynqa/sig) - Interactive grep (for streaming) [![ci](https://github.com/ynqa/sig/actions/workflows/ci.yml/badge.svg)](https://github.com/ynqa/sig/actions/workflows/ci.yml) Stars:`512`.
* [sitkevij/hex](https://github.com/sitkevij/hex) - A colorized hexdump terminal utility. Stars:`504`.
* [lotabout/rargs](https://github.com/lotabout/rargs) [[rargs](https://crates.io/crates/rargs)] - xargs + awk with pattern matching support Stars:`474`.
* [ddh](https://github.com/darakian/ddh) - Fast duplicate file finder Stars:`442`.
* [brocode/fblog](https://github.com/brocode/fblog) - Small command-line JSON Log viewer Stars:`405`.
* [j0ru/kickoff](https://github.com/j0ru/kickoff) - Fast and snappy wayland program launcher [![build](https://github.com/j0ru/kickoff/actions/workflows/ci.yml/badge.svg)](https://github.com/j0ru/kickoff/actions) Stars:`359`.
* [crabz](https://github.com/sstadick/crabz) - Multi-threaded compression and decompression CLI tool [![Build Status](https://github.com/sstadick/crabz/workflows/Check/badge.svg)](https://github.com/sstadick/crabz/actions?query=workflow%3ACheck) Stars:`321`.
* [supercilex/fuc](https://github.com/supercilex/fuc) - Fast `cp` and `rm` commands Stars:`317`.
* [mmstick/fontfinder](https://github.com/mmstick/fontfinder) - GTK3 application for previewing and installing Google's fonts Stars:`276`.
* [nickgerace/gfold](https://github.com/nickgerace/gfold) [[gfold](https://crates.io/crates/gfold)] - CLI tool to help keep track of multiple Git repositories [![build](https://img.shields.io/github/workflow/status/nickgerace/gfold/merge/main)](https://github.com/nickgerace/gfold/actions?query=workflow%3Amerge+branch%3Amain) Stars:`274`.
* [cristianoliveira/funzzy](https://github.com/cristianoliveira/funzzy) - A configurable filesystem watcher inspired by [entr](http://eradman.com/entrproject/) Stars:`225`.
* [bustd](https://github.com/vrmiguel/bustd) - Lightweight process killer daemon to handle out-of-memory scenarios on Linux. [![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/vrmiguel/bustd/build-and-test)](https://github.com/vrmiguel/bustd/actions?query=branch%3Amaster) Stars:`213`.
* [mmstick/tv-renamer](https://github.com/mmstick/tv-renamer) - A tv series renaming application with an optional GTK3 frontend. Stars:`149`.
* [buster/rrun](https://github.com/buster/rrun) - A command launcher for Linux, similar to gmrun Stars:`113`.
* [mmstick/concurr](https://github.com/mmstick/concurr) - Alternative to GNU Parallel w/ a client-server architecture Stars:`103`.
* [mxseev/logram](https://github.com/mxseev/logram) - Push log files' updates to Telegram Stars:`97`.
* [lodosgroup/lpm](https://github.com/lodosgroup/lpm) - An experimental system package manager Stars:`65`.
* [fselect](https://crates.io/crates/fselect) - Find files with SQL-like queries

### Task scheduling

* [delicate](https://github.com/BinChengZhao/delicate) - A lightweight and distributed task scheduling platform. [![Build Status](https://github.com/BinChengZhao/delicate/workflows/CI/badge.svg)](https://github.com/BinChengZhao/delicate/actions) Stars:`687`.
* [tasklet](https://github.com/stav121/tasklet) [[tasklet](https://crates.io/crates/tasklet)] - A task scheduling library written in Rust ![Build Status](https://img.shields.io/github/actions/workflow/status/stav121/tasklet/rust.yml) Stars:`22`.

### Text editors

* [zed](https://github.com/zed-industries/zed) - A high-performance, multiplayer code editor from the creators of Atom and Tree-sitter. Stars:`39.7K`.
* [Lapce](https://github.com/lapce/lapce) - A modern editor with a backend. Taking inspiration from the discontinued [xi-editor](https://github.com/xi-editor/xi-editor). Stars:`33.8K`.
* [helix](https://github.com/helix-editor/helix) - A post-modern modal text editor inspired by Neovim/Kakoune. [![build badge](https://github.com/helix-editor/helix/actions/workflows/build.yml/badge.svg)](https://github.com/helix-editor/helix/actions) Stars:`31.3K`.
* [ox](https://github.com/curlpipe/ox) - An independent Rust text editor that runs in your terminal! Stars:`3.3K`.
* [emacs-ng](https://github.com/emacs-ng/emacs-ng) - Complementing the C codebase with rust code to introduce new features. Stars:`1.6K`.
* [gchp/iota](https://github.com/gchp/iota) - A simple text editor Stars:`1.6K`.
* [ilai-deutel/kibi](https://github.com/ilai-deutel/kibi) - A tiny (≤1024 LOC) text editor with syntax highlighting, incremental search and more. [![build badge](https://github.com/ilai-deutel/kibi/workflows/CI/badge.svg?branch=master)](https://github.com/ilai-deutel/kibi/actions?query=branch%3Amaster) Stars:`1.4K`.
* [mathall/rim](https://github.com/mathall/rim) - Vim-like text editor. Stars:`601`.
* [amp](https://amp.rs) - Inspired by Vi/Vim.
* [vamolessa/pepper](https://git.sr.ht/~lessa/pepper) [[pepper](https://crates.io/crates/pepper)] - An opinionated modal editor to simplify code editing from the terminal

### Text processing

* [grex](https://github.com/pemistahl/grex) - A command-line tool and library for generating regular expressions from user-provided test cases Stars:`7.0K`.
* [phiresky/ripgrep-all](https://github.com/phiresky/ripgrep-all) - ripgrep, but also search in PDFs, E-Books, Office documents, zip, tar.gz, etc. Stars:`6.4K`.
* [Melody](https://github.com/yoav-lavi/melody) - A language that compiles to regular expressions and aims to be more easily readable and maintainable [![build badge](https://github.com/yoav-lavi/melody/actions/workflows/rust.yml/badge.svg)](https://github.com/yoav-lavi/melody/actions/workflows/rust.yml) [![crates.io](https://img.shields.io/crates/v/melody_compiler?label=compiler)](https://crates.io/crates/melody_compiler) Stars:`4.6K`.
* [jqnatividad/qsv](https://github.com/jqnatividad/qsv) [[qsv](https://crates.io/crates/qsv)] - A high performance CSV data-wrangling toolkit. Forked from xsv, with 34+ additional commands & more. [![Linux build status](https://github.com/jqnatividad/qsv/actions/workflows/rust.yml/badge.svg)](https://github.com/jqnatividad/qsv/actions/workflows/rust.yml) [![Windows build status](https://github.com/jqnatividad/qsv/actions/workflows/rust-windows.yml/badge.svg)](https://github.com/jqnatividad/qsv/actions/workflows/rust-windows.yml) [![macOS build status](https://github.com/jqnatividad/qsv/actions/workflows/rust-macos.yml/badge.svg)](https://github.com/jqnatividad/qsv/actions/workflows/rust-macos.yml) Stars:`2.3K`.
* [ashvardanian/stringzilla](https://github.com/ashvardanian/StringZilla) - SIMD-accelerated string search, sort, edit distances, alignments, and generators for x86 AVX2 & AVX-512, and Arm NEON [![crates.io](https://img.shields.io/crates/v/stringzilla.svg)](https://crates.io/crates/stringzilla) Stars:`1.9K`.
* [dominikwilkowski/cfonts](https://github.com/dominikwilkowski/cfonts) [[cfonts](https://crates.io/crates/cfonts)] - Sexy ANSI fonts for the console ![build badge](https://github.com/dominikwilkowski/cfonts/actions/workflows/testing.yml/badge.svg) Stars:`1.6K`.
* [sstadick/hck](https://github.com/sstadick/hck) - A faster and more featureful drop in replacement for `cut` [![build badge](https://github.com/sstadick/hck/workflows/Check/badge.svg?branch=master)](https://github.com/sstadick/hck) Stars:`687`.
* [ruplacer](https://github.com/your-tools/ruplacer) - Find and replace text in source files [![Run tests](https://github.com/your-tools/ruplacer/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/your-tools/ruplacer/actions/workflows/test.yml) Stars:`420`.
* [vishaltelangre/ff](https://github.com/vishaltelangre/ff) - Find files (ff) by name! Stars:`328`.
* [whitfin/runiq](https://github.com/whitfin/runiq) - an efficient way to filter duplicate lines from unsorted input. Stars:`206`.
* [Lisprez/so_stupid_search](https://github.com/Lisprez/so_stupid_search) - A simple and fast string search tool for human beings Stars:`162`.
* [whitfin/bytelines](https://github.com/whitfin/bytelines) [[bytelines](https://crates.io/crates/bytelines)] - Read input lines as byte slices for high efficiency. Stars:`59`.
* [replicadse/complate](https://github.com/replicadse/complate) - An in-terminal text templating tool designed for standardizing messages (like for GIT commits). [![crates.io](https://img.shields.io/crates/v/complate.svg)](https://crates.io/crates/complate) [![crates.io](https://img.shields.io/crates/d/complate?label=crates.io%20downloads)](https://crates.io/crates/complate) [![build badge](https://github.com/replicadse/complate/workflows/pipeline/badge.svg?branch=master)](https://github.com/replicadse/complate/actions) Stars:`35`.
* [sd](https://crates.io/crates/sd) - Intuitive find & replace CLI
* [ripgrep](https://crates.io/crates/ripgrep) - combines the usability of The Silver Searcher with the raw speed of grep
* [xsv](https://crates.io/crates/xsv) - A fast CSV command line tool (slicing, indexing, selecting, searching, sampling, etc.)

### Utilities

* [rustdesk/rustdesk](https://github.com/rustdesk/rustdesk) - A remote desktop software, great alternative to TeamViewer and AnyDesk. Stars:`68.2K`.
* [vaultwarden](https://github.com/dani-garcia/vaultwarden#readme) [![Build](https://github.com/dani-garcia/vaultwarden/actions/workflows/build.yml/badge.svg)](https://github.com/dani-garcia/vaultwarden/actions/workflows/build.yml) - Alternative implementation of the Bitwarden server API written in Rust Stars:`34.8K`.
* [warpdotdev/Warp](https://github.com/warpdotdev/Warp) - :heavy_dollar_sign: Warp is a blazingly-fast modern GPU-accelerated terminal built to make you and your team more productive. Stars:`20.3K`.
* [str4d/rage](https://github.com/str4d/rage) [[rage](https://crates.io/crates/rage)] - Rust implementation of [age](https://github.com/FiloSottile/age). Stars:`2.4K`.
* [rustic-rs/rustic](https://github.com/rustic-rs/rustic) [[rustic-rs](https://crates.io/crates/rustic-rs)] - Fast, encrypted, deduplicated backups powered by Rust. [![Version](https://img.shields.io/crates/v/rustic-rs.svg)](https://crates.io/crates/rustic-rs) Stars:`1.6K`.
* [mprocs](https://github.com/pvolok/mprocs) - TUI for running multiple processes Stars:`1.4K`.
* [fcsonline/tmux-thumbs](https://github.com/fcsonline/tmux-thumbs) - A lightning fast version of tmux-fingers, copy/pasting tmux like vimium/vimperator. Stars:`877`.
* [nix-community/nix-init](https://github.com/nix-community/nix-init) - Generate Nix packages from URLs with hash prefetching, dependency inference, license detection, and more [![build-badge](https://github.com/nix-community/nix-init/actions/workflows/ci.yml/badge.svg)](https://github.com/nix-community/nix-init/actions/workflows/ci.yml) Stars:`778`.
* [suckit](https://github.com/Skallwar/suckit) - Recursively visit and download a website's content to your disk. [![Crate](https://img.shields.io/crates/v/suckit.svg?logo=rust)](https://crates.io/crates/suckit) [![Build Status](https://github.com/Skallwar/suckit/workflows/Build%20and%20test/badge.svg)](https://github.com/Skallwar/suckit/blob/master/.github/workflows/build_and_test.yml) Stars:`722`.
* [nomino](https://github.com/yaa110/nomino) - Batch rename utility for developers Stars:`556`.
* [mrjackwills/oxker](https://github.com/mrjackwills/oxker) [[oxker](https://crates.io/crates/oxker)] - A simple tui to view & control docker containers. Stars:`495`.
* [Vibe](https://github.com/thewh1teagle/vibe) - Transcribe audio or video in every language on every platform. Stars:`444`.
* [1History](https://github.com/1History/1History) - Command line interface to backup Firefox/Chrome/Safari history to one SQLite file [![Build Status](https://github.com/1History/1History/actions/workflows/CI.yml/badge.svg)](https://github.com/1History/1History/actions/workflows/CI.yml) Stars:`434`.
* [nix-community/nurl](https://github.com/nix-community/nurl) [[nurl](https://crates.io/crates/nurl)] - Generate Nix fetcher calls from repository URLs [![build-badge](https://github.com/nix-community/nurl/actions/workflows/ci.yml/badge.svg)](https://github.com/nix-community/nurl/actions/workflows/ci.yml) Stars:`416`.
* [guoxbin/dtool](https://github.com/guoxbin/dtool) - A useful command-line tool collection to assist development including conversion, codec, hashing, encryption, etc. Stars:`364`.
* [Epic Asset Manager](https://github.com/AchetaGames/Epic-Asset-Manager) - An unofficial client to install Unreal Engine, download and manage purchased assets, projects, plugins and games from the Epic Games Store. Stars:`364`.
* [Mobslide](https://github.com/thewh1teagle/mobslide) - Desktop application that turns your smartphone into presentation remote controller. Stars:`336`.
* [wthrr](https://github.com/ttytm/wthrr-the-weathercrab) - Weather companion for the terminal. [![crates.io](https://img.shields.io/crates/v/wthrr?logo=rust)](https://crates.io/crates/wthrr) Stars:`331`.
* [evansmurithi/cloak](https://github.com/evansmurithi/cloak) - A Command Line OTP (One Time Password) Authenticator application. ![CI](https://github.com/evansmurithi/cloak/workflows/CI/badge.svg) [![build badge](https://ci.appveyor.com/api/projects/status/9mlfpfru3ng4c689/branch/master?svg=true)](https://ci.appveyor.com/project/evansmurithi/cloak) Stars:`283`.
* [tversteeg/emplace](https://github.com/tversteeg/emplace) - Synchronize installed packages on multiple machines Stars:`251`.
* [vamolessa/verco](https://github.com/vamolessa/verco) [[verco](https://crates.io/crates/verco)] - A simple Git/Hg tui client focused on keyboard shortcuts Stars:`229`.
* [nix-community/nix-melt](https://github.com/nix-community/nix-melt) - A ranger-like flake.lock viewer [![build-badge](https://github.com/nix-community/nix-melt/actions/workflows/ci.yml/badge.svg)](https://github.com/nix-community/nix-melt/actions/workflows/ci.yml) Stars:`202`.
* [Tabiew](https://github.com/shshemi/tabiew) - A lightweight TUI app to view and query CSV files. Stars:`198`.
* [raftario/licensor](https://github.com/raftario/licensor) - write licenses to stdout [![GitHub Actions](https://github.com/raftario/licensor/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/raftario/licensor/actions/workflows/build.yml) Stars:`193`.
* [rust-parallel](https://github.com/aaronriekenberg/rust-parallel) - Fast command line app using Tokio to execute commands in parallel.  Similar interface to GNU Parallel or xargs. [![Crate](https://img.shields.io/crates/v/rust-parallel.svg?logo=rust)](https://crates.io/crates/rust-parallel) [![Build Status](https://github.com/aaronriekenberg/rust-parallel/actions/workflows/CI.yml/badge.svg)](https://github.com/aaronriekenberg/rust-parallel/actions/workflows/CI.yml) Stars:`126`.
* [sorairolake/qrtool](https://github.com/sorairolake/qrtool) [[qrtool](https://crates.io/crates/qrtool)] - A utility for encoding and decoding QR code images. [![CI](https://github.com/sorairolake/qrtool/workflows/CI/badge.svg?branch=develop)](https://github.com/sorairolake/qrtool/actions?query=workflow%3ACI) Stars:`121`.
* [brycx/checkpwn](https://github.com/brycx/checkpwn) - A Have I Been Pwned (HIBP) command-line utility tool that lets you easily check for compromised accounts and passwords. Stars:`119`.
* [wrestic](https://github.com/alvaro17f/wrestic) - A wrapper around restic. Stars:`76`.

### Video

* [gyroflow/gyroflow](https://github.com/gyroflow/gyroflow) - Video stabilization application using gyroscope data Stars:`6.4K`.
* [xiph/rav1e](https://github.com/xiph/rav1e) - The fastest and safest AV1 encoder. Stars:`3.6K`.
* [harlanc/xiu](https://github.com/harlanc/xiu) - A powerful and secure live server (rtmp/httpflv/hls/relay). [![crates.io](https://img.shields.io/crates/v/xiu.svg)](https://crates.io/crates/xiu) Stars:`1.6K`.
* [dertuxmalwieder/yaydl](https://github.com/dertuxmalwieder/yaydl) [[yaydl](https://crates.io/crates/yaydl)] - A simple video downloader Stars:`266`.
* [vidmerger](https://github.com/TGotwig/vidmerger) - Merge video & audio files via CLI Stars:`113`.

### Virtualization

* [firecracker-microvm/firecracker](https://github.com/firecracker-microvm/firecracker) - A lightweight virtual machine for container workload [Firecracker Microvm](https://firecracker-microvm.github.io/) Stars:`24.5K`.
* [containers/youki](https://github.com/containers/youki) - A container runtime [![build badge](https://github.com/containers/youki/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/containers/youki/actions) Stars:`6.0K`.
* [kata-containers/kata-containers](https://github.com/kata-containers/kata-containers) - A implementation of lightweight Virtual Machines (VMs) that feel and perform like containers, but provide the workload isolation and security advantages of VMs. Stars:`5.1K`.
* [tailhook/vagga](https://github.com/tailhook/vagga) - A containerization tool without daemons Stars:`1.9K`.

### Web

* [LemmyNet/lemmy](https://github.com/LemmyNet/lemmy) - A link aggregator / reddit clone for the fediverse [![Build Status](https://cloud.drone.io/api/badges/LemmyNet/lemmy/status.svg)](https://cloud.drone.io/LemmyNet/lemmy) Stars:`13.0K`.
* [libreddit](https://github.com/libreddit/libreddit) - An alternative private front-end to Reddit Stars:`5.0K`.
* [Plume-org/Plume](https://github.com/Plume-org/Plume) - ActivityPub federating blogging application Stars:`2.1K`.
* [Revolt/backend](https://github.com/revoltchat/backend) - User-first chat platform built with modern web technologies. Stars:`1.1K`.
* [cfal/tobaru](https://github.com/cfal/tobaru) - Port forwarder with allowlists, IP and TLS SNI/ALPN rule-based routing, iptables support, round-robin forwarding (load balancing), and hot reloading. Stars:`173`.
* [MASQ-Project/Node](https://github.com/MASQ-Project/Node) - MASQ Node software provides a decentralized mesh-network of nodes for global users to access normal internet content - next evolution of tech beyond Tor & VPN [![build badge](https://github.com/MASQ-Project/Node/actions/workflows/ci-matrix.yml/badge.svg)](https://github.com/MASQ-Project/Node/actions) Stars:`169`.
* [importantimport/hatsu](https://github.com/importantimport/hatsu) - 🩵 Self-hosted and fully-automated ActivityPub bridge for static sites. [![release](https://github.com/importantimport/hatsu/actions/workflows/release.yml/badge.svg)](https://github.com/importantimport/hatsu/actions/workflows/release.yml) Stars:`110`.

### Web Servers

* [cloudflare/pingora](https://github.com/cloudflare/pingora) - A library for building fast, reliable and evolvable network services. Stars:`20.3K`.
* [svenstaro/miniserve](https://github.com/svenstaro/miniserve) - A small, self-contained cross-platform CLI tool that allows you to just grab the binary and serve some file(s) via HTTP [![build badge](https://github.com/svenstaro/miniserve/workflows/CI/badge.svg?branch=master)](https://github.com/svenstaro/miniserve/actions) Stars:`5.8K`.
* [TheWaWaR/simple-http-server](https://github.com/TheWaWaR/simple-http-server) - simple static http server Stars:`2.5K`.
* [static-web-server](https://github.com/static-web-server/static-web-server) - A blazing fast and asynchronous web server for static files-serving. ⚡ [![CI](https://github.com/static-web-server/static-web-server/actions/workflows/devel.yml/badge.svg)](https://github.com/static-web-server/static-web-server/actions/workflows/devel.yml?query=branch%3Amaster) Stars:`1.3K`.
* [mufeedvh/binserve](https://github.com/mufeedvh/binserve) - A blazingly fast static web server with routing, templating, and security in a single binary you can set up with zero code [![build badge](https://github.com/mufeedvh/binserve/workflows/CICD/badge.svg?branch=master)](https://github.com/mufeedvh/binserve/actions) Stars:`973`.
* [orhun/rustypaste](https://github.com/orhun/rustypaste) - A minimal file upload/pastebin service ![https://github.com/orhun/rustypaste/actions](https://img.shields.io/github/actions/workflow/status/orhun/rustypaste/ci.yml?branch=master&label=build) Stars:`722`.
* [thecoshman/http](https://github.com/thecoshman/http) - Host These Things Please - A basic http server for hosting a folder fast and simply Stars:`438`.
* [emanuele-em/proxelar](https://github.com/emanuele-em/proxelar) - A MITM Proxy 🦀! Toolkit for HTTP/1, HTTP/2, and WebSockets with SSL/TLS Capabilities [![Rust](https://github.com/emanuele-em/proxelar/actions/workflows/rust.yml/badge.svg)](https://github.com/emanuele-em/proxelar/actions/workflows/rust.yml) Stars:`373`.
* [wyhaya/see](https://github.com/wyhaya/see) - Static HTTP file server Stars:`204`.
* [ronanyeah/rust-hasura](https://github.com/ronanyeah/rust-hasura) - A demonstration of how a GraphQL server can be used as a remote schema with [Hasura](https://hasura.io/) ![Rust](https://github.com/ronanyeah/rust-hasura/workflows/Rust/badge.svg?branch=master) Stars:`140`.
* [mu-arch/skyfolder](https://github.com/mu-arch/skyfolder) - 🪂 Beautiful HTTP/Bittorrent server without the hassle. Secure - GUI - Pretty - Fast Stars:`116`.

## Development tools

* [delta](https://crates.io/crates/git-delta) - A syntax-highlighter for git and diff output[![build badge](https://github.com/dandavison/delta/workflows/Continuous%20Integration/badge.svg)](https://github.com/dandavison/delta//actions) Stars:`21.4K`.
* [just](https://github.com/casey/just) - A handy command runner for project-specific tasks Stars:`18.5K`.
* [git-cliff](https://github.com/orhun/git-cliff) - A highly customizable Changelog Generator that follows Conventional Commit specifications ![https://github.com/orhun/git-cliff/actions](https://img.shields.io/github/actions/workflow/status/orhun/git-cliff/ci.yml?branch=main&label=build) Stars:`8.3K`.
* [Rustup](https://github.com/rust-lang/rustup) - the Rust toolchain installer [![build badge](https://github.com/rust-lang/rustup/workflows/Linux%20(master)/badge.svg?branch=master)](https://github.com/rust-lang/rustup/actions) Stars:`6.0K`.
* [Racer](https://github.com/racer-rust/racer) - code completion for Rust Stars:`3.4K`.
* [typos](https://github.com/crate-ci/typos) [[typos-cli](https://crates.io/crates/typos-cli)] - Source code spell checker Stars:`2.3K`.
* [dotenv-linter](https://github.com/dotenv-linter/dotenv-linter) - Linter for `.env` files [![build badge](https://github.com/dotenv-linter/dotenv-linter/workflows/CI/badge.svg?branch=master)](https://github.com/dotenv-linter/dotenv-linter/actions?query=workflow%3ACI+branch%3Amaster) Stars:`1.8K`.
* [bacon](https://github.com/Canop/bacon) - background rust code checker, similar to cargo-watch Stars:`1.6K`.
* [create-rust-app](https://github.com/Wulf/create-rust-app) - Set up a modern rust+react web app by running one command. [![crate](https://img.shields.io/crates/v/create-rust-app.svg)](https://crates.io/crates/create-rust-app) Stars:`1.5K`.
* [geiger](https://github.com/geiger-rs/cargo-geiger) - A program that list statistics related to usage of unsafe code in a crate and all its dependencies [![Build Status](https://dev.azure.com/cargo-geiger/cargo-geiger/_apis/build/status/geiger-rs.cargo-geiger?branchName=master)](https://dev.azure.com/cargo-geiger/cargo-geiger/_build/latest?definitionId=1&branchName=master) Stars:`1.4K`.
* [ATAC](https://github.com/Julien-cpsn/ATAC) - A feature-full TUI API client made in Rust. ATAC is free, open-source, offline and account-less. Stars:`1.2K`.
* [Rust Search Extension](https://github.com/huhu/rust-search-extension) - A handy browser extension to search crates and docs in address bar (omnibox). [![Build Status](https://github.com/huhu/rust-search-extension/workflows/build/badge.svg?branch=master)](https://github.com/huhu/rust-search-extension/actions) Stars:`1.2K`.
* [mask](https://github.com/jacobdeichert/mask) - A CLI task runner defined by a simple markdown file [![build badge](https://github.com/jacobdeichert/mask/workflows/CI/badge.svg?branch=master)](https://github.com/jacobdeichert/mask/actions?query=workflow%3ACI) Stars:`1.0K`.
* [scriptisto](https://github.com/igor-petruk/scriptisto) - A language-agnostic "shebang interpreter" that enables you to write one file scripts in compiled languages. [![Build Status](https://cloud.drone.io/api/badges/igor-petruk/scriptisto/status.svg)](https://cloud.drone.io/igor-petruk/scriptisto) Stars:`904`.
* [clog-tool/clog-cli](https://github.com/clog-tool/clog-cli) - generates a changelog from git metadata ([conventional changelog](https://blog.thoughtram.io/announcements/tools/2014/09/18/announcing-clog-a-conventional-changelog-generator-for-the-rest-of-us.html)) Stars:`839`.
* [envio-cli/envio](https://github.com/envio-cli/envio) - A Modern And Secure CLI Tool For Managing Environment Variables [![build badge](https://github.com/envio-cli/envio/actions/workflows/CICD.yml/badge.svg?branch=main)](https://github.com/envio-cli/envio/actions/workflows/CICD.yml) Stars:`761`.
* [git-journal](https://github.com/saschagrunert/git-journal/) - The Git Commit Message and Changelog Generation Framework Stars:`608`.
* [hot-lib-reloader](https://github.com/rksm/hot-lib-reloader-rs) - Hot reload Rust code [![build badge](https://github.com/rksm/hot-lib-reloader-rs/actions/workflows/ci.yml/badge.svg)](https://github.com/rksm/hot-lib-reloader-rs/actions/workflows/ci.yml) Stars:`575`.
* [fw](https://github.com/brocode/fw) - workspace productivity booster [![Rust](https://github.com/brocode/fw/actions/workflows/rust.yml/badge.svg)](https://github.com/brocode/fw/actions/workflows/rust.yml) Stars:`530`.
* [datanymizer/datanymizer](https://github.com/datanymizer/datanymizer) - Powerful database anonymizer with flexible rules [![build badge](https://github.com/datanymizer/datanymizer/workflows/CI/badge.svg?branch=main)](https://github.com/datanymizer/datanymizer/actions?query=workflow%3ACI+branch%3Amain) Stars:`500`.
* [comtrya](https://github.com/comtrya/comtrya) - A configuration management tool for localhost / dotfiles [![build badge](https://github.com/comtrya/comtrya/actions/workflows/main.yaml/badge.svg)](https://github.com/comtrya/comtrya/actions) Stars:`470`.
* [dan-t/rusty-tags](https://github.com/dan-t/rusty-tags) - create ctags/etags for a cargo project and all of its dependencies Stars:`403`.
* [Module Linker](https://github.com/fiatjaf/module-linker) - Extension that adds `<a>` links to references in `mod`, `use` and `extern crate` statements at GitHub. Stars:`250`.
* [intelli-shell](https://github.com/lasantosr/intelli-shell) - Bookmark commands with placeholders and search or autocomplete at any time [![crate](https://img.shields.io/crates/v/intelli-shell.svg)](https://crates.io/crates/intelli-shell) [![build badge](https://github.com/lasantosr/intelli-shell/actions/workflows/release.yml/badge.svg)](https://github.com/lasantosr/intelli-shell/actions/workflows/release.yml) Stars:`206`.
* [frolic](https://github.com/FrolicOrg/Frolic) - An API layer to build customer facing dashboards 10x faster Stars:`180`.
* [ptags](https://github.com/dalance/ptags) - A parallel universal-ctags wrapper for git repository Stars:`124`.
* [fzf-make](https://github.com/kyu08/fzf-make) [[fzf-make](https://crates.io/crates/fzf-make)] - A command line tool that executes make target using fuzzy finder with preview window. [![crates.io](https://img.shields.io/crates/v/fzf-make?style=flatflat-square)](https://crates.io/crates/fzf-make) Stars:`63`.
* [clippy](https://crates.io/crates/clippy) - Rust lints

### Build system

  * [dtolnay/cargo-expand](https://github.com/dtolnay/cargo-expand) - Expand macros in your source code Stars:`2.5K`.
  * [cargo-make](https://crates.io/crates/cargo-make) - Task runner and build tool. [![build badge](https://github.com/sagiegurari/cargo-make/workflows/CI/badge.svg?branch=master)](https://github.com/sagiegurari/cargo-make/actions) Stars:`2.5K`.
  * [cargo-generate](https://github.com/cargo-generate/cargo-generate) - A generator of a rust project by leveraging a pre-existing git repository as a template. Stars:`1.9K`.
  * [cargo-udeps](https://github.com/est31/cargo-udeps) [[cargo-udeps](https://crates.io/crates/cargo-udeps)] - find unused dependencies Stars:`1.6K`.
  * [cargo-release](https://crates.io/crates/cargo-release) - tool for releasing git-managed cargo project, build, tag, publish, doc and push [![Rust](https://github.com/crate-ci/cargo-release/actions/workflows/ci.yml/badge.svg)](https://github.com/crate-ci/cargo-release/actions/workflows/rust.yml) Stars:`1.3K`.
  * [cargo-cache](https://crates.io/crates/cargo-cache) - inspect/manage/clean your cargo cache (`~/.cargo/`/`${CARGO_HOME}`), print sizes etc [![Build Status](https://github.com/matthiaskrgr/cargo-cache/workflows/ci/badge.svg?branch=master)](https://github.com/matthiaskrgr/cargo-cache/actions) Stars:`830`.
  * [cargo-all-features](https://github.com/frewsxcv/cargo-all-features) - A configurable subcommand to simplify testing, building and much more for all combinations of features [![CI](https://github.com/frewsxcv/cargo-all-features/actions/workflows/ci.yml/badge.svg)](https://github.com/frewsxcv/cargo-all-features/actions/workflows/ci.yml) Stars:`134`.
  * [cargo-rdme](https://github.com/orium/cargo-rdme) [[cargo-rdme](https://crates.io/crates/cargo-rdme)] - Cargo subcommand to create your README from your crate’s documentation. [![build badge](https://github.com/orium/cargo-rdme/workflows/CI/badge.svg)](https://github.com/orium/cargo-rdme/actions?query=workflow%3ACI) Stars:`125`.
  * [cargo-commander](https://crates.io/crates/cargo-commander) - A subcommand for `cargo` to run CLI commands similar to how the scripts section in `package.json` works [![Build and test](https://github.com/simonhyll/cargo-commander/actions/workflows/build.yml/badge.svg)](https://github.com/simonhyll/cargo-commander/actions/workflows/build.yml) Stars:`16`.
  * [cargo-graph](https://crates.io/crates/cargo-graph) - updated fork of `cargo-dot` with additional features. Unmaintained, see `cargo-deps`
  * [cargo-deb](https://crates.io/crates/cargo-deb) - Generates binary Debian packages
  * [cargo-ebuild](https://crates.io/crates/cargo-ebuild) - cargo extension that can generate ebuilds using the in-tree eclasses
  * [cargo-edit](https://crates.io/crates/cargo-edit) - allows you to add and list dependencies by reading/writing to your Cargo.toml file from the command line
  * [cargo-deps](https://crates.io/crates/cargo-deps) - build dependency graphs
* [Cargo](https://crates.io/) - the Rust package manager
  * [cargo-info](https://crates.io/crates/cargo-info) - queries crates.io for crates details from command line
  * [cargo-license](https://crates.io/crates/cargo-license) - A cargo subcommand to quickly view the licenses of all dependencies.
  * [cargo-limit](https://crates.io/crates/cargo-limit) - Cargo with less noise: warnings are skipped until errors are fixed, Neovim integration, etc. [![build badge](https://github.com/cargo-limit//cargo-limit/actions/workflows/rust.yml/badge.svg)](https://github.com/cargo-limit//cargo-limit/actions)
  * [cargo-do](https://crates.io/crates/cargo-do) - run multiple cargo commands in a row
  * [cargo-modules](https://crates.io/crates/cargo-modules) - A cargo plugin for showing a tree-like overview of a crate's modules.
  * [cargo-multi](https://crates.io/crates/cargo-multi) - runs specified cargo command on multiple crates
  * [cargo-outdated](https://crates.io/crates/cargo-outdated) - displays when newer versions of Rust dependencies are available, or out of date
  * [cargo-count](https://crates.io/crates/cargo-count) - lists source code counts and details about cargo projects, including unsafe statistics
  * [cargo-check](https://crates.io/crates/cargo-check) - A wrapper around `cargo rustc -- -Zno-trans` which can be helpful for running a faster compile if you only need correctness checks
  * [cargo-script](https://crates.io/crates/cargo-script) - lets people quickly and easily run Rust "scripts" which can make use of Cargo's package ecosystem
  * [cargo-bitbake](https://crates.io/crates/cargo-bitbake) - A cargo extension that can generate BitBake recipes utilizing the classes from meta-rust
  * [cargo-update](https://crates.io/crates/cargo-update) - cargo subcommand for checking and applying updates to installed executables
  * [cargo-watch](https://crates.io/crates/cargo-watch) - utility for cargo to compile projects when sources change
  * [cargo-benchcmp](https://crates.io/crates/cargo-benchcmp) - A utility to compare micro-benchmarks
* CMake
* [Fleet](https://github.com/dimensionhq/fleet) [[fleet-rs](https://crates.io/crates/fleet-rs)] - The blazing fast build tool for Rust. Stars:`2.4K`.
  * [Devolutions/CMakeRust](https://github.com/Devolutions/CMakeRust) - useful for integrating a Rust library into a CMake project Stars:`168`.
  * [SiegeLord/RustCMake](https://github.com/SiegeLord/RustCMake) - an example project showing usage of CMake with Rust Stars:`109`.
* GitHub actions
  * [nix-community/fenix](https://github.com/nix-community/fenix) - Rust toolchains and rust analyzer nightly for nix [![build-badge](https://github.com/nix-community/fenix/actions/workflows/ci.yml/badge.svg)](https://github.com/nix-community/fenix/actions/workflows/ci.yml) Stars:`611`.
  * [peaceiris/actions-mdbook](https://github.com/peaceiris/actions-mdbook) - GitHub Actions for mdBook Stars:`284`.
  * [icepuma/rust-action](https://github.com/icepuma/rust-action) - rust github action Stars:`78`.
* [Nix](https://nixos.org/)

### Debugging

* GDB
  * [gdbgui](https://github.com/cs01/gdbgui) - Browser based frontend for gdb to debug C, C++, Rust, and go. Stars:`9.8K`.
* [kxxt/tracexec](https://github.com/kxxt/tracexec) [[tracexec](https://crates.io/crates/tracexec)] - Tracer for execve{,at} and pre-exec behavior, launcher for debuggers. Stars:`195`.
* LLDB
  * [CodeLLDB](https://marketplace.visualstudio.com/items?itemName=vadimcn.vscode-lldb) - A LLDB extension for [Visual Studio Code](https://code.visualstudio.com/).

### Deployment

* Docker
  * [LukeMathWalker/cargo-chef](https://github.com/LukeMathWalker/cargo-chef) - A tool and pre-built images for caching compiling remote dependencies between Docker builds. Stars:`1.6K`.
  * [emk/rust-musl-builder](https://github.com/emk/rust-musl-builder) - Docker images for compiling static Rust binaries using musl-libc and musl-gcc, with static versions of useful C libraries Stars:`1.5K`.
  * [rust-cross/rust-musl-cross](https://github.com/rust-cross/rust-musl-cross) - Docker images for compiling static Rust binaries using musl-cross [![Build](https://github.com/rust-cross/rust-musl-cross/workflows/Build/badge.svg)](https://github.com/rust-cross/rust-musl-cross/actions?query=workflow%3ABuild) Stars:`598`.
  * [rust-lang-nursery/docker-rust](https://github.com/rust-lang/docker-rust) - the official Rust Docker image Stars:`428`.
  * [kpcyrd/mini-docker-rust](https://github.com/kpcyrd/mini-docker-rust) - An example project for very small rust docker images Stars:`206`.
  * [Stavrospanakakis/is_ready](https://github.com/Stavrospanakakis/is_ready) - Wait for multiple services to become available ![Build](https://github.com/Stavrospanakakis/is_ready/actions/workflows/release.yml/badge.svg) Stars:`122`.
  * [liuchong/docker-rustup](https://github.com/liuchong/docker-rustup) - A multiple version (with musl tools) Rust Docker image Stars:`93`.
* Heroku
* [MarcoIeni/release-plz](https://github.com/MarcoIeni/release-plz) [[release-plz](https://crates.io/crates/release-plz)] - Release crates from CI, with changelog generation and semver check. [![build badge](https://github.com/MarcoIeni/release-plz/workflows/CI/badge.svg)](https://github.com/MarcoIeni/release-plz/actions) Stars:`721`.
  * [emk/heroku-buildpack-rust](https://github.com/emk/heroku-buildpack-rust) - A buildpack for Rust applications on Heroku Stars:`519`.

### Embedded

[Rust Embedded](https://rust-embedded.org/) focuses on improving the end-to-end experience of using Rust in resource-constrained environments and non-traditional platforms. See [awesome-embedded-rust](https://github.com/rust-embedded/awesome-embedded-rust) for a curated, and more extended list of embedded Rust resources.

* Arduino
  * [avr-rust/ruduino](https://github.com/avr-rust/ruduino) - Reusable components for the Arduino Uno. Stars:`695`.
* Cross compiling
  * [japaric/rust-cross](https://github.com/japaric/rust-cross) - everything you need to know about cross compiling Rust programs Stars:`2.5K`.
  * [japaric/xargo](https://github.com/japaric/xargo) - effortless cross compilation of Rust programs to custom bare-metal targets like ARM Cortex-M Stars:`1.1K`.
* Espressif
  * [esp-rs](https://github.com/esp-rs) - home to a number of community projects enabling the use of the Rust programming language on various SoCs and modules produced by Espressif Systems.
* Firmware
  * [oreboot/oreboot](https://github.com/oreboot/oreboot) - oreboot is a fork of coreboot, with C removed, written in Rust Stars:`1.5K`.
* nRF
  * [nrf-rs/nrf-hal](https://github.com/nrf-rs/nrf-hal) - A Rust HAL for the nRF family of devices Stars:`483`.

### FFI

See also [Foreign Function Interface](https://doc.rust-lang.org/book/first-edition/ffi.html), [The Rust FFI Omnibus](http://jakegoulding.com/rust-ffi-omnibus/) (a collection of examples of using code written in Rust from other languages) and [FFI examples written in Rust](https://github.com/alexcrichton/rust-ffi-examples).

* C
  * [mozilla/cbindgen](https://github.com/mozilla/cbindgen) - generates C header files from Rust source files. Used in Gecko for WebRender Stars:`2.3K`.
  * [Sean1708/rusty-cheddar](https://github.com/Sean1708/rusty-cheddar) - generates C header files from Rust source files Stars:`191`.
* C#
  * [csbindgen](https://github.com/Cysharp/csbindgen) - generates C# bindings for Rust source files Stars:`609`.
* C++
  * [dtolnay/cxx](https://github.com/dtolnay/cxx) - Safe interop between Rust and C++ [![build badge](https://img.shields.io/badge/github-dtolnay/cxx-8da0cb?style=for-the-badge&labelColor=555555&logo=github)](https://github.com/dtolnay/cxx) Stars:`5.7K`.
  * [rust-lang/rust-bindgen](https://github.com/rust-lang/rust-bindgen) - A Rust bindings generator Stars:`4.2K`.
  * [rust-cpp](https://crates.io/crates/cpp) - Embed C++ code directly in Rust. [![Build status](https://ci.appveyor.com/api/projects/status/uu76vmcrwnjqra0u/branch/master?svg=true)](https://ci.appveyor.com/project/mystor/rust-cpp/branch/master)
* Erlang
  * [rusterlium/rustler](https://github.com/rusterlium/rustler) - safe Rust bridge for creating Erlang NIF functions Stars:`4.2K`.
* Java
  * [drrb/java-rust-example](https://github.com/drrb/java-rust-example) - use Rust from Java Stars:`336`.
  * [bennettanderson/rjni](https://github.com/benanders/rjni) - use Java from Rust Stars:`72`.
  * [j4rs](https://crates.io/crates/j4rs) - use Java from Rust
  * [jni](https://crates.io/crates/jni) - use Rust from Java
  * [jni-sys](https://crates.io/crates/jni-sys) - Rust definitions corresponding to jni.h
  * [rucaja](https://crates.io/crates/rucaja) - use Java from Rust
* Lua
  * [mlua-rs/mlua](https://github.com/mlua-rs/mlua) - High level Lua 5.4/5.3/5.2/5.1 (including LuaJIT) and Roblox Luau bindings to Rust with async/await support [![build badge](https://github.com/mlua-rs/mlua/workflows/CI/badge.svg)](https://github.com/mlua-rs/mlua/actions) Stars:`1.5K`.
  * [tomaka/hlua](https://github.com/tomaka/hlua) - Rust library to interface with Lua Stars:`503`.
  * [jcmoyer/rust-lua53](https://github.com/jcmoyer/rust-lua53) - Lua 5.3 bindings for Rust Stars:`158`.
  * [lilyball/rust-lua](https://github.com/lilyball/rust-lua) - Safe Rust bindings to Lua 5.1 Stars:`126`.
  * [tickbh/td_rlua](https://github.com/tickbh/td_rlua) [[td_rlua](https://crates.io/crates/td_rlua)] - Zero-cost high-level lua 5.3 wrapper for Rust Stars:`53`.
* mruby
  * [anima-engine/mrusty](https://github.com/anima-engine/mrusty) - mruby safe bindings for Rust Stars:`205`.
* Node.js
  * [neon-bindings/neon](https://github.com/neon-bindings/neon) - Rust bindings for writing safe and fast native Node.js modules Stars:`7.9K`.
  * [infinyon/node-bindgen](https://github.com/infinyon/node-bindgen) - Easy way to generate nodejs module using Rust Stars:`491`.
  * [zhangyuang/node-ffi-rs](https://github.com/zhangyuang/node-ffi-rs) - A module written in Rust and N-API provides interface (FFI) features for Node.js Stars:`118`.
* Objective-C
  * [SSheldon/rust-objc](https://github.com/SSheldon/rust-objc) - Objective-C Runtime bindings and wrapper for Rust Stars:`382`.
* PHP
  * [phper-framework/phper](https://github.com/phper-framework/phper) - The framework that allows us to write PHP extensions using pure and safe Rust whenever possible Stars:`264`.
* Prolog
  * [mthom/scryer-prolog](https://github.com/mthom/scryer-prolog/) - Scryer Prolog is a free software ISO Prolog system written in Rust Stars:`1.9K`.
* Python
  * [RustPython](https://github.com/RustPython/RustPython) - A Python Interpreter written in Rust [![Build Status](https://github.com/RustPython/RustPython/workflows/CI/badge.svg)](https://github.com/RustPython/RustPython/actions?query=workflow%3ACI) Stars:`18.0K`.
  * [PyO3/PyO3](https://github.com/PyO3/PyO3) - Rust bindings for the Python interpreter Stars:`11.5K`.
  * [dgrunwald/rust-cpython](https://github.com/dgrunwald/rust-cpython) - Python bindings Stars:`1.8K`.
  * [getsentry/milksnake](https://github.com/getsentry/milksnake) - extension for python setuptools that allows you to distribute dynamic linked libraries in Python wheels in the most portable way imaginable. Stars:`789`.
* Ruby
  * [danielpclark/rutie](https://github.com/danielpclark/rutie) - native Ruby extensions written in Rust and vice versa Stars:`930`.
  * [d-unsed/ruru](https://github.com/d-unsed/ruru) - native Ruby extensions written in Rust Stars:`830`.
* Web Assembly
  * [rustwasm/wasm-bindgen](https://github.com/rustwasm/wasm-bindgen) - A project for facilitating high-level interactions between wasm modules and JS. Stars:`7.5K`.
  * [rustwasm/wasm-pack](https://github.com/rustwasm/wasm-pack) - :package: :sparkles: pack up the wasm and publish it to npm! Stars:`6.1K`.
  * [rhysd/wain](https://github.com/rhysd/wain) - wain: WebAssembly INterpreter from scratch in Safe Rust with zero dependency [![build badge](https://github.com/rhysd/wain/workflows/CI/badge.svg?branch=master&event=push)](https://github.com/rhysd/wain/actions?query=workflow%3ACI+branch%3Amaster+event%3Apush) Stars:`418`.

### Formatters

* [rustfmt](https://github.com/rust-lang/rustfmt) - Rust code formatter maintained by the Rust team and included in cargo Stars:`5.9K`.
* [dprint](https://github.com/dprint/dprint) - A pluggable and configurable code formatting platform [![build badge](https://github.com/dprint/dprint/workflows/CI/badge.svg)](https://github.com/dprint/dprint/actions?query=workflow%3ACI) Stars:`3.1K`.
* [Prettier Rust](https://github.com/jinxdash/prettier-plugin-rust) - An opinionated Rust code formatter that autofixes bad syntax ([Prettier](https://prettier.io/) community plugin) Stars:`173`.

### IDEs

See also [Are we (I)DE yet?](https://areweideyet.com/) and [Rust Tools](https://www.rust-lang.org/tools).

  * [lapce](https://github.com/lapce/lapce) - Lightning-fast and Powerful Code Editor written in Rust. [![build badge](https://github.com/lapce/lapce/actions/workflows/release.yml/badge.svg)](https://github.com/lapce/lapce/actions/workflows/release.yml) Stars:`33.8K`.
    * [intellij-rust/intellij-rust](https://github.com/intellij-rust/intellij-rust) - Rust plugin for the IntelliJ Platform Stars:`4.5K`.
    * [rust.vim](https://github.com/rust-lang/rust.vim) - provides file detection, syntax highlighting, formatting, Syntastic integration, and more. Stars:`3.8K`.
    * [autozimu/LanguageClient-neovim](https://github.com/autozimu/LanguageClient-neovim) - [LSP](https://microsoft.github.io/language-server-protocol/) client. Implemented in Rust and supports rls out of the box. Stars:`3.5K`.
    * [rust-mode](https://github.com/rust-lang/rust-mode) - Rust Major Mode Stars:`1.1K`.
    * [crates.nvim](https://github.com/Saecki/crates.nvim) - plugin that helps to managing crates.io dependencies. Stars:`801`.
    * [rust-lang/rust-enhanced](https://github.com/rust-lang/rust-enhanced) - official Rust package Stars:`774`.
    * [rustic](https://github.com/brotzeit/rustic) - Rust development environment for Emacs [![build badge](https://github.com/brotzeit/rustic/workflows/CI/badge.svg)](https://github.com/brotzeit/rustic/actions?query=workflow%3ACI) Stars:`712`.
    * [vim-racer](https://github.com/racer-rust/vim-racer) - allows vim to use [Racer](https://github.com/racer-rust/racer) for Rust code completion and navigation. Stars:`627`.
    * [kakoune-lsp](https://github.com/kakoune-lsp/kakoune-lsp/) - [LSP](https://microsoft.github.io/language-server-protocol/) client. Implemented in Rust and supports rls out of the box. Stars:`591`.
    * [emacs-racer](https://github.com/racer-rust/emacs-racer) - Autocompletion (see also [company](https://company-mode.github.io) and [auto-complete](https://github.com/auto-complete/auto-complete)) Stars:`399`.
    * [Eclipse Corrosion](https://github.com/eclipse-corrosion/corrosion) - a Rust development plugin for the Eclipse IDE, providing a rich edition experience through integration with the Rust Analyzer language server, Cargo runner and gdb debugger Stars:`218`.
  * [Ride](https://github.com/madeso/ride) - A Rust IDE Stars:`171`.
    * [flycheck-rust](https://github.com/flycheck/flycheck-rust) - Rust support for [Flycheck](https://github.com/flycheck/flycheck) Stars:`119`.
  * [IntelliJ](https://www.jetbrains.com/idea/)
  * [Sublime Text](https://www.sublimetext.com/)
  * [Kakoune](http://kakoune.org/)
  * [Vim](https://vim.sourceforge.io/) - the ubiquitous text editor
  * [Emacs](https://www.gnu.org/software/emacs/)
  * [Eclipse](https://www.eclipse.org/)
  * [gnome-builder](https://wiki.gnome.org/Apps/Builder) - native support for rust and cargo since Version 3.22.2
  * [gitpod.io](https://gitpod.io) - Online IDE with full Rust support based on Rust Language Server
  * Visual Studio
    * [PistonDevelopers/VisualRust](https://github.com/PistonDevelopers/VisualRust) - A Visual Studio extension for Rust [![Build status](https://ci.appveyor.com/api/projects/status/5nw5no10jj0y4p3f?svg=true)](https://ci.appveyor.com/project/vosen/visualrust) Stars:`701`.
    * [crates](https://github.com/serayuzgur/crates) - crates is an extension for crates.io dependencies. [![build badge](https://img.shields.io/vscode-marketplace/v/serayuzgur.crates.svg)](https://github.com/serayuzgur/crates) Stars:`237`.
    * [dgriffen/rls-vs2017](https://github.com/ZoeyR/rls-vs2017) - Rust support for Visual Studio 2017 Preview [![build badge](https://ci.appveyor.com/api/projects/status/d2lxlincwninhsng?svg=true)](https://ci.appveyor.com/project/dgriffen/rls-vs2017) Stars:`110`.
  * [Visual Studio Code](https://code.visualstudio.com/)
    * [Better TOML](https://marketplace.visualstudio.com/items?itemName=bungcip.better-toml) - TOML support in vscode
    * [CodeLLDB](https://marketplace.visualstudio.com/items?itemName=vadimcn.vscode-lldb) - A LLDB extension
    * [Prettier - Code formatter (Rust)](https://marketplace.visualstudio.com/items?itemName=jinxdash.prettier-rust) - Opinionated Rust code formatter that autofixes bad syntax ([Prettier](https://prettier.io/) community plugin)
    * [rust-analyzer](https://marketplace.visualstudio.com/items?itemName=rust-lang.rust-analyzer) - An alternative rust language server to the RLS

### Profiling

* [Bytehound](https://github.com/koute/bytehound) - A memory profiler for Linux Stars:`4.4K`.
* [bheisler/criterion.rs](https://github.com/bheisler/criterion.rs) - Statistics-driven benchmarking library Stars:`4.3K`.
* [Divan](https://github.com/nvzqz/divan) - Simple yet powerful benchmarking library with allocation profiling Stars:`849`.
* [Bencher](https://github.com/bencherdev/bencher) - A suite of continuous benchmarking tools designed to catch performance regressions in CI Stars:`457`.
* [ellisonch/rust-stopwatch](https://github.com/ellisonch/rust-stopwatch) - A stopwatch library Stars:`80`.
* FlameGraphs
* [sharkdp/hyperfine](https://github.com/sharkdp/hyperfine) - A command-line benchmarking tool Stars:`20.8K`.
  * [llogiq/flame](https://github.com/llogiq/flame) - An intrusive flamegraph profiling tool for rust Stars:`689`.
  * [mrhooray/torch](https://github.com/mrhooray/torch) - generates FlameGraphs based on DWARF Debug Info Stars:`131`.

### Services

* [deps.rs](https://github.com/deps-rs/deps.rs) - Detect outdated or insecure dependencies Stars:`421`.
* [docs.rs](https://docs.rs) - Automatic documentation generation of crates

### Static analysis

[[assert](https://crates.io/keywords/assert), [static](https://crates.io/keywords/static)]

* [facebookexperimental/MIRAI](https://github.com/facebookexperimental/mirai) - an abstract interpreter operating on Rust's mid-level intermediate representation (MIR) [![Continuous Integration](https://github.com/facebookexperimental/mirai/actions/workflows/rust.yml/badge.svg)](https://github.com/facebookexperimental/mirai/actions/workflows/rust.yml) Stars:`986`.
* [static_assertions](https://crates.io/crates/static_assertions) - Compile-time assertions to ensure that invariants are met

### Testing

[[test](https://crates.io/keywords/test), [testing](https://crates.io/keywords/testing)]

* Code Coverage
  * [tarpaulin](https://crates.io/crates/cargo-tarpaulin) - A code coverage tool
* Continuous Integration
  * [trust](https://github.com/japaric/trust) - A Travis CI and AppVeyor template to test your Rust crate on 5 architectures and publish binary releases of it for Linux, macOS and Windows Stars:`1.3K`.
* Frameworks and Runners
  * [rstest](https://crates.io/crates/rstest) - Fixture-based test framework [![Build Status](https://github.com/la10736/rstest/workflows/Test/badge.svg?branch=master)](https://github.com/la10736/rstest/actions) Stars:`1.0K`.
  * [cucumber](https://crates.io/crates/cucumber) [![Latest Version](https://img.shields.io/crates/v/cucumber.svg)](https://crates.io/crates/cucumber) - An implementation of the Cucumber testing framework for Rust. Fully native, no external test runners or dependencies. [![Build Status](https://github.com/cucumber-rs/cucumber/workflows/CI/badge.svg?branch=master)](https://github.com/cucumber-rs/cucumber) Stars:`551`.
  * [GoogleTest Rust](https://crates.io/crates/googletest) - Powerful test assertion framework based on the C++ test library GoogleTest [![Build Status](https://github.com/google/googletest-rust/workflows/CI/badge.svg)](https://github.com/google/googletest-rust/actions?query=workflow%3ACI+branch%3Amain) Stars:`185`.
  * [rlt](https://github.com/wfxr/rlt) - A universal load testing framework, with real-time tui support. Stars:`126`.
  * [d-e-s-o/test-log](https://github.com/d-e-s-o/test-log) [[test-log](https://crates.io/crates/test-log)] - A replacement of the `#[test]` attribute that initializes logging and/or tracing infrastructure before running tests. [![GitHub Workflow Status](https://github.com/d-e-s-o/test-log/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/d-e-s-o/test-log/actions/workflows/test.yml) Stars:`99`.
  * [AlKass/polish](https://github.com/AlKass/polish) - Mini Testing/Test-Driven Framework [![Crates Package Status](https://img.shields.io/crates/v/polish.svg)](https://crates.io/crates/polish) Stars:`52`.
  * [demonstrate](https://crates.io/crates/demonstrate) - Declarative Testing Framework [![Build Status](https://github.com/aubaugh/demonstrate/workflows/Continuous%20Integration/badge.svg?branch=master)](https://github.com/aubaugh/demonstrate) Stars:`49`.
  * [cargo-dinghy](https://crates.io/crates/cargo-dinghy/) - A cargo extension to simplify running library tests and benches on smartphones and other small processor devices.
  * [speculate](https://crates.io/crates/speculate) - An RSpec inspired minimal testing framework
* Mocking and Test Data
  * [asomers/mockall](https://github.com/asomers/mockall) [[mockall](https://crates.io/crates/mockall)] - A powerful mock object library. [![Cirrus Build Status](https://api.cirrus-ci.com/github/asomers/mockall.svg)](https://cirrus-ci.com/github/asomers/mockall) Stars:`1.4K`.
  * [synth](https://github.com/shuttle-hq/synth/) - Generate database data declaratively. [![build](https://github.com/shuttle-hq/synth/actions/workflows/synth-test.yml/badge.svg)](https://github.com/shuttle-hq/synth) Stars:`1.4K`.
  * [fake-rs](https://github.com/cksac/fake-rs) - A library for generating fake data Stars:`849`.
  * [httpmock](https://github.com/alexliesenfeld/httpmock) - HTTP mocking [![build badge](https://dev.azure.com/alexliesenfeld/httpmock/_apis/build/status/alexliesenfeld.httpmock?branchName=master)](https://dev.azure.com/alexliesenfeld/httpmock/_build/latest?definitionId=2&branchName=master) Stars:`436`.
  * [nrxus/faux](https://github.com/nrxus/faux/) [![Latest Version](https://img.shields.io/crates/v/faux.svg)](https://crates.io/crates/faux) - A library to create mocks out of structs. ![build](https://github.com/nrxus/faux/workflows/test/badge.svg?branch=master) Stars:`410`.
  * [goldenfile](https://github.com/calder/rust-goldenfile) [[goldenfile](https://crates.io/crates/goldenfile)] - A library providing a simple API for goldenfile testing. Stars:`36`.
  * [mockiato](https://crates.io/crates/mockiato) - A strict, yet friendly mocking library for unstable Rust 2018
  * [mockito](https://crates.io/crates/mockito) - HTTP mocking
* Mutation Testing
  * [mutagen](https://github.com/llogiq/mutagen) [[mutagen](https://crates.io/crates/mutagen)] - A source-level mutation testing framework (nightly only) Stars:`621`.
  * [cargo-mutants](https://github.com/sourcefrog/cargo-mutants) [[cargo-mutants](https://crates.io/crates/cargo-mutants)] - Finds inadequately tested code by injecting mutations, no source changes required. [![build badge](https://github.com/sourcefrog/cargo-mutants/actions/workflows/tests.yml/badge.svg?branch=main&event=push)](https://github.com/sourcefrog/cargo-mutants/actions/workflows/tests.yml?query=branch%3Amain) Stars:`473`.
* Property Testing and Fuzzing
  * [rust-fuzz/afl.rs](https://github.com/rust-fuzz/afl.rs) - A Rust fuzzer, using [AFL](https://lcamtuf.coredump.cx/afl/) Stars:`1.6K`.
  * [proptest](https://crates.io/crates/proptest) - property testing framework inspired by the [Hypothesis](https://hypothesis.works/) framework for Python
  * [quickcheck](https://crates.io/crates/quickcheck) - A Rust implementation of [QuickCheck](https://wiki.haskell.org/Introduction_to_QuickCheck1)

### Transpiling

* [immunant/c2rust](https://github.com/immunant/c2rust) - C to Rust translator and cross checker built atop Clang/LLVM. Stars:`3.8K`.
* [BayesWitnesses/m2cgen](https://github.com/BayesWitnesses/m2cgen) - A CLI tool to transpile trained classic machine learning models into a native Rust code with zero dependencies. [![GitHub Actions Status](https://github.com/BayesWitnesses/m2cgen/workflows/GitHub%20Actions/badge.svg?branch=master)](https://github.com/BayesWitnesses/m2cgen/actions) Stars:`2.7K`.
* [jameysharp/corrode](https://github.com/jameysharp/corrode) - A C to Rust translator written in Haskell. Stars:`2.1K`.

## Libraries

* [perf-monitor-rs](https://github.com/larksuite/perf-monitor-rs) - A toolkit designed to be a foundation for applications to monitor their performance. [![crates.io](https://img.shields.io/crates/v/perf_monitor.svg)](https://crates.io/crates/perf_monitor) Stars:`200`.

### Artificial Intelligence

#### Genetic algorithms

* [innoave/genevo](https://github.com/innoave/genevo) - Execute genetic algorithm (GA) simulations in a customizable and extensible way. Stars:`172`.
* [Martin1887/oxigen](https://github.com/Martin1887/oxigen) - Fast, parallel, extensible and adaptable genetic algorithm library. A example using this library solves the N Queens problem for N = 255 in only few seconds and using less than 1 MB of RAM. Stars:`163`.
* [pkalivas/radiate](https://github.com/pkalivas/radiate) - A customizable parallel genetic programming engine capable of evolving solutions for supervised, unsupervised, and reinforcement learning problems. Comes with complete and customizable implementation of NEAT and Evtree.![Crates.io](https://img.shields.io/crates/v/radiate) Stars:`146`.
* [willi-kappler/darwin-rs](https://github.com/willi-kappler/darwin-rs) - Evolutionary algorithms Stars:`114`.
* [m-decoster/RsGenetic](https://github.com/m-decoster/RsGenetic) - Genetic Algorithm library. In maintenance mode. Stars:`76`.

#### Machine learning

See [[Machine learning](https://crates.io/keywords/machine-learning)]

See also [About Rust’s Machine Learning Community](https://medium.com/@autumn_eng/about-rust-s-machine-learning-community-4cda5ec8a790#.hvkp56j3f) and [Are we learning yet?](https://www.arewelearningyet.com).

* [huggingface/candle](https://github.com/huggingface/candle) [[candle-core](https://crates.io/crates/candle-core)] - a minimalist ML framework with a focus on easiness of use and on performance (including GPU support) Stars:`14.4K`.
* [huggingface/tokenizers](https://github.com/huggingface/tokenizers) - Hugging Face's tokenizers for modern NLP pipelines (original implementation) with bindings for Python. [![Build Status](https://github.com/huggingface/tokenizers/workflows/Rust/badge.svg?branch=master)](https://github.com/huggingface/tokenizers/actions) Stars:`8.7K`.
* [burn](https://github.com/tracel-ai/burn) - A Flexible and Comprehensive Deep Learning Framework. Stars:`7.8K`.
* [autumnai/leaf](https://github.com/autumnai/leaf) - Open Machine Intelligence framework.. Abandoned project. The most updated fork is [juice](https://github.com/fff-rs/juice). Stars:`5.6K`.
* [tensorflow/rust](https://github.com/tensorflow/rust) - Bindings for TensorFlow. Stars:`5.1K`.
* [LaurentMazare/tch-rs](https://github.com/LaurentMazare/tch-rs) - Bindings for PyTorch. Stars:`4.0K`.
* [rust-ml/linfa](https://github.com/rust-ml/linfa) - Machine learning framework. Stars:`3.5K`.
* [guillaume-be/rust-bert](https://github.com/guillaume-be/rust-bert) [[rust_bert](https://crates.io/crates/rust_bert)] - Ready-to-use NLP pipelines and language models Stars:`2.5K`.
* [coreylowman/dfdx](https://github.com/coreylowman/dfdx) - CUDA accelerated machine learning framework that leverages many of Rust's unique features. ![Crates.io](https://img.shields.io/crates/v/dfdx) Stars:`1.7K`.
* [smartcorelib/smartcore](https://github.com/smartcorelib/smartcore) - Machine Learning Library [![Build Status](https://img.shields.io/circleci/build/github/smartcorelib/smartcore)](https://smartcorelib.org/) Stars:`671`.
* [maciejkula/rustlearn](https://github.com/maciejkula/rustlearn) - Machine learning library. [![Circle CI](https://circleci.com/gh/maciejkula/rustlearn.svg?style=svg)](https://app.circleci.com/pipelines/github/maciejkula/rustlearn) Stars:`613`.

#### OpenAI

* [64bit/async-openai](https://github.com/64bit/async-openai) [[async-openai](https://crates.io/crates/async-openai)] - Ergonomic Rust bindings for OpenAI API based on OpenAPI spec. Stars:`1.0K`.
* [zurawiki/tiktoken-rs](https://github.com/zurawiki/tiktoken-rs) [[tiktoken-rs](https://crates.io/crates/tiktoken-rs)] - Library for tokenizing text with OpenAI models using tiktoken. [![CI](https://github.com/zurawiki/tiktoken-rs/actions/workflows/ci.yml/badge.svg)](https://github.com/zurawiki/tiktoken-rs/actions/workflows/ci.yml) Stars:`221`.

### Astronomy

[[astronomy](https://crates.io/keywords/astronomy)]

* [saurvs/astro-rust](https://github.com/saurvs/astro-rust) - astronomy Stars:`258`.
* [cds-astro/aladin-lite](https://github.com/cds-astro/aladin-lite) - Web application for visualizing spatial and planetary image surveys in different projections Stars:`94`.
* [flosse/rust-sun](https://github.com/flosse/rust-sun) [[sun](https://crates.io/crates/sun)] - A rust port of the JS library suncalc Stars:`46`.
* [fitsio](https://crates.io/crates/fitsio) - fits interface library wrapping cfitsio

### Asynchronous

* [tokio-rs/tokio](https://github.com/tokio-rs/tokio) - A runtime for writing reliable, asynchronous, and slim applications with the Rust programming language. Stars:`25.4K`.
* [mio](https://github.com/tokio-rs/mio) - MIO is a lightweight IO library, with a focus on adding as little overhead as possible over the OS abstractions Stars:`6.2K`.
* [rust-lang/futures-rs](https://github.com/rust-lang/futures-rs) - Zero-cost futures Stars:`5.3K`.
* [async-std](https://async.rs/) [[async-std](https://crates.io/crates/async-std)] - Async version of the Rust standard library [![CI](https://github.com/async-rs/async-std/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/async-rs/async-std/actions/workflows/ci.yml) Stars:`3.9K`.
* [Xudong-Huang/may](https://github.com/Xudong-Huang/may) - Stackful coroutine library Stars:`1.9K`.
* [zonyitoo/coio-rs](https://github.com/zonyitoo/coio-rs) - A coroutine I/O library with a working-stealing scheduler Stars:`455`.
* [dpc/mioco](https://github.com/dpc/mioco) - Scalable, coroutine-based, asynchronous IO handling library Stars:`145`.
* [TeaEntityLab/fpRust](https://github.com/TeaEntityLab/fpRust) - Monad/MonadIO, Handler, Coroutine/doNotation, Functional Programming features for Rust Stars:`115`.
* [t3hmrman/async-dropper](https://github.com/t3hmrman/async-dropper) [[async-dropper](https://crates.io/crates/async-dropper)] - Implementation of `AsyncDrop` Stars:`29`.
* [igumnoff/gabriel2](https://github.com/igumnoff/gabriel2) [[gabriel2](https://crates.io/crates/gabriel2)] - Gabriel2: An actor-model library based on Tokio Stars:`16`.

### Audio and Music

[[audio](https://crates.io/keywords/audio)]

  * [RustAudio/cpal](https://github.com/RustAudio/cpal) - Low-level cross-platform audio I/O library. [![Actions Status](https://github.com/RustAudio/cpal/workflows/cpal/badge.svg?branch=master)](https://github.com/RustAudio/cpal/actions) Stars:`2.5K`.
* [pdeljanov/Symphonia](https://github.com/pdeljanov/Symphonia) - Audio decoding and media demuxing library supporting AAC, FLAC, MP3, MP4, OGG, Vorbis, and WAV. Stars:`2.2K`.
  * [RustAudio/rodio](https://github.com/RustAudio/rodio) - Audio playback library Stars:`1.7K`.
* [ozankasikci/rust-music-theory](https://github.com/ozankasikci/rust-music-theory) - Music theory library Stars:`621`.
  * [RustAudio/rust-portaudio](https://github.com/RustAudio/rust-portaudio) - PortAudio bindings Stars:`368`.
* [Serial-ATA/lofty-rs](https://github.com/Serial-ATA/lofty-rs) [[lofty](https://crates.io/crates/lofty)] - A library for reading and editing the metadata of various audio formats [![build badge](https://github.com/Serial-ATA/lofty-rs/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/Serial-ATA/lofty-rs/actions) Stars:`179`.
* [jhasse/ears](https://github.com/jhasse/ears) - A simple library to play Sounds and Musics, on top of OpenAL and libsndfile Stars:`89`.
* [musitdev/portmidi-rs](https://github.com/musitdev/portmidi-rs) - [PortMidi](https://portmedia.sourceforge.net/portmidi/) bindings Stars:`75`.
* [insomnimus/nodi](https://github.com/insomnimus/nodi) [[nodi](https://crates.io/crates/nodi)] - A library for playback and abstraction of MIDI files. [![build badge](https://github.com/insomnimus/nodi/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/insomnimus/nodi/actions) Stars:`17`.
* [hound](https://crates.io/crates/hound) - A WAV encoding and decoding library
* [RustAudio](https://github.com/RustAudio)

### Authentication

* [Keats/jsonwebtoken](https://github.com/Keats/jsonwebtoken) - [JSON Web Token](https://en.wikipedia.org/wiki/JSON_Web_Token) library Stars:`1.6K`.
* [oauth2](https://github.com/ramosbugs/oauth2-rs) - Extensible, strongly-typed OAuth2 client library Stars:`865`.
* [oxide-auth](https://github.com/HeroicKatora/oxide-auth) - A OAuth2 server library, for use in combination with actix or other frontends, featuring a set of configurable and pluggable backends [![Build Status](https://api.cirrus-ci.com/github/HeroicKatora/oxide-auth.svg?branch=master)](https://cirrus-ci.com/github/HeroicKatora/oxide-auth) Stars:`657`.
* [yup-oauth2](https://github.com/dermesser/yup-oauth2) - An oauth2 client implementation providing the Device, Installed and Service Account flows Stars:`213`.
* [constantoine/totp-rs](https://github.com/constantoine/totp-rs) [[totp-rs](https://crates.io/crates/totp-rs)] - 2fa library to generate and verify TOTP-based tokens ![Build Status](https://github.com/constantoine/totp-rs/workflows/Rust/badge.svg) Stars:`151`.
* [sgrust01/jwtvault](https://github.com/sgrust01/jwtvault) - Async library to manage and orchestrate JWT workflow Stars:`68`.

### Automotive

* [mbr/socketcan](https://github.com/socketcan-rs/socketcan-rs) [[socketcan](https://crates.io/crates/socketcan)] - Linux SocketCAN library Stars:`121`.
* [marcelbuesing/can-dbc](https://github.com/marcelbuesing/can-dbc) [[can-dbc](https://crates.io/crates/can-dbc)] - A parser for the DBC format Stars:`59`.
* [idletea/tokio-socketcan](https://github.com/idletea/tokio-socketcan) [[tokio-socketcan](https://crates.io/crates/tokio-socketcan)] - Linux SocketCAN support for tokio based on the socketcan crate Stars:`34`.
* [Sensirion/lin-bus](https://github.com/Sensirion/lin-bus-rs) [[lin-bus](https://crates.io/crates/lin-bus)] - LIN bus driver traits and protocol implementation [![build badge](https://circleci.com/gh/Sensirion/lin-bus-rs.svg?style=svg)](https://app.circleci.com/pipelines/github/Sensirion/lin-bus-rs) Stars:`16`.
* [marcelbuesing/tokio-socketcan-bcm](https://github.com/marcelbuesing/tokio-socketcan-bcm) [[tokio-socketcan-bcm](https://crates.io/crates/tokio-socketcan-bcm)] - Linux SocketCAN BCM support for tokio Stars:`7`.

### Bioinformatics

* [Rust-Bio](https://github.com/rust-bio) - bioinformatics libraries.

### Caching

* [mozilla/sccache](https://github.com/mozilla/sccache/) - Shared Compilation Cache, great compilation Stars:`5.6K`.
* [jaemk/cached](https://github.com/jaemk/cached) - Simple function caching/memoization Stars:`1.4K`.
* [moka-rs/moka](https://github.com/moka-rs/moka) - A high performance concurrent caching library inspired by the Caffeine library for Java [![build badge](https://github.com/moka-rs/moka/workflows/CI/badge.svg)](https://github.com/moka-rs/moka/actions/workflows/CI.yml) Stars:`1.4K`.
* [zkat/cacache-rs](https://github.com/zkat/cacache-rs) - A high-performance, concurrent, content-addressable disk cache, optimized for async APIs [![build badge](https://github.com/zkat/cacache-rs/workflows/CI/badge.svg)](https://github.com/zkat/cacache-rs/actions/workflows/ci.yml) Stars:`505`.
* [al8n/stretto](https://github.com/al8n/stretto) - A high performance thread-safe memory-bound cache [![build badge](https://github.com/al8n/stretto/actions/workflows/ci.yml/badge.svg)](https://github.com/al8n/stretto/actions/workflows/ci.yml) Stars:`407`.
* [aisk/rust-memcache](https://github.com/aisk/rust-memcache) - Memcached client library Stars:`128`.
* [06chaynes/http-cache](https://github.com/06chaynes/http-cache) [[http-cache](https://crates.io/crates/http-cache)] - A caching middleware that follows HTTP caching rules [![build badge](https://github.com/06chaynes/http-cache/workflows/http-cache/badge.svg)](https://github.com/06chaynes/http-cache/actions/workflows/http-cache.yml) Stars:`65`.

### Cloud

* AWS [[aws](https://crates.io/keywords/aws)]
  * [awslabs/aws-lambda-rust-runtime](https://github.com/awslabs/aws-lambda-rust-runtime) [[lambda_runtime](https://crates.io/crates/lambda_runtime)] - Runtime for AWS Lambda [![build badge](https://github.com/awslabs/aws-lambda-rust-runtime/workflows/Rust/badge.svg)](https://github.com/awslabs/aws-lambda-rust-runtime/actions) Stars:`3.3K`.
  * [awslabs/aws-sdk-rust](https://github.com/awslabs/aws-sdk-rust) - The new AWS SDK Stars:`2.9K`.
  * [rusoto/rusoto](https://github.com/rusoto/rusoto) - An AWS SDK for Rust Stars:`2.7K`.
* Load Balancer
  * [Convey](https://github.com/bparli/convey) - Layer 4 Load Balancer with dynamic configuration loading. Stars:`336`.
* Multi Cloud
  * [Qovery/engine](https://github.com/Qovery/engine) - Abstraction layer library that turns easy application deployment on Cloud providers in just a few minutes Stars:`2.2K`.

### Command-line

* Argument parsing
  * [clap-rs](https://github.com/clap-rs/clap) [[clap](https://crates.io/crates/clap)] - A simple to use, full featured command-line argument parser Stars:`13.7K`.
  * [TeXitoi/structopt](https://github.com/TeXitoi/structopt) [[structopt](https://crates.io/crates/structopt)] - parse command line argument by defining a struct Stars:`2.7K`.
  * [google/argh](https://github.com/google/argh) [[argh](https://crates.io/crates/argh)] - An opinionated Derive-based argument parser optimized for code size [![build badge](https://github.com/google/argh/workflows/Argh/badge.svg?branch=master)](https://github.com/google/argh/actions) Stars:`1.6K`.
  * [docopt/docopt.rs](https://github.com/docopt/docopt.rs) [[docopt](https://crates.io/crates/docopt)] - Implementation of [DocOpt](http://docopt.org) Stars:`756`.
  * [killercup/quicli](https://github.com/killercup/quicli) [[quicli](https://crates.io/crates/quicli)] - quickly build cool CLI apps Stars:`541`.
  * [ksk001100/seahorse](https://github.com/ksk001100/seahorse) [[seahorse](https://crates.io/crates/seahorse)] - A minimal CLI framework [![Build status](https://github.com/ksk001100/seahorse/workflows/CI/badge.svg?branch=master)](https://github.com/ksk001100/seahorse/actions) Stars:`280`.
  * [cliparser](https://crates.io/crates/cliparser) - Simple command line parser. [![build badge](https://github.com/sagiegurari/cliparser/workflows/CI/badge.svg?branch=master)](https://github.com/sagiegurari/cliparser/actions) Stars:`8`.
* Data visualization
  * [zhiburt/tabled](https://github.com/zhiburt/tabled) [[tabled](https://crates.io/crates/tabled)] - An easy to use library for pretty print tables of structs and enums. [![Build Status](https://github.com/zhiburt/tabled/actions/workflows/ci.yml/badge.svg)](https://github.com/zhiburt/tabled/actions) Stars:`1.9K`.
  * [nukesor/comfy-table](https://github.com/nukesor/comfy-table) [[comfy-table](https://crates.io/crates/comfy-table)] - Beautiful dynamic tables for your cli tools. [![Build status](https://github.com/Nukesor/comfy-table/workflows/Tests/badge.svg?branch=master)](https://github.com/nukesor/comfy-table/actions) Stars:`897`.
* Human-centered design
  * [rust-cli/human-panic](https://github.com/rust-cli/human-panic) [[human-panic](https://crates.io/crates/human-panic)] - panic messages for humans Stars:`1.6K`.
* Line editor
  * [kkawakam/rustyline](https://github.com/kkawakam/rustyline) [[rustyline](https://crates.io/crates/rustyline)] - readline implementation Stars:`1.5K`.
  * [murarth/linefeed](https://github.com/murarth/linefeed) [[linefeed](https://crates.io/crates/linefeed)] - Configurable, extensible, interactive line reader Stars:`184`.
  * [MovingtoMars/liner](https://github.com/MovingtoMars/liner) [[liner](https://crates.io/crates/liner)] - A library offering readline-like functionality Stars:`74`.
  * [srijs/rust-copperline](https://github.com/srijs/rust-copperline) [[copperline](https://crates.io/crates/copperline)] - command line editing library Stars:`27`.
* Other
  * [mgrachev/update-informer](https://github.com/mgrachev/update-informer) [[update-informer](https://crates.io/crates/update-informer)] - Update informer for CLI applications. It checks for a new version on Crates.io and GitHub [![build badge](https://github.com/mgrachev/update-informer/workflows/CI/badge.svg)](https://github.com/mgrachev/update-informer/actions) Stars:`205`.
* Pipeline
  * [oconnor663/duct.rs](https://github.com/oconnor663/duct.rs) [[duct](https://crates.io/crates/duct)] - A builder for subprocess pipelines and IO redirection Stars:`795`.
  * [hniksic/rust-subprocess](https://github.com/hniksic/rust-subprocess) [[subprocess](https://crates.io/crates/subprocess)] - facilities for interaction with external pipelines Stars:`425`.
  * [rust-cli/rexpect](https://github.com/rust-cli/rexpect) [[rexpect](https://crates.io/crates/rexpect)] - automate interactive applications such as ssh, ftp, passwd, etc [![CI](https://github.com/rust-cli/rexpect/actions/workflows/ci.yml/badge.svg)](https://github.com/rust-cli/rexpect/actions/workflows/ci.yml) Stars:`317`.
  * [zhiburt/expectrl](https://github.com/zhiburt/expectrl) [[expectrl](https://crates.io/crates/expectrl)] - A library for controlling interactive programs in a pseudo-terminal [![build badge](https://github.com/zhiburt/expectrl/actions/workflows/ci.yml/badge.svg)](https://github.com/zhiburt/expectrl/actions/workflows/ci.yml) Stars:`167`.
  * [imp/pager-rs](https://gitlab.com/imp/pager-rs) [[pager](https://crates.io/crates/pager)] - pipe your output through an external pager
* Progress
  * [console-rs/indicatif](https://github.com/console-rs/indicatif) [[indicatif](https://crates.io/crates/indicatif)] - indicate progress to users Stars:`4.2K`.
  * [a8m/pb](https://github.com/a8m/pb) [[pbr](https://crates.io/crates/pbr)] - console progress bar Stars:`581`.
  * [FGRibreau/spinners](https://github.com/FGRibreau/spinners) [[spinners](https://crates.io/crates/spinners)] - 60+ elegant terminal spinners Stars:`531`.
  * [etienne-napoleone/spinach](https://github.com/etienne-napoleone/spinach) [[spinach](https://crates.io/crates/spinach)] - Practical spinner. [![CI](https://github.com/etienne-napoleone/spinach/actions/workflows/ci.yml/badge.svg)](https://github.com/etienne-napoleone/spinach/actions/workflows/ci.yml) Stars:`91`.
* Prompt
  * [starship/starship](https://starship.rs/) [[starship](https://crates.io/crates/starship)] - A minimal, blazing fast, and extremely customizable prompt for any shell [![Build status](https://github.com/starship/starship/workflows/Main%20workflow/badge.svg?branch=master)](https://github.com/starship/starship/actions) Stars:`43.0K`.
  * [mikaelmello/inquire](https://github.com/mikaelmello/inquire) [[inquire](https://crates.io/crates/inquire)] - A library for building interactive prompts on terminals. [![Build status](https://github.com/mikaelmello/inquire/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/mikaelmello/inquire/actions) Stars:`1.7K`.
  * [ynqa/promkit](https://github.com/ynqa/promkit) [[promkit](https://crates.io/crates/promkit)] - A toolkit for building interactive command-line tools [![ci](https://github.com/ynqa/promkit/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/ynqa/promkit/actions/workflows/ci.yml) Stars:`238`.
  * [hashmismatch/terminal_cli.rs](https://github.com/hashmismatch/terminal_cli.rs) [[terminal_cli](https://crates.io/crates/terminal_cli)] - build an interactive command prompt Stars:`54`.
* Style
  * [colored](https://github.com/colored-rs/colored) [[colored](https://crates.io/crates/colored)] - Coloring terminal so simple, you already know how to do it! Stars:`1.7K`.
  * [console-rs/dialoguer](https://github.com/console-rs/dialoguer) [[dialoguer](https://crates.io/crates/dialoguer)] - Library for command line prompts and similar things. Stars:`1.3K`.
  * [ogham/rust-ansi-term](https://github.com/ogham/rust-ansi-term) [[ansi_term](https://crates.io/crates/ansi_term)] - control colours and formatting on ANSI terminals Stars:`447`.
  * [SergioBenitez/yansi](https://github.com/SergioBenitez/yansi) [[yansi](https://crates.io/crates/yansi)] - A dead simple ANSI terminal color painting library Stars:`232`.
  * [LukasKalbertodt/bunt](https://github.com/LukasKalbertodt/bunt) [[bunt](https://crates.io/crates/bunt)] - cross-platform terminal colors and styling with macros [![Build status](https://github.com/LukasKalbertodt/bunt/actions/workflows/ci.yml/badge.svg)](https://github.com/LukasKalbertodt/bunt/actions?query=workflow%3ACI+branch%3Amaster) Stars:`219`.
  * [LukasKalbertodt/term-painter](https://github.com/LukasKalbertodt/term-painter) [[term-painter](https://crates.io/crates/term-painter)] - cross-platform styled terminal output Stars:`78`.
* TUI
  * BearLibTerminal
  * [gyscos/Cursive](https://github.com/gyscos/Cursive) [[cursive](https://crates.io/crates/cursive)] - build rich TUI applications Stars:`4.2K`.
  * [ivanceras/titik](https://github.com/ivanceras/titik) - a crossplatform TUI widget library with the goal of providing interactive widgets Stars:`120`.
    * [cfyzium/bearlibterminal](https://github.com/nabijaczleweli/BearLibTerminal.rs) [[bear-lib-terminal](https://crates.io/crates/bear-lib-terminal)] - [BearLibTerminal](https://github.com/tommyettinger/BearLibTerminal) bindings Stars:`32`.
  * ncurses
  * [ratatui-org/ratatui](https://github.com/ratatui-org/ratatui) [[ratatui](https://crates.io/crates/ratatui)] - Library that's all about cooking up terminal user interfaces (TUIs) Stars:`8.9K`.
  * [redox-os/termion](https://github.com/redox-os/termion) [[termion](https://crates.io/crates/termion)] - bindless library for controlling terminals/TTY Stars:`2.1K`.
    * [jeaye/ncurses-rs](https://github.com/jeaye/ncurses-rs) [[ncurses](https://crates.io/crates/ncurses)] - [ncurses](https://www.gnu.org/software/ncurses/) bindings Stars:`671`.
    * [ihalila/pancurses](https://github.com/ihalila/pancurses) [[pancurses](https://crates.io/crates/pancurses)] - curses library, supports linux and windows Stars:`391`.
  * [ogham/rust-term-grid](https://github.com/ogham/rust-term-grid) [[term_grid](https://crates.io/crates/term_grid)] - Library for putting things in a grid Stars:`65`.
  * [ruterm](https://crates.io/crates/ruterm) - tiny & simple library for working with TTY
  * Termbox
  * [TimonPost/crossterm](https://github.com/crossterm-rs/crossterm) [[crossterm](https://crates.io/crates/crossterm)] - crossplatform terminal library Stars:`3.1K`.
    * [gchp/rustbox](https://github.com/gchp/rustbox) [[rustbox](https://crates.io/crates/rustbox)] - bindings to [Termbox](https://github.com/nsf/termbox) Stars:`467`.

### Compression

  * [dropbox/rust-brotli](https://github.com/dropbox/rust-brotli) - Brotli decompressor that optionally avoids the stdlib Stars:`798`.
  * [dyz1990/sevenz-rust](https://github.com/dyz1990/sevenz-rust) [[sevenz-rust](https://crates.io/crates/sevenz-rust)] - A 7z decompressor/compressor written in pure rust. [![Rust](https://github.com/dyz1990/sevenz-rust/workflows/Rust/badge.svg?branch=main)](https://github.com/dyz1990/sevenz-rust/actions) Stars:`145`.
  * [ende76/brotli-rs](https://github.com/ende76/brotli-rs) - implementation of Brotli compression Stars:`64`.
* [7z](https://7-zip.org/7z.html)
* [Brotli](https://opensource.googleblog.com/2015/09/introducing-brotli-new-compression.html)
* bzip2
  * [alexcrichton/bzip2-rs](https://github.com/alexcrichton/bzip2-rs) - [libbz2](https://www.sourceware.org/bzip2/) bindings Stars:`94`.
* gzip
  * [zopfli](https://github.com/zopfli-rs/zopfli) [[zopfli](https://crates.io/crates/zopfli)] - implementation of the Zopfli compression algorithm for higher quality deflate or zlib compression Stars:`32`.
* gzp
  * [sstadick/gzp](https://github.com/sstadick/gzp/) - multi-threaded encoding and decoding of deflate formats and snappy Stars:`151`.
* miniz
  * [rust-lang/flate2-rs](https://github.com/rust-lang/flate2-rs) - [miniz](https://code.google.com/archive/p/miniz) bindings [![build badge](https://github.com/rust-lang/flate2-rs/workflows/CI/badge.svg?branch=master)](https://github.com/rust-lang/flate2-rs/actions) Stars:`860`.
* tar
  * [alexcrichton/tar-rs](https://github.com/alexcrichton/tar-rs) - tar archive reading/writing Stars:`615`.
* zip
  * [zip-rs/zip2](https://github.com/zip-rs/zip2) [[zip](https://crates.io/crates/zip)] - read and write  ZIP archives Stars:`75`.
* zstd
  * [gyscos/zstd-rs](https://github.com/gyscos/zstd-rs) - rust binding for the zstd compression library Stars:`478`.

### Computation

* [dimforge/nalgebra](https://github.com/dimforge/nalgebra) - low-dimensional linear algebra library Stars:`3.8K`.
* [faer-rs](https://github.com/sarah-ek/faer-rs) [[faer](https://crates.io/crates/faer)] - Linear algebra foundation for Rust Stars:`1.7K`.
* [calebwin/emu](https://github.com/calebwin/emu) - A language for GPGPU numerical computing Stars:`1.6K`.
* [argmin-rs/argmin](https://github.com/argmin-rs/argmin) [[argmin](https://crates.io/crates/argmin)] - Optimization library Stars:`937`.
  * [GuillaumeGomez/rust-GSL](https://github.com/GuillaumeGomez/rust-GSL) - GSL bindings Stars:`189`.
  * [mikkyang/rust-blas](https://github.com/mikkyang/rust-blas) - BLAS bindings Stars:`82`.
  * [stainless-steel/lapack](https://github.com/blas-lapack-rs/lapack) - LAPACK bindings Stars:`81`.
* [BLAS](https://en.wikipedia.org/wiki/Basic_Linear_Algebra_Subprograms) [[blas](https://crates.io/keywords/blas)]
* [GSL](http://www.gnu.org/software/gsl/)
* [LAPACK](https://en.wikipedia.org/wiki/LAPACK)
* Parallel
  * [arrayfire/arrayfire-rust](https://github.com/arrayfire/arrayfire-rust) - [Arrayfire](https://github.com/arrayfire) bindings Stars:`810`.
  * [autumnai/collenchyma](https://github.com/autumnai/collenchyma) - An extensible, pluggable, backend-agnostic framework for parallel, high-performance computations on CUDA, OpenCL and common host CPU. Stars:`474`.
  * [luqmana/rust-opencl](https://github.com/luqmana/rust-opencl) - [OpenCL](https://www.khronos.org/opencl/) bindings Stars:`171`.
* Science
  * [indigits/scirust](https://github.com/indigits/scirust) - Scientific computing library Stars:`262`.
  * [cpmech/russell](https://github.com/cpmech/russell) - Rust Scientific Library for numerical mathematics, ordinary differential equations, special math functions, high-performance (sparse) linear algebra Stars:`92`.
* Statrs
  * [statrs-dev/statrs](https://github.com/statrs-dev/statrs) - Robust statistical computation library Stars:`546`.

### Concurrency

* [Rayon](https://github.com/rayon-rs/rayon) - A data parallelism library Stars:`10.5K`.
* [crossbeam-rs/crossbeam](https://github.com/crossbeam-rs/crossbeam) - Support for parallelism and low-level concurrency Stars:`7.1K`.
* [zonyitoo/coio-rs](https://github.com/zonyitoo/coio-rs) - Coroutine I/O Stars:`455`.
* [rustcc/coroutine-rs](https://github.com/rustcc/coroutine-rs) - Coroutine Library Stars:`415`.
* [orium/archery](https://github.com/orium/archery) [[archery](https://crates.io/crates/archery)] - Library to abstract from `Rc`/`Arc` pointer types. [![build badge](https://github.com/orium/archery/workflows/CI/badge.svg)](https://github.com/orium/archery/actions?query=workflow%3ACI) Stars:`134`.

### Configuration

* [mehcode/config-rs](https://github.com/mehcode/config-rs) [[config](https://crates.io/crates/config)] - Layered configuration system (with strong support for 12-factor applications). Stars:`2.4K`.
* [softprops/envy](https://github.com/softprops/envy) - deserialize env vars into typesafe structs [![Main](https://github.com/softprops/envy/actions/workflows/main.yml/badge.svg)](https://github.com/softprops/envy/actions/workflows/main.yml) Stars:`831`.
* [SergioBenitez/Figment](https://github.com/SergioBenitez/Figment) [[figment](https://crates.io/crates/figment)] - A configuration library so con-free, it's unreal. Stars:`510`.
* [Kixunil/configure_me](https://github.com/Kixunil/configure_me) [[configure_me](https://crates.io/crates/configure_me)] - library for processing application configuration easily Stars:`60`.
* [andoriyu/uclicious](https://github.com/andoriyu/uclicious) [[uclicious](https://crates.io/crates/uclicious)] - [libUCL](https://github.com/vstakhov/libucl) based feature-rich configuration library. [![CircleCI](https://circleci.com/gh/vstakhov/libucl.svg?style=svg)](https://app.circleci.com/pipelines/github/vstakhov/libucl) Stars:`17`.

### Cryptography

[[crypto](https://crates.io/keywords/crypto), [cryptography](https://crates.io/keywords/cryptography)]

* [rustls/rustls](https://github.com/rustls/rustls) - Implementation of TLS Stars:`5.7K`.
* [briansmith/ring](https://github.com/briansmith/ring) - Safe, fast, small crypto using Rust and BoringSSL's cryptography primitives. Stars:`3.6K`.
* [cossacklabs/themis](https://github.com/cossacklabs/themis) [[themis](https://crates.io/crates/themis)] - a high-level cryptographic library for solving typical data security tasks, best fit for multi-platform apps. [![build badge](https://circleci.com/gh/cossacklabs/themis/tree/master.svg?style=shield)](https://app.circleci.com/pipelines/github/cossacklabs/themis) Stars:`1.8K`.
* [RustCrypto/hashes](https://github.com/RustCrypto/hashes) - Collection of cryptographic hash functions Stars:`1.8K`.
* [DaGenix/rust-crypto](https://github.com/DaGenix/rust-crypto) - cryptographic algorithms Stars:`1.4K`.
* [sfackler/rust-openssl](https://github.com/sfackler/rust-openssl) - [OpenSSL](https://www.openssl.org/) bindings Stars:`1.4K`.
* [exonum/exonum](https://github.com/exonum/exonum) [[exonum](https://crates.io/crates/exonum)] - extensible framework for blockchain projects Stars:`1.2K`.
* [dalek-cryptography/curve25519-dalek](https://github.com/dalek-cryptography/curve25519-dalek) - Curve25519 operations Stars:`850`.
* [dalek-cryptography/ed25519-dalek](https://github.com/dalek-cryptography/ed25519-dalek) - Ed25519 digital signatures Stars:`672`.
* [orion-rs/orion](https://github.com/orion-rs/orion) - This library aims to provide easy and usable crypto. 'Usable' meaning exposing high-level API's that are easy to use and hard to misuse. [![Tests](https://github.com/orion-rs/orion/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/orion-rs/orion/actions/workflows/test.yml) Stars:`539`.
* [sfackler/rust-native-tls](https://github.com/sfackler/rust-native-tls) - Bindings for native TLS libraries Stars:`466`.
* [briansmith/webpki](https://github.com/briansmith/webpki) - Web PKI TLS X.509 certificate validation. Stars:`457`.
* [dalek-cryptography/x25519-dalek](https://github.com/dalek-cryptography/x25519-dalek) - X25519 key exchange Stars:`326`.
* [w3f/schnorrkel](https://github.com/w3f/schnorrkel) - Schnorr VRFs and signatures on the Ristretto group Stars:`305`.
* [facebook/opaque-ke](https://github.com/facebook/opaque-ke) - Implementation of the recent [OPAQUE](https://datatracker.ietf.org/doc/draft-krawczyk-cfrg-opaque/) password-authenticated key exchange. [![build badge](https://github.com/facebook/opaque-ke/workflows/Rust%20CI/badge.svg?branch=master)](https://github.com/facebook/opaque-ke) Stars:`285`.
* [kornelski/rust-security-framework](https://github.com/kornelski/rust-security-framework) - Bindings for Security Framework (OSX native) Stars:`229`.
* [arkworks-rs/circom-compat](https://github.com/arkworks-rs/circom-compat) - Arkworks bindings to Circom's R1CS, for Groth16 Proof and Witness generation. Stars:`227`.
* [debris/tiny-keccak](https://github.com/debris/tiny-keccak) - Keccak family (SHA3) Stars:`192`.
* [conradkleinespel/rooster](https://github.com/conradkleinespel/rooster) [[rooster](https://crates.io/crates/rooster)] - Simple password manager to use in your terminal Stars:`151`.
* [libOctavo/octavo](https://github.com/libOctavo/octavo) - Modular hash and crypto library Stars:`140`.
* [klutzy/suruga](https://github.com/klutzy/suruga) - Implementation of [TLS 1.2](https://datatracker.ietf.org/doc/html/rfc5246) Stars:`124`.
* [racum/rust-djangohashers](https://github.com/racum/rust-djangohashers) [[djangohashers](https://crates.io/crates/djangohashers)] - Port of the password primitives used in the Django Project. It doesn't require Django, only hashes and validates passwords according to its style. Stars:`56`.
* [sorairolake/abcrypt](https://github.com/sorairolake/abcrypt) [[abcrypt](https://crates.io/crates/abcrypt)] - A simple, modern and secure file encryption library. [![CI](https://github.com/sorairolake/abcrypt/workflows/CI/badge.svg?branch=develop)](https://github.com/sorairolake/abcrypt/actions?query=workflow%3ACI) Stars:`10`.
* [iddm/randomorg](https://github.com/iddm/randomorg) - A random.org client library. [![Crates badge](https://img.shields.io/crates/v/randomorg.svg)](https://crates.io/crates/randomorg) Stars:`8`.
* [sorairolake/scryptenc-rs](https://github.com/sorairolake/scryptenc-rs) [[scryptenc](https://crates.io/crates/scryptenc)] - An implementation of the scrypt encrypted data format. [![CI](https://github.com/sorairolake/scryptenc-rs/workflows/CI/badge.svg?branch=develop)](https://github.com/sorairolake/scryptenc-rs/actions?query=workflow%3ACI) Stars:`1`.

### Data processing

* [pola-rs/polars](https://github.com/pola-rs/polars) - Fast feature complete DataFrame library ![Build and test](https://github.com/pola-rs/polars/workflows/Build%20and%20test/badge.svg?branch=master) Stars:`27.9K`.
* [pg_analytics](https://github.com/paradedb/paradedb/tree/dev/pg_analytics) - PostgreSQL extension that accelerates analytical query processing inside Postgres to a performance level comparable to dedicated OLAP databases. Stars:`4.7K`.
* [pg_lakehouse](https://github.com/paradedb/paradedb/tree/dev/pg_lakehouse) - PostgreSQL extension that transforms Postgres into an analytical query engine over object stores like AWS S3/GCS and table formats like Delta Lake/Iceberg. Stars:`4.7K`.
* [bluss/ndarray](https://github.com/rust-ndarray/ndarray) - N-dimensional array with array views, multidimensional slicing, and efficient operations Stars:`3.4K`.
* [weld-project/weld](https://github.com/weld-project/weld) - High-performance runtime for data analytics applications Stars:`3.0K`.
* [amv-dev/yata](https://github.com/amv-dev/yata) - high performance technical analysis library [![Build Status](https://img.shields.io/github/workflow/status/amv-dev/yata/Rust?branch=master)](https://github.com/amv-dev/yata/actions?query=workflow%3ARust) Stars:`306`.
* [kernelmachine/utah](https://github.com/kernelmachine/utah) - Dataframe structure and operations Stars:`142`.

### Data streaming

* [ArroyoSystems/arroyo](https://github.com/ArroyoSystems/arroyo) - High-performance real-time analytics in Rust and SQL [![CI](https://github.com/ArroyoSystems/arroyo/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/ArroyoSystems/arroyo/actions) Stars:`3.5K`.
* [infinyon/fluvio](https://github.com/infinyon/fluvio) - Programmable data streaming platform [![CI](https://github.com/infinyon/fluvio/workflows/CI/badge.svg?branch=stable)](https://github.com/infinyon/fluvio/actions) Stars:`2.8K`.
* [iggy-rs/iggy](https://github.com/iggy-rs/iggy) [[iggy](https://crates.io/crates/iggy)] - Persistent message streaming platform, supporting QUIC, TCP and HTTP transport protocols [![CI](https://github.com/iggy-rs/iggy/actions/workflows/test.yml/badge.svg)](https://github.com/iggy-rs/iggy/actions/workflows/test.yml) Stars:`1.8K`.

### Data structures

* [rust-itertools/itertools](https://github.com/rust-itertools/itertools) - Extra iterator adaptors, functions and macros Stars:`2.6K`.
* [greyblake/nutype](https://github.com/greyblake/nutype) [[nutype](https://crates.io/crates/nutype)] - define newtype structures with validation constraints. [![build status](https://github.com/greyblake/nutype/actions/workflows/ci.yml/badge.svg)](https://github.com/greyblake/nutype/actions) Stars:`1.3K`.
* [orium/rpds](https://github.com/orium/rpds) [[rpds](https://crates.io/crates/rpds)] - Persistent data structures. [![build badge](https://github.com/orium/rpds/workflows/CI/badge.svg)](https://github.com/orium/rpds/actions?query=workflow%3ACI) Stars:`1.2K`.
* [ashvardanian/simsimd](https://github.com/ashvardanian/SimSIMD) - SIMD-accelerated vector distances and similarity functions for x86 AVX2 & AVX-512, and Arm NEON [![crates.io](https://img.shields.io/crates/v/simsimd.svg)](https://crates.io/crates/simsimd) Stars:`802`.
* [RoaringBitmap/roaring-rs](https://github.com/RoaringBitmap/roaring-rs) - Roaring Bitmaps Stars:`713`.
* [fizyk20/generic-array](https://github.com/fizyk20/generic-array) - a hack to allow for arrays sized by typenums Stars:`402`.
* [yamafaktory/hypergraph](https://github.com/yamafaktory/hypergraph) [[hypergraph](https://crates.io/crates/hypergraph)] - Hypergraph is a data structure library to generate directed hypergraphs. [![ci](https://github.com/yamafaktory/hypergraph/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/yamafaktory/hypergraph/actions/workflows/ci.yml) Stars:`273`.
* [tnballo/scapegoat](https://github.com/tnballo/scapegoat) [[scapegoat](https://crates.io/crates/scapegoat)] - Safe, fallible, stack-only alternative to `BTreeSet` and `BTreeMap`. [![GitHub Actions](https://github.com/tnballo/scapegoat/workflows/test/badge.svg?branch=master)](https://github.com/tnballo/scapegoat/actions) Stars:`243`.
* [mrhooray/kdtree-rs](https://github.com/mrhooray/kdtree-rs) - K-dimensional tree for fast geospatial indexing and nearest neighbors lookup Stars:`221`.
* [becheran/grid](https://github.com/becheran/grid) [[grid](https://crates.io/crates/grid)] - Provide a two dimensional data structure that is easy to use and fast. [![build status](https://github.com/becheran/grid/actions/workflows/rust.yml/badge.svg)](https://github.com/becheran/grid/actions) Stars:`80`.
* [danielpclark/array_tool](https://github.com/danielpclark/array_tool) - Array helpers. Some of the most common methods you would use on Arrays made available on Vectors. Polymorphic implementations for handling most of your use cases. Stars:`77`.
* [billyevans/tst](https://github.com/billyevans/tst) [[tst](https://crates.io/crates/tst)] - Ternary search tree collection Stars:`23`.
* [contain-rs](https://github.com/contain-rs) - Extension of Rust's std::collections
* [xfix/enum-map](https://codeberg.org/xfix/enum-map) [[enum-map](https://crates.io/crates/enum-map)] - An optimized map implementation for enums using an array to store values.
* [garro95/priority-queue](https://github.com/garro95/priority-queue)[[priority-queue](https://crates.io/crates/priority-queue)] - A priority queue that implements priority changes.

### Data visualization

* [rerun](https://github.com/rerun-io/rerun) - [[rerun](https://crates.io/crates/rerun)] - An SDK for logging computer vision and robotics data (tensors, point clouds, etc) paired with a visualizer for exploring that data over time. Stars:`5.7K`.
* [plotters](https://github.com/plotters-rs/plotters) - [![build badge](https://github.com/plotters-rs/plotters/workflows/CI/badge.svg)](https://github.com/plotters-rs/plotters/actions) Stars:`3.7K`.
* [plotly](https://github.com/plotly/plotly.rs) - Plotly for Rust Stars:`1.0K`.
* [milliams/plotlib](https://github.com/milliams/plotlib) - Data plotting library for Rust Stars:`465`.
* [blitzarx1/egui_graphs](https://github.com/blitzarx1/egui_graphs) [[egui_graphs](https://crates.io/crates/egui_graphs)] - Interactive graph visualization widget powered by egui and petgraph. [![Crates.io](https://img.shields.io/crates/v/egui_graphs)](https://crates.io/crates/egui_graphs) [![docs.rs](https://img.shields.io/docsrs/egui_graphs)](https://docs.rs/egui_graphs) Stars:`346`.
* [mazznoer/colorgrad-rs](https://github.com/mazznoer/colorgrad-rs) [[colorgrad](https://crates.io/crates/colorgrad)] - Color scales library for data visualization, charts, games, maps, generative art and others. Stars:`272`.
* [saresend/gust](https://github.com/saresend/Gust) - A small charting/visualization tool and partial vega implementation Stars:`131`.
* [djduque/pgfplots](https://github.com/djduque/pgfplots) [[pgfplots](https://crates.io/crates/pgfplots)] - Library to generate publication-quality figures. [![build](https://github.com/DJDuque/pgfplots/actions/workflows/rust.yml/badge.svg)](https://github.com/DJDuque/pgfplots/actions/workflows/rust.yml) Stars:`118`.
* [plotpy](https://github.com/cpmech/plotpy) [[plotpy](https://crates.io/crates/plotpy)] - Rust plotting library using Python (Matplotlib) Stars:`51`.

### Database

[[database](https://crates.io/keywords/database)]

* NoSQL [[nosql](https://crates.io/keywords/nosql)]

    * [AlexPikalov/cdrs](https://github.com/AlexPikalov/cdrs) [[cdrs](https://crates.io/crates/cdrs)] - native client Stars:`346`.
    * [krojew/cdrs-tokio](https://github.com/krojew/cdrs-tokio) - High-level async Cassandra client written in 100% Rust. [![build badge](https://github.com/krojew/cdrs-tokio/actions/workflows/rust.yml/badge.svg)](https://github.com/krojew/cdrs-tokio/actions) Stars:`131`.
    * [cassandra-rs](https://github.com/cassandra-rs/cassandra-rs) - bindings to the DataStax C/C++ Stars:`129`.
    * [Arangors](https://github.com/fMeow/arangors) [[arangors](https://crates.io/crates/arangors)] - An ArangoDB driver Stars:`125`.
  * [ArangoDB](https://arangodb.com)
    * [Aragog](https://gitlab.com/qonfucius/aragog) [[aragog](https://crates.io/crates/aragog)] - A Lightweight ArangoDB Object document, relational and graph mapper [![pipeline status](https://gitlab.com/qonfucius/aragog/badges/master/pipeline.svg)](https://gitlab.com/qonfucius/aragog/-/commits/master)
  * [Cassandra](https://cassandra.apache.org/_/index.html) [[cassandra](https://crates.io/keywords/cassandra), [cql](https://crates.io/keywords/cql)]
      * [[cassandra-protocol](https://crates.io/crates/cassandra-protocol)] - Cassandra protocol implementation.
      * [[cdrs-tokio](https://crates.io/crates/cdrs-tokio)] - production-ready async Apache Cassandra driverclient
  * CouchDB [[couchdb](https://crates.io/keywords/couchdb)]
    * [softprops/dynomite](https://github.com/softprops/dynomite) - A library for strongly-typed and convenient interaction with `rusoto_dynamodb` [![build badge](https://github.com/softprops/dynomite/workflows/Main/badge.svg?branch=master)](https://github.com/softprops/dynomite/actions) Stars:`215`.
    * [chill-rs/chill](https://github.com/chill-rs/chill) [[couchdb](https://crates.io/crates/chill)] - Client for the CouchDB REST API Stars:`35`.
  * [DynamoDB](https://aws.amazon.com/dynamodb/) [[dynamodb](https://crates.io/keywords/dynamodb)]
  * Elasticsearch [[elasticsearch](https://crates.io/keywords/elasticsearch)]
    * [elastic-rs/elastic](https://github.com/elastic-rs/elastic) [[elastic](https://crates.io/crates/elastic)] - elastic is an efficient, modular API client for Elasticsearch written in Rust [![build badge](https://ci.appveyor.com/api/projects/status/csa78tcumdpnbur2?svg=true)](https://ci.appveyor.com/project/KodrAus/elastic) Stars:`253`.
    * [benashford/rs-es](https://github.com/benashford/rs-es) [[rs-es](https://crates.io/crates/rs-es)] - Client for the [Elastic](https://www.elastic.co/) REST API Stars:`217`.
  * etcd
    * [jimmycuadra/rust-etcd](https://github.com/jimmycuadra/rust-etcd) [[etcd](https://crates.io/crates/etcd)] - A client library for CoreOS's etcd. Stars:`143`.
    * [driftluo/InfluxDBClient-rs](https://github.com/driftluo/InfluxDBClient-rs) - Synchronization interface Stars:`82`.
  * [InfluxDB](https://www.influxdata.com/)
  * LevelDB
    * [skade/leveldb](https://github.com/skade/leveldb) - [LevelDB](https://github.com/google/leveldb) bindings Stars:`182`.
  * LMDB [[lmdb](https://crates.io/keywords/lmdb)]
    * [vhbit/lmdb-rs](https://github.com/vhbit/lmdb-rs) [[lmdb-rs](https://crates.io/crates/lmdb-rs)] - [LMDB](https://www.symas.com/symas-embedded-database-lmdb) bindings Stars:`113`.
  * MongoDB [[mongodb](https://crates.io/keywords/mongodb)]
    * [Redb](https://github.com/cberner/redb) - An embedded key-value database. It provides a similar interface to other embedded key-value stores such as rocksdb and lmdb. ![GitHub Workflow Status](https://github.com/cberner/redb/actions/workflows/ci.yml/badge.svg) Stars:`3.1K`.
    * [mongodb/mongo-rust-driver](https://github.com/mongodb/mongo-rust-driver) [[mongodb](https://crates.io/crates/mongodb)] - [MongoDB](https://www.mongodb.com/) bindings Stars:`1.4K`.
    * [PoloDB](https://github.com/PoloDB/PoloDB) - An embedded JSON-based database has API similar to MongoDB. ![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/PoloDB/PoloDB/rust.yml) Stars:`813`.
    * [seladb/pickledb-rs](https://github.com/seladb/pickledb-rs) - a lightweight and simple key-value store, heavily inspired by Python's PickleDB. Stars:`253`.
  * [PickleDB](https://pythonhosted.org/pickleDB/)
  * [PoloDB](https://www.polodb.org/)
  * [Redb](https://www.redb.org/)
  * Redis [[redis](https://crates.io/keywords/redis)]
    * [surrealdb/surrealdb](https://github.com/surrealdb/surrealdb) - SurrealDB embedded document-graph database Stars:`26.2K`.
    * [redis-rs](https://github.com/redis-rs/redis-rs) - [Redis](https://redis.io/) library [![Rust](https://github.com/redis-rs/redis-rs/actions/workflows/rust.yml/badge.svg)](https://github.com/redis-rs/redis-rs/actions/workflows/rust.yml) Stars:`3.5K`.
  * [UnQLite](https://github.com/symisc/unqlite) Stars:`2.1K`.
    * [rust-rocksdb/rust-rocksdb](https://github.com/rust-rocksdb/rust-rocksdb) - RocksDB bindings [![RocksDB CI](https://github.com/rust-rocksdb/rust-rocksdb/actions/workflows/rust.yml/badge.svg?branch=master)](https://github.com/rust-rocksdb/rust-rocksdb/actions/workflows/rust.yml) Stars:`1.8K`.
    * [aembke/fred](https://github.com/aembke/fred.rs) [[fred](https://crates.io/crates/fred)] - A high level async [Redis](https://redis.io/) client for Rust with Tokio. [![CircleCI](https://circleci.com/gh/aembke/fred.rs/tree/main.svg?style=svg)]([https://circleci.com/gh/aembke/fred.rs/tree/main](https://app.circleci.com/pipelines/github/aembke/fred.rs?branch=main)) Stars:`348`.
    * [bonifaido/rust-zookeeper](https://github.com/bonifaido/rust-zookeeper) [[zookeeper](https://crates.io/crates/zookeeper)] - A client library for Apache ZooKeeper. Stars:`205`.
    * [zitsen/unqlite.rs](https://github.com/zitsen/unqlite.rs) - UnQLite bindings Stars:`109`.
    * [krojew/rust-zookeeper](https://github.com/krojew/rust-zookeeper) [[zookeeper-async](https://crates.io/crates/zookeeper-async)] - Async Zookeeper client, based on tokio.  ![build status](https://github.com/krojew/rust-zookeeper/actions/workflows/rust.yml/badge.svg) Stars:`22`.
  * [RocksDB](https://rocksdb.org/)
  * [SurrealDB](https://surrealdb.com/)
  * [ZooKeeper](https://zookeeper.apache.org/)
* OGM [[ogm](https://crates.io/keywords/ogm)]
    * [Aragog](https://gitlab.com/qonfucius/aragog) [[aragog](https://crates.io/crates/aragog)] - A Lightweight ArangoDB Object document, relational and graph mapper [![pipeline status](https://gitlab.com/qonfucius/aragog/badges/master/pipeline.svg)](https://gitlab.com/qonfucius/aragog/-/commits/master)
* ORM [[orm](https://crates.io/keywords/orm)]
  * [diesel-rs/diesel](https://github.com/diesel-rs/diesel) - an ORM and Query builder Stars:`12.3K`.
  * [SeaQL/sea-orm](https://github.com/SeaQL/sea-orm) - 🐚 An async & dynamic ORM  [![crate](https://img.shields.io/crates/v/sea-orm.svg)](https://crates.io/crates/sea-orm) [![docs](https://img.shields.io/docsrs/sea-orm/latest)](https://docs.rs/sea-orm) [![build status](https://github.com/SeaQL/sea-orm/actions/workflows/rust.yml/badge.svg)](https://github.com/SeaQL/sea-orm/actions/workflows/rust.yml) Stars:`6.6K`.
  * [rbatis/rbatis](https://github.com/rbatis/rbatis) - ORM Framework High Performance(JSON based) Stars:`2.2K`.
  * [Brendonovich/prisma-client-rust](https://github.com/Brendonovich/prisma-client-rust) - An autogenerated query builder that provides simple and fully type-safe database access using the Prisma ecosystem. [![Test Status](https://img.shields.io/github/workflow/status/Brendonovich/prisma-client-rust/CI?label=tests&style=flat-square)](https://github.com/Brendonovich/prisma-client-rust/actions) Stars:`1.8K`.
* [sfackler/r2d2](https://github.com/sfackler/r2d2) - generic connection pool Stars:`1.5K`.
  * [SeaQL/seaography](https://github.com/SeaQL/seaography) - 🧭 GraphQL framework for SeaORM [![crate](https://img.shields.io/crates/v/seaography.svg)](https://crates.io/crates/seaography) [![docs](https://img.shields.io/docsrs/seaography/latest)](https://docs.rs/seaography) [![build status](https://github.com/SeaQL/seaography/actions/workflows/tests.yaml/badge.svg)](https://github.com/SeaQL/seaography/actions/workflows/tests.yaml) Stars:`358`.
  * [mjovanc/njord](https://github.com/mjovanc/njord) - ⛵ A lightweight ORM library for Rust [![build status](https://github.com/mjovanc/njord/actions/workflows/ci.yml/badge.svg)](https://github.com/mjovanc/njord/actions/workflows/ci.yml) ![crates.io](https://img.shields.io/crates/v/njord.svg) Stars:`299`.
  * [ivanceras/rustorm](https://github.com/ivanceras/rustorm) - an ORM Stars:`250`.
* SQL [[sql](https://crates.io/keywords/sql)]
  * Generic
    * [launchbadge/sqlx](https://github.com/launchbadge/sqlx) - async PostgreSQL/MySQL/SQLite connection pool with strong typing support [![build badge](https://img.shields.io/github/workflow/status/launchbadge/sqlx/Rust/master?style=flat-square)](https://github.com/launchbadge/sqlx) Stars:`12.4K`.
    * [SeaQL/sea-query](https://github.com/SeaQL/sea-query) - 🔱 A dynamic SQL query builder for MySQL, Postgres and SQLite [![crate](https://img.shields.io/crates/v/sea-query.svg)](https://crates.io/crates/sea-query) [![docs](https://img.shields.io/docsrs/sea-query/latest)](https://docs.rs/sea-query) [![build status](https://github.com/SeaQL/sea-query/actions/workflows/rust.yml/badge.svg)](https://github.com/SeaQL/sea-query/actions/workflows/rust.yml) Stars:`1.1K`.
    * [SeaQL/sea-schema](https://github.com/SeaQL/sea-schema) - 🌿 SQL schema definition and discovery [![crate](https://img.shields.io/crates/v/sea-schema.svg)](https://crates.io/crates/sea-schema) [![docs](https://img.shields.io/docsrs/sea-schema/latest)](https://docs.rs/sea-schema) [![build status](https://github.com/SeaQL/sea-schema/actions/workflows/rust.yml/badge.svg)](https://github.com/SeaQL/sea-schema/actions/workflows/rust.yml) Stars:`180`.
  * Microsoft SQL
    * [prisma/tiberius](https://github.com/prisma/tiberius) - [![Cargo tests](https://github.com/prisma/tiberius/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/prisma/tiberius/actions/workflows/test.yml) Stars:`306`.
  * MySql [[mysql](https://crates.io/keywords/mysql)]
    * [blackbeam/rust-mysql-simple](https://github.com/blackbeam/rust-mysql-simple) [[mysql](https://crates.io/crates/mysql)] - A native MySql client Stars:`651`.
    * [blackbeam/mysql_async](https://github.com/blackbeam/mysql_async) [[mysql_async](https://crates.io/crates/mysql_async)] - asynchronous Mysql driver based on Tokio. [![CircleCI](https://circleci.com/gh/blackbeam/mysql_async/tree/master.svg?style=shield)](https://app.circleci.com/pipelines/github/blackbeam/mysql_async?branch=master) Stars:`372`.
    * [AgilData/mysql-proxy-rs](https://github.com/AgilData/mysql-proxy-rs) - A MySQL Proxy [![CircleCI](https://circleci.com/gh/AgilData/mysql-proxy-rs/tree/master.svg?style=svg)](https://app.circleci.com/pipelines/github/AgilData/mysql-proxy-rs?branch=master) Stars:`193`.
  * Oracle
    * [kubo/rust-oracle](https://github.com/kubo/rust-oracle) [[oracle](https://crates.io/crates/oracle)] - Oracle driver [![build badge](https://github.com/kubo/rust-oracle/actions/workflows/run-tests.yml/badge.svg?branch=master)](https://github.com/kubo/rust-oracle/actions/workflows/run-tests.yml) Stars:`179`.
  * PostgreSql [[postgres](https://crates.io/keywords/postgres), [postgresql](https://crates.io/keywords/postgresql)]
    * [sfackler/rust-postgres](https://github.com/sfackler/rust-postgres) [[postgres](https://crates.io/crates/postgres)] - A native [PostgreSQL](https://www.postgresql.org/) client Stars:`3.4K`.
  * Sqlite [[sqlite](https://crates.io/keywords/sqlite)]
    * [rusqlite](https://github.com/rusqlite/rusqlite) - [Sqlite3](https://www.sqlite.org/index.html) bindings Stars:`2.9K`.

### Date and time

[[date](https://crates.io/keywords/date), [time](https://crates.io/keywords/time)]

* [chronotope/chrono](https://github.com/chronotope/chrono) - Date and time library Stars:`3.2K`.
* [time-rs/time](https://github.com/time-rs/time) - [![build badge](https://github.com/time-rs/time/workflows/Build/badge.svg)](https://github.com/time-rs/time/actions) Stars:`1.1K`.
* [Mnwa/ms](https://github.com/Mnwa/ms) [[ms-converter](https://crates.io/crates/ms-converter)] - it's a library for converting human-like times to milliseconds [![build badge](https://github.com/Mnwa/ms/workflows/build/badge.svg?branch=master)](https://github.com/Mnwa/ms/actions?query=workflow%3Abuild) Stars:`36`.
* [sorairolake/nt-time](https://github.com/sorairolake/nt-time) [[nt-time](https://crates.io/crates/nt-time)] - A Windows file time library. [![CI](https://github.com/sorairolake/nt-time/workflows/CI/badge.svg?branch=develop)](https://github.com/sorairolake/nt-time/actions?query=workflow%3ACI) Stars:`5`.

### Distributed systems

* Antimony
  * [antimonyproject/antimony](https://github.com/antimonyproject/antimony) [[antimony](https://crates.io/crates/antimony)] - stream processing / distributed computation platform Stars:`64`.
* Apache Kafka
  * [fede1024/rust-rdkafka](https://github.com/fede1024/rust-rdkafka) [[rdkafka](https://crates.io/crates/rdkafka)] - [librdkafka](https://github.com/confluentinc/librdkafka) bindings Stars:`1.5K`.
  * [kafka-rust/kafka-rust](https://github.com/kafka-rust/kafka-rust) - Rust client for Apache Kafka Stars:`1.2K`.
  * [gklijs/schema_registry_converter](https://github.com/gklijs/schema_registry_converter) [[schema_registry_converter](https://crates.io/crates/schema_registry_converter)] - to integrate with [confluent schema registry](https://www.confluent.io/product/confluent-platform/data-compatibility/) Stars:`98`.
* HDFS
  * [hyunsik/hdfs-rs](https://github.com/hyunsik/hdfs-rs) [[hdfs](https://crates.io/crates/hdfs)] - libhdfs bindings Stars:`33`.
* Other
  * [build-trust/ockam](https://github.com/build-trust/ockam) [[ockam](https://crates.io/crates/ockam)] - End-to-End Encryption, Mutual Authentication, and ABAC for distributed applications [![build badge](https://github.com/build-trust/ockam/workflows/Rust/badge.svg)](https://github.com/build-trust/ockam) Stars:`4.4K`.

### Domain driven design

  * [serverlesstechnology/cqrs](https://github.com/serverlesstechnology/cqrs) [[cqrs-es](https://crates.io/crates/cqrs-es)] - A framework for CQRS and event sourcing with [user guide](https://doc.rust-cqrs.org/) Stars:`346`.

### eBPF

* [aya/aya-rs](https://github.com/aya-rs/aya) - Built with a focus on developer experience and operability. Stars:`2.9K`.
* [libbpf/libbpf-rs](https://github.com/libbpf/libbpf-rs) - A minimal and opinionated eBPF tooling. Stars:`678`.

### Email

[[email](https://crates.io/keywords/email), [imap](https://crates.io/keywords/imap), [smtp](https://crates.io/keywords/smtp)]

* [lettre/lettre](https://github.com/lettre/lettre) - an SMTP-library [![CI](https://github.com/lettre/lettre/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/lettre/lettre/actions/workflows/test.yml) Stars:`1.8K`.
* [tweedegolf/mailcrab](https://github.com/tweedegolf/mailcrab) - Email test server for development. Stars:`673`.
* [meli/meli](https://github.com/meli/meli) - 🐝 terminal mail client Stars:`624`.
* [jdrouet/mrml](https://github.com/jdrouet/mrml) - A library to generate nice email templates working on any mail client. Stars:`319`.
* [stalwartlabs/mail-parser](https://github.com/stalwartlabs/mail-parser) [[mail-parser](https://crates.io/crates/mail-parser)] - A fast and robust e-mail parsing library with full MIME support [![build badge](https://github.com/stalwartlabs/mail-parser/actions/workflows/rust.yml/badge.svg)](https://github.com/stalwartlabs/mail-parser/actions/workflows/rust.yml) Stars:`272`.
* [stalwartlabs/mail-send](https://github.com/stalwartlabs/mail-send) [[mail-send](https://crates.io/crates/mail-send)] - E-mail builder and SMTP client library with DKIM support [![build badge](https://github.com/stalwartlabs/mail-send/actions/workflows/rust.yml/badge.svg)](https://github.com/stalwartlabs/mail-send/actions/workflows/rust.yml) Stars:`192`.
* [staktrace/mailparse](https://github.com/staktrace/mailparse) [[mailparse](https://crates.io/crates/mailparse)] - A library for parsing real-world email files Stars:`180`.
* [mailtutan/mailtutan](https://github.com/mailtutan/mailtutan) - An SMTP server for test and development environment. Stars:`159`.
* [jdrouet/catapulte](https://github.com/jdrouet/catapulte) - A microservice to send emails using [MRML](https://github.com/jdrouet/mrml) templates. Stars:`136`.
* [jdrouet/jolimail](https://github.com/jdrouet/jolimail) - A web application to build [MRML](https://github.com/jdrouet/mrml) templates. Stars:`133`.
* [gsquire/sendgrid-rs](https://github.com/gsquire/sendgrid-rs) - Library for SendGrid API Stars:`103`.
* [stalwartlabs/mail-auth](https://github.com/stalwartlabs/mail-auth) [[mail-auth](https://crates.io/crates/mail-auth)] - DKIM, ARC, SPF and DMARC message authentication library [![build badge](https://github.com/stalwartlabs/mail-auth/actions/workflows/rust.yml/badge.svg)](https://github.com/stalwartlabs/mail-auth/actions/workflows/rust.yml) Stars:`78`.
* [duesee/imap-codec](https://github.com/duesee/imap-codec) [[imap-codec](https://crates.io/crates/imap-codec)] - Rock-solid and complete codec for IMAP [![Build & Test](https://github.com/duesee/imap-codec/actions/workflows/build_and_test.yml/badge.svg)](https://github.com/duesee/imap-codec/actions/workflows/build_and_test.yml) Stars:`35`.

### Encoding

[[encoding](https://crates.io/keywords/encoding)]

* ASN.1
  * [alex/rust-asn1](https://github.com/alex/rust-asn1) - ASN.1 (DER) serializer Stars:`97`.
* Binary
  * [bincode-org/bincode](https://github.com/bincode-org/bincode) - A binary encoder/decoder [![CI](https://github.com/bincode-org/bincode/actions/workflows/rust.yml/badge.svg?branch=trunk)](https://github.com/bincode-org/bincode/actions/workflows/rust.yml) Stars:`2.6K`.
  * [m4b/goblin](https://github.com/m4b/goblin) [[goblin](https://crates.io/crates/goblin)] - cross-platform, zero-copy, and endian-aware binary parsing Stars:`1.2K`.
  * [jamesmunns/postcard](https://github.com/jamesmunns/postcard) [[postcard](https://crates.io/crates/postcard)] - Postcard is a #![no_std] focused serializer and deserializer for Serde. Stars:`783`.
* BSON
  * [mongodb/bson-rust](https://github.com/mongodb/bson-rust) - Encoding and decoding support for BSON Stars:`389`.
* Byte swapping
  * [BurntSushi/byteorder](https://github.com/BurntSushi/byteorder) - Supports big-endian, little-endian and native byte orders Stars:`951`.
* Cap'n Proto
  * [capnproto/capnproto-rust](https://github.com/capnproto/capnproto-rust) - Cap'n Proto is a type system for distributed systems Stars:`2.0K`.
* CBOR
  * [serde_cbor](https://crates.io/crates/serde_cbor) - CBOR support for serde
* Character Encoding
  * [hsivonen/encoding_rs](https://github.com/hsivonen/encoding_rs) [[encoding_rs](https://crates.io/crates/encoding_rs)] - A Gecko-oriented implementation of the Encoding Standard Stars:`369`.
  * [lifthrasiir/rust-encoding](https://github.com/lifthrasiir/rust-encoding) - Character encoding support for Rust. (also known as rust-encoding) It is based on WHATWG Encoding Standard, and also provides an advanced interface for error detection and recovery. Stars:`284`.
* CRC
  * [mrhooray/crc-rs](https://github.com/mrhooray/crc-rs) - Rust implementation of CRC(16, 32, 64) with support of various standards Stars:`182`.
* CSV
  * [BurntSushi/rust-csv](https://github.com/BurntSushi/rust-csv) - A fast and flexible CSV reader and writer, with support for Serde Stars:`1.6K`.
* EDN
  * [frol/flatc-rust](https://github.com/frol/flatc-rust) - FlatBuffers compiler (flatc) integration for Cargo build scripts Stars:`105`.
  * [edn-rs](https://github.com/naomijub/edn-rs) [[edn-rs](https://crates.io/crates/edn-rs)] - crate to parse and emit EDN format into Rust types. Stars:`81`.
* [FlatBuffers](https://flatbuffers.dev/)
* HAR
  * [mandrean/har-rs](https://github.com/mandrean/har-rs) [[har](https://crates.io/crates/har)] - A HTTP Archive Format (HAR) serialization & deserialization library Stars:`40`.
* HTML
  * [servo/html5ever](https://github.com/servo/html5ever) - High-performance browser-grade HTML5 parser Stars:`2.1K`.
* JSON
  * [serde-rs/json](https://github.com/serde-rs/json) [[serde\_json](https://crates.io/crates/serde_json)] - JSON support for [Serde](https://github.com/serde-rs/serde) framework Stars:`4.7K`.
  * [simd-lite/simd-json](https://github.com/simd-lite/simd-json) [[simd-json](https://crates.io/crates/simd-json)] - High performance JSON parser based on a port of simdjson Stars:`1.1K`.
  * [pikkr/pikkr](https://github.com/pikkr/pikkr) [[pikkr](https://crates.io/crates/pikkr)] - JSON parser which picks up values directly without performing tokenization Stars:`631`.
  * [maciejhirsz/json-rust](https://github.com/maciejhirsz/json-rust) [[json](https://crates.io/crates/json)] - JSON implementation Stars:`563`.
  * [importcjj/rust-ajson](https://github.com/importcjj/rust-ajson) [[ajson](https://crates.io/crates/ajson)] - Get JSON values quickly Stars:`104`.
* MsgPack
  * [3Hren/msgpack-rust](https://github.com/3Hren/msgpack-rust) - Low/high level MessagePack implementation Stars:`1.1K`.
* NetCDF
  * [georust/netcdf](https://github.com/georust/netcdf) [[netcdf](https://crates.io/crates/netcdf)] - Medium-level netCDF bindings, allowing easy reading and writing of array-like structures to a file. Stars:`80`.
* PEM
  * [jcreekmore/pem-rs](https://github.com/jcreekmore/pem-rs) [[pem](https://crates.io/crates/pem)] - Parse and encode PEM-encoded data Stars:`52`.
* ProtocolBuffers
  * [tokio-rs/prost](https://github.com/tokio-rs/prost) - [![continuous integration](https://github.com/tokio-rs/prost/workflows/continuous%20integration/badge.svg?branch=master)](https://github.com/tokio-rs/prost/actions) Stars:`3.7K`.
  * [stepancheg/rust-protobuf](https://github.com/stepancheg/rust-protobuf) - Rust implementation of Google protocol buffers Stars:`2.7K`.
* rkyv
  * [rkyv/rkyv](https://github.com/rkyv/rkyv) [[rkyv](https://crates.io/crates/rkyv)] - rkyv (archive) is a zero-copy deserialization framework Stars:`2.7K`.
* RON (Rusty Object Notation)
  * [https://github.com/ron-rs/ron](https://github.com/ron-rs/ron) - Rusty Object Notation Stars:`3.2K`.
* Serde
  * [iddm/serde-aux](https://github.com/iddm/serde-aux/) - additional tools for using with the serde library. [![CI](https://github.com/iddm/serde-aux/actions/workflows/ci.yml/badge.svg)](https://github.com/iddm/serde-aux/actions/workflows/ci.yml) [![Crates badge](https://img.shields.io/crates/v/serde-aux.svg)](https://crates.io/crates/serde-aux) Stars:`145`.
* TOML
  * [tamasfe/taplo](https://github.com/tamasfe/taplo) [[taplo](https://crates.io/crates/taplo)] - A TOML toolkit [![CI](https://github.com/tamasfe/taplo/workflows/Continuous%20integration/badge.svg)](https://github.com/tamasfe/taplo/actions?query=workflow%3A%22Continuous+integration%22) Stars:`1.3K`.
  * [toml-rs/toml](https://github.com/toml-rs/toml) - [![CI](https://github.com/toml-rs/toml/actions/workflows/ci.yml/badge.svg)](https://github.com/toml-rs/toml/actions/workflows/ci.yml) Stars:`658`.
* XML
  * [tafia/quick-xml](https://github.com/tafia/quick-xml) - High performance XML pull reader/writer Stars:`1.1K`.
  * [netvl/xml-rs](https://github.com/netvl/xml-rs) - A streaming XML library Stars:`459`.
  * [media-io/yaserde](https://github.com/media-io/yaserde) - Yet Another Serializer/Deserializer specialized for XML Stars:`174`.
  * [shepmaster/sxd-document](https://github.com/shepmaster/sxd-document) - An XML library Stars:`155`.
  * [shepmaster/sxd-xpath](https://github.com/shepmaster/sxd-xpath) - An XPath library Stars:`121`.
  * [Florob/RustyXML](https://github.com/Florob/RustyXML) - an XML parser Stars:`103`.
* YAML
  * [dtolnay/serde-yaml](https://github.com/dtolnay/serde-yaml) [[serde\_yaml](https://crates.io/crates/serde_yaml)] - YAML support for [Serde](https://github.com/serde-rs/serde) framework [![build](https://img.shields.io/github/workflow/status/dtolnay/serde-yaml/CI/master)](https://github.com/dtolnay/serde-yaml/actions?query=branch%3Amaster) Stars:`949`.
  * [chyh1990/yaml-rust](https://github.com/chyh1990/yaml-rust) - The missing YAML 1.2 implementation. Stars:`598`.
  * [vitiral/stfu8](https://github.com/vitiral/stfu8) [[stfu8](https://crates.io/crates/stfu8)] - Sorta Text Format in UTF-8 Stars:`25`.

### Filesystem

[[filesystem](https://crates.io/keywords/filesystem)]
* Operations
  * [OpenDAL](https://github.com/apache/opendal) [[opendal](https://crates.io/crates/opendal)] - A unified data access layer, empowering users to seamlessly and efficiently retrieve data from diverse storage services. [![build](https://img.shields.io/github/actions/workflow/status/apache/opendal/ci_core.yml?branch=main)](https://github.com/apache/opendal/actions?query=branch%3Amain) Stars:`3.0K`.
  * [Camino](https://github.com/camino-rs/camino) [[camino](https://crates.io/crates/camino)] - Like Rust's std::path::Path, but UTF-8. Stars:`406`.
  * [webdesus/fs_extra](https://github.com/webdesus/fs_extra) - expanding opportunities standard library std::fs and std::io Stars:`241`.
  * [ParthJadhav/Rust_Search](https://github.com/ParthJadhav/Rust_Search) [[rust_search](https://crates.io/crates/rust_search)] - Blazingly fast file search library. Stars:`134`.
  * [vitiral/path_abs](https://github.com/vitiral/path_abs) [[path_abs](https://crates.io/crates/path_abs)] - Absolute serializable path types and associated methods. Stars:`49`.
  * [pop-os/sys-mount](https://github.com/pop-os/sys-mount) [[sys-mount](https://crates.io/crates/sys-mount)] - High level abstraction for the `mount` / `umount2` system calls. Stars:`41`.
  * [pop-os/dbus-udisks2](https://github.com/pop-os/dbus-udisks2) [[dbus-udisks2](https://crates.io/crates/dbus-udisks2)] - UDisks2 DBus API Stars:`18`.
* Temporary Files
  * [zboxfs/zbox](https://github.com/zboxfs/zbox) [[zbox](https://crates.io/crates/zbox)] - Zero-details, privacy-focused embeddable file system. Stars:`1.5K`.
  * [Stebalien/tempfile](https://github.com/Stebalien/tempfile) - temporary file library Stars:`1.1K`.
  * [Stebalien/xattr](https://github.com/Stebalien/xattr) [[xattr](https://crates.io/crates/xattr)] - list and manipulate unix extended file attributes Stars:`57`.

### Finance

* [avhz/RustQuant](https://github.com/avhz/RustQuant) [[RustQuant](https://crates.io/crates/RustQuant)] - A quantitative finance library. ![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/avhz/RustQuant/build.yml) Stars:`952`.
* [d-e-s-o/apca](https://github.com/d-e-s-o/apca) [[apca](https://crates.io/crates/apca)] - Opinionated and comprehensive bindings to the [Alpaca API](https://alpaca.markets/) for stock trading and more. ![GitHub Workflow Status](https://github.com/d-e-s-o/apca/actions/workflows/test.yml/badge.svg?branch=main) Stars:`130`.

### Functional Programming

[[functional programming](https://crates.io/keywords/fp)]
* Prelude
  * [JasonShin/fp-core.rs](https://github.com/JasonShin/fp-core.rs) - A library for functional programming Stars:`1.3K`.
  * [myrrlyn/tap](https://github.com/myrrlyn/tap) - Suffix-Position Pipeline Behavior Stars:`387`.

### Game development

See also [Are we game yet?](https://arewegameyet.rs)
* Allegro
* [Awesome wgpu](https://github.com/rofrol/awesome-wgpu) - A curated list of wgpu code and resources Stars:`404`.
* [Awesome Quads](https://github.com/ozkriff/awesome-quads) - A curated list of links to miniquad/macroquad-related code & resources Stars:`173`.
  * [SiegeLord/RustAllegro](https://github.com/SiegeLord/RustAllegro) - [Allegro 5](https://liballeg.org/) bindings Stars:`95`.
* bracket-lib (previously RLTK)
  * [bracket-lib](https://github.com/amethyst/bracket-lib) [[bracket-lib](https://crates.io/crates/bracket-lib)] - The Roguelike Toolkit (RLTK). [![Rust](https://github.com/amethyst/bracket-lib/actions/workflows/rust.yml/badge.svg)](https://github.com/amethyst/bracket-lib/actions/workflows/rust.yml) Stars:`1.5K`.
* Challonge
  * [iddm/challonge-rs](https://github.com/iddm/challonge-rs) [[challonge](https://crates.io/crates/challonge)] - Client library for the Challonge REST API. Helps to organize tournaments. [![CI](https://github.com/iddm/challonge-rs/actions/workflows/ci.yml/badge.svg)](https://github.com/iddm/challonge-rs/actions/workflows/ci.yml) Stars:`2`.
* Entity-Component Systems (ECS)
  * [amethyst/specs](https://github.com/amethyst/specs) - Specs Parallel ECS Stars:`2.4K`.
  * [legion](https://github.com/amethyst/legion) - A feature rich high performance ECS library with minimal boilerplate [![build badge](https://github.com/amethyst/legion/workflows/CI/badge.svg?branch=master)](https://github.com/amethyst/legion/actions) Stars:`1.6K`.
* Game Engines
  * [Bevy](https://github.com/bevyengine/bevy) - is a refreshingly simple data-driven game engine. - [![Crates.io](https://img.shields.io/crates/v/bevy.svg)](https://crates.io/crates/bevy) [![Crates.io](https://img.shields.io/crates/d/bevy.svg)](https://crates.io/crates/bevy) Stars:`33.9K`.
  * [Fyrox](https://fyrox.rs/) - Game engine 3D [![Crates.io](https://img.shields.io/crates/v/fyrox.svg)](https://crates.io/crates/fyrox) [![license](https://img.shields.io/crates/l/fyrox.svg)](https://github.com/FyroxEngine/Fyrox/blob/master/LICENSE.md) [![Crates.io](https://img.shields.io/crates/d/fyrox.svg)](https://crates.io/crates/fyrox) Stars:`7.5K`.
  * [Piston](https://www.piston.rs/) - [![Crates.io](https://img.shields.io/crates/v/piston.svg?style=flat-square)](https://crates.io/crates/piston) [![Crates.io](https://img.shields.io/crates/l/piston.svg)](https://github.com/PistonDevelopers/piston/blob/master/LICENSE) [![Crates.io](https://img.shields.io/crates/d/piston.svg)](https://crates.io/crates/piston) Stars:`4.6K`.
  * [ggez](https://github.com/ggez/ggez) - A lightweight game framework for making 2D games with minimum friction - [![Crates.io](https://img.shields.io/crates/v/ggez.svg)](https://crates.io/crates/ggez) [![license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/ggez/ggez/blob/master/LICENSE) [![Crates.io](https://img.shields.io/crates/d/ggez.svg)](https://crates.io/crates/ggez) Stars:`4.2K`.
  * [Unrust](https://github.com/unrust/unrust) - Webgl 2.0 / native game engine Stars:`374`.
  * [oxidator](https://github.com/Ruddle/oxidator) - A real time strategy game/engine supporting WebGPU Stars:`292`.
  * [Kiss3d](http://kiss3d.org) - A Keep It Simple, Stupid 3d graphics engine [![Crates.io](https://img.shields.io/crates/d/kiss3d.svg)](https://crates.io/crates/kiss3d)
* Game Servers
  * [godot-rust/gdnative](https://github.com/godot-rust/gdnative) [[gdnative](https://crates.io/crates/gdnative)] - Bindings to the Godot game engine [![CI](https://github.com/godot-rust/gdnative/actions/workflows/full-ci.yml/badge.svg)](https://github.com/godot-rust/gdnative/actions/workflows/full-ci.yml) Stars:`3.6K`.
  * [Rust-SDL2/rust-sdl2](https://github.com/Rust-SDL2/rust-sdl2) - SDL2 bindings Stars:`2.7K`.
  * [deltaphc/raylib-rs](https://github.com/deltaphc/raylib-rs) [[raylib](https://crates.io/crates/raylib)] - Bindings for raylib Stars:`717`.
  * [brson/rust-sdl](https://github.com/brson/rust-sdl) - SDL1 bindings Stars:`178`.
  * [gamedig/rust-gamedig](https://github.com/gamedig/rust-gamedig) [[gamedig](https://crates.io/crates/gamedig)] - Query game servers for informations such as name, players online, max players count etc. [![Crates.io](https://img.shields.io/crates/v/gamedig.svg)](https://crates.io/crates/gamedig) [![Crates.io](https://img.shields.io/crates/d/gamedig.svg)](https://crates.io/crates/gamedig) Stars:`36`.
* [Godot](https://godotengine.org/)
* [Raylib](https://www.raylib.com/)
* [SDL](http://www.libsdl.org/) [[sdl](https://crates.io/keywords/sdl)]
* SFML
  * [jeremyletang/rust-sfml](https://github.com/jeremyletang/rust-sfml) - [SFML](https://www.sfml-dev.org/) bindings Stars:`627`.
* Skillratings
  * [atomflunder/skillratings](https://github.com/atomflunder/skillratings) [[skillratings](https://crates.io/crates/skillratings)] - Collection of skill rating algorithms for multiplayer games like Elo, Glicko-2, TrueSkill etc. [![crates.io badge](https://img.shields.io/crates/v/skillratings)](https://crates.io/crates/skillratings) [![CI](https://github.com/atomflunder/skillratings/actions/workflows/ci.yml/badge.svg)](https://github.com/atomflunder/skillratings/actions/workflows/ci.yml) Stars:`37`.
* Tcod-rs
  * [tomassedovic/tcod-rs](https://github.com/tomassedovic/tcod-rs) - Libtcod bindings. Stars:`229`.
  * Warning: Not maintained anymore
* Toornament-rs
  * [iddm/toornament-rs](https://github.com/iddm/toornament-rs) - Toornament.com API bindings. [![CI](https://github.com/iddm/toornament-rs/actions/workflows/ci.yml/badge.svg)](https://github.com/iddm/toornament-rs/actions/workflows/ci.yml) [![Crates badge](https://img.shields.io/crates/v/toornament.svg)](https://crates.io/crates/toornament) Stars:`4`.
* Victorem
  * [VictoremWinbringer/Victorem](https://github.com/VictoremWinbringer/Victorem) [[Victorem](https://crates.io/crates/Victorem)] - Easy UDP Game Server and UDP Client framework for creating simple 2D and 3D online game prototype Stars:`29`.

### Geospatial

[[geo](https://crates.io/keywords/geo), [gis](https://crates.io/keywords/gis)]

* [MapLibre/Martin](https://github.com/maplibre/martin) - Map tile server with PostGIS, MBTiles, PMTiles, and sprites support. [![CI build](https://github.com/maplibre/martin/actions/workflows/ci.yml/badge.svg)](https://github.com/maplibre/martin/actions)[![crates.io version](https://img.shields.io/crates/v/martin.svg)](https://crates.io/crates/martin)[![Book](https://img.shields.io/badge/docs-Book-informational)](https://maplibre.org/martin/) Stars:`2.0K`.
* [rust-reverse-geocoder](https://github.com/gx0r/rrgeo) - A fast, offline reverse geocoder, inspired by [thampiman/reverse-geocoder](https://github.com/thampiman/reverse-geocoder) Stars:`120`.
* [DaveKram/coord_transforms](https://github.com/DaveKram/coord_transforms) [[coord_transforms](https://crates.io/crates/coord_transforms)] - coordinate transformations (2-d, 3-d, and geospatial) Stars:`32`.
* [vlopes11/geomorph](https://github.com/vlopes11/geomorph) [[geomorph](https://crates.io/crates/geomorph)] - conversion between UTM, LatLon and MGRS coordinates Stars:`15`.
* [Georust](https://github.com/georust) - geospatial tools and libraries written

### Graph algorithms

* [petgraph/petgraph](https://github.com/petgraph/petgraph) - Graph data structure library. [![graph CI status](https://github.com/petgraph/petgraph/workflows/Continuous%20integration/badge.svg?branch=master)](https://github.com/petgraph/petgraph/actions/workflows/ci.yml) Stars:`2.8K`.
* [neo4j-labs/graph](https://github.com/neo4j-labs/graph) - A library for high-performant graph algorithms [![graph CI status](https://img.shields.io/github/workflow/status/neo4j-labs/graph/CI/main?label=CI)](https://github.com/neo4j-labs/graph/actions/workflows/rust.yml) Stars:`364`.

### Graphics

[[graphics](https://crates.io/keywords/graphics)]

* Font
* [gfx-rs/wgpu](https://github.com/gfx-rs/wgpu) - Native WebGPU implementation based on gfx-hal. [![build badge](https://github.com/gfx-rs/wgpu/workflows/CI/badge.svg?branch=master)](https://github.com/gfx-rs/wgpu/actions) Stars:`11.5K`.
* [gfx-rs/gfx](https://github.com/gfx-rs/gfx) - A high-performance, bindless graphics API. Stars:`5.3K`.
  * [redox-os/rusttype](https://github.com/redox-os/rusttype) - Alternative to libraries like FreeType Stars:`605`.
  * [RazrFalcon/rustybuzz](https://github.com/RazrFalcon/rustybuzz) - An incremental harfbuzz port Stars:`501`.
* OpenGL [[opengl](https://crates.io/keywords/opengl)]
  * [glium/glium](https://github.com/glium/glium) - safe OpenGL wrapper. Stars:`3.4K`.
  * [brendanzab/gl-rs](https://github.com/brendanzab/gl-rs) - An OpenGL function pointer loader Stars:`678`.
  * [PistonDevelopers/glfw-rs](https://github.com/PistonDevelopers/glfw-rs) - GLFW3 bindings and idiomatic wrapper Stars:`641`.
  * [glutin](https://crates.io/crates/glutin) - Alternative to [GLFW](https://www.glfw.org/)
  * [Kiss3d](http://kiss3d.org) - draw simple geometric figures and play with them with one-liners
* PDF
  * [vulkano](https://github.com/vulkano-rs/vulkano) [[vulkano](https://crates.io/crates/vulkano)] - Safe and rich Rust wrapper around the Vulkan API Stars:`4.4K`.
  * [J-F-Liu/lopdf](https://github.com/J-F-Liu/lopdf) - PDF document manipulation Stars:`1.5K`.
  * [fschutt/printpdf](https://github.com/fschutt/printpdf) - PDF writing library Stars:`779`.
  * [kaj/rust-pdf](https://github.com/kaj/rust-pdf) - Generating PDF files in pure Rust Stars:`145`.
  * [bastibense/libharu_ng](https://github.com/bastibense/libharu_ng) [[libharu_ng](https://crates.io/crates/libharu_ng)] - Easily generate PDFs from your Rust app. Stars:`13`.
* [Vulkan](https://www.vulkan.org/) [[vulkan](https://crates.io/keywords/vulkan)]
  * [erupt](https://gitlab.com/Friz64/erupt) [[erupt](https://crates.io/crates/erupt)] - [![build badge](https://gitlab.com/Friz64/erupt/badges/main/pipeline.svg)](https://gitlab.com/Friz64/erupt/-/pipelines)

### GUI

[[gui](https://crates.io/keywords/gui)]

* [autopilot-rs/autopilot-rs](https://github.com/autopilot-rs/autopilot-rs) - A simple, cross-platform GUI automation library. Stars:`371`.
* Cocoa
* [tauri-apps/tauri](https://github.com/tauri-apps/tauri) - Build smaller, faster, and more secure desktop applications with a web frontend, powered by [WRY](https://github.com/tauri-apps/wry). [![test library](https://img.shields.io/github/workflow/status/tauri-apps/tauri/test%20library?label=test%20library)](https://github.com/tauri-apps/tauri/actions?query=workflow%3A%22test+library%22) Stars:`79.7K`.
* [ImGui](https://github.com/ocornut/imgui) Stars:`57.9K`.
* [iced-rs/iced](https://github.com/iced-rs/iced) [[iced](https://crates.io/crates/iced)] - A cross-platform GUI library, focused on simplicity and type-safety. Inspired by Elm. Stars:`23.4K`.
* [emilk/egui](https://github.com/emilk/egui) - Simple, fast, and highly portable immediate mode GUI library. egui runs on the web, natively, and in your favorite game engine. [![Build Status](https://github.com/emilk/egui/workflows/CI/badge.svg)](https://github.com/emilk/egui/actions?workflow=CI) Stars:`20.8K`.
* [DioxusLabs/dioxus](https://github.com/dioxuslabs/dioxus) - a portable, performant, and ergonomic framework for building cross-platform user interfaces in Rust. ![rust ci](https://github.com/dioxuslabs/dioxus/actions/workflows/main.yml/badge.svg) Stars:`19.4K`.
* [slint-ui/slint](https://github.com/slint-ui/slint) [slint](https://crates.io/crates/slint) - [Slint](https://slint.dev/) is a toolkit to efficiently develop fluid graphical user interfaces for embedded devices and desktop applications. [![Build Status](https://github.com/slint-ui/slint/workflows/CI/badge.svg?branch=master)](https://github.com/slint-ui/slint/actions?query=workflow%3ACI) Stars:`16.1K`.
* [libui](https://github.com/andlabs/libui) Stars:`10.7K`.
* [Nuklear](https://github.com/Immediate-Mode-UI/Nuklear) Stars:`8.9K`.
* [fschutt/azul](https://github.com/fschutt/azul) - A free, functional, IMGUI-oriented GUI framework for rapid development of desktop applications written in Rust, supported by the Mozilla WebRender rendering engine. Stars:`5.8K`.
* [makepad/makepad](https://github.com/makepad/makepad) [[makepad-widgets](https://crates.io/crates/makepad-widgets)] - Makepad is a creative software development platform that compiles to wasm/webGL, osx/metal, windows/dx11 linux/opengl. Stars:`5.0K`.
* [OrbTk](https://github.com/redox-os/orbtk) - The Orbital Widget Toolkit is a multi platform (G)UI toolkit using SDL2 [![Build and test](https://github.com/redox-os/orbtk/workflows/build/badge.svg?branch=develop)](https://github.com/redox-os/orbtk/actions) Stars:`3.8K`.
  * [fzyzcjy/flutter_rust_bridge](https://github.com/fzyzcjy/flutter_rust_bridge) - High-level memory-safe binding generator for Flutter/Dart <-> Rust Stars:`3.7K`.
* [tauri-apps/wry](https://github.com/tauri-apps/wry) - Webview Rendering librarY. Stars:`3.4K`.
* [PistonDevelopers/conrod](https://github.com/PistonDevelopers/conrod/) - An easy-to-use, immediate-mode, 2D GUI library Stars:`3.3K`.
* [xilem](https://github.com/linebender/xilem) - Successor of the data-first UI design toolkit [druid](https://github.com/linebender/druid). Stars:`3.2K`.
  * [imgui-rs](https://github.com/imgui-rs/imgui-rs) - Bindings for ImGui [![Build Status](https://github.com/imgui-rs/imgui-rs/workflows/ci/badge.svg?branch=master)](https://github.com/imgui-rs/imgui-rs/actions) Stars:`2.6K`.
  * [relm](https://github.com/antoyo/relm) - Asynchronous, GTK+-based, GUI library, inspired by Elm Stars:`2.4K`.
  * [flutter-rs](https://github.com/flutter-rs/flutter-rs) - Build flutter desktop app in dart & rust. Stars:`2.1K`.
  * [gtk-rs/gtk4-rs](https://github.com/gtk-rs/gtk4-rs) - GTK4 binding ![CI](https://github.com/gtk-rs/gtk4-rs/workflows/CI/badge.svg) Stars:`1.8K`.
  * [cunarist/rinf](https://github.com/cunarist/rinf) - Rust as your Flutter backend, Flutter as your Rust frontend [![Build Test](https://github.com/cunarist/rinf/actions/workflows/build_test.yaml/badge.svg)](https://github.com/cunarist/rinf/actions/workflows/build_test.yaml?query=branch%3Amain) Stars:`1.7K`.
  * [fltk-rs](https://github.com/fltk-rs/fltk-rs) - FLTK bindings [![Build](https://github.com/fltk-rs/fltk-rs/workflows/Build/badge.svg?branch=master)](https://github.com/fltk-rs/fltk-rs/actions) Stars:`1.6K`.
* [emoon/rust_minifb](https://github.com/emoon/rust_minifb) - minifb is a cross-platform window setup with optional bitmap rendering. It also comes with easy mouse and keyboard input. Primarily designed for prototyping Stars:`980`.
  * [servo/core-foundation-rs](https://github.com/servo/core-foundation-rs) - Rust bindings to Core Foundation and other low level libraries on Mac OS X and iOS Stars:`953`.
  * [rust-native-ui/libui-rs](https://github.com/rust-native-ui/libui-rs) - libui bindings. Stars:`927`.
  * [sciter-sdk/rust-sciter](https://github.com/sciter-sdk/rust-sciter) - Sciter bindings [![build badge](https://ci.appveyor.com/api/projects/status/github/sciter-sdk/rust-sciter?svg=true)](https://ci.appveyor.com/project/sciter-sdk/rust-sciter) Stars:`805`.
* [ivanceras/sauron-native](https://github.com/ivanceras/sauron-native) - A truly native and cross platform GUI library. One unified code can be run as native GUI, Html Web and TUI. Stars:`631`.
  * [woboq/qmetaobject-rs](https://github.com/woboq/qmetaobject-rs) - Integrate Qml and Rust by building the QMetaObject at compile time. Stars:`619`.
  * [cyndis/qmlrs](https://github.com/cyndis/qmlrs) - QtQuick bindings Stars:`436`.
  * [nuklear-rust](https://github.com/snuk182/nuklear-rust) - Bindings for Nuklear Stars:`358`.
  * [Kiss-ui](https://github.com/KISS-UI/kiss-ui) - A simple UI framework built on IUP Stars:`342`.
* [saurvs/nfd-rs](https://github.com/saurvs/nfd-rs) - [nativefiledialog](https://github.com/mlabbe/nativefiledialog) bindings Stars:`155`.
* [rise-ui](https://github.com/rise-ui/rise) - Simple component-based cross-Platform GUI Toolkit for developing beautiful and user-friendly interfaces. Stars:`72`.
  * [rust-qt](https://github.com/rust-qt) - Qt bindings for Rust
* [Sciter](https://sciter.com/)
* [Qt](https://doc.qt.io)
* [IUP](http://webserver2.tecgraf.puc-rio.br/iup/)
* [GTK+](https://www.gtk.org/) [[gtk](https://crates.io/keywords/gtk)]
* [Flutter](https://flutter.dev/)
* [FLTK](https://www.fltk.org/)

### Image processing

* [image-rs/image](https://github.com/image-rs/image) - Basic imaging processing functions and methods for converting to and from image formats Stars:`4.8K`.
* [twistedfall/opencv-rust](https://github.com/twistedfall/opencv-rust) - Bindings for OpenCV Stars:`1.9K`.
* [rust-cv/cv](https://github.com/rust-cv/cv) - Implement computer vision algorithms, abstractions, and systems. `#[no_std]` is supported where possible. ![build badge](https://github.com/rust-cv/cv/workflows/tests/badge.svg) Stars:`794`.
* [image-rs/imageproc](https://github.com/image-rs/imageproc) - An image processing library, based on the `image` library. Stars:`716`.
* [abonander/img_hash](https://github.com/abonander/img_hash) - Perceptual image hashing and comparison for equality and similarity. Stars:`302`.
* [teovoinea/steganography](https://github.com/teovoinea/steganography) [[steganography](https://crates.io/crates/steganography)] - A simple steganography library Stars:`93`.
* [marekm4/dominant_color](https://github.com/marekm4/dominant_color) [[dominant_color](https://crates.io/crates/dominant_color)] - Dominant color extractor ![build badge](https://github.com/marekm4/dominant_color/actions/workflows/rust.yml/badge.svg?branch=master) Stars:`31`.

### Language specification

* [shnewto/bnf](https://github.com/shnewto/bnf) - A library for parsing Backus–Naur form context-free grammars. Stars:`252`.

### Logging

[[log](https://crates.io/keywords/log)]

* [tokio-rs/tracing](https://github.com/tokio-rs/tracing) - An application level tracing framework for async-aware structured logging, error handling, metrics, and more [![Build Status](https://github.com/tokio-rs/tracing/workflows/CI/badge.svg?branch=master)](https://github.com/tokio-rs/tracing/actions?query=workflow%3ACI) Stars:`5.1K`.
* [rust-lang/log](https://github.com/rust-lang/log) - Logging implementation Stars:`2.1K`.
* [slog-rs/slog](https://github.com/slog-rs/slog) - Structured, composable logging Stars:`1.6K`.
* [estk/log4rs](https://github.com/estk/log4rs) - highly configurable logging framework modeled after Java's Logback and log4j libraries [![CircleCI](https://circleci.com/gh/estk/log4rs.svg?style=shield)](https://app.circleci.com/pipelines/github/estk/log4rs) Stars:`969`.
* [seanmonstar/pretty-env-logger](https://github.com/seanmonstar/pretty-env-logger) - A pretty, easy-to-use logger. Stars:`474`.
* [rbatis/fast_log](https://github.com/rbatis/fast_log) - Async log High-performance asynchronous logging Stars:`221`.

### Macro

* cute
  * [mattgathu/cute](https://github.com/mattgathu/cute) - Macro for Python-esque list comprehensions. Stars:`336`.
* [Linq-in-Rust](https://github.com/StardustDL/Linq-in-Rust) - Macro and methods for C#-LINQ-like expressions. [![CI](https://github.com/StardustDL/Linq-in-Rust/workflows/CI/badge.svg?branch=master)](https://github.com/StardustDL/Linq-in-Rust/actions?query=workflow%3ACI) Stars:`125`.

### Markup language

* CommonMark
  * [pulldown-cmark/pulldown-cmark](https://github.com/pulldown-cmark/pulldown-cmark) - [CommonMark](https://commonmark.org/) parser Stars:`2.0K`.
* [insomnimus/tidier](https://github.com/insomnimus/tidier) [[tidier](https://crates.io/crates/tidier)] - A library to format HTML, XHTML and XML documents. [![build badge](https://github.com/insomnimus/tidier/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/insomnimus/tidier/actions) Stars:`3`.

### Mobile

* Android / iOS
  * [ivnsch/rust_android_ios](https://github.com/ivnsch/rust_android_ios) - An example of using a shared lib for Android and iOS using rust-swig and cbindgen respectively. Stars:`238`.
* Generic
  * [redbadger/crux](https://github.com/redbadger/crux) [[crux_core](https://crates.io/crates/crux_core)] - Cross-platform app development. Crux helps you share your app's business logic and behavior across mobile (iOS/Android) and web - as a single reusable core. [![Build status](https://img.shields.io/github/actions/workflow/status/redbadger/crux/build.yaml)](https://github.com/redbadger/crux/actions) Stars:`1.6K`.
  * [Geal/rust_on_mobile](https://github.com/Geal/rust_on_mobile) - iOS CocoaPods / Android JNI Stars:`170`.
* iOS
  * [TimNN/cargo-lipo](https://github.com/TimNN/cargo-lipo) - A cargo lipo subcommand which automatically creates a universal library for use with your iOS application. Stars:`524`.

### Network programming

* Bluetooth
  * [bluez/bluer](https://github.com/bluez/bluer) [[bluer](https://crates.io/crates/bluer)] - Official BlueZ bindings. [![build badge](https://github.com/bluez/bluer/actions/workflows/rust.yml/badge.svg?branch=master)](https://github.com/bluez/bluer/actions/workflows/rust.yml) Stars:`272`.
* CoAP
  * [Covertness/coap-rs](https://github.com/Covertness/coap-rs) - A [Constrained Application Protocol(CoAP)](https://datatracker.ietf.org/doc/html/rfc7252) library. Stars:`202`.
* Docker
  * [fussybeaver/bollard](https://github.com/fussybeaver/bollard) - Docker daemon API Stars:`782`.
* FTP
  * [mattnenterprise/rust-ftp](https://github.com/mattnenterprise/rust-ftp) - an [FTP](https://en.wikipedia.org/wiki/File_Transfer_Protocol) client Stars:`179`.
* gRPC
  * [hyperium/tonic](https://github.com/hyperium/tonic) - A native gRPC client & server implementation with async/await support [![Crates.io](https://img.shields.io/crates/v/tonic)](https://crates.io/crates/tonic) Stars:`9.3K`.
  * [tikv/grpc-rs](https://github.com/tikv/grpc-rs) - The gRPC library built on C Core library and futures Stars:`1.8K`.
* HTTP
  * [Hurl](https://github.com/Orange-OpenSource/hurl) - Run and test HTTP requests with plain text and libcurl [![CI](https://github.com/Orange-OpenSource/hurl/workflows/CI/badge.svg)](https://github.com/Orange-OpenSource/hurl/actions) Stars:`12.3K`.
* IPNetwork
  * [candrew/netsim](https://github.com/canndrew/netsim) - A library for network simulation and testing Stars:`135`.
  * [achanda/ipnetwork](https://github.com/achanda/ipnetwork) - A library to work with IP networks Stars:`120`.
* Low level
  * [actix/actix](https://github.com/actix/actix) - Actor library Stars:`8.5K`.
  * [smoltcp-rs/smoltcp](https://github.com/smoltcp-rs/smoltcp) - A standalone, event-driven TCP/IP stack that is designed for bare-metal, real-time systems Stars:`3.6K`.
  * [libpnet/libpnet](https://github.com/libpnet/libpnet) - A cross-platform, low level networking Stars:`2.2K`.
  * [dylanmckay/protocol](https://github.com/dylanmckay/protocol) - Custom TCP/UDP protocol definitions Stars:`185`.
* message-io
  * [lemunozm/message-io](https://github.com/lemunozm/message-io) - Event-driven message library to build network applications easy and fast. Supports TCP, UDP and WebSockets. [![build badge](https://img.shields.io/github/workflow/status/lemunozm/message-io/message-io%20ci)](https://github.com/lemunozm/message-io/actions?query=workflow%3A%22message-io+ci%22) Stars:`1.1K`.
* MQTT
  * [bytebeamio/rumqtt](https://github.com/bytebeamio/rumqtt) - A library for developers to build applications that communicate with the [MQTT protocol](https://mqtt.org) over TCP and WebSockets, with or without TLS. [![Build and Test](https://github.com/bytebeamio/rumqtt/actions/workflows/build.yml/badge.svg)](https://github.com/bytebeamio/rumqtt/actions/workflows/build.yml) Stars:`1.5K`.
* NanoMsg
  * [thehydroimpulse/nanomsg.rs](https://github.com/thehydroimpulse/nanomsg.rs) - [nanomsg](https://nanomsg.org/) bindings Stars:`387`.
* NATS
  * [nats-io/nats.rs](https://github.com/nats-io/nats.rs) - Client for NATS, the cloud native messaging system. [![Build Status](https://github.com/nats-io/nats.rs/workflows/Rust/badge.svg?branch=master)](https://github.com/nats-io/nats.rs/actions) Stars:`983`.
* Nng
  * [neachdainn/nng-rs](https://gitlab.com/neachdainn/nng-rs) [[Nng](https://crates.io/crates/nng)] - [Nng (nanomsg v2)](https://nng.nanomsg.org/index.html) bindings [![build badge](https://gitlab.com/neachdainn/nng-rs/badges/master/pipeline.svg)](https://gitlab.com/neachdainn/nng-rs/-/pipelines)
* NNTP
  * [mattnenterprise/rust-nntp](https://github.com/mattnenterprise/rust-nntp) [[nntp](https://crates.io/crates/nntp)] - an [NNTP](https://en.wikipedia.org/wiki/Network_News_Transfer_Protocol) client Stars:`17`.
* P2P
  * [libp2p/rust-libp2p](https://github.com/libp2p/rust-libp2p) - Implementation of libp2p networking stack. [![Circle CI](https://circleci.com/gh/libp2p/rust-libp2p.svg?style=svg)](https://app.circleci.com/pipelines/github/libp2p/rust-libp2p) Stars:`4.3K`.
* POP3
  * [mattnenterprise/rust-pop3](https://github.com/mattnenterprise/rust-pop3) [[pop3](https://crates.io/crates/pop3)] - A [POP3](https://en.wikipedia.org/wiki/Post_Office_Protocol) client Stars:`31`.
* QUIC
  * [cloudflare/quiche](https://github.com/cloudflare/quiche) - cloudflare implementation of the QUIC transport protocol and HTTP/3 ![build](https://img.shields.io/github/actions/workflow/status/cloudflare/quiche/stable.yml?branch=master) Stars:`9.1K`.
  * [quinn-rs/quinn](https://github.com/quinn-rs/quinn) - Futures-based QUIC implementation [![build badge](https://dev.azure.com/dochtman/Projects/_apis/build/status/Quinn?branchName=master)](https://dev.azure.com/dochtman/Projects/_build) Stars:`3.6K`.
  * [mozilla/neqo](https://github.com/mozilla/neqo) - an Implementation of QUIC Stars:`1.8K`.
  * [aws/s2n-quic](https://github.com/aws/s2n-quic) - An implementation of the IETF QUIC protocol ![ci](https://img.shields.io/github/actions/workflow/status/aws/s2n-quic/ci.yml?branch=main) Stars:`1.1K`.
  * [tencent/tquic](https://github.com/Tencent/tquic) - A high-performance, lightweight, and cross-platform QUIC library [![Build Status](https://img.shields.io/github/actions/workflow/status/tencent/tquic/rust.yml)](https://github.com/Tencent/tquic/actions/workflows/rust.yml) Stars:`946`.
* Raknet
  * [b23r0/rust-raknet](https://github.com/b23r0/rust-raknet) - RakNet Protocol implementation [![Build Status](https://img.shields.io/github/workflow/status/b23r0/rust-raknet/Rust)](https://github.com/b23r0/rust-raknet/actions/workflows/rust.yml) Stars:`217`.
* RPC
  * [ENQT-GmbH/remoc](https://github.com/ENQT-GmbH/remoc) [[remoc](https://crates.io/crates/remoc)] - Remoc provides channels (broadcast, mpsc, oneshot, watch) similar to Tokio's and trait calling over any remote transport. [![build badge](https://github.com/ENQT-GmbH/remoc/actions/workflows/rust.yml/badge.svg?branch=master)](https://github.com/ENQT-GmbH/remoc/actions/workflows/rust.yml) Stars:`161`.
  * [smallnest/rpcx-rs](https://github.com/smallnest/rpcx-rs) - A RPC library for developing microservices in easy and simple way. Stars:`125`.
* Socket.io
  * [1c3t3a/rust-socketio](https://github.com/1c3t3a/rust-socketio) [[rust_socketio](https://crates.io/crates/rust_socketio)] - an implementation of a [socket.io](https://socket.io) client written in Rust. [![build badge](https://github.com/1c3t3a/rust-socketio/actions/workflows/build.yml/badge.svg)](https://github.com/1c3t3a/rust-socketio/actions/workflows/build.yml) Stars:`389`.
* SSH
  * [alexcrichton/ssh2-rs](https://github.com/alexcrichton/ssh2-rs) - [libssh2](https://libssh2.org/) bindings Stars:`467`.
  * [Thrussh](https://pijul.org/thrussh) [[thrussh](https://crates.io/crates/thrussh)] - an SSH library, backed by [libsodium](https://doc.libsodium.org/)
* Stomp
  * [zslayton/stomp-rs](https://github.com/zslayton/stomp-rs) - A [STOMP 1.2](http://stomp.github.io/stomp-specification-1.2.html) client implementation Stars:`89`.
* VPN
  * [defguard/wireguard-rs](https://github.com/DefGuard/wireguard-rs) - A multi-platform library providing a unified high-level API for managing WireGuard interfaces using native OS kernel and userspace WireGuard protocol implementations Stars:`124`.
* ZeroMQ
  * [erickt/rust-zmq](https://github.com/erickt/rust-zmq) - [ZeroMQ](https://zeromq.org/) bindings Stars:`888`.

### Parsing

  * [tree-sitter/tree-sitter](https://github.com/tree-sitter/tree-sitter) - A parser generator tool and an incremental parsing library geared towards programming tools Stars:`17.2K`.
  * [rust-bakery/nom](https://github.com/rust-bakery/nom) - parser combinator library Stars:`9.2K`.
  * [pest-parser/pest](https://github.com/pest-parser/pest) - The Elegant Parser Stars:`4.5K`.
  * [lalrpop/lalrpop](https://github.com/lalrpop/lalrpop) - LR(1) parser generator Stars:`2.9K`.
  * [kevinmehall/rust-peg](https://github.com/kevinmehall/rust-peg) - Parsing Expression Grammar (PEG) parser generator Stars:`1.4K`.
  * [Marwes/combine](https://github.com/Marwes/combine) - parser combinator library Stars:`1.3K`.
  * [softdevteam/grmtools](https://github.com/softdevteam/grmtools/) - A LR parser with better error correction Stars:`494`.
  * [Eliah-Lakhin/lady-deirdre](https://github.com/Eliah-Lakhin/lady-deirdre) - A framework for new programming languages and LSP servers. Stars:`358`.
  * [m4rw3r/chomp](https://github.com/m4rw3r/chomp) - A fast monadic-style parser combinator Stars:`242`.
  * [igumnoff/shiva](https://github.com/igumnoff/shiva) - Shiva library: Implementation in Rust of a parser and generator for documents of any type (Plain text, Markdown, HTML, PDF and etc) Stars:`142`.
  * [ptal/oak](https://github.com/ptal/oak) - A typed PEG parser generator (compiler plugin) Stars:`141`.
  * [freestrings/jsonpath](https://github.com/freestrings/jsonpath) - [JsonPath](https://goessner.net/articles/JsonPath/) engine. Webassembly and Javascript support too Stars:`122`.
  * [comex/rust-shlex](https://github.com/comex/rust-shlex) [[shlex](https://crates.io/crates/shlex)] - Split a string into shell words, like Python's shlex. [![build badge](https://github.com/comex/rust-shlex/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/comex/rust-shlex/actions/workflows/test.yml) Stars:`97`.
  * [Folyd/robotstxt](https://github.com/Folyd/robotstxt) - Port of Google's robots.txt parser and matcher C++ library Stars:`85`.
  * [s-panferov/queryst](https://github.com/s-panferov/queryst) - A query string parsing library inspired by [gs](https://github.com/ljharb/qs#readme) Stars:`71`.
  * [nrc/zero](https://github.com/nrc/zero) [[zero](https://crates.io/crates/zero/)] - zero-allocation parsing of binary data Stars:`47`.
  * [replicadse/wavefront_rs](https://github.com/replicadse/wavefront_rs) - A parser for the Wavefront OBJ format. [![crates.io](https://img.shields.io/crates/v/wavefront_rs.svg)](https://crates.io/crates/wavefront_rs) [![crates.io](https://img.shields.io/crates/d/wavefront_rs?label=crates.io%20downloads)](https://crates.io/crates/wavefront_rs) [![build badge](https://github.com/replicadse/wavefront_rs/workflows/pipeline/badge.svg?branch=master)](https://github.com/replicadse/wavefront_rs/actions) Stars:`4`.
  * [hmeyer/stl_io](https://crates.io/crates/stl_io) - A parser for STL (STereoLithography) files

### Peripherals

* Fingerprint reader
  * [alvaroparker/libfprint-rs](https://github.com/alvaroparker/libfprint-rs) [[libfprint-rs](https://crates.io/crates/libfprint-rs)] - Libfprint-rs provides a wrapper around the Linux libfprint library. Stars:`11`.
* Serial Port
  * [serialport/serialport-rs](https://github.com/serialport/serialport-rs) [[serialport](https://crates.io/crates/serialport)] - A cross-platform library that provides access to a serial port Stars:`433`.

### Platform specific

* Cross-platform
  * [iddm/thread-priority](https://github.com/iddm/thread-priority/) - Simple, crossplatform thread priority management. [![CI](https://github.com/iddm/thread-priority/actions/workflows/ci.yml/badge.svg)](https://github.com/iddm/thread-priority/actions/workflows/ci.yml) [![Crates badge](https://img.shields.io/crates/v/thread-priority.svg)](https://crates.io/crates/thread-priority) Stars:`107`.
  * [svartalf/rust-battery](https://crates.io/crates/battery) - Cross-platform information about the notebook batteries
* FreeBSD
  * [fubarnetes/libjail-rs](https://github.com/fubarnetes/libjail-rs/) [[jail](https://crates.io/crates/jail)] - FreeBSD jail library Stars:`49`.
* Linux
  * [hannobraun/inotify-rs](https://github.com/hannobraun/inotify-rs) - [inotify](https://en.wikipedia.org/wiki/Inotify) bindings [![Rust](https://github.com/hannobraun/inotify-rs/actions/workflows/rust.yml/badge.svg)](https://github.com/hannobraun/inotify-rs/actions/workflows/rust.yml) Stars:`252`.
  * [pop-os/distinst](https://github.com/pop-os/distinst/) - Linux distribution installer Stars:`221`.
  * [yaa110/rust-iptables](https://github.com/yaa110/rust-iptables) [[iptables](https://crates.io/crates/iptables)] - [iptables](https://www.netfilter.org/projects/iptables/index.html) bindings Stars:`81`.
* Unix-like
  * [nix-rust/nix](https://github.com/nix-rust/nix) - Unix-like API bindings [![Cirrus Build Status](https://api.cirrus-ci.com/github/nix-rust/nix.svg)](https://cirrus-ci.com/github/nix-rust/nix) Stars:`2.6K`.
  * [rustix](https://github.com/bytecodealliance/rustix) - Safe bindings to POSIX/Unix/Linux/Winsock2 syscalls [![Actions Status](https://github.com/bytecodealliance/rustix/workflows/CI/badge.svg)](https://github.com/bytecodealliance/rustix/actions?query=workflow%3ACI) Stars:`1.4K`.
  * [zargony/fuse-rs](https://github.com/zargony/fuse-rs) - [FUSE](https://github.com/libfuse/libfuse) bindings Stars:`1.0K`.
* Windows
  * [microsoft/windows-rs](https://github.com/microsoft/windows-rs) - Rust for Windows [![Actions Status](https://github.com/microsoft/windows-rs/workflows/CI/badge.svg)](https://github.com/microsoft/windows-rs/actions) Stars:`10.1K`.
  * [retep998/winapi-rs](https://github.com/retep998/winapi-rs) - Windows API bindings [![Rust](https://github.com/retep998/winapi-rs/actions/workflows/rust.yml/badge.svg?branch=dev)](https://github.com/retep998/winapi-rs/actions/workflows/rust.yml) Stars:`1.8K`.

### Scripting

[[scripting](https://crates.io/keywords/scripting)]

* [rhaiscript/rhai](https://github.com/rhaiscript/rhai) - A tiny and fast embedded scripting language resembling a combination of JavaScript and Rust [![build badge](https://github.com/rhaiscript/rhai/workflows/Build/badge.svg)](https://github.com/rhaiscript/rhai/actions) Stars:`3.6K`.
* [gluon-lang/gluon](https://github.com/gluon-lang/gluon) - A small, statically-typed, functional programming language Stars:`3.2K`.
* [mun](https://github.com/mun-lang/mun) - A compiled, statically-typed scripting language with first class hot reloading support Stars:`1.8K`.
* [PistonDevelopers/dyon](https://github.com/PistonDevelopers/dyon) - A rusty dynamically typed scripting language Stars:`1.7K`.
* [rune-rs/rune](https://github.com/rune-rs/rune) - An embeddable dynamic programming language Stars:`1.6K`.
* [metacall/core](https://github.com/metacall/core) [[metacall](https://crates.io/crates/metacall)] - Cross-platform Polyglot Runtime which supports NodeJS, JavaScript, TypeScript, Python, Ruby, C#, Wasm, Java, Cobol and more. [![build badge](https://gitlab.com/metacall/core/badges/master/pipeline.svg)](https://gitlab.com/metacall/core) Stars:`1.5K`.
* [kcl](https://github.com/kcl-lang/kcl) - A constraint-based record & functional language mainly used in configuration and policy scenarios. Stars:`1.4K`.
* [murarth/ketos](https://github.com/murarth/ketos) - A Lisp dialect functional programming language serving as a scripting and extension language for rust Stars:`749`.
* [duckscript](https://crates.io/crates/duckscript) - [Simple, extendable and embeddable scripting language.](https://github.com/sagiegurari/duckscript) [![build badge](https://github.com/sagiegurari/duckscript/workflows/CI/badge.svg?branch=master)](https://github.com/sagiegurari/duckscript/actions) Stars:`498`.
* [fleabitdev/gamelisp](https://github.com/fleabitdev/glsp) - A Lisp-like scripting language for game development Stars:`391`.
* [3body-lang](https://github.com/rustq/3body-lang) - The Three Body Language Stars:`169`.

### Simulation

[[simulation](https://crates.io/keywords/simulation)]

* [nyx-space](https://crates.io/crates/nyx-space) - High fidelity, fast, reliable and validated astrodynamical toolkit library, used for spacecraft mission design and orbit determination [![Build Status](https://gitlab.com/nyx-space/nyx/badges/master/pipeline.svg)](https://gitlab.com/nyx-space/nyx/-/pipelines)

### System

* [GuillaumeGomez/sysinfo](https://github.com/GuillaumeGomez/sysinfo) [[sysinfo](https://crates.io/crates/sysinfo)] - Cross-platform library to fetch system information [![build badge](https://github.com/GuillaumeGomez/sysinfo/actions/workflows/CI.yml/badge.svg?branch=master)](https://github.com/GuillaumeGomez/sysinfo/actions/workflows/CI.yml) Stars:`1.9K`.
* [ardaku/whoami](https://github.com/ardaku/whoami) [[whoami](https://crates.io/crates/whoami)] - crate to get the current user and environment. [![build badge](https://github.com/ardaku/whoami/actions/workflows/ci.yml/badge.svg?branch=stable)](https://github.com/ardaku/whoami/actions/workflows/ci.yml) Stars:`172`.
* [Phate6660/nixinfo](https://github.com/Phate6660/nixinfo) [[nixinfo](https://crates.io/crates/nixinfo)] - A lib crate for gathering system info such as cpu, distro, environment, kernel, etc. Stars:`46`.
* [sorairolake/sysexits-rs](https://github.com/sorairolake/sysexits-rs) [[sysexits](https://crates.io/crates/sysexits)] - The system exit codes as defined by [`<sysexits.h>`](https://manpages.ubuntu.com/manpages/lunar/man3/sysexits.h.3head.html). [![CI](https://github.com/sorairolake/sysexits-rs/workflows/CI/badge.svg?branch=develop)](https://github.com/sorairolake/sysexits-rs/actions?query=workflow%3ACI) Stars:`21`.

### Task scheduling

* [delay-timer](https://github.com/BinChengZhao/delay-timer) - Time-manager of delayed tasks. Like crontab, but asynchronous tasks are possible. [![Build](https://github.com/BinChengZhao/delay-timer/actions/workflows/rust.yml/badge.svg)]( https://github.com/BinChengZhao/delay-timer/actions) Stars:`301`.

### Template engine

* Handlebars
  * [sunng87/handlebars-rust](https://github.com/sunng87/handlebars-rust) - Handlebars template engine with inheritance, custom helper support. Stars:`1.2K`.
  * [zzau13/yarte](https://github.com/zzau13/yarte) - Yarte stands for **Y**et **A**nother **R**ust **T**emplate **E**ngine, is the fastest template engine. Stars:`279`.
* HTML
  * [Keats/tera](https://github.com/Keats/tera) - template engine based on Jinja2 and the Django template language. [![Actions Status](https://github.com/Keats/tera/workflows/ci/badge.svg?branch=master)](https://github.com/Keats/tera/actions) Stars:`3.3K`.
  * [djc/askama](https://github.com/djc/askama) - template rendering engine based on Jinja Stars:`3.2K`.
  * [lambda-fairy/maud](https://github.com/lambda-fairy/maud) - compile-time HTML templates Stars:`2.0K`.
  * [kaj/ructe](https://github.com/kaj/ructe) - HTML template system Stars:`443`.
  * [Stebalien/horrorshow-rs](https://github.com/Stebalien/horrorshow-rs) - compile-time HTML templates Stars:`318`.
* Mustache
  * [rustache/rustache](https://github.com/rustache/rustache) - a Rust implementation of the Mustache spec Stars:`210`.

### Text processing

* [rust-lang/regex](https://github.com/rust-lang/regex) - Regular expressions (RE2 style) Stars:`3.4K`.
* [greyblake/whatlang-rs](https://github.com/greyblake/whatlang-rs) - Natural language detection library based on trigrams Stars:`963`.
* [mgeisler/textwrap](https://github.com/mgeisler/textwrap) [[textwrap](https://crates.io/crates/textwrap)] - Word wrap text (with support for hyphenation) Stars:`441`.
* [fancy-regex/fancy-regex](https://github.com/fancy-regex/fancy-regex) [[fancy-regex](https://crates.io/crates/fancy-regex)] - Regular expressions implementation designed to support a relatively rich set of features such as look-around and backtracking. [![crates](https://img.shields.io/crates/v/fancy-regex.svg)](https://crates.io/crates/fancy-regex) [![build badge](https://github.com/fancy-regex/fancy-regex/workflows/ci/badge.svg)](https://github.com/fancy-regex/fancy-regex/actions/workflows/ci.yml) Stars:`402`.
* [BurntSushi/suffix](https://github.com/BurntSushi/suffix) - Linear time suffix array construction (with Unicode support) Stars:`256`.
* [BurntSushi/tabwriter](https://github.com/BurntSushi/tabwriter) - Elastic tab stops (i.e., text column alignment) Stars:`243`.
* [cpc](https://github.com/probablykasper/cpc) - Parses and calculates strings of math with support for units and unit conversion, from `1+2` to `1% of round(1 lightyear / 14!s to km/h)`. Stars:`119`.
* [Daniel-Liu-c0deb0t/triple_accel](https://github.com/Daniel-Liu-c0deb0t/triple_accel) [[triple_accel](https://crates.io/crates/triple_accel)] - Rust edit distance routines accelerated using SIMD; supports fast Hamming, Levenshtein, restricted Damerau-Levenshtein, etc. distance calculations and string search [![build badge](https://github.com/Daniel-Liu-c0deb0t/triple_accel/workflows/Test/badge.svg?branch=master)](https://github.com/Daniel-Liu-c0deb0t/triple_accel/actions) Stars:`98`.
* [null8626/decancer](https://github.com/null8626/decancer) [[decancer](https://crates.io/crates/decancer)] - A tiny package that removes common unicode confusables/homoglyphs from strings. [![crates](https://img.shields.io/crates/v/decancer.svg)](https://crates.io/crates/decancer) [![build badge](https://github.com/null8626/decancer/workflows/CI/badge.svg)](https://github.com/null8626/decancer/actions/workflows/CI.yml) Stars:`97`.
* [Lucretiel/joinery](https://github.com/Lucretiel/joinery) [[joinery](https://crates.io/crates/joinery)] - Generic string + iterable joining Stars:`92`.
* [ps1dr3x/easy_reader](https://github.com/ps1dr3x/easy_reader) - A reader that allows forwards, backwards and random navigations through the lines of huge files without consuming iterators Stars:`85`.
* [becheran/wildmatch](https://github.com/becheran/wildmatch) [[wildmatch](https://crates.io/crates/wildmatch)] - Simple string matching with questionmark- and star-wildcard operator [![Actions Status](https://github.com/becheran/wildmatch/workflows/Build/badge.svg?branch=master)](https://github.com/becheran/wildmatch/actions) Stars:`69`.
* [yaa110/rake-rs](https://github.com/yaa110/rake-rs) [[rake](https://crates.io/crates/rake)] - Multilingual implementation of RAKE algorithm for Rust Stars:`33`.
* [pwoolcoc/ngrams](https://github.com/pwoolcoc/ngrams) [[ngrams](https://crates.io/crates/ngrams)] - Construct [n-grams](https://en.wikipedia.org/wiki/N-gram) from arbitrary iterators Stars:`27`.
* [strsim-rs](https://crates.io/crates/strsim) - String similarity metrics

### Text search

* [meilisearch/MeiliSearch](https://github.com/meilisearch/MeiliSearch) - Ultra relevant, instant and typo-tolerant full-text search API. [![Build Status](https://github.com/meilisearch/MeiliSearch/workflows/Cargo%20test/badge.svg?branch=master)](https://github.com/meilisearch/MeiliSearch/actions) Stars:`45.1K`.
* [tantivy](https://github.com/quickwit-oss/tantivy) [[tantivy](https://crates.io/crates/tantivy)] - A horse-speed full-text search engine library written in Rust. [![Build Status](https://github.com/quickwit-oss/tantivy/actions/workflows/test.yml/badge.svg)](https://github.com/quickwit-oss/tantivy/actions/workflows/test.yml) Stars:`11.3K`.
* [pg_search](https://github.com/paradedb/paradedb/tree/dev/pg_search) - PostgreSQL extension that enables full-text search over SQL tables using the BM25 algorithm, the state-of-the-art ranking function for full-text search. Stars:`4.7K`.
* [BurntSushi/fst](https://github.com/BurntSushi/fst) [[fst](https://crates.io/crates/fst)] - a fast implementation of ordered sets and maps using finite state machines Stars:`1.7K`.
* [andylokandy/simsearch-rs](https://github.com/andylokandy/simsearch-rs) [[simsearch](https://crates.io/crates/simsearch)] - A simple and lightweight fuzzy search engine that works in memory, searching for similar strings Stars:`159`.
* [CurrySoftware/perlin](https://github.com/CurrySoftware/perlin) [[perlin](https://crates.io/crates/perlin)] - A lazy, zero-allocation and data-agnostic Information Retrieval library Stars:`75`.

### Unsafe

* [zerocopy](https://crates.io/crates/zerocopy) - "Zerocopy makes zero-cost memory manipulation effortless. We write `unsafe` so you don’t have to."

### Video

* [ffmpeg-sidecar](https://github.com/nathanbabcock/ffmpeg-sidecar) - Wrap a standalone FFmpeg binary in an intuitive Iterator interface. [![Build Status](https://github.com/nathanbabcock/ffmpeg-sidecar/actions/workflows/rust.yml/badge.svg)](https://github.com/nathanbabcock/ffmpeg-sidecar/actions/workflows/rust.yml) Stars:`192`.

### Virtualization

* [bytecodealliance/wasmtime](https://github.com/bytecodealliance/wasmtime) - A standalone runtime for WebAssembly [![Build Status](https://github.com/bytecodealliance/wasmtime/workflows/CI/badge.svg)](https://github.com/bytecodealliance/wasmtime/actions?query=workflow%3ACI) Stars:`14.8K`.
* [beneills/quantum](https://github.com/beneills/quantum) - Advanced quantum computer simulator Stars:`264`.
* [oxidecomputer/propolis](https://github.com/oxidecomputer/propolis) - Userspace program for illumos bhyve kernel modules Stars:`172`.
* [saurvs/hypervisor-rs](https://github.com/saurvs/hypervisor-rs) - Hardware-accelerated virtualization on OS X Stars:`60`.
* [chromium/chromiumos/platform/crosvm](https://chromium.googlesource.com/chromiumos/platform/crosvm/) - CrOSVM Enables Chrome OS to run Linux apps inside a fast, secure virtualized environment

### Web programming

See also [Are we web yet?](https://www.arewewebyet.org) and [Rust web framework comparison](https://github.com/flosse/rust-web-framework-comparison).

* Client-side / WASM
  * [leptos](https://github.com/leptos-rs/leptos) - Leptos is a full-stack, isomorphic web framework leveraging fine-grained reactivity to build declarative user interfaces.[![crate](https://img.shields.io/crates/v/create-rust-app.svg)](https://crates.io/crates/leptos) Stars:`15.3K`.
  * [seed](https://github.com/seed-rs/seed) - A framework for creating web apps Stars:`3.8K`.
  * [sauron](https://github.com/ivanceras/sauron) - Client side web framework which closely adheres to The Elm Architecture. Stars:`1.9K`.
  * [cargo-web](https://crates.io/crates/cargo-web) - A Cargo subcommand for the client-side Web
  * [stdweb](https://crates.io/crates/stdweb) - A standard library for the client-side Web
  * [yew](https://crates.io/crates/yew) - A framework for making client web apps
* HTTP Client
  * [hyperium/hyper](https://github.com/hyperium/hyper) - an HTTP implementation [![CI](https://github.com/hyperium/hyper/workflows/CI/badge.svg?branch=master)](https://github.com/hyperium/hyper/actions?query=workflow%3ACI) Stars:`14.1K`.
  * [seanmonstar/reqwest](https://github.com/seanmonstar/reqwest) - an ergonomic HTTP Client. Stars:`9.4K`.
  * [ducaale/xh](https://github.com/ducaale/xh) - Friendly and fast tool for sending HTTP requests [![crate](https://img.shields.io/crates/v/create-rust-app.svg)](https://crates.io/crates/xh) [![GitHub actions Status](https://github.com/ducaale/xh/workflows/CI/badge.svg?branch=master)](https://github.com/ducaale/xh/actions) Stars:`5.0K`.
  * [async-graphql](https://github.com/async-graphql/async-graphql) - A GraphQL server library [![Build Status](https://dev.azure.com/graphql-rust/GraphQL%20Rust/_apis/build/status/graphql-rust.juniper)](https://dev.azure.com/graphql-rust/GraphQL%20Rust/_build/latest?definitionId=1) Stars:`3.3K`.
  * [graphql-client](https://github.com/graphql-rust/graphql-client) - Typed, correct GraphQL requests and responses. [![GitHub actions Status](https://github.com/graphql-rust/graphql-client/workflows/CI/badge.svg?branch=master)](https://github.com/graphql-rust/graphql-client/actions) Stars:`1.1K`.
  * [alexcrichton/curl-rust](https://github.com/alexcrichton/curl-rust) - [libcurl](https://curl.se/libcurl/) bindings Stars:`1.0K`.
  * [DoumanAsh/yukikaze](https://gitlab.com/Douman/yukikaze) [[yukikaze](https://crates.io/crates/yukikaze)] - Beautiful and elegant Yukikaze is little HTTP client library based on hyper. [![build badge](https://gitlab.com/Douman/yukikaze/badges/master/pipeline.svg)](https://gitlab.com/Douman/yukikaze)
* HTTP Server
  * [Rocket](https://github.com/rwf2/Rocket) - Rocket is a web framework with a focus on ease-of-use, expressability, and speed Stars:`23.8K`.
  * [actix/actix-web](https://github.com/actix/actix-web) - A lightweight async web framework with websocket support Stars:`20.7K`.
  * [tokio/axum](https://github.com/tokio-rs/axum) - Ergonomic and modular web framework built with Tokio, Tower, and Hyper [![Build badge](https://github.com/tokio-rs/axum/actions/workflows/CI.yml/badge.svg?branch=main)](https://github.com/tokio-rs/axum/actions/workflows/CI.yml) Stars:`17.3K`.
  * [hyperium/hyper](https://github.com/hyperium/hyper) - an HTTP implementation [![CI](https://github.com/hyperium/hyper/workflows/CI/badge.svg?branch=master)](https://github.com/hyperium/hyper/actions?query=workflow%3ACI) Stars:`14.1K`.
  * [seanmonstar/warp](https://github.com/seanmonstar/warp) - A super-easy, composable, web server framework for warp speeds. [![crate](https://img.shields.io/crates/v/create-rust-app.svg)](https://crates.io/crates/warp) Stars:`9.3K`.
  * [Iron](https://github.com/iron/iron) - A middleware-based server framework Stars:`6.1K`.
  * [Juniper](https://github.com/graphql-rust/juniper) - GraphQL server library Stars:`5.6K`.
  * [poem-web/poem](https://github.com/poem-web/poem) - A full-featured and easy-to-use web framework. [![CI](https://github.com/poem-web/poem/actions/workflows/ci.yml/badge.svg)](https://github.com/poem-web/poem/actions/workflows/ci.yml) Stars:`3.3K`.
  * [Nickel](https://github.com/nickel-org/nickel.rs/) - inspired by [Express](http://expressjs.com/) Stars:`3.0K`.
  * [Salvo](https://github.com/salvo-rs/salvo) - an easy to use webframework base on hyper and tokio. [![build build](https://github.com/salvo-rs/salvo/workflows/CI%20(Linux)/badge.svg?branch=master&event=push)](https://github.com/salvo-rs/salvo/actions) Stars:`2.9K`.
  * [Gotham](https://github.com/gotham-rs/gotham) - A flexible web framework that does not sacrifice safety, security or speed. Stars:`2.2K`.
  * [handlebars-rust](https://github.com/sunng87/handlebars-rust) - an Iron web framework middleware. Stars:`1.2K`.
  * [tomaka/rouille](https://github.com/tomaka/rouille) - Web framework Stars:`1.1K`.
  * [carllerche/tower-web](https://github.com/carllerche/tower-web) [[tower-web](https://crates.io/crates/tower-web)] - A fast, boilerplate free, web framework Stars:`979`.
  * [tiny-http](https://github.com/tiny-http/tiny-http) - Low level HTTP server library Stars:`970`.
  * [Zino](https://github.com/zino-rs/zino) - Next-generation framework for composable applications Stars:`697`.
  * [miketang84/sapper](https://github.com/miketang84/sapper) - A lightweight web framework built on async hyper. Stars:`617`.
  * [Rustless](https://github.com/rustless/rustless) - A REST-like API micro-framework inspired by [Grape](https://github.com/ruby-grape/grape) and [Hyper](https://github.com/hyperium/hyper) Stars:`617`.
  * [Graphul](https://github.com/graphul-rs/graphul) - An Express-inspired web framework. [![crate](https://img.shields.io/crates/v/create-rust-app.svg)](https://crates.io/crates/graphul) Stars:`440`.
  * [GildedHonour/frank_jwt](https://github.com/GildedHonour/frank_jwt) - JSON Web Token implementation. Stars:`251`.
  * [Anansi](https://github.com/saru-tora/anansi) - A simple full-stack web framework Stars:`112`.
  * [danclive/sincere](https://github.com/danclive/sincere) - A micro web framework based on hyper and multithreading. Stars:`96`.
  * [Saphir](https://github.com/richerarc/saphir) - A progressive web framework with low-level control, without the pain. Stars:`90`.
  * [branca](https://crates.io/crates/branca) - Implementation of Branca for Authenticated and Encrypted API tokens.
* Miscellaneous
  * [serenity-rs/serenity](https://github.com/serenity-rs/serenity) [[serenity](https://crates.io/crates/serenity)] - A library for the Discord API Stars:`4.6K`.
  * [osohq/oso](https://github.com/osohq/oso) [[oso](https://crates.io/crates/oso)] - A policy engine for authorization that's embedded in your application. [![Build Status](https://github.com/osohq/oso/workflows/Development/badge.svg?branch=main)](https://github.com/osohq/oso/actions?query=branch%3Amain+workflow%3ADevelopment) Stars:`3.4K`.
  * [teloxide/teloxide](https://github.com/teloxide/teloxide/) - An elegant Telegram bots framework [![Build Status](https://github.com/teloxide/teloxide/workflows/Continuous%20integration/badge.svg?branch=master)](https://github.com/teloxide/teloxide/actions) Stars:`2.8K`.
  * [svix/svix-webhooks](https://github.com/svix/svix-webhooks) [[svix](https://crates.io/crates/svix)] - A library for sending webhooks and verifying signatures. Stars:`2.2K`.
  * [juhaku/utoipa](https://github.com/juhaku/utoipa) - Simple, Fast, Code first and Compile time generated OpenAPI documentation [![crates.io](https://img.shields.io/crates/v/utoipa.svg?label=crates.io&color=orange&logo=rust)](https://crates.io/crates/utoipa) [![Utoipa build](https://github.com/juhaku/utoipa/actions/workflows/build.yaml/badge.svg)](https://github.com/juhaku/utoipa/actions/workflows/build.yaml) Stars:`2.0K`.
  * [causal-agent/scraper](https://github.com/causal-agent/scraper) [[scraper](https://crates.io/crates/scraper)] - HTML parsing and querying with CSS selectors. [![Build Status](https://github.com/causal-agent/scraper/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/causal-agent/scraper/actions) Stars:`1.8K`.
  * [pyrossh/rust-embed](https://github.com/pyrossh/rust-embed) - A macro to embed static assets into the rust binary Stars:`1.6K`.
  * [utkarshkukreti/select.rs](https://github.com/utkarshkukreti/select.rs) [[select](https://crates.io/crates/select)] - A library to extract useful data from HTML documents, suitable for web scraping. Stars:`952`.
  * [cargonauts](https://github.com/cargonauts-rs/cargonauts) - A web framework intended for building maintainable, well-factored web apps. Stars:`179`.
  * [hominee/dyer](https://github.com/hominee/dyer) [[dyer](https://crates.io/crates/dyer)] - dyer is designed for reliable, flexible and fast Request-Response based service, including data processing, web-crawling and so on, providing some friendly, flexible, comprehensive features without compromising speed. Stars:`133`.
  * [softprops/openapi](https://github.com/softprops/openapi) - A library for processing openapi spec files Stars:`125`.
  * [rookie](https://github.com/thewh1teagle/rookie) - Load cookies from any browser on any platform. ![crates.io](https://img.shields.io/crates/v/rookie.svg) Stars:`110`.
  * [tu6ge/valitron](https://github.com/tu6ge/valitron) [[valitron](https://crates.io/crates/valitron)] - An ergonomic, functional and configurable validator Stars:`58`.
  * [pwoolcoc/soup](https://gitlab.com/pwoolcoc/soup) [[soup](https://crates.io/crates/soup)] - A library similar to Python's BeautifulSoup, designed to enable quick and easy manipulation and querying of HTML documents. [![Build Status](https://gitlab.com/pwoolcoc/soup/badges/master/pipeline.svg)](https://gitlab.com/pwoolcoc/soup/badges/master/pipeline.svg)
  * [tbot](https://gitlab.com/SnejUgal/tbot) [[tbot](https://crates.io/crates/tbot)] - Make cool Telegram bots easily [![pipeline status](https://gitlab.com/SnejUgal/tbot/badges/master/pipeline.svg)](https://gitlab.com/SnejUgal/tbot/-/commits/master)
* Reverse Proxy
  * [sozu-proxy/sozu](https://github.com/sozu-proxy/sozu) [[sozu](https://crates.io/crates/sozu)] - A HTTP reverse proxy. [![CI](https://github.com/sozu-proxy/sozu/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/sozu-proxy/sozu/actions/workflows/ci.yml) Stars:`2.9K`.
* Static Site Generators
  * [getzola/zola](https://github.com/getzola/zola) [[zola](https://www.getzola.org/)] - An opinionated static site generator with everything built-in. [![Build Status](https://dev.azure.com/getzola/zola/_apis/build/status/getzola.zola?branchName=master)](https://dev.azure.com/getzola/zola/_build) Stars:`13.0K`.
  * [vi/websocat](https://github.com/vi/websocat) - CLI for interacting with WebSockets, with functionality of Netcat, Curl and Socat. Stars:`6.7K`.
  * [snapview/tungstenite-rs](https://github.com/snapview/tungstenite-rs) - Lightweight stream-based WebSocket implementation. Stars:`1.8K`.
  * [rust-websocket](https://github.com/websockets-rs/rust-websocket) - A framework for dealing with WebSocket connections (both clients and servers) Stars:`1.5K`.
  * [housleyjk/ws-rs](https://github.com/housleyjk/ws-rs) - lightweight, event-driven WebSockets Stars:`1.5K`.
  * [cobalt-org/cobalt.rs](https://github.com/cobalt-org/cobalt.rs) - Static site generator [![Build Status](https://dev.azure.com/cobalt-org/cobalt-org/_apis/build/status/cobalt.rs?branchName=master)](https://dev.azure.com/cobalt-org/cobalt-org/_build?definitionId=2) Stars:`1.3K`.
  * [grego/blades](https://github.com/grego/blades) [[blades](https://getblades.org/)] - Blazing fast dead simple static site generator. Stars:`327`.
  * [FuGangqiang/mdblog.rs](https://github.com/FuGangqiang/mdblog.rs) [[mdblog](https://crates.io/crates/mdblog)] - Static site generator from markdown files. Stars:`58`.
  * [leven-the-blog/leven](https://github.com/leven-the-blog/leven) [[leven](https://crates.io/crates/leven)] - A simple, parallelized blog generator. Stars:`57`.
  * [iddm/urlshortener-rs](https://github.com/iddm/urlshortener-rs) - A very simple urlshortener library. [![CI](https://github.com/iddm/urlshortener-rs/actions/workflows/ci.yml/badge.svg)](https://github.com/iddm/urlshortener-rs/actions/workflows/ci.yml) [![Crates badge](https://img.shields.io/crates/v/urlshortener.svg)](https://crates.io/crates/urlshortener) Stars:`49`.
* [WebSocket](https://datatracker.ietf.org/doc/rfc6455/)

## Registries

A registry allows you to publish your Rust libraries as crate packages, to share them with others publicly and privately.

* [w4/chartered](https://github.com/w4/chartered) - A private, authenticated, permissioned Cargo registry [![CI](https://github.com/w4/chartered/actions/workflows/ci.yml/badge.svg)](https://github.com/w4/chartered/actions/workflows/ci.yml) Stars:`129`.
* [Cloudsmith :heavy_dollar_sign:](https://cloudsmith.com/product/formats/cargo-registry) - A fully managed package management SaaS, with first-class support for public and private Cargo/Rust registries (plus many others). Has a generous free-tier and is also completely free for open-source.
* [Crates](https://crates.io) - The official public registry for Rust/Cargo.

## Resources

* Benchmarks
  * [TeXitoi/benchmarksgame-rs](https://github.com/TeXitoi/benchmarksgame-rs) - Implementations for the [The Computer Language Benchmarks Game](https://benchmarksgame-team.pages.debian.net/benchmarksgame/) Stars:`68`.
* Decks & Presentations
  * [Learning systems programming with Rust](https://speakerdeck.com/jvns/learning-systems-programming-with-rust) - Presented by [Julia Evans](https://twitter.com/@b0rk) @ Rustconf 2016.
  * [Rust: Hack Without Fear!](https://www.youtube.com/watch?v=lO1z-7cuRYI) - Presented by [Nicholas Matsakis](https://github.com/nikomatsakis) @ C++Now 2018
  * [Shipping a Solid Rust Crate](https://www.youtube.com/watch?v=t4CyEKb-ywA) - Presented by [Michael Gattozzi](https://github.com/mgattozzi) @ RustConf 2017
* [Discover Rust Libraries & Code Snippets](https://kandi.openweaver.com/explorelibrary/rust) - A curated list of libraries, authors, kits, tutorials & learning resources on kandi
* Learning
  * [Rustlings](https://github.com/rust-lang/rustlings) - small exercises to get you used to reading and writing Rust code Stars:`51.2K`.
  * [rust-learning](https://github.com/ctjhoa/rust-learning) - A collection of useful resources to learn Rust Stars:`11.3K`.
  * [Easy Rust](https://github.com/Dhghomon/easy_rust) - Learn Rust in easy English. Stars:`7.9K`.
  * [Idiomatic Rust](https://github.com/mre/idiomatic-rust) - A peer-reviewed collection of articles/talks/repos which teach idiomatic Rust. Stars:`6.2K`.
  * [stdx](https://github.com/brson/stdx) - Learn these crates first as an extension to std Stars:`2.0K`.
  * [Aquascope](https://github.com/cognitive-engineering-lab/aquascope) - Interactive visualizations of Rust at compile-time and run-time Stars:`1.9K`.
  * [rust-how-do-i-start](https://github.com/jondot/rust-how-do-i-start) - A repo dedicated to answering the question: "So, Rust. How do I _start_?". A beginner only hand-picked resources and learning track. Stars:`1.0K`.
  * [Rusty CS](https://github.com/AbdesamedBendjeddou/Rusty-CS) - A Computer Science Curriculum that helps practice the acquired academic knowledge in Rust Stars:`875`.
  * [Rust Gym](https://github.com/warycat/rustgym) - A big collection of coding interview problems solved in Rust. Stars:`861`.
  * [Awesome Rust Streaming](https://github.com/jamesmunns/awesome-rust-streaming) - A community curated list of livestreams. Stars:`685`.
  * [Rust Flashcards](https://github.com/ad-si/Rust-Flashcards) - Over 550 flashcards to learn Rust from first principles. Stars:`512`.
  * [Learn Rust by 500 lines code](https://github.com/cuppar/rtd) - Learn Rust by 500 lines code, build a Todo Cli Application from scratch. Stars:`443`.
  * [Rust by Example](https://doc.rust-lang.org/rust-by-example/) - a collection of runnable examples that illustrate various Rust concepts and standard libraries.
  * [Rust in Action](https://www.manning.com/books/rust-in-action) - A hands-on guide to systems programming with Rust by [Tim McNamara](https://github.com/timClicks) (paid)
  * [Programming Community Curated Resources for Learning Rust](https://hackr.io/tutorials/learn-rust) - A list of recommended resources voted by the programming community.
  * [Refactoring to Rust](https://www.manning.com/books/refactoring-to-rust) - A book that introduces to Rust language.
  * [Learning Rust With Entirely Too Many Linked Lists](https://rust-unofficial.github.io/too-many-lists/) - in-depth exploration of Rust's memory management rules, through implementing a few different types of list structures.
  * [Rust Cookbook](https://rust-lang-nursery.github.io/rust-cookbook/) - A collection of simple examples that demonstrate good practices to accomplish common programming tasks, using the crates of the Rust ecosystem.
  * [Hands-on Rust](https://pragprog.com/titles/hwrust/hands-on-rust/) - A hands-on guide to learning Rust by making games - by [Herbert Wolverson](https://github.com/thebracket/) (paid)
  * [Rust for professionals](https://overexact.com/rust-for-professionals/) - A quick introduction to Rust for experienced software developers.
  * [exercism.org](https://exercism.org/tracks/rust) - programming exercises that help you learn new concepts in Rust.
  * [Little Book of Rust Books](https://lborb.github.io/book/) - Curated list of rust books and how-tos.
  * [Rust in Motion](https://www.manning.com/livevideo/rust-in-motion?a_aid=cnichols&a_bid=6a993c2e) - A video series by [Carol Nichols](https://github.com/carols10cents) and [Jake Goulding](https://github.com/shepmaster) (paid)
  * [Rust Language Cheat Sheet](https://cheats.rs/) - Rust Language Cheat Sheet
  * [Rust Tiếng Việt](https://rust-tieng-viet.github.io/) - Learn Rust in Vietnamese.
  * [Comprehensive Rust 🦀](https://google.github.io/comprehensive-rust/) - A 3-day course on Rust Fundamentals plus 1-day courses on Android, Bare-metal Rust, and Concurrency. Available in English, [Brazilian Portuguese](https://google.github.io/comprehensive-rust/pt-BR/), and [Korean](https://google.github.io/comprehensive-rust/ko/).
  * [CodeCrafters.io](https://app.codecrafters.io/tracks/rust) - Build your own Redis, Git, Docker, or SQLite
  * [CIS 198: Rust Programming](http://cis198-2016s.github.io/schedule/) - University of Pennsylvania's Comp Sci Rust Programming Course
  * [Build a language VM](https://blog.subnetzero.io/post/building-language-vm-part-00/) - a series of posts to detailing how to build a language VM.
  * [awesome-rust-mentors](https://rustbeginners.github.io/awesome-rust-mentors/) - A list of helpful mentors willing to take mentees and educate them about Rust and programming.
  * [Take your first steps with Rust](https://learn.microsoft.com/en-us/training/paths/rust-first-steps/) - Lay the foundation of knowledge you need to build fast and effective programs in Rust.
  * [Tour of Rust](https://tourofrust.com) - This is meant to be an interactive step by step guide through the features of the Rust programming language.
* Podcasts
* [Rust Design Patterns](https://github.com/rust-unofficial/patterns) - A catalogue of Rust design patterns, anti-patterns and idioms Stars:`7.9K`.
* [RustBooks](https://github.com/sger/RustBooks) - list of RustBooks Stars:`4.2K`.
* [RustViz](https://github.com/rustviz/rustviz) - generates visualizations from simple Rust programs to assist users in better understanding the Rust Lifetime and Borrowing mechanism. Stars:`2.7K`.
  * [New Rustacean](https://newrustacean.com) - A podcast about learning Rust
  * [Rustacean Station](https://rustacean-station.org/) - A community project for creating podcast content for Rust
* [Rust Guidelines](http://aturon.github.io/) - Aaron Turon's blog posts on rust
* [Rust Servers, Services and Apps - MEAP](https://www.manning.com/books/rust-servers-services-and-apps) - Build backend servers, services, and front-ends in Rust to get fast, reliable, and maintainable applications.
* [Rust Subreddit](https://www.reddit.com/r/rust/) - A subreddit(forum) where rust related questions, articles and resources are posted and discussed
* [RustCamp 2015 Talks](https://www.youtube.com/playlist?list=PLE7tQUdRKcybdIw61JpCoo89i4pWU5f_t) - Recorded talks from RustCamp 2015
* [Watch Jon Gjengset Implement BitTorrent in Rust](https://www.youtube.com/watch?v=jf_ddGnum_4) - Implementing (part of) a BitTorrent client in Rust

## License

[![CC0](https://licensebuttons.net/p/zero/1.0/88x31.png)](https://creativecommons.org/publicdomain/zero/1.0/)
