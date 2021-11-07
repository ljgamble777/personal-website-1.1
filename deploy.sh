#!/bin/bash

GOOS=linux go build -o webserver &&
ssh ec2-184-169-208-236.us-west-1.compute.amazonaws.com 'sudo systemctl stop webserver && mkdir -p personal-website' &&
rsync -v webserver ec2-184-169-208-236.us-west-1.compute.amazonaws.com:/home/ec2-user/personal-website/ &&
ssh ec2-184-169-208-236.us-west-1.compute.amazonaws.com 'sudo systemctl start webserver' &&
rm webserver
