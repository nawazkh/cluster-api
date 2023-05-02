Changes since v1.4.1
---
## :chart_with_upwards_trend: Overview

- 25 new commits merged
- 1 breaking changes :warning:
- 1 feature additions ✨
- 9 bugs fixed 🐛

## :warning: Breaking Changes

- CAPD: Use k8s v1.27.0 in quickstart docs and CAPD (#8538)

## :sparkles: New Features

- ClusterClass: add support or concurrent MD upgrades in classy clusters (#8528)

## :bug: Bug Fixes

- ClusterClass: avoid errors when MHC and upgrade occur together in classy clusters (#8568)
- Dependency: Update kindnetd and kindest/haproxy (#8470)
- e2e: Catch failure when fetching component URLs (#8489)
- IPAM: fix gateway being required for IPAddress (#8574)
- KCP: Ensure nil-pointer check in KCP syncMachines (#8488)
- MachinePool: Add node watcher to MachinePool controller (#8474)
- Testing: internal/machine-controller fix flakes in phases test (#8590)
- Testing: machine-controller: fix phase tests race condition in tests on lastUpdated field (#8491)

## :seedling: Others

- Dependency: Update controller-tools to v1.11.4 (#8555)
- Dependency: Update cert manager to v1.11.1 (#8532)
- Dependency: Update kind to v1.18.0 (#8434)
- Dependency: Update docker to v20.10.24 (#8476)
- e2e: assert E2E error responses when waiting for MD nodes (#8516)
- e2e: use topology flavor for workload clusters in clusterctl upgrade test (#8550)
- IPAM: Make IPAddressClaim.Status.AddressRef optional (#8530)

:book: Additionally, there have been 8 contributions to our documentation and book. (#8582, #8562, #8537, #8512, #8513, #8482, #8556, #8589) 


_Thanks to all our contributors!_ 😊
