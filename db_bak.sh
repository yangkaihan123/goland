#!/bin/bash
HOSTNAME="172.30.80.21"
SOCKET="/data/mysql/mysql3306/data/mysql.sock"
PORT="3306"
USERNAME="root"
PASSWORD="123456"
OPTIONS="/usr/local/mysql"
DATE=`date '+%Y%m%d'`  #日期格式（作为文件名）
ARCHIVE=mysql_$DATE.tar.gz #压缩文件名
AGODATE=mysql_`date -d '30 days ago' +%Y%m%d` #前30天
all_databases="show databases;"
BAK_DIR="/home/backups/mysqlbak"
BAK_DATA_DIR="$BAK_DIR/$DATE"  #备份文件存储路径
LOGFILE=$BAK_DIR/data_backup.log #日记文件路径

#判断备份文件存储目录是否存在，否则创建该目录
if [ -d ${BAK_DATA_DIR} ];then
   rm -rf ${BAK_DATA_DIR}
fi
mkdir -p ${BAK_DATA_DIR}

#开始备份之前，将备份信息头写入日记文件
echo " " >> $LOGFILE
echo "———————————————–——————–——————–" >> $LOGFILE
echo "BACKUP DATE:" $(date +"%Y-%m-%d %H:%M:%S") >> $LOGFILE
echo "———————————————–——————–——————– " >> $LOGFILE

# Backup Database:3306
for i in `$OPTIONS/bin/mysql -h$HOSTNAME -u${USERNAME} -p${PASSWORD} --socket=${SOCKET} --port=${PORT} -e "${all_databases}"`
do
  if [ $i != "Database" -a $i != "mysql" -a $i != "information_schema" -a $i != "performance_schema" -a $i != "test" -a $i != "sys" ]; then
     #echo $i
     $OPTIONS/bin/mysqldump -u${USERNAME} -p${PASSWORD} --socket=${SOCKET} -B --default-character-set=utf8 --set-gtid-purged=off --add-drop-database -R $i > ${BAK_DATA_DIR}/${i}_$DATE.sql
  fi
done

echo "备份成功!" >> $LOGFILE

#切换至备份目录
cd $BAK_DATA_DIR
cd ..
#判断数据库备份是否成功
if [ -d ${BAK_DATA_DIR} ]; then
#创建备份文件的压缩包
tar czvf $ARCHIVE $DATE  >> $LOGFILE
echo "———————————————–——————–——————– " >> $LOGFILE
echo "打包完成! $(date +"%Y-%m-%d %H:%M:%S")" >> $LOGFILE
#删除原始备份文件，只需保留数据库备份文件的压缩包
rm -rf $BAK_DATA_DIR
echo "———————————————–——————–——————– " >> $LOGFILE
echo "删除备份原文件成功!" >> $LOGFILE

if [[ -f $AGODATE.tar.gz ]]; then
   #statements
   rm -f $AGODATE.tar.gz
   echo "删除前N天备份文件成功!" >> $LOGFILE
fi
else
   echo "Database Backup Fail!(备份失败)" >> $LOGFILE
fi