基于gost修改，socks5部分使用redis认证

```
redisPassword, err := RedisClient.HGet("socks5",req.Username).Result()
		if err != nil {
			log.Log("[socks5]", " redis error")
		}else{
			if req.Password == redisPassword {
				valid = true
				log.Log("[socks5]", "login successs user: "+req.Username)
			} else {
				log.Log("[socks5]", "login fail user:"+req.Username+" pass:"+req.Password)
			}
		}
```
