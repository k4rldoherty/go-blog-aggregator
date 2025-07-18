# Gator 🐊

Gator is a powerful and efficient command-line interface (CLI) tool for aggregating and managing RSS blog feeds. Built with Go, it leverages `sqlc` for type-safe database interactions and `goose` for robust database migrations, ensuring a smooth and reliable experience for your blog aggregation needs.

## Features

* **RSS Feed Aggregation:** Fetch and store content from your favorite RSS feeds.
* **User Management:** Create and manage user accounts for personalized feed subscriptions.
* **Feed Management:** Add, list, and manage the RSS feeds available in the system.
* **Subscription Management:** Users can subscribe to feeds to create a personalized reading list.
* **Post Browse:** Browse posts from all aggregated feeds or from your subscribed feeds.
* **Type-Safe Database Access:** Utilizes `sqlc` to generate Go code from SQL queries, providing compile-time safety and reducing common database errors.
* **Database Migrations:** Employs `goose` for reliable and version-controlled database schema management.

## Prerequisites

Before you can install and run `gator`, ensure you have the following installed on your system:

* **Go (Golang):** Version 1.20 or higher. You can download it from [go.dev/doc/install](https://go.dev/doc/install).
* **PostgreSQL:** A running PostgreSQL database instance. `gator` relies on PostgreSQL for data storage. You can install it via your operating system's package manager, Docker, or download from [postgresql.org/download/](https://www.postgresql.org/download/).

## Installation

### 1. Install `gator` CLI

You can install the `gator` CLI tool directly using `go install`. This command compiles the source code and places the executable in your `$GOPATH/bin` directory, making it globally accessible.

```bash
go install [github.com/k4rldoherty/rss-blog-aggregator](https://www.github.com/k4rldoherty/rss-blog-aggregator)
