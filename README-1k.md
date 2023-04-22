# Awesome Rust [![build badge](https://github.com/rust-unofficial/awesome-rust/actions/workflows/rust.yml/badge.svg?branch=main)](https://github.com/rust-unofficial/awesome-rust/actions/workflows/rust.yml) [![Track Awesome List](https://www.trackawesomelist.com/badge.svg)](https://www.trackawesomelist.com/rust-unofficial/awesome-rust/)

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

See also [Rust — Production](https://www.rust-lang.org/production) organizations running Rust in production.

* [denoland/deno](https://github.com/denoland/deno) — A secure JavaScript/TypeScript runtime built with V8, Rust, and Tokio [![Build Status](https://github.com/denoland/deno/workflows/ci/badge.svg?branch=master&event=push)](https://github.com/denoland/deno/actions) Stars:`89.0K`.
* [alacritty](https://github.com/alacritty/alacritty) — A cross-platform, GPU enhanced terminal emulator Stars:`46.2K`.
* [SWC](https://github.com/swc-project/swc) — super-fast TypeScript / JavaScript compiler Stars:`27.1K`.
* [Servo](https://github.com/servo/servo) — A prototype web browser engine Stars:`22.9K`.
* [wasmer](https://github.com/wasmerio/wasmer) — A safe and fast WebAssembly runtime supporting WASI and Emscripten [![Build Status](https://github.com/wasmerio/wasmer/workflows/build/badge.svg?style=flat-square)](https://github.com/wasmerio/wasmer/actions) Stars:`15.0K`.
* [zellij](https://github.com/zellij-org/zellij) — A terminal multiplexer (workspace) with batteries included Stars:`11.5K`.
* [broot](https://github.com/Canop/broot) A new way to see and navigate directory trees (get an overview of a directory, even a big one; find a directory then `cd` to it; never lose track of file hierarchy while you search; manipulate your files, ...), further reading [dystroy.org/broot](https://dystroy.org/broot/) [![Latest Version](https://img.shields.io/crates/v/broot.svg)](https://crates.io/crates/broot) Stars:`8.6K`.
* [wezterm](https://github.com/wez/wezterm) — A GPU-accelerated cross-platform terminal emulator and multiplexer Stars:`8.3K`.
* [cloudflare/boringtun](https://github.com/cloudflare/boringtun) — A Userspace WireGuard VPN Implementation [![build badge](https://img.shields.io/badge/crates.io-v0.2.0-orange.svg)](https://crates.io/crates/boringtun) Stars:`5.1K`.
* [Sniffnet](https://github.com/GyulyVGC/sniffnet) — Cross-platform application to monitor your network traffic with ease [![build badge](https://img.shields.io/github/actions/workflow/status/gyulyvgc/sniffnet/rust.yml?logo=github)](https://github.com/GyulyVGC/sniffnet/blob/main/.github/workflows/rust.yml) [![crate](https://img.shields.io/crates/v/sniffnet?logo=rust)](https://crates.io/crates/sniffnet) Stars:`4.9K`.
* [datafusion](https://github.com/apache/arrow-datafusion) — Apache Arrow DataFusion and Ballista query engines Stars:`3.5K`.
* [rx](https://github.com/cloudhead/rx) — Vi inspired Modern Pixel Art Editor Stars:`2.7K`.
* [habitat](https://github.com/habitat-sh/habitat) — A tool created by Chef to build, deploy, and manage applications. Stars:`2.5K`.
* [notty](https://github.com/withoutboats/notty) — A new kind of terminal Stars:`2.2K`.
* [shuttle](https://github.com/shuttle-hq/shuttle) — A serverless platform built for Rust Stars:`2.2K`.
* [linkerd/linkerd2-proxy](https://github.com/linkerd/linkerd2-proxy) — Ultralight service mesh for Kubernetes. Stars:`1.7K`.
* [fcsonline/drill](https://github.com/fcsonline/drill) — A HTTP load testing application inspired by Ansible syntax Stars:`1.6K`.
* [kalker](https://github.com/PaddiM8/kalker) - A scientific calculator that supports math-like syntax with user-defined variables, functions, derivation, integration, and complex numbers. Cross platform + WASM support [![Build Status](https://github.com/PaddiM8/kalker/workflows/Release/badge.svg)](https://github.com/PaddiM8/kalker/actions) Stars:`1.3K`.

### Audio and Music

* [Spotify TUI](https://github.com/Rigellute/spotify-tui) — A Spotify client for the terminal written in Rust. ![Continuous Integration](https://github.com/Rigellute/spotify-tui/workflows/Continuous%20Integration/badge.svg?branch=master) Stars:`14.7K`.
* [Spotifyd](https://github.com/Spotifyd/spotifyd) — An open source Spotify client running as a UNIX daemon. ![Continuous Integration](https://github.com/Spotifyd/spotifyd/workflows/Continuous%20Integration/badge.svg?branch=master) Stars:`8.8K`.
* [ncspot](https://github.com/hrkfdn/ncspot) - Cross-platform ncurses Spotify client, inspired by ncmpc and the likes. [![build badge](https://github.com/hrkfdn/ncspot/workflows/Build/badge.svg)](https://github.com/hrkfdn/ncspot/actions?query=workflow%3ABuild) Stars:`3.7K`.
* [Glicol](https://github.com/chaosprint/glicol) — Graph-oriented live coding language written in Rust for collaborative musicking in browsers. Stars:`1.5K`.
* [Polaris](https://github.com/agersant/polaris) — A music streaming application. Stars:`1.1K`.

### Cryptocurrencies

* [Diem](https://github.com/diem/diem) — Diem’s mission is to enable a simple global currency and financial infrastructure that empowers billions of people. Stars:`16.7K`.
* [Solana](https://github.com/solana-labs/solana) — Incredibly fast, highly scalable blockchain using Proof-of-History. Stars:`10.2K`.
* [Substrate](https://github.com/paritytech/substrate) — Generic modular blockchain template written in Rust Stars:`8.1K`.
* [Polkadot](https://github.com/paritytech/polkadot) — Heterogeneous multi‑chain technology with pooled security Stars:`6.7K`.
* [Foundry](https://github.com/foundry-rs/foundry) - Foundry is a blazing fast, portable and modular toolkit for Ethereum application development written in Rust. ![Build Status](https://img.shields.io/github/workflow/status/foundry-rs/foundry/test?style=flat-square) Stars:`5.8K`.
* [zcash](https://github.com/zcash/zcash) — Zcash is an implementation of the "Zerocash" protocol. Stars:`4.7K`.
* [Lighthouse](https://github.com/sigp/lighthouse) — Rust Ethereum Consensus Layer (CL) Client [![Build Status](https://github.com/sigp/lighthouse/workflows/test-suite/badge.svg?branch=master)](https://github.com/sigp/lighthouse/actions) Stars:`2.3K`.
* [near/nearcore](https://github.com/near/nearcore) — decentralized smart-contract platform for low-end mobile devices. Stars:`2.0K`.
* [ethers-rs](https://github.com/gakonst/ethers-rs) - Complete Ethereum & Celo library and wallet implementation in Rust. ![Build Status](https://github.com/gakonst/ethers-rs/workflows/Tests/badge.svg) Stars:`1.7K`.
* [rust-bitcoin](https://github.com/rust-bitcoin/rust-bitcoin) — Library with support for de/serialization, parsing and executing on data structures and network messages related to Bitcoin. Stars:`1.4K`.
* [CITA](https://github.com/citahub/cita) — A high performance blockchain kernel for enterprise users. Stars:`1.3K`.
* [Joystream](https://github.com/Joystream/joystream) — A user governed video platform Stars:`1.2K`.
* [Nervos CKB](https://github.com/nervosnetwork/ckb) — Nervos CKB is a public permissionless blockchain, the common knowledge layer of Nervos network. Stars:`1.1K`.

### Database

* [SurrealDB](https://github.com/surrealdb/surrealdb) — A scalable, distributed, document-graph database [![Build Status](https://img.shields.io/github/workflow/status/surrealdb/surrealdb/Continuous%20integration/main)](https://github.com/surrealdb/surrealdb/actions) Stars:`19.7K`.
* [tikv](https://github.com/tikv/tikv) — A distributed KV database in Rust [![Build Status](https://ci.pingcap.net/job/tikv_ghpr_test/badge/icon)](https://ci.pingcap.net/job/tikv_ghpr_test/) Stars:`13.0K`.
* [Qdrant](https://github.com/qdrant/qdrant) - An open source vector similarity search engine with extended filtering support [![Tests](https://github.com/qdrant/qdrant/workflows/Tests/badge.svg)](https://github.com/qdrant/qdrant/actions) Stars:`6.2K`.
* [Databend](https://github.com/datafuselabs/databend) - A Modern Real-Time Data Processing & Analytics DBMS with Cloud-Native Architecture [![Release](https://github.com/datafuselabs/databend/actions/workflows/databend-release.yml/badge.svg)](https://github.com/datafuselabs/databend/actions/workflows/databend-release.yml) Stars:`5.9K`.
* [Materialize](https://github.com/MaterializeInc/materialize) - Streaming SQL database powered by Timely Dataflow :heavy_dollar_sign: [![Build status](https://badge.buildkite.com/97d6604e015bf633d1c2a12d166bb46f3b43a927d3952c999a.svg?branch=main)](https://buildkite.com/materialize/tests) Stars:`5.0K`.
* [noria](https://github.com/mit-pdos/noria) [[noria](https://crates.io/crates/noria)] — Dynamically changing, partially-stateful data-flow for web application backends Stars:`4.7K`.
* [RisingWaveLabs/RisingWave](https://github.com/RisingWaveLabs/risingwave) - the next-generation streaming database in the cloud [![CI](https://github.com/RisingWaveLabs/risingwave/actions/workflows/main.yml/badge.svg)](https://github.com/RisingWaveLabs/risingwave/actions/workflows/main.yml/badge.svg?branch=main) Stars:`4.3K`.
* [CozoDB](https://github.com/cozodb/cozo) - A transactional, relational database that uses Datalog and focuses on graph data and algorithms. Time-travel-capable, and fast! [![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/cozodb/cozo/build.yml?branch=main)](https://github.com/cozodb/cozo/actions/workflows/build.yml) Stars:`2.3K`.
* [seppo0010/rsedis](https://github.com/seppo0010/rsedis) — A Redis reimplementation in Rust Stars:`1.7K`.
* [Skytable](https://github.com/skytable/skytable) — A multi-model NoSQL database ![GitHub Workflow Status](https://img.shields.io/github/workflow/status/skytable/skytable/Tests?style=flat-square) Stars:`1.6K`.
* [PumpkinDB](https://github.com/PumpkinDB/PumpkinDB) — an event sourcing database engine Stars:`1.3K`.

### Emulators

See also [crates matching keyword 'emulator'](https://crates.io/keywords/emulator).

* CHIP-8
* Commodore 64
* Flash Player
  * [Ruffle](https://github.com/ruffle-rs/ruffle) — Ruffle is an Adobe Flash Player emulator written in the Rust programming language. Ruffle targets both the desktop and the web using WebAssembly. [![CI](https://github.com/ruffle-rs/ruffle/actions/workflows/test_rust.yml/badge.svg)](https://github.com/ruffle-rs/ruffle/actions/workflows/test_rust.yml)[![CI](https://github.com/ruffle-rs/ruffle/actions/workflows/test_web.yml/badge.svg)](https://github.com/ruffle-rs/ruffle/actions/workflows/test_web.yml) Stars:`12.3K`.
* Gameboy
  * [mohanson/gameboy](https://github.com/mohanson/gameboy) — Full featured Cross-platform GameBoy emulator. Forever boys!. Stars:`1.3K`.
* Gameboy Advance
* GameMaker
* Intel 8080 CPU
* iOS
  * [touchHLE](https://github.com/hikari-no-yume/touchHLE) — High-level emulator for iPhone OS apps Stars:`1.4K`.
* iPod
* NES
* PlayStation 4
* ZX Spectrum

### Games

See also [Games Made With Piston](https://github.com/PistonDevelopers/piston/wiki/Games-Made-With-Piston).

* [citybound](https://github.com/citybound/citybound) — The city sim you deserve Stars:`7.3K`.
* [cristicbz/rust-doom](https://github.com/cristicbz/rust-doom) — A renderer for Doom, may progress to being a playable game Stars:`2.2K`.
* [ozkriff/zemeroth](https://github.com/ozkriff/zemeroth) — A small 2D turn-based hexagonal strategy game Stars:`1.3K`.

### Graphics

* [ivanceras/svgbob](https://github.com/ivanceras/svgbob) — converts ASCII diagrams into SVG graphics Stars:`3.5K`.
* [RazrFalcon/resvg](https://github.com/RazrFalcon/resvg) — An SVG rendering library. Stars:`2.0K`.

### Image processing

* [shssoichiro/oxipng](https://github.com/shssoichiro/oxipng) [[oxipng](https://crates.io/crates/oxipng)] — Multithreaded PNG optimizer written in Rust. [![Build Status](https://github.com/shssoichiro/oxipng/workflows/oxipng/badge.svg)](https://github.com/shssoichiro/oxipng/actions?query=branch%3Amaster) [![Version](https://img.shields.io/crates/v/oxipng.svg)](https://crates.io/crates/oxipng) Stars:`2.0K`.

### Industrial automation


### Observability

* [vectordotdev/vector](https://github.com/vectordotdev/vector) — A High-Performance, Logs, Metrics, & Events Router. Stars:`13.3K`.
* [Quickwit-oss/quickwit](https://github.com/quickwit-oss/quickwit) - Cloud-native and highly cost-efficient search engine for log management. [![CI](https://github.com/quickwit-oss/quickwit/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/quickwit-oss/quickwit/actions?query=workflow%3ACI) Stars:`3.3K`.
* [Scaphandre](https://github.com/hubblo-org/scaphandre) - A power consumption monitoring agent, to track host and each service power consumption and enable designing systems and applications for more sustainability. Designed to fit any monitoring toolchain (already supports prometheus, warp10, riemann...). Stars:`1.1K`.

### Operating systems

See also [A comparison of operating systems written in Rust](https://github.com/flosse/rust-os-comparison).
* [tock/tock](https://github.com/tock/tock) — A secure embedded operating system for Cortex-M based microcontrollers Stars:`4.3K`.
* [theseus-os/Theseus](https://github.com/theseus-os/Theseus) — A safe-language, single address space and single privilege level OS written from scratch in pure Rust - [![build badge](https://img.shields.io/github/workflow/status/theseus-os/Theseus/Documentation?label=docs%20build)](https://www.theseus-os.com/Theseus/book/index.html) Stars:`2.4K`.

### Payments

* [hyperswitch](https://github.com/juspay/hyperswitch) — An open source payments orchestrator that lets you connect with multiple payment processors and route payment traffic effortlessly, all with a single API integration ![GitHub last commit](https://img.shields.io/github/last-commit/juspay/hyperswitch?style=flat-square) Stars:`2.9K`.

### Productivity

* [espanso](https://github.com/espanso/espanso) — A cross-platform Text Expander written in Rust[![CI](https://github.com/espanso/espanso/actions/workflows/ci.yml/badge.svg?branch=dev&event=push)](https://github.com/espanso/espanso/actions/workflows/ci.yml) Stars:`7.3K`.

### Security tools

* [rustscan/rustscan](https://github.com/RustScan/RustScan) — Make Nmap faster with this port scanning tool [![build badge](https://github.com/RustScan/RustScan/workflows/Continuous%20integration/badge.svg?branch=master)](https://github.com/RustScan/RustScan/actions?query=workflow%3A%22Continuous+integration%22) Stars:`9.8K`.
* [epi052/feroxbuster](https://github.com/epi052/feroxbuster) - A simple, fast, recursive content discovery tool written in Rust ( Stars:`4.1K`.
* [kpcyrd/sn0int](https://github.com/kpcyrd/sn0int) — A semi-automatic OSINT framework and package manager Stars:`1.5K`.
* [AFLplusplus/LibAFL](https://github.com/AFLplusplus/LibAFL) - Advanced Fuzzing Library - Slot your Fuzzer together in Rust! Scales across cores and machines. For Windows, Android, MacOS, Linux, no_std, etc. [![build and test](https://github.com/AFLplusplus/LibAFL/actions/workflows/build_and_test.yml/badge.svg)](https://github.com/AFLplusplus/LibAFL/actions/workflows/build_and_test.yml) Stars:`1.4K`.

### Simulation


### Social networks

* Mastodon

### System tools

* [sharkdp/bat](https://github.com/sharkdp/bat) — A cat(1) clone with wings. [![CICD](https://github.com/sharkdp/bat/actions/workflows/CICD.yml/badge.svg?branch=master)](https://github.com/sharkdp/bat/actions/workflows/CICD.yml) Stars:`40.8K`.
* [sharkdp/fd](https://github.com/sharkdp/fd) — A simple, fast and user-friendly alternative to find. [![CICD](https://github.com/sharkdp/fd/actions/workflows/CICD.yml/badge.svg)](https://github.com/sharkdp/fd/actions/workflows/CICD.yml) Stars:`27.2K`.
* [nushell/nushell](https://github.com/nushell/nushell) - A new type of shell Stars:`24.3K`.
* [ogham/exa](https://github.com/ogham/exa) — A replacement for 'ls' Stars:`21.4K`.
* [uutils/coreutils](https://github.com/uutils/coreutils) — A cross-platform Rust rewrite of the GNU coreutils [[![CICD](https://github.com/uutils/coreutils/actions/workflows/CICD.yml/badge.svg)](https://github.com/uutils/coreutils/actions/workflows/CICD.yml) Stars:`14.5K`.
* [gitui](https://github.com/extrawurst/gitui) - Blazing fast terminal client for git written in Rust. [![build](https://github.com/extrawurst/gitui/workflows/CI/badge.svg?branch=master)](https://github.com/extrawurst/gitui/actions) Stars:`12.8K`.
* [qarmin/cakawka](https://github.com/qarmin/czkawka) - Multi-functional app to find duplicates, empty folders, similar images, etc. [![GitHub Actions Workflow](https://github.com/qarmin/czkawka/actions/workflows/pages/pages-build-deployment/badge.svg?branch=master)](https://github.com/qarmin/czkawka/actions) Stars:`10.1K`.
* [lsd](https://github.com/lsd-rs/lsd) — An ls with a lot of pretty colors and awesome icons [![build](https://github.com/lsd-rs/lsd/workflows/CICD/badge.svg?branch=master)](https://github.com/lsd-rs/lsd/actions) Stars:`9.7K`.
* [XAMPPRocky/tokei](https://github.com/XAMPPRocky/tokei) — counts the lines of code Stars:`8.2K`.
* [bandwhich](https://github.com/imsnif/bandwhich) — Terminal bandwidth utilization tool Stars:`7.8K`.
* [bottom](https://github.com/ClementTsang/bottom) - Yet another cross-platform graphical process/system monitor. [![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/ClementTsang/bottom/ci/master)](https://github.com/ClementTsang/bottom/actions?query=branch%3Amaster) Stars:`6.6K`.
* [dust](https://github.com/bootandy/dust) — A more intuitive version of du Stars:`6.1K`.
* [cantino/mcfly](https://github.com/cantino/mcfly) - Fly through your shell history. Great Scott! Stars:`5.1K`.
* [lotabout/skim](https://github.com/lotabout/skim) — A fuzzy finder in pure rust Stars:`4.1K`.
* [dalance/procs](https://github.com/dalance/procs) — A modern replacement for 'ps' written by Rust [![Regression](https://github.com/dalance/procs/actions/workflows/regression.yml/badge.svg)](https://github.com/dalance/procs/actions/workflows/regression.yml) Stars:`3.9K`.
* [watchexec](https://github.com/watchexec/watchexec) — Executes commands in response to file modifications Stars:`3.9K`.
* [pueue](https://github.com/nukesor/pueue) — Manage your long running shell commands. [![GitHub Actions Workflow](https://github.com/nukesor/pueue/workflows/Test%20build/badge.svg?branch=master)](https://github.com/nukesor/pueue/actions) Stars:`3.5K`.
* [orhun/kmon](https://github.com/orhun/kmon) — Linux Kernel Manager and Activity Monitor ![https://github.com/orhun/kmon/actions](https://img.shields.io/github/workflow/status/orhun/kmon/Continuous%20Integration/master?label=build) Stars:`1.8K`.
* [diskonaut](https://github.com/imsnif/diskonaut) — Terminal visual disk space navigator Stars:`1.7K`.
* [m4b/bingrep](https://github.com/m4b/bingrep) — Greps through binaries from various OSs and architectures, and colors them. Stars:`1.6K`.
* [redox-os/ion](https://github.com/redox-os/ion) — Next-generation system shell Stars:`1.4K`.
* [ouch](https://github.com/ouch-org/ouch) - Painless compression and decompression on the command-line [![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/ouch-org/ouch/build-and-test)](https://github.com/ouch-org/ouch/actions?query=branch%3Amaster) Stars:`1.3K`.
* [Kondo](https://github.com/tbillington/kondo) - CLI & GUI tool for deleting software project artifacts and reclaiming disk space Stars:`1.1K`.

### Task scheduling


### Text editors

* [Lapce](https://github.com/lapce/lapce) — A modern editor with a backend written in Rust. Taking inspiration from the discontinued [xi-editor](https://github.com/xi-editor/xi-editor). Stars:`25.2K`.
* [helix](https://github.com/helix-editor/helix) — A post-modern modal text editor inspired by Neovim/Kakoune. [![build badge](https://github.com/helix-editor/helix/actions/workflows/build.yml/badge.svg)](https://github.com/helix-editor/helix/actions) Stars:`21.3K`.
* [ox](https://github.com/curlpipe/ox) — An independent Rust text editor that runs in your terminal! Stars:`2.9K`.
* [gchp/iota](https://github.com/gchp/iota) — A simple text editor Stars:`1.6K`.
* [emacs-ng](https://github.com/emacs-ng/emacs-ng) — Complementing the C codebase with rust code to introduce new features. Stars:`1.4K`.

### Text processing

* [grex](https://github.com/pemistahl/grex) — A command-line tool and library for generating regular expressions from user-provided test cases Stars:`6.0K`.
* [phiresky/ripgrep-all](https://github.com/phiresky/ripgrep-all) — ripgrep, but also search in PDFs, E-Books, Office documents, zip, tar.gz, etc. Stars:`5.5K`.
* [Melody](https://github.com/yoav-lavi/melody) - A language that compiles to regular expressions and aims to be more easily readable and maintainable [![build badge](https://github.com/yoav-lavi/melody/actions/workflows/rust.yml/badge.svg)](https://github.com/yoav-lavi/melody/actions/workflows/rust.yml) [![crates.io](https://img.shields.io/crates/v/melody_compiler?label=compiler)](https://crates.io/crates/melody_compiler) Stars:`4.1K`.
* [dominikwilkowski/cfonts](https://github.com/dominikwilkowski/cfonts) [[cfonts](https://crates.io/crates/cfonts)] — Sexy ANSI fonts for the console ![build badge](https://github.com/dominikwilkowski/cfonts/actions/workflows/testing.yml/badge.svg) Stars:`1.3K`.

### Utilities

![CI](https://github.com/evansmurithi/cloak/workflows/CI/badge.svg) [![build badge](https://ci.appveyor.com/api/projects/status/9mlfpfru3ng4c689/branch/master?svg=true)](https://ci.appveyor.com/project/evansmurithi/cloak)
* [rustdesk/rustdesk](https://github.com/rustdesk/rustdesk) — A remote desktop software, great alternative to TeamViewer and AnyDesk. Stars:`40.9K`.
* [vaultwarden](https://github.com/dani-garcia/vaultwarden#readme) [![Build](https://github.com/dani-garcia/vaultwarden/actions/workflows/build.yml/badge.svg)](https://github.com/dani-garcia/vaultwarden/actions/workflows/build.yml) — Alternative implementation of the Bitwarden server API written in Rust Stars:`24.5K`.

### Video

* [xiph/rav1e](https://github.com/xiph/rav1e) — The fastest and safest AV1 encoder. Stars:`3.2K`.

### Virtualization

* [firecracker-microvm/firecracker](https://github.com/firecracker-microvm/firecracker) — A lightweight virtual machine for container workload [Firecracker Microvm](https://firecracker-microvm.github.io/) Stars:`21.4K`.
* [containers/youki](https://github.com/containers/youki) — A container runtime in Rust [![build badge](https://github.com/containers/youki/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/containers/youki/actions) Stars:`4.6K`.
* [tailhook/vagga](https://github.com/tailhook/vagga) — A containerization tool without daemons Stars:`1.8K`.

### Web

* [LemmyNet/lemmy](https://github.com/LemmyNet/lemmy) — A link aggregator / reddit clone for the fediverse [![Build Status](https://cloud.drone.io/api/badges/LemmyNet/lemmy/status.svg)](https://cloud.drone.io/LemmyNet/lemmy) Stars:`7.5K`.
* [libreddit](https://github.com/libreddit/libreddit) - An alternative private front-end to Reddit Stars:`4.3K`.
* [Plume-org/Plume](https://github.com/Plume-org/Plume) — ActivityPub federating blogging application Stars:`1.9K`.

### Web Servers

* [svenstaro/miniserve](https://github.com/svenstaro/miniserve) — A small, self-contained cross-platform CLI tool that allows you to just grab the binary and serve some file(s) via HTTP [![build badge](https://github.com/svenstaro/miniserve/workflows/CI/badge.svg?branch=master)](https://github.com/svenstaro/miniserve/actions) Stars:`4.6K`.

## Development tools

* [just](https://github.com/casey/just) — A handy command runner for project-specific tasks Stars:`11.6K`.
* [git-cliff](https://github.com/orhun/git-cliff) — A highly customizable Changelog Generator that follows Conventional Commit specifications ![https://github.com/orhun/git-cliff/actions](https://img.shields.io/github/workflow/status/orhun/git-cliff/Continuous%20Integration/main?label=build) Stars:`5.5K`.
* [Rustup](https://github.com/rust-lang/rustup) — the Rust toolchain installer [![build badge](https://github.com/rust-lang/rustup/workflows/Linux%20(master)/badge.svg?branch=master)](https://github.com/rust-lang/rustup/actions) Stars:`5.3K`.
* [Racer](https://github.com/racer-rust/racer) — code completion for Rust Stars:`3.4K`.
* [dotenv-linter](https://github.com/dotenv-linter/dotenv-linter) — Linter for `.env` files [![build badge](https://github.com/dotenv-linter/dotenv-linter/workflows/CI/badge.svg?branch=master)](https://github.com/dotenv-linter/dotenv-linter/actions?query=workflow%3ACI+branch%3Amaster) Stars:`1.6K`.
* [geiger](https://github.com/rust-secure-code/cargo-geiger) — A program that list statistics related to usage of unsafe Rust code in a Rust crate and all its dependencies [![Build Status](https://dev.azure.com/cargo-geiger/cargo-geiger/_apis/build/status/rust-secure-code.cargo-geiger?branchName=master)](https://dev.azure.com/cargo-geiger/cargo-geiger/_build/latest?definitionId=1&branchName=master) Stars:`1.2K`.
* [create-rust-app](https://github.com/Wulf/create-rust-app) — Set up a modern rust+react web app by running one command. [![crate](https://img.shields.io/crates/v/create-rust-app.svg)](https://crates.io/crates/create-rust-app) Stars:`1.2K`.
* [Rust Search Extension](https://github.com/huhu/rust-search-extension) — A handy browser extension to search crates and docs in address bar (omnibox). [![Build Status](https://github.com/huhu/rust-search-extension/workflows/build/badge.svg?branch=master)](https://github.com/huhu/rust-search-extension/actions) Stars:`1.0K`.

### Build system

  * [dtolnay/cargo-expand](https://github.com/dtolnay/cargo-expand) — Expand macros in your source code Stars:`2.0K`.
  * [cargo-generate](https://github.com/cargo-generate/cargo-generate) A generator of a rust project by leveraging a pre-existing git repository as a template. Stars:`1.4K`.
  * [cargo-udeps](https://github.com/est31/cargo-udeps) [[cargo-udeps](https://crates.io/crates/cargo-udeps)] — find unused dependencies Stars:`1.1K`.
* CMake
* [Fleet](https://github.com/dimensionhq/fleet) [[fleet-rs](https://crates.io/crates/fleet-rs)] - The blazing fast build tool for Rust. Stars:`2.3K`.
* Github actions

### Debugging

* GDB
  * [gdbgui](https://github.com/cs01/gdbgui) — Browser based frontend for gdb to debug C, C++, Rust, and go. Stars:`9.2K`.
* LLDB

### Deployment

* Docker
  * [emk/rust-musl-builder](https://github.com/emk/rust-musl-builder) — Docker images for compiling static Rust binaries using musl-libc and musl-gcc, with static versions of useful C libraries Stars:`1.4K`.
  * [LukeMathWalker/cargo-chef](https://github.com/LukeMathWalker/cargo-chef) - A tool and pre-built images for caching compiling remote dependencies between Docker builds. Stars:`1.0K`.
* Heroku

### Embedded

[Rust Embedded](https://rust-embedded.org/) focuses on improving the end-to-end experience of using Rust in resource-constrained environments and non-traditional platforms. See [awesome-embedded-rust](https://github.com/rust-embedded/awesome-embedded-rust) for a curated, and more extended list of embedded Rust resources.

* Arduino
* Cross compiling
  * [japaric/rust-cross](https://github.com/japaric/rust-cross) — everything you need to know about cross compiling Rust programs Stars:`2.4K`.
  * [japaric/xargo](https://github.com/japaric/xargo) — effortless cross compilation of Rust programs to custom bare-metal targets like ARM Cortex-M Stars:`1.0K`.
* Espressif
* nRF

### FFI

See also [Foreign Function Interface](https://doc.rust-lang.org/book/first-edition/ffi.html), [The Rust FFI Omnibus](http://jakegoulding.com/rust-ffi-omnibus/) (a collection of examples of using code written in Rust from other languages) and [FFI examples written in Rust](https://github.com/alexcrichton/rust-ffi-examples).

* C
  * [rlhunt/cbindgen](https://github.com/eqrion/cbindgen) — generates C header files from Rust source files. Used in Gecko for WebRender Stars:`1.8K`.
* C++
  * [dtolnay/cxx](https://github.com/dtolnay/cxx) — Safe interop between Rust and C++ [![build badge](https://img.shields.io/badge/github-dtolnay/cxx-8da0cb?style=for-the-badge&labelColor=555555&logo=github)](https://github.com/dtolnay/cxx) Stars:`4.7K`.
  * [rust-lang/rust-bindgen](https://github.com/rust-lang/rust-bindgen) — A Rust bindings generator Stars:`3.4K`.
* Erlang
  * [rusterlium/rustler](https://github.com/rusterlium/rustler) — safe Rust bridge for creating Erlang NIF functions Stars:`3.7K`.
* Java
* Lua
* mruby
* Node.js
  * [neon-bindings/neon](https://github.com/neon-bindings/neon) — Rust bindings for writing safe and fast native Node.js modules Stars:`7.1K`.
* Objective-C
* PHP
* Python
  * [RustPython](https://github.com/RustPython/RustPython) — A Python Interpreter written in Rust [![Build Status](https://github.com/RustPython/RustPython/workflows/CI/badge.svg)](https://github.com/RustPython/RustPython/actions?query=workflow%3ACI) Stars:`14.8K`.
  * [PyO3/PyO3](https://github.com/PyO3/PyO3) — Rust bindings for the Python interpreter Stars:`8.2K`.
  * [dgrunwald/rust-cpython](https://github.com/dgrunwald/rust-cpython) — Python bindings Stars:`1.7K`.
* Ruby
* Web Assembly
  * [rustwasm/wasm-bindgen](https://github.com/rustwasm/wasm-bindgen) — A project for facilitating high-level interactions between wasm modules and JS. Stars:`6.3K`.
  * [rustwasm/wasm-pack](https://github.com/rustwasm/wasm-pack) — :package: :sparkles: pack up the wasm and publish it to npm! Stars:`5.1K`.

### Formatters

* [rustfmt](https://github.com/rust-lang/rustfmt) — Rust code formatter maintained by the Rust team and included in cargo Stars:`5.1K`.
* [dprint](https://github.com/dprint/dprint) — A pluggable and configurable code formatting platform [![build badge](https://github.com/dprint/dprint/workflows/CI/badge.svg)](https://github.com/dprint/dprint/actions?query=workflow%3ACI) Stars:`1.8K`.

### IDEs

See also [Are we (I)DE yet?](https://areweideyet.com/) and [Rust Tools](https://www.rust-lang.org/tools).

  * [lapce](https://github.com/lapce/lapce) — Lightning-fast and Powerful Code Editor written in Rust. [![build badge](https://github.com/lapce/lapce/actions/workflows/release.yml/badge.svg)](https://github.com/lapce/lapce/actions/workflows/release.yml) Stars:`25.2K`.
    * [intellij-rust/intellij-rust](https://github.com/intellij-rust/intellij-rust) — Stars:`4.4K`.
    * [autozimu/LanguageClient-neovim](https://github.com/autozimu/LanguageClient-neovim) — [LSP](https://microsoft.github.io/language-server-protocol/) client. Implemented in Rust and supports rls out of the box. Stars:`3.5K`.
    * [rust.vim](https://github.com/rust-lang/rust.vim) — provides file detection, syntax highlighting, formatting, Syntastic integration, and more. Stars:`3.5K`.
    * [rust-tools.nvim](https://github.com/simrat39/rust-tools.nvim) - Tools for better development in rust using neovim's builtin lsp Stars:`1.7K`.
  * Visual Studio

### Profiling

* [Bytehound](https://github.com/koute/bytehound) — A memory profiler for Linux Stars:`3.5K`.
* [bheisler/criterion.rs](https://github.com/bheisler/criterion.rs) — Statistics-driven benchmarking library for Rust Stars:`3.4K`.
* FlameGraphs
* [sharkdp/hyperfine](https://github.com/sharkdp/hyperfine) — A command-line benchmarking tool Stars:`15.5K`.

### Services


### Static analysis

[[assert](https://crates.io/keywords/assert), [static](https://crates.io/keywords/static)]


### Testing

[[test](https://crates.io/keywords/test), [testing](https://crates.io/keywords/testing)]

* Code Coverage
* Continuous Integration
  * [trust](https://github.com/japaric/trust) — A Travis CI and AppVeyor template to test your Rust crate on 5 architectures and publish binary releases of it for Linux, macOS and Windows Stars:`1.2K`.
* Frameworks and Runners
* Mocking and Test Data
  * [asomers/mockall](https://github.com/asomers/mockall) [[mockall](https://crates.io/crates/mockall)] — A powerful mock object library for Rust. [![Cirrus Build Status](https://api.cirrus-ci.com/github/asomers/mockall.svg)](https://cirrus-ci.com/github/asomers/mockall) Stars:`1.0K`.
* Mutation Testing
* Property Testing and Fuzzing
  * [rust-fuzz/afl.rs](https://github.com/rust-fuzz/afl.rs) — A Rust fuzzer, using [AFL](https://lcamtuf.coredump.cx/afl/) Stars:`1.4K`.

### Transpiling

* [immunant/c2rust](https://github.com/immunant/c2rust) — C to Rust translator and cross checker built atop Clang/LLVM. Stars:`3.2K`.
* [BayesWitnesses/m2cgen](https://github.com/BayesWitnesses/m2cgen) — A CLI tool to transpile trained classic machine learning models into a native Rust code with zero dependencies. [![GitHub Actions Status](https://github.com/BayesWitnesses/m2cgen/workflows/GitHub%20Actions/badge.svg?branch=master)](https://github.com/BayesWitnesses/m2cgen/actions) Stars:`2.5K`.
* [jameysharp/corrode](https://github.com/jameysharp/corrode) — A C to Rust translator written in Haskell. Stars:`2.1K`.

## Libraries


### Artificial Intelligence

#### Genetic algorithms


#### Machine learning

See [[Machine learning](https://crates.io/keywords/machine-learning)]

See also [About Rust’s Machine Learning Community](https://medium.com/@autumn_eng/about-rust-s-machine-learning-community-4cda5ec8a790#.hvkp56j3f) and [Are we learning yet?](https://www.arewelearningyet.com).

* [huggingface/tokenizers](https://github.com/huggingface/tokenizers) - Hugging Face's tokenizers for modern NLP pipelines written in Rust (original implementation) with bindings for Python. [![Build Status](https://github.com/huggingface/tokenizers/workflows/Rust/badge.svg?branch=master)](https://github.com/huggingface/tokenizers/actions) Stars:`6.8K`.
* [autumnai/leaf](https://github.com/autumnai/leaf) — Open Machine Intelligence framework.. Abandoned project. The most updated fork is [spearow/juice]( https://github.com/spearow/juice). Stars:`5.5K`.
* [tensorflow/rust](https://github.com/tensorflow/rust) — Rust language bindings for TensorFlow. Stars:`4.4K`.
* [LaurentMazare/tch-rs](https://github.com/LaurentMazare/tch-rs) — Rust language bindings for PyTorch. Stars:`2.7K`.
* [rust-ml/linfa](https://github.com/rust-ml/linfa) — Machine learning framework. Stars:`2.5K`.

#### OpenAI


### Astronomy

[[astronomy](https://crates.io/keywords/astronomy)]


### Asynchronous

* [mio](https://github.com/tokio-rs/mio) — MIO is a lightweight IO library for Rust with a focus on adding as little overhead as possible over the OS abstractions Stars:`5.5K`.
* [rust-lang/futures-rs](https://github.com/rust-lang/futures-rs) — Zero-cost futures in Rust Stars:`4.8K`.
* [Xudong-Huang/may](https://github.com/Xudong-Huang/may) — rust stackful coroutine library Stars:`1.4K`.

### Audio and Music

[[audio](https://crates.io/keywords/audio)]

  * [RustAudio/cpal](https://github.com/RustAudio/cpal) - Low-level cross-platform audio I/O library in pure Rust. [![Actions Status](https://github.com/RustAudio/cpal/workflows/cpal/badge.svg?branch=master)](https://github.com/RustAudio/cpal/actions) Stars:`1.9K`.
* [pdeljanov/Symphonia](https://github.com/pdeljanov/Symphonia) — A pure Rust audio decoding and media demuxing library supporting AAC, FLAC, MP3, MP4, OGG, Vorbis, and WAV. Stars:`1.5K`.
  * [RustAudio/rodio](https://github.com/RustAudio/rodio) — A Rust audio playback library Stars:`1.3K`.

### Authentication

* [Keats/jsonwebtoken](https://github.com/Keats/jsonwebtoken) — [JSON Web Token](https://en.wikipedia.org/wiki/JSON_Web_Token) lib in rust Stars:`1.2K`.

### Automotive


### Bioinformatics


### Caching

* [jaemk/cached](https://github.com/jaemk/cached) — Simple function caching/memoization Stars:`1.1K`.

### Cloud

* AWS [[aws](https://crates.io/keywords/aws)]
  * [awslabs/aws-lambda-rust-runtime](https://github.com/awslabs/aws-lambda-rust-runtime) [[lambda_runtime](https://crates.io/crates/lambda_runtime)] — A Rust runtime for AWS Lambda [![build badge](https://github.com/awslabs/aws-lambda-rust-runtime/workflows/Rust/badge.svg)](https://github.com/awslabs/aws-lambda-rust-runtime/actions) Stars:`2.7K`.
  * [rusoto/rusoto](https://github.com/rusoto/rusoto) — Stars:`2.6K`.
  * [awslabs/aws-sdk-rust](https://github.com/awslabs/aws-sdk-rust) - The new AWS SDK for Rust Stars:`2.2K`.
* Load Balancer
* Multi Cloud
  * [Qovery/engine](https://github.com/Qovery/engine) - Abstraction layer library that turns easy application deployment on Cloud providers in just a few minutes Stars:`2.0K`.

### Command-line

* Argument parsing
  * [clap-rs](https://github.com/clap-rs/clap) [[clap](https://crates.io/crates/clap)] — A simple to use, full featured command-line argument parser Stars:`11.2K`.
  * [TeXitoi/structopt](https://github.com/TeXitoi/structopt) [[structopt](https://crates.io/crates/structopt)] — parse command line argument by defining a struct Stars:`2.6K`.
  * [google/argh](https://github.com/google/argh) [[argh](https://crates.io/crates/argh)] — An opinionated Derive-based argument parser optimized for code size [![build badge](https://github.com/google/argh/workflows/Argh/badge.svg?branch=master)](https://github.com/google/argh/actions) Stars:`1.4K`.
* Data visualization
  * [zhiburt/tabled](https://github.com/zhiburt/tabled) [[tabled](https://crates.io/crates/tabled)] — An easy to use library for pretty print tables of Rust structs and enums. [![Build Status](https://github.com/zhiburt/tabled/actions/workflows/ci.yml/badge.svg)](https://github.com/zhiburt/tabled/actions) Stars:`1.4K`.
* Human-centered design
  * [rust-cli/human-panic](https://github.com/rust-cli/human-panic) [[human-panic](https://crates.io/crates/human-panic)] — panic messages for humans Stars:`1.2K`.
* Line editor
  * [kkawakam/rustyline](https://github.com/kkawakam/rustyline) [[rustyline](https://crates.io/crates/rustyline)] — readline implementation in Rust Stars:`1.2K`.
* Other
* Pipeline
* Progress
  * [console-rs/indicatif](https://github.com/console-rs/indicatif) [[indicatif](https://crates.io/crates/indicatif)] — indicate progress to users Stars:`3.5K`.
* Prompt
* Style
  * [mackwic/colored](https://github.com/mackwic/colored) [[colored](https://crates.io/crates/colored)] — Coloring terminal so simple, you already know how to do it! Stars:`1.3K`.
* TUI
  * BearLibTerminal
  * [fdehau/tui-rs](https://github.com/fdehau/tui-rs) [[tui](https://crates.io/crates/tui)] — A TUI library inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib) and [termui](https://github.com/gizak/termui) Stars:`10.2K`.
  * [gyscos/Cursive](https://github.com/gyscos/Cursive) [[cursive](https://crates.io/crates/cursive)] — build rich TUI applications Stars:`3.5K`.
  * ncurses
  * [redox-os/termion](https://github.com/redox-os/termion) [[termion](https://crates.io/crates/termion)] — bindless library for controlling terminals/TTY Stars:`1.9K`.
  * Termbox
  * [TimonPost/crossterm](https://github.com/crossterm-rs/crossterm) [[crossterm](https://crates.io/crates/crossterm)] — crossplatform terminal library Stars:`2.3K`.

### Compression

* bzip2
* gzip
* gzp
* miniz
* snappy
* tar
* zip

### Computation

* [dimforge/nalgebra](https://github.com/dimforge/nalgebra) — low-dimensional linear algebra library Stars:`3.2K`.
* [calebwin/emu](https://github.com/calebwin/emu) — A language for GPGPU numerical computing from a Rust macro Stars:`1.6K`.
* Parallel
* Scirust
* Statrs

### Concurrency

* [Rayon](https://github.com/rayon-rs/rayon) – A data parallelism library for Rust Stars:`8.5K`.
* [crossbeam-rs/crossbeam](https://github.com/crossbeam-rs/crossbeam) – Support for parallelism and low-level concurrency in Rust Stars:`6.0K`.

### Configuration

* [mehcode/config-rs](https://github.com/mehcode/config-rs) [[config](https://crates.io/crates/config)] — Layered configuration system for Rust applications (with strong support for 12-factor applications). Stars:`1.9K`.

### Cryptography

[[crypto](https://crates.io/keywords/crypto), [cryptography](https://crates.io/keywords/cryptography)]

* [rustls/rustls](https://github.com/rustls/rustls) — A Rust implementation of TLS Stars:`4.4K`.
* [briansmith/ring](https://github.com/briansmith/ring) — Safe, fast, small crypto using Rust and BoringSSL's cryptography primitives. Stars:`3.1K`.
* [cossacklabs/themis](https://github.com/cossacklabs/themis) [[themis](https://crates.io/crates/themis)] — a high-level cryptographic library for solving typical data security tasks, best fit for multi-platform apps. [![build badge](https://circleci.com/gh/cossacklabs/themis/tree/master.svg?style=shield)](https://app.circleci.com/pipelines/github/cossacklabs/themis) Stars:`1.7K`.
* [RustCrypto/hashes](https://github.com/RustCrypto/hashes) — Collection of cryptographic hash functions written in pure Rust Stars:`1.3K`.
* [DaGenix/rust-crypto](https://github.com/DaGenix/rust-crypto) — cryptographic algorithms in Rust Stars:`1.3K`.
* [exonum/exonum](https://github.com/exonum/exonum) [[exonum](https://crates.io/crates/exonum)] — extensible framework for blockchain projects Stars:`1.2K`.
* [sfackler/rust-openssl](https://github.com/sfackler/rust-openssl) — [OpenSSL](https://www.openssl.org/) bindings Stars:`1.1K`.

### Data processing

* [pola-rs/polars](https://github.com/pola-rs/polars) - Fast feature complete DataFrame library ![Build and test](https://github.com/pola-rs/polars/workflows/Build%20and%20test/badge.svg?branch=master) Stars:`16.6K`.
* [weld-project/weld](https://github.com/weld-project/weld) — High-performance runtime for data analytics applications Stars:`2.9K`.
* [bluss/ndarray](https://github.com/rust-ndarray/ndarray) — N-dimensional array with array views, multidimensional slicing, and efficient operations Stars:`2.7K`.

### Data streaming

* [infinyon/fluvio](https://github.com/infinyon/fluvio) - Programmable data streaming platform [![CI](https://github.com/infinyon/fluvio/workflows/CI/badge.svg?branch=stable)](https://github.com/infinyon/fluvio/actions) Stars:`1.7K`.

### Data structures

* [rust-itertools/itertools](https://github.com/rust-itertools/itertools) — Stars:`2.1K`.

### Data visualization

* [plotters](https://github.com/plotters-rs/plotters) — [![build badge](https://github.com/plotters-rs/plotters/workflows/CI/badge.svg)](https://github.com/plotters-rs/plotters/actions) Stars:`2.9K`.
* [rerun](https://github.com/rerun-io/rerun) — [[rerun](https://crates.io/crates/rerun)] — An SDK for logging computer vision and robotics data (tensors, point clouds, etc) paired with a visualizer for exploring that data over time. Stars:`2.0K`.

### Database

[[database](https://crates.io/keywords/database)]

* NoSQL [[nosql](https://crates.io/keywords/nosql)]

  * CouchDB [[couchdb](https://crates.io/keywords/couchdb)]
  * Elasticsearch [[elasticsearch](https://crates.io/keywords/elasticsearch)]
  * etcd
  * ForestDB
  * LevelDB
  * LMDB [[lmdb](https://crates.io/keywords/lmdb)]
  * MongoDB [[mongodb](https://crates.io/keywords/mongodb)]
    * [mongodb/mongo-rust-driver](https://github.com/mongodb/mongo-rust-driver) [[mongodb](https://crates.io/crates/mongodb)] — [MongoDB](https://www.mongodb.com/) bindings Stars:`1.2K`.
  * Redis [[redis](https://crates.io/keywords/redis)]
    * [surrealdb/surrealdb](https://github.com/surrealdb/surrealdb) — SurrealDB embedded document-graph database Stars:`19.7K`.
    * [redis-rs](https://github.com/redis-rs/redis-rs) — [Redis](https://redis.io/) library in Rust [![Rust](https://github.com/redis-rs/redis-rs/actions/workflows/rust.yml/badge.svg)](https://github.com/redis-rs/redis-rs/actions/workflows/rust.yml) Stars:`3.0K`.
    * [rust-rocksdb/rust-rocksdb](https://github.com/rust-rocksdb/rust-rocksdb) — RocksDB bindings [![RocksDB CI](https://github.com/rust-rocksdb/rust-rocksdb/actions/workflows/rust.yml/badge.svg?branch=master)](https://github.com/rust-rocksdb/rust-rocksdb/actions/workflows/rust.yml) Stars:`1.4K`.
* OGM [[ogm](https://crates.io/keywords/ogm)]
* ORM [[orm](https://crates.io/keywords/orm)]
  * [diesel-rs/diesel](https://github.com/diesel-rs/diesel) — an ORM and Query builder for Rust Stars:`10.3K`.
  * [SeaQL/sea-orm](https://github.com/SeaQL/sea-orm) — 🐚 An async & dynamic ORM for Rust [![crate](https://img.shields.io/crates/v/sea-orm.svg)](https://crates.io/crates/sea-orm) [![docs](https://img.shields.io/docsrs/sea-orm/latest)](https://docs.rs/sea-orm) [![build status](https://github.com/SeaQL/sea-orm/actions/workflows/rust.yml/badge.svg)](https://github.com/SeaQL/sea-orm/actions/workflows/rust.yml) Stars:`4.1K`.
  * [rbatis/rbatis](https://github.com/rbatis/rbatis) — Rust ORM Framework High Performance(JSON based) Stars:`1.8K`.
* [sfackler/r2d2](https://github.com/sfackler/r2d2) — generic connection pool Stars:`1.3K`.
  * [Brendonovich/prisma-client-rust](https://github.com/Brendonovich/prisma-client-rust) — An autogenerated query builder that provides simple and fully type-safe database access using the Prisma ecosystem. [![Test Status](https://img.shields.io/github/workflow/status/Brendonovich/prisma-client-rust/CI?label=tests&style=flat-square)](https://github.com/Brendonovich/prisma-client-rust/actions) Stars:`1.1K`.
* SQL [[sql](https://crates.io/keywords/sql)]
  * Generic
    * [launchbadge/sqlx](https://github.com/launchbadge/sqlx) - async PostgreSQL/MySQL/SQLite connection pool with strong typing support [![build badge](https://img.shields.io/github/workflow/status/launchbadge/sqlx/Rust/master?style=flat-square)](https://github.com/launchbadge/sqlx) Stars:`8.6K`.
  * Microsoft SQL
  * MySql [[mysql](https://crates.io/keywords/mysql)]
  * Oracle
  * PostgreSql [[postgres](https://crates.io/keywords/postgres), [postgresql](https://crates.io/keywords/postgresql)]
    * [sfackler/rust-postgres](https://github.com/sfackler/rust-postgres) [[postgres](https://crates.io/crates/postgres)] — A native [PostgreSQL](https://www.postgresql.org/) client Stars:`2.9K`.
  * Sqlite [[sqlite](https://crates.io/keywords/sqlite)]
    * [rusqlite](https://github.com/rusqlite/rusqlite) — [Sqlite3](https://www.sqlite.org/index.html) bindings Stars:`2.0K`.

### Date and time

[[date](https://crates.io/keywords/date), [time](https://crates.io/keywords/time)]

* [chronotope/chrono](https://github.com/chronotope/chrono) — Stars:`2.7K`.

### Distributed systems

* Antimony
* Apache Kafka
  * [fede1024/rust-rdkafka](https://github.com/fede1024/rust-rdkafka) [[rdkafka](https://crates.io/crates/rdkafka)] — [librdkafka](https://github.com/confluentinc/librdkafka) bindings Stars:`1.2K`.
* Beanstalkd
* HDFS

### Domain driven design


### Email

[[email](https://crates.io/keywords/email), [imap](https://crates.io/keywords/imap), [smtp](https://crates.io/keywords/smtp)]

* [lettre/lettre](https://github.com/lettre/lettre) — an SMTP-library for Rust [![CI](https://github.com/lettre/lettre/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/lettre/lettre/actions/workflows/test.yml) Stars:`1.4K`.

### Encoding

[[encoding](https://crates.io/keywords/encoding)]

* ASN.1
* Binary
  * [bincode-org/bincode](https://github.com/bincode-org/bincode) — A binary encoder/decoder in Rust [![CI](https://github.com/bincode-org/bincode/actions/workflows/rust.yml/badge.svg?branch=trunk)](https://github.com/bincode-org/bincode/actions/workflows/rust.yml) Stars:`2.0K`.
* BSON
* Byte swapping
* Cap'n Proto
  * [capnproto/capnproto-rust](https://github.com/capnproto/capnproto-rust) — Stars:`1.6K`.
* CBOR
* Character Encoding
* CRC
* CSV
  * [BurntSushi/rust-csv](https://github.com/BurntSushi/rust-csv) — A fast and flexible CSV reader and writer, with support for Serde Stars:`1.4K`.
* EDN
* HAR
* HTML
  * [servo/html5ever](https://github.com/servo/html5ever) — High-performance browser-grade HTML5 parser Stars:`1.8K`.
* JSON
  * [serde-rs/json](https://github.com/serde-rs/json) [[serde\_json](https://crates.io/crates/serde_json)] — JSON support for [Serde](https://github.com/serde-rs/serde) framework Stars:`3.8K`.
* MsgPack
* NetCDF
* PEM
* ProtocolBuffers
  * [tokio-rs/prost](https://github.com/tokio-rs/prost) — [![continuous integration](https://github.com/tokio-rs/prost/workflows/continuous%20integration/badge.svg?branch=master)](https://github.com/tokio-rs/prost/actions) Stars:`2.7K`.
  * [stepancheg/rust-protobuf](https://github.com/stepancheg/rust-protobuf) — Stars:`2.4K`.
* rkyv
  * [rkyv/rkyv](https://github.com/rkyv/rkyv) [[rkyv](https://crates.io/crates/rkyv)] — rkyv (archive) is a zero-copy deserialization framework for Rust Stars:`1.9K`.
* RON (Rusty Object Notation)
  * [https://github.com/ron-rs/ron](https://github.com/ron-rs/ron) — Stars:`2.5K`.
* Serde
* TOML
* XML
* YAML

### Filesystem

[[filesystem](https://crates.io/keywords/filesystem)]
* Operations
* Temporary Files
  * [zboxfs/zbox](https://github.com/zboxfs/zbox) [[zbox](https://crates.io/crates/zbox)] — Zero-details, privacy-focused embeddable file system. Stars:`1.4K`.

### Functional Programming

[[functional programming](https://crates.io/keywords/fp)]
* Prelude
  * [JasonShin/fp-core.rs](https://github.com/JasonShin/fp-core.rs) — A library for functional programming in Rust Stars:`1.2K`.

### Game development

See also [Are we game yet?](https://arewegameyet.rs)
* Allegro
* bracket-lib (previously RLTK)
  * [bracket-lib](https://github.com/amethyst/bracket-lib) [[bracket-lib](https://crates.io/crates/bracket-lib)] - The Roguelike Toolkit (RLTK), implemented for Rust. [![Rust](https://github.com/amethyst/bracket-lib/actions/workflows/rust.yml/badge.svg)](https://github.com/amethyst/bracket-lib/actions/workflows/rust.yml) Stars:`1.2K`.
* Challonge
* Corange
* Entity-Component Systems (ECS)
  * [amethyst/specs](https://github.com/amethyst/specs) — Specs Parallel ECS Stars:`2.2K`.
  * [legion](https://github.com/amethyst/legion) — A feature rich high performance ECS library with minimal boilerplate [![build badge](https://github.com/amethyst/legion/workflows/CI/badge.svg?branch=master)](https://github.com/amethyst/legion/actions) Stars:`1.5K`.
* Game Engines
  * [Bevy](https://github.com/bevyengine/bevy) is a refreshingly simple data-driven game engine built in Rust. - [![Crates.io](https://img.shields.io/crates/v/bevy.svg)](https://crates.io/crates/bevy) Stars:`23.3K`.
[![Crates.io](https://img.shields.io/crates/d/bevy.svg)](https://crates.io/crates/bevy)
  * [ggez](https://github.com/ggez/ggez) — A lightweight game framework for making 2D games with minimum friction - [![Crates.io](https://img.shields.io/crates/v/ggez.svg)](https://crates.io/crates/ggez) [![license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/ggez/ggez/blob/master/LICENSE) [![Crates.io](https://img.shields.io/crates/d/ggez.svg)](https://crates.io/crates/ggez) Stars:`3.7K`.
  * [godot-rust/gdnative](https://github.com/godot-rust/gdnative) [[gdnative](https://crates.io/crates/gdnative)] - Rust bindings to the Godot game engine [![CI](https://github.com/godot-rust/gdnative/actions/workflows/full-ci.yml/badge.svg)](https://github.com/godot-rust/gdnative/actions/workflows/full-ci.yml) Stars:`3.3K`.
  * [Rust-SDL2/rust-sdl2](https://github.com/Rust-SDL2/rust-sdl2) — SDL2 bindings Stars:`2.3K`.
* SFML
* Skillratings
* Tcod-rs
  * Warning: Not maintained anymore
* Toornament-rs
* Victorem

### Geospatial

[[geo](https://crates.io/keywords/geo), [gis](https://crates.io/keywords/gis)]


### Graph algorithms

* [petgraph/petgraph](https://github.com/petgraph/petgraph) - Graph data structure library for Rust. [![graph CI status](https://github.com/petgraph/petgraph/workflows/Continuous%20integration/badge.svg?branch=master)](https://github.com/petgraph/petgraph/actions/workflows/ci.yml) Stars:`2.1K`.

### Graphics

[[graphics](https://crates.io/keywords/graphics)]

* Font
* [gfx-rs/wgpu](https://github.com/gfx-rs/wgpu) - Native WebGPU implementation based on gfx-hal. [![build badge](https://github.com/gfx-rs/wgpu/workflows/CI/badge.svg?branch=master)](https://github.com/gfx-rs/wgpu/actions) Stars:`7.4K`.
* [gfx-rs/gfx](https://github.com/gfx-rs/gfx) — A high-performance, bindless graphics API for Rust. Stars:`5.3K`.
* OpenGL [[opengl](https://crates.io/keywords/opengl)]
  * [glium/glium](https://github.com/glium/glium) — safe OpenGL wrapper for the Rust language. Stars:`3.2K`.
* PDF
  * [vulkano](https://github.com/vulkano-rs/vulkano) [[vulkano](https://crates.io/crates/vulkano)] — Stars:`3.8K`.
  * [J-F-Liu/lopdf](https://github.com/J-F-Liu/lopdf) — PDF document manipulation Stars:`1.2K`.

### GUI

[[gui](https://crates.io/keywords/gui)]

* Cocoa
* [tauri-apps/tauri](https://github.com/tauri-apps/tauri) — Build smaller, faster, and more secure desktop applications with a web frontend, powered by [WRY](https://github.com/tauri-apps/wry). [![test library](https://img.shields.io/github/workflow/status/tauri-apps/tauri/test%20library?label=test%20library)](https://github.com/tauri-apps/tauri/actions?query=workflow%3A%22test+library%22) Stars:`62.5K`.
* [ImGui](https://github.com/ocornut/imgui) Stars:`46.9K`.
* [iced-rs/iced](https://github.com/iced-rs/iced) [[iced](https://crates.io/crates/iced)] — A cross-platform GUI library for Rust focused on simplicity and type-safety. Inspired by Elm. Stars:`18.8K`.
* [emilk/egui](https://github.com/emilk/egui) - Simple, fast, and highly portable immediate mode GUI library for Rust. egui runs on the web, natively, and in your favorite game engine. [![Build Status](https://github.com/emilk/egui/workflows/CI/badge.svg)](https://github.com/emilk/egui/actions?workflow=CI) Stars:`14.6K`.
* [libui](https://github.com/andlabs/libui) Stars:`10.4K`.
* [DioxusLabs/dioxus](https://github.com/dioxuslabs/dioxus) - a portable, performant, and ergonomic framework for building cross-platform user interfaces in Rust. ![rust ci](https://github.com/dioxuslabs/dioxus/actions/workflows/main.yml/badge.svg) Stars:`8.6K`.
* [slint-ui/slint](https://github.com/slint-ui/slint) [[slint](https://crates.io/crates/slint)] — [Slint](https://slint-ui.com) is a toolkit to efficiently develop fluid graphical user interfaces for embedded devices and desktop applications. [![Build Status](https://github.com/slint-ui/slint/workflows/CI/badge.svg?branch=master)](https://github.com/slint-ui/slint/actions?query=workflow%3ACI) Stars:`7.8K`.
* [Nuklear](https://github.com/Immediate-Mode-UI/Nuklear) Stars:`7.3K`.
* [fschutt/azul](https://github.com/fschutt/azul) — A free, functional, IMGUI-oriented GUI framework for rapid development of desktop applications written in Rust, supported by the Mozilla WebRender rendering engine. Stars:`5.5K`.
* [OrbTk](https://github.com/redox-os/orbtk) — The Orbital Widget Toolkit is a multi platform (G)UI toolkit using SDL2 [![Build and test](https://github.com/redox-os/orbtk/workflows/build/badge.svg?branch=develop)](https://github.com/redox-os/orbtk/actions) Stars:`3.8K`.
  * [fzyzcjy/flutter_rust_bridge](https://github.com/fzyzcjy/flutter_rust_bridge) — High-level memory-safe binding generator for Flutter/Dart <-> Rust Stars:`2.5K`.
* [tauri-apps/wry](https://github.com/tauri-apps/wry) - Webview Rendering librarY. Stars:`2.5K`.
  * [relm](https://github.com/antoyo/relm) — Asynchronous, GTK+-based, GUI library, inspired by Elm Stars:`2.3K`.
  * [imgui-rs](https://github.com/imgui-rs/imgui-rs) — Rust bindings for ImGui [![Build Status](https://github.com/imgui-rs/imgui-rs/workflows/ci/badge.svg?branch=master)](https://github.com/imgui-rs/imgui-rs/actions) Stars:`2.2K`.
  * [flutter-rs](https://github.com/flutter-rs/flutter-rs) — Build flutter desktop app in dart & rust. Stars:`2.0K`.
* [xilem](https://github.com/linebender/xilem) — Successor of the data-first Rust-native UI design toolkit [druid](https://github.com/linebender/druid). Stars:`1.4K`.
  * [fltk-rs](https://github.com/fltk-rs/fltk-rs) — FLTK Rust bindings [![Build](https://github.com/fltk-rs/fltk-rs/workflows/Build/badge.svg?branch=master)](https://github.com/fltk-rs/fltk-rs/actions) Stars:`1.2K`.

### Image processing

* [image-rs/image](https://github.com/image-rs/image) — Basic imaging processing functions and methods for converting to and from image formats Stars:`3.8K`.
* [twistedfall/opencv-rust](https://github.com/twistedfall/opencv-rust) — Rust bindings for OpenCV Stars:`1.4K`.

### Language specification


### Logging

[[log](https://crates.io/keywords/log)]

* [tokio-rs/tracing](https://github.com/tokio-rs/tracing) — An application level tracing framework for async-aware structured logging, error handling, metrics, and more [![Build Status](https://github.com/tokio-rs/tracing/workflows/CI/badge.svg?branch=master)](https://github.com/tokio-rs/tracing/actions?query=workflow%3ACI) Stars:`3.8K`.
* [rust-lang/log](https://github.com/rust-lang/log) — Logging implementation for Rust Stars:`1.7K`.
* [slog-rs/slog](https://github.com/slog-rs/slog) — Structured, composable logging for Rust Stars:`1.4K`.

### Macro

* cute

### Markup language

* CommonMark
  * [raphlinus/pulldown-cmark](https://github.com/raphlinus/pulldown-cmark) — [CommonMark](https://commonmark.org/) parser in Rust Stars:`1.6K`.

### Mobile

* Android / iOS
* Generic
* iOS

### Network programming

* Bluetooth
* CoAP
* Docker
* FTP
* gRPC
  * [hyperium/tonic](https://github.com/hyperium/tonic) — A native gRPC client & server implementation with async/await support [![Crates.io](https://img.shields.io/crates/v/tonic)](https://crates.io/crates/tonic) Stars:`7.2K`.
  * [tikv/grpc-rs](https://github.com/tikv/grpc-rs) — The gRPC library for Rust built on C Core library and futures Stars:`1.7K`.
* HTTP
  * [Hurl](https://github.com/Orange-OpenSource/hurl) — Run and test HTTP requests with plain text and libcurl [![CI](https://github.com/Orange-OpenSource/hurl/workflows/CI/badge.svg)](https://github.com/Orange-OpenSource/hurl/actions) Stars:`4.0K`.
* IPNetwork
* Low level
  * [tokio-rs/tokio](https://github.com/tokio-rs/tokio) — A network application framework for rapid development and highly scalable production deployments of clients and servers. Stars:`20.2K`.
  * [actix/actix](https://github.com/actix/actix) — Actor library for Rust Stars:`7.8K`.
  * [smoltcp-rs/smoltcp](https://github.com/smoltcp-rs/smoltcp) — A standalone, event-driven TCP/IP stack that is designed for bare-metal, real-time systems Stars:`3.0K`.
  * [libpnet/libpnet](https://github.com/libpnet/libpnet) — A cross-platform, low level networking Stars:`1.9K`.
* message-io
* MQTT
* NanoMsg
* NATS
* Nng
* NNTP
* P2P
  * [libp2p/rust-libp2p](https://github.com/libp2p/rust-libp2p) — The Rust Implementation of libp2p networking stack. [![Circle CI](https://circleci.com/gh/libp2p/rust-libp2p.svg?style=svg)](https://app.circleci.com/pipelines/github/libp2p/rust-libp2p) Stars:`3.2K`.
* POP3
* QUIC
  * [cloudflare/quiche](https://github.com/cloudflare/quiche) — cloudflare implementation of the QUIC transport protocol and HTTP/3 ![build](https://img.shields.io/github/workflow/status/cloudflare/quiche/Stable) Stars:`7.5K`.
  * [quinn-rs/quinn](https://github.com/quinn-rs/quinn) — Futures-based QUIC implementation in Rust [![build badge](https://dev.azure.com/dochtman/Projects/_apis/build/status/Quinn?branchName=master)](https://dev.azure.com/dochtman/Projects/_build) Stars:`2.8K`.
  * [mozilla/neqo](https://github.com/mozilla/neqo) — an Implementation of QUIC written in Rust Stars:`1.6K`.
* Raknet
* RPC
* Socket.io
* SSH
* Stomp
* ZeroMQ

### Parsing

  * [rust-bakery/nom](https://github.com/rust-bakery/nom) — parser combinator library Stars:`8.0K`.
  * [pest-parser/pest](https://github.com/pest-parser/pest) — The Elegant Parser Stars:`3.8K`.
  * [lalrpop/lalrpop](https://github.com/lalrpop/lalrpop) — LR(1) parser generator for Rust Stars:`2.5K`.
  * [kevinmehall/rust-peg](https://github.com/kevinmehall/rust-peg) — Parsing Expression Grammar (PEG) parser generator Stars:`1.3K`.
  * [Marwes/combine](https://github.com/Marwes/combine) — parser combinator library Stars:`1.2K`.

### Peripherals

* Serial Port

### Platform specific

* Cross-platform
* FreeBSD
* Linux
* Unix-like
  * [nix-rust/nix](https://github.com/nix-rust/nix) — Unix-like API bindings [![Cirrus Build Status](https://api.cirrus-ci.com/github/nix-rust/nix.svg)](https://cirrus-ci.com/github/nix-rust/nix) Stars:`2.1K`.
* Windows
  * [retep998/winapi-rs](https://github.com/retep998/winapi-rs) — Windows API bindings [![Rust](https://github.com/retep998/winapi-rs/actions/workflows/rust.yml/badge.svg?branch=dev)](https://github.com/retep998/winapi-rs/actions/workflows/rust.yml) Stars:`1.6K`.

### Scripting

[[scripting](https://crates.io/keywords/scripting)]

* [gluon-lang/gluon](https://github.com/gluon-lang/gluon) —  A small, statically-typed, functional programming language Stars:`2.8K`.
* [rhaiscript/rhai](https://github.com/rhaiscript/rhai) — A tiny and fast embedded scripting language resembling a combination of JavaScript and Rust [![build badge](https://github.com/rhaiscript/rhai/workflows/Build/badge.svg)](https://github.com/rhaiscript/rhai/actions) Stars:`2.7K`.
* [PistonDevelopers/dyon](https://github.com/PistonDevelopers/dyon) — A rusty dynamically typed scripting language Stars:`1.6K`.
* [mun](https://github.com/mun-lang/mun) — A compiled, statically-typed scripting language with first class hot reloading support Stars:`1.5K`.
* [metacall/core](https://github.com/metacall/core) [[metacall](https://crates.io/crates/metacall)] — Cross-platform Polyglot Runtime which supports NodeJS, JavaScript, TypeScript, Python, Ruby, C#, Wasm, Java, Cobol and more. [![build badge](https://gitlab.com/metacall/core/badges/master/pipeline.svg)](https://gitlab.com/metacall/core) Stars:`1.3K`.
* [rune-rs/rune](https://github.com/rune-rs/rune) — An embeddable dynamic programming language for Rust Stars:`1.2K`.

### Simulation

[[simulation](https://crates.io/keywords/simulation)]


### System

* [GuillaumeGomez/sysinfo](https://github.com/GuillaumeGomez/sysinfo) [[sysinfo](https://crates.io/crates/sysinfo)] — Cross-platform library to fetch system information [![build badge](https://github.com/GuillaumeGomez/sysinfo/actions/workflows/CI.yml/badge.svg?branch=master)](https://github.com/GuillaumeGomez/sysinfo/actions/workflows/CI.yml) Stars:`1.3K`.

### Task scheduling

[![Build](https://github.com/BinChengZhao/delay-timer/actions/workflows/rust.yml/badge.svg)](
https://github.com/BinChengZhao/delay-timer/actions)

### Template engine

* Handlebars
  * [sunng87/handlebars-rust](https://github.com/sunng87/handlebars-rust) — Handlebars template engine with inheritance, custom helper support. Stars:`1.0K`.
* HTML
  * [Keats/tera](https://github.com/Keats/tera) — template engine based on Jinja2 and the Django template language. [![Actions Status](https://github.com/Keats/tera/workflows/ci/badge.svg?branch=master)](https://github.com/Keats/tera/actions) Stars:`2.7K`.
  * [djc/askama](https://github.com/djc/askama) — template rendering engine based on Jinja Stars:`2.1K`.
  * [lambda-fairy/maud](https://github.com/lambda-fairy/maud) — compile-time HTML templates Stars:`1.5K`.
* Mustache

### Text processing

* [rust-lang/regex](https://github.com/rust-lang/regex) — Regular expressions (RE2 style) Stars:`2.8K`.

### Text search

* [meilisearch/MeiliSearch](https://github.com/meilisearch/MeiliSearch) — Ultra relevant, instant and typo-tolerant full-text search API. [![Build Status](https://github.com/meilisearch/MeiliSearch/workflows/Cargo%20test/badge.svg?branch=master)](https://github.com/meilisearch/MeiliSearch/actions) Stars:`35.9K`.
* [tantivy](https://github.com/quickwit-oss/tantivy) [[tantivy](https://crates.io/crates/tantivy)] — A horse-speed full-text search engine library written in Rust. [![Build Status](https://github.com/quickwit-oss/tantivy/actions/workflows/test.yml/badge.svg)](https://github.com/quickwit-oss/tantivy/actions/workflows/test.yml) Stars:`8.0K`.
* [BurntSushi/fst](https://github.com/BurntSushi/fst) [[fst](https://crates.io/crates/fst)] — Stars:`1.6K`.

### Unsafe


### Video


### Virtualization

* [bytecodealliance/wasmtime](https://github.com/bytecodealliance/wasmtime) — A standalone runtime for WebAssembly [![Build Status](https://github.com/bytecodealliance/wasmtime/workflows/CI/badge.svg)](https://github.com/bytecodealliance/wasmtime/actions?query=workflow%3ACI) Stars:`12.0K`.

### Web programming

See also [Are we web yet?](https://www.arewewebyet.org) and [Rust web framework comparison](https://github.com/flosse/rust-web-framework-comparison).

* Client-side / WASM
  * [leptos](https://github.com/leptos-rs/leptos) — Leptos is a full-stack, isomorphic Rust web framework leveraging fine-grained reactivity to build declarative user interfaces.[![crate](https://img.shields.io/crates/v/create-rust-app.svg)](https://crates.io/crates/leptos) Stars:`8.2K`.
  * [sauron](https://github.com/ivanceras/sauron) - Client side web framework which closely adheres to The Elm Architecture. Stars:`1.8K`.
* HTTP Client
  * [hyperium/hyper](https://github.com/hyperium/hyper) — an HTTP implementation [![CI](https://github.com/hyperium/hyper/workflows/CI/badge.svg?branch=master)](https://github.com/hyperium/hyper/actions?query=workflow%3ACI) Stars:`11.7K`.
  * [seanmonstar/reqwest](https://github.com/seanmonstar/reqwest) — an ergonomic HTTP Client for Rust. Stars:`7.4K`.
  * [async-graphql](https://github.com/async-graphql/async-graphql) - A GraphQL server library implemented in Rust [![Build Status](https://dev.azure.com/graphql-rust/GraphQL%20Rust/_apis/build/status/graphql-rust.juniper)](https://dev.azure.com/graphql-rust/GraphQL%20Rust/_build/latest?definitionId=1) Stars:`2.8K`.
* HTTP Server
  * [Rocket](https://github.com/SergioBenitez/Rocket) — Rocket is web framework for Rust (nightly) with a focus on ease-of-use, expressability, and speed Stars:`20.4K`.
  * [actix/actix-web](https://github.com/actix/actix-web) — A lightweight async web framework for Rust with websocket support Stars:`17.3K`.
  * [hyperium/hyper](https://github.com/hyperium/hyper) — an HTTP implementation [![CI](https://github.com/hyperium/hyper/workflows/CI/badge.svg?branch=master)](https://github.com/hyperium/hyper/actions?query=workflow%3ACI) Stars:`11.7K`.
  * [tokio/axum](https://github.com/tokio-rs/axum) - Ergonomic and modular web framework built with Tokio, Tower, and Hyper [![Build badge](https://github.com/tokio-rs/axum/actions/workflows/CI.yml/badge.svg?branch=main)](https://github.com/tokio-rs/axum/actions/workflows/CI.yml) Stars:`9.9K`.
  * [seanmonstar/warp](https://github.com/seanmonstar/warp) — A super-easy, composable, web server framework for warp speeds. [![crate](https://img.shields.io/crates/v/create-rust-app.svg)](https://crates.io/crates/warp) Stars:`8.0K`.
  * [Iron](https://github.com/iron/iron) — A middleware-based server framework Stars:`6.1K`.
  * [Juniper](https://github.com/graphql-rust/juniper) — GraphQL server library for Rust Stars:`5.1K`.
  * [poem-web/poem](https://github.com/poem-web/poem) - A full-featured and easy-to-use web framework with the Rust programming language. [![CI](https://github.com/poem-web/poem/actions/workflows/ci.yml/badge.svg)](https://github.com/poem-web/poem/actions/workflows/ci.yml) Stars:`2.4K`.
  * [Gotham](https://github.com/gotham-rs/gotham) — A flexible web framework that does not sacrifice safety, security or speed. Stars:`2.1K`.
  * [Salvo](https://github.com/salvo-rs/salvo) — an easy to use webframework base on hyper and tokio. [![build build](https://github.com/salvo-rs/salvo/workflows/CI%20(Linux)/badge.svg?branch=master&event=push)](https://github.com/salvo-rs/salvo/actions) Stars:`1.4K`.
  * [handlebars-rust](https://github.com/sunng87/handlebars-rust) — an Iron web framework middleware. Stars:`1.0K`.
* Miscellaneous
  * [serenity-rs/serenity](https://github.com/serenity-rs/serenity) [[serenity](https://crates.io/crates/serenity)] - A Rust library for the Discord API Stars:`3.6K`.
  * [osohq/oso](https://github.com/osohq/oso) [[oso](https://crates.io/crates/oso)] - A policy engine for authorization that's embedded in your application. [![Build Status](https://github.com/osohq/oso/workflows/Development/badge.svg?branch=main)](https://github.com/osohq/oso/actions?query=branch%3Amain+workflow%3ADevelopment) Stars:`3.0K`.
  * [svix/svix-webhooks](https://github.com/svix/svix-webhooks) [[svix](https://crates.io/crates/svix)]- A library for sending webhooks and verifying signatures. Stars:`1.4K`.
  * [causal-agent/scraper](https://github.com/causal-agent/scraper) [[scraper](https://crates.io/crates/scraper)] - HTML parsing and querying with CSS selectors. [![Build Status](https://github.com/causal-agent/scraper/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/causal-agent/scraper/actions) Stars:`1.4K`.
  * [pyrossh/rust-embed](https://github.com/pyrossh/rust-embed) — A macro to embed static assets into the rust binary Stars:`1.1K`.
* Reverse Proxy
  * [sozu-proxy/sozu](https://github.com/sozu-proxy/sozu) [[sozu](https://crates.io/crates/sozu)] — A HTTP reverse proxy. [![CI](https://github.com/sozu-proxy/sozu/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/sozu-proxy/sozu/actions/workflows/ci.yml) Stars:`2.3K`.
* Static Site Generators
  * [getzola/zola](https://github.com/getzola/zola) [[zola](https://www.getzola.org/)] — An opinionated static site generator with everything built-in. [![Build Status](https://dev.azure.com/getzola/zola/_apis/build/status/getzola.zola?branchName=master)](https://dev.azure.com/getzola/zola/_build) Stars:`10.8K`.
  * [vi/websocat](https://github.com/vi/websocat) — CLI for interacting with WebSockets, with functionality of Netcat, Curl and Socat. Stars:`5.4K`.
  * [snapview/tungstenite-rs](https://github.com/snapview/tungstenite-rs) — Lightweight stream-based WebSocket implementation for Rust. Stars:`1.4K`.
  * [housleyjk/ws-rs](https://github.com/housleyjk/ws-rs) — lightweight, event-driven WebSockets for Rust Stars:`1.4K`.
  * [rust-websocket](https://github.com/websockets-rs/rust-websocket) — A framework for dealing with WebSocket connections (both clients and servers) Stars:`1.3K`.
  * [cobalt-org/cobalt.rs](https://github.com/cobalt-org/cobalt.rs) — Static site generator written in Rust [![Build Status](https://dev.azure.com/cobalt-org/cobalt-org/_apis/build/status/cobalt.rs?branchName=master)](https://dev.azure.com/cobalt-org/cobalt-org/_build?definitionId=2) Stars:`1.2K`.

## Registries

A registry allows you to publish your Rust libraries as crate packages, to share them with others publicly and privately.


## Resources

* Benchmarks
* Decks & Presentations
* Learning
  * [Rustlings](https://github.com/rust-lang/rustlings) — small exercises to get you used to reading and writing Rust code Stars:`37.3K`.
  * [rust-learning](https://github.com/ctjhoa/rust-learning) — A collection of useful resources to learn Rust Stars:`9.7K`.
  * [Easy Rust](https://github.com/Dhghomon/easy_rust) - Learn Rust in easy English. Stars:`7.5K`.
  * [Idiomatic Rust](https://github.com/mre/idiomatic-rust) —  A peer-reviewed collection of articles/talks/repos which teach idiomatic Rust. Stars:`4.6K`.
  * [stdx](https://github.com/brson/stdx) — Learn these crates first as an extension to std Stars:`1.9K`.
* Podcasts
* [Rust Design Patterns](https://github.com/rust-unofficial/patterns) Stars:`6.9K`.
* [RustBooks](https://github.com/sger/RustBooks) — list of RustBooks Stars:`3.2K`.
* [RustViz](https://github.com/rustviz/rustviz) — generates visualizations from simple Rust programs to assist users in better understanding the Rust Lifetime and Borrowing mechanism. Stars:`2.4K`.

## License

[![CC0](https://licensebuttons.net/p/zero/1.0/88x31.png)](https://creativecommons.org/publicdomain/zero/1.0/)
