# Go Slackbot File Upload

This repository contains a Go-based Slackbot that enables file upload functionality to a specified Slack channel. The Slackbot listens for file upload events and uploads the received files to the designated channel.


## Overview

This repository showcases a powerful and efficient Slackbot built using the Go programming language. The bot is designed to handle file uploads and provide seamless integration with Slack, making it an ideal choice for automating file management tasks within Slack channels.

## Key Features

- **File Upload**: The Slackbot enables users to easily upload files directly to Slack channels, eliminating the need for manual file sharing.

- **Seamless Integration**: The bot seamlessly integrates with Slack, leveraging Slack's APIs to facilitate smooth communication and file handling.

- **Efficiency**: Built using the Go programming language, the bot is highly efficient and lightweight, ensuring optimal performance even with large file uploads.

- **User-Friendly Commands**: The bot provides intuitive and user-friendly commands to interact with the Slack workspace, simplifying file uploads and management for all users.

## Prerequisites

Before running the Slackbot, make sure you have the following prerequisites installed:

- Go programming language (version 1.16+): [Install Go](https://golang.org/doc/install)
- Slack account: [Create a Slack workspace](https://slack.com/get-started)

## Setup

To set up and run the Slackbot, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/Coffeedragon96/06-go-slackbot-file-upload.git
   ```

2. Navigate to the project directory:

   ```bash
   cd 06-go-slackbot-file-upload
   ```

3. Install the project dependencies:

   ```bash
   go mod download
   ```

4. Create a Slack app and bot user in your Slack workspace. Obtain the bot token and channel ID.

   - Visit the [Slack API website](https://api.slack.com/apps) and create a new app.
   - Enable the "Bots" feature and add a new bot user to your app.
   - Install the app to your workspace.
   - Copy the bot token and note the channel ID where the files will be uploaded.

5. Create a `.env` file in the project directory and add the following content:

   ```
   SLACK_BOT_TOKEN=<your-bot-token>
   SLACK_CHANNEL_ID=<your-channel-id>
   ```

   Replace `<your-bot-token>` with the bot token obtained from Slack and `<your-channel-id>` with the channel ID where the files will be uploaded.

6. Run the Slackbot:

   ```bash
   go run main.go
   ```

   The Slackbot will start and connect to your Slack workspace using the provided bot token.

## Usage

Once the Slackbot is running and connected to your Slack workspace, you can upload files to the designated channel using the following steps:

1. In any Slack channel or direct message, click on the file upload icon or use the `/upload` command.

2. Select the file you want to upload from your local machine.

3. The Slackbot will receive the file and automatically upload it to the specified channel.

   > **Note:** Only files uploaded after starting the Slackbot will be processed.

## Contributing

Contributions to this repository are welcome. If you find a bug or have suggestions for improvements, feel free to open an issue or submit a pull request.

Please follow the existing coding style and provide clear commit messages when contributing.

## License

This repository is licensed under the [MIT License](LICENSE). Feel free to use and modify the code according to the terms of the license.

## Acknowledgments

This Slackbot project was inspired by the need to streamline file uploads in Slack conversations. Special thanks to the Slack API documentation and the Go community for providing the necessary tools and resources to build this project.

