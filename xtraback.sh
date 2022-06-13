#!/bin/sh
 xtrabackup --defaults-file=/etc/my.cnf --user=xtrabk --password='1234qwer' --stream=tar --parallel=5 --no-timestamp --target-dir /data/tmp/backup_full_mysql | ssh -p 22 root@172.16.163.22 " gzip - > /data/tmp/backup_full_mysql/backup.tar.gz"


 gateway-operation.fat.svc.spaceiot.local:8887


export JAVA_HOME=/usr/lib/jvm/java-1.8.0-openjdk-1.8.0.302.b08-0.el7_9.x86_64
ROCKETMQ_HOME=/root/rocketmq-all-4.7.1-source-release/distribution/target/rocketmq-4.7.1/rocketmq-4.7.1
#nohup sh $ROCKETMQ_HOME/bin/mqnamesrv >> $ROCKETMQ_HOME/namesrv.log 2>&1 &
nohup sh $ROCKETMQ_HOME/bin/mqbroker -n "127.0.0.1:9876" -c $ROCKETMQ_HOME/conf/2m-noslave/broker-a.properties >> $ROCKETMQ_HOME/broker.log 2>&1 &
