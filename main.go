package main

import (
	"bufio"
	"crypto/sha256"
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type connectionEntry struct {
	params connectionParams
	expiry int64
}

type connectionParams struct {
	netProtocolName     string
	netProtocolNumber   string
	transProtocolName   string
	transProtocolNumber string
	// this spot would be expiry in seconds - column is in seprate structure
	connState string
	fromIP    string
	toIP      string
	fromPort  string
	toPort    string
	destIP    string
	replyIP   string
	destPort  string
	replyPort string
}

func newConnectionParams(netProtocolName, netProtocolNumber, transProtocolName, transProtocolNumber, connState, fromIP, toIP, fromPort, toPort, destIP, replyIP, destPort, replyPort string) *connectionParams {
	return &connectionParams{
		netProtocolName:     netProtocolName,
		netProtocolNumber:   netProtocolNumber,
		transProtocolName:   transProtocolName,
		transProtocolNumber: transProtocolNumber,
		connState:           connState,
		fromIP:              fromIP,
		toIP:                toIP,
		fromPort:            fromPort,
		toPort:              toPort,
		destIP:              destIP,
		replyIP:             replyIP,
		destPort:            destPort,
		replyPort:           replyPort,
	}
}

func asSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func newConnectionEntry(params *connectionParams, expiry int64) *connectionEntry {
	return &connectionEntry{
		params: *params,
		expiry: expiry,
	}
}

func parseList(data string, prevState map[string]*connectionEntry, filter string) map[string]*connectionEntry {
	connections := make(map[string]*connectionEntry)
	i := 0

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		e := parseRow(words, filter)
		if e == nil {
			continue
		}
		connections[asSha256(e.params)] = e
		if prevConnectionEntry, ok := prevState[asSha256(e.params)]; ok {
			// Conn was reused
			if prevConnectionEntry.expiry < e.expiry {
				i++
				now := time.Now()

				fmt.Printf("%s - from: %s:%s - to %s:%s, dest: %s:%s %d-%d\n", now, e.params.fromIP, e.params.fromPort, e.params.toIP, e.params.toPort, e.params.destIP, e.params.destPort, prevConnectionEntry.expiry, e.expiry)
			}
		}
	}
	return connections
}

func removeKey(data string) string {
	if idx := strings.Index(data, "="); idx != -1 {
		return data[idx+1:]
	}
	return data
}

func parseRow(words []string, filter string) *connectionEntry {

	// Skip row if it has incorrect length or UDP
	if len(words) < 15 {
		return nil
	}

	// Discard UDP entries
	if words[2] != "tcp" {
		return nil
	}

	// Parse string into params structure
	params := newConnectionParams(words[0], words[1], words[2], words[3], removeKey(words[5]), removeKey(words[6]), removeKey(words[7]), removeKey(words[8]), removeKey(words[9]), removeKey(words[10]), removeKey(words[11]), removeKey(words[12]), removeKey(words[13]))
	if filter != "" && params.toIP != filter {
		return nil
	}

	expiry, err := strconv.ParseInt(words[4], 10, 64)
	if err != nil {
		panic(err)
	}

	if expiry >= 120 {
		return nil
	}

	return newConnectionEntry(params, expiry)
}

func main() {

	sleepTime := flag.Int("interval", 5, "evaluation interval in seconds")
	filterIP := flag.String("service-ip", "", "IP address of the target service")
	connTrackFile := flag.String("path", "/proc/net/nf_conntrack", "Path to nf_conntrack file")
	flag.Parse()

	// create a map of entries
	prevState := make(map[string]*connectionEntry)

	for {
		file, err := ioutil.ReadFile(*connTrackFile)
		fileContent := string(file)
		if err != nil {
			panic(err)
		}

		prevState = parseList(fileContent, prevState, *filterIP)
		time.Sleep(time.Duration(*sleepTime) * time.Second)
	}

}
