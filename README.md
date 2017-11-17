#
##Requirements

* 

##Docker

```bash
docker run -p 6060:8080 \
-e REDIS_URL='redis://10.0.1.10:6379' \
roam-server
```