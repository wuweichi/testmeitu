# -*- coding: utf-8 -*-
"""
run command classes
"""
import os
import datetime
import subprocess
import mydb
import utils
import atexit
import threading


class RunCommand:
    def __init__(self, config, pid_file):
        self.config = config
        self.pid_file = pid_file
        atexit.register(self.delete_pid)

    # 执行命令
    def run_command(self, command, jobs_id):
        # 检查任务是否执行中
        try:
            pf = file(self.pid_file, 'r')
            pid = int(pf.read().strip())
            pf.close()
        except IOError:
            pid = None

        if pid is None:
            # 写pid文件
            file(self.pid_file, 'w+').write("%s\n" % jobs_id)

            result = dict()
            result['job_id'] = jobs_id
            result['start_at'] = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")
            result['log_detail'] = \
            subprocess.Popen(command, shell=True, stderr=subprocess.PIPE, stdout=subprocess.PIPE).communicate()[0]
            result['end_at'] = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")
            result['created_at'] = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")
            result['running_at'] = utils.get_ip_address()
            # 执行结果入库
            try:
                my_db = mydb.MyDb(self.config)
                my_db.insert(table_name='execution_log', data=result)
            except:
                pass
            self.delete_pid()
            # 退出子线程
        return True

    def delete_pid(self):
        os.remove(self.pid_file)
