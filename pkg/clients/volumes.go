package clients

import (
	"context"
	"fmt"
	"net/url"

	"git.containerum.net/ch/kube-api/pkg/model"
	"git.containerum.net/ch/resource-service/pkg/util/coblog"
	"github.com/containerum/cherry"
	"github.com/containerum/cherry/adaptors/cherrylog"
	kubtypes "github.com/containerum/kube-client/pkg/model"
	"github.com/containerum/kube-importer/pkg/kierrors"
	"github.com/go-resty/resty"
	"github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

// Volumes is an interface to resc-api service
type Volumes interface {
	ImportStorages(ctx context.Context, storages model.StorageList) error
	ImportVolumes(ctx context.Context, volumes kubtypes.VolumesList) error
}

type volc struct {
	client *resty.Client
	log    *cherrylog.LogrusAdapter
}

// NewResourceHTTP creates http client to resc-api service.
func NewVolumesHTTP(u *url.URL) Volumes {
	log := logrus.WithField("component", "volumes_client")
	client := resty.New().
		SetHostURL(u.String()).
		SetLogger(log.WriterLevel(logrus.DebugLevel)).
		SetDebug(true).
		SetError(cherry.Err{}).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("X-User-Role", "admin").
		SetHeader("X-User-Id", "00000000-0000-0000-0000-000000000000")
	client.JSONMarshal = jsoniter.Marshal
	client.JSONUnmarshal = jsoniter.Unmarshal
	return volc{
		client: client,
		log:    cherrylog.NewLogrusAdapter(log),
	}
}

func (vol volc) ImportStorages(ctx context.Context, storages model.StorageList) error {
	vol.log.Debugln("import storages")
	coblog.Std.Debugln(storages)

	resp, err := vol.client.R().
		SetBody(storages).
		SetContext(ctx).
		Post("/import/storages")
	if err != nil {
		return kierrors.ErrInternalError().Log(err, vol.log)
	}
	if resp.Error() != nil {
		return resp.Error().(*cherry.Err)
	}
	return nil
}

func (vol volc) ImportVolumes(ctx context.Context, volumes kubtypes.VolumesList) error {
	vol.log.Debugln("import volumes")
	coblog.Std.Struct(volumes)

	resp, err := vol.client.R().
		SetBody(volumes).
		SetContext(ctx).
		Post("/import/volumes")
	if err != nil {
		return kierrors.ErrInternalError().Log(err, vol.log)
	}
	if resp.Error() != nil {
		return resp.Error().(*cherry.Err)
	}
	return nil
}

func (vol volc) String() string {
	return fmt.Sprintf("permissions http client: url=%v", vol.client.HostURL)
}

// Dummy implementation

type volcDummy struct {
	log *logrus.Entry
}

// NewDummyResource creates a dummy client to resc-api service. It does nothing but logs actions.
func NewDummyVolumes() Volumes {
	return volcDummy{log: logrus.WithField("component", "volumes_client_dummy")}
}

func (vol volcDummy) ImportStorages(ctx context.Context, storages model.StorageList) error {
	vol.log.Debugln("import storages")
	coblog.Std.Debugln(storages)

	return nil
}

func (vol volcDummy) ImportVolumes(ctx context.Context, volumes kubtypes.VolumesList) error {
	vol.log.Debugln("import volumes")
	coblog.Std.Struct(volumes)

	return nil
}

func (volcDummy) String() string {
	return "permissions dummy client"
}
