package mocks

//go:generate mockgen -destination gen/mockclock/mocks.go -package mockclock github.com/seanbao/can-go/internal/clock Clock,Ticker
//go:generate mockgen -destination gen/mocksocketcan/mocks.go -package mocksocketcan -source ../../pkg/socketcan/fileconn.go
//go:generate mockgen -destination gen/mockcanrunner/mocks.go -package mockcanrunner github.com/seanbao/can-go/pkg/canrunner Node,TransmittedMessage,ReceivedMessage,FrameTransmitter,FrameReceiver
