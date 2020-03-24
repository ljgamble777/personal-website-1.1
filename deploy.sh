#!/bin/bash

GOOS=linux go build -o webserver &&
ssh ec2-18-144-24-104.us-west-1.compute.amazonaws.com 'sudo systemctl stop webserver && mkdir -p personal-website' &&
rsync -v webserver ec2-18-144-24-104.us-west-1.compute.amazonaws.com:/home/ec2-user/personal-website/ &&
rsync -rv --exclude=.DS_Store static ec2-18-144-24-104.us-west-1.compute.amazonaws.com:/home/ec2-user/personal-website/ &&
ssh ec2-18-144-24-104.us-west-1.compute.amazonaws.com 'sudo systemctl start webserver' &&
rm webserver