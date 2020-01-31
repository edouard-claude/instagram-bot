
# Instagram Bot

Another **[Instagram Bot](https://github.com/edouard-claude/instagram-bot)** to increase your followers in [Golang](https://golang.org/).

This **[Instagram Bot](https://github.com/edouard-claude/instagram-bot)** will browse all the followers of the target `-t`, then it will like 5 `-l` recent publications > 90 days `-d` he skips accounts that are already liked, and and he sleep between 1 and 15 sec like an human.

`go build && ./insta -u <your_username> -p <your_password> -t <target>`

 
### Bot Instagram Output :
```
👥Fetch XXX followers from <target>
👥Append XX publics followers of XXX from <target>
⏩Paginate next followers from <target>
😴Sleep 1sec...
👥Append XX publics followers of XXX from <target>
⏩Paginate next followers from <target>
😴Sleep 4sec...
👥Append XXX publics followers of XXX from <target>
⏩Paginate next followers from <target>
😴Sleep 2sec...
👥Append XXX publics followers of XXX from <target>
⏩Paginate next followers from <target>
[...]
😅Already liked skip 1 of XXX...
🌅Fetch feed from <follower_of_target> 2/XXX
🧓Item are too old skip 2 of XXX...
🌅Fetch feed from <follower_of_target> 3/XXX
🧓Item are too old skip 3 of XXX...
🌅Fetch feed from <follower_of_target> 4/XXX
💚Like 1/2
😴Sleep 3sec...
🌅Fetch feed from <follower_of_target> 5/XXX
💚Like 1/2
😴Sleep 12sec...
🧓Item are too old skip 4 of XXX...
🌅Fetch feed from <follower_of_target> 6/XXX
🧓Item are too old skip 5 of XXX...
🌅Fetch feed from <follower_of_target>._ 7/XXX
😅Already liked skip 6 of XXX...
🌅Fetch feed from <follower_of_target> 8/XXX
[...]
```
Enjoy 😘