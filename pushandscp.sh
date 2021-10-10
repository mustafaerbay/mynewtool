#!/bin/bash
echo "build for linux started"
GOOS=linux GOARCH=386 go build -o mynewtool_linux_86 .
echo "build for windows started"
GOOS=windows GOARCH=386 go build -o mynewtool_windows .

if [[ ${1:-""} == "send" ]]; then
    git config user.name "mustafaerbay"
    git config user.email "mustafaerbay365@gmail.com"
    git add .
    git commit -m "${2}"
    git push
    exit 1
fi

echo "sending to remote machine"
chmod +x mynewtool_*
#scp mynewtool_linux root@10.243.231.38:/root/erbay
#scp mynewtool_linux_86 root@7.189.98.104:/root/