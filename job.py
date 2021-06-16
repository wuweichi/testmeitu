# -*- coding: utf-8 -*-
"""
job classes for crontab
"""
import os
import utils
import subprocess
import datetime
import threading
import mydb
from runCommand import RunCommand


from apscheduler.schedulers.background import BackgroundScheduler
from socketServers import *

class Job(object):
    def __init__(self, config):
        self.scheduler = BackgroundScheduler()
        self.config = config

    def parse_time(self):
        """cron time syntax::
            # * * * * *  command to execute
            # ┬ ┬ ┬ ┬ ┬
            # │ │ │ │ │
            # │ │ │ │ │
            # │ │ │ │ └─── day of week (0 - 7) (0 to 6 are Sunday to Saturday)
            # │ │ │ └───── month (1 - 12)
            # │ │ └─────── day of month (1 - 31)
            # │ └───────── hour (0 - 23)
            # └─────────── minute (0 - 59)
        """

    # 添加任务
    def add_job(self, update_job = False):
        self.scheduler.print_jobs()

        ip = utils.get_ip_address()
        cmd_lines = utils.read_config_file()
        cmd_list = cmd_lines.split("\n")
        for cmd in cmd_list:
            cmd_detail = cmd.split(' ', 8)
            if len(cmd_detail) > 7:
                status = cmd_detail[0]
                jobs_id = cmd_detail[1]
                target_ip = cmd_detail[2]
                minutes = cmd_detail[3]
                hours = cmd_detail[4]
                days = cmd_detail[5]
                months = cmd_detail[6]
                weeks = cmd_detail[7]
                command = cmd_detail[8]
                job = self.scheduler.get_job(job_id=jobs_id)
                if status != '1':
                    if job is not None:
                        self.scheduler.remove_job(job_id=jobs_id)
                    continue
                if target_ip != ip:
                    if job is not None:
                        self.scheduler.remove_job(job_id=jobs_id)
                    continue
                if job is not None:
                    self.scheduler.modify_job(job_id=jobs_id, args=[command, jobs_id])

                    self.scheduler.reschedule_job(job_id=jobs_id, trigger='cron',
                                       minute=minutes, hour=hours, day=days, month=months, day_of_week=weeks)

                    print 'modify:' + jobs_id
                    continue

                self.scheduler.add_job(self.run_cron, args=[command, jobs_id], id=jobs_id, trigger='cron',
                                       minute=minutes, hour=hours, day=days, month=months, day_of_week=weeks)

                print 'add job' + jobs_id
        self.scheduler.print_jobs()
        if update_job is not True:
            self.scheduler.start()

    # 执行命令
    def run_cron(self, command, jobs_id):
        pid_file = '/tmp/crontab_' + jobs_id + '.pid'
        exec_command_thread = threading.Thread(target=RunCommand(self.config, pid_file=pid_file).run_command, name="run_command", args=[command, jobs_id])
        exec_command_thread.setDaemon(True)
        exec_command_thread.start()


