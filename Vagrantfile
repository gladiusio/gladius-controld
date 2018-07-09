$init = <<SCRIPT
sudo apt-get -y update
sudo apt-get -y install python-pip python-dev git jq
cd /home/vagrant && git clone https://github.com/mininet/mininet.git
cd /home/vagrant && sudo mininet/util/install.sh
sudo mkdir /gladius/
SCRIPT

Vagrant.configure("2") do |config|
  config.vm.provider "virtualbox" do |v|
    v.memory = 4096
    v.cpus = 4
  end
  config.vm.define "mininet" do |v|
    v.vm.box = "ubuntu/xenial64"
    v.vm.hostname = "mininet"
    v.vm.network "private_network", ip: "192.168.33.223"
    v.ssh.forward_agent = true
    v.vm.provision :shell, :inline => $init
    v.trigger.after :up do |trigger|
      trigger.run_remote = {inline: "python /vagrant/mininet/mininet_topo.py"}
    end
  end
end
