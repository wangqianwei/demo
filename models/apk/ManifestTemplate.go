package apk

import "encoding/xml"

type Manifest struct {
	manifest                         xml.Name          `xml:"manifest"`
	Xmlns                            string            `xml:"xmlns:android,attr,omitempty"`
	AndroidCompileSdkVersion         int32             `xml:"http://schemas.android.com/apk/res/android compileSdkVersion,attr,omitempty"`
	AndroidCompileSdkVersionCodename int32             `xml:"http://schemas.android.com/apk/res/android compileSdkVersionCodename,attr,omitempty"`
	Package                          string            `xml:"package,attr,omitempty"`
	PlatformBuildVersionCode         int32             `xml:"platformBuildVersionCode,attr,omitempty"`
	PlatformBuildVersionName         int32             `xml:"platformBuildVersionName,attr,omitempty"`
	UsesPermissionList               []UsesPermission  `xml:"uses-permission"`
	UsesFeatureList                  []UsesFeature     `xml:"uses-feature"`
	PermissionList                   []Permission      `xml:"permission"`
	SupportsScreensList              []SupportsScreens `xml:"supports-screens"`
	Application                      Application       `xml:"application"`
}

type UsesPermission struct {
	usesPermission xml.Name `xml:"uses-permission"`
	AndroidName    string   `xml:"http://schemas.android.com/apk/res/android name,attr,omitempty"`
}

type UsesFeature struct {
	usesFeature xml.Name `xml:"uses-feature"`
	AndroidName string   `xml:"http://schemas.android.com/apk/res/android name,attr,omitempty"`
}

type Permission struct {
	permission      xml.Name `xml:"permission"`
	AndroidName     string   `xml:"http://schemas.android.com/apk/res/android name,attr,omitempty"`
	ProtectionLevel string   `xml:"http://schemas.android.com/apk/res/android protectionLevel,attr,omitempty"`
}

type SupportsScreens struct {
	supportsScreens xml.Name `xml:"supports-screens"`
	AnyDensity      string   `xml:"http://schemas.android.com/apk/res/android anyDensity,attr,omitempty"`
	LargeScreens    string   `xml:"http://schemas.android.com/apk/res/android largeScreens,attr,omitempty"`
	NormalScreens   string   `xml:"http://schemas.android.com/apk/res/android normalScreens,attr,omitempty"`
	Resizeable      string   `xml:"http://schemas.android.com/apk/res/android resizeable,attr,omitempty"`
	SmallScreens    string   `xml:"http://schemas.android.com/apk/res/android smallScreens,attr,omitempty"`
}

type Application struct {
	application          xml.Name   `xml:"application"`
	AllowBackup          string     `xml:"http://schemas.android.com/apk/res/android allowBackup,attr,omitempty"`
	HardwareAccelerated  string     `xml:"http://schemas.android.com/apk/res/android hardwareAccelerated,attr,omitempty"`
	Icon                 string     `xml:"http://schemas.android.com/apk/res/android icon,attr,omitempty"`
	Label                string     `xml:"http://schemas.android.com/apk/res/android label,attr,omitempty"`
	LargeHeap            string     `xml:"http://schemas.android.com/apk/res/android largeHeap,attr,omitempty"`
	Name                 string     `xml:"http://schemas.android.com/apk/res/android name,attr,omitempty"`
	ResizeableActivity   string     `xml:"http://schemas.android.com/apk/res/android resizeableActivity,attr,omitempty"`
	RoundIcon            string     `xml:"http://schemas.android.com/apk/res/android roundIcon,attr,omitempty"`
	SupportsRtl          string     `xml:"http://schemas.android.com/apk/res/android supportsRtl,attr,omitempty"`
	Theme                string     `xml:"http://schemas.android.com/apk/res/android theme,attr,omitempty"`
	UsesCleartextTraffic string     `xml:"http://schemas.android.com/apk/res/android usesCleartextTraffic,attr,omitempty"`
	MetaData             []MetaData `xml:"meta-data"`
	Activity             []Activity `xml:"activity"`
	Receiver             []Receiver `xml:"receiver"`
	Service              []Service  `xml:"service"`
	Provider             []Provider `xml:"provider"`
}

type MetaData struct {
	metaData xml.Name `xml:"meta-data"`
	Name     string   `xml:"http://schemas.android.com/apk/res/android name,attr,omitempty"`
	Value    string   `xml:"http://schemas.android.com/apk/res/android value,attr,omitempty"`
	Resource string   `xml:"http://schemas.android.com/apk/res/android resource,attr,omitempty"`
}

type Activity struct {
	activity            xml.Name        `xml:"activity"`
	ConfigChanges       string          `xml:"http://schemas.android.com/apk/res/android configChanges,attr,omitempty"`
	Exported            string          `xml:"http://schemas.android.com/apk/res/android exported,attr,omitempty"`
	LaunchMode          string          `xml:"http://schemas.android.com/apk/res/android launchMode,attr,omitempty"`
	Name                string          `xml:"http://schemas.android.com/apk/res/android name,attr,omitempty"`
	ScreenOrientation   string          `xml:"http://schemas.android.com/apk/res/android screenOrientation,attr,omitempty"`
	WindowSoftInputMode string          `xml:"http://schemas.android.com/apk/res/android windowSoftInputMode,attr,omitempty"`
	IntentFilter        [] IntentFilter `xml:"intent-filter"`
}

type Receiver struct {
	receiver     xml.Name        `xml:"receiver"`
	Name         string          `xml:"http://schemas.android.com/apk/res/android name,attr,omitempty"`
	Exported     string          `xml:"http://schemas.android.com/apk/res/android exported,attr,omitempty"`
	Enabled      string          `xml:"http://schemas.android.com/apk/res/android enabled,attr,omitempty"`
	Process      string          `xml:"http://schemas.android.com/apk/res/android process,attr,omitempty"`
	IntentFilter [] IntentFilter `xml:"intent-filter"`
}
type Service struct {
	service      xml.Name        `xml:"service"`
	Name         string          `xml:"http://schemas.android.com/apk/res/android name,attr,omitempty"`
	Exported     string          `xml:"http://schemas.android.com/apk/res/android exported,attr,omitempty"`
	Enabled      string          `xml:"http://schemas.android.com/apk/res/android enabled,attr,omitempty"`
	Process      string          `xml:"http://schemas.android.com/apk/res/android process,attr,omitempty"`
	IntentFilter [] IntentFilter `xml:"intent-filter"`
}

type Provider struct {
	provider            xml.Name        `xml:"provider"`
	Name                string          `xml:"http://schemas.android.com/apk/res/android name,attr,omitempty"`
	Authorities         string          `xml:"http://schemas.android.com/apk/res/android authorities,attr,omitempty"`
	Exported            string          `xml:"http://schemas.android.com/apk/res/android exported,attr,omitempty"`
	GrantUriPermissions string          `xml:"http://schemas.android.com/apk/res/android grantUriPermissions,attr,omitempty"`
	Process             string          `xml:"http://schemas.android.com/apk/res/android process,attr,omitempty"`
	IntentFilter        [] IntentFilter `xml:"intent-filter"`
}

type IntentFilter struct {
	intentfilter xml.Name   `xml:"intent-filter"`
	Priority     string     `xml:"http://schemas.android.com/apk/res/android priority,attr,omitempty"`
	Data         [] Data    `xml:"data"`
	Category     []Category `xml:"category"`
	Action       []Action   `xml:"action"`
}

type Data struct {
	data        xml.Name `xml:"data"`
	Name        string   `xml:"http://schemas.android.com/apk/res/android name,attr,omitempty"`
	MimeType    string   `xml:"http://schemas.android.com/apk/res/android mimeType,attr,omitempty"`
	Scheme      string   `xml:"http://schemas.android.com/apk/res/android scheme,attr,omitempty"`
	Host        string   `xml:"http://schemas.android.com/apk/res/android host,attr,omitempty"`
	PathPattern string   `xml:"http://schemas.android.com/apk/res/android pathPattern,attr,omitempty"`
}

type Category struct {
	category xml.Name `xml:"category"`
	Name     string   `xml:"http://schemas.android.com/apk/res/android name,attr,omitempty"`
}

type Action struct {
	action xml.Name `xml:"action"`
	Name   string   `xml:"http://schemas.android.com/apk/res/android name,attr,omitempty"`
}
