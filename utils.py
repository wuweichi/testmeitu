# -*- coding: utf-8 -*-

import socket
import struct
import fcntl
import os

from datetime import datetime


# 获取本机IP
def get_ip_address(ifname='em2'):
    try:
        s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        ip = socket.inet_ntoa(fcntl.ioctl(
            s.fileno(),
            0x8915,
            struct.pack('256s', ifname[:15])
        )[20:24])
    except:
        if ifname != 'eth0':
            ip = get_ip_address('eth0')
    return ip


# 读取配置文件
def read_config_file():
    config_file = get_config_file()
    print(str('Read config file in: %s' % datetime.now()))
    
    config_file_object = open(config_file, 'r')
    try:
        all_config_text = config_file_object.read()
    finally:
        config_file_object.close()
    return all_config_text

# 获取配置文件,不存在则创建
def get_config_file():
    config_file = "%s/manage-task/config/cron_config.conf" % os.environ['HOME']
    file_dir = os.path.split(config_file)[0]
    if not os.path.isdir(file_dir):
        os.makedirs(file_dir)
    if not os.path.exists(config_file ):
        os.system(r'touch %s' % config_file)
    return config_file