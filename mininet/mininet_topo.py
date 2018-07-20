#!/usr/bin/python

from mininet.topo import Topo
from mininet.topolib import TreeTopo
from mininet.net import Mininet
from mininet.node import CPULimitedHost
from mininet.link import TCLink
from mininet.util import dumpNodeConnections
from mininet.log import setLogLevel
from mininet.cli import CLI
from mininet.log import info, warn, output
from time import sleep
import argparse
import sys


class SingleSwitchTopo(Topo):
    "Single switch connected to n hosts."

    def build(self, n=2):
        switch = self.addSwitch('s1')
        query = self.addHost('query_node')
        self.addLink(query, switch, bw=1000, delay='10ms')

        for h in range(n):
            host = self.addHost('h%s' % (h + 1), privateDirs=['/gladius'])
            self.addLink(host, switch, bw=1000, delay='20ms')

def setupNetwork(topology="flat", num_of_hosts=10, depth=3, fanout=2):
    if topology == "flat":
        topo = SingleSwitchTopo(n=num_of_hosts)
   
    net = Mininet(topo=topo, link=TCLink)
    num_of_hosts = len(net.hosts)
    info('Host count: %d\n' % num_of_hosts)

    net.start()

    between_nodes = 5

    # seed node is always 10.0.0.1
    info("Setting up seed node\n")
    h1 = net.get('h1')
    h1.cmd('python /vagrant/mininet/setup_seed.py ' +
           h1.name + ' >> /tmp/' + h1.name + '_log.out 2>&1 &')
    seed_ip = h1.IP()

    sleep(20)

    info("Setting up accounts\n")
    for node_num in range(1, num_of_hosts):
        h = net.get('h%s' % (node_num + 1))
        h.cmd('python /vagrant/mininet/setup_peer.py ' +
              h.name + ' >> /tmp/' + h.name + '_log.out 2>&1 &')

    sleep(25)

    info("Starting peers\n")
    for node_num in range(1, num_of_hosts):
        info("\rStarting node: %d" % node_num)
        h = net.get('h%s' % (node_num + 1))
        h.cmd('python /vagrant/mininet/start_peer.py ' + h.name + ' ' +
              seed_ip + ' >> /tmp/' + h.name + '_log.out 2>&1 &')
        sleep(between_nodes)

    # Wait for the network to reach equalibrium
    sleep(10)

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

    ap = argparse.ArgumentParser()
    ap.add_argument("--topology", action="store", dest="topology", required=True, choices=('flat', 'tree'), help="Network Topology")
    ap.add_argument("--nodes", action="store", dest="nodes", required='flat' in sys.argv, type=int, help="Number of nodes in a flat topology")
    ap.add_argument("--depth", action="store", dest="depth", required='tree' in sys.argv, type=int, help="Depth of tree topology")
    ap.add_argument("--fanout", action="store", dest="fanout", required='tree' in sys.argv, type=int, help="Fanout of tree topology")

    args = ap.parse_args()

    setupNetwork(args.topology, args.nodes, args.depth, args.fanout)
