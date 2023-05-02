Changes since v1.4.1
---
## :chart_with_upwards_trend: Fun stats
- 23 new commits merged
- 1 breaking changes :warning:
- 1 feature additions ✨
- 7 bugs fixed 🐛

## :memo: Proposals
- docs/proposal: drop broken link to ClusterClass google doc (#8556)

## :warning: Breaking Changes
- Use k8s v1.27.0 in quickstart docs and CAPD (#8538)

## :sparkles: New Features
- add support or concurrent MD upgrades in classy clusters (#8528)

## :bug: Bug Fixes
- ipam: fix gateway being required for IPAddress (#8574)
- avoid errors when MHC and upgrade occur together in classy clusters (#8568)
- machine-controller: fix phase tests race condition in tests on lastUpdated field (#8491)
- Catch failure when fetching component URLs (#8489)
- Ensure nil-pointer check in KCP syncMachines (#8488)
- Add node watcher to MachinePool controller (#8474)
- Update kindnetd and kindest/haproxy (#8470)

## :seedling: Others
- test/e2e: use topology flavor for workload clusters in clusterctl upgrade test (#8550)
- Update controller-tools to v1.11.4 (#8555)
- Update cert manager to v1.11.1 (#8532)
- Make IPAddressClaim.Status.AddressRef optional (#8530)
- assert E2E error responses when waiting for MD nodes (#8516)
- Update kind to v1.18.0 (#8434)
- Update docker to v20.10.24 (#8476)

:book: Additionally, there have been 6 contributions to our documentation and book. (#8582, #8562, #8537, #8512, #8513, #8482) 


_Thanks to all our contributors!_ 😊
