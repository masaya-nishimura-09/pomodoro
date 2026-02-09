#!/bin/bash

gcc main.c -o main `pkg-config --cflags --libs libnotify`
