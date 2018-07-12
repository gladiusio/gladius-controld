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
    sleep(1)

    # Create an account
    url = "http://localhost:3001/api/keystore/account/create"
    data = '''{"passphrase":"password"}'''
    response = requests.post(url, data=data).text

    print "account: " + response

    # Sign the initial message
    url = "http://localhost:3001/api/p2p/message/sign"
    # Get our local IP address
    s = subprocess.check_output(
        "ifconfig | grep -Eo 'inet (addr:)?([0-9]*\.){3}[0-9]*' | grep -Eo '([0-9]*\.){3}[0-9]*' | grep -v '127.0.0.1' | head -n 1 | tail -n 2", shell=True).rstrip()
    data = '''{"message": {"node": {"ip_address": "''' + \
        s + '''"}}, "passphrase": "password"}'''

    print data
    singed_message = requests.post(url, data=data)
    singed_message_string = json.dumps(singed_message.json()['response'])

    print "sm: " + singed_message.text

    # Set up our state
    url = "http://localhost:3001/api/p2p/state/push_message"
    data = singed_message_string
    response = requests.post(url, data=data).text
    print "push: " + response

    # Sleep for a bit before checking our state
    sleep(float(wait_time))

    # Check our state and write it to a file
    url = "http://localhost:3001/api/p2p/state/"
    state = requests.get(url).text
    with open('/tmp/%s_state.out' % node_name, 'w') as output:
        output.write(state)


if __name__ == "__main__":
    setup_peer(sys.argv[1], sys.argv[2])
