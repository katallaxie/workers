#!/bin/bash

echo "Starting localhost->::1 tunnel..."
socat TCP4-LISTEN:8976,fork TCP6:[::1]:8976 &
tunnel_proc=$!

npx -y wrangler login

kill $tunnel_proc