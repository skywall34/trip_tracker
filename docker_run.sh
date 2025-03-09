#!/bin/bash

docker run -d \
    --name trip_tracker \
    -p 3000:3000 \
    -v /home/mshin/trip-tracker/internal/database/database.db:/app/internal/database/database.db \
    trip_tracker
