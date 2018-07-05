Vagrant.configure("2") do |config|
  config.vm.define "node_seed" do |seed|
    seed.vm.box = "ubuntu/trusty64"
    # Setup networking
    seed.vm.network "private_network", ip: "172.28.128.2"
    seed.vm.synced_folder "./", "/gladius/"

    seed.trigger.after :up do |trigger|
      trigger.run_remote = {inline: "/gladius/test_setup_p2p.sh"}
    end
  end
  $i = 0
  $num_of_vms = 10
  while $i < $num_of_vms do
    config.vm.define "node_#$i" do |node|
       node.vm.box = "ubuntu/trusty64"
       # Setup networking
       node.vm.network "private_network", type: "dhcp"
       node.vm.synced_folder "./", "/gladius/"

       node.trigger.after :up do |trigger|
         trigger.run_remote = {inline: "/gladius/test_setup_p2p.sh"}
       end
    end
   $i +=1
  end
end
