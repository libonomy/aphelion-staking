<br />
<p align="center">
  <h3 align="center">Libonomy Aphelion</h3>

Core consensus module responsible for processing blockchain data and validating the behaviour of blockchain.


<b>Halt Status:</b> Development was halted: Remained in progress for R&D work along with the implementation of the Layer 0 cross chain modules for ETH & BTC and AI simulators integration and autonomous agents.


</p>

## Getting Started

Follow the instructions below to run the project locally.

## Prerequisites

Make sure you have the following installed:

- **GO**

  Install GO from:

  ```sh
  https://go.dev/doc/install
  ```

  ⚠️ **Supported GO version:** `1.12.x`


## Installation

1. **Clone the repository**

2. **Navigate to the project directory**

3. **Install dependencies**

```sh
go mod tidy
```

4. **Make build**

```sh
make install
```

After building process is complete you can use the consensus amd its functions independently. It is also a sub package used in cuspd build as the consensus engine for performing all decision based operations.

```
Copyright © 2025 — Libonomy
```