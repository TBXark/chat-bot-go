read -p "Please input server name: " server_name
scp ./bin/chatbot-linux-amd64 $server_name:~
ssh $server_name "sh update-chatbot.sh"