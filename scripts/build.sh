#!/bin/bash
OUTDIR=build/

go build -o $OUTDIR

GOOS=windows go build -o $OUTDIR
