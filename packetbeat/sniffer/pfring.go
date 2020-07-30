// +build linux,havepfring

package sniffer

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pfring"
	"github.com/google/gopacket/layers"
)

type pfringHandle struct {
	Ring *pfring.Ring
	ZeroCopy bool
}

func newPfringHandle(device string, snaplen int, promisc bool, zc bool) (*pfringHandle, error) {

	var h pfringHandle
	var err error

	if device == "any" {
		return nil, fmt.Errorf("Pfring sniffing doesn't support 'any' as interface")
	}

	var flags pfring.Flag

	if promisc {
		flags = pfring.FlagPromisc
	}

	h.Ring, err = pfring.NewRing(device, uint32(snaplen), flags)
	h.ZeroCopy = zc
	return &h, err
}

func (h *pfringHandle) PrepareHandle() (_ error) {
	return h.Ring.SetSocketMode(pfring.ReadOnly)
}

func (h *pfringHandle) SetBPFFilter(expr string) (_ error) {
	return h.Ring.SetBPFFilter(expr)
}

func (h *pfringHandle) ReadPacketData() (data []byte, ci gopacket.CaptureInfo, err error) {
	if !h.ZeroCopy {
		return h.Ring.ReadPacketData()
	} else {
		return h.Ring.ZeroCopyReadPacketData()
	}
}

func (h *pfringHandle) SetCluster(cluster_id int, cluster_mode int) (_ error) {
	return h.Ring.SetCluster(cluster_id, pfring.ClusterType(cluster_mode))
}

func (h *pfringHandle) Enable() (_ error) {
	return h.Ring.Enable()
}

func (h *pfringHandle) LinkType() layers.LinkType {
	return layers.LinkTypeEthernet
}

func (h *pfringHandle) Close() {
	h.Ring.Close()
}