#!/usr/bin/python
# -*- coding: UTF-8 -*-

import MySQLdb
import traceback

class MyDb():
    def __init__(self, config):
        # 打开数据库连接
        self.db = MySQLdb.connect(host=config.MYSQL_HOST, user=config.MYSQL_USER, passwd=config.MYSQL_PASS, db=config.MYSQL_DB, port=config.MYSQL_PORT, charset="utf8" )

    def insert(self, table_name, data):
        try:
            cursor = self.db.cursor()
            fields = ','.join(data.keys())
            data_length = len(data)
            pre_value = ('%s,' * data_length).strip(',')
            sql = "INSERT INTO " + table_name + "(" + fields + ") VALUES (" + pre_value + ")"
            args = data.values()

            cursor.execute(sql, args)
            self.db.commit()

            cursor.close()
            self.db.close()
        except:
            traceback.print_exc()


    # 关闭数据库连接
    def close_db(self):
        self.db.close()
