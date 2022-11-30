Vagrant.configure("2") do |config|
  config.vm.box = "debian/bullseye64"

  config.vm.provider "virtualbox" do |v|
    v.name = "Go heap dev"
  end

  config.vm.provision "shell" do |s|
    s.path = "vagrant/provision.sh"
  end
end
