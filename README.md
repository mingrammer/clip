<br><br>

<h1 align="center">Clip</h1>

<p align="center">
  <a href="/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg"/></a>
  <a href="https://goreportcard.com/report/github.com/mingrammer/clip"><img src="https://goreportcard.com/badge/github.com/mingrammer/clip"/></a>
</p>

<p align="center">
A simple key-value store for clipboard
</p>

<br><br><br>

It is a command line tool for simple key-value store for clipboard. I often use frequently used text, especially long texts that are hard to remember, on the clipboard. But it is cumbersome to copy the text manually by using trackpad or mouse. So `clip` was developed to solve this problem.

# Installation

```bash
go get github.com/mingrammer/clip
```

Or using [Homebrew](https://brew.sh)

```bash
brew tap mingrammer/clip https://github.com/mingrammer/clip.git
brew install clip
```

# Usage

```bash
# List all key-value pairs
clip list

# Get a value of a specific key
clip get <key>

# Set a key-value pair
clip set <key> <value>

# Delete a key
clip del <key>

# Copy a value of a specific key to clipboard
clip cp <key>
```

# License

Details on [LICENSE](/LICENSE)
