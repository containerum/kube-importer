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

// Resource is an interface to resource service
type Resource interface {
	ImportDeployments(ctx context.Context, deploy kubtypes.DeploymentsList) error

	ImportServices(ctx context.Context, service model.ServiceWithParamList) error

	ImportIngresses(ctx context.Context, ingress kubtypes.IngressesList) error

	ImportConfigMaps(ctx context.Context, cm kubtypes.ConfigMapsList) error
}

type resc struct {
	client *resty.Client
	log    *cherrylog.LogrusAdapter
}

// NewResourceHTTP creates http client to resource service.
func NewResourceHTTP(u *url.URL) Resource {
	log := logrus.WithField("component", "resource_client")
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
	return resc{
		client: client,
		log:    cherrylog.NewLogrusAdapter(log),
	}
}

func (res resc) ImportDeployments(ctx context.Context, deploy kubtypes.DeploymentsList) error {
	res.log.Debugln("import deployments")
	coblog.Std.Struct(deploy)

	resp, err := res.client.R().
		SetBody(deploy).
		SetContext(ctx).
		Post("/import/deployments")
	if err != nil {
		return kierrors.ErrInternalError().Log(err, res.log)
	}
	if resp.Error() != nil {
		return resp.Error().(*cherry.Err)
	}
	return nil
}

func (res resc) ImportServices(ctx context.Context, svc model.ServiceWithParamList) error {
	res.log.Debugln("import services")
	coblog.Std.Struct(svc)

	resp, err := res.client.R().
		SetBody(svc).
		SetContext(ctx).
		Post("/import/services")
	if err != nil {
		return kierrors.ErrInternalError().Log(err, res.log)
	}
	if resp.Error() != nil {
		return resp.Error().(*cherry.Err)
	}
	return nil
}

func (res resc) ImportIngresses(ctx context.Context, ingr kubtypes.IngressesList) error {
	res.log.Debugln("import ingresses")
	coblog.Std.Struct(ingr)

	resp, err := res.client.R().
		SetBody(ingr).
		SetContext(ctx).
		Post("/import/ingresses")
	if err != nil {
		return kierrors.ErrInternalError().Log(err, res.log)
	}
	if resp.Error() != nil {
		return resp.Error().(*cherry.Err)
	}
	return nil
}

func (res resc) ImportConfigMaps(ctx context.Context, cm kubtypes.ConfigMapsList) error {
	res.log.Debugln("import config maps")
	coblog.Std.Struct(cm)

	resp, err := res.client.R().
		SetBody(cm).
		SetContext(ctx).
		Post("/import/configmaps")
	if err != nil {
		return kierrors.ErrInternalError().Log(err, res.log)
	}
	if resp.Error() != nil {
		return resp.Error().(*cherry.Err)
	}
	return nil
}

func (res resc) String() string {
	return fmt.Sprintf("resource service http client: url=%v", res.client.HostURL)
}

// Dummy implementation

type rescDummy struct {
	log *logrus.Entry
}

// NewDummyResource creates a dummy client to resource service. It does nothing but logs actions.
func NewDummyResource() Resource {
	return rescDummy{log: logrus.WithField("component", "resource_client_dummy")}
}

func (res rescDummy) ImportDeployments(ctx context.Context, deploy kubtypes.DeploymentsList) error {
	res.log.Debugln("import deployments")
	coblog.Std.Struct(deploy)

	return nil
}

func (res rescDummy) ImportServices(ctx context.Context, svc model.ServiceWithParamList) error {
	res.log.Debugln("import services")
	coblog.Std.Struct(svc)

	return nil
}

func (res rescDummy) ImportIngresses(ctx context.Context, ingr kubtypes.IngressesList) error {
	res.log.Debugln("import ingresses")
	coblog.Std.Struct(ingr)

	return nil
}

func (res rescDummy) ImportConfigMaps(ctx context.Context, cm kubtypes.ConfigMapsList) error {
	res.log.Debugln("import config maps")
	coblog.Std.Struct(cm)

	return nil
}

func (rescDummy) String() string {
	return "resource service dummy client"
}
