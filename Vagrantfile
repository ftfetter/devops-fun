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
    # config.vm.provision "shell", path: "script/setup/ansible.sh"
    config.vm.provision "file", source: "~/go/src/github.com/ftfetter/devops-fun", destination: "~/go/src/github.com/ftfetter/devops-fun"
    # config.vm.provision "shell", inline: "ansible-playbook /home/vagrant/go/scr/github.com/ftfetter/devops-fun/playbook.yml"
    config.vm.provision "ansible_local" do |ansible|
        ansible.provisioning_path = "/home/vagrant/go/src/github.com/ftfetter/devops-fun"
        ansible.playbook = "playbook.yml"
        ansible.verbose = true
        ansible.install = true
        ansible.groups = {
            "localhost" => ["localhost"]
        }
    end
end