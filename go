j=1
while true
do
addr=`cat /wallet/$j |grep Address|awk '{print $3 }'`
echo $addr
for ((i=1;i<=30;i++)) do
cp /etc/systemd/system/sub.service /etc/systemd/system/sub$i.service 
sed -i 's#btfsl#'$i'#g' /etc/systemd/system/sub$i.service 
sed -i 's#tihuan#'$i'#g' /etc/systemd/system/sub$i.service 
sed -i 's#wallet#'$addr'#g' /etc/systemd/system/sub$i.service 
systemctl daemon-reload
service sub$i restart
sleep 30s
echo "第'$i'个完成"
done
sleep 1440m
echo "'$j'lun"
j=`expr $j + 1`
done
