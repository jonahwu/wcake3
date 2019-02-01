# concept of directory
ETCDCTL_API=3 etcdctl get --prefix '/chan'

ETCDCTL_API=3 etcdctl get --prefix '/chan' --keys-only
# need sub aaa/ and only key; some time we just need key as an index
ETCDCTL_API=3 etcdctl get aaa/ --keys-only --prefix

# global search
ETCDCTL_API=3 etcdctl get --from-key '/chan/p1'
