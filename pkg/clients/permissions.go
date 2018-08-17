package clients

import (
	"context"
	"fmt"
	"net/url"

	"time"

	"git.containerum.net/ch/resource-service/pkg/util/coblog"
	"github.com/containerum/cherry"
	"github.com/containerum/cherry/adaptors/cherrylog"
	kubtypes "github.com/containerum/kube-client/pkg/model"
	"github.com/containerum/kube-importer/pkg/kierrors"
	"github.com/go-resty/resty"
	"github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

// Permissions is an interface to permissions service
type Permissions interface {
	ImportNamespaces(ctx context.Context, ns kubtypes.NamespacesList) (*kubtypes.ImportResponse, error)
}

type permc struct {
	client *resty.Client
	log    *cherrylog.LogrusAdapter
}

// NewPermissionsHTTP creates http client to permissions service.
func NewPermissionsHTTP(u *url.URL) Permissions {
	log := logrus.WithField("component", "permissions_client")
	client := resty.New().
		SetHostURL(u.String()).
		SetLogger(log.WriterLevel(logrus.DebugLevel)).
		SetDebug(true).
		SetError(cherry.Err{}).
		SetTimeout(10*time.Second).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("X-User-Role", "admin").
		SetHeader("X-User-Id", "00000000-0000-0000-0000-000000000000")
	client.JSONMarshal = jsoniter.Marshal
	client.JSONUnmarshal = jsoniter.Unmarshal
	return permc{
		client: client,
		log:    cherrylog.NewLogrusAdapter(log),
	}
}

func (perm permc) ImportNamespaces(ctx context.Context, ns kubtypes.NamespacesList) (*kubtypes.ImportResponse, error) {
	perm.log.Debugln("import namespaces")
	coblog.Std.Struct(ns)

	var ret kubtypes.ImportResponse
	resp, err := perm.client.R().
		SetBody(ns).
		SetContext(ctx).
		SetResult(&ret).
		Post("/import/namespaces")
	if err != nil {
		return nil, kierrors.ErrInternalError().Log(err, perm.log)
	}
	if resp.Error() != nil {
		return nil, resp.Error().(*cherry.Err)
	}
	return &ret, nil
}

func (perm permc) String() string {
	return fmt.Sprintf("permissions http client: url=%v", perm.client.HostURL)
}

// Dummy implementation

type permcDummy struct {
	log *logrus.Entry
}

// NewDummyResource creates a dummy client to permissions service. It does nothing but logs actions.
func NewDummyPermissions() Permissions {
	return permcDummy{log: logrus.WithField("component", "permissions_client_dummy")}
}

func (perm permcDummy) ImportNamespaces(ctx context.Context, ns kubtypes.NamespacesList) (*kubtypes.ImportResponse, error) {
	perm.log.Debugln("import namespaces")
	coblog.Std.Struct(ns)

	return nil, nil
}

func (permcDummy) String() string {
	return "permissions dummy client"
}
