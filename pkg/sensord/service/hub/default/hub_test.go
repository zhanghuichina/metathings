package default_hub

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"

	log_helper "github.com/nayotta/metathings/pkg/common/log"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
	sensord_pb "github.com/nayotta/metathings/pkg/proto/sensord"
	"github.com/nayotta/metathings/pkg/sensord/service/hub"
)

type defaultHubTestSuite struct {
	suite.Suite
	hub hub.Hub
}

func (suite *defaultHubTestSuite) SetupTest() {
	logger, _ := log_helper.NewLogger("hub", "debug")
	opt := opt_helper.NewOption("logger", logger)
	suite.hub, _ = NewHub(opt)
}

func (suite *defaultHubTestSuite) TestNewSubscriber() {
	sub, err := suite.hub.NewSubscriber(opt_helper.NewOption("sensor_id", "test"))
	suite.Nil(err)
	suite.NotEqual(0, sub.Id())
}

func (suite *defaultHubTestSuite) TestCloseSubscriber() {
	sub, _ := suite.hub.NewSubscriber(opt_helper.NewOption("sensor_id", "test"))
	err := sub.Close()
	suite.Nil(err)
}

func (suite *defaultHubTestSuite) TestNewPublisher() {
	pub, err := suite.hub.NewPublisher(opt_helper.NewOption("sensor_id", "test"))
	suite.Nil(err)
	suite.NotEqual(0, pub.Id())
}

func (suite *defaultHubTestSuite) TestClosePublisher() {
	pub, _ := suite.hub.NewPublisher(opt_helper.NewOption("sensor_id", "test"))
	err := pub.Close()
	suite.Nil(err)
}

func (suite *defaultHubTestSuite) TestPublishEmptyData() {
	pub, _ := suite.hub.NewPublisher(opt_helper.NewOption("sensor_id", "test"))
	dat := &sensord_pb.SensorData{}
	err := pub.Publish(dat)
	suite.Nil(err)
}

func (suite *defaultHubTestSuite) TestSubscribeData() {
	pub, _ := suite.hub.NewPublisher(opt_helper.NewOption("sensor_id", "test"))
	sub, _ := suite.hub.NewSubscriber(opt_helper.NewOption("sensor_id", "test"))
	go func() {
		pub.Publish(&sensord_pb.SensorData{
			SensorId: "test",
		})
	}()
	dat, err := sub.Subscribe()
	suite.Nil(err)
	suite.Equal("test", dat.SensorId)
}

func (suite *defaultHubTestSuite) Test1Pub2Sub() {
	wg := new(sync.WaitGroup)

	pub, _ := suite.hub.NewPublisher(opt_helper.NewOption("sensor_id", "test"))
	sub0, _ := suite.hub.NewSubscriber(opt_helper.NewOption("sensor_id", "test"))
	sub1, _ := suite.hub.NewSubscriber(opt_helper.NewOption("sensor_id", "test"))

	wg.Add(2)
	go func() {
		dat, err := sub0.Subscribe()
		suite.Nil(err)
		suite.Equal("test", dat.SensorId)
		wg.Done()
	}()

	go func() {
		dat, err := sub1.Subscribe()
		suite.Nil(err)
		suite.Equal("test", dat.SensorId)
		wg.Done()
	}()

	pub.Publish(&sensord_pb.SensorData{SensorId: "test"})
	wg.Wait()
}

func TestDefaultHubTestSuite(t *testing.T) {
	suite.Run(t, new(defaultHubTestSuite))
}
