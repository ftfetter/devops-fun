# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
    config.vm.provider "virtualbox"
    config.vm.provider "virtualbox" do |v|
        v.memory = 1024
        v.cpus = 2
    end
    config.vm.box = "centos/7"
    config.vm.network "private_network", ip: "55.55.55.150"
    config.vm.synced_folder ".", "/home/vagrant/shared/"
    config.vm.provision "shell", path: "script/setup/ansible.sh"
end