# lazynix

A Terminal User Interface (TUI) for Nix, built with [Bubble Tea](https://github.com/charmbracelet/bubbletea).

> **Note**: This project is currently a work in progress.

## Overview

lazynix aims to provide a simple and efficient terminal interface for managing Nix configurations and operations.

## Prerequisites

- [Go](https://go.dev/) (1.25 or later)
- [Nix](https://nixos.org/) (optional, for development environment)

## Getting Started

### Running from source

1. Clone the repository:
   ```bash
   git clone https://github.com/Foroughi/lazynix.git
   cd lazynix
   ```

2. Run the application:
   ```bash
   go run src/main.go
   ```

## Development

This project uses Nix flakes for the development environment.

To enter the development shell:
```bash
nix develop
```
