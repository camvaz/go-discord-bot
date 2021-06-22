#!/bin/bash
yum update
yum install -y golang epel-release git screen
cd /usr/local/bin || return
mkdir ffmpeg && cd/ffmpeg
wget https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-amd64-static.tar.xz
tar -xf ffmpeg-release-amd64-static.tar.xz
cp -a /usr/local/bin/ffmpeg/ffmpeg-4.2.1-amd64-static/ . /usr/local/bin/ffmpeg/
ln -s /usr/local/bin/ffmpeg/ffmpeg /usr/bin/ffmpeg

echo "
 ## keys go here
" > .env

cd /home/ec2-user || return
git clone https://github.com/camvaz/go-discord-bot.git
cd go-discord-bot || return
go get
go build
chmod 755 ./go-discord-bot
./go-discord-bot