package models

type LatestTlm struct {
	LatestTlmTime    string
	LatestTlmPackets []TlmPacket
}

type TlmPacket struct {
	PacketInfo  PacketInfo
	Telemetries []Tlm
}

type PacketInfo struct {
	TlmApid        string
	Id             string
	CompoName      string
	Name           string
	IsRealtimeData bool
	IsRestricted   bool
}

type Tlm struct {
	TlmInfo  TlmInfo
	TlmValue TlmValue
}

type TlmInfo struct {
	Name        string
	Type        string
	Unit        string
	OctetPos    int
	BitPos      int
	BitLen      int
	ConvType    string
	Poly        []float64
	Status      map[string]string
	Description string
}

type TlmValue struct {
	Time     string
	TI       uint32
	Value    string
	RawValue string
}

type TlmPacketHistory struct {
	PacketInfo   PacketInfo
	TlmHistories []TlmHistory
}

type TlmHistory struct {
	TlmInfo   TlmInfo
	TlmValues []TlmValue
}
