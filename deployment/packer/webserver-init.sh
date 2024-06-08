cd ~
wget -c https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
sudo tar -C /usr/local/ -xzf go1.22.0.linux-amd64.tar.gz
git clone https://github.com/GenerateNU/sac
export PATH=$PATH:/usr/local/go/bin
cd sac/backend
/usr/local/go/bin/go build -o ~/webserver