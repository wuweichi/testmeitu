# !/usr/bin/env python
# coding=utf-8
import logging
import os
import sys
import getopt
import threading
import utils
import glob
from job import *
from monitor import *

opts, args = getopt.getopt(sys.argv[1:], "e:", "env=")
env = 'test'
for opt, value in opts:
    if opt == '--env':
        env = value
    elif opt == '-e':
        env = value
if env == 'test':
    import config.test_config as config
else:
    import config.online_config as config


# 根据开发环境配置
opts, args = getopt.getopt(sys.argv[1:], "e:", "env=")
env = 'test'
for opt, value in opts:
    if opt == '--env':
        env = value
    elif opt == '-e':
        env = value

if env == 'test':
    import config.test_config as config
elif env == 'local':
    import config.local_config as config
else:
    import config.online_config as config

logging.basicConfig()

if __name__ == '__main__':
    print 'Crontab Server(PID:%s) Started!' % os.getpid()
    try:
        # 删除各个任务pid文件
        print "清空pid文件"
        pid_list = glob.glob('/tmp/crontab_*')
        for pid_file in pid_list:
            os.remove(pid_file)
            
        # 计划任务配置表
        crontab_list = utils.get_config_file()
        
        # 初始化任务
        job_handle = Job(config)
        job_handle.add_job()

        # 监听配置文件改动
        file_monitor_thread = threading.Thread(target=Monitor(crontab_list, job_handle).auto_compile, name="monitor_files")
        file_monitor_thread.start()

        # 监听后台传输
        socket_thread = threading.Thread(target=sock(config).run, name="run_socket")
        socket_thread.start()
    except (KeyboardInterrupt, SystemExit):
        traceback.print_exc()