// +build !linux !havepfring

package sniffer

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type pfringHandle struct {
}

func newPfringHandle(device string, snaplen int, promisc bool, zc bool) (*pfringHandle, error) {
	return nil, fmt.Errorf("Pfring sniffing is not compiled in")
}

func (h *pfringHandle) PrepareHandle() (_ error) {
	return fmt.Errorf("Pfring sniffing is not compiled in")
}

func (h *pfringHandle) SetBPFFilter(expr string) (_ error) {
	return fmt.Errorf("Pfring sniffing is not compiled in")
}

func (h *pfringHandle) ReadPacketData() (data []byte, ci gopacket.CaptureInfo, err error) {
	return data, ci, fmt.Errorf("Pfring sniffing is not compiled in")
}

func (h *pfringHandle) SetCluster(cluster_id int, cluster_mode int) (_ error) {
	return fmt.Errorf("Pfring sniffing is not compiled in")
}

func (h *pfringHandle) Enable() (_ error) {
	return fmt.Errorf("Pfring sniffing is not compiled in")
}

func (h *pfringHandle) LinkType() layers.LinkType {
	return fmt.Errorf("Pfring sniffing is not compiled in")
}

func (h *pfringHandle) Close() {
	return fmt.Errorf("Pfring sniffing is not compiled in")
}