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
        query = self.addHost('query_node')
        self.addLink(query, switch, bw=1000, delay='10ms')

        for h in range(n):
            host = self.addHost('h%s' % (h + 1), privateDirs=['/gladius'])
            self.addLink(host, switch, bw=1000, delay='20ms')


def setupNetwork(num_of_nodes=10):
    topo = SingleSwitchTopo(n=num_of_nodes)
    net = Mininet(topo=topo, link=TCLink)

    net.start()
    # net.pingAll()
    between_nodes = 5

    info("Setting up seed node\n")
    h1 = net.get('h1')
    h1.cmd('python /vagrant/mininet/setup_seed.py ' +
           h1.name + ' >> /tmp/' + h1.name + '_log.out 2>&1 &')
    seed_ip = h1.IP()

    sleep(15)

    info("Setting up accounts\n")
    for node_num in range(1, num_of_nodes):
        h = net.get('h%s' % (node_num + 1))
        h.cmd('python /vagrant/mininet/setup_peer.py ' +
              h.name + ' >> /tmp/' + h.name + '_log.out 2>&1 &')

    sleep(25)

    info("Starting peers\n")
    for node_num in range(1, num_of_nodes):
        info("\rStarting node: %d" % node_num)
        h = net.get('h%s' % (node_num + 1))
        h.cmd('python /vagrant/mininet/start_peer.py ' + h.name + ' ' +
              seed_ip + ' >> /tmp/' + h.name + '_log.out 2>&1 &')
        sleep(between_nodes)

    # Give some time for the last few nodes
    sleep(40)

    info("\nRunning query on all nodes\n")
    query_node = net.get('query_node')
    result = query_node.cmd(
        'python /vagrant/mininet/query_all.py ' + ' '.join([host.IP() for host in net.hosts[:len(net.hosts) - 1]]))

    with open('/tmp/final_output.log', 'w') as f:
        f.write(result)

    CLI(net)
    net.stop()


if __name__ == '__main__':
    setLogLevel('info')
    setupNetwork(20)
