$SERVER_IP = "84.201.159.214"
$USER = "great_ustug"
$KEY_PATH = "C:\Users\maxim\.ssh\id_ed25519"

scp -i $KEY_PATH ../.././backend/cmd/server/server ${USER}@${SERVER_IP}:/home/${USER}/server

ssh -i $KEY_PATH ${USER}@${SERVER_IP} "chmod +x server exit"