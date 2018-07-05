Vagrant.configure("2") do |config|
  $seed_ip = "10.0.2.15"
  $num_of_vms = 10

  config.vm.define "node_seed" do |seed|
    seed.vm.box = "ubuntu/trusty64"
    # Setup networking
    seed.vm.network "private_network", ip: "#$seed_ip"
    seed.vm.synced_folder "./", "/gladius/"

    seed.trigger.after :up do |trigger|
      trigger.run_remote = {inline: "/gladius/vagrant_scripts/setup_seed.sh #$seed_ip"}
    end
  end
  $i = 0
  while $i < $num_of_vms do
    config.vm.define "node_#$i" do |node|
       node.vm.box = "ubuntu/trusty64"
       # Setup networking
       node.vm.network "private_network", type: "dhcp"
       node.vm.synced_folder "./", "/gladius/"

       node.trigger.after :up do |trigger|
         trigger.run_remote = {inline: "/gladius/vagrant_scripts/setup_peer.sh "}
       end
    end
   $i +=1
  end
end
