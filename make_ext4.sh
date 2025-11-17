#!/bin/bash

dd if=/dev/zero of=rootfs.img bs=1M count=2048

mkfs.ext4 rootfs.img

mkdir -p mnt/rootfs_img

sudo mount -o loop rootfs.img mnt/rootfs_img

sudo rsync -aHAX --numeric-ids  \
    --exclude='/proc/*'         \
    --exclude='/sys/*'          \
    --exclude='/dev/*'          \
    --exclude='/run/*'          \
    ./debian-rootfs/ mnt/rootfs_img/

sudo umount mnt/rootfs_img