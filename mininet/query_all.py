#!/usr/bin/python
import sys
import requests


def query_nodes(nodes):
    results = set()
    for node in nodes:
        url = "http://%s:3001/api/p2p/state/" % node
        state = requests.get(url).text
        results.add(state)

    if (len(results) > 1):
        print "Test failed, there were %d results" % len(results)
        print "State was: " + str(results)
    else:
        print "Test passed!"
        print "State was: " + str(results)


if __name__ == '__main__':
    query_nodes(sys.argv[1:])
