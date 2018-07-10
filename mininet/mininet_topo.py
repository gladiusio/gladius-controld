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
            self.addLink(host, switch, bw=1000, delay='25ms')


def setupNetwork(num_of_nodes=10):
    topo = SingleSwitchTopo(n=num_of_nodes)
    net = Mininet(topo=topo, link=TCLink)

    net.start()
    # net.pingAll()

    h1 = net.get('h1')
    h1.cmd('/vagrant/mininet/setup_seed.sh >> /tmp/%s.out &' % h1.name)
    seed_ip = h1.IP()

    sleep(10)

    for node_num in range(1, num_of_nodes):
        sleep(.3)
        h = net.get('h%s' % (node_num + 1))
        h.cmd('/vagrant/mininet/setup_peer.sh ' +
              seed_ip + ' >> /tmp/%s.out &' % h.name)

    # Give some time for the nodes to complete their work
    sleep(15)

    responses = set()
    for node in net.hosts:
        filename = "/tmp/final_" + node.name + ".out"
        jsonResponse = node.cmd(
            "curl --request GET --url http://localhost:3001/api/p2p/state/ > " + filename)
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
    setupNetwork(3)

topos = { 'mytopo': ( lambda: SingleSwitchTopo(5) ) }
