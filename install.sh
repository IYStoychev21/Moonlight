#!/bin/bash

cargo build --release

sudo cp target/release/moonlight /bin
