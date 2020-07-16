package transactions

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"sync"
	"time"
)

type Transaction struct {
	XMLName string `xml:"transactions"`
	Id      string `json:"id" xml:"id"`
	From    string `json:"from" xml:"from"`
	To      string `json:"to" xml:"to"`
	Amount  int64  `json:"amount" xml:"amount"`
	Created int64  `json:"created" xml:"created"`
}

type Transactions struct {
	XMLName      string         `xml:"transactions"`
	Transactions []*Transaction `xml:"transaction"`
}

type Service struct {
	mu           sync.Mutex
	Transactions []*Transaction
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Register(from, to string, amount int64) (string, error) {
	t := &Transaction{
		Id:      "x",
		From:    from,
		To:      to,
		Amount:  amount,
		Created: time.Now().Unix(),
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.Transactions = append(s.Transactions, t)

	return t.Id, nil
}

func (s *Service) ExportJSON() ([]byte, error) {
	encodedJson, err := json.Marshal(s.Transactions)
	if err != nil {
		log.Print(err)
		return []byte{}, err
	}

	return encodedJson, nil
}

func (t *Transactions) ExportXML() ([]byte, error) {
	encodedXML, err := xml.Marshal(t)
	if err != nil {
		log.Print(err)
		return []byte{}, err
	}
	encodedXML = append([]byte(xml.Header), encodedXML...)

	return encodedXML, nil
}
