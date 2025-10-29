

# CommitFeed
---
![GitHub Release](https://img.shields.io/github/v/release/kurtiz/commit-feed)
![GitHub commit activity](https://img.shields.io/github/commit-activity/t/kurtiz/commit-feed)
[![wakatime](https://wakatime.com/badge/user/9657174f-2430-4dfd-aaef-2b316eb71a36/project/e422536c-8bfe-4432-a7d6-42811dc51487.svg)](https://wakatime.com/badge/user/9657174f-2430-4dfd-aaef-2b316eb71a36/project/e422536c-8bfe-4432-a7d6-42811dc51487)


> ✨ *Turn your Git commits into engaging social media updates — automatically.*

**CommitFeed** is a command-line tool written in Go that reads your Git commit history, summarizes recent changes, and uses AI to generate ready-to-post content for platforms like **LinkedIn** and **Twitter/X**.

Perfect for open-source maintainers, indie hackers, or dev teams who want to share progress updates directly from their terminal.


## ⚡Features

* 🪄 **AI-powered post generation** — uses Hugging Face (or any compatible LLM) to craft natural, developer-friendly posts.
* 🧾 **Reads real Git history** — pulls your recent commits and formats them into summaries.
* 🌍 **Multi-platform support** — generates platform-optimized versions for LinkedIn and Twitter.
* ⚙️ **Configurable AI providers** — choose between Hugging Face, OpenAI, Gemini, DeepSeek, or Grok.
* 🏡 **First-time setup wizard** — built with [Charm’s BubbleTea](https://github.com/charmbracelet/bubbletea) for a smooth CLI experience.
* 🔐 **Secure local config** — stores your API keys safely in `~/.commit-feed/config.json`. (_plans in place to encrypt the keys_)
* 🧩 **Post automation** — optionally publish posts directly with the `--post` flag (coming soon).


## 📦 Installation

### From binary

[Download the latest binary for your platform](https://github.com/kurtiz/commit-feed/releases/latest).

```bash
chmod +x commitfeed
./commitfeed generate
```



### Quick Install



```bash
curl -fsSL https://raw.githubusercontent.com/kurtiz/commit-feed/main/install.sh | bash
```
<img src="./assets/install.gif">

### From source

You’ll need **Go 1.21+** installed.

```bash
git clone https://github.com/kurtiz/commit-feed
cd commit-feed
go build -o commitfeed .
```

Then run it:

```bash
./commitfeed generate
```

Or install directly into your `$GOPATH/bin`:

```bash
go install github.com/kurtiz/commit-feed@latest
```

## 🧰 First Run Setup

When you run `commitfeed` for the first time, it launches a **beautiful interactive setup wizard** that guides you through:

1. Selecting your preferred AI provider (Hugging Face, OpenAI, Gemini, DeepSeek, Grok, etc.)
2. Providing your API key (or skipping for default free mode)
3. Saving preferences to `~/.commit-feed/config.json`

Example config file:

```json
{
  "provider": "huggingface",
  "api_key": "API_KEY",
  "default_platforms": [
    "linkedin",
    "twitter"
  ]
}
```

## 💻 Usage

### Basic command

```bash
commitfeed generate
```

This scans your latest commits and generates two posts — one for **LinkedIn** and one for **Twitter**.

Output example:

```
📦 Using AI Provider: huggingface
📰 Target Platforms: [linkedin twitter]

✅ Generated Posts:
🔗 LinkedIn:
We just added a .gitignore to keep our repo tidy and introduced a Git log reader feature that pulls commit history straight into your app. Plus, a brand-new README gives an overview, build steps, and quick usage examples. Happy hacking!

🐦 Twitter:
Just dropped a .gitignore + a Git log reader 📚 + a fresh README with build & usage steps. Clean repo, ready-to-go docs, and instant commit history in your app. 🚀 #devtools #opensource

💡 Preview only (not posted). Use --post to share automatically.
```


### 🗃️ Commands

| Flag          | Description                                           | Example                      |
| ------------- | ----------------------------------------------------- | ---------------------------- |
| `generate`    | Generates posts for the latest commits                | `commitfeed generate`        |
| `init`        | Initializes your config file                          | `commitfeed init`            |

### 🎛️ Generate flags/Options

| Flag          | Description                                           | Example                      |
| ------------- | ----------------------------------------------------- | ---------------------------- |
| `--platforms` | Specify target platforms (`linkedin,twitter`)         | `--platforms=twitter`        |
| `--range`     | Specify commit range                                  | `--range HEAD~5..HEAD`       |
| `--post, -p`  | Automatically publish generated posts *(coming soon)* | `--post`                     |
| `--help`      | Show all available options                            | `commitfeed generate --help` |

---

## ⚙️ Configuration

CommitFeed reads its config from:

```bash
~/.commit-feed/config.json
```

You can edit it manually or re-run the setup wizard:

```bash
commitfeed init
```

---

## 🧩 Project Structure

```
commit-feed/
├── cmd/                  # Cobra command definitions
│   ├── root.go
│   ├── generate.go
│   └── init.go
├── internal/
│   ├── ai/               # AI provider logic (Hugging Face, OpenAI, etc.)
│   ├── git/              # Git log parsing utilities
│   └── config/           # Config loader & setup wizard
├── main.go
├── go.mod
└── README.md
```

---

## 🔌 Supported AI Providers

| Provider            | Model Example             | Free Tier | Notes                       |
| ------------------- | ------------------------- | --------- | --------------------------- |
| **Hugging Face**    | `openai/gpt-oss-20b:groq` | ✅ Yes     | Default option              |
| **OpenAI**          | `gpt-4-turbo`             | ❌ Paid    | Needs OpenAI API key        |
| **Gemini (Google)** | `gemini-1.5-pro`          | ✅ Limited | Requires Google Cloud setup |
| **DeepSeek**        | `deepseek-coder`          | ✅ Yes     | Great for dev summaries     |
| **Grok (xAI)**      | `grok-1`                  | ❌ Paid    | Requires xAI API key        |

---

## 🧠 How It Works

1. CommitFeed checks that you’re in a valid Git repository.
2. It extracts recent commits with author, date, and message.
3. The commit messages are formatted into an AI prompt.
4. The selected AI model generates short social media posts.
5. The output is displayed (or optionally posted).

---

## 🧪 Development

```bash
# Run locally
go run main.go generate

# Add new subcommands
cobra-cli add <command>

# Clean dependencies
go mod tidy
```

---

## 🤝 Contributing

Pull requests are welcome!
If you’d like to add support for another AI provider or posting platform, open an issue or submit a PR.

### Ideas

* Support for Mastodon, Threads, or Bluesky.
* Markdown-to-Post formatter.
* Scheduling & auto-posting.

---

## 🪪 License

MIT License © 2025 [Your Name / GitHub handle]
See the [LICENSE](./LICENSE) file for details.

---

## 🌟 Acknowledgments

CommitFeed is powered by:

* [Cobra](https://github.com/spf13/cobra) – CLI framework for Go
* [Charm](https://charm.sh) – for terminal UI magic
* [Hugging Face Inference API](https://huggingface.co/inference-api)
* [OpenAI-Compatible Router](https://huggingface.co/docs/api-inference/openai_compatibility)



curl -fsSL https://raw.githubusercontent.com/kurtiz/commit-feed/main/install.sh | bash

cd ~/Documents/node/bvault-js
commitfeed init
commitfeed generate --platforms=twitter