package metathings_camerad_storage

import (
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var schemas = `
CREATE TABLE IF NOT EXISTS camera (
    id varchar(255),
    name varchar(255),
    core_id varchar(255),
    entity_name varchar(255),
    owner_id varchar(255),
    state varchar(255),
    url varchar(1024),
    device varchar(255),
    width int,
    height int,
    bitrate int,
    framerate int,

    created_at datetime,
    updated_at datetime
);
`

type storageImpl struct {
	logger log.FieldLogger
	db     *sqlx.DB
}

func (s *storageImpl) CreateCamera(cam Camera) (Camera, error) {
	c := Camera{}

	now := time.Now()
	cam.CreatedAt = now
	cam.UpdatedAt = now

	_, err := s.db.NamedExec(`
INSERT INTO camera (id, name, core_id, entity_name, owner_id, state, url, device, width, height, bitrate, framerate)
VALUES (:id, :name, :core_id, :entity_name, :owner_id, :state, :url, :device, :width, :height, :bitrate, :framerate)`, &cam)
	if err != nil {
		s.logger.WithError(err).Errorf("failed to create camera")
		return c, err
	}

	s.db.Get(&c, "SELECT * FROM camera WHERE id=$1", *cam.Id)

	s.logger.WithField("cam_id", *cam.Id).Debugf("create camera")

	return c, nil
}

func (s *storageImpl) DeleteCamera(cam_id string) error {
	_, err := s.db.Exec("DELETE FROM camera WHERE id=$1", cam_id)
	if err != nil {
		s.logger.WithError(err).WithField("cam_id", cam_id).Errorf("failed to delete camera")
		return err
	}

	s.logger.WithField("cam_id", cam_id).Debugf("delete camera")

	return nil
}
func (s *storageImpl) PatchCamera(cam_id string, cam Camera) (Camera, error) {
	values := []string{}
	arguments := []interface{}{}
	i := 1
	c := Camera{}

	if cam.Name != nil {
		values = append(values, fmt.Sprintf("name=$%v", i))
		arguments = append(arguments, *cam.Name)
		i++
	}

	if cam.State != nil {
		values = append(values, fmt.Sprintf("state=$%v", i))
		arguments = append(arguments, *cam.State)
		i++
	}

	if cam.Url != nil {
		values = append(values, fmt.Sprintf("url=$%v", i))
		arguments = append(arguments, *cam.Url)
		i++
	}

	if cam.Device != nil {
		values = append(values, fmt.Sprintf("device=$%v", i))
		arguments = append(arguments, *cam.Device)
		i++
	}

	if cam.Width != nil && cam.Height != nil {
		values = append(values, fmt.Sprintf("width=$%v", i))
		arguments = append(arguments, *cam.Width)
		i++

		values = append(values, fmt.Sprintf("height=$%v", i))
		arguments = append(arguments, *cam.Height)
		i++
	}

	if cam.Bitrate != nil {
		values = append(values, fmt.Sprintf("bitrate=$%v", i))
		arguments = append(arguments, *cam.Bitrate)
		i++
	}

	if cam.Framerate != nil {
		values = append(values, fmt.Sprintf("framerate=$%v", i))
		arguments = append(arguments, *cam.Framerate)
		i++
	}

	if len(values) > 0 {
		values = append(values, fmt.Sprintf("updated_at=$%v", i))
		arguments = append(arguments, time.Now())
		i++

		val := strings.Join(values, ", ")
		arguments = append(arguments, cam_id)

		sql_str := "UPDATE camera SET " + val + fmt.Sprintf(" WHRE id=$%v", i)
		s.logger.WithFields(log.Fields{
			"sql":  sql_str,
			"args": arguments,
		}).Debugf("execute sql")
		_, err := s.db.Exec(sql_str, arguments...)
		if err != nil {
			s.logger.WithError(err).WithField("cam_id", cam_id).Errorf("failed to patch camera")
			return c, err
		}
		s.db.Get(&c, "SELECT * FROM camera WHERE id=$%1", cam_id)
		s.logger.WithField("cam_id", cam_id).Debugf("update camera")
		return c, nil
	}
	return c, ErrNothingChanged
}
func (s *storageImpl) GetCamera(cam_id string) (Camera, error) {
	c := Camera{}
	err := s.db.Get(&c, "SELECT * FROM camera WHERE id=$1", cam_id)
	if err != nil {
		s.logger.WithError(err).WithField("cam_id", cam_id).Debugf("failed to get camera")
		return c, err
	}

	s.logger.WithField("cam_id", cam_id).Debugf("get camera")
	return c, nil
}
func (s *storageImpl) ListCameras(cam Camera) ([]Camera, error) {
	cs, err := s.list_cameras(cam)
	if err != nil {
		s.logger.WithError(err).Errorf("failed to list cameras")
		return nil, err
	}
	s.logger.Debugf("list cameras")
	return cs, nil
}
func (s *storageImpl) ListCamerasForUser(owner_id string, cam Camera) ([]Camera, error) {
	cam.OwnerId = &owner_id
	cs, err := s.list_cameras(cam)
	if err != nil {
		s.logger.WithField("owner_id", owner_id).WithError(err).Errorf("failed to list cameras for user")
		return nil, err
	}
	s.logger.WithField("owner_id", owner_id).Debugf("list cameras for user")
	return cs, nil
}

func (s *storageImpl) list_cameras(cam Camera) ([]Camera, error) {
	var err error
	values := []string{}
	arguments := []interface{}{}
	i := 0
	cams := []Camera{}

	if cam.Name != nil {
		values = append(values, fmt.Sprintf("name=$%v", i))
		arguments = append(arguments, *cam.Name)
		i++
	}

	if cam.State != nil {
		values = append(values, fmt.Sprintf("state=$%v", i))
		arguments = append(arguments, *cam.State)
		i++
	}

	if cam.Url != nil {
		values = append(values, fmt.Sprintf("url=$%v", i))
		arguments = append(arguments, *cam.Url)
		i++
	}

	if cam.Device != nil {
		values = append(values, fmt.Sprintf("device=$%v", i))
		arguments = append(arguments, *cam.Device)
		i++
	}

	if cam.Width != nil && cam.Height != nil {
		values = append(values, fmt.Sprintf("width=$%v", i))
		arguments = append(arguments, *cam.Width)
		i++

		values = append(values, fmt.Sprintf("height=$%v", i))
		arguments = append(arguments, *cam.Height)
		i++
	}

	if cam.Bitrate != nil {
		values = append(values, fmt.Sprintf("bitrate=$%v", i))
		arguments = append(arguments, *cam.Bitrate)
		i++
	}

	if cam.Framerate != nil {
		values = append(values, fmt.Sprintf("framerate=$%v", i))
		arguments = append(arguments, *cam.Framerate)
		i++
	}

	if len(values) == 0 {
		err = s.db.Select(&cams, "SELECT * FROM camera")
	} else {
		val := strings.Join(values, " and ")
		sql_str := fmt.Sprintf("SELECT * FROM camera WHERE %v", val)
		s.logger.WithFields(log.Fields{
			"sql":  sql_str,
			"args": arguments,
		}).Debugf("execute sql")
		err = s.db.Select(&cams, sql_str, arguments...)
	}
	if err != nil {
		return nil, err
	}

	return cams, nil
}

func newStorageImpl(driver, uri string, logger log.FieldLogger) (*storageImpl, error) {
	if driver != "sqlite3" {
		logger.WithField("driver", driver).Errorf("not supported driver")
		return nil, ErrUnknownStorageDriver
	}
	db, err := sqlx.Connect(driver, uri)
	if err != nil {
		logger.WithFields(log.Fields{
			"driver": driver,
			"uri":    uri,
		}).WithError(err).Errorf("failed to connect database")
	}
	db.MustExec(schemas)
	return &storageImpl{
		logger: logger.WithField("#module", "storage"),
		db:     db,
	}, nil
}
