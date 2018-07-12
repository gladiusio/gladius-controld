#!/usr/bin/python

from mininet.topo import Topo
from mininet.net import Mininet
from mininet.node import CPULimitedHost
from mininet.link import TCLink
from mininet.util import dumpNodeConnections
from mininet.log import setLogLevel
from mininet.cli import CLI
from mininet.log import info, warn, output
from time import sleep


class SingleSwitchTopo(Topo):
    "Single switch connected to n hosts."

    def build(self, n=2):
        switch = self.addSwitch('s1')
        for h in range(n):
            host = self.addHost('h%s' % (h + 1), privateDirs=['/gladius'])
            self.addLink(host, switch, bw=100, delay='10ms')


def setupNetwork(num_of_nodes=10):
    topo = SingleSwitchTopo(n=num_of_nodes)
    net = Mininet(topo=topo, link=TCLink)

    net.start()
    # net.pingAll()
    between_nodes = 3
    completion_time = (between_nodes * num_of_nodes) + 10

    h1 = net.get('h1')
    h1.cmd('python /vagrant/mininet/setup_seed.py ' +
           h1.name + ' ' + str(completion_time + 15) + ' >> ' + h1.name + '_log.out 2>&1 &')
    seed_ip = h1.IP()

    sleep(15)

    for node_num in range(1, num_of_nodes):
        h = net.get('h%s' % (node_num + 1))
        h.cmd('python /vagrant/mininet/setup_peer.py ' + h.name + ' ' +
              str(completion_time) + ' >> ' + h.name + '_log.out 2>&1 &')

    sleep(5)

    for node_num in range(1, num_of_nodes):
        h = net.get('h%s' % (node_num + 1))
        h.cmd('python /vagrant/mininet/start_peer.py ' + h.name + ' ' +
              seed_ip + ' >> ' + h.name + '_log.out 2>&1 &')
        sleep(between_nodes)

    # Give some time for the nodes to complete their work
    sleep(completion_time + 10)

    responses = set()
    for node in net.hosts:
        filename = "/tmp/%s_state.out" % node
        file = open(filename, "r")

        responses.add(file.read())

    if len(responses) > 1:
        warn("Test failed, there were %d different responses." % len(responses))
    else:
        info("Test passed")
    CLI(net)
    net.stop()


if __name__ == '__main__':
    setLogLevel('info')
    setupNetwork(50)

topos = {'mytopo': (lambda: SingleSwitchTopo(50))}
