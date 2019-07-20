# -*- mode: ruby -*-
# vi: set ft=ruby

Vagrant.configure(2) do |config|

    config.vm.provider "virtualbox"
    config.vm.provider "virtualbox" do |v|
        v.memory = 1024
        v.cpus = 2
    end

    config.vm.box = "centos/7"
    config.vm.network "private_network", ip: "55.55.55.5"
    config.vm.network "forwarded_port", guest: 8080, host: 80

    config.vm.synced_folder ".", "/home/vagrant/shared/"

    config.vm.provision "shell", path: "setup.sh"
    config.vm.provision "file", source: "calc-service.go", destination: "/tmp/calc-service.go"
    config.vm.provision "shell", inline: "sudo mv /tmp/calc-service.go /home/vagrant/go/src/calc-service"
    config.vm.provision "file", source: "setup.sh", destination: "$HOME"
    config.vm.provision "shell", path: "runApplication.sh"

end
  