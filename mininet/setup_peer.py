import subprocess
import os
import sys
import requests
import json
from time import sleep


def setup_peer(node_name, wait_time):
    # Start the controld in the background
    subprocess.Popen("/vagrant/build/gladius-controld >> /tmp/controld_%s.out 2>&1" % node_name,
                     env={"GLADIUSBASE": "/gladius"},
                     shell=True)

    # Wait for controld to start
    sleep(10)

    # Create an account
    url = "http://localhost:3001/api/keystore/account/create"
    data = '''{"passphrase":"password"}'''
    response = requests.post(url, data=data).text

    print "account: " + response

    # Sleep for a bit before checking our state
    sleep(float(wait_time))

    # Check our state and write it to a file
    url = "http://localhost:3001/api/p2p/state/"
    state = requests.get(url).text
    with open('/tmp/%s_state.out' % node_name, 'w') as output:
        output.write(state)

if __name__ == "__main__":
    setup_peer(sys.argv[1], sys.argv[2])
