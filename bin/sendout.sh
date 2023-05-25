#!/usr/bin/env zsh
    echo "sending exutables into remote sites"
    go build ../server/
    go build ../client/
    scp server client gaify@sp1:~/github/FHS_PoSpace_IoT/
    scp server client gaify@sp2:~/github/FHS_PoSpace_IoT/
    scp server client gaify@sp3:~/github/FHS_PoSpace_IoT/
    scp server client gaify@sp4:~/github/FHS_PoSpace_IoT/
    scp server client gaify@sp5:~/github/FHS_PoSpace_IoT/
    echo "Servers are already started in this folder."
