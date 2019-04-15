# AgentSmith
Golang C2 Agent PoC utilizing web and social media paltforms to issue command and control and pasting results to PasteBin or other basic HTTP Endpoints.

## Usage (Basic Setup)
1. Set URL for issuing commands in the GetCmd handler function. (Example: a Github Gist link)
2. Set URL for posting command output in the SendResponse handler function. (Example: A simple HTTP Server URL/IP that logs GET Request output)
3. If using Pastebin instead of step 2, enter a valid API key to the SendtoPB handler function and turn on the flag in the SendResponse handler function.

## Usage (Server)
1. Issue a command using the following syntax:
    - (cmd)some_command(cmd)
2. To issue more complex commands please review the official golang documentation on [exec](https://golang.org/pkg/os/exec/) and use the following syntax to satisfy those paramters:
    - (cmd)ls(cmd)
    - (arg)-la(arg)
    - (val)/etc(val)
    - The above example constructs the command string `ls -la /etc`

## Usage (Agent)
1. Follow the Basic Setup usage to configure your agent
2. Build the Agent for deployment: `make {os}-agent` (Be sure you have Go installed and in your path! See Makefile for types of agents you can build)
3. Deploy to target
