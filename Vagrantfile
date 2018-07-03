Vagrant.configure("2") do |config|
  # Use alpine so we can make lots and lots of VMs
  config.vm.box = "generic/alpine36"

  # Setup networking
  config.vm.network "private_network", type: "dhcp"

  config.vm.synced_folder "./", "/gladius/"
end
