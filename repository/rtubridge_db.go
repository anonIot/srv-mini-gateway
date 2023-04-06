package repository

import (
	"time"

	"github.com/goburrow/modbus"
)

type acRepositoryDB struct {
	Cli *modbus.RTUClientHandler
}

func NewAcRespositoryDB(Client *modbus.RTUClientHandler) acRepositoryDB {

	return acRepositoryDB{Cli: Client}
}

func (r acRepositoryDB) AcAction(slaveID int, bmsId int) (*AcPacketRepository, error) {
	sid := slaveID
	bms := bmsId

	handler := r.Cli
	handler.SlaveId = byte(sid)
	acAddress := (1000 + (bms * 10)) - 1
	client := modbus.NewClient(handler)
	results, err := client.ReadHoldingRegisters(uint16(acAddress), uint16(10))

	if err != nil {
		return nil, err
	}
	now := time.Now()

	acInfo := AcPacketRepository{
		SlaveId:   sid,
		Bms:       bmsId,
		Value1000: results,
		Timer:     now.String(),
	}

	return &acInfo, nil
}
