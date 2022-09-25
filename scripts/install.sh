#!/bin/bash
mv appgo.service /etc/systemd/system
chmod +x myapp.sh
mv myapp.sh /usr/local/bin
systemctl daemon-reload
systemctl enable appgo.service
systemctl start appgo.service
