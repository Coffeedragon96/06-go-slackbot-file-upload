# 06-go-slackbot-file-upload

Building a simple slackbot to upload files on slack

# Prerequisites:
- Slack Channel
- Go installed environment

# Setup Slack channel bot
- Create a slack channel
- Goto api.slack.com/apps
- Create new app
- Enable socket mode
- Create a socket token
- Go to OAuth & Permissions
- Add permissions (Some may be not required, and we can remove them later):
    - channels:read
    - chat:write
    - files:read
    - files:write
    - im:read
    - im:write
    - mpim:history
    - remote_files:read
    - remote_files:share
    - remote_files:write