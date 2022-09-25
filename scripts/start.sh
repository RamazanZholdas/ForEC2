#!/bin/bash
systemctl daemon-reload
systemctl enable appgo.service
systemctl start appgo.service