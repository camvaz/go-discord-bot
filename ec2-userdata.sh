#!/bin/bash
yum update
yum install -y golang epel-release git screen
cd /usr/local/bin || return
mkdir ffmpeg && cd ffmpeg || return
wget https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-amd64-static.tar.xz
tar -xf ffmpeg-release-amd64-static.tar.xz
mv /usr/local/bin/ffmpeg/ffmpeg-4.4-amd64-static/* . 
ln -s /usr/local/bin/ffmpeg/ffmpeg /usr/bin/ffmpeg
cd /home/ec2-user || return
git clone https://github.com/camvaz/go-discord-bot.git
cd go-discord-bot || return

echo "
 ## keys go here
" > .env

go get
go build
chmod 755 ./go-discord-bot
./go-discord-bot