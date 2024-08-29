 sudo nano /etc/redis/redis.conf

 bind -> ke ip WSL (172.23.96.128)
 protected-mode no


 sudo service redis-server restart
 sudo service redis-server start

 redis-cli -h 172.23.96.128