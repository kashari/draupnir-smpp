# draupnir-smpp
smpp client handler (spec v3.4) that runs on port 4444 and can be used to make smpp traffic to a SMSC or ESME
written in golang making use of goroutines and channels 

pdu base code sourced from (gosmpp) lib

first version (core functionality making the "middle man" between an ESME and a SMSC)