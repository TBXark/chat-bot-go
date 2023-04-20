read -p "Please input server name: " server_name
scp ./build/bot $server_name:~
ssh $server_name "sh update-chatbot.sh"