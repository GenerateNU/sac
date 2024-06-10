sudo yum update -y
sudo yum install -y docker
sudo service docker enable
sudo service docker start
sudo usermod -a -G docker ec2-user

# Install the docker compose plugin for all users
sudo mkdir -p /usr/local/lib/docker/cli-plugins

sudo curl -sL https://github.com/docker/compose/releases/latest/download/docker-compose-linux-"$(uname -m)" \
  -o /usr/local/lib/docker/cli-plugins/docker-compose

# Set ownership to root and make executable
test -f /usr/local/lib/docker/cli-plugins/docker-compose \
  && sudo chown root:root /usr/local/lib/docker/cli-plugins/docker-compose
test -f /usr/local/lib/docker/cli-plugins/docker-compose \
  && sudo chmod +x /usr/local/lib/docker/cli-plugins/docker-compose
  