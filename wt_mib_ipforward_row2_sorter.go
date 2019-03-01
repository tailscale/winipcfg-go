/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import "sort"

type wtMibIpforwardRow2ByMetric []*wtMibIpforwardRow2

func (byMetric wtMibIpforwardRow2ByMetric) Len() int {
	return len(byMetric)
}

func (byMetric wtMibIpforwardRow2ByMetric) Swap(i, j int) {
	byMetric[i], byMetric[j] = byMetric[j], byMetric[i]
}

func (byMetric wtMibIpforwardRow2ByMetric) Less(i, j int) bool {
	return byMetric[i].Metric < byMetric[j].Metric
}

func sortWtMibIpforwardRow2sByMetric(rows []*wtMibIpforwardRow2) {
	sort.Sort(wtMibIpforwardRow2ByMetric(rows))
}
