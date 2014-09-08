#!/bin/bash

echo "Deploying Pocket"

goapp deploy -application=czb-goingdutch -oauth app/dispatch.yaml app/default/default.yaml app/manager/manager.yaml
appcfg.py --application=czb-goingdutch --oauth2 update_queues app/
appcfg.py --application=czb-goingdutch --oauth2 update_indexes app/
