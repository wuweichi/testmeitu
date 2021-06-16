# -*- coding: utf-8 -*-
"""
monitor classes for config files
"""
import functools
import pyinotify
import sys
import threading

from job import *


class Monitor(object):
    def __init__(self, config_path, job_handle):
        self.config_path = config_path
        self.jobs = job_handle

    # 自动检查配置文件是否有修改
    def auto_compile(self):
        watch_manager = pyinotify.WatchManager()
        notifier = pyinotify.Notifier(watch_manager, default_proc_fun=self.update_jobs)
        watch_manager.add_watch(self.config_path, pyinotify.IN_MODIFY)
        try:
            notifier.loop()
        except pyinotify.NotifierError, err:
            print >> sys.stderr, err

    # 修改计划任务列表后,重新执行计划任务
    def update_jobs(self, event):
        mask_name = event.maskname;
        if mask_name in ['IN_ATTRIB', 'IN_MODIFY']:
            self.jobs.add_job(update_job=True)
        print mask_name