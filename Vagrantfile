Vagrant.configure("2") do |config|
  $i = 0
  $num_of_vms = 10
  while $i < $num_of_vms do
    config.vm.define "node_#$i" do |node|
       # Use alpine so we can make lots and lots of VMs
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
