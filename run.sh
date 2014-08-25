#!/bin/bash

echo "Starting Pocket Backend"

goapp serve app/dispatch.yaml app/default/default.yaml app/manager/manager.yaml


