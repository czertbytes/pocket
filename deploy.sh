#!/bin/bash

echo "Deploying Pocket"

goapp deploy -oauth app/dispatch.yaml app/default/default.yaml app/manager/manager.yaml
appcfg.py --oauth2 update_queues app/
appcfg.py --oauth2 update_indexes app/
