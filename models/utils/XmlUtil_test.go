package utils

import (
	"demo/models/apk"
	"encoding/xml"
	"testing"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func TestParse(t *testing.T) {
	v := apk.Manifest{}
	XmlParse("H:\\aapt2-3.2.0-alpha18-4804415-windows\\abc\\AndroidManifest.xml", &v);
	LogInfo("%d", len(v.Application.Activity))
	LogInfo("%d", len(v.Application.MetaData))
	LogInfo("%d", len(v.Application.Service))
	LogInfo("%d", len(v.Application.Provider))
	LogInfo("%d", len(v.Application.Receiver))

	LogInfo("%d", len(v.UsesPermissionList))
	LogInfo("%d", len(v.UsesFeatureList))
	LogInfo("%d", len(v.PermissionList))
	LogInfo("%d", len(v.SupportsScreensList))
}
