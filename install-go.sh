echo "Installing Go...."
rm -f go1.17.linux-amd64.tar.gz
wget https://dl.google.com/go/go1.17.linux-amd64.tar.gz
sudo sh -c 'sudo tar -C /usr/local/ -xzf go1.17.linux-amd64.tar.gz'

echo GOROOT=/usr/local/go
export GOROOT=/usr/local/go
echo PATH=$PATH:$GOROOT/bin
export PATH=$PATH:$GOROOT/bin

go version
echo "Done"
