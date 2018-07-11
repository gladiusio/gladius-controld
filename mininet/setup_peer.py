import subprocess
import os
import sys
import requests
import json
from time import sleep


def setup_peer(node_name, discovery_ip):
    # Start the controld in the background
    os.environ["GLADIUS_BASE"] = "/gladius"
    subprocess.Popen("/vagrant/build/gladius-controld >> /tmp/controld_%s" % node_name,
                     env={"GLAIUD_BASE": "/gladius"},
                     shell=True)

    sleep(1)
    # Create an account
    url = "http://localhost:3001/api/keystore/account/create"
    data = '''{"passphrase":"password"}'''
    response = requests.post(url, data=data)

    # Sign the intorduction message
    url = "http://localhost:3001/api/p2p/message/sign"
    s = subprocess.check_output(
        "ifconfig | grep -Eo 'inet (addr:)?([0-9]*\.){3}[0-9]*' | grep -Eo '([0-9]*\.){3}[0-9]*' | grep -v '127.0.0.1' | head -n 1 | tail -n 2", shell=True)
    data = '''{"message": {"node": {"ip_address": "''' + \
        s + '''"}}, "passphrase": "password"}'''
    singed_message = requests.post(url, data=data).json()['response']
    singed_message_string = json.dumps(singed_message)

    # Introduce to the discovery peer
    url = "http://localhost:3001/api/p2p/discovery/introduce"
    data = '''{"ip": "''' + discovery_ip + '''","passphrase": "password","signed_message": ''' + \
        singed_message_string + '''}'''
    requests.post(url, data=data)

    # For good measure inform the peers we just learned about
    url = "http://localhost:3001/api/p2p/state/push_message"
    data = singed_message_string
    requests.post(url, data=data)

    # Sleep for a minute before checking our state
    sleep(20)

    # Check our state and write it to a file
    url = "http://localhost:3001/api/p2p/state/"
    state = requests.get(url).text
    with open('/tmp/%s_state.out' % node_name, 'w') as output:
        output.write(state)


if __name__ == "__main__":
    setup_peer(sys.argv[1], sys.argv[2])
