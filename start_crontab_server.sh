#!/usr/bin/env bash
host = `ifconfig em2 |grep "inet addr"| cut -f 2 -d ":"|cut -f 1 -d " "`
port = 31111
appName = "main.py"

while true;
do
    process_count = `ps -fe | grep "$appName" | grep -v "grep" | wc -l`
    port_count =`echo ""|telnet $host $port 2>/dev/null|grep "\^]"|wc -l`

    echo "程序进程数：" $process_count "端口是否通：" $port_count

    if [ "$process_count" != "1" -a $port_count -eq 0 ]; then
        echo "重新启动程序:"
        nohup /opt/python2.7/bin/python main.py >> /tmp/task_log  2>&1 &
    fi
    sleep 2
done