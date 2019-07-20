#!/bin/bash

##############################################
# Script for provisioning Ansible for CentOS #
##############################################

# update repos 
sudo yum update -y

# installing ansible
sudo yum install -y ansible