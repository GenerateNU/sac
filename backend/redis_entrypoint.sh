#!/bin/sh

# set up redis configuration directory
mkdir -p /usr/local/etc/redis

# dynamically generate redis configuration and ACL files here, using environment variables
echo "aclfile /usr/local/etc/redis/custom_aclfile.acl" > /usr/local/etc/redis/redis.conf

# generate ACL file using environment variables
if [ -n ${REDIS_USERNAME} ] && [ -n ${REDIS_PASSWORD} ]; then
    echo "user ${REDIS_USERNAME} on allkeys allchannels allcommands >${REDIS_PASSWORD} " > /usr/local/etc/redis/custom_aclfile.acl
fi

# disable default user
if [ $(echo ${REDIS_DISABLE_DEFAULT_USER}) == "true" ]; then
    echo "user default off nopass nocommands" >> /usr/local/etc/redis/custom_aclfile.acl
fi

# call the original Docker entrypoint script with redis-server and the path to the custom redis configuration
exec docker-entrypoint.sh redis-server /usr/local/etc/redis/redis.conf