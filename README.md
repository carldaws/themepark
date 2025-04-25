# Themepark

> A fast, minimal universal theme switcher for your terminal and apps.

Themepark generates theme configuration files from built-in templates and themes.
It outputs config files to `~/.themepark/` so programs like Ghostty (and soon Neovim, Wezterm, etc.) can use a single, consistent theme — automatically.

## ✨ Features

- 🛠️ Single static binary — no setup, no runtime dependencies
- 🎨 Built-in themes and templates — nothing to install
- 🚀 Fast theme switching — just a single CLI command
- 🗂️ Clean output — only generates necessary files in ~/.themepark/
- 📦 Ready for expansion — easily add support for more tools in the future

## 📦 Installation

Clone and build:

```bash
git clone https://github.com/carldaws/themepark.git
cd themepark
go build -o themepark
```
Binaries coming soon!

## 🚀 Usage

### Switch to a theme

```bash
themepark use <theme-name>
```

This generates the appropriate config file(s) in `~/.themepark/`.

### List available themes

```bash
themepark list
```

Example output:

```diff
Available themes:
- gruvbox
- solarized
- dracula
```

### Find where a generated file is

```bash
themepark where ghostty
```
Example output:

```diff
/Users/yourname/.themepark/ghostty.conf
```
You can then configure Ghostty to load that path directly.
