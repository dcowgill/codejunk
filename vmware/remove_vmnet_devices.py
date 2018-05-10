#!/usr/bin/env python

from __future__ import print_function

import os.path
import re
import subprocess

vmware_dir = '/' + os.path.join(
    'Applications',
    'VMware Fusion.app',
    'Contents',
    'Library',
)

config_entries = [
    'DHCP',
    'HOSTONLY_NETMASK',
    'HOSTONLY_SUBNET',
    'NAT',
    'VIRTUAL_ADAPTER',
]

def vmware_exec(program, *args):
    cmd = ['/usr/bin/sudo', os.path.join(vmware_dir, program)]
    cmd += args
    print(' '.join(cmd))
    # Don't use check_call: the vmware command-line tools return
    # non-zero exit codes for some successful outcomes.
    subprocess.call(cmd)

def remove_vmnet_device(n):
    for cfg in config_entries:
        vmware_exec('vmnet-cfgcli', 'vnetcfgremove', 'VNET_%d_%s' % (n, cfg))

def restart_vmnet():
    for flag in ('--configure', '--stop', '--start'):
        vmware_exec('vmnet-cli', flag)

def get_vmnet_ids():
    ifconfig = subprocess.check_output(['/sbin/ifconfig'])
    for line in ifconfig.split('\n'):
        m = re.match(r'^vmnet(\d+):', line)
        if m:
            yield int(m.group(1))

def main():
    for n in get_vmnet_ids():
        remove_vmnet_device(n)
    restart_vmnet()

if __name__ == '__main__':
    main()
