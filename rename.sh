#!/bin/bash

for FILE in model_*.go
do
    NEWNAME=${FILE:6}
    mv $FILE $NEWNAME
done