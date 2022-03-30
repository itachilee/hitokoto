package util

import (
	"github.com/nullseed/logruseq"
	log "github.com/sirupsen/logrus"
)

func setup() {
	log.SetReportCaller(true)
	log.AddHook(logruseq.NewSeqHook("http://localhost:5341",
		logruseq.OptionAPIKey("vqbAePvtL92GpF1qtFDc")))
}
