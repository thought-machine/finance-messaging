package main

import (
	"io/ioutil"
	"github.com/op/go-logging"
	"os"
	"encoding/xml"
	"github.com/fairxio/finance-messaging/iso20022/pacs"
)

var LOGGER = logging.MustGetLogger("main")

func main() {

	pacsMessage, err := ioutil.ReadFile("./example-message.xml")
	if err != nil {
		LOGGER.Fatalf("Unable to read file:  %v", err)
		os.Exit(1)
	}

	var pacsMessageParsed pacs.Document00800106
	err = xml.Unmarshal(pacsMessage, &pacsMessageParsed)
	if err != nil {
		LOGGER.Fatalf("Unable to parse file:  %v", err)
		os.Exit(1)
	}

	LOGGER.Infof("Interbank Settlement Date:  %v", pacsMessageParsed.Message.GroupHeader.InterbankSettlementDate)




}
