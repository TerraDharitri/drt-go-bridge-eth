package timer

import (
	"github.com/TerraDharitri/drt-go-chain/ntp"
)

func newNTPTimerWithInnerSyncTimer(ntpSyncTimer ntp.SyncTimer) *ntpTimer {
	return &ntpTimer{
		ntpSyncTimer: ntpSyncTimer,
	}
}
