package api

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/yogyrahmawan/logger_service/src/mockspb"
	"github.com/yogyrahmawan/logger_service/src/pb"
	"golang.org/x/net/context"
)

var postMsg *pb.LoggerMessage
var getMsgWithServiceName *pb.GetLoggerRequest
var getMsgWithLevel *pb.GetLoggerRequest
var loggerResponse *pb.LoggerResponse
var loggerResponses *pb.LoggerResponsesMessage

func createLoggerResponses() {
	lgs := []*pb.LoggerMessage{postMsg}
	loggerResponses = &pb.LoggerResponsesMessage{
		LoggerMessages: lgs,
	}
}

func createLoggerResponse() {
	loggerResponse = &pb.LoggerResponse{
		Status: "ok",
	}
}

func createGetMsgWithLevel() {
	getMsgWithServiceName = &pb.GetLoggerRequest{
		Level: "error",
	}
}

func createGetMsgWithServiceName() {
	getMsgWithServiceName = &pb.GetLoggerRequest{
		ServiceName: "bbyb_service",
	}
}

func createLogPostMessage() {
	cvtTimestamp, err := ptypes.TimestampProto(time.Now())
	if err != nil {
		log.Fatal("error converting timestamp, err = " + err.Error())
		return
	}
	postMsg = &pb.LoggerMessage{
		IpPort:      "localhost:8080",
		ServiceName: "bbyb_service",
		Level:       "error",
		Text:        "error construct iso request",
		CreatedAt:   cvtTimestamp,
	}
}

func TestMain(m *testing.M) {
	createLoggerResponses()
	createLoggerResponse()
	createGetMsgWithLevel()
	createGetMsgWithServiceName()
	createLogPostMessage()

	m.Run()
}

func TestSendLog(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// create mock
	stream := mockspb.NewMockLoggerService_SendLogClient(ctrl)

	// set expectation on sending
	stream.EXPECT().Send(
		gomock.Any(),
	).Return(nil)

	stream.EXPECT().Send(
		postMsg,
	).Return(nil)

	// set expectation on receiving
	stream.EXPECT().RecvMsg(postMsg).Return(nil)
	stream.EXPECT().CloseAndRecv().Return(loggerResponse, nil)

	// client interface
	client := mockspb.NewMockLoggerServiceClient(ctrl)
	client.EXPECT().SendLog(
		gomock.Any(),
	).Return(stream, nil)

	client.EXPECT().SendLog(
		postMsg,
	).Return(stream, nil)

	if err := testSendLog(client); err != nil {
		t.Fatalf("Test failed: %v", err)
	}
}

func testSendLog(client pb.LoggerServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.SendLog(ctx)
	if err != nil {
		return err
	}

	if err := stream.Send(postMsg); err != nil {
		return err
	}

	got, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	if !proto.Equal(got, loggerResponse) {
		return fmt.Errorf("stream.Recv() = %v, want %v", got, loggerResponses)
	}

	if err := stream.CloseSend(); err != nil {
		return err
	}

	return nil
}

func TestGetLog(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mockspb.NewMockLoggerServiceClient(ctrl)
	client.EXPECT().GetLog(
		gomock.Any(),
		getMsgWithServiceName,
	).Return(loggerResponses, nil)

	client.EXPECT().GetLog(
		gomock.Any(),
		getMsgWithLevel,
	).Return(loggerResponses, nil)

	if err := testGetLog(client); err != nil {
		t.Fatalf("Test Failed: %v", err)
	}
}

func testGetLog(client pb.LoggerServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.GetLog(ctx, getMsgWithLevel)
	if err != nil {
		return err
	}

	if !proto.Equal(res, loggerResponses) {
		return fmt.Errorf("received = %v, want %v", res, loggerResponses)
	}

	res, err = client.GetLog(ctx, getMsgWithServiceName)
	if err != nil {
		return err
	}

	if !proto.Equal(res, loggerResponses) {
		return fmt.Errorf("received = %v, want %v", res, loggerResponses)
	}

	return nil
}
