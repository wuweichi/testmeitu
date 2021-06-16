# !/usr/bin/env python
# coding=utf-8

import os
import json
import threading

from twisted.protocols import basic
from twisted.internet import protocol
from twisted.protocols.policies import TimeoutMixin
from twisted.internet import reactor

listenPort = 21110


class SvrProtocol(basic.LineReceiver, TimeoutMixin):
    def __init__(self):
        self.setTimeout(60)

    def timeoutConnection(self):
        self.sendLine('Error timeoutConnection')
        self.transport.loseConnection()

    def lineReceived(self, line):
        print line
        exit()
        self.resetTimeout()
        line = line.strip()
        if line == 'quit':
            self.transport.loseConnection()
        else:
            config_file_object = open('cron_config.conf', 'w')
            try:
                decode_json = json.loads(line)
                for line in decode_json:
                    config_string = line['id'] + ' ' + line['target_ip'] + ' ' + line['planning_time'] + ' ' + line['command']
                    print config_string
                    config_file_object.write(config_string+"\n")
            except:
                pass


class SvrServerFactory(protocol.ServerFactory):
    protocol = SvrProtocol


class Monitor:
    def run(self):
        pidFile = '/tmp/testsvr.pid'
        open(pidFile, 'w').write(str(os.getpid()))
        print 'config file watch Server(PID:%s) Started!' % os.getpid()
        reactor.listenTCP(listenPort, SvrServerFactory(), 500)
        reactor.suggestThreadPoolSize(40)
        reactor.run(installSignalHandlers=0)


