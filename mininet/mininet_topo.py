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
            # Each host gets 50%/n of system CPU
            host = self.addHost('h%s' % (h + 1), privateDirs=['/gladius'])
            # 10 Mbps, 5ms delay, 2% loss, 1000 packet queue
            self.addLink(host, switch, bw=10, delay='5ms', loss=0,
                         max_queue_size=1000, use_htb=True)


def setupNetwork(num_of_nodes=10):
    topo = SingleSwitchTopo(n=num_of_nodes)
    net = Mininet(topo=topo, link=TCLink)

    net.start()

    h1 = net.get('h1')
    h1.cmd('/vagrant/mininet/setup_seed.sh >> /tmp/%s.out &' % h1.name)
    seed_ip = h1.IP()

    for node_num in range(1, num_of_nodes):
        h = net.get('h%s' % (node_num + 1))
        h.cmd('/vagrant/mininet/setup_peer.sh ' +
              seed_ip + ' >> /tmp/%s.out &' % h.name)

    sleep(10)
    h1.cmd("curl --request GET --url http://localhost:3001/api/p2p/state/ >> /tmp/final.out")
    h2 = net.get('h2')
    h2.cmd(
        "curl --request GET --url http://localhost:3001/api/p2p/state/ >> /tmp/final_h2.out")


if __name__ == '__main__':
    setLogLevel('info')
    setupNetwork(10)
